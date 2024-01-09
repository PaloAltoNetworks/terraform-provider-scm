package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	wugpput "github.com/paloaltonetworks/scm-go/netsec/schemas/ike/crypto/profiles"
	uIHLJPY "github.com/paloaltonetworks/scm-go/netsec/services/ikecryptoprofiles"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Data source (listing).
var (
	_ datasource.DataSource              = &ikeCryptoProfileListDataSource{}
	_ datasource.DataSourceWithConfigure = &ikeCryptoProfileListDataSource{}
)

func NewIkeCryptoProfileListDataSource() datasource.DataSource {
	return &ikeCryptoProfileListDataSource{}
}

type ikeCryptoProfileListDataSource struct {
	client *scm.Client
}

// ikeCryptoProfileListDsModel is the model.
type ikeCryptoProfileListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Device  types.String `tfsdk:"device"`
	Folder  types.String `tfsdk:"folder"`
	Limit   types.Int64  `tfsdk:"limit"`
	Name    types.String `tfsdk:"name"`
	Offset  types.Int64  `tfsdk:"offset"`
	Snippet types.String `tfsdk:"snippet"`

	// Output.
	Data []ikeCryptoProfileListDsModel_wugpput_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type ikeCryptoProfileListDsModel_wugpput_Config struct {
	AuthenticationMultiple types.Int64                                         `tfsdk:"authentication_multiple"`
	DhGroups               types.List                                          `tfsdk:"dh_groups"`
	Encryptions            types.List                                          `tfsdk:"encryptions"`
	Hashes                 types.List                                          `tfsdk:"hashes"`
	Id                     types.String                                        `tfsdk:"id"`
	Lifetime               *ikeCryptoProfileListDsModel_wugpput_LifetimeObject `tfsdk:"lifetime"`
	Name                   types.String                                        `tfsdk:"name"`
}

