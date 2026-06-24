package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE for SCM Subnet (Package: ztna_connector_all)
var (
	_ datasource.DataSource              = &SubnetDataSource{}
	_ datasource.DataSourceWithConfigure = &SubnetDataSource{}
)

func NewSubnetDataSource() datasource.DataSource {
	return &SubnetDataSource{}
}

// SubnetDataSource defines the data source implementation.
type SubnetDataSource struct {
	client *ztna_connector_all.APIClient
}

func (d *SubnetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: SubnetDataSource.Metadata ---")
	resp.TypeName = "ztna_subnet"
}

func (d *SubnetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: SubnetDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.SubnetsDataSourceSchema
}

func (d *SubnetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: SubnetsDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["ztna_connector_all"].(*ztna_connector_all.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *ztna_connector_all.APIClient for 'ztna_connector_all' client."))
		return
	}
	d.client = client
}

func (d *SubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: SubnetDataSource.Read ---")

	var data models.Subnets

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
	if !data.Oid.IsNull() {
		objectId := data.Oid.ValueString()
		tflog.Debug(ctx, "Reading Subnets data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.SubnetAPI.GetSubnetByID(ctx, objectId)
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
			resp.Diagnostics.AddError("Error Reading Subnets", fmt.Sprintf("Could not read Subnets with ID %s: %s", objectId, err.Error()))
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
		packedObject, diags := packSubnetsFromSdk(ctx, *scmObject)
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
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the Subnets data source.")
		return
	}

	// Create the composite Tfid for consistency.
	var idBuilder strings.Builder

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Oid.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
