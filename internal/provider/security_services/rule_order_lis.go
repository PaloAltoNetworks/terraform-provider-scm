// Source of truth for the LIS-based rule ordering algorithm used by
// scm_security_rule_order. Embedded by the generator
// (cmd/generate/terraform/convertor_generator.go) and emitted verbatim to
// terraform-provider-scm/internal/provider/security_services/rule_order_lis.go.
// The .go.tmpl extension keeps the file out of the generator's own compile
// unit while still letting //go:embed read it as raw bytes.
//
// This file is deliberately dependency-free (no SCM SDK, no
// terraform-plugin-framework) so that it can be unit tested in isolation and
// reasoned about as pure Go on string slices. The resource wrapper in
// resource_security_rule_order.go translates the returned []Move slice into
// actual SCM /security-rules/{id}:move API calls.
//
// See tech spec: "Security Rule Ordering — Feature Overview and Technical
// Specification" for the algorithm design, examples, and rationale.
package provider

import (
	"fmt"
	"sort"
	"strings"
)

// Move describes a single :move API call.
//
//   - Position "top"    → move Rule to the top of the rulebase.
//   - Position "bottom" → move Rule to the bottom.
//   - Position "after"  → move Rule immediately after DestinationRule.
//   - Position "before" → move Rule immediately before DestinationRule.
type Move struct {
	Rule            string
	Position        string // "top" | "bottom" | "after" | "before"
	DestinationRule string // empty for top/bottom
}

// OrderMode selects the LIS variant.
const (
	OrderModeConsecutive       = "consecutive"
	OrderModeStrictConsecutive = "strictly_consecutive"
)

// PlanMoves computes the ordered list of :move calls needed to bring `current`
// into `desired` order at the specified anchor.
//
//	desired     — the target ordered list (D) of rule names
//	current     — the full current SCM order (C) of all rules in the rulebase
//	destination — one of "null", "top", "bottom", "after", "before"
//	anchor      — external anchor rule name (required for after/before)
//	mode        — "consecutive" (default) or "strictly_consecutive"
//
// Returns the moves in the order they should be issued.
func PlanMoves(desired, current []string, destination, anchor, mode string) ([]Move, error) {
	if len(desired) == 0 {
		return nil, nil
	}

	// Validation.
	switch destination {
	case "null", "top", "bottom":
		if anchor != "" {
			return nil, fmt.Errorf("destination=%q must not be combined with destination_rule", destination)
		}
	case "after", "before":
		if anchor == "" {
			return nil, fmt.Errorf("destination=%q requires destination_rule", destination)
		}
		if containsString(desired, anchor) {
			return nil, fmt.Errorf("destination_rule %q must not appear in ordered_rules", anchor)
		}
		if !containsString(current, anchor) {
			return nil, fmt.Errorf("destination_rule %q not found in current rulebase", anchor)
		}
	default:
		return nil, fmt.Errorf("invalid destination %q (want null|top|bottom|after|before)", destination)
	}
	for _, r := range desired {
		if !containsString(current, r) {
			return nil, fmt.Errorf("rule %q from ordered_rules not found in current rulebase", r)
		}
	}
	// Detect duplicates in desired.
	seen := make(map[string]bool, len(desired))
	for _, r := range desired {
		if seen[r] {
			return nil, fmt.Errorf("duplicate rule %q in ordered_rules", r)
		}
		seen[r] = true
	}

	switch mode {
	case OrderModeConsecutive, "":
		return planConsecutive(desired, current, destination, anchor)
	case OrderModeStrictConsecutive:
		return planStrictConsecutive(desired, current, destination, anchor)
	default:
		return nil, fmt.Errorf("invalid order %q (want %s|%s)", mode, OrderModeConsecutive, OrderModeStrictConsecutive)
	}
}

// -----------------------------------------------------------------------------
// Consecutive mode: LIS on P (projection of C onto D), unmanaged rules invisible.
// -----------------------------------------------------------------------------