type ikeCryptoProfileListDsModel_wugpput_LifetimeObject struct {
	Days    types.Int64 `tfsdk:"days"`
	Hours   types.Int64 `tfsdk:"hours"`
	Minutes types.Int64 `tfsdk:"minutes"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// Metadata returns the data source type name.
func (d *ikeCryptoProfileListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ike_crypto_profile_list"
}

// Schema defines the schema for this listing data source.
func (d *ikeCryptoProfileListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"authentication_multiple":true, "dh_group":true, "encryption":true, "hash":true, "id":true, "lifetime":true, "name":true} forceNew:map[string]bool(nil)
						"authentication_multiple": dsschema.Int64Attribute{
							Description: "IKEv2 SA reauthentication interval equals authetication-multiple * rekey-lifetime; 0 means reauthentication disabled. Value must be less than or equal to 50. Default: `0`.",
							Computed:    true,
						},
						"dh_groups": dsschema.ListAttribute{
							Description: "The DhGroups param.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"encryptions": dsschema.ListAttribute{
							Description: "Encryption algorithm.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"hashes": dsschema.ListAttribute{
							Description: "The Hashes param.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"id": dsschema.StringAttribute{
							Description: "UUID of the resource.",
							Computed:    true,
						},
						"lifetime": dsschema.SingleNestedAttribute{
							Description: "The Lifetime param.",
							Computed:    true,
							Attributes: map[string]dsschema.Attribute{
								// inputs:map[string]bool{} outputs:map[string]bool{"days":true, "hours":true, "minutes":true, "seconds":true} forceNew:map[string]bool(nil)
								"days": dsschema.Int64Attribute{
									Description: "specify lifetime in days. Value must be between 1 and 365.",
									Computed:    true,
								},
								"hours": dsschema.Int64Attribute{
									Description: "specify lifetime in hours. Value must be between 1 and 65535.",
									Computed:    true,
								},
								"minutes": dsschema.Int64Attribute{
									Description: "specify lifetime in minutes. Value must be between 3 and 65535.",
									Computed:    true,
								},
								"seconds": dsschema.Int64Attribute{
									Description: "specify lifetime in seconds. Value must be between 180 and 65535.",
									Computed:    true,
								},
							},
						},
						"name": dsschema.StringAttribute{
							Description: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]. String length must not exceed 31 characters.",
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
func (d *ikeCryptoProfileListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *ikeCryptoProfileListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ikeCryptoProfileListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_ike_crypto_profile_list",
		"terraform_provider_function": "Read",
		"name":                        state.Name.ValueString(),
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
		"offset":                      state.Offset.ValueInt64(),
		"limit":                       state.Limit.ValueInt64(),
	})

	// Prepare to run the command.
	svc := uIHLJPY.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := uIHLJPY.ListInput{}

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
		state.Data = make([]ikeCryptoProfileListDsModel_wugpput_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := ikeCryptoProfileListDsModel_wugpput_Config{}

			var1.AuthenticationMultiple = types.Int64PointerValue(var0.AuthenticationMultiple)

			var2, var3 := types.ListValueFrom(ctx, types.StringType, var0.DhGroups)
			var1.DhGroups = var2
			resp.Diagnostics.Append(var3.Errors()...)

			var4, var5 := types.ListValueFrom(ctx, types.StringType, var0.Encryptions)
			var1.Encryptions = var4
			resp.Diagnostics.Append(var5.Errors()...)

			var6, var7 := types.ListValueFrom(ctx, types.StringType, var0.Hashes)
			var1.Hashes = var6
			resp.Diagnostics.Append(var7.Errors()...)

			var1.Id = types.StringPointerValue(var0.Id)

			if var0.Lifetime == nil {
				var1.Lifetime = nil
			} else {
				var1.Lifetime = &ikeCryptoProfileListDsModel_wugpput_LifetimeObject{}

				var1.Lifetime.Days = types.Int64PointerValue(var0.Lifetime.Days)

				var1.Lifetime.Hours = types.Int64PointerValue(var0.Lifetime.Hours)

				var1.Lifetime.Minutes = types.Int64PointerValue(var0.Lifetime.Minutes)

				var1.Lifetime.Seconds = types.Int64PointerValue(var0.Lifetime.Seconds)
			}

			var1.Name = types.StringValue(var0.Name)
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
	_ datasource.DataSource              = &ikeCryptoProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &ikeCryptoProfileDataSource{}
)

func NewIkeCryptoProfileDataSource() datasource.DataSource {
	return &ikeCryptoProfileDataSource{}
}

type ikeCryptoProfileDataSource struct {
	client *scm.Client
}

// ikeCryptoProfileDsModel is the model.
type ikeCryptoProfileDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Id types.String `tfsdk:"id"`

	// Output.
	AuthenticationMultiple types.Int64 `tfsdk:"authentication_multiple"`
	DhGroups               types.List  `tfsdk:"dh_groups"`
	Encryptions            types.List  `tfsdk:"encryptions"`
	Hashes                 types.List  `tfsdk:"hashes"`
	// omit input: id
	Lifetime *ikeCryptoProfileDsModel_wugpput_LifetimeObject `tfsdk:"lifetime"`
	Name     types.String                                    `tfsdk:"name"`
}

type ikeCryptoProfileDsModel_wugpput_LifetimeObject struct {
	Days    types.Int64 `tfsdk:"days"`
	Hours   types.Int64 `tfsdk:"hours"`
	Minutes types.Int64 `tfsdk:"minutes"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// Metadata returns the data source type name.
func (d *ikeCryptoProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ike_crypto_profile"
}

// Schema defines the schema for this data source.
func (d *ikeCryptoProfileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"id":true} outputs:map[string]bool{"authentication_multiple":true, "dh_group":true, "encryption":true, "hash":true, "id":true, "lifetime":true, "name":true, "tfid":true} forceNew:map[string]bool{"id":true}
			"authentication_multiple": dsschema.Int64Attribute{
				Description: "IKEv2 SA reauthentication interval equals authetication-multiple * rekey-lifetime; 0 means reauthentication disabled. Value must be less than or equal to 50. Default: `0`.",
				Computed:    true,
			},
			"dh_groups": dsschema.ListAttribute{
				Description: "The DhGroups param.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"encryptions": dsschema.ListAttribute{
				Description: "Encryption algorithm.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"hashes": dsschema.ListAttribute{
				Description: "The Hashes param.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"id": dsschema.StringAttribute{
				Description: "The Id param.",
				Required:    true,
			},
			"lifetime": dsschema.SingleNestedAttribute{
				Description: "The Lifetime param.",
				Computed:    true,
				Attributes: map[string]dsschema.Attribute{
					// inputs:map[string]bool{} outputs:map[string]bool{"days":true, "hours":true, "minutes":true, "seconds":true} forceNew:map[string]bool(nil)
					"days": dsschema.Int64Attribute{
						Description: "specify lifetime in days. Value must be between 1 and 365.",
						Computed:    true,
					},
					"hours": dsschema.Int64Attribute{
						Description: "specify lifetime in hours. Value must be between 1 and 65535.",
						Computed:    true,
					},
					"minutes": dsschema.Int64Attribute{
						Description: "specify lifetime in minutes. Value must be between 3 and 65535.",
						Computed:    true,
					},
					"seconds": dsschema.Int64Attribute{
						Description: "specify lifetime in seconds. Value must be between 180 and 65535.",
						Computed:    true,
					},
				},
			},
			"name": dsschema.StringAttribute{
				Description: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]. String length must not exceed 31 characters.",
				Computed:    true,
			},
			"tfid": dsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
			},
		},
	}
}

