package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	qhaCvMo "github.com/paloaltonetworks/scm-go/netsec/schemas/profile/groups"
	ampruGo "github.com/paloaltonetworks/scm-go/netsec/services/profilegroups"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Data source (listing).
var (
	_ datasource.DataSource              = &profileGroupListDataSource{}
	_ datasource.DataSourceWithConfigure = &profileGroupListDataSource{}
)

func NewProfileGroupListDataSource() datasource.DataSource {
	return &profileGroupListDataSource{}
}

type profileGroupListDataSource struct {
	client *scm.Client
}

// profileGroupListDsModel is the model.
type profileGroupListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Device  types.String `tfsdk:"device"`
	Folder  types.String `tfsdk:"folder"`
	Limit   types.Int64  `tfsdk:"limit"`
	Name    types.String `tfsdk:"name"`
	Offset  types.Int64  `tfsdk:"offset"`
	Snippet types.String `tfsdk:"snippet"`

	// Output.
	Data []profileGroupListDsModel_qhaCvMo_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type profileGroupListDsModel_qhaCvMo_Config struct {
	AiSecurities             types.List   `tfsdk:"ai_securities"`
	DnsSecurities            types.List   `tfsdk:"dns_securities"`
	FileBlockings            types.List   `tfsdk:"file_blockings"`
	Id                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"`
	SaasSecurities           types.List   `tfsdk:"saas_securities"`
	Spywares                 types.List   `tfsdk:"spywares"`
	UrlFilterings            types.List   `tfsdk:"url_filterings"`
	VirusAndWildfireAnalyses types.List   `tfsdk:"virus_and_wildfire_analyses"`
	Vulnerabilities          types.List   `tfsdk:"vulnerabilities"`
}

// Metadata returns the data source type name.
func (d *profileGroupListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_profile_group_list"
}

// Schema defines the schema for this listing data source.
func (d *profileGroupListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"ai_security":true, "dns_security":true, "file_blocking":true, "id":true, "name":true, "saas_security":true, "spyware":true, "url_filtering":true, "virus_and_wildfire_analysis":true, "vulnerability":true} forceNew:map[string]bool(nil)
						"ai_securities": dsschema.ListAttribute{
							Description: "List of AI security profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"dns_securities": dsschema.ListAttribute{
							Description: "List of DNS security profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"file_blockings": dsschema.ListAttribute{
							Description: "List of file blocking profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"id": dsschema.StringAttribute{
							Description: "UUID of the resource.",
							Computed:    true,
						},
						"name": dsschema.StringAttribute{
							Description: "The name of the profile group.",
							Computed:    true,
						},
						"saas_securities": dsschema.ListAttribute{
							Description: "List of HTTP header insertion profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"spywares": dsschema.ListAttribute{
							Description: "List of anti-spyware profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"url_filterings": dsschema.ListAttribute{
							Description: "List of URL filtering profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"virus_and_wildfire_analyses": dsschema.ListAttribute{
							Description: "List of anti-virus and Wildfire analysis profiles.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"vulnerabilities": dsschema.ListAttribute{
							Description: "List of vulnerability protection profiles.",
							Computed:    true,
							ElementType: types.StringType,
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
func (d *profileGroupListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *profileGroupListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state profileGroupListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_profile_group_list",
		"terraform_provider_function": "Read",
		"name":                        state.Name.ValueString(),
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
		"offset":                      state.Offset.ValueInt64(),
		"limit":                       state.Limit.ValueInt64(),
	})

	// Prepare to run the command.
	svc := ampruGo.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := ampruGo.ListInput{}

	input.Name = state.Name.ValueStringPointer()

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()

	input.Offset = state.Offset.ValueInt64Pointer()

	input.Limit = state.Limit.ValueInt64Pointer()

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
	if input.Offset != nil {
		idBuilder.WriteString(strconv.FormatInt(*input.Offset, 10))
	}

	idBuilder.WriteString(IdSeparator)
	if input.Limit != nil {
		idBuilder.WriteString(strconv.FormatInt(*input.Limit, 10))
	}

	// Store the answer to state.

	state.Tfid = types.StringValue(idBuilder.String())

	if len(ans.Data) == 0 {
		state.Data = nil
	} else {
		state.Data = make([]profileGroupListDsModel_qhaCvMo_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := profileGroupListDsModel_qhaCvMo_Config{}

			var2, var3 := types.ListValueFrom(ctx, types.StringType, var0.AiSecurities)
			var1.AiSecurities = var2
			resp.Diagnostics.Append(var3.Errors()...)

			var4, var5 := types.ListValueFrom(ctx, types.StringType, var0.DnsSecurities)
			var1.DnsSecurities = var4
			resp.Diagnostics.Append(var5.Errors()...)

			var6, var7 := types.ListValueFrom(ctx, types.StringType, var0.FileBlockings)
			var1.FileBlockings = var6
			resp.Diagnostics.Append(var7.Errors()...)

			var1.Id = types.StringPointerValue(var0.Id)

			var1.Name = types.StringValue(var0.Name)

			var8, var9 := types.ListValueFrom(ctx, types.StringType, var0.SaasSecurities)
			var1.SaasSecurities = var8
			resp.Diagnostics.Append(var9.Errors()...)

			var10, var11 := types.ListValueFrom(ctx, types.StringType, var0.Spywares)
			var1.Spywares = var10
			resp.Diagnostics.Append(var11.Errors()...)

			var12, var13 := types.ListValueFrom(ctx, types.StringType, var0.UrlFilterings)
			var1.UrlFilterings = var12
			resp.Diagnostics.Append(var13.Errors()...)

			var14, var15 := types.ListValueFrom(ctx, types.StringType, var0.VirusAndWildfireAnalyses)
			var1.VirusAndWildfireAnalyses = var14
			resp.Diagnostics.Append(var15.Errors()...)

			var16, var17 := types.ListValueFrom(ctx, types.StringType, var0.Vulnerabilities)
			var1.Vulnerabilities = var16
			resp.Diagnostics.Append(var17.Errors()...)
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
	_ datasource.DataSource              = &profileGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &profileGroupDataSource{}
)

func NewProfileGroupDataSource() datasource.DataSource {
	return &profileGroupDataSource{}
}

type profileGroupDataSource struct {
	client *scm.Client
}

// profileGroupDsModel is the model.
type profileGroupDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Id types.String `tfsdk:"id"`

	// Output.
	AiSecurities  types.List `tfsdk:"ai_securities"`
	DnsSecurities types.List `tfsdk:"dns_securities"`
	FileBlockings types.List `tfsdk:"file_blockings"`
	// omit input: id
	Name                     types.String `tfsdk:"name"`
	SaasSecurities           types.List   `tfsdk:"saas_securities"`
	Spywares                 types.List   `tfsdk:"spywares"`
	UrlFilterings            types.List   `tfsdk:"url_filterings"`
	VirusAndWildfireAnalyses types.List   `tfsdk:"virus_and_wildfire_analyses"`
	Vulnerabilities          types.List   `tfsdk:"vulnerabilities"`
}