func planConsecutive(desired, current []string, destination, anchor string) ([]Move, error) {
	var moves []Move

	// Phase 1 — Seed.
	postSeed, seed := applySeed(desired, current, destination, anchor)
	if seed != nil {
		moves = append(moves, *seed)
	}

	// Phase 2 — LIS on P (target-only projection).
	//
	//   P[i] = position of desired[i] in postSeed
	//
	// The anchored rule (first for top/after, last for bottom/before) is always
	// pinned into the LIS: after the seed it sits at its target slot whether or
	// not a seed move was needed, and the emit phase never moves it.
	pinIdx := pinnedIndex(desired, destination)
	P := make([]int, len(desired))
	for i, r := range desired {
		P[i] = indexOf(postSeed, r)
	}
	lisMask := lisWithPinAndPrefer(P, pinIdx, nil) // no target-preference; targets == all
	inLIS := make(map[string]bool, len(desired))
	for i, r := range desired {
		if lisMask[i] {
			inLIS[r] = true
		}
	}

	// Phase 3 — Emit.
	return append(moves, emitChain(desired, inLIS, destination)...), nil
}

// pinnedIndex returns the index in desired of the rule that the destination
// anchors (first for top/after, last for bottom/before), or -1 for "null".
func pinnedIndex(desired []string, destination string) int {
	switch destination {
	case "top", "after":
		return 0
	case "bottom", "before":
		return len(desired) - 1
	}
	return -1
}

// emitChain walks seq and emits one :move per rule that is not in the LIS,
// chaining each to its list-neighbor.
//
//	forward  for destination in {top, after}    (seq[0] is pinned, never moved)
//	backward for destination in {bottom, before} (seq[last] is pinned)
//	null     pivots on the leftmost LIS member and chains both directions, so
//	         every non-LIS rule is moved regardless of where it sits in seq.
func emitChain(seq []string, inLIS map[string]bool, destination string) []Move {
	var moves []Move
	switch destination {
	case "null":
		pivot := -1
		for i, r := range seq {
			if inLIS[r] {
				pivot = i
				break
			}
		}
		// Pivot exists for len(seq) >= 1: any single element is an LIS of length 1.
		for i := pivot - 1; i >= 0; i-- {
			if inLIS[seq[i]] {
				continue
			}
			moves = append(moves, Move{Rule: seq[i], Position: "before", DestinationRule: seq[i+1]})
		}
		for i := pivot + 1; i < len(seq); i++ {
			if inLIS[seq[i]] {
				continue
			}
			moves = append(moves, Move{Rule: seq[i], Position: "after", DestinationRule: seq[i-1]})
		}
	case "bottom", "before":
		for i := len(seq) - 2; i >= 0; i-- {
			if inLIS[seq[i]] {
				continue
			}
			moves = append(moves, Move{Rule: seq[i], Position: "before", DestinationRule: seq[i+1]})
		}
	default:
		for i := 1; i < len(seq); i++ {
			if inLIS[seq[i]] {
				continue
			}
			moves = append(moves, Move{Rule: seq[i], Position: "after", DestinationRule: seq[i-1]})
		}
	}
	return moves
}

// -----------------------------------------------------------------------------
// Strict Consecutive mode: LIS on P_ext (projection of C onto D_ext).
// Unmanaged rules may be moved (never explicitly pinned in LIS; the only pin
// is the seed).
// -----------------------------------------------------------------------------

func planStrictConsecutive(desired, current []string, destination, anchor string) ([]Move, error) {
	var moves []Move

	// Phase 1 — Seed.
	postSeed, seed := applySeed(desired, current, destination, anchor)
	if seed != nil {
		moves = append(moves, *seed)
	}

	// Phase 2/3 — Build D_ext (prefix + D + suffix), compute LIS on P_ext, emit.
	if destination == "null" {
		// No absolute position is requested, so the block may settle anywhere
		// among the unmanaged rules. Try each candidate gap (those adjacent to
		// current targets, plus both ends) and keep the plan with fewest moves.
		// Already-contiguous, in-order targets therefore yield zero moves.
		var best []Move
		for i, k := range nullGapCandidates(desired, postSeed) {
			dExt := buildDExtAtGap(desired, postSeed, k)
			cand := emitChain(dExt, strictLIS(postSeed, desired, dExt, ""), destination)
			if i == 0 || len(cand) < len(best) {
				best = cand
			}
		}
		return append(moves, best...), nil
	}

	dExt := buildDExt(desired, postSeed, destination, anchor)
	pinRule := ""
	if pi := pinnedIndex(desired, destination); pi >= 0 {
		pinRule = desired[pi]
	}
	inLIS := strictLIS(postSeed, desired, dExt, pinRule)
	return append(moves, emitChain(dExt, inLIS, destination)...), nil
}

