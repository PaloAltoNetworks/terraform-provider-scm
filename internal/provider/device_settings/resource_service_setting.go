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

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// RESOURCE for SCM ServiceSetting (Package: device_settings)
var (
	_ resource.Resource                = &ServiceSettingResource{}
	_ resource.ResourceWithConfigure   = &ServiceSettingResource{}
	_ resource.ResourceWithImportState = &ServiceSettingResource{}
)

func NewServiceSettingResource() resource.Resource {
	return &ServiceSettingResource{}
}

// ServiceSettingResource defines the resource implementation.
type ServiceSettingResource struct {
	client *device_settings.APIClient
}

func (r *ServiceSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_setting"
}

func (r *ServiceSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.ServiceSettingsResourceSchema
}

func (r *ServiceSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["device_settings"].(*device_settings.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *device_settings.APIClient for 'device_settings' client."))
		return
	}
	r.client = client
}

func (r *ServiceSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for ServiceSetting")
	var data models.ServiceSettings

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a patcher to temporarily store sensitive values.
	patcher := &serviceSettingsSensitiveValuePatcher{}

	// Stash plaintext values from the plan.

	{ // Stash plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.SecureProxyPassword
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_stash_Services_SecureProxyPassword_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_stash_Services_SecureProxyPassword_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					finalVal = temp_stash_Services_SecureProxyPassword_0.SecureProxyPassword

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_secure_proxy_password_plaintext = finalVal
		}
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.ServiceSettings{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackServiceSettingsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating service_settings on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.ServiceSettingsAPI.CreateServiceSettings(ctx).ServiceSettings(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating service_settings", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packServiceSettingsFromSdk(ctx, *createdObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 6a. Normalize null lists: when the API returns null for a list field that
	// the plan had as [] (empty list), coerce to empty list to satisfy Terraform's
	// consistency requirement. This recurses into nested objects and list elements.
	packedObject, diags = utils.NormalizeNullLists(ctx, planObject, packedObject)
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

	{ // Patch plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted = temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey
																	temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted = temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey
																	temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_create_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted = temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey
																	temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted = temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey
																	temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_create_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.SecureProxyPassword
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_create_Services_SecureProxyPassword_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_create_Services_SecureProxyPassword_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					patcher.services_secure_proxy_password_encrypted = temp_patch_create_Services_SecureProxyPassword_0.SecureProxyPassword
					temp_patch_create_Services_SecureProxyPassword_0.SecureProxyPassword = patcher.services_secure_proxy_password_plaintext

					// Repack the modified structs.

					if !resp.Diagnostics.HasError() {

						data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_create_Services_SecureProxyPassword_0)

						resp.Diagnostics.Append(diags...)
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

	tflog.Debug(ctx, "Created service_settings", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.ServiceSettings - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for ServiceSetting")
	var data, savestate models.ServiceSettings

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
	tflog.Debug(ctx, "Reading service_settings from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.ServiceSettingsAPI.GetServiceSettingsByID(ctx, objectId)
	scmObject, httpErr, err := getReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no service_settings on read SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on read SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error reading service_settings", err.Error())
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
	patcher := &serviceSettingsSensitiveValuePatcher{}
	resp.Diagnostics.Append(patcher.populatePatcherFromState(ctx, savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packServiceSettingsFromSdk(ctx, *scmObject)
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

	{ // Patch plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block: Perform comparison and patch.
																	if patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted.Equal(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey) {
																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext
																	} else {
																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = basetypes.NewStringNull()
																	}

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block: Perform comparison and patch.
																	if patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted.Equal(temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey) {
																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext
																	} else {
																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = basetypes.NewStringNull()
																	}

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_read_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block: Perform comparison and patch.
																	if patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted.Equal(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey) {
																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext
																	} else {
																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = basetypes.NewStringNull()
																	}

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block: Perform comparison and patch.
																	if patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted.Equal(temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey) {
																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext
																	} else {
																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = basetypes.NewStringNull()
																	}

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_read_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.SecureProxyPassword
		if !resp.Diagnostics.HasError() {

			if !data.Services.IsNull() && !data.Services.IsUnknown() {
				var temp_patch_read_Services_SecureProxyPassword_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(data.Services.As(ctx, &temp_patch_read_Services_SecureProxyPassword_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block: Perform comparison and patch.
					if patcher.services_secure_proxy_password_encrypted.Equal(temp_patch_read_Services_SecureProxyPassword_0.SecureProxyPassword) {
						temp_patch_read_Services_SecureProxyPassword_0.SecureProxyPassword = patcher.services_secure_proxy_password_plaintext
					} else {
						temp_patch_read_Services_SecureProxyPassword_0.SecureProxyPassword = basetypes.NewStringNull()
					}

					// Repack the modified structs.

					if !resp.Diagnostics.HasError() {

						data.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_read_Services_SecureProxyPassword_0)

						resp.Diagnostics.Append(diags...)
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

func (r *ServiceSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.ServiceSettings which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for ServiceSetting")
	var plan, state models.ServiceSettings

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
	patcher := &serviceSettingsSensitiveValuePatcher{}

	{ // Stash plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	finalVal = temp_stash_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey

																}
															}

														}
													}

												}
											}

										}
									}

								}
							}

						}
					}

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext = finalVal
		}
	}

	{ // Stash plaintext for Services.SecureProxyPassword
		var finalVal basetypes.StringValue
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_stash_upd_Services_SecureProxyPassword_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_stash_upd_Services_SecureProxyPassword_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					finalVal = temp_stash_upd_Services_SecureProxyPassword_0.SecureProxyPassword

				}
			}

		}
		if !resp.Diagnostics.HasError() && !finalVal.IsUnknown() && !finalVal.IsNull() {
			patcher.services_secure_proxy_password_plaintext = finalVal
		}
	}

	// Step 3: Creates a plan object from the plan
	planObject, diags := types.ObjectValueFrom(ctx, models.ServiceSettings{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackServiceSettingsToSdk(ctx, planObject)
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

	tflog.Debug(ctx, "Updating service_settings on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.ServiceSettingsAPI.UpdateServiceSettingsByID(ctx, objectId).ServiceSettings(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no service_settings on update SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on update SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error updating service_settings", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Update Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packServiceSettingsFromSdk(ctx, *updatedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 9a: Normalize null lists: when the API returns null for a list field that
	// the plan had as [] (empty list), coerce to empty list to satisfy Terraform's
	// consistency requirement. This recurses into nested objects and list elements.
	packedObject, diags = utils.NormalizeNullLists(ctx, planObject, packedObject)
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

	{ // Patch plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted = temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey
																	temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.PrimaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		plan.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.IsUnknown() {
								var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted = temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey
																	temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.PrimaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		plan.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_PrimaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.IsUnknown() {
																var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted = temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey
																	temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5.Md5, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1.SecondaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		plan.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Md5_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.IsUnknown() {
						var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1 models.ServiceSettingsServicesNtpServers
						resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1, basetypes.ObjectAsOptions{})...)
						if !resp.Diagnostics.HasError() {

							if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.IsUnknown() {
								var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2 models.ServiceSettingsServicesNtpServersPrimaryNtpServer
								resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2, basetypes.ObjectAsOptions{})...)
								if !resp.Diagnostics.HasError() {

									if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.IsUnknown() {
										var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
										resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3, basetypes.ObjectAsOptions{})...)
										if !resp.Diagnostics.HasError() {

											if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.IsUnknown() {
												var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
												resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4, basetypes.ObjectAsOptions{})...)
												if !resp.Diagnostics.HasError() {

													if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.IsUnknown() {
														var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
														resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5, basetypes.ObjectAsOptions{})...)
														if !resp.Diagnostics.HasError() {

															if !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsNull() && !temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.IsUnknown() {
																var temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6 models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
																resp.Diagnostics.Append(temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1.As(ctx, &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6, basetypes.ObjectAsOptions{})...)
																if !resp.Diagnostics.HasError() {

																	// Innermost block
																	patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted = temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey
																	temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6.AuthenticationKey = patcher.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext

																	// Repack the modified structs.

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5.Sha1, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_6)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4.Algorithm, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_5)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3.SymmetricKey, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_4)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2.AuthenticationType, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_3)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1.SecondaryNtpServer, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_2)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0.NtpServers, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_1)

																		resp.Diagnostics.Append(diags...)
																	}

																	if !resp.Diagnostics.HasError() {

																		plan.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_upd_Services_NtpServers_SecondaryNtpServer_AuthenticationType_SymmetricKey_Algorithm_Sha1_AuthenticationKey_0)

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

								}
							}

						}
					}

				}
			}

		}
	}

	{ // Patch plaintext for Services.SecureProxyPassword
		if !resp.Diagnostics.HasError() {

			if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
				var temp_patch_upd_Services_SecureProxyPassword_0 models.ServiceSettingsServices
				resp.Diagnostics.Append(plan.Services.As(ctx, &temp_patch_upd_Services_SecureProxyPassword_0, basetypes.ObjectAsOptions{})...)
				if !resp.Diagnostics.HasError() {

					// Innermost block
					patcher.services_secure_proxy_password_encrypted = temp_patch_upd_Services_SecureProxyPassword_0.SecureProxyPassword
					temp_patch_upd_Services_SecureProxyPassword_0.SecureProxyPassword = patcher.services_secure_proxy_password_plaintext

					// Repack the modified structs.

					if !resp.Diagnostics.HasError() {

						plan.Services, diags = types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &temp_patch_upd_Services_SecureProxyPassword_0)

						resp.Diagnostics.Append(diags...)
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

	tflog.Debug(ctx, "Updated service_settings", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ServiceSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.ServiceSettings
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

	tflog.Debug(ctx, "Deleting service_settings", map[string]interface{}{"id": objectId})
	deleteReq := r.client.ServiceSettingsAPI.DeleteServiceSettingsByID(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting service_settings", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Deleteion Failed: API Request Failed",
			detailedMessage,
		)
	}
}

func (r *ServiceSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Parse and validate the import ID before storing it - block deprecated folder values
	// Note: Both "Shared" and "Prisma Access" are accepted as valid values
	importID := req.ID
	tokens := strings.Split(importID, ":::")

	if len(tokens) > 0 && tokens[0] != "" {
		folderValue := tokens[0]
		if folderValue == "All Firewalls" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'All Firewalls' is not allowed in import IDs. Please use 'ngfw-shared' instead.\nExample: terraform import scm_service_setting.example \"ngfw-shared\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Global" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Global' is not allowed in import IDs. Please use 'All' instead.\nExample: terraform import scm_service_setting.example \"All\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Explicit Proxy" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Explicit Proxy' is not allowed in import IDs. Please use 'Mobile Users Explicit Proxy' instead.\nExample: terraform import scm_service_setting.example \"Mobile Users Explicit Proxy\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Access Agent" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Access Agent' is not allowed in import IDs. Please use 'Mobile Users' instead.\nExample: terraform import scm_service_setting.example \"Mobile Users\":::\"<id>\"",
			)
			return
		}
	}

	// If validation passes, store the import ID in tfid
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
