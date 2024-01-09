package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	mhJDwSQ "github.com/paloaltonetworks/scm-go/netsec/schemas/regions"
	aUYGFNI "github.com/paloaltonetworks/scm-go/netsec/services/regions"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	_ datasource.DataSource              = &regionListDataSource{}
	_ datasource.DataSourceWithConfigure = &regionListDataSource{}
)

func NewRegionListDataSource() datasource.DataSource {
	return &regionListDataSource{}
}

type regionListDataSource struct {
	client *scm.Client
}

// regionListDsModel is the model.
type regionListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Device  types.String `tfsdk:"device"`
	Folder  types.String `tfsdk:"folder"`
	Limit   types.Int64  `tfsdk:"limit"`
	Name    types.String `tfsdk:"name"`
	Offset  types.Int64  `tfsdk:"offset"`
	Snippet types.String `tfsdk:"snippet"`

	// Output.
	Data []regionListDsModel_mhJDwSQ_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type regionListDsModel_mhJDwSQ_Config struct {
	Addresses   types.List                                   `tfsdk:"addresses"`
	GeoLocation *regionListDsModel_mhJDwSQ_GeoLocationObject `tfsdk:"geo_location"`
	Id          types.String                                 `tfsdk:"id"`
	Name        types.String                                 `tfsdk:"name"`
}

type regionListDsModel_mhJDwSQ_GeoLocationObject struct {
	Latitude  types.Float64 `tfsdk:"latitude"`
	Longitude types.Float64 `tfsdk:"longitude"`
}

// Metadata returns the data source type name.
func (d *regionListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region_list"
}

// Schema defines the schema for this listing data source.
func (d *regionListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"device":true, "folder":true, "limit":true, "name":true, "offset":true, "snippet":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"address":true, "geo_location":true, "id":true, "name":true} forceNew:map[string]bool(nil)
						"addresses": dsschema.ListAttribute{
							Description: "The Addresses param.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"geo_location": dsschema.SingleNestedAttribute{
							Description: "The GeoLocation param.",
							Computed:    true,
							Attributes: map[string]dsschema.Attribute{
								// inputs:map[string]bool{} outputs:map[string]bool{"latitude":true, "longitude":true} forceNew:map[string]bool(nil)
								"latitude": dsschema.Float64Attribute{
									Description: "latitude coordinate. Value must be between -90 and 90.",
									Computed:    true,
								},
								"longitude": dsschema.Float64Attribute{
									Description: "longitude coordinate. Value must be between -180 and 180.",
									Computed:    true,
								},
							},
						},
						"id": dsschema.StringAttribute{
							Description: "UUID of the resource.",
							Computed:    true,
						},
						"name": dsschema.StringAttribute{
							Description: "Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 31 characters.",
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
func (d *regionListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *regionListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state regionListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_region_list",
		"terraform_provider_function": "Read",
		"name":                        state.Name.ValueString(),
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
		"offset":                      state.Offset.ValueInt64(),
		"limit":                       state.Limit.ValueInt64(),
	})

	// Prepare to run the command.
	svc := aUYGFNI.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := aUYGFNI.ListInput{}

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
		state.Data = make([]regionListDsModel_mhJDwSQ_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := regionListDsModel_mhJDwSQ_Config{}

			var2, var3 := types.ListValueFrom(ctx, types.StringType, var0.Addresses)
			var1.Addresses = var2
			resp.Diagnostics.Append(var3.Errors()...)

			if var0.GeoLocation == nil {
				var1.GeoLocation = nil
			} else {
				var1.GeoLocation = &regionListDsModel_mhJDwSQ_GeoLocationObject{}

				var1.GeoLocation.Latitude = types.Float64Value(var0.GeoLocation.Latitude)

				var1.GeoLocation.Longitude = types.Float64Value(var0.GeoLocation.Longitude)
			}

			var1.Id = types.StringPointerValue(var0.Id)

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
	_ datasource.DataSource              = &regionDataSource{}
	_ datasource.DataSourceWithConfigure = &regionDataSource{}
)

func NewRegionDataSource() datasource.DataSource {
	return &regionDataSource{}
}

type regionDataSource struct {
	client *scm.Client
}

// regionDsModel is the model.
type regionDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Id types.String `tfsdk:"id"`

	// Output.
	Addresses   types.List                               `tfsdk:"addresses"`
	GeoLocation *regionDsModel_mhJDwSQ_GeoLocationObject `tfsdk:"geo_location"`
	// omit input: id
	Name types.String `tfsdk:"name"`
}

