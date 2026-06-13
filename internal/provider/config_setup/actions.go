package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
)

// GetActions returns all actions for this package.
func GetActions() []func() action.Action {
	return []func() action.Action{
		NewSnippetSnapshotCompareAction,
		NewSharedSnippetsLoadAction,
		NewTrustValidationAction,
		NewSnippetSnapshotConvertAction,
		NewSnippetSnapshotUpdatesAction,
		NewSnippetSnapshotDiffAction,
		NewSnippetSnapshotLoadAction,
		NewSnippetSnapshotSaveAction,
		NewSnippetSnapshotPublishAction,
	}
}
