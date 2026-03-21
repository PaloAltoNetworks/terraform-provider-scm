package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE for SCM Device (Package: config_setup)
var (
	_ datasource.DataSource              = &DeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDataSource{}
)

func NewDeviceDataSource() datasource.DataSource {
	return &DeviceDataSource{}
}

// DeviceDataSource defines the data source implementation.
type DeviceDataSource struct {
	client *config_setup.APIClient
}

func (d *DeviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (d *DeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.DevicesDataSourceSchema
}

func (d *DeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
		resp.Diagnostics.AddError("Unexpected Client Type", "Expected *config_setup.APIClient for 'config_setup' client.")
		return
	}
	d.client = client
}

func (d *DeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: DeviceDataSource.Read ---")

	var data models.DevicesTf
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Missing Identifier", "'id' must be provided for the Device data source.")
		return
	}

	objectId := data.Id.ValueString()
	tflog.Debug(ctx, "Reading Device data source by ID", map[string]interface{}{"id": objectId})

	getReq := d.client.DevicesAPI.GetDeviceByID(ctx, objectId)
	scmObject, httpRes, err := getReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Device", fmt.Sprintf("Could not read Device with ID %s: %s", objectId, err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError("Resource Get Failed: API Request Failed", detailedMessage)
		return
	}
	if httpRes.StatusCode != 200 {
		resp.Diagnostics.AddError("Unexpected HTTP status code", fmt.Sprintf("Expected 200, got %d", httpRes.StatusCode))
		return
	}

	// Pack the response into the Terraform model.
	model, diags := packDevicesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the composite Tfid.
	var idBuilder strings.Builder
	idBuilder.WriteString(":::")
	idBuilder.WriteString(model.Id.ValueString())
	model.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}