type regionDsModel_mhJDwSQ_GeoLocationObject struct {
	Latitude  types.Float64 `tfsdk:"latitude"`
	Longitude types.Float64 `tfsdk:"longitude"`
}

// Metadata returns the data source type name.
func (d *regionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

// Schema defines the schema for this data source.
func (d *regionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"id":true} outputs:map[string]bool{"address":true, "geo_location":true, "id":true, "name":true, "tfid":true} forceNew:map[string]bool{"id":true}
			"addresses": dsschema.ListAttribute{
				Description: "The Addresses param.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"geo_location": dsschema.SingleNestedAttribute{
				Description: "The GeoLocation param.",
				Computed:    true,
				Attributes: map[string]dsschema.Attribute{
					// inputs:map[string]bool{} outputs:map[string]bool{"latitude":true, "longitude":true} forceNew:map[string]bool(nil)
					"latitude": dsschema.Float64Attribute{
						Description: "latitude coordinate. Value must be between -90 and 90.",
						Computed:    true,
					},
					"longitude": dsschema.Float64Attribute{
						Description: "longitude coordinate. Value must be between -180 and 180.",
						Computed:    true,
					},
				},
			},
			"id": dsschema.StringAttribute{
				Description: "The Id param.",
				Required:    true,
			},
			"name": dsschema.StringAttribute{
				Description: "Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 31 characters.",
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
func (d *regionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *regionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state regionDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source read", map[string]any{
		"data_source_name":            "scm_region",
		"terraform_provider_function": "Read",
		"id":                          state.Id.ValueString(),
	})

	// Prepare to run the command.
	svc := aUYGFNI.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := aUYGFNI.ReadInput{}

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

	var0, var1 := types.ListValueFrom(ctx, types.StringType, ans.Addresses)
	state.Addresses = var0
	resp.Diagnostics.Append(var1.Errors()...)

	if ans.GeoLocation == nil {
		state.GeoLocation = nil
	} else {
		state.GeoLocation = &regionDsModel_mhJDwSQ_GeoLocationObject{}

		state.GeoLocation.Latitude = types.Float64Value(ans.GeoLocation.Latitude)

		state.GeoLocation.Longitude = types.Float64Value(ans.GeoLocation.Longitude)
	}

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Resource.
var (
	_ resource.Resource                = &regionResource{}
	_ resource.ResourceWithConfigure   = &regionResource{}
	_ resource.ResourceWithImportState = &regionResource{}
)

func NewRegionResource() resource.Resource {
	return &regionResource{}
}

type regionResource struct {
	client *scm.Client
}

// regionRsModel is the model.
type regionRsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Addresses   types.List                               `tfsdk:"addresses"`
	Device      types.String                             `tfsdk:"device"`
	Folder      types.String                             `tfsdk:"folder"`
	GeoLocation *regionRsModel_mhJDwSQ_GeoLocationObject `tfsdk:"geo_location"`
	Id          types.String                             `tfsdk:"id"`
	Name        types.String                             `tfsdk:"name"`
	Snippet     types.String                             `tfsdk:"snippet"`

	// Output.
	// omit input: addresses
	// omit input: geo_location
	// omit input: id
	// omit input: name
}

type regionRsModel_mhJDwSQ_GeoLocationObject struct {
	Latitude  types.Float64 `tfsdk:"latitude"`
	Longitude types.Float64 `tfsdk:"longitude"`
}

