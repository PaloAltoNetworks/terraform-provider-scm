package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	hhIWLbI "github.com/paloaltonetworks/scm-go/netsec/schemas/authentication/portals"
	ljRBvXJ "github.com/paloaltonetworks/scm-go/netsec/services/authenticationportals"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Data source (listing).
var (
	_ datasource.DataSource              = &authenticationPortalListDataSource{}
	_ datasource.DataSourceWithConfigure = &authenticationPortalListDataSource{}
)

func NewAuthenticationPortalListDataSource() datasource.DataSource {
	return &authenticationPortalListDataSource{}
}

type authenticationPortalListDataSource struct {
	client *scm.Client
}

// authenticationPortalListDsModel is the model.
type authenticationPortalListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Device  types.String `tfsdk:"device"`
	Folder  types.String `tfsdk:"folder"`
	Limit   types.Int64  `tfsdk:"limit"`
	Name    types.String `tfsdk:"name"`
	Offset  types.Int64  `tfsdk:"offset"`
	Snippet types.String `tfsdk:"snippet"`

	// Output.
	Data []authenticationPortalListDsModel_hhIWLbI_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type authenticationPortalListDsModel_hhIWLbI_Config struct {
	AuthenticationProfile types.String `tfsdk:"authentication_profile"`
	CertificateProfile    types.String `tfsdk:"certificate_profile"`
	GpUdpPort             types.Int64  `tfsdk:"gp_udp_port"`
	Id                    types.String `tfsdk:"id"`
	IdleTimer             types.Int64  `tfsdk:"idle_timer"`
	RedirectHost          types.String `tfsdk:"redirect_host"`
	Timer                 types.Int64  `tfsdk:"timer"`
	TlsServiceProfile     types.String `tfsdk:"tls_service_profile"`
}

// Metadata returns the data source type name.
func (d *authenticationPortalListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_portal_list"
}

// Schema defines the schema for this listing data source.
func (d *authenticationPortalListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"authentication_profile":true, "certificate_profile":true, "gp_udp_port":true, "id":true, "idle_timer":true, "redirect_host":true, "timer":true, "tls_service_profile":true} forceNew:map[string]bool(nil)
						"authentication_profile": dsschema.StringAttribute{
							Description: "The AuthenticationProfile param.",
							Computed:    true,
						},
						"certificate_profile": dsschema.StringAttribute{
							Description: "The CertificateProfile param.",
							Computed:    true,
						},
						"gp_udp_port": dsschema.Int64Attribute{
							Description: "The GpUdpPort param. Value must be between 1 and 65535.",
							Computed:    true,
						},
						"id": dsschema.StringAttribute{
							Description: "UUID of the resource.",
							Computed:    true,
						},
						"idle_timer": dsschema.Int64Attribute{
							Description: "The IdleTimer param. Value must be between 1 and 1440.",
							Computed:    true,
						},
						"redirect_host": dsschema.StringAttribute{
							Description: "The RedirectHost param.",
							Computed:    true,
						},
						"timer": dsschema.Int64Attribute{
							Description: "The Timer param. Value must be between 1 and 1440.",
							Computed:    true,
						},
						"tls_service_profile": dsschema.StringAttribute{
							Description: "The TlsServiceProfile param.",
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
func (d *authenticationPortalListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *authenticationPortalListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state authenticationPortalListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_authentication_portal_list",
		"terraform_provider_function": "Read",
		"name":                        state.Name.ValueString(),
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
		"limit":                       state.Limit.ValueInt64(),
		"offset":                      state.Offset.ValueInt64(),
	})

	// Prepare to run the command.
	svc := ljRBvXJ.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := ljRBvXJ.ListInput{}

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
		state.Data = make([]authenticationPortalListDsModel_hhIWLbI_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := authenticationPortalListDsModel_hhIWLbI_Config{}

			var1.AuthenticationProfile = types.StringPointerValue(var0.AuthenticationProfile)

			var1.CertificateProfile = types.StringPointerValue(var0.CertificateProfile)

			var1.GpUdpPort = types.Int64PointerValue(var0.GpUdpPort)

			var1.Id = types.StringPointerValue(var0.Id)

			var1.IdleTimer = types.Int64PointerValue(var0.IdleTimer)

			var1.RedirectHost = types.StringPointerValue(var0.RedirectHost)

			var1.Timer = types.Int64PointerValue(var0.Timer)

			var1.TlsServiceProfile = types.StringPointerValue(var0.TlsServiceProfile)
			state.Data = append(state.Data, var1)
		}
	}

	state.Limit = types.Int64PointerValue(ans.Limit)

	state.Offset = types.Int64PointerValue(ans.Offset)

	state.Total = types.Int64PointerValue(ans.Total)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Data source.
