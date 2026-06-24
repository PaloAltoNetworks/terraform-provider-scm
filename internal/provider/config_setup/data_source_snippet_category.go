package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE for SCM SnippetCategory (Package: config_setup)
var (
	_ datasource.DataSource              = &SnippetCategoryDataSource{}
	_ datasource.DataSourceWithConfigure = &SnippetCategoryDataSource{}
)

func NewSnippetCategoryDataSource() datasource.DataSource {
	return &SnippetCategoryDataSource{}
}

// SnippetCategoryDataSource defines the data source implementation.
type SnippetCategoryDataSource struct {
	client *config_setup.APIClient
}

func (d *SnippetCategoryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: SnippetCategoryDataSource.Metadata ---")
	resp.TypeName = "scm_snippet_category"
}

func (d *SnippetCategoryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: SnippetCategoryDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.SnippetCategoriesDataSourceSchema
}

func (d *SnippetCategoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: SnippetCategoriesDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["config_setup"].(*config_setup.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *config_setup.APIClient for 'config_setup' client."))
		return
	}
	d.client = client
}

func (d *SnippetCategoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: SnippetCategoryDataSource.Read ---")

	var data models.SnippetCategories

	// TF LOGGING ADDED: Log before the potentially crashing line.
	tflog.Debug(ctx, "--- VERIFICATION LOG: About to call req.Config.Get() ---")

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// TF LOGGING ADDED: This log will ONLY print if the line above succeeds.
	tflog.Debug(ctx, "--- VERIFICATION LOG: Call to req.Config.Get() succeeded. ---")

	if resp.Diagnostics.HasError() {
		// TF LOGGING ADDED: Log if diagnostics has an error after Get().
		tflog.Debug(ctx, "--- VERIFICATION LOG: req.Config.Get() resulted in a diagnostic error.")
		return
	}

	// Logic to handle read by ID or by name/list.
	// We prioritize reading by ID if it is provided.
	if !data.Id.IsNull() {
		objectId := data.Id.ValueString()
		tflog.Debug(ctx, "Reading SnippetCategories data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.SnippetCategoriesAPI.GetSnippetCategoryByID(ctx, objectId)
		scmObject, httpRes, err := getReq.Execute()

		var statusCode int
		if httpRes != nil {
			statusCode = httpRes.StatusCode
		}
		tflog.Debug(ctx, "--- SCM API CALL: Execution complete.", map[string]interface{}{
			"error":          err,
			"statusCode":     statusCode,
			"scmObjectIsNil": scmObject == nil,
		})

		if err != nil {
			resp.Diagnostics.AddError("Error Reading SnippetCategories", fmt.Sprintf("Could not read SnippetCategories with ID %s: %s", objectId, err.Error()))
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError(
				"Resource Get Failed: API Request Failed",
				detailedMessage,
			)
			return
		}
		if httpRes.StatusCode != 200 {
			resp.Diagnostics.AddError("Unexpected HTTP status code", fmt.Sprintf("Expected 200, got %d", httpRes.StatusCode))
			return
		}

		tflog.Debug(ctx, "--- DATA SOURCE READ: API call successful. About to call the packer.")

		// Create a packed object from the SCM response.
		packedObject, diags := packSnippetCategoriesFromSdk(ctx, *scmObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Load the packed object into the data model.
		resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

	} else if !data.Name.IsNull() {
		objectName := data.Name.ValueString()
		tflog.Debug(ctx, "Reading SnippetCategories data source by Name", map[string]interface{}{"name": objectName})

		listReq := d.client.SnippetCategoriesAPI.ListSnippetCategories(ctx)

		listResponse, httpRes, err := listReq.Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error Listing SnippetCategoriess", fmt.Sprintf("Could not list SnippetCategoriess: %s", err.Error()))
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError(
				"Resource Listing Failed: API Request Failed",
				detailedMessage,
			)
			return
		}
		if httpRes.StatusCode != 200 {
			resp.Diagnostics.AddError("Unexpected HTTP status code", fmt.Sprintf("Expected 200, got %d", httpRes.StatusCode))
			return
		}

		// Find the specific object from the list.
		var foundObject *config_setup.SnippetCategories
		for i := range listResponse.GetData() {
			item := listResponse.GetData()[i]
			if item.GetName() == objectName {
				foundObject = &item
				break
			}
		}

		if foundObject == nil {
			resp.Diagnostics.AddError("SnippetCategories Not Found", fmt.Sprintf("No SnippetCategories found with name: %s", objectName))
			return
		}

		// Create a packed object from the SCM response.
		packedObject, diags := packSnippetCategoriesFromSdk(ctx, *foundObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Load the packed object into the data model.
		resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

	} else {
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the SnippetCategories data source.")
		return
	}

	// Create the composite Tfid for consistency.
	var idBuilder strings.Builder

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