// Configure prepares the struct.
func (d *ikeCryptoProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *ikeCryptoProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ikeCryptoProfileDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source read", map[string]any{
		"data_source_name":            "scm_ike_crypto_profile",
		"terraform_provider_function": "Read",
		"id":                          state.Id.ValueString(),
	})

	// Prepare to run the command.
	svc := uIHLJPY.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := uIHLJPY.ReadInput{}

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

	state.AuthenticationMultiple = types.Int64PointerValue(ans.AuthenticationMultiple)

	var0, var1 := types.ListValueFrom(ctx, types.StringType, ans.DhGroups)
	state.DhGroups = var0
	resp.Diagnostics.Append(var1.Errors()...)

	var2, var3 := types.ListValueFrom(ctx, types.StringType, ans.Encryptions)
	state.Encryptions = var2
	resp.Diagnostics.Append(var3.Errors()...)

	var4, var5 := types.ListValueFrom(ctx, types.StringType, ans.Hashes)
	state.Hashes = var4
	resp.Diagnostics.Append(var5.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	if ans.Lifetime == nil {
		state.Lifetime = nil
	} else {
		state.Lifetime = &ikeCryptoProfileDsModel_wugpput_LifetimeObject{}

		state.Lifetime.Days = types.Int64PointerValue(ans.Lifetime.Days)

		state.Lifetime.Hours = types.Int64PointerValue(ans.Lifetime.Hours)

		state.Lifetime.Minutes = types.Int64PointerValue(ans.Lifetime.Minutes)

		state.Lifetime.Seconds = types.Int64PointerValue(ans.Lifetime.Seconds)
	}

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Resource.
var (
	_ resource.Resource                = &ikeCryptoProfileResource{}
	_ resource.ResourceWithConfigure   = &ikeCryptoProfileResource{}
	_ resource.ResourceWithImportState = &ikeCryptoProfileResource{}
)

func NewIkeCryptoProfileResource() resource.Resource {
	return &ikeCryptoProfileResource{}
}

type ikeCryptoProfileResource struct {
	client *scm.Client
}

// ikeCryptoProfileRsModel is the model.
type ikeCryptoProfileRsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	AuthenticationMultiple types.Int64                                     `tfsdk:"authentication_multiple"`
	Device                 types.String                                    `tfsdk:"device"`
	DhGroups               types.List                                      `tfsdk:"dh_groups"`
	Encryptions            types.List                                      `tfsdk:"encryptions"`
	Folder                 types.String                                    `tfsdk:"folder"`
	Hashes                 types.List                                      `tfsdk:"hashes"`
	Id                     types.String                                    `tfsdk:"id"`
	Lifetime               *ikeCryptoProfileRsModel_wugpput_LifetimeObject `tfsdk:"lifetime"`
	Name                   types.String                                    `tfsdk:"name"`
	Snippet                types.String                                    `tfsdk:"snippet"`

	// Output.
	// omit input: authentication_multiple
	// omit input: dh_groups
	// omit input: encryptions
	// omit input: hashes
	// omit input: id
	// omit input: lifetime
	// omit input: name
}