var (
	_ datasource.DataSource              = &authenticationPortalDataSource{}
	_ datasource.DataSourceWithConfigure = &authenticationPortalDataSource{}
)

func NewAuthenticationPortalDataSource() datasource.DataSource {
	return &authenticationPortalDataSource{}
}

type authenticationPortalDataSource struct {
	client *scm.Client
}

// authenticationPortalDsModel is the model.
type authenticationPortalDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Id types.String `tfsdk:"id"`

	// Output.
	AuthenticationProfile types.String `tfsdk:"authentication_profile"`
	CertificateProfile    types.String `tfsdk:"certificate_profile"`
	GpUdpPort             types.Int64  `tfsdk:"gp_udp_port"`
	// omit input: id
	IdleTimer         types.Int64  `tfsdk:"idle_timer"`
	RedirectHost      types.String `tfsdk:"redirect_host"`
	Timer             types.Int64  `tfsdk:"timer"`
	TlsServiceProfile types.String `tfsdk:"tls_service_profile"`
}

// Metadata returns the data source type name.
func (d *authenticationPortalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_portal"
}

// Schema defines the schema for this data source.
func (d *authenticationPortalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"id":true} outputs:map[string]bool{"authentication_profile":true, "certificate_profile":true, "gp_udp_port":true, "id":true, "idle_timer":true, "redirect_host":true, "tfid":true, "timer":true, "tls_service_profile":true} forceNew:map[string]bool{"id":true}
			"authentication_profile": dsschema.StringAttribute{
				Description: "The AuthenticationProfile param.",
				Computed:    true,
			},
			"certificate_profile": dsschema.StringAttribute{
				Description: "The CertificateProfile param.",
				Computed:    true,
			},
			"gp_udp_port": dsschema.Int64Attribute{
				Description: "The GpUdpPort param. Value must be between 1 and 65535.",
				Computed:    true,
			},
			"id": dsschema.StringAttribute{
				Description: "The Id param.",
				Required:    true,
			},
			"idle_timer": dsschema.Int64Attribute{
				Description: "The IdleTimer param. Value must be between 1 and 1440.",
				Computed:    true,
			},
			"redirect_host": dsschema.StringAttribute{
				Description: "The RedirectHost param.",
				Computed:    true,
			},
			"tfid": dsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
			},
			"timer": dsschema.Int64Attribute{
				Description: "The Timer param. Value must be between 1 and 1440.",
				Computed:    true,
			},
			"tls_service_profile": dsschema.StringAttribute{
				Description: "The TlsServiceProfile param.",
				Computed:    true,
			},
		},
	}
}

// Configure prepares the struct.
func (d *authenticationPortalDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *authenticationPortalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state authenticationPortalDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source read", map[string]any{
		"data_source_name":            "scm_authentication_portal",
		"terraform_provider_function": "Read",
		"id":                          state.Id.ValueString(),
	})

	// Prepare to run the command.
	svc := ljRBvXJ.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := ljRBvXJ.ReadInput{}

	input.Id = state.Id.ValueString()

	// Perform the operation.
	ans, err := svc.Read(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError("Error reading config", err.Error())
		return
	}

	// Create the Terraform ID.
	var idBuilder strings.Builder
	idBuilder.WriteString(input.Id)

	// Store the answer to state.

	state.Tfid = types.StringValue(idBuilder.String())

	state.AuthenticationProfile = types.StringPointerValue(ans.AuthenticationProfile)

	state.CertificateProfile = types.StringPointerValue(ans.CertificateProfile)

	state.GpUdpPort = types.Int64PointerValue(ans.GpUdpPort)

	state.Id = types.StringPointerValue(ans.Id)

	state.IdleTimer = types.Int64PointerValue(ans.IdleTimer)

	state.RedirectHost = types.StringPointerValue(ans.RedirectHost)

	state.Timer = types.Int64PointerValue(ans.Timer)

	state.TlsServiceProfile = types.StringPointerValue(ans.TlsServiceProfile)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Resource.
var (
	_ resource.Resource                = &authenticationPortalResource{}
	_ resource.ResourceWithConfigure   = &authenticationPortalResource{}
	_ resource.ResourceWithImportState = &authenticationPortalResource{}
)

func NewAuthenticationPortalResource() resource.Resource {
	return &authenticationPortalResource{}
}

type authenticationPortalResource struct {
	client *scm.Client
}

