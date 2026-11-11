package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/cloud_ngfw"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/cloud_ngfw"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE for SCM Ngfirewall (Package: cloud_ngfw)
var (
	_ datasource.DataSource              = &NgfirewallDataSource{}
	_ datasource.DataSourceWithConfigure = &NgfirewallDataSource{}
)

func NewNgfirewallDataSource() datasource.DataSource {
	return &NgfirewallDataSource{}
}

// NgfirewallDataSource defines the data source implementation.
type NgfirewallDataSource struct {
	client *cloud_ngfw.APIClient
}

func (d *NgfirewallDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: NgfirewallDataSource.Metadata ---")
	resp.TypeName = "scm_cngfw_ngfirewall"
}

func (d *NgfirewallDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: NgfirewallDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.NgfirewallsDataSourceSchema
}

func (d *NgfirewallDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: NgfirewallsDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["cloud_ngfw"].(*cloud_ngfw.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *cloud_ngfw.APIClient for 'cloud_ngfw' client."))
		return
	}
	d.client = client
}

func (d *NgfirewallDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: NgfirewallDataSource.Read ---")

	var data models.Ngfirewalls

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
		tflog.Debug(ctx, "Reading Ngfirewalls data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.NgfirewallsAPI.GetNgfirewallsByID(ctx, objectId)
		if !data.Region.IsNull() {
			getReq = getReq.Region(data.Region.ValueString())
		}
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
			resp.Diagnostics.AddError("Error Reading Ngfirewalls", fmt.Sprintf("Could not read Ngfirewalls with ID %s: %s", objectId, err.Error()))
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
		// Save operation parameters before packing — the packer only knows about
		// response body fields and will null-out path/query params otherwise.
		// If these values provided by the response also then these will be auto recovered even after packing. But
		// for Cngfw resources, these values are not returned at the response level rather nested inside the response
		// body.
		savedId := data.Id
		savedRegion := data.Region

		// Create a packed object from the SCM response.
		packedObject, diags := packNgfirewallsFromSdk(ctx, *scmObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Load the packed object into the data model.
		resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		// Restore operation parameters overwritten by the packer.
		data.Id = savedId
		data.Region = savedRegion

	} else {
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the Ngfirewalls data source.")
		return
	}

	// Create the composite Tfid for consistency.
	var idBuilder strings.Builder
	idBuilder.WriteString(data.Region.ValueString())
	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