type ikeCryptoProfileRsModel_wugpput_LifetimeObject struct {
	Days    types.Int64 `tfsdk:"days"`
	Hours   types.Int64 `tfsdk:"hours"`
	Minutes types.Int64 `tfsdk:"minutes"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// Metadata returns the data source type name.
func (r *ikeCryptoProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ike_crypto_profile"
}

// Schema defines the schema for this data source.
func (r *ikeCryptoProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]rsschema.Attribute{
			// inputs:map[string]bool{"authentication_multiple":true, "device":true, "dh_group":true, "encryption":true, "folder":true, "hash":true, "id":true, "lifetime":true, "name":true, "snippet":true} outputs:map[string]bool{"authentication_multiple":true, "dh_group":true, "encryption":true, "hash":true, "id":true, "lifetime":true, "name":true, "tfid":true} forceNew:map[string]bool{"device":true, "folder":true, "snippet":true}
			"authentication_multiple": rsschema.Int64Attribute{
				Description: "IKEv2 SA reauthentication interval equals authetication-multiple * rekey-lifetime; 0 means reauthentication disabled. Value must be less than or equal to 50. Default: `0`.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.AtMost(50),
				},
			},
			"device": rsschema.StringAttribute{
				Description: "The Device param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dh_groups": rsschema.ListAttribute{
				Description: "The DhGroups param.",
				Required:    true,
				ElementType: types.StringType,
			},
			"encryptions": rsschema.ListAttribute{
				Description: "Encryption algorithm.",
				Required:    true,
				ElementType: types.StringType,
			},
			"folder": rsschema.StringAttribute{
				Description: "The Folder param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"hashes": rsschema.ListAttribute{
				Description: "The Hashes param.",
				Required:    true,
				ElementType: types.StringType,
			},
			"id": rsschema.StringAttribute{
				Description: "UUID of the resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"lifetime": rsschema.SingleNestedAttribute{
				Description: "The Lifetime param.",
				Optional:    true,
				Attributes: map[string]rsschema.Attribute{
					// inputs:map[string]bool{"days":true, "hours":true, "minutes":true, "seconds":true} outputs:map[string]bool{"days":true, "hours":true, "minutes":true, "seconds":true} forceNew:map[string]bool(nil)
					"days": rsschema.Int64Attribute{
						Description: "specify lifetime in days. Value must be between 1 and 365.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 365),
						},
					},
					"hours": rsschema.Int64Attribute{
						Description: "specify lifetime in hours. Value must be between 1 and 65535.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
					},
					"minutes": rsschema.Int64Attribute{
						Description: "specify lifetime in minutes. Value must be between 3 and 65535.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(3, 65535),
						},
					},
					"seconds": rsschema.Int64Attribute{
						Description: "specify lifetime in seconds. Value must be between 180 and 65535.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(180, 65535),
						},
					},
				},
			},
			"name": rsschema.StringAttribute{
				Description: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]. String length must not exceed 31 characters.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(31),
				},
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
		},
	}
}

// Configure prepares the struct.
func (r *ikeCryptoProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*scm.Client)
}

// Create resource.
func (r *ikeCryptoProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state ikeCryptoProfileRsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name":               "scm_ike_crypto_profile",
		"terraform_provider_function": "Create",
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
	})

	// Prepare to create the config.
	svc := uIHLJPY.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := uIHLJPY.CreateInput{}

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()
	input.Request = &wugpput.Config{}

	input.Request.AuthenticationMultiple = state.AuthenticationMultiple.ValueInt64Pointer()

	resp.Diagnostics.Append(state.DhGroups.ElementsAs(ctx, &input.Request.DhGroups, false)...)
	//if len(state.DhGroups) != 0 {
	//    input.Request.DhGroups = make([]string, 0, len(state.DhGroups))
	//    for _, var0 := range state.DhGroups {
	//        input.Request.DhGroups = append(input.Request.DhGroups, var0.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.Encryptions.ElementsAs(ctx, &input.Request.Encryptions, false)...)
	//if len(state.Encryptions) != 0 {
	//    input.Request.Encryptions = make([]string, 0, len(state.Encryptions))
	//    for _, var1 := range state.Encryptions {
	//        input.Request.Encryptions = append(input.Request.Encryptions, var1.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(state.Hashes.ElementsAs(ctx, &input.Request.Hashes, false)...)
	//if len(state.Hashes) != 0 {
	//    input.Request.Hashes = make([]string, 0, len(state.Hashes))
	//    for _, var2 := range state.Hashes {
	//        input.Request.Hashes = append(input.Request.Hashes, var2.ValueString())
	//    }
	//}

	if state.Lifetime != nil {
		input.Request.Lifetime = &wugpput.LifetimeObject{}

		input.Request.Lifetime.Days = state.Lifetime.Days.ValueInt64Pointer()

		input.Request.Lifetime.Hours = state.Lifetime.Hours.ValueInt64Pointer()

		input.Request.Lifetime.Minutes = state.Lifetime.Minutes.ValueInt64Pointer()

		input.Request.Lifetime.Seconds = state.Lifetime.Seconds.ValueInt64Pointer()
	}

	input.Request.Name = state.Name.ValueString()

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

	state.AuthenticationMultiple = types.Int64PointerValue(ans.AuthenticationMultiple)

	var3, var4 := types.ListValueFrom(ctx, types.StringType, ans.DhGroups)
	state.DhGroups = var3
	resp.Diagnostics.Append(var4.Errors()...)

	var5, var6 := types.ListValueFrom(ctx, types.StringType, ans.Encryptions)
	state.Encryptions = var5
	resp.Diagnostics.Append(var6.Errors()...)

	var7, var8 := types.ListValueFrom(ctx, types.StringType, ans.Hashes)
	state.Hashes = var7
	resp.Diagnostics.Append(var8.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	if ans.Lifetime == nil {
		state.Lifetime = nil
	} else {
		state.Lifetime = &ikeCryptoProfileRsModel_wugpput_LifetimeObject{}

		state.Lifetime.Days = types.Int64PointerValue(ans.Lifetime.Days)

		state.Lifetime.Hours = types.Int64PointerValue(ans.Lifetime.Hours)

		state.Lifetime.Minutes = types.Int64PointerValue(ans.Lifetime.Minutes)

		state.Lifetime.Seconds = types.Int64PointerValue(ans.Lifetime.Seconds)
	}

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read performs Read for the struct.
func (r *ikeCryptoProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var savestate, state ikeCryptoProfileRsModel
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
		"resource_name":               "scm_ike_crypto_profile",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	// Prepare to read the config.
	svc := uIHLJPY.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := uIHLJPY.ReadInput{}

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

	state.AuthenticationMultiple = types.Int64PointerValue(ans.AuthenticationMultiple)

	var0, var1 := types.ListValueFrom(ctx, types.StringType, ans.DhGroups)
	state.DhGroups = var0
	resp.Diagnostics.Append(var1.Errors()...)

	var2, var3 := types.ListValueFrom(ctx, types.StringType, ans.Encryptions)
	state.Encryptions = var2
	resp.Diagnostics.Append(var3.Errors()...)

	var4, var5 := types.ListValueFrom(ctx, types.StringType, ans.Hashes)
	state.Hashes = var4
	resp.Diagnostics.Append(var5.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	if ans.Lifetime == nil {
		state.Lifetime = nil
	} else {
		state.Lifetime = &ikeCryptoProfileRsModel_wugpput_LifetimeObject{}

		state.Lifetime.Days = types.Int64PointerValue(ans.Lifetime.Days)

		state.Lifetime.Hours = types.Int64PointerValue(ans.Lifetime.Hours)

		state.Lifetime.Minutes = types.Int64PointerValue(ans.Lifetime.Minutes)

		state.Lifetime.Seconds = types.Int64PointerValue(ans.Lifetime.Seconds)
	}

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update performs the Update for the struct.
func (r *ikeCryptoProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ikeCryptoProfileRsModel
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
		"resource_name":               "scm_ike_crypto_profile",
		"tfid":                        state.Tfid.ValueString(),
	})

	// Prepare to update the config.
	svc := uIHLJPY.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := uIHLJPY.UpdateInput{}

	if tokens[3] != "" {
		input.Id = tokens[3]
	}
	input.Request = &wugpput.Config{}

	input.Request.AuthenticationMultiple = plan.AuthenticationMultiple.ValueInt64Pointer()

	resp.Diagnostics.Append(plan.DhGroups.ElementsAs(ctx, &input.Request.DhGroups, false)...)
	//if len(plan.DhGroups) != 0 {
	//    input.Request.DhGroups = make([]string, 0, len(plan.DhGroups))
	//    for _, var0 := range plan.DhGroups {
	//        input.Request.DhGroups = append(input.Request.DhGroups, var0.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.Encryptions.ElementsAs(ctx, &input.Request.Encryptions, false)...)
	//if len(plan.Encryptions) != 0 {
	//    input.Request.Encryptions = make([]string, 0, len(plan.Encryptions))
	//    for _, var1 := range plan.Encryptions {
	//        input.Request.Encryptions = append(input.Request.Encryptions, var1.ValueString())
	//    }
	//}

	resp.Diagnostics.Append(plan.Hashes.ElementsAs(ctx, &input.Request.Hashes, false)...)
	//if len(plan.Hashes) != 0 {
	//    input.Request.Hashes = make([]string, 0, len(plan.Hashes))
	//    for _, var2 := range plan.Hashes {
	//        input.Request.Hashes = append(input.Request.Hashes, var2.ValueString())
	//    }
	//}

	if plan.Lifetime != nil {
		input.Request.Lifetime = &wugpput.LifetimeObject{}

		input.Request.Lifetime.Days = plan.Lifetime.Days.ValueInt64Pointer()

		input.Request.Lifetime.Hours = plan.Lifetime.Hours.ValueInt64Pointer()

		input.Request.Lifetime.Minutes = plan.Lifetime.Minutes.ValueInt64Pointer()

		input.Request.Lifetime.Seconds = plan.Lifetime.Seconds.ValueInt64Pointer()
	}

	input.Request.Name = plan.Name.ValueString()

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

	state.AuthenticationMultiple = types.Int64PointerValue(ans.AuthenticationMultiple)

	var3, var4 := types.ListValueFrom(ctx, types.StringType, ans.DhGroups)
	state.DhGroups = var3
	resp.Diagnostics.Append(var4.Errors()...)

	var5, var6 := types.ListValueFrom(ctx, types.StringType, ans.Encryptions)
	state.Encryptions = var5
	resp.Diagnostics.Append(var6.Errors()...)

	var7, var8 := types.ListValueFrom(ctx, types.StringType, ans.Hashes)
	state.Hashes = var7
	resp.Diagnostics.Append(var8.Errors()...)

	state.Id = types.StringPointerValue(ans.Id)

	if ans.Lifetime == nil {
		state.Lifetime = nil
	} else {
		state.Lifetime = &ikeCryptoProfileRsModel_wugpput_LifetimeObject{}

		state.Lifetime.Days = types.Int64PointerValue(ans.Lifetime.Days)

		state.Lifetime.Hours = types.Int64PointerValue(ans.Lifetime.Hours)

		state.Lifetime.Minutes = types.Int64PointerValue(ans.Lifetime.Minutes)

		state.Lifetime.Seconds = types.Int64PointerValue(ans.Lifetime.Seconds)
	}

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete performs delete for the struct.
func (r *ikeCryptoProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
		"resource_name":               "scm_ike_crypto_profile",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	svc := uIHLJPY.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := uIHLJPY.DeleteInput{}

	input.Id = tokens[3]

	// Perform the operation.
	if _, err := svc.Delete(ctx, input); err != nil && !IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error in delete", err.Error())
	}
}

func (r *ikeCryptoProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
