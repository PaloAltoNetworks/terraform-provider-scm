package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
)

// GetActions returns all actions for this package.
func GetActions() []func() action.Action {
	return []func() action.Action{
		NewTrustValidationAction,
		NewSnippetSnapshotUpdatesAction,
		NewSharedSnippetsLoadAction,
		NewSnippetSnapshotLoadAction,
		NewSnippetSnapshotCompareAction,
		NewSnippetSnapshotPublishAction,
		NewSnippetSnapshotConvertAction,
		NewSnippetSnapshotDiffAction,
		NewSnippetSnapshotSaveAction,
	}
}
