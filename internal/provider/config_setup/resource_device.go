package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// RESOURCE for SCM Device (Package: config_setup)
// This resource manages device-to-folder assignment, labels, snippets, description, and display_name.
// The device itself is NOT created or deleted by Terraform — it already exists in SCM.
var (
	_ resource.Resource                = &DeviceResource{}
	_ resource.ResourceWithConfigure   = &DeviceResource{}
	_ resource.ResourceWithImportState = &DeviceResource{}
)

func NewDeviceResource() resource.Resource {
	return &DeviceResource{}
}

// DeviceResource defines the resource implementation.
type DeviceResource struct {
	client *config_setup.APIClient
}

func (r *DeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (r *DeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.DevicesResourceSchema
}

func (r *DeviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["config_setup"].(*config_setup.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", "Expected *config_setup.APIClient for 'config_setup' client.")
		return
	}
	r.client = client
}

// Create adopts an existing device by setting the desired folder, labels, snippets, etc.
// Since the device already exists in SCM, Create performs an Update to set the desired state.
func (r *DeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for Device")
	var plan models.DevicesTf

	// 1. Get the plan from Terraform.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceId := plan.Id.ValueString()
	tflog.Debug(ctx, "Adopting device into Terraform management", map[string]interface{}{"id": deviceId})

	// 2. Unpack the plan into an SDK DevicesPut object.
	sdkPut, diags := unpackDevicesToSdkPut(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 3. Update the device to set the desired folder/labels/snippets.
	updateReq := r.client.DevicesAPI.UpdateDeviceByID(ctx, deviceId).DevicesPut(*sdkPut)
	_, httpErr, err := updateReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating device during create", err.Error())
		if httpErr != nil {
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Creation Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// 4. Read back the device to get the full state.
	getReq := r.client.DevicesAPI.GetDeviceByID(ctx, deviceId)
	scmObject, _, err := getReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading device after create", err.Error())
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		return
	}

	// 5. Pack the response into the Terraform model.
	model, diags := packDevicesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 6. Set the Terraform ID.
	var idBuilder strings.Builder
	idBuilder.WriteString(":::")
	idBuilder.WriteString(model.Id.ValueString())
	model.Tfid = types.StringValue(idBuilder.String())

	tflog.Debug(ctx, "Created (adopted) device", map[string]interface{}{"tfid": model.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *DeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Starting Read function for Device")
	var savestate models.DevicesTf

	// 1. Fetch the state.
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tokens := strings.Split(savestate.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), savestate.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	// 2. Read the device from the API.
	tflog.Debug(ctx, "Reading device from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.DevicesAPI.GetDeviceByID(ctx, objectId)
	scmObject, httpErr, err := getReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Device not found on read. Removing from state.", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Error reading device", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// 3. Pack the response.
	model, diags := packDevicesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 4. Carry over the TFID from existing state.
	model.Tfid = savestate.Tfid

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *DeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Starting Update function for Device")
	var plan models.DevicesTf
	var state models.DevicesTf

	// 1. Get plan and state.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the plan into DevicesPut.
	sdkPut, diags := unpackDevicesToSdkPut(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 3. Get the ID from the TFID.
	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	// 4. Make the update call.
	tflog.Debug(ctx, "Updating device on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.DevicesAPI.UpdateDeviceByID(ctx, objectId).DevicesPut(*sdkPut)
	_, httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Device not found on update. Removing from state.", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Error updating device", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Update Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// 5. Read back the device to get the full state after update.
	getReq := r.client.DevicesAPI.GetDeviceByID(ctx, objectId)
	scmObject, _, err := getReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading device after update", err.Error())
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		return
	}

	// 6. Pack the response.
	model, diags := packDevicesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 7. Preserve the ID from the plan and TFID from state.
	_ = req.Plan.GetAttribute(ctx, path.Root("id"), &model.Id)
	model.Tfid = state.Tfid

	tflog.Debug(ctx, "Updated device", map[string]interface{}{"tfid": model.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

// Delete removes the device from Terraform state only.
// The physical device is NOT deleted from SCM — it was onboarded externally.
func (r *DeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Starting Delete function for Device")
	var data models.DevicesTf
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Removing device from Terraform state. The device remains in SCM.", map[string]interface{}{
		"id":   data.Id.ValueString(),
		"name": data.Name.ValueString(),
	})
	// No API call — just allow Terraform to remove the resource from state.
}

func (r *DeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
