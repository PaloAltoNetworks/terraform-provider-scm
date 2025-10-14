package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewFolderDataSource,
		NewLabelDataSource,
		NewSnippetDataSource,
		NewVariableDataSource,
		NewFolderListDataSource,
		NewLabelListDataSource,
		NewSnippetListDataSource,
		NewVariableListDataSource,
	}
}
