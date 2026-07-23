package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
)

// GetActions returns all actions for this package.
func GetActions() []func() action.Action {
	return []func() action.Action{
		NewSnippetSnapshotConvertAction,
		NewSnippetSnapshotSaveAction,
		NewSharedSnippetsLoadAction,
		NewSnippetSnapshotCompareAction,
		NewSnippetSnapshotLoadAction,
		NewTrustValidationAction,
		NewSnippetSnapshotPublishAction,
		NewSnippetSnapshotDiffAction,
		NewSnippetSnapshotUpdatesAction,
	}
}