// strictLIS computes the LIS of P_ext (the projection of current onto dExt)
// and returns the set of rule names that stay put. pinRule, if non-empty, is
// forced into the LIS. Target rules are preferred over unmanaged in tie-breaks.
func strictLIS(current, desired, dExt []string, pinRule string) map[string]bool {
	// P_ext = for each rule in current (in order), its index in dExt.
	// Skip rules that aren't in dExt (only the anchor rule for after/before).
	var pExt []int
	var pExtRule []string
	dExtIdx := make(map[string]int, len(dExt))
	for i, r := range dExt {
		dExtIdx[r] = i
	}
	for _, r := range current {
		if idx, ok := dExtIdx[r]; ok {
			pExt = append(pExt, idx)
			pExtRule = append(pExtRule, r)
		}
	}

	pinIdx := -1
	if pinRule != "" {
		pinIdx = indexOf(pExtRule, pinRule)
	}

	targetSet := make(map[string]bool, len(desired))
	for _, r := range desired {
		targetSet[r] = true
	}
	prefer := make([]bool, len(pExtRule))
	for i, r := range pExtRule {
		prefer[i] = targetSet[r]
	}
	lisMask := lisWithPinAndPrefer(pExt, pinIdx, prefer)

	inLIS := make(map[string]bool, len(pExtRule))
	for i, r := range pExtRule {
		if lisMask[i] {
			inLIS[r] = true
		}
	}
	return inLIS
}

// nullGapCandidates returns the distinct gap indices k (0..|U|, where U is the
// unmanaged rules in current order) at which the block could be placed:
// the gap just before each current target, plus both ends. Ascending.
func nullGapCandidates(desired, current []string) []int {
	targetSet := make(map[string]bool, len(desired))
	for _, r := range desired {
		targetSet[r] = true
	}
	seen := map[int]bool{}
	var ks []int
	add := func(k int) {
		if !seen[k] {
			seen[k] = true
			ks = append(ks, k)
		}
	}
	add(0)
	unmanaged := 0
	for _, r := range current {
		if targetSet[r] {
			add(unmanaged)
		} else {
			unmanaged++
		}
	}
	add(unmanaged)
	sort.Ints(ks)
	return ks
}

// buildDExtAtGap returns U[:k] + D + U[k:] where U is the unmanaged rules in
// current order.
func buildDExtAtGap(desired, current []string, k int) []string {
	targetSet := make(map[string]bool, len(desired))
	for _, r := range desired {
		targetSet[r] = true
	}
	var u []string
	for _, r := range current {
		if !targetSet[r] {
			u = append(u, r)
		}
	}
	out := make([]string, 0, len(u)+len(desired))
	out = append(out, u[:k]...)
	out = append(out, desired...)
	out = append(out, u[k:]...)
	return out
}

// buildDExt constructs D_ext = prefix + D + suffix per the tech spec table.
// The anchor rule itself (for after/before) is deliberately excluded from D_ext.
func buildDExt(desired, current []string, destination, anchor string) []string {
	targetSet := make(map[string]bool, len(desired))
	for _, r := range desired {
		targetSet[r] = true
	}

	// Walk current in order, split unmanaged rules into prefix vs suffix per rule.
	var prefix, suffix []string
	for _, r := range current {
		if targetSet[r] {
			continue
		}
		if r == anchor {
			// Excluded from D_ext regardless of mode.
			continue
		}
		side := unmanagedSide(r, current, destination, anchor)
		switch side {
		case "prefix":
			prefix = append(prefix, r)
		case "suffix":
			suffix = append(suffix, r)
		}
	}

	out := make([]string, 0, len(prefix)+len(desired)+len(suffix))
	out = append(out, prefix...)
	out = append(out, desired...)
	out = append(out, suffix...)
	return out
}

// unmanagedSide returns "prefix" or "suffix" for an unmanaged rule based on
// destination. The anchor rule itself is handled by the caller (excluded).
func unmanagedSide(rule string, current []string, destination, anchor string) string {
	switch destination {
	case "null", "top":
		return "suffix"
	case "bottom":
		return "prefix"
	case "after", "before":
		idxX := indexOf(current, anchor)
		if indexOf(current, rule) < idxX {
			return "prefix"
		}
		return "suffix"
	}
	return "suffix"
}

// -----------------------------------------------------------------------------
// Seed application.
// -----------------------------------------------------------------------------

