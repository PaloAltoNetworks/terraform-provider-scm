// Unit tests for the LIS-based rule ordering algorithm.
// Embedded and emitted alongside rule_order_lis.go via the generator.
package provider

import (
	"reflect"
	"testing"
)

// simulateMoves applies the sequence of moves to a starting state and returns
// the resulting order. Used by tests to verify that the planned moves actually
// achieve the target ordering.
func simulateMoves(start []string, moves []Move) []string {
	state := append([]string(nil), start...)
	for _, m := range moves {
		switch m.Position {
		case "top":
			state = moveTo(state, m.Rule, 0)
		case "bottom":
			state = moveTo(state, m.Rule, len(state)-1)
		case "after":
			state = moveRelativeTo(state, m.Rule, m.DestinationRule, +1)
		case "before":
			state = moveRelativeTo(state, m.Rule, m.DestinationRule, -1)
		}
	}
	return state
}

// projection returns the sub-sequence of state consisting of members of `keep`
// (order preserved from state).
func projection(state, keep []string) []string {
	set := make(map[string]bool, len(keep))
	for _, k := range keep {
		set[k] = true
	}
	var out []string
	for _, s := range state {
		if set[s] {
			out = append(out, s)
		}
	}
	return out
}

// -----------------------------------------------------------------------------
// Consecutive mode — Examples 1–5 from the tech spec.
// -----------------------------------------------------------------------------

func TestConsecutive_Example1_NullNoSeed(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC", "RuleD", "RuleE"}
	current := []string{"RuleC", "RuleA", "RuleB", "RuleE", "RuleD"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 2 {
		t.Fatalf("want 2 moves, got %d: %+v", len(moves), moves)
	}

	final := simulateMoves(current, moves)
	if got := projection(final, desired); !reflect.DeepEqual(got, desired) {
		t.Fatalf("projection want %v, got %v (final=%v)", desired, got, final)
	}
}

func TestConsecutive_Example2_AfterOnlyManagedBlockers(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC", "RuleD"}
	current := []string{"RuleB", "AnchorRule", "RuleD", "RuleA", "RuleY", "RuleC"}

	moves, err := PlanMoves(desired, current, "after", "AnchorRule", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 3 {
		t.Fatalf("want 3 moves (1 seed + 2 chain), got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "RuleA" || moves[0].Position != "after" || moves[0].DestinationRule != "AnchorRule" {
		t.Fatalf("expected seed move :move RuleA after AnchorRule, got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	// Final state per spec: [AnchorRule, RuleA, RuleB, RuleY, RuleC, RuleD].
	want := []string{"AnchorRule", "RuleA", "RuleB", "RuleY", "RuleC", "RuleD"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestConsecutive_Example3_AfterUnmanagedBlocker(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC", "RuleD"}
	current := []string{"RuleB", "AnchorRule", "ThirdPartyRule", "RuleD", "RuleA", "RuleY", "RuleC"}

	moves, err := PlanMoves(desired, current, "after", "AnchorRule", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 3 {
		t.Fatalf("want 3 moves, got %d: %+v", len(moves), moves)
	}

	final := simulateMoves(current, moves)
	want := []string{"AnchorRule", "RuleA", "RuleB", "ThirdPartyRule", "RuleY", "RuleC", "RuleD"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestConsecutive_Example4_Top(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC", "RuleD"}
	current := []string{"ThirdPartyRule", "RuleD", "RuleA", "RuleY", "RuleB", "RuleC"}

	moves, err := PlanMoves(desired, current, "top", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 2 {
		t.Fatalf("want 2 moves (1 seed + 1 chain), got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "RuleA" || moves[0].Position != "top" {
		t.Fatalf("expected seed :move RuleA top, got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	want := []string{"RuleA", "ThirdPartyRule", "RuleY", "RuleB", "RuleC", "RuleD"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestConsecutive_Example5_Bottom(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC", "RuleD"}
	current := []string{"RuleB", "RuleA", "RuleY", "RuleC", "RuleD", "ThirdPartyRule"}

	moves, err := PlanMoves(desired, current, "bottom", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 2 {
		t.Fatalf("want 2 moves (1 seed + 1 chain), got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "RuleD" || moves[0].Position != "bottom" {
		t.Fatalf("expected seed :move RuleD bottom, got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	want := []string{"RuleA", "RuleY", "RuleB", "RuleC", "ThirdPartyRule", "RuleD"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

// -----------------------------------------------------------------------------
// Strict Consecutive mode — Examples 6–10 from the tech spec.
// -----------------------------------------------------------------------------

func TestStrict_Example6_NullConsecutive(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC"}
	current := []string{"RuleA", "u", "RuleB", "v", "RuleC"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 2 {
		t.Fatalf("want 2 moves, got %d: %+v", len(moves), moves)
	}

	final := simulateMoves(current, moves)
	want := []string{"RuleA", "RuleB", "RuleC", "u", "v"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestStrict_Example7_AfterConsecutive(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC"}
	current := []string{"AnchorRule", "u", "RuleC", "v", "RuleB", "w", "RuleA"}

	moves, err := PlanMoves(desired, current, "after", "AnchorRule", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 3 {
		t.Fatalf("want 3 moves, got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "RuleA" || moves[0].Position != "after" {
		t.Fatalf("expected seed :move RuleA after AnchorRule, got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	want := []string{"AnchorRule", "RuleA", "RuleB", "RuleC", "u", "v", "w"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestStrict_Example8_BottomConsecutive(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC"}
	current := []string{"RuleC", "u", "RuleA", "v", "RuleB", "w"}

	moves, err := PlanMoves(desired, current, "bottom", "", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 3 {
		t.Fatalf("want 3 moves, got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "RuleC" || moves[0].Position != "bottom" {
		t.Fatalf("expected seed :move RuleC bottom, got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	want := []string{"u", "v", "w", "RuleA", "RuleB", "RuleC"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestStrict_Example9_BeforeConsecutive(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC"}
	current := []string{"w", "RuleC", "RuleA", "u", "RuleB", "v", "AnchorRule"}

	moves, err := PlanMoves(desired, current, "before", "AnchorRule", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 3 {
		t.Fatalf("want 3 moves, got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "RuleC" || moves[0].Position != "before" {
		t.Fatalf("expected seed :move RuleC before AnchorRule, got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	want := []string{"w", "u", "v", "RuleA", "RuleB", "RuleC", "AnchorRule"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

func TestStrict_Example10_PerformanceBenefit(t *testing.T) {
	desired := []string{"RuleA", "RuleB", "RuleC", "RuleD", "RuleE"}
	current := []string{"RuleA", "u", "RuleB", "RuleC", "RuleD", "RuleE"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 1 {
		t.Fatalf("want 1 move (perf demo), got %d: %+v", len(moves), moves)
	}
	if moves[0].Rule != "u" {
		t.Fatalf("expected :move u after ..., got %+v", moves[0])
	}

	final := simulateMoves(current, moves)
	want := []string{"RuleA", "RuleB", "RuleC", "RuleD", "RuleE", "u"}
	if !reflect.DeepEqual(final, want) {
		t.Fatalf("final want %v, got %v", want, final)
	}
}

// -----------------------------------------------------------------------------
// Edge cases.
// -----------------------------------------------------------------------------

func TestEdge_EmptyDesired(t *testing.T) {
	moves, err := PlanMoves(nil, []string{"a", "b"}, "null", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if moves != nil {
		t.Fatalf("want no moves, got %+v", moves)
	}
}

func TestEdge_SingleRule(t *testing.T) {
	moves, err := PlanMoves([]string{"A"}, []string{"A"}, "null", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 0 {
		t.Fatalf("want 0 moves for single rule already present, got %+v", moves)
	}
}

func TestEdge_AlreadyInOrder_Null(t *testing.T) {
	desired := []string{"A", "B", "C"}
	current := []string{"A", "B", "C"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 0 {
		t.Fatalf("want 0 moves, got %+v", moves)
	}
}

func TestEdge_AlreadyInOrder_Top_SeedNoop(t *testing.T) {
	desired := []string{"A", "B", "C"}
	current := []string{"A", "B", "C", "x"}

	moves, err := PlanMoves(desired, current, "top", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 0 {
		t.Fatalf("want 0 moves (A already at top, all in order), got %+v", moves)
	}
}

func TestEdge_WorstCase_ReverseOrder(t *testing.T) {
	desired := []string{"A", "B", "C", "D"}
	current := []string{"D", "C", "B", "A"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// LIS length = 1 (single element). n - |LIS| = 3 chain moves.
	if len(moves) != 3 {
		t.Fatalf("want 3 moves, got %d: %+v", len(moves), moves)
	}

	final := simulateMoves(current, moves)
	if got := projection(final, desired); !reflect.DeepEqual(got, desired) {
		t.Fatalf("projection want %v, got %v (final=%v)", desired, got, final)
	}
}

func TestEdge_UnmanagedInterleaversPreserveRelativeOrder(t *testing.T) {
	// From "Key Property" example: strictly-in-order targets with interleavers,
	// consecutive mode → zero moves.
	desired := []string{"A", "B", "C"}
	current := []string{"u", "A", "v", "B", "w", "C", "x"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 0 {
		t.Fatalf("want 0 moves, got %+v", moves)
	}
}

func TestEdge_ValidationErrors(t *testing.T) {
	cases := []struct {
		name        string
		desired     []string
		current     []string
		destination string
		anchor      string
		mode        string
		wantErr     bool
	}{
		{
			name:        "after without anchor",
			desired:     []string{"A"},
			current:     []string{"A"},
			destination: "after",
			wantErr:     true,
		},
		{
			name:        "anchor not in current",
			desired:     []string{"A"},
			current:     []string{"A"},
			destination: "after",
			anchor:      "X",
			wantErr:     true,
		},
		{
			name:        "anchor in desired",
			desired:     []string{"A", "X"},
			current:     []string{"A", "X"},
			destination: "after",
			anchor:      "X",
			wantErr:     true,
		},
		{
			name:        "desired rule missing from current",
			desired:     []string{"A", "B"},
			current:     []string{"A"},
			destination: "null",
			wantErr:     true,
		},
		{
			name:        "duplicate in desired",
			desired:     []string{"A", "A"},
			current:     []string{"A"},
			destination: "null",
			wantErr:     true,
		},
		{
			name:        "invalid destination",
			desired:     []string{"A"},
			current:     []string{"A"},
			destination: "middle",
			wantErr:     true,
		},
		{
			name:        "invalid mode",
			desired:     []string{"A"},
			current:     []string{"A"},
			destination: "null",
			mode:        "weird",
			wantErr:     true,
		},
		{
			name:        "top with anchor set is rejected",
			desired:     []string{"A"},
			current:     []string{"A", "X"},
			destination: "top",
			anchor:      "X",
			wantErr:     true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mode := tc.mode
			if mode == "" {
				mode = OrderModeConsecutive
			}
			_, err := PlanMoves(tc.desired, tc.current, tc.destination, tc.anchor, mode)
			if tc.wantErr && err == nil {
				t.Fatalf("want error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestConsecutive_NoSeedWhenPinAlreadyPlaced(t *testing.T) {
	// RuleA is already immediately after AnchorRule; the seed should not fire.
	desired := []string{"RuleA", "RuleB"}
	current := []string{"AnchorRule", "RuleA", "RuleB"}

	moves, err := PlanMoves(desired, current, "after", "AnchorRule", OrderModeConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 0 {
		t.Fatalf("want 0 moves (already in target order), got %+v", moves)
	}
}

func TestStrict_PreferTargetTieBreak(t *testing.T) {
	// Example 3 from tech spec (null, consecutive mode has multiple length-3
	// LIS witnesses). Verify that the prefer-target tie-break picks the
	// witness that keeps all targets stationary and moves the unmanaged rules.
	desired := []string{"RuleA", "RuleB", "RuleC"}
	current := []string{"RuleA", "u", "RuleB", "v", "RuleC"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// prefer-target picks {RuleA, RuleB, RuleC} → move set = {u, v}.
	for _, m := range moves {
		if m.Rule != "u" && m.Rule != "v" {
			t.Fatalf("expected only unmanaged rules to move under prefer-target tie-break, got move on %q", m.Rule)
		}
	}
}

// -----------------------------------------------------------------------------
// Determinism: same inputs → same outputs.
// -----------------------------------------------------------------------------

func TestDeterminism_RepeatedRuns(t *testing.T) {
	desired := []string{"A", "B", "C", "D"}
	current := []string{"D", "u", "A", "C", "v", "B"}

	for _, mode := range []string{OrderModeConsecutive, OrderModeStrictConsecutive} {
		var first []Move
		for i := 0; i < 5; i++ {
			m, err := PlanMoves(desired, current, "null", "", mode)
			if err != nil {
				t.Fatalf("[mode=%s] unexpected error: %v", mode, err)
			}
			if i == 0 {
				first = m
				continue
			}
			if !reflect.DeepEqual(m, first) {
				t.Fatalf("[mode=%s] non-deterministic plan; run %d differs from run 0: %+v vs %+v", mode, i, m, first)
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Regression tests: apply the planned moves and assert the end state, using the
// same observeCurrentBlock check the resource uses for post-apply verification.
// -----------------------------------------------------------------------------

func TestStrict_NullDestination_FirstTargetNotInLIS(t *testing.T) {
	desired := []string{"A", "B", "C"}
	current := []string{"B", "A", "u1", "C"}

	moves, err := PlanMoves(desired, current, "null", "", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	final := simulateMoves(current, moves)
	if got := observeCurrentBlock(final, desired, "null", "", OrderModeStrictConsecutive); !reflect.DeepEqual(got, desired) {
		t.Fatalf("block want %v, got %v (final=%v, moves=%+v)", desired, got, final, moves)
	}
}

// strictly_consecutive + null: the block may sit anywhere, so targets that are
// already contiguous and in order must produce zero moves.
func TestStrict_NullDestination_AlreadyContiguousNoMoves(t *testing.T) {
	desired := []string{"A", "B", "C"}
	cases := [][]string{
		{"A", "B", "C"},
		{"u1", "u2", "u3", "u4", "A", "B", "C"},
		{"u1", "A", "B", "C", "u2"},
		{"A", "B", "C", "u1", "u2"},
	}
	for _, current := range cases {
		moves, err := PlanMoves(desired, current, "null", "", OrderModeStrictConsecutive)
		if err != nil {
			t.Fatalf("%v: unexpected error: %v", current, err)
		}
		if len(moves) != 0 {
			t.Errorf("%v: want 0 moves, got %+v", current, moves)
		}
	}
}

// Interleaved unmanaged rules must be moved out, and unmanaged relative order
// must be preserved.
func TestStrict_NullDestination_InterleaverMovedOut(t *testing.T) {
	desired := []string{"A", "B", "C"}
	current := []string{"u1", "A", "u2", "B", "C", "u3"}
	moves, err := PlanMoves(desired, current, "null", "", OrderModeStrictConsecutive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(moves) != 1 {
		t.Fatalf("want 1 move (u2 out), got %+v", moves)
	}
	final := simulateMoves(current, moves)
	if got := observeCurrentBlock(final, desired, "null", "", OrderModeStrictConsecutive); !reflect.DeepEqual(got, desired) {
		t.Fatalf("block want %v, got %v (final=%v)", desired, got, final)
	}
	if got, want := projection(final, []string{"u1", "u2", "u3"}), []string{"u1", "u2", "u3"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unmanaged order changed: %v", got)
	}
}

// Consecutive mode with an absolute/anchored destination must detect a block
// that has been moved out of position, even when relative order is intact.
func TestObserve_ConsecutiveDetectsDestinationDrift(t *testing.T) {
	desired := []string{"A", "B"}
	cases := []struct {
		name        string
		current     []string
		destination string
		anchor      string
	}{
		{"top", []string{"u1", "A", "B"}, "top", ""},
		{"bottom", []string{"A", "B", "u1"}, "bottom", ""},
		{"after", []string{"X", "u1", "A", "B"}, "after", "X"},
		{"before", []string{"A", "B", "u1", "X"}, "before", "X"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := observeCurrentBlock(tc.current, desired, tc.destination, tc.anchor, OrderModeConsecutive)
			if reflect.DeepEqual(got, desired) {
				t.Fatalf("destination drift not detected: current=%v observed=%v", tc.current, got)
			}
		})
	}
}

// Property test: for every permutation of a small rulebase and every
// destination/mode combination, applying PlanMoves must satisfy the same
// check the resource runs after apply.
func TestProperty_PlannedMovesSatisfyObservation(t *testing.T) {
	universe := []string{"A", "B", "C", "u1", "u2", "X"}
	desired := []string{"A", "B", "C"}
	dests := []struct{ dest, anchor string }{
		{"null", ""}, {"top", ""}, {"bottom", ""}, {"after", "X"}, {"before", "X"},
	}
	for _, perm := range permutations(universe) {
		for _, d := range dests {
			for _, mode := range []string{OrderModeConsecutive, OrderModeStrictConsecutive} {
				moves, err := PlanMoves(desired, perm, d.dest, d.anchor, mode)
				if err != nil {
					t.Fatalf("%v %s/%s: unexpected error: %v", perm, d.dest, mode, err)
				}
				final := simulateMoves(perm, moves)
				if got := observeCurrentBlock(final, desired, d.dest, d.anchor, mode); !reflect.DeepEqual(got, desired) {
					t.Errorf("start=%v dest=%s anchor=%s mode=%s: moves=%+v final=%v observed=%v want=%v",
						perm, d.dest, d.anchor, mode, moves, final, got, desired)
				}
			}
		}
	}
}

func permutations(in []string) [][]string {
	if len(in) <= 1 {
		return [][]string{append([]string(nil), in...)}
	}
	var out [][]string
	for i := range in {
		rest := append(append([]string(nil), in[:i]...), in[i+1:]...)
		for _, p := range permutations(rest) {
			out = append(out, append([]string{in[i]}, p...))
		}
	}
	return out
}

// A single target rule must still honor destination (the resource used to
// skip planning entirely for fewer than two rules).
func TestSingleRule_HonorsDestination(t *testing.T) {
	current := []string{"u1", "R1", "u2", "X"}
	cases := []struct{ dest, anchor string }{{"top", ""}, {"bottom", ""}, {"after", "X"}, {"before", "X"}}
	for _, tc := range cases {
		t.Run(tc.dest, func(t *testing.T) {
			moves, err := PlanMoves([]string{"R1"}, current, tc.dest, tc.anchor, OrderModeConsecutive)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(moves) == 0 {
				t.Fatalf("expected a move for destination %s", tc.dest)
			}
			final := simulateMoves(current, moves)
			if got := observeCurrentBlock(final, []string{"R1"}, tc.dest, tc.anchor, OrderModeConsecutive); !reflect.DeepEqual(got, []string{"R1"}) {
				t.Fatalf("final=%v observed=%v", final, got)
			}
		})
	}
}

func TestParseRuleOrderImportID(t *testing.T) {
	ok := []struct {
		id                    string
		kind, value, rulebase string
		rules                 []string
	}{
		{"GLENCORE:::pre", "folder", "GLENCORE", "pre", nil},
		{":S1::post:A,B,C", "snippet", "S1", "post", []string{"A", "B", "C"}},
		{"::fw 1:pre:Rule One, Rule-Two", "device", "fw 1", "pre", []string{"Rule One", "Rule-Two"}},
		{"Mobile Users:::pre:A", "folder", "Mobile Users", "pre", []string{"A"}},
	}
	for _, tc := range ok {
		kind, value, rb, rules, err := ParseRuleOrderImportID(tc.id)
		if err != nil {
			t.Errorf("%q: unexpected error: %v", tc.id, err)
			continue
		}
		if kind != tc.kind || value != tc.value || rb != tc.rulebase || !reflect.DeepEqual(rules, tc.rules) {
			t.Errorf("%q: got (%s,%s,%s,%v)", tc.id, kind, value, rb, rules)
		}
	}
	bad := []string{
		"", "foo", "GLENCORE", "GLENCORE:pre", "GLENCORE::pre",
		":::pre",            // no container
		"A:B::pre",          // two containers
		"GLENCORE:::middle", // bad rulebase
		"GLENCORE:::pre:A,,B", "GLENCORE:::pre:A,A",
	}
	for _, id := range bad {
		if _, _, _, _, err := ParseRuleOrderImportID(id); err == nil {
			t.Errorf("%q: expected an error", id)
		}
	}
}
