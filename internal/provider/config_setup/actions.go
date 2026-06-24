package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
)

// GetActions returns all actions for this package.
func GetActions() []func() action.Action {
	return []func() action.Action{
		NewSnippetSnapshotDiffAction,
		NewSharedSnippetsLoadAction,
		NewSnippetSnapshotLoadAction,
		NewSnippetSnapshotConvertAction,
		NewSnippetSnapshotCompareAction,
		NewTrustValidationAction,
		NewSnippetSnapshotSaveAction,
		NewSnippetSnapshotPublishAction,
		NewSnippetSnapshotUpdatesAction,
	}
}