// Metadata returns the data source type name.
func (r *regionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

// Schema defines the schema for this data source.
func (r *regionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]rsschema.Attribute{
			// inputs:map[string]bool{"address":true, "device":true, "folder":true, "geo_location":true, "id":true, "name":true, "snippet":true} outputs:map[string]bool{"address":true, "geo_location":true, "id":true, "name":true, "tfid":true} forceNew:map[string]bool{"device":true, "folder":true, "snippet":true}
			"addresses": rsschema.ListAttribute{
				Description: "The Addresses param.",
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
			"folder": rsschema.StringAttribute{
				Description: "The Folder param.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"geo_location": rsschema.SingleNestedAttribute{
				Description: "The GeoLocation param.",
				Optional:    true,
				Attributes: map[string]rsschema.Attribute{
					// inputs:map[string]bool{"latitude":true, "longitude":true} outputs:map[string]bool{"latitude":true, "longitude":true} forceNew:map[string]bool(nil)
					"latitude": rsschema.Float64Attribute{
						Description: "latitude coordinate. Value must be between -90 and 90.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(-90, 90),
						},
					},
					"longitude": rsschema.Float64Attribute{
						Description: "longitude coordinate. Value must be between -180 and 180.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(-180, 180),
						},
					},
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
				Description: "Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 31 characters.",
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
func (r *regionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*scm.Client)
}

// Create resource.
func (r *regionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state regionRsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name":               "scm_region",
		"terraform_provider_function": "Create",
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
	})

	// Prepare to create the config.
	svc := aUYGFNI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := aUYGFNI.CreateInput{}

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()
	input.Request = &mhJDwSQ.Config{}

	resp.Diagnostics.Append(state.Addresses.ElementsAs(ctx, &input.Request.Addresses, false)...)
	//if len(state.Addresses) != 0 {
	//    input.Request.Addresses = make([]string, 0, len(state.Addresses))
	//    for _, var0 := range state.Addresses {
	//        input.Request.Addresses = append(input.Request.Addresses, var0.ValueString())
	//    }
	//}

	if state.GeoLocation != nil {
		input.Request.GeoLocation = &mhJDwSQ.GeoLocationObject{}

		input.Request.GeoLocation.Latitude = state.GeoLocation.Latitude.ValueFloat64()

		input.Request.GeoLocation.Longitude = state.GeoLocation.Longitude.ValueFloat64()
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

	var1, var2 := types.ListValueFrom(ctx, types.StringType, ans.Addresses)
	state.Addresses = var1
	resp.Diagnostics.Append(var2.Errors()...)

	if ans.GeoLocation == nil {
		state.GeoLocation = nil
	} else {
		state.GeoLocation = &regionRsModel_mhJDwSQ_GeoLocationObject{}

		state.GeoLocation.Latitude = types.Float64Value(ans.GeoLocation.Latitude)

		state.GeoLocation.Longitude = types.Float64Value(ans.GeoLocation.Longitude)
	}

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read performs Read for the struct.
func (r *regionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var savestate, state regionRsModel
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
		"resource_name":               "scm_region",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	// Prepare to read the config.
	svc := aUYGFNI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := aUYGFNI.ReadInput{}

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

	var0, var1 := types.ListValueFrom(ctx, types.StringType, ans.Addresses)
	state.Addresses = var0
	resp.Diagnostics.Append(var1.Errors()...)

	if ans.GeoLocation == nil {
		state.GeoLocation = nil
	} else {
		state.GeoLocation = &regionRsModel_mhJDwSQ_GeoLocationObject{}

		state.GeoLocation.Latitude = types.Float64Value(ans.GeoLocation.Latitude)

		state.GeoLocation.Longitude = types.Float64Value(ans.GeoLocation.Longitude)
	}

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update performs the Update for the struct.
func (r *regionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state regionRsModel
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
		"resource_name":               "scm_region",
		"tfid":                        state.Tfid.ValueString(),
	})

	// Prepare to update the config.
	svc := aUYGFNI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := aUYGFNI.UpdateInput{}

	if tokens[3] != "" {
		input.Id = tokens[3]
	}
	input.Request = &mhJDwSQ.Config{}

	resp.Diagnostics.Append(plan.Addresses.ElementsAs(ctx, &input.Request.Addresses, false)...)
	//if len(plan.Addresses) != 0 {
	//    input.Request.Addresses = make([]string, 0, len(plan.Addresses))
	//    for _, var0 := range plan.Addresses {
	//        input.Request.Addresses = append(input.Request.Addresses, var0.ValueString())
	//    }
	//}

	if plan.GeoLocation != nil {
		input.Request.GeoLocation = &mhJDwSQ.GeoLocationObject{}

		input.Request.GeoLocation.Latitude = plan.GeoLocation.Latitude.ValueFloat64()

		input.Request.GeoLocation.Longitude = plan.GeoLocation.Longitude.ValueFloat64()
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

	var1, var2 := types.ListValueFrom(ctx, types.StringType, ans.Addresses)
	state.Addresses = var1
	resp.Diagnostics.Append(var2.Errors()...)

	if ans.GeoLocation == nil {
		state.GeoLocation = nil
	} else {
		state.GeoLocation = &regionRsModel_mhJDwSQ_GeoLocationObject{}

		state.GeoLocation.Latitude = types.Float64Value(ans.GeoLocation.Latitude)

		state.GeoLocation.Longitude = types.Float64Value(ans.GeoLocation.Longitude)
	}

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete performs delete for the struct.
func (r *regionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
		"resource_name":               "scm_region",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	svc := aUYGFNI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := aUYGFNI.DeleteInput{}

	input.Id = tokens[3]

	// Perform the operation.
	if _, err := svc.Delete(ctx, input); err != nil && !IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error in delete", err.Error())
	}
}

func (r *regionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}