// authenticationPortalRsModel is the model.
type authenticationPortalRsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	AuthenticationProfile types.String `tfsdk:"authentication_profile"`
	CertificateProfile    types.String `tfsdk:"certificate_profile"`
	Device                types.String `tfsdk:"device"`
	Folder                types.String `tfsdk:"folder"`
	GpUdpPort             types.Int64  `tfsdk:"gp_udp_port"`
	Id                    types.String `tfsdk:"id"`
	IdleTimer             types.Int64  `tfsdk:"idle_timer"`
	RedirectHost          types.String `tfsdk:"redirect_host"`
	Snippet               types.String `tfsdk:"snippet"`
	Timer                 types.Int64  `tfsdk:"timer"`
	TlsServiceProfile     types.String `tfsdk:"tls_service_profile"`

	// Output.
	// omit input: authentication_profile
	// omit input: certificate_profile
	// omit input: gp_udp_port
	// omit input: id
	// omit input: idle_timer
	// omit input: redirect_host
	// omit input: timer
	// omit input: tls_service_profile
}

// Metadata returns the data source type name.
func (r *authenticationPortalResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_portal"
}

// Schema defines the schema for this data source.
func (r *authenticationPortalResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]rsschema.Attribute{
			// inputs:map[string]bool{"authentication_profile":true, "certificate_profile":true, "device":true, "folder":true, "gp_udp_port":true, "id":true, "idle_timer":true, "redirect_host":true, "snippet":true, "timer":true, "tls_service_profile":true} outputs:map[string]bool{"authentication_profile":true, "certificate_profile":true, "gp_udp_port":true, "id":true, "idle_timer":true, "redirect_host":true, "tfid":true, "timer":true, "tls_service_profile":true} forceNew:map[string]bool{"device":true, "folder":true, "snippet":true}
			"authentication_profile": rsschema.StringAttribute{
				Description: "The AuthenticationProfile param.",
				Optional:    true,
			},
			"certificate_profile": rsschema.StringAttribute{
				Description: "The CertificateProfile param.",
				Optional:    true,
			},
			"device": rsschema.StringAttribute{
				Description: "The Device param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"folder": rsschema.StringAttribute{
				Description: "The Folder param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"gp_udp_port": rsschema.Int64Attribute{
				Description: "The GpUdpPort param. Value must be between 1 and 65535.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"id": rsschema.StringAttribute{
				Description: "UUID of the resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"idle_timer": rsschema.Int64Attribute{
				Description: "The IdleTimer param. Value must be between 1 and 1440.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1440),
				},
			},
			"redirect_host": rsschema.StringAttribute{
				Description: "The RedirectHost param.",
				Optional:    true,
			},
			"snippet": rsschema.StringAttribute{
				Description: "The Snippet param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"tfid": rsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"timer": rsschema.Int64Attribute{
				Description: "The Timer param. Value must be between 1 and 1440.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1440),
				},
			},
			"tls_service_profile": rsschema.StringAttribute{
				Description: "The TlsServiceProfile param.",
				Optional:    true,
			},
		},
	}
}

// Configure prepares the struct.
func (r *authenticationPortalResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*scm.Client)
}

// Create resource.
func (r *authenticationPortalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state authenticationPortalRsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name":               "scm_authentication_portal",
		"terraform_provider_function": "Create",
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
	})

	// Prepare to create the config.
	svc := ljRBvXJ.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ljRBvXJ.CreateInput{}

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()
	input.Request = &hhIWLbI.Config{}

	input.Request.AuthenticationProfile = state.AuthenticationProfile.ValueStringPointer()

	input.Request.CertificateProfile = state.CertificateProfile.ValueStringPointer()

	input.Request.GpUdpPort = state.GpUdpPort.ValueInt64Pointer()

	input.Request.IdleTimer = state.IdleTimer.ValueInt64Pointer()

	input.Request.RedirectHost = state.RedirectHost.ValueStringPointer()

	input.Request.Timer = state.Timer.ValueInt64Pointer()

	input.Request.TlsServiceProfile = state.TlsServiceProfile.ValueStringPointer()

	// Perform the operation.
	ans, err := svc.Create(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError("Error creating config", err.Error())
		return
	}

	// Create the Terraform ID.
	var idBuilder strings.Builder
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
	if ans.Id == nil {
		resp.Diagnostics.AddError("Undefined param required for ID", "Id")
		return
	}
	if ans.Id != nil {
		idBuilder.WriteString(*ans.Id)
	}

	// Store the answer to state.

	state.Tfid = types.StringValue(idBuilder.String())

	state.AuthenticationProfile = types.StringPointerValue(ans.AuthenticationProfile)

	state.CertificateProfile = types.StringPointerValue(ans.CertificateProfile)

	state.GpUdpPort = types.Int64PointerValue(ans.GpUdpPort)

	state.Id = types.StringPointerValue(ans.Id)

	state.IdleTimer = types.Int64PointerValue(ans.IdleTimer)

	state.RedirectHost = types.StringPointerValue(ans.RedirectHost)

	state.Timer = types.Int64PointerValue(ans.Timer)

	state.TlsServiceProfile = types.StringPointerValue(ans.TlsServiceProfile)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read performs Read for the struct.