// applySeed returns (postSeedCurrent, seedMove). If the seed is a no-op or
// destination is "null", seedMove is nil and postSeedCurrent is unchanged.
func applySeed(desired, current []string, destination, anchor string) ([]string, *Move) {
	if destination == "null" {
		return current, nil
	}
	var pinRule, position, destRule string
	switch destination {
	case "top":
		pinRule, position = desired[0], "top"
	case "bottom":
		pinRule, position = desired[len(desired)-1], "bottom"
	case "after":
		pinRule, position, destRule = desired[0], "after", anchor
	case "before":
		pinRule, position, destRule = desired[len(desired)-1], "before", anchor
	default:
		return current, nil
	}

	// Check whether pinRule is already at its target position (seed no-op).
	curPos := indexOf(current, pinRule)
	switch position {
	case "top":
		if curPos == 0 {
			return current, nil
		}
	case "bottom":
		if curPos == len(current)-1 {
			return current, nil
		}
	case "after":
		if curPos == indexOf(current, anchor)+1 {
			return current, nil
		}
	case "before":
		if curPos == indexOf(current, anchor)-1 {
			return current, nil
		}
	}

	// Simulate the :move.
	var post []string
	switch position {
	case "top":
		post = moveTo(current, pinRule, 0)
	case "bottom":
		post = moveTo(current, pinRule, len(current)-1)
	case "after":
		post = moveRelativeTo(current, pinRule, anchor, +1)
	case "before":
		post = moveRelativeTo(current, pinRule, anchor, -1)
	}
	return post, &Move{Rule: pinRule, Position: position, DestinationRule: destRule}
}

// moveTo removes `rule` from list and inserts it at absolute position `dstPos`
// in the resulting list (0-indexed, clamped to [0, len]).
func moveTo(list []string, rule string, dstPos int) []string {
	reduced := make([]string, 0, len(list))
	for _, r := range list {
		if r != rule {
			reduced = append(reduced, r)
		}
	}
	if dstPos < 0 {
		dstPos = 0
	}
	if dstPos > len(reduced) {
		dstPos = len(reduced)
	}
	out := make([]string, 0, len(list))
	out = append(out, reduced[:dstPos]...)
	out = append(out, rule)
	out = append(out, reduced[dstPos:]...)
	return out
}

// moveRelativeTo removes `rule` and inserts it before (offset=-1) or after
// (offset=+1) `anchor` in the resulting list. Anchor must exist in `list`
// after removing `rule`.
func moveRelativeTo(list []string, rule, anchor string, offset int) []string {
	reduced := make([]string, 0, len(list))
	for _, r := range list {
		if r != rule {
			reduced = append(reduced, r)
		}
	}
	aidx := indexOf(reduced, anchor)
	insertAt := aidx
	if offset > 0 {
		insertAt = aidx + 1
	}
	if insertAt < 0 {
		insertAt = 0
	}
	if insertAt > len(reduced) {
		insertAt = len(reduced)
	}
	out := make([]string, 0, len(list))
	out = append(out, reduced[:insertAt]...)
	out = append(out, rule)
	out = append(out, reduced[insertAt:]...)
	return out
}

// -----------------------------------------------------------------------------
// LIS with optional pin and prefer-target tie-break.
// -----------------------------------------------------------------------------

// lisWithPinAndPrefer returns a boolean mask (same length as arr) marking
// members of the longest strictly increasing subsequence of arr.
//
//	pinIdx  — if >= 0, arr[pinIdx] must be included in the LIS
//	prefer  — if non-nil, secondary tie-break: prefer subsequences with more
//	          "true" entries (used to bias toward target rules in strict mode)
//
// Reconstruction is deterministic: ties are broken (1) by prefer-count, then
// (2) by earliest predecessor index. Runs in O(n^2), which is fine for the
// small rulesets typical of Terraform-managed containers.
func lisWithPinAndPrefer(arr []int, pinIdx int, prefer []bool) []bool {
	n := len(arr)
	mask := make([]bool, n)
	if n == 0 {
		return mask
	}

	if pinIdx >= 0 {
		// Split around the pin.
		// LIS = (best strictly increasing subseq of arr[:pinIdx] with values < arr[pinIdx])
		//     + [pin]
		//     + (best strictly increasing subseq of arr[pinIdx+1:] with values > arr[pinIdx])
		pinVal := arr[pinIdx]
		leftIdx := lisSegment(arr, 0, pinIdx, -1, pinVal, prefer)
		rightIdx := lisSegment(arr, pinIdx+1, n, pinVal, -1, prefer)
		for _, i := range leftIdx {
			mask[i] = true
		}
		mask[pinIdx] = true
		for _, i := range rightIdx {
			mask[i] = true
		}
		return mask
	}

	idxs := lisSegment(arr, 0, n, -1, -1, prefer)
	for _, i := range idxs {
		mask[i] = true
	}
	return mask
}