// Metadata returns the data source type name.
func (d *profileGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_profile_group"
}

// Schema defines the schema for this data source.
func (d *profileGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"id":true} outputs:map[string]bool{"ai_security":true, "dns_security":true, "file_blocking":true, "id":true, "name":true, "saas_security":true, "spyware":true, "tfid":true, "url_filtering":true, "virus_and_wildfire_analysis":true, "vulnerability":true} forceNew:map[string]bool{"id":true}
			"ai_securities": dsschema.ListAttribute{
				Description: "List of AI security profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"dns_securities": dsschema.ListAttribute{
				Description: "List of DNS security profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"file_blockings": dsschema.ListAttribute{
				Description: "List of file blocking profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"id": dsschema.StringAttribute{
				Description: "The Id param.",
				Required:    true,
			},
			"name": dsschema.StringAttribute{
				Description: "The name of the profile group.",
				Computed:    true,
			},
			"saas_securities": dsschema.ListAttribute{
				Description: "List of HTTP header insertion profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"spywares": dsschema.ListAttribute{
				Description: "List of anti-spyware profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"tfid": dsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
			},
			"url_filterings": dsschema.ListAttribute{
				Description: "List of URL filtering profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"virus_and_wildfire_analyses": dsschema.ListAttribute{
				Description: "List of anti-virus and Wildfire analysis profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"vulnerabilities": dsschema.ListAttribute{
				Description: "List of vulnerability protection profiles.",
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// Configure prepares the struct.
func (d *profileGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *profileGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state profileGroupDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source read", map[string]any{
		"data_source_name":            "scm_profile_group",
		"terraform_provider_function": "Read",
		"id":                          state.Id.ValueString(),
	})

	// Prepare to run the command.
	svc := ampruGo.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := ampruGo.ReadInput{}

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

	var0, var1 := types.ListValueFrom(ctx, types.StringType, ans.AiSecurities)
	state.AiSecurities = var0
	resp.Diagnostics.Append(var1.Errors()...)

	var2, var3 := types.ListValueFrom(ctx, types.StringType, ans.DnsSecurities)
	state.DnsSecurities = var2
	resp.Diagnostics.Append(var3.Errors()...)

	var4, var5 := types.ListValueFrom(ctx, types.StringType, ans.FileBlockings)
	state.FileBlockings = var4
	resp.Diagnostics.Append(var5.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	var6, var7 := types.ListValueFrom(ctx, types.StringType, ans.SaasSecurities)
	state.SaasSecurities = var6
	resp.Diagnostics.Append(var7.Errors()...)

	var8, var9 := types.ListValueFrom(ctx, types.StringType, ans.Spywares)
	state.Spywares = var8
	resp.Diagnostics.Append(var9.Errors()...)

	var10, var11 := types.ListValueFrom(ctx, types.StringType, ans.UrlFilterings)
	state.UrlFilterings = var10
	resp.Diagnostics.Append(var11.Errors()...)

	var12, var13 := types.ListValueFrom(ctx, types.StringType, ans.VirusAndWildfireAnalyses)
	state.VirusAndWildfireAnalyses = var12
	resp.Diagnostics.Append(var13.Errors()...)

	var14, var15 := types.ListValueFrom(ctx, types.StringType, ans.Vulnerabilities)
	state.Vulnerabilities = var14
	resp.Diagnostics.Append(var15.Errors()...)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Resource.
var (
	_ resource.Resource                = &profileGroupResource{}
	_ resource.ResourceWithConfigure   = &profileGroupResource{}
	_ resource.ResourceWithImportState = &profileGroupResource{}
)

func NewProfileGroupResource() resource.Resource {
	return &profileGroupResource{}
}

type profileGroupResource struct {
	client *scm.Client
}

// profileGroupRsModel is the model.
type profileGroupRsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	AiSecurities             types.List   `tfsdk:"ai_securities"`
	Device                   types.String `tfsdk:"device"`
	DnsSecurities            types.List   `tfsdk:"dns_securities"`
	FileBlockings            types.List   `tfsdk:"file_blockings"`
	Folder                   types.String `tfsdk:"folder"`
	Id                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"`
	SaasSecurities           types.List   `tfsdk:"saas_securities"`
	Snippet                  types.String `tfsdk:"snippet"`
	Spywares                 types.List   `tfsdk:"spywares"`
	UrlFilterings            types.List   `tfsdk:"url_filterings"`
	VirusAndWildfireAnalyses types.List   `tfsdk:"virus_and_wildfire_analyses"`
	Vulnerabilities          types.List   `tfsdk:"vulnerabilities"`

	// Output.
	// omit input: ai_securities
	// omit input: dns_securities
	// omit input: file_blockings
	// omit input: id
	// omit input: name
	// omit input: saas_securities
	// omit input: spywares
	// omit input: url_filterings
	// omit input: virus_and_wildfire_analyses
	// omit input: vulnerabilities
}