func (r *authenticationPortalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var savestate, state authenticationPortalRsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tfid := savestate.Tfid.ValueString()
	tokens := strings.Split(tfid, IdSeparator)
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error in resource ID format", "Expected 4 tokens")
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource read", map[string]any{
		"terraform_provider_function": "Read",
		"resource_name":               "scm_authentication_portal",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	// Prepare to read the config.
	svc := ljRBvXJ.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ljRBvXJ.ReadInput{}

	input.Id = tokens[3]

	// Perform the operation.
	ans, err := svc.Read(ctx, input)
	if err != nil {
		if IsObjectNotFound(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Error reading config", err.Error())
		}
		return
	}

	// Store the answer to state.

	if tokens[0] == "" {
		state.Folder = types.StringNull()
	} else {
		state.Folder = types.StringValue(tokens[0])
	}

	if tokens[1] == "" {
		state.Snippet = types.StringNull()
	} else {
		state.Snippet = types.StringValue(tokens[1])
	}

	if tokens[2] == "" {
		state.Device = types.StringNull()
	} else {
		state.Device = types.StringValue(tokens[2])
	}
	state.Tfid = savestate.Tfid

	state.AuthenticationProfile = types.StringPointerValue(ans.AuthenticationProfile)

	state.CertificateProfile = types.StringPointerValue(ans.CertificateProfile)

	state.GpUdpPort = types.Int64PointerValue(ans.GpUdpPort)

	state.Id = types.StringPointerValue(ans.Id)

	state.IdleTimer = types.Int64PointerValue(ans.IdleTimer)

	state.RedirectHost = types.StringPointerValue(ans.RedirectHost)

	state.Timer = types.Int64PointerValue(ans.Timer)

	state.TlsServiceProfile = types.StringPointerValue(ans.TlsServiceProfile)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update performs the Update for the struct.
func (r *authenticationPortalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state authenticationPortalRsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tfid := state.Tfid.ValueString()
	tokens := strings.Split(tfid, IdSeparator)
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error in resource ID format", "Expected 4 tokens")
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource update", map[string]any{
		"terraform_provider_function": "Update",
		"resource_name":               "scm_authentication_portal",
		"tfid":                        state.Tfid.ValueString(),
	})

	// Prepare to update the config.
	svc := ljRBvXJ.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ljRBvXJ.UpdateInput{}

	if tokens[3] != "" {
		input.Id = tokens[3]
	}
	input.Request = &hhIWLbI.Config{}

	input.Request.AuthenticationProfile = plan.AuthenticationProfile.ValueStringPointer()

	input.Request.CertificateProfile = plan.CertificateProfile.ValueStringPointer()

	input.Request.GpUdpPort = plan.GpUdpPort.ValueInt64Pointer()

	input.Request.IdleTimer = plan.IdleTimer.ValueInt64Pointer()

	input.Request.RedirectHost = plan.RedirectHost.ValueStringPointer()

	input.Request.Timer = plan.Timer.ValueInt64Pointer()

	input.Request.TlsServiceProfile = plan.TlsServiceProfile.ValueStringPointer()

	// Perform the operation.
	ans, err := svc.Update(ctx, input)
	if err != nil {
		if IsObjectNotFound(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Error updating resource", err.Error())
		}
		return
	}

	// Store the answer to state.
	// Note: when supporting importing a resource, this will need to change to taking
	// values from the savestate.Tfid param and locMap.

	state.AuthenticationProfile = types.StringPointerValue(ans.AuthenticationProfile)

	state.CertificateProfile = types.StringPointerValue(ans.CertificateProfile)

	state.GpUdpPort = types.Int64PointerValue(ans.GpUdpPort)

	state.Id = types.StringPointerValue(ans.Id)

	state.IdleTimer = types.Int64PointerValue(ans.IdleTimer)

	state.RedirectHost = types.StringPointerValue(ans.RedirectHost)

	state.Timer = types.Int64PointerValue(ans.Timer)

	state.TlsServiceProfile = types.StringPointerValue(ans.TlsServiceProfile)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete performs delete for the struct.
func (r *authenticationPortalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var idType types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("tfid"), &idType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tfid := idType.ValueString()
	tokens := strings.Split(tfid, IdSeparator)
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error in resource ID format", "Expected 4 tokens")
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource delete", map[string]any{
		"terraform_provider_function": "Delete",
		"resource_name":               "scm_authentication_portal",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	svc := ljRBvXJ.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ljRBvXJ.DeleteInput{}

	input.Id = tokens[3]

	// Perform the operation.
	if _, err := svc.Delete(ctx, input); err != nil && !IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error in delete", err.Error())
	}
}

func (r *authenticationPortalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}