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

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// RESOURCE for SCM MfaServer (Package: identity_services)
var (
	_ resource.Resource                = &MfaServerResource{}
	_ resource.ResourceWithConfigure   = &MfaServerResource{}
	_ resource.ResourceWithImportState = &MfaServerResource{}
)

func NewMfaServerResource() resource.Resource {
	return &MfaServerResource{}
}

// MfaServerResource defines the resource implementation.
type MfaServerResource struct {
	client *identity_services.APIClient
}

func (r *MfaServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mfa_server"
}

func (r *MfaServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.MfaServersResourceSchema
}

func (r *MfaServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["identity_services"].(*identity_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *identity_services.APIClient for 'identity_services' client."))
		return
	}
	r.client = client
}

func (r *MfaServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for MfaServer")
	var data models.MfaServers

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a patcher to temporarily store sensitive values.
	patcher := &mfaServersSensitiveValuePatcher{}

	// Stash plaintext values from the plan.

	{ // Stash plaintext for MfaVendorType.DuoSecurityV2.DuoSecretKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsNull() && !temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsUnknown() {
						var temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_1 models.MfaServersMfaVendorTypeDuoSecurityV2
						resp.Diagnostics.Append(temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.As(ctx, &temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for MfaVendorType.OktaAdaptiveV1.OktaToken
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsNull() && !temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsUnknown() {
						var temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_1 models.MfaServersMfaVendorTypeOktaAdaptiveV1
						resp.Diagnostics.Append(temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.As(ctx, &temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext = finalVal
		}
	}

	{ // Stash plaintext for MfaVendorType.PingIdentityV1.PingUseBase64Key
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsNull() && !temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsUnknown() {
						var temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_1 models.MfaServersMfaVendorTypePingIdentityV1
						resp.Diagnostics.Append(temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.As(ctx, &temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for MfaVendorType.RsaSecuridAccessV1.RsaAccesskey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsNull() && !temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsUnknown() {
						var temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1 models.MfaServersMfaVendorTypeRsaSecuridAccessV1
						resp.Diagnostics.Append(temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.As(ctx, &temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext = finalVal
		}
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.MfaServers{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackMfaServersToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mfa_servers on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.MFAServersAPI.CreateMFAServers(ctx).MfaServers(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating mfa_servers", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packMfaServersFromSdk(ctx, *createdObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Normalize null lists from API response to match the plan
	packedObject, diags = utils.NormalizeNullLists(ctx, packedObject, planObject)
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

	{ // Patch plaintext for MfaVendorType.DuoSecurityV2.DuoSecretKey
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsNull() && !temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsUnknown() {
						var temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_1 models.MfaServersMfaVendorTypeDuoSecurityV2
						resp.Diagnostics.Append(temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.As(ctx, &temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted = temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey
							temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey = patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrTypes(), &temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_create_MfaVendorType_DuoSecurityV2_DuoSecretKey_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.OktaAdaptiveV1.OktaToken
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsNull() && !temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsUnknown() {
						var temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_1 models.MfaServersMfaVendorTypeOktaAdaptiveV1
						resp.Diagnostics.Append(temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.As(ctx, &temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted = temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken
							temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken = patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrTypes(), &temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_create_MfaVendorType_OktaAdaptiveV1_OktaToken_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.PingIdentityV1.PingUseBase64Key
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsNull() && !temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsUnknown() {
						var temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_1 models.MfaServersMfaVendorTypePingIdentityV1
						resp.Diagnostics.Append(temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.As(ctx, &temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted = temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key
							temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key = patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypePingIdentityV1{}.AttrTypes(), &temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_create_MfaVendorType_PingIdentityV1_PingUseBase64Key_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.RsaSecuridAccessV1.RsaAccesskey
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsNull() && !temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsUnknown() {
						var temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1 models.MfaServersMfaVendorTypeRsaSecuridAccessV1
						resp.Diagnostics.Append(temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.As(ctx, &temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted = temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey
							temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey = patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrTypes(), &temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_create_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0)

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

	tflog.Debug(ctx, "Created mfa_servers", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MfaServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.MfaServers - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for MfaServer")
	var data, savestate models.MfaServers

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
	tflog.Debug(ctx, "Reading mfa_servers from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.MFAServersAPI.GetMFAServersByID(ctx, objectId)
	scmObject, httpErr, err := getReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no mfa_servers on read SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on read SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error reading mfa_servers", err.Error())
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
	patcher := &mfaServersSensitiveValuePatcher{}
	resp.Diagnostics.Append(patcher.populatePatcherFromState(ctx, savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packMfaServersFromSdk(ctx, *scmObject)
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

	{ // Patch plaintext for MfaVendorType.DuoSecurityV2.DuoSecretKey
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsNull() && !temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsUnknown() {
						var temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_1 models.MfaServersMfaVendorTypeDuoSecurityV2
						resp.Diagnostics.Append(temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.As(ctx, &temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block: Perform comparison and patch.
							if patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted.Equal(temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey) {
								temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey = patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext
							} else {
								temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey = basetypes.NewStringNull()
							}

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrTypes(), &temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_read_MfaVendorType_DuoSecurityV2_DuoSecretKey_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.OktaAdaptiveV1.OktaToken
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsNull() && !temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsUnknown() {
						var temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_1 models.MfaServersMfaVendorTypeOktaAdaptiveV1
						resp.Diagnostics.Append(temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.As(ctx, &temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block: Perform comparison and patch.
							if patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted.Equal(temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken) {
								temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken = patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext
							} else {
								temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken = basetypes.NewStringNull()
							}

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrTypes(), &temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_read_MfaVendorType_OktaAdaptiveV1_OktaToken_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.PingIdentityV1.PingUseBase64Key
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsNull() && !temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsUnknown() {
						var temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_1 models.MfaServersMfaVendorTypePingIdentityV1
						resp.Diagnostics.Append(temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.As(ctx, &temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block: Perform comparison and patch.
							if patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted.Equal(temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key) {
								temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key = patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext
							} else {
								temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key = basetypes.NewStringNull()
							}

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypePingIdentityV1{}.AttrTypes(), &temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_read_MfaVendorType_PingIdentityV1_PingUseBase64Key_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.RsaSecuridAccessV1.RsaAccesskey
		if !resp.Diagnostics.HasError() {

			if !data.MfaVendorType.IsNull() && !data.MfaVendorType.IsUnknown() {
				var temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(data.MfaVendorType.As(ctx, &temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsNull() && !temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsUnknown() {
						var temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1 models.MfaServersMfaVendorTypeRsaSecuridAccessV1
						resp.Diagnostics.Append(temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.As(ctx, &temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block: Perform comparison and patch.
							if patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted.Equal(temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey) {
								temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey = patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext
							} else {
								temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey = basetypes.NewStringNull()
							}

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrTypes(), &temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								data.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_read_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0)

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

	// Use reflection to safely restore the Snippet field from the TFID token 1.
	vSnippet := reflect.ValueOf(&data).Elem()   // Unique variable: vSnippet
	fSnippet := vSnippet.FieldByName("Snippet") // Unique variable: fSnippet

	if fSnippet.IsValid() && fSnippet.CanSet() {
		tokenValue := tokens[1]

		if tokenValue != "" {
			newStringValue := basetypes.NewStringValue(tokenValue)
			fSnippet.Set(reflect.ValueOf(newStringValue))
		} else {
			newNullValue := basetypes.NewStringNull()
			fSnippet.Set(reflect.ValueOf(newNullValue))
		}
	}

	// --- DEVICE RESTORATION (tokens[2]) ---

	// Use reflection to safely restore the Device field from the TFID token 2.
	vDevice := reflect.ValueOf(&data).Elem() // Unique variable: vDevice
	fDevice := vDevice.FieldByName("Device") // Unique variable: fDevice

	if fDevice.IsValid() && fDevice.CanSet() {
		tokenValue := tokens[2]

		if tokenValue != "" {
			newStringValue := basetypes.NewStringValue(tokenValue)
			fDevice.Set(reflect.ValueOf(newStringValue))
		} else {
			newNullValue := basetypes.NewStringNull()
			fDevice.Set(reflect.ValueOf(newNullValue))
		}
	}

	// Step 10 - Set data back into tf state and done
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MfaServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.MfaServers which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for MfaServer")
	var plan, state models.MfaServers

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
	patcher := &mfaServersSensitiveValuePatcher{}

	{ // Stash plaintext for MfaVendorType.DuoSecurityV2.DuoSecretKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsNull() && !temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsUnknown() {
						var temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1 models.MfaServersMfaVendorTypeDuoSecurityV2
						resp.Diagnostics.Append(temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.As(ctx, &temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for MfaVendorType.OktaAdaptiveV1.OktaToken
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsNull() && !temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsUnknown() {
						var temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1 models.MfaServersMfaVendorTypeOktaAdaptiveV1
						resp.Diagnostics.Append(temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.As(ctx, &temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext = finalVal
		}
	}

	{ // Stash plaintext for MfaVendorType.PingIdentityV1.PingUseBase64Key
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsNull() && !temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsUnknown() {
						var temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1 models.MfaServersMfaVendorTypePingIdentityV1
						resp.Diagnostics.Append(temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.As(ctx, &temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for MfaVendorType.RsaSecuridAccessV1.RsaAccesskey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsNull() && !temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsUnknown() {
						var temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1 models.MfaServersMfaVendorTypeRsaSecuridAccessV1
						resp.Diagnostics.Append(temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.As(ctx, &temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							finalVal = temp_stash_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext = finalVal
		}
	}

	// Step 3: Creates a plan object from the plan
	planObject, diags := types.ObjectValueFrom(ctx, models.MfaServers{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackMfaServersToSdk(ctx, planObject)
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

	tflog.Debug(ctx, "Updating mfa_servers on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.MFAServersAPI.UpdateMFAServersByID(ctx, objectId).MfaServers(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no mfa_servers on update SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on update SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error updating mfa_servers", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Update Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packMfaServersFromSdk(ctx, *updatedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Normalize null lists from API response to match the plan
	packedObject, diags = utils.NormalizeNullLists(ctx, packedObject, planObject)
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

	{ // Patch plaintext for MfaVendorType.DuoSecurityV2.DuoSecretKey
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsNull() && !temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.IsUnknown() {
						var temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1 models.MfaServersMfaVendorTypeDuoSecurityV2
						resp.Diagnostics.Append(temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2.As(ctx, &temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted = temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey
							temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1.DuoSecretKey = patcher.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0.DuoSecurityV2, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrTypes(), &temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								plan.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_upd_MfaVendorType_DuoSecurityV2_DuoSecretKey_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.OktaAdaptiveV1.OktaToken
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsNull() && !temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.IsUnknown() {
						var temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1 models.MfaServersMfaVendorTypeOktaAdaptiveV1
						resp.Diagnostics.Append(temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1.As(ctx, &temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted = temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken
							temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1.OktaToken = patcher.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0.OktaAdaptiveV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrTypes(), &temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								plan.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_upd_MfaVendorType_OktaAdaptiveV1_OktaToken_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.PingIdentityV1.PingUseBase64Key
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsNull() && !temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.IsUnknown() {
						var temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1 models.MfaServersMfaVendorTypePingIdentityV1
						resp.Diagnostics.Append(temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1.As(ctx, &temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted = temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key
							temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1.PingUseBase64Key = patcher.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0.PingIdentityV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypePingIdentityV1{}.AttrTypes(), &temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								plan.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_upd_MfaVendorType_PingIdentityV1_PingUseBase64Key_0)

								resp.Diagnostics.Append(diags...)
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for MfaVendorType.RsaSecuridAccessV1.RsaAccesskey
		if !resp.Diagnostics.HasError() {

			if !plan.MfaVendorType.IsNull() && !plan.MfaVendorType.IsUnknown() {
				var temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0 models.MfaServersMfaVendorType
				resp.Diagnostics.Append(plan.MfaVendorType.As(ctx, &temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsNull() && !temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.IsUnknown() {
						var temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1 models.MfaServersMfaVendorTypeRsaSecuridAccessV1
						resp.Diagnostics.Append(temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1.As(ctx, &temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							// Innermost block
							patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted = temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey
							temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1.RsaAccesskey = patcher.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext

							// Repack the modified structs.

							if !resp.Diagnostics.HasError() {

								temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0.RsaSecuridAccessV1, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrTypes(), &temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_1)

								resp.Diagnostics.Append(diags...)
							}

							if !resp.Diagnostics.HasError() {

								plan.MfaVendorType, diags = types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &temp_patch_upd_MfaVendorType_RsaSecuridAccessV1_RsaAccesskey_0)

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

	tflog.Debug(ctx, "Updated mfa_servers", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MfaServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.MfaServers
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

	tflog.Debug(ctx, "Deleting mfa_servers", map[string]interface{}{"id": objectId})
	deleteReq := r.client.MFAServersAPI.DeleteMFAServersByID(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting mfa_servers", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Deleteion Failed: API Request Failed",
			detailedMessage,
		)
	}
}

func (r *MfaServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Parse and validate the import ID before storing it - block deprecated folder values
	// Note: Both "Shared" and "Prisma Access" are accepted as valid values
	importID := req.ID
	tokens := strings.Split(importID, ":::")

	if len(tokens) > 0 && tokens[0] != "" {
		folderValue := tokens[0]
		if folderValue == "All Firewalls" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'All Firewalls' is not allowed in import IDs. Please use 'ngfw-shared' instead.\nExample: terraform import scm_mfa_server.example \"ngfw-shared\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Global" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Global' is not allowed in import IDs. Please use 'All' instead.\nExample: terraform import scm_mfa_server.example \"All\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Explicit Proxy" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Explicit Proxy' is not allowed in import IDs. Please use 'Mobile Users Explicit Proxy' instead.\nExample: terraform import scm_mfa_server.example \"Mobile Users Explicit Proxy\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Access Agent" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Access Agent' is not allowed in import IDs. Please use 'Mobile Users' instead.\nExample: terraform import scm_mfa_server.example \"Mobile Users\":::\"<id>\"",
			)
			return
		}
	}

	// If validation passes, store the import ID in tfid
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
