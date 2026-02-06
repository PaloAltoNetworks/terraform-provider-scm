package provider

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// RESOURCE for SCM ServiceConnection (Package: deployment_services)
var (
	_ resource.Resource                = &ServiceConnectionResource{}
	_ resource.ResourceWithConfigure   = &ServiceConnectionResource{}
	_ resource.ResourceWithImportState = &ServiceConnectionResource{}
)

func NewServiceConnectionResource() resource.Resource {
	return &ServiceConnectionResource{}
}

// ServiceConnectionResource defines the resource implementation.
type ServiceConnectionResource struct {
	client *deployment_services.APIClient
}

func (r *ServiceConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_connection"
}

func (r *ServiceConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.ServiceConnectionsResourceSchema
}

func (r *ServiceConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["deployment_services"].(*deployment_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *deployment_services.APIClient for 'deployment_services' client."))
		return
	}
	r.client = client
}

func (r *ServiceConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for ServiceConnection")
	var data models.ServiceConnections

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a patcher to temporarily store sensitive values.
	patcher := &serviceConnectionsSensitiveValuePatcher{}

	// Stash plaintext values from the plan.

	{ // Stash plaintext for BgpPeer.Secret
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.BgpPeer.IsNull() && !data.BgpPeer.IsUnknown() {
				var temp_stash_BgpPeer_Secret_0 models.ServiceConnectionsBgpPeer
				resp.Diagnostics.Append(data.BgpPeer.As(ctx, &temp_stash_BgpPeer_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					finalVal = temp_stash_BgpPeer_Secret_0.Secret

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.bgp_peer_secret_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Protocol.Bgp.Secret
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
				var temp_stash_Protocol_Bgp_Secret_0 models.ServiceConnectionsProtocol
				resp.Diagnostics.Append(data.Protocol.As(ctx, &temp_stash_Protocol_Bgp_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_Protocol_Bgp_Secret_0.Bgp.IsNull() && !temp_stash_Protocol_Bgp_Secret_0.Bgp.IsUnknown() {
						var temp_stash_Protocol_Bgp_Secret_1 models.ServiceConnectionsProtocolBgp
						resp.Diagnostics.Append(temp_stash_Protocol_Bgp_Secret_0.Bgp.As(ctx, &temp_stash_Protocol_Bgp_Secret_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_Protocol_Bgp_Secret_1.Secret

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.protocol_bgp_secret_plaintext = finalVal
		}
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.ServiceConnections{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackServiceConnectionsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating service_connections on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.ServiceConnectionsAPI.CreateServiceConnections(ctx).ServiceConnections(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating service_connections", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packServiceConnectionsFromSdk(ctx, *createdObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 7. BLOCK 2: Restore the PARAMETER values from the original plan.
	//    This is necessary for parameters that are sent to the API but not returned in the response.

	// FOLDER NORMALIZATION: Handle folder value translation and normalization.
	// This handles both deprecated value translation and Shared/Prisma Access normalization.
	// Get the user's configured folder value from the request plan.
	var userConfiguredFolder basetypes.StringValue
	_ = req.Plan.GetAttribute(ctx, path.Root("folder"), &userConfiguredFolder)

	if !userConfiguredFolder.IsNull() && !userConfiguredFolder.IsUnknown() && !data.Folder.IsNull() {
		userFolderStr := userConfiguredFolder.ValueString()
		apiReturnedFolder := data.Folder.ValueString()

		// First, translate any deprecated values from the API
		translatedFolder := utils.FolderAPIToState(apiReturnedFolder)

		// Then, if user configured "Shared" or "Prisma Access", normalize to match user's choice
		if (userFolderStr == "Shared" || userFolderStr == "Prisma Access") &&
			(translatedFolder == "Shared" || translatedFolder == "Prisma Access") {
			translatedFolder = userFolderStr
		}

		// Set the final normalized value
		data.Folder = basetypes.NewStringValue(translatedFolder)
		tflog.Debug(ctx, "Normalized folder value", map[string]interface{}{
			"configured":   userFolderStr,
			"api_returned": apiReturnedFolder,
			"normalized":   translatedFolder,
		})
	}

	// Stash the encrypted values from the API response and apply the patch.

	{ // Patch plaintext for BgpPeer.Secret
		if !resp.Diagnostics.HasError() {

			if !data.BgpPeer.IsNull() && !data.BgpPeer.IsUnknown() {
				var temp_patch_create_BgpPeer_Secret_0 models.ServiceConnectionsBgpPeer
				resp.Diagnostics.Append(data.BgpPeer.As(ctx, &temp_patch_create_BgpPeer_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					patcher.bgp_peer_secret_encrypted = temp_patch_create_BgpPeer_Secret_0.Secret
					temp_patch_create_BgpPeer_Secret_0.Secret = patcher.bgp_peer_secret_plaintext

					// Repack the modified structs.

					if !resp.Diagnostics.HasError() {

						data.BgpPeer, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsBgpPeer{}.AttrTypes(), &temp_patch_create_BgpPeer_Secret_0)

						resp.Diagnostics.Append(diags...)
					}

				}
			}

		}
	}

	{ // Patch plaintext for Protocol.Bgp.Secret
		if !resp.Diagnostics.HasError() {

			if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
				var temp_patch_create_Protocol_Bgp_Secret_0 models.ServiceConnectionsProtocol
				resp.Diagnostics.Append(data.Protocol.As(ctx, &temp_patch_create_Protocol_Bgp_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_Protocol_Bgp_Secret_0.Bgp.IsNull() && !temp_patch_create_Protocol_Bgp_Secret_0.Bgp.IsUnknown() {
						var temp_patch_create_Protocol_Bgp_Secret_1 models.ServiceConnectionsProtocolBgp
						resp.Diagnostics.Append(temp_patch_create_Protocol_Bgp_Secret_0.Bgp.As(ctx, &temp_patch_create_Protocol_Bgp_Secret_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.protocol_bgp_secret_encrypted = temp_patch_create_Protocol_Bgp_Secret_1.Secret
							temp_patch_create_Protocol_Bgp_Secret_1.Secret = patcher.protocol_bgp_secret_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_create_Protocol_Bgp_Secret_0.Bgp, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocolBgp{}.AttrTypes(), &temp_patch_create_Protocol_Bgp_Secret_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.Protocol, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocol{}.AttrTypes(), &temp_patch_create_Protocol_Bgp_Secret_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	// Save the patcher's data to the EncryptedValues map in the state.
	evMap := patcher.populateEncryptedValuesMap()
	if len(evMap) > 0 {
		// Create the map using NewMapValueFrom with explicit context and type
		data.EncryptedValues, diags = basetypes.NewMapValueFrom(ctx, basetypes.StringType{}, evMap)
	} else {
		data.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	}
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the Terraform ID and save the final state.
	var idBuilder strings.Builder

	v := reflect.ValueOf(data)

	if f := v.FieldByName("Folder"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	tflog.Debug(ctx, "Created service_connections", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.ServiceConnections - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for ServiceConnection")
	var data, savestate models.ServiceConnections

	// Step 2 - Fetch the state into savestate
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

	// Step 3 - Make read api call with id = id from state tfid
	tflog.Debug(ctx, "Reading service_connections from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.ServiceConnectionsAPI.GetServiceConnectionsByID(ctx, objectId)
	scmObject, httpErr, err := getReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no service_connections on read SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on read SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error reading service_connections", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Read Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// Step 4 - Encrypted values logic
	// Populate a patcher from the state's encrypted_values map.
	patcher := &serviceConnectionsSensitiveValuePatcher{}
	resp.Diagnostics.Append(patcher.populatePatcherFromState(ctx, savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packServiceConnectionsFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// FOLDER NORMALIZATION: Handle folder value translation and normalization.
	// This handles both deprecated value translation and Shared/Prisma Access normalization.
	if !savestate.Folder.IsNull() && !savestate.Folder.IsUnknown() && !data.Folder.IsNull() {
		savedFolderValue := savestate.Folder.ValueString()
		apiReturnedFolder := data.Folder.ValueString()

		// First, translate any deprecated values from the API
		translatedFolder := utils.FolderAPIToState(apiReturnedFolder)

		// Then, if saved state had "Shared" or "Prisma Access", normalize to match saved state
		if (savedFolderValue == "Shared" || savedFolderValue == "Prisma Access") &&
			(translatedFolder == "Shared" || translatedFolder == "Prisma Access") {
			translatedFolder = savedFolderValue
		}

		// Set the final normalized value
		data.Folder = basetypes.NewStringValue(translatedFolder)
		tflog.Debug(ctx, "Normalized folder value", map[string]interface{}{
			"saved_state":  savedFolderValue,
			"api_returned": apiReturnedFolder,
			"normalized":   translatedFolder,
		})
	}

	// Step 6 - Encrypted values logic
	// Check for out-of-band changes and apply the patch.

	{ // Patch plaintext for BgpPeer.Secret
		if !resp.Diagnostics.HasError() {

			if !data.BgpPeer.IsNull() && !data.BgpPeer.IsUnknown() {
				var temp_patch_read_BgpPeer_Secret_0 models.ServiceConnectionsBgpPeer
				resp.Diagnostics.Append(data.BgpPeer.As(ctx, &temp_patch_read_BgpPeer_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block: Perform comparison and patch.
					if patcher.bgp_peer_secret_encrypted.Equal(temp_patch_read_BgpPeer_Secret_0.Secret) {
						temp_patch_read_BgpPeer_Secret_0.Secret = patcher.bgp_peer_secret_plaintext
					} else {
						temp_patch_read_BgpPeer_Secret_0.Secret = basetypes.NewStringNull()
					}

					// Repack the modified structs.

					if !resp.Diagnostics.HasError() {

						data.BgpPeer, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsBgpPeer{}.AttrTypes(), &temp_patch_read_BgpPeer_Secret_0)

						resp.Diagnostics.Append(diags...)
					}

				}
			}

		}
	}

	{ // Patch plaintext for Protocol.Bgp.Secret
		if !resp.Diagnostics.HasError() {

			if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
				var temp_patch_read_Protocol_Bgp_Secret_0 models.ServiceConnectionsProtocol
				resp.Diagnostics.Append(data.Protocol.As(ctx, &temp_patch_read_Protocol_Bgp_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_Protocol_Bgp_Secret_0.Bgp.IsNull() && !temp_patch_read_Protocol_Bgp_Secret_0.Bgp.IsUnknown() {
						var temp_patch_read_Protocol_Bgp_Secret_1 models.ServiceConnectionsProtocolBgp
						resp.Diagnostics.Append(temp_patch_read_Protocol_Bgp_Secret_0.Bgp.As(ctx, &temp_patch_read_Protocol_Bgp_Secret_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block: Perform comparison and patch.
							if patcher.protocol_bgp_secret_encrypted.Equal(temp_patch_read_Protocol_Bgp_Secret_1.Secret) {
								temp_patch_read_Protocol_Bgp_Secret_1.Secret = patcher.protocol_bgp_secret_plaintext
							} else {
								temp_patch_read_Protocol_Bgp_Secret_1.Secret = basetypes.NewStringNull()
							}

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_read_Protocol_Bgp_Secret_0.Bgp, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocolBgp{}.AttrTypes(), &temp_patch_read_Protocol_Bgp_Secret_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.Protocol, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocol{}.AttrTypes(), &temp_patch_read_Protocol_Bgp_Secret_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	// Persist the patcher's data to the EncryptedValues map in the state.
	evMap := patcher.populateEncryptedValuesMap()
	if len(evMap) > 0 {
		// Create the map using NewMapValueFrom with explicit context and type
		data.EncryptedValues, diags = basetypes.NewMapValueFrom(ctx, basetypes.StringType{}, evMap)
	} else {
		data.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	}
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 7 - Carry over tfid from state back into data
	data.Tfid = savestate.Tfid

	// Step 8 - Set things in params back into data object from the savestate - things like position of security rule

	// Step 9 - Set folder, snippet, device from params back into data if present

	// --- FOLDER RESTORATION (tokens[0]) ---

	// Use reflection to safely restore the Folder field from the TFID token 0.
	vFolder := reflect.ValueOf(&data).Elem() // Unique variable: vFolder
	fFolder := vFolder.FieldByName("Folder") // Unique variable: fFolder

	if fFolder.IsValid() && fFolder.CanSet() {
		tokenValue := tokens[0]

		// No validation here - just restore the value from TFID
		// Validation happens during ImportState for actual import operations
		// This allows backward compatibility for existing state with old folder values

		if tokenValue != "" {
			newStringValue := basetypes.NewStringValue(tokenValue)
			fFolder.Set(reflect.ValueOf(newStringValue))
		} else {
			newNullValue := basetypes.NewStringNull()
			fFolder.Set(reflect.ValueOf(newNullValue))
		}
	}

	// --- SNIPPET RESTORATION (tokens[1]) ---

	// --- DEVICE RESTORATION (tokens[2]) ---

	// Step 10 - Set data back into tf state and done
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.ServiceConnections which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for ServiceConnection")
	var plan, state models.ServiceConnections

	// Step 2: Get the plan from plan file (resource.tf) into plan and state from tfstate into state
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 3: Encrypted values logic
	patcher := &serviceConnectionsSensitiveValuePatcher{}

	{ // Stash plaintext for BgpPeer.Secret
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.BgpPeer.IsNull() && !plan.BgpPeer.IsUnknown() {
				var temp_stash_upd_BgpPeer_Secret_0 models.ServiceConnectionsBgpPeer
				resp.Diagnostics.Append(plan.BgpPeer.As(ctx, &temp_stash_upd_BgpPeer_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					finalVal = temp_stash_upd_BgpPeer_Secret_0.Secret

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.bgp_peer_secret_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Protocol.Bgp.Secret
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() {
				var temp_stash_upd_Protocol_Bgp_Secret_0 models.ServiceConnectionsProtocol
				resp.Diagnostics.Append(plan.Protocol.As(ctx, &temp_stash_upd_Protocol_Bgp_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_Protocol_Bgp_Secret_0.Bgp.IsNull() && !temp_stash_upd_Protocol_Bgp_Secret_0.Bgp.IsUnknown() {
						var temp_stash_upd_Protocol_Bgp_Secret_1 models.ServiceConnectionsProtocolBgp
						resp.Diagnostics.Append(temp_stash_upd_Protocol_Bgp_Secret_0.Bgp.As(ctx, &temp_stash_upd_Protocol_Bgp_Secret_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_upd_Protocol_Bgp_Secret_1.Secret

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.protocol_bgp_secret_plaintext = finalVal
		}
	}

	// Step 3: Creates a plan object from the plan
	planObject, diags := types.ObjectValueFrom(ctx, models.ServiceConnections{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackServiceConnectionsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5: Update calls cannot have id sent in payload, so remove it
	// ID is a string, so we set it to its zero value ("") to omit it from the update payload.
	unpackedScmObject.Id = ""

	// Step 6: Get id from token and make update call
	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Updating service_connections on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.ServiceConnectionsAPI.UpdateServiceConnectionsByID(ctx, objectId).ServiceConnections(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no service_connections on update SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on update SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error updating service_connections", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Update Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packServiceConnectionsFromSdk(ctx, *updatedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve any operation parameter values from the plan (folder, snippet, device).
	// This ensures the user's configured value is preserved regardless of what the API returns.
	_ = req.Plan.GetAttribute(ctx, path.Root("id"), &plan.Id)

	// FOLDER NORMALIZATION: Handle folder value translation and normalization.
	// This handles both deprecated value translation and Shared/Prisma Access normalization.
	// We need to get the user's configured folder value from the request plan.
	var userConfiguredFolder basetypes.StringValue
	_ = req.Plan.GetAttribute(ctx, path.Root("folder"), &userConfiguredFolder)

	if !userConfiguredFolder.IsNull() && !userConfiguredFolder.IsUnknown() && !plan.Folder.IsNull() {
		userFolderStr := userConfiguredFolder.ValueString()
		apiReturnedFolder := plan.Folder.ValueString()

		// First, translate any deprecated values from the API
		translatedFolder := utils.FolderAPIToState(apiReturnedFolder)

		// Then, if user configured "Shared" or "Prisma Access", normalize to match user's choice
		if (userFolderStr == "Shared" || userFolderStr == "Prisma Access") &&
			(translatedFolder == "Shared" || translatedFolder == "Prisma Access") {
			translatedFolder = userFolderStr
		}

		// Set the final normalized value
		plan.Folder = basetypes.NewStringValue(translatedFolder)
		tflog.Debug(ctx, "Normalized folder value", map[string]interface{}{
			"configured":   userFolderStr,
			"api_returned": apiReturnedFolder,
			"normalized":   translatedFolder,
		})
	}

	// Step 10: Encrypted values logic

	{ // Patch plaintext for BgpPeer.Secret
		if !resp.Diagnostics.HasError() {

			if !plan.BgpPeer.IsNull() && !plan.BgpPeer.IsUnknown() {
				var temp_patch_upd_BgpPeer_Secret_0 models.ServiceConnectionsBgpPeer
				resp.Diagnostics.Append(plan.BgpPeer.As(ctx, &temp_patch_upd_BgpPeer_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					patcher.bgp_peer_secret_encrypted = temp_patch_upd_BgpPeer_Secret_0.Secret
					temp_patch_upd_BgpPeer_Secret_0.Secret = patcher.bgp_peer_secret_plaintext

					// Repack the modified structs.

					if !resp.Diagnostics.HasError() {

						plan.BgpPeer, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsBgpPeer{}.AttrTypes(), &temp_patch_upd_BgpPeer_Secret_0)

						resp.Diagnostics.Append(diags...)
					}

				}
			}

		}
	}

	{ // Patch plaintext for Protocol.Bgp.Secret
		if !resp.Diagnostics.HasError() {

			if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() {
				var temp_patch_upd_Protocol_Bgp_Secret_0 models.ServiceConnectionsProtocol
				resp.Diagnostics.Append(plan.Protocol.As(ctx, &temp_patch_upd_Protocol_Bgp_Secret_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_Protocol_Bgp_Secret_0.Bgp.IsNull() && !temp_patch_upd_Protocol_Bgp_Secret_0.Bgp.IsUnknown() {
						var temp_patch_upd_Protocol_Bgp_Secret_1 models.ServiceConnectionsProtocolBgp
						resp.Diagnostics.Append(temp_patch_upd_Protocol_Bgp_Secret_0.Bgp.As(ctx, &temp_patch_upd_Protocol_Bgp_Secret_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.protocol_bgp_secret_encrypted = temp_patch_upd_Protocol_Bgp_Secret_1.Secret
							temp_patch_upd_Protocol_Bgp_Secret_1.Secret = patcher.protocol_bgp_secret_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_upd_Protocol_Bgp_Secret_0.Bgp, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocolBgp{}.AttrTypes(), &temp_patch_upd_Protocol_Bgp_Secret_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								plan.Protocol, diags = types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocol{}.AttrTypes(), &temp_patch_upd_Protocol_Bgp_Secret_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}
	evMap := patcher.populateEncryptedValuesMap()
	if len(evMap) > 0 {
		// Create the map using NewMapValueFrom with explicit context and type
		plan.EncryptedValues, diags = basetypes.NewMapValueFrom(ctx, basetypes.StringType{}, evMap)
	} else {
		plan.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	}
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

	// Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule

	tflog.Debug(ctx, "Updated service_connections", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ServiceConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.ServiceConnections
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tokens := strings.Split(data.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), data.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Deleting service_connections", map[string]interface{}{"id": objectId})
	deleteReq := r.client.ServiceConnectionsAPI.DeleteServiceConnectionsByID(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting service_connections", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Deleteion Failed: API Request Failed",
			detailedMessage,
		)
	}
}

func (r *ServiceConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Parse and validate the import ID before storing it - block deprecated folder values
	// Note: Both "Shared" and "Prisma Access" are accepted as valid values
	importID := req.ID
	tokens := strings.Split(importID, ":::")

	if len(tokens) > 0 && tokens[0] != "" {
		folderValue := tokens[0]
		if folderValue == "All Firewalls" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'All Firewalls' is not allowed in import IDs. Please use 'ngfw-shared' instead.\nExample: terraform import scm_service_connection.example \"ngfw-shared\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Global" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Global' is not allowed in import IDs. Please use 'All' instead.\nExample: terraform import scm_service_connection.example \"All\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Explicit Proxy" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Explicit Proxy' is not allowed in import IDs. Please use 'Mobile Users Explicit Proxy' instead.\nExample: terraform import scm_service_connection.example \"Mobile Users Explicit Proxy\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Access Agent" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Access Agent' is not allowed in import IDs. Please use 'Mobile Users' instead.\nExample: terraform import scm_service_connection.example \"Mobile Users\":::\"<id>\"",
			)
			return
		}
	}

	// If validation passes, store the import ID in tfid
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