// Metadata returns the data source type name.
func (r *profileGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_profile_group"
}

// Schema defines the schema for this data source.
func (r *profileGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]rsschema.Attribute{
			// inputs:map[string]bool{"ai_security":true, "device":true, "dns_security":true, "file_blocking":true, "folder":true, "id":true, "name":true, "saas_security":true, "snippet":true, "spyware":true, "url_filtering":true, "virus_and_wildfire_analysis":true, "vulnerability":true} outputs:map[string]bool{"ai_security":true, "dns_security":true, "file_blocking":true, "id":true, "name":true, "saas_security":true, "spyware":true, "tfid":true, "url_filtering":true, "virus_and_wildfire_analysis":true, "vulnerability":true} forceNew:map[string]bool{"device":true, "folder":true, "snippet":true}
			"ai_securities": rsschema.ListAttribute{
				Description: "List of AI security profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"device": rsschema.StringAttribute{
				Description: "The Device param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dns_securities": rsschema.ListAttribute{
				Description: "List of DNS security profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"file_blockings": rsschema.ListAttribute{
				Description: "List of file blocking profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"folder": rsschema.StringAttribute{
				Description: "The Folder param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": rsschema.StringAttribute{
				Description: "UUID of the resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": rsschema.StringAttribute{
				Description: "The name of the profile group.",
				Required:    true,
			},
			"saas_securities": rsschema.ListAttribute{
				Description: "List of HTTP header insertion profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"snippet": rsschema.StringAttribute{
				Description: "The Snippet param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"spywares": rsschema.ListAttribute{
				Description: "List of anti-spyware profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"tfid": rsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url_filterings": rsschema.ListAttribute{
				Description: "List of URL filtering profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"virus_and_wildfire_analyses": rsschema.ListAttribute{
				Description: "List of anti-virus and Wildfire analysis profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"vulnerabilities": rsschema.ListAttribute{
				Description: "List of vulnerability protection profiles.",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// Configure prepares the struct.
func (r *profileGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*scm.Client)
}

// Create resource.
func (r *profileGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state profileGroupRsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name":               "scm_profile_group",
		"terraform_provider_function": "Create",
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
	})

	// Prepare to create the config.
	svc := ampruGo.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ampruGo.CreateInput{}

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()
	input.Request = &qhaCvMo.Config{}

	resp.Diagnostics.Append(state.AiSecurities.ElementsAs(ctx, &input.Request.AiSecurities, false)...)
	//if len(state.AiSecurities) != 0 {
	//    input.Request.AiSecurities = make([]string, 0, len(state.AiSecurities))
	//    for _, var0 := range state.AiSecurities {
	//        input.Request.AiSecurities = append(input.Request.AiSecurities, var0.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.DnsSecurities.ElementsAs(ctx, &input.Request.DnsSecurities, false)...)
	//if len(state.DnsSecurities) != 0 {
	//    input.Request.DnsSecurities = make([]string, 0, len(state.DnsSecurities))
	//    for _, var1 := range state.DnsSecurities {
	//        input.Request.DnsSecurities = append(input.Request.DnsSecurities, var1.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.FileBlockings.ElementsAs(ctx, &input.Request.FileBlockings, false)...)
	//if len(state.FileBlockings) != 0 {
	//    input.Request.FileBlockings = make([]string, 0, len(state.FileBlockings))
	//    for _, var2 := range state.FileBlockings {
	//        input.Request.FileBlockings = append(input.Request.FileBlockings, var2.ValueString())
	//    }
	//}

	input.Request.Name = state.Name.ValueString()

	resp.Diagnostics.Append(state.SaasSecurities.ElementsAs(ctx, &input.Request.SaasSecurities, false)...)
	//if len(state.SaasSecurities) != 0 {
	//    input.Request.SaasSecurities = make([]string, 0, len(state.SaasSecurities))
	//    for _, var3 := range state.SaasSecurities {
	//        input.Request.SaasSecurities = append(input.Request.SaasSecurities, var3.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.Spywares.ElementsAs(ctx, &input.Request.Spywares, false)...)
	//if len(state.Spywares) != 0 {
	//    input.Request.Spywares = make([]string, 0, len(state.Spywares))
	//    for _, var4 := range state.Spywares {
	//        input.Request.Spywares = append(input.Request.Spywares, var4.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.UrlFilterings.ElementsAs(ctx, &input.Request.UrlFilterings, false)...)
	//if len(state.UrlFilterings) != 0 {
	//    input.Request.UrlFilterings = make([]string, 0, len(state.UrlFilterings))
	//    for _, var5 := range state.UrlFilterings {
	//        input.Request.UrlFilterings = append(input.Request.UrlFilterings, var5.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.VirusAndWildfireAnalyses.ElementsAs(ctx, &input.Request.VirusAndWildfireAnalyses, false)...)
	//if len(state.VirusAndWildfireAnalyses) != 0 {
	//    input.Request.VirusAndWildfireAnalyses = make([]string, 0, len(state.VirusAndWildfireAnalyses))
	//    for _, var6 := range state.VirusAndWildfireAnalyses {
	//        input.Request.VirusAndWildfireAnalyses = append(input.Request.VirusAndWildfireAnalyses, var6.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.Vulnerabilities.ElementsAs(ctx, &input.Request.Vulnerabilities, false)...)
	//if len(state.Vulnerabilities) != 0 {
	//    input.Request.Vulnerabilities = make([]string, 0, len(state.Vulnerabilities))
	//    for _, var7 := range state.Vulnerabilities {
	//        input.Request.Vulnerabilities = append(input.Request.Vulnerabilities, var7.ValueString())
	//    }
	//}

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

	var8, var9 := types.ListValueFrom(ctx, types.StringType, ans.AiSecurities)
	state.AiSecurities = var8
	resp.Diagnostics.Append(var9.Errors()...)

	var10, var11 := types.ListValueFrom(ctx, types.StringType, ans.DnsSecurities)
	state.DnsSecurities = var10
	resp.Diagnostics.Append(var11.Errors()...)

	var12, var13 := types.ListValueFrom(ctx, types.StringType, ans.FileBlockings)
	state.FileBlockings = var12
	resp.Diagnostics.Append(var13.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	var14, var15 := types.ListValueFrom(ctx, types.StringType, ans.SaasSecurities)
	state.SaasSecurities = var14
	resp.Diagnostics.Append(var15.Errors()...)

	var16, var17 := types.ListValueFrom(ctx, types.StringType, ans.Spywares)
	state.Spywares = var16
	resp.Diagnostics.Append(var17.Errors()...)

	var18, var19 := types.ListValueFrom(ctx, types.StringType, ans.UrlFilterings)
	state.UrlFilterings = var18
	resp.Diagnostics.Append(var19.Errors()...)

	var20, var21 := types.ListValueFrom(ctx, types.StringType, ans.VirusAndWildfireAnalyses)
	state.VirusAndWildfireAnalyses = var20
	resp.Diagnostics.Append(var21.Errors()...)

	var22, var23 := types.ListValueFrom(ctx, types.StringType, ans.Vulnerabilities)
	state.Vulnerabilities = var22
	resp.Diagnostics.Append(var23.Errors()...)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read performs Read for the struct.
func (r *profileGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var savestate, state profileGroupRsModel
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
		"resource_name":               "scm_profile_group",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	// Prepare to read the config.
	svc := ampruGo.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ampruGo.ReadInput{}

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

	var0, var1 := types.ListValueFrom(ctx, types.StringType, ans.AiSecurities)
	state.AiSecurities = var0
	resp.Diagnostics.Append(var1.Errors()...)

	var2, var3 := types.ListValueFrom(ctx, types.StringType, ans.DnsSecurities)
	state.DnsSecurities = var2
	resp.Diagnostics.Append(var3.Errors()...)

	var4, var5 := types.ListValueFrom(ctx, types.StringType, ans.FileBlockings)
	state.FileBlockings = var4
	resp.Diagnostics.Append(var5.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	var6, var7 := types.ListValueFrom(ctx, types.StringType, ans.SaasSecurities)
	state.SaasSecurities = var6
	resp.Diagnostics.Append(var7.Errors()...)

	var8, var9 := types.ListValueFrom(ctx, types.StringType, ans.Spywares)
	state.Spywares = var8
	resp.Diagnostics.Append(var9.Errors()...)

	var10, var11 := types.ListValueFrom(ctx, types.StringType, ans.UrlFilterings)
	state.UrlFilterings = var10
	resp.Diagnostics.Append(var11.Errors()...)

	var12, var13 := types.ListValueFrom(ctx, types.StringType, ans.VirusAndWildfireAnalyses)
	state.VirusAndWildfireAnalyses = var12
	resp.Diagnostics.Append(var13.Errors()...)

	var14, var15 := types.ListValueFrom(ctx, types.StringType, ans.Vulnerabilities)
	state.Vulnerabilities = var14
	resp.Diagnostics.Append(var15.Errors()...)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update performs the Update for the struct.
func (r *profileGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state profileGroupRsModel
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
		"resource_name":               "scm_profile_group",
		"tfid":                        state.Tfid.ValueString(),
	})

	// Prepare to update the config.
	svc := ampruGo.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ampruGo.UpdateInput{}

	if tokens[3] != "" {
		input.Id = tokens[3]
	}
	input.Request = &qhaCvMo.Config{}

	resp.Diagnostics.Append(plan.AiSecurities.ElementsAs(ctx, &input.Request.AiSecurities, false)...)
	//if len(plan.AiSecurities) != 0 {
	//    input.Request.AiSecurities = make([]string, 0, len(plan.AiSecurities))
	//    for _, var0 := range plan.AiSecurities {
	//        input.Request.AiSecurities = append(input.Request.AiSecurities, var0.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.DnsSecurities.ElementsAs(ctx, &input.Request.DnsSecurities, false)...)
	//if len(plan.DnsSecurities) != 0 {
	//    input.Request.DnsSecurities = make([]string, 0, len(plan.DnsSecurities))
	//    for _, var1 := range plan.DnsSecurities {
	//        input.Request.DnsSecurities = append(input.Request.DnsSecurities, var1.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.FileBlockings.ElementsAs(ctx, &input.Request.FileBlockings, false)...)
	//if len(plan.FileBlockings) != 0 {
	//    input.Request.FileBlockings = make([]string, 0, len(plan.FileBlockings))
	//    for _, var2 := range plan.FileBlockings {
	//        input.Request.FileBlockings = append(input.Request.FileBlockings, var2.ValueString())
	//    }
	//}

	input.Request.Name = plan.Name.ValueString()

	resp.Diagnostics.Append(plan.SaasSecurities.ElementsAs(ctx, &input.Request.SaasSecurities, false)...)
	//if len(plan.SaasSecurities) != 0 {
	//    input.Request.SaasSecurities = make([]string, 0, len(plan.SaasSecurities))
	//    for _, var3 := range plan.SaasSecurities {
	//        input.Request.SaasSecurities = append(input.Request.SaasSecurities, var3.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.Spywares.ElementsAs(ctx, &input.Request.Spywares, false)...)
	//if len(plan.Spywares) != 0 {
	//    input.Request.Spywares = make([]string, 0, len(plan.Spywares))
	//    for _, var4 := range plan.Spywares {
	//        input.Request.Spywares = append(input.Request.Spywares, var4.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.UrlFilterings.ElementsAs(ctx, &input.Request.UrlFilterings, false)...)
	//if len(plan.UrlFilterings) != 0 {
	//    input.Request.UrlFilterings = make([]string, 0, len(plan.UrlFilterings))
	//    for _, var5 := range plan.UrlFilterings {
	//        input.Request.UrlFilterings = append(input.Request.UrlFilterings, var5.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.VirusAndWildfireAnalyses.ElementsAs(ctx, &input.Request.VirusAndWildfireAnalyses, false)...)
	//if len(plan.VirusAndWildfireAnalyses) != 0 {
	//    input.Request.VirusAndWildfireAnalyses = make([]string, 0, len(plan.VirusAndWildfireAnalyses))
	//    for _, var6 := range plan.VirusAndWildfireAnalyses {
	//        input.Request.VirusAndWildfireAnalyses = append(input.Request.VirusAndWildfireAnalyses, var6.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.Vulnerabilities.ElementsAs(ctx, &input.Request.Vulnerabilities, false)...)
	//if len(plan.Vulnerabilities) != 0 {
	//    input.Request.Vulnerabilities = make([]string, 0, len(plan.Vulnerabilities))
	//    for _, var7 := range plan.Vulnerabilities {
	//        input.Request.Vulnerabilities = append(input.Request.Vulnerabilities, var7.ValueString())
	//    }
	//}

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

	var8, var9 := types.ListValueFrom(ctx, types.StringType, ans.AiSecurities)
	state.AiSecurities = var8
	resp.Diagnostics.Append(var9.Errors()...)

	var10, var11 := types.ListValueFrom(ctx, types.StringType, ans.DnsSecurities)
	state.DnsSecurities = var10
	resp.Diagnostics.Append(var11.Errors()...)

	var12, var13 := types.ListValueFrom(ctx, types.StringType, ans.FileBlockings)
	state.FileBlockings = var12
	resp.Diagnostics.Append(var13.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	var14, var15 := types.ListValueFrom(ctx, types.StringType, ans.SaasSecurities)
	state.SaasSecurities = var14
	resp.Diagnostics.Append(var15.Errors()...)

	var16, var17 := types.ListValueFrom(ctx, types.StringType, ans.Spywares)
	state.Spywares = var16
	resp.Diagnostics.Append(var17.Errors()...)

	var18, var19 := types.ListValueFrom(ctx, types.StringType, ans.UrlFilterings)
	state.UrlFilterings = var18
	resp.Diagnostics.Append(var19.Errors()...)

	var20, var21 := types.ListValueFrom(ctx, types.StringType, ans.VirusAndWildfireAnalyses)
	state.VirusAndWildfireAnalyses = var20
	resp.Diagnostics.Append(var21.Errors()...)

	var22, var23 := types.ListValueFrom(ctx, types.StringType, ans.Vulnerabilities)
	state.Vulnerabilities = var22
	resp.Diagnostics.Append(var23.Errors()...)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete performs delete for the struct.
func (r *profileGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
		"resource_name":               "scm_profile_group",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	svc := ampruGo.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := ampruGo.DeleteInput{}

	input.Id = tokens[3]

	// Perform the operation.
	if _, err := svc.Delete(ctx, input); err != nil && !IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error in delete", err.Error())
	}
}

func (r *profileGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