// lisSegment returns indices (into arr) of the longest strictly increasing
// subsequence of arr[lo:hi] whose values fall strictly in (lowerExcl,
// upperExcl). Use lowerExcl=-1 for no lower bound and upperExcl=-1 for no
// upper bound. Tie-break: prefer higher prefer-count, then earliest predecessor.
func lisSegment(arr []int, lo, hi, lowerExcl, upperExcl int, prefer []bool) []int {
	n := hi - lo
	if n <= 0 {
		return nil
	}
	// Build sub-array of eligible entries with original indices.
	type entry struct {
		origIdx int
		val     int
	}
	elig := make([]entry, 0, n)
	for i := lo; i < hi; i++ {
		v := arr[i]
		if lowerExcl >= 0 && v <= lowerExcl {
			continue
		}
		if upperExcl >= 0 && v >= upperExcl {
			continue
		}
		elig = append(elig, entry{origIdx: i, val: v})
	}
	if len(elig) == 0 {
		return nil
	}

	// DP over the eligible entries.
	m := len(elig)
	length := make([]int, m)
	preferCount := make([]int, m)
	prev := make([]int, m)
	for i := range length {
		length[i] = 1
		prev[i] = -1
		if prefer != nil && prefer[elig[i].origIdx] {
			preferCount[i] = 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			if elig[j].val < elig[i].val {
				candLen := length[j] + 1
				candPref := preferCount[j]
				if prefer != nil && prefer[elig[i].origIdx] {
					candPref++
				}
				if candLen > length[i] ||
					(candLen == length[i] && candPref > preferCount[i]) {
					length[i] = candLen
					preferCount[i] = candPref
					prev[i] = j
				}
			}
		}
	}

	// Find best endpoint.
	best := 0
	for i := 1; i < m; i++ {
		if length[i] > length[best] ||
			(length[i] == length[best] && preferCount[i] > preferCount[best]) {
			best = i
		}
	}

	// Reconstruct.
	var out []int
	for i := best; i != -1; i = prev[i] {
		out = append([]int{elig[i].origIdx}, out...)
	}
	return out
}

// -----------------------------------------------------------------------------
// Utilities.
// -----------------------------------------------------------------------------

func indexOf(list []string, needle string) int {
	for i, s := range list {
		if s == needle {
			return i
		}
	}
	return -1
}

func containsString(list []string, needle string) bool {
	return indexOf(list, needle) >= 0
}

// -----------------------------------------------------------------------------
// Drift observation (pure; shared by Read and post-apply verification).
// -----------------------------------------------------------------------------

