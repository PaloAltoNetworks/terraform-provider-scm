package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	ebBZfXA "github.com/paloaltonetworks/scm-go/netsec/services/trustedcertificateauthorities"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Data source (listing).
var (
	_ datasource.DataSource              = &trustedCertificateAuthorityListDataSource{}
	_ datasource.DataSourceWithConfigure = &trustedCertificateAuthorityListDataSource{}
)

func NewTrustedCertificateAuthorityListDataSource() datasource.DataSource {
	return &trustedCertificateAuthorityListDataSource{}
}

type trustedCertificateAuthorityListDataSource struct {
	client *scm.Client
}

// trustedCertificateAuthorityListDsModel is the model.
type trustedCertificateAuthorityListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Device  types.String `tfsdk:"device"`
	Folder  types.String `tfsdk:"folder"`
	Limit   types.Int64  `tfsdk:"limit"`
	Name    types.String `tfsdk:"name"`
	Offset  types.Int64  `tfsdk:"offset"`
	Snippet types.String `tfsdk:"snippet"`

	// Output.
	Data []trustedCertificateAuthorityListDsModel_iYmUVvF_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type trustedCertificateAuthorityListDsModel_iYmUVvF_Config struct {
	CommonName     types.String `tfsdk:"common_name"`
	ExpiryEpoch    types.String `tfsdk:"expiry_epoch"`
	Filename       types.String `tfsdk:"filename"`
	Id             types.String `tfsdk:"id"`
	Issuer         types.String `tfsdk:"issuer"`
	Name           types.String `tfsdk:"name"`
	NotValidAfter  types.String `tfsdk:"not_valid_after"`
	NotValidBefore types.String `tfsdk:"not_valid_before"`
	SerialNumber   types.String `tfsdk:"serial_number"`
	Subject        types.String `tfsdk:"subject"`
}

// Metadata returns the data source type name.
func (d *trustedCertificateAuthorityListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_trusted_certificate_authority_list"
}

// Schema defines the schema for this listing data source.
func (d *trustedCertificateAuthorityListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"common_name":true, "expiry_epoch":true, "filename":true, "id":true, "issuer":true, "name":true, "not_valid_after":true, "not_valid_before":true, "serial_number":true, "subject":true} forceNew:map[string]bool(nil)
						"common_name": dsschema.StringAttribute{
							Description: "The CommonName param. String length must not exceed 255 characters.",
							Computed:    true,
						},
						"expiry_epoch": dsschema.StringAttribute{
							Description: "The ExpiryEpoch param.",
							Computed:    true,
						},
						"filename": dsschema.StringAttribute{
							Description: "The Filename param.",
							Computed:    true,
						},
						"id": dsschema.StringAttribute{
							Description: "UUID of the resource.",
							Computed:    true,
						},
						"issuer": dsschema.StringAttribute{
							Description: "The Issuer param.",
							Computed:    true,
						},
						"name": dsschema.StringAttribute{
							Description: "The Name param. String length must not exceed 63 characters.",
							Computed:    true,
						},
						"not_valid_after": dsschema.StringAttribute{
							Description: "The NotValidAfter param.",
							Computed:    true,
						},
						"not_valid_before": dsschema.StringAttribute{
							Description: "The NotValidBefore param.",
							Computed:    true,
						},
						"serial_number": dsschema.StringAttribute{
							Description: "The SerialNumber param.",
							Computed:    true,
						},
						"subject": dsschema.StringAttribute{
							Description: "The Subject param.",
							Computed:    true,
						},
					},
				},
			},
			"device": dsschema.StringAttribute{
				Description: "The Device param.",
				Optional:    true,
			},
			"folder": dsschema.StringAttribute{
				Description: "The Folder param.",
				Optional:    true,
			},
			"limit": dsschema.Int64Attribute{
				Description: "The Limit param. A limit of -1 will return all configured items. Default: `200`.",
				Optional:    true,
				Computed:    true,
			},
			"name": dsschema.StringAttribute{
				Description: "The Name param.",
				Optional:    true,
			},
			"offset": dsschema.Int64Attribute{
				Description: "The Offset param. Default: `0`.",
				Optional:    true,
				Computed:    true,
			},
			"snippet": dsschema.StringAttribute{
				Description: "The Snippet param.",
				Optional:    true,
			},
			"tfid": dsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
			},
			"total": dsschema.Int64Attribute{
				Description: "The Total param.",
				Computed:    true,
			},
		},
	}
}

// Configure prepares the struct.
func (d *trustedCertificateAuthorityListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *trustedCertificateAuthorityListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state trustedCertificateAuthorityListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_trusted_certificate_authority_list",
		"terraform_provider_function": "Read",
		"name":                        state.Name.ValueString(),
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
		"limit":                       state.Limit.ValueInt64(),
		"offset":                      state.Offset.ValueInt64(),
	})

	// Prepare to run the command.
	svc := ebBZfXA.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := ebBZfXA.ListInput{}

	input.Name = state.Name.ValueStringPointer()

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()

	input.Limit = state.Limit.ValueInt64Pointer()

	input.Offset = state.Offset.ValueInt64Pointer()

	// Perform the operation.
	ans, err := svc.List(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError("Error getting listing", err.Error())
		return
	}

	// Create the Terraform ID.
	var idBuilder strings.Builder
	if input.Name != nil {
		idBuilder.WriteString(*input.Name)
	}

	idBuilder.WriteString(IdSeparator)
	if input.Folder != nil {
		idBuilder.WriteString(*input.Folder)
	}

	idBuilder.WriteString(IdSeparator)
	if input.Snippet != nil {
		idBuilder.WriteString(*input.Snippet)
	}

	idBuilder.WriteString(IdSeparator)
	if input.Device != nil {
		idBuilder.WriteString(*input.Device)
	}

	idBuilder.WriteString(IdSeparator)
	if input.Limit != nil {
		idBuilder.WriteString(strconv.FormatInt(*input.Limit, 10))
	}

	idBuilder.WriteString(IdSeparator)
	if input.Offset != nil {
		idBuilder.WriteString(strconv.FormatInt(*input.Offset, 10))
	}

	// Store the answer to state.

	state.Tfid = types.StringValue(idBuilder.String())

	if len(ans.Data) == 0 {
		state.Data = nil
	} else {
		state.Data = make([]trustedCertificateAuthorityListDsModel_iYmUVvF_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := trustedCertificateAuthorityListDsModel_iYmUVvF_Config{}

			var1.CommonName = types.StringPointerValue(var0.CommonName)

			var1.ExpiryEpoch = types.StringPointerValue(var0.ExpiryEpoch)

			var1.Filename = types.StringPointerValue(var0.Filename)

			var1.Id = types.StringPointerValue(var0.Id)

			var1.Issuer = types.StringPointerValue(var0.Issuer)

			var1.Name = types.StringPointerValue(var0.Name)

			var1.NotValidAfter = types.StringPointerValue(var0.NotValidAfter)

			var1.NotValidBefore = types.StringPointerValue(var0.NotValidBefore)

			var1.SerialNumber = types.StringPointerValue(var0.SerialNumber)

			var1.Subject = types.StringPointerValue(var0.Subject)
			state.Data = append(state.Data, var1)
		}
	}

	state.Limit = types.Int64PointerValue(ans.Limit)

	state.Offset = types.Int64PointerValue(ans.Offset)

	state.Total = types.Int64PointerValue(ans.Total)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
