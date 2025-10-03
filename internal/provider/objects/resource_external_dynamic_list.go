package provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// RESOURCE for SCM ExternalDynamicList (Package: objects)
var (
	_ resource.Resource                = &ExternalDynamicListResource{}
	_ resource.ResourceWithConfigure   = &ExternalDynamicListResource{}
	_ resource.ResourceWithImportState = &ExternalDynamicListResource{}
)

func NewExternalDynamicListResource() resource.Resource {
	return &ExternalDynamicListResource{}
}

// ExternalDynamicListResource defines the resource implementation.
type ExternalDynamicListResource struct {
	client *objects.APIClient
}

func (r *ExternalDynamicListResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_dynamic_list"
}

func (r *ExternalDynamicListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.ExternalDynamicListsResourceSchema
}

func (r *ExternalDynamicListResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["objects"].(*objects.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *objects.APIClient for 'objects' client."))
		return
	}
	r.client = client
}

func (r *ExternalDynamicListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for ExternalDynamicList")
	var data models.ExternalDynamicLists

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a patcher to temporarily store sensitive values.
	patcher := &externalDynamicListsSensitiveValuePatcher{}

	// Stash plaintext values from the plan.

	{ // Stash plaintext for Type.Ip.Auth.Password
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Type.IsNull() && !data.Type.IsUnknown() {
				var temp_stash_Type_Ip_Auth_Password_0 models.ExternalDynamicListsType
				resp.Diagnostics.Append(data.Type.As(ctx, &temp_stash_Type_Ip_Auth_Password_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_Type_Ip_Auth_Password_0.Ip.IsNull() && !temp_stash_Type_Ip_Auth_Password_0.Ip.IsUnknown() {
						var temp_stash_Type_Ip_Auth_Password_1 models.ExternalDynamicListsTypeIp
						resp.Diagnostics.Append(temp_stash_Type_Ip_Auth_Password_0.Ip.As(ctx, &temp_stash_Type_Ip_Auth_Password_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_Type_Ip_Auth_Password_1.Auth.IsNull() && !temp_stash_Type_Ip_Auth_Password_1.Auth.IsUnknown() {
								var temp_stash_Type_Ip_Auth_Password_2 models.ExternalDynamicListsTypeIpAuth
								resp.Diagnostics.Append(temp_stash_Type_Ip_Auth_Password_1.Auth.As(ctx, &temp_stash_Type_Ip_Auth_Password_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									// Innermost block
									finalVal = temp_stash_Type_Ip_Auth_Password_2.Password

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.type_ip_auth_password_plaintext = finalVal
		}
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.ExternalDynamicLists{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackExternalDynamicListsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating external_dynamic_lists on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.ExternalDynamicListsAPI.CreateExternalDynamicLists(ctx).ExternalDynamicLists(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating external_dynamic_lists", err.Error())
		return
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packExternalDynamicListsFromSdk(ctx, *createdObject)
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

	// Stash the encrypted values from the API response and apply the patch.

	{ // Patch plaintext for Type.Ip.Auth.Password
		if !resp.Diagnostics.HasError() {

			if !data.Type.IsNull() && !data.Type.IsUnknown() {
				var temp_patch_create_Type_Ip_Auth_Password_0 models.ExternalDynamicListsType
				resp.Diagnostics.Append(data.Type.As(ctx, &temp_patch_create_Type_Ip_Auth_Password_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_Type_Ip_Auth_Password_0.Ip.IsNull() && !temp_patch_create_Type_Ip_Auth_Password_0.Ip.IsUnknown() {
						var temp_patch_create_Type_Ip_Auth_Password_1 models.ExternalDynamicListsTypeIp
						resp.Diagnostics.Append(temp_patch_create_Type_Ip_Auth_Password_0.Ip.As(ctx, &temp_patch_create_Type_Ip_Auth_Password_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_create_Type_Ip_Auth_Password_1.Auth.IsNull() && !temp_patch_create_Type_Ip_Auth_Password_1.Auth.IsUnknown() {
								var temp_patch_create_Type_Ip_Auth_Password_2 models.ExternalDynamicListsTypeIpAuth
								resp.Diagnostics.Append(temp_patch_create_Type_Ip_Auth_Password_1.Auth.As(ctx, &temp_patch_create_Type_Ip_Auth_Password_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									// Innermost block
									patcher.type_ip_auth_password_encrypted = temp_patch_create_Type_Ip_Auth_Password_2.Password
									temp_patch_create_Type_Ip_Auth_Password_2.Password = patcher.type_ip_auth_password_plaintext

									// Repack the modified structs.

									if !resp.Diagnostics.HasError() {

										temp_patch_create_Type_Ip_Auth_Password_1.Auth, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpAuth{}.AttrTypes(), &temp_patch_create_Type_Ip_Auth_Password_2)

										resp.Diagnostics.Append(diags...)
									}

									if !resp.Diagnostics.HasError() {

										temp_patch_create_Type_Ip_Auth_Password_0.Ip, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIp{}.AttrTypes(), &temp_patch_create_Type_Ip_Auth_Password_1)

										resp.Diagnostics.Append(diags...)
									}

									if !resp.Diagnostics.HasError() {

										data.Type, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsType{}.AttrTypes(), &temp_patch_create_Type_Ip_Auth_Password_0)

										resp.Diagnostics.Append(diags...)
									}

								}
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

	if f := v.FieldByName("Snippet"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

	idBuilder.WriteString(":")

	if f := v.FieldByName("Device"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	tflog.Debug(ctx, "Created external_dynamic_lists", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExternalDynamicListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.ExternalDynamicLists - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for ExternalDynamicList")
	var data, savestate models.ExternalDynamicLists

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
	tflog.Debug(ctx, "Reading external_dynamic_lists from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.ExternalDynamicListsAPI.GetExternalDynamicListsByID(ctx, objectId)
	scmObject, _, err := getReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading external_dynamic_lists", err.Error())
		return
	}

	// Step 4 - Encrypted values logic
	// Populate a patcher from the state's encrypted_values map.
	patcher := &externalDynamicListsSensitiveValuePatcher{}
	resp.Diagnostics.Append(patcher.populatePatcherFromState(ctx, savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packExternalDynamicListsFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 6 - Encrypted values logic
	// Check for out-of-band changes and apply the patch.

	{ // Patch plaintext for Type.Ip.Auth.Password
		if !resp.Diagnostics.HasError() {

			if !data.Type.IsNull() && !data.Type.IsUnknown() {
				var temp_patch_read_Type_Ip_Auth_Password_0 models.ExternalDynamicListsType
				resp.Diagnostics.Append(data.Type.As(ctx, &temp_patch_read_Type_Ip_Auth_Password_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_Type_Ip_Auth_Password_0.Ip.IsNull() && !temp_patch_read_Type_Ip_Auth_Password_0.Ip.IsUnknown() {
						var temp_patch_read_Type_Ip_Auth_Password_1 models.ExternalDynamicListsTypeIp
						resp.Diagnostics.Append(temp_patch_read_Type_Ip_Auth_Password_0.Ip.As(ctx, &temp_patch_read_Type_Ip_Auth_Password_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_read_Type_Ip_Auth_Password_1.Auth.IsNull() && !temp_patch_read_Type_Ip_Auth_Password_1.Auth.IsUnknown() {
								var temp_patch_read_Type_Ip_Auth_Password_2 models.ExternalDynamicListsTypeIpAuth
								resp.Diagnostics.Append(temp_patch_read_Type_Ip_Auth_Password_1.Auth.As(ctx, &temp_patch_read_Type_Ip_Auth_Password_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									// Innermost block: Perform comparison and patch.
									if patcher.type_ip_auth_password_encrypted.Equal(temp_patch_read_Type_Ip_Auth_Password_2.Password) {
										temp_patch_read_Type_Ip_Auth_Password_2.Password = patcher.type_ip_auth_password_plaintext
									} else {
										temp_patch_read_Type_Ip_Auth_Password_2.Password = basetypes.NewStringNull()
									}

									// Repack the modified structs.

									if !resp.Diagnostics.HasError() {

										temp_patch_read_Type_Ip_Auth_Password_1.Auth, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpAuth{}.AttrTypes(), &temp_patch_read_Type_Ip_Auth_Password_2)

										resp.Diagnostics.Append(diags...)
									}

									if !resp.Diagnostics.HasError() {

										temp_patch_read_Type_Ip_Auth_Password_0.Ip, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIp{}.AttrTypes(), &temp_patch_read_Type_Ip_Auth_Password_1)

										resp.Diagnostics.Append(diags...)
									}

									if !resp.Diagnostics.HasError() {

										data.Type, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsType{}.AttrTypes(), &temp_patch_read_Type_Ip_Auth_Password_0)

										resp.Diagnostics.Append(diags...)
									}

								}
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

	if tokens[0] != "" {
		data.Folder = basetypes.NewStringValue(tokens[0])
	} else {
		data.Folder = basetypes.NewStringNull()
	}

	if tokens[1] != "" {
		data.Snippet = basetypes.NewStringValue(tokens[1])
	} else {
		data.Snippet = basetypes.NewStringNull()
	}

	if tokens[2] != "" {
		data.Device = basetypes.NewStringValue(tokens[2])
	} else {
		data.Device = basetypes.NewStringNull()
	}

	// Step 10 - Set data back into tf state and done
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExternalDynamicListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.ExternalDynamicLists which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for ExternalDynamicList")
	var plan, state models.ExternalDynamicLists

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
	patcher := &externalDynamicListsSensitiveValuePatcher{}

	{ // Stash plaintext for Type.Ip.Auth.Password
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
				var temp_stash_upd_Type_Ip_Auth_Password_0 models.ExternalDynamicListsType
				resp.Diagnostics.Append(plan.Type.As(ctx, &temp_stash_upd_Type_Ip_Auth_Password_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_Type_Ip_Auth_Password_0.Ip.IsNull() && !temp_stash_upd_Type_Ip_Auth_Password_0.Ip.IsUnknown() {
						var temp_stash_upd_Type_Ip_Auth_Password_1 models.ExternalDynamicListsTypeIp
						resp.Diagnostics.Append(temp_stash_upd_Type_Ip_Auth_Password_0.Ip.As(ctx, &temp_stash_upd_Type_Ip_Auth_Password_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_upd_Type_Ip_Auth_Password_1.Auth.IsNull() && !temp_stash_upd_Type_Ip_Auth_Password_1.Auth.IsUnknown() {
								var temp_stash_upd_Type_Ip_Auth_Password_2 models.ExternalDynamicListsTypeIpAuth
								resp.Diagnostics.Append(temp_stash_upd_Type_Ip_Auth_Password_1.Auth.As(ctx, &temp_stash_upd_Type_Ip_Auth_Password_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									// Innermost block
									finalVal = temp_stash_upd_Type_Ip_Auth_Password_2.Password

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.type_ip_auth_password_plaintext = finalVal
		}
	}

	// Step 3: Creates a plan object from the plan
	planObject, diags := types.ObjectValueFrom(ctx, models.ExternalDynamicLists{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackExternalDynamicListsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5: Update calls cannot have id sent in payload, so remove it
	// ID is a pointer, so we nil it out to omit it from the update payload.
	unpackedScmObject.Id = nil

	// Step 6: Get id from token and make update call
	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Updating external_dynamic_lists on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.ExternalDynamicListsAPI.UpdateExternalDynamicListsByID(ctx, objectId).ExternalDynamicLists(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, _, err := updateReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating external_dynamic_lists", err.Error())
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packExternalDynamicListsFromSdk(ctx, *updatedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// // Preserve any write-only parameter values from the plan.
	//
	// _ = req.Plan.GetAttribute(ctx, path.Root("id"), &plan.Id)
	//

	// Step 10: Encrypted values logic

	{ // Patch plaintext for Type.Ip.Auth.Password
		if !resp.Diagnostics.HasError() {

			if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
				var temp_patch_upd_Type_Ip_Auth_Password_0 models.ExternalDynamicListsType
				resp.Diagnostics.Append(plan.Type.As(ctx, &temp_patch_upd_Type_Ip_Auth_Password_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_Type_Ip_Auth_Password_0.Ip.IsNull() && !temp_patch_upd_Type_Ip_Auth_Password_0.Ip.IsUnknown() {
						var temp_patch_upd_Type_Ip_Auth_Password_1 models.ExternalDynamicListsTypeIp
						resp.Diagnostics.Append(temp_patch_upd_Type_Ip_Auth_Password_0.Ip.As(ctx, &temp_patch_upd_Type_Ip_Auth_Password_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_upd_Type_Ip_Auth_Password_1.Auth.IsNull() && !temp_patch_upd_Type_Ip_Auth_Password_1.Auth.IsUnknown() {
								var temp_patch_upd_Type_Ip_Auth_Password_2 models.ExternalDynamicListsTypeIpAuth
								resp.Diagnostics.Append(temp_patch_upd_Type_Ip_Auth_Password_1.Auth.As(ctx, &temp_patch_upd_Type_Ip_Auth_Password_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									// Innermost block
									patcher.type_ip_auth_password_encrypted = temp_patch_upd_Type_Ip_Auth_Password_2.Password
									temp_patch_upd_Type_Ip_Auth_Password_2.Password = patcher.type_ip_auth_password_plaintext

									// Repack the modified structs.

									if !resp.Diagnostics.HasError() {

										temp_patch_upd_Type_Ip_Auth_Password_1.Auth, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpAuth{}.AttrTypes(), &temp_patch_upd_Type_Ip_Auth_Password_2)

										resp.Diagnostics.Append(diags...)
									}

									if !resp.Diagnostics.HasError() {

										temp_patch_upd_Type_Ip_Auth_Password_0.Ip, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIp{}.AttrTypes(), &temp_patch_upd_Type_Ip_Auth_Password_1)

										resp.Diagnostics.Append(diags...)
									}

									if !resp.Diagnostics.HasError() {

										plan.Type, diags = types.ObjectValueFrom(ctx, models.ExternalDynamicListsType{}.AttrTypes(), &temp_patch_upd_Type_Ip_Auth_Password_0)

										resp.Diagnostics.Append(diags...)
									}

								}
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

	tflog.Debug(ctx, "Updated external_dynamic_lists", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ExternalDynamicListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.ExternalDynamicLists
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

	tflog.Debug(ctx, "Deleting external_dynamic_lists", map[string]interface{}{"id": objectId})
	deleteReq := r.client.ExternalDynamicListsAPI.DeleteExternalDynamicListsByID(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting external_dynamic_lists", err.Error())
	}
}

func (r *ExternalDynamicListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