// observeCurrentBlock returns the observed "block state" to store in
// `ordered_rules` for drift detection. If it equals `desired`, no drift.
// If it differs, Terraform's plan will show the diff and trigger Update.
//
// Consecutive mode: projection of currentOrder onto desired (detects order
// drift but ignores interleavers — which is the mode's contract).
//
// Strict Consecutive mode: the raw slice of currentOrder covering the
// enforced block window — target rules AND anything sitting between them,
// AND anything sitting in the position-required prefix/suffix (top slot,
// bottom slot, after/before anchor). This detects interleaver drift AND
// block-position drift.
func observeCurrentBlock(currentOrder, desired []string, destination, anchor, mode string) []string {
	desiredSet := make(map[string]bool, len(desired))
	for _, n := range desired {
		desiredSet[n] = true
	}

	// Consecutive: projection onto desired, plus a destination check. The
	// projection alone cannot see an out-of-position block, so when the
	// destination slot is held by a rule that is not the expected target, that
	// rule is added to the observed list (top/after: prepended; bottom/before:
	// appended) so it differs from desired and surfaces as drift.
	if mode != OrderModeStrictConsecutive {
		var out []string
		for _, n := range currentOrder {
			if desiredSet[n] {
				out = append(out, n)
			}
		}
		if len(desired) == 0 {
			return out
		}
		first, last := desired[0], desired[len(desired)-1]
		switch destination {
		case "top":
			if len(currentOrder) > 0 && currentOrder[0] != first && !desiredSet[currentOrder[0]] {
				out = append([]string{currentOrder[0]}, out...)
			}
		case "bottom":
			if n := len(currentOrder); n > 0 && currentOrder[n-1] != last && !desiredSet[currentOrder[n-1]] {
				out = append(out, currentOrder[n-1])
			}
		case "after":
			if aIdx := indexOf(currentOrder, anchor); aIdx >= 0 {
				slot := anchor // anchor is last: nothing follows it
				if aIdx+1 < len(currentOrder) {
					slot = currentOrder[aIdx+1]
				}
				if slot != first && !desiredSet[slot] {
					out = append([]string{slot}, out...)
				}
			}
		case "before":
			if aIdx := indexOf(currentOrder, anchor); aIdx >= 0 {
				slot := anchor // anchor is first: nothing precedes it
				if aIdx > 0 {
					slot = currentOrder[aIdx-1]
				}
				if slot != last && !desiredSet[slot] {
					out = append(out, slot)
				}
			}
		}
		return out
	}

	// Strict Consecutive: return the slice of currentOrder that covers the
	// enforced block window. Any rule in this slice that isn't in desired
	// is an interleaver or position violation → shows up as drift.
	firstIdx, lastIdx := -1, -1
	for i, n := range currentOrder {
		if desiredSet[n] {
			if firstIdx == -1 {
				firstIdx = i
			}
			lastIdx = i
		}
	}
	if firstIdx == -1 {
		// No target rules present. Return empty; Terraform will diff against
		// the non-empty desired and trigger Update.
		return nil
	}
	lo, hi := firstIdx, lastIdx
	switch destination {
	case "top":
		lo = 0
	case "bottom":
		hi = len(currentOrder) - 1
	case "after":
		if aIdx := indexOf(currentOrder, anchor); aIdx >= 0 {
			lo = aIdx + 1
		}
	case "before":
		if aIdx := indexOf(currentOrder, anchor); aIdx >= 0 {
			hi = aIdx - 1
		}
	}
	if lo < 0 {
		lo = 0
	}
	if hi >= len(currentOrder) {
		hi = len(currentOrder) - 1
	}
	if lo > hi {
		return nil
	}
	return append([]string(nil), currentOrder[lo:hi+1]...)
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// -----------------------------------------------------------------------------
// Import ID parsing (pure; used by ImportState).
// -----------------------------------------------------------------------------

// ParseRuleOrderImportID parses an import ID in the same shape as the
// scm_security_rule tfid/import ID (`folder:snippet:device:<id>`, with the two
// unused container slots left empty), where <id> is the rulebase optionally
// followed by the managed rule names:
//
//	<folder>:::<pre|post>[:<rule>,<rule>,...]
//	:<snippet>::<pre|post>[:<rule>,<rule>,...]
//	::<device>:<pre|post>[:<rule>,<rule>,...]
//
// Rule names are given explicitly because a container holds both managed and
// unmanaged rules and the live rulebase alone cannot say which are managed.
// SCM rule names cannot contain ":" or ",", so both are safe separators.
func ParseRuleOrderImportID(id string) (kind, value, rulebase string, rules []string, err error) {
	const usage = "expected <folder>:::<pre|post>[:<rule1>,<rule2>,...] (or :<snippet>::<pre|post>..., ::<device>:<pre|post>...)"
	parts := strings.SplitN(id, ":", 5)
	if len(parts) != 4 && len(parts) != 5 {
		return "", "", "", nil, fmt.Errorf("invalid import ID %q: %s", id, usage)
	}
	folder, snippet, device := parts[0], parts[1], parts[2]
	rulebase = parts[3]

	set := 0
	for _, c := range []struct{ k, v string }{{"folder", folder}, {"snippet", snippet}, {"device", device}} {
		if c.v != "" {
			kind, value = c.k, c.v
			set++
		}
	}
	if set != 1 {
		return "", "", "", nil, fmt.Errorf("invalid import ID %q: exactly one of folder, snippet or device must be set: %s", id, usage)
	}
	if rulebase != "pre" && rulebase != "post" {
		return "", "", "", nil, fmt.Errorf("invalid rulebase %q: want pre|post", rulebase)
	}
	if len(parts) == 5 && parts[4] != "" {
		seen := map[string]bool{}
		for _, name := range strings.Split(parts[4], ",") {
			name = strings.TrimSpace(name)
			if name == "" {
				return "", "", "", nil, fmt.Errorf("empty rule name in import ID %q", id)
			}
			if seen[name] {
				return "", "", "", nil, fmt.Errorf("duplicate rule %q in import ID", name)
			}
			seen[name] = true
			rules = append(rules, name)
		}
	}
	return kind, value, rulebase, rules, nil
}
