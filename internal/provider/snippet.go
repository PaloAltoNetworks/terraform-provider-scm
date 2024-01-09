package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	ujXZojh "github.com/paloaltonetworks/scm-go/netsec/schemas/snippets"
	zAHtTyI "github.com/paloaltonetworks/scm-go/netsec/services/snippets"

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
	_ datasource.DataSource              = &snippetListDataSource{}
	_ datasource.DataSourceWithConfigure = &snippetListDataSource{}
)

func NewSnippetListDataSource() datasource.DataSource {
	return &snippetListDataSource{}
}

type snippetListDataSource struct {
	client *scm.Client
}

// snippetListDsModel is the model.
type snippetListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Limit  types.Int64  `tfsdk:"limit"`
	Name   types.String `tfsdk:"name"`
	Offset types.Int64  `tfsdk:"offset"`

	// Output.
	Data []snippetListDsModel_ujXZojh_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type snippetListDsModel_ujXZojh_Config struct {
	CreatedIn   types.String                              `tfsdk:"created_in"`
	Description types.String                              `tfsdk:"description"`
	DisplayName types.String                              `tfsdk:"display_name"`
	Folders     []snippetListDsModel_ujXZojh_FolderObject `tfsdk:"folders"`
	Id          types.String                              `tfsdk:"id"`
	Labels      types.List                                `tfsdk:"labels"`
	LastUpdate  types.String                              `tfsdk:"last_update"`
	Name        types.String                              `tfsdk:"name"`
	SharedIn    types.String                              `tfsdk:"shared_in"`
	Type        types.String                              `tfsdk:"type"`
}

type snippetListDsModel_ujXZojh_FolderObject struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (d *snippetListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet_list"
}

// Schema defines the schema for this listing data source.
func (d *snippetListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"limit":true, "name":true, "offset":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"limit":true, "name":true, "offset":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"created_in":true, "description":true, "display_name":true, "folders":true, "id":true, "labels":true, "last_update":true, "name":true, "shared_in":true, "type":true} forceNew:map[string]bool(nil)
						"created_in": dsschema.StringAttribute{
							Description: "The CreatedIn param.",
							Computed:    true,
						},
						"description": dsschema.StringAttribute{
							Description: "The Description param.",
							Computed:    true,
						},
						"display_name": dsschema.StringAttribute{
							Description: "The DisplayName param.",
							Computed:    true,
						},
						"folders": dsschema.ListNestedAttribute{
							Description: "The Folders param.",
							Computed:    true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									// inputs:map[string]bool{} outputs:map[string]bool{"id":true, "name":true} forceNew:map[string]bool(nil)
									"id": dsschema.StringAttribute{
										Description: "The Id param.",
										Computed:    true,
									},
									"name": dsschema.StringAttribute{
										Description: "The Name param.",
										Computed:    true,
									},
								},
							},
						},
						"id": dsschema.StringAttribute{
							Description: "The Id param.",
							Computed:    true,
						},
						"labels": dsschema.ListAttribute{
							Description: "The Labels param.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"last_update": dsschema.StringAttribute{
							Description: "The LastUpdate param.",
							Computed:    true,
						},
						"name": dsschema.StringAttribute{
							Description: "The Name param.",
							Computed:    true,
						},
						"shared_in": dsschema.StringAttribute{
							Description: "The SharedIn param.",
							Computed:    true,
						},
						"type": dsschema.StringAttribute{
							Description: "The Type param. String must be one of these: `\"predefined\"`.",
							Computed:    true,
						},
					},
				},
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
func (d *snippetListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *snippetListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state snippetListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_snippet_list",
		"terraform_provider_function": "Read",
		"name":                        state.Name.ValueString(),
		"limit":                       state.Limit.ValueInt64(),
		"offset":                      state.Offset.ValueInt64(),
	})

	// Prepare to run the command.
	svc := zAHtTyI.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := zAHtTyI.ListInput{}

	input.Name = state.Name.ValueStringPointer()

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
		state.Data = make([]snippetListDsModel_ujXZojh_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := snippetListDsModel_ujXZojh_Config{}

			var1.CreatedIn = types.StringPointerValue(var0.CreatedIn)

			var1.Description = types.StringPointerValue(var0.Description)

			var1.DisplayName = types.StringPointerValue(var0.DisplayName)

			if len(var0.Folders) == 0 {
				var1.Folders = nil
			} else {
				var1.Folders = make([]snippetListDsModel_ujXZojh_FolderObject, 0, len(var0.Folders))
				for _, var2 := range var0.Folders {
					var3 := snippetListDsModel_ujXZojh_FolderObject{}

					var3.Id = types.StringPointerValue(var2.Id)

					var3.Name = types.StringPointerValue(var2.Name)
					var1.Folders = append(var1.Folders, var3)
				}
			}

			var1.Id = types.StringPointerValue(var0.Id)

			var4, var5 := types.ListValueFrom(ctx, types.StringType, var0.Labels)
			var1.Labels = var4
			resp.Diagnostics.Append(var5.Errors()...)

			var1.LastUpdate = types.StringPointerValue(var0.LastUpdate)

			var1.Name = types.StringValue(var0.Name)

			var1.SharedIn = types.StringPointerValue(var0.SharedIn)

			var1.Type = types.StringPointerValue(var0.Type)
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
	_ datasource.DataSource              = &snippetDataSource{}
	_ datasource.DataSourceWithConfigure = &snippetDataSource{}
)

func NewSnippetDataSource() datasource.DataSource {
	return &snippetDataSource{}
}

type snippetDataSource struct {
	client *scm.Client
}

// snippetDsModel is the model.
type snippetDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Id types.String `tfsdk:"id"`

	// Output.
	CreatedIn   types.String                          `tfsdk:"created_in"`
	Description types.String                          `tfsdk:"description"`
	DisplayName types.String                          `tfsdk:"display_name"`
	Folders     []snippetDsModel_ujXZojh_FolderObject `tfsdk:"folders"`
	// omit input: id
	Labels     types.List   `tfsdk:"labels"`
	LastUpdate types.String `tfsdk:"last_update"`
	Name       types.String `tfsdk:"name"`
	SharedIn   types.String `tfsdk:"shared_in"`
	Type       types.String `tfsdk:"type"`
}

type snippetDsModel_ujXZojh_FolderObject struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (d *snippetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet"
}

// Schema defines the schema for this data source.
func (d *snippetDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"id":true} outputs:map[string]bool{"created_in":true, "description":true, "display_name":true, "folders":true, "id":true, "labels":true, "last_update":true, "name":true, "shared_in":true, "tfid":true, "type":true} forceNew:map[string]bool{"id":true}
			"created_in": dsschema.StringAttribute{
				Description: "The CreatedIn param.",
				Computed:    true,
			},
			"description": dsschema.StringAttribute{
				Description: "The Description param.",
				Computed:    true,
			},
			"display_name": dsschema.StringAttribute{
				Description: "The DisplayName param.",
				Computed:    true,
			},
			"folders": dsschema.ListNestedAttribute{
				Description: "The Folders param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"id":true, "name":true} forceNew:map[string]bool(nil)
						"id": dsschema.StringAttribute{
							Description: "The Id param.",
							Computed:    true,
						},
						"name": dsschema.StringAttribute{
							Description: "The Name param.",
							Computed:    true,
						},
					},
				},
			},
			"id": dsschema.StringAttribute{
				Description: "The Id param.",
				Required:    true,
			},
			"labels": dsschema.ListAttribute{
				Description: "The Labels param.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"last_update": dsschema.StringAttribute{
				Description: "The LastUpdate param.",
				Computed:    true,
			},
			"name": dsschema.StringAttribute{
				Description: "The Name param.",
				Computed:    true,
			},
			"shared_in": dsschema.StringAttribute{
				Description: "The SharedIn param.",
				Computed:    true,
			},
			"tfid": dsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
			},
			"type": dsschema.StringAttribute{
				Description: "The Type param. String must be one of these: `\"predefined\"`.",
				Computed:    true,
			},
		},
	}
}

// Configure prepares the struct.
func (d *snippetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *snippetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state snippetDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source read", map[string]any{
		"data_source_name":            "scm_snippet",
		"terraform_provider_function": "Read",
		"id":                          state.Id.ValueString(),
	})

	// Prepare to run the command.
	svc := zAHtTyI.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := zAHtTyI.ReadInput{}

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

	state.CreatedIn = types.StringPointerValue(ans.CreatedIn)

	state.Description = types.StringPointerValue(ans.Description)

	state.DisplayName = types.StringPointerValue(ans.DisplayName)

	if len(ans.Folders) == 0 {
		state.Folders = nil
	} else {
		state.Folders = make([]snippetDsModel_ujXZojh_FolderObject, 0, len(ans.Folders))
		for _, var0 := range ans.Folders {
			var1 := snippetDsModel_ujXZojh_FolderObject{}

			var1.Id = types.StringPointerValue(var0.Id)

			var1.Name = types.StringPointerValue(var0.Name)
			state.Folders = append(state.Folders, var1)
		}
	}

	state.Id = types.StringPointerValue(ans.Id)

	var2, var3 := types.ListValueFrom(ctx, types.StringType, ans.Labels)
	state.Labels = var2
	resp.Diagnostics.Append(var3.Errors()...)

	state.LastUpdate = types.StringPointerValue(ans.LastUpdate)

	state.Name = types.StringValue(ans.Name)

	state.SharedIn = types.StringPointerValue(ans.SharedIn)

	state.Type = types.StringPointerValue(ans.Type)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Resource.
var (
	_ resource.Resource                = &snippetResource{}
	_ resource.ResourceWithConfigure   = &snippetResource{}
	_ resource.ResourceWithImportState = &snippetResource{}
)

func NewSnippetResource() resource.Resource {
	return &snippetResource{}
}

type snippetResource struct {
	client *scm.Client
}

// snippetRsModel is the model.
type snippetRsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	CreatedIn   types.String                          `tfsdk:"created_in"`
	Description types.String                          `tfsdk:"description"`
	DisplayName types.String                          `tfsdk:"display_name"`
	Folders     []snippetRsModel_ujXZojh_FolderObject `tfsdk:"folders"`
	Id          types.String                          `tfsdk:"id"`
	Labels      types.List                            `tfsdk:"labels"`
	LastUpdate  types.String                          `tfsdk:"last_update"`
	Name        types.String                          `tfsdk:"name"`
	SharedIn    types.String                          `tfsdk:"shared_in"`
	Type        types.String                          `tfsdk:"type"`

	// Output.
	// omit input: created_in
	// omit input: description
	// omit input: display_name
	// omit input: folders
	// omit input: id
	// omit input: labels
	// omit input: last_update
	// omit input: name
	// omit input: shared_in
	// omit input: type
}

type snippetRsModel_ujXZojh_FolderObject struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (r *snippetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet"
}

// Schema defines the schema for this data source.
func (r *snippetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]rsschema.Attribute{
			// inputs:map[string]bool{"created_in":true, "description":true, "display_name":true, "folders":true, "id":true, "labels":true, "last_update":true, "name":true, "shared_in":true, "type":true} outputs:map[string]bool{"created_in":true, "description":true, "display_name":true, "folders":true, "id":true, "labels":true, "last_update":true, "name":true, "shared_in":true, "tfid":true, "type":true} forceNew:map[string]bool{}
			"created_in": rsschema.StringAttribute{
				Description: "The CreatedIn param.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": rsschema.StringAttribute{
				Description: "The Description param.",
				Optional:    true,
			},
			"display_name": rsschema.StringAttribute{
				Description: "The DisplayName param.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"folders": rsschema.ListNestedAttribute{
				Description: "The Folders param.",
				Computed:    true,
				NestedObject: rsschema.NestedAttributeObject{
					Attributes: map[string]rsschema.Attribute{
						// inputs:map[string]bool{"id":true, "name":true} outputs:map[string]bool{"id":true, "name":true} forceNew:map[string]bool(nil)
						"id": rsschema.StringAttribute{
							Description: "The Id param.",
							Optional:    true,
						},
						"name": rsschema.StringAttribute{
							Description: "The Name param.",
							Optional:    true,
						},
					},
				},
			},
			"id": rsschema.StringAttribute{
				Description: "The Id param.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"labels": rsschema.ListAttribute{
				Description: "The Labels param.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"last_update": rsschema.StringAttribute{
				Description: "The LastUpdate param.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": rsschema.StringAttribute{
				Description: "The Name param.",
				Required:    true,
			},
			"shared_in": rsschema.StringAttribute{
				Description: "The SharedIn param.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tfid": rsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": rsschema.StringAttribute{
				Description: "The Type param. String must be one of these: `\"predefined\"`.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("predefined"),
				},
			},
		},
	}
}

// Configure prepares the struct.
func (r *snippetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*scm.Client)
}

// Create resource.
func (r *snippetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state snippetRsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name":               "scm_snippet",
		"terraform_provider_function": "Create",
	})

	// Prepare to create the config.
	svc := zAHtTyI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := zAHtTyI.CreateInput{}

	input.Request = &ujXZojh.Config{}

	input.Request.Description = state.Description.ValueStringPointer()

	resp.Diagnostics.Append(state.Labels.ElementsAs(ctx, &input.Request.Labels, false)...)
	//if len(state.Labels) != 0 {
	//    input.Request.Labels = make([]string, 0, len(state.Labels))
	//    for _, var0 := range state.Labels {
	//        input.Request.Labels = append(input.Request.Labels, var0.ValueString())
	//    }
	//}

	input.Request.Name = state.Name.ValueString()

	// Perform the operation.
	ans, err := svc.Create(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError("Error creating config", err.Error())
		return
	}

	// Create the Terraform ID.
	var idBuilder strings.Builder
	if ans.Id == nil {
		resp.Diagnostics.AddError("Undefined param required for ID", "Id")
		return
	}
	if ans.Id != nil {
		idBuilder.WriteString(*ans.Id)
	}

	// Store the answer to state.

	state.Tfid = types.StringValue(idBuilder.String())

	state.CreatedIn = types.StringPointerValue(ans.CreatedIn)

	state.Description = types.StringPointerValue(ans.Description)

	state.DisplayName = types.StringPointerValue(ans.DisplayName)

	if len(ans.Folders) == 0 {
		state.Folders = nil
	} else {
		state.Folders = make([]snippetRsModel_ujXZojh_FolderObject, 0, len(ans.Folders))
		for _, var1 := range ans.Folders {
			var2 := snippetRsModel_ujXZojh_FolderObject{}

			var2.Id = types.StringPointerValue(var1.Id)

			var2.Name = types.StringPointerValue(var1.Name)
			state.Folders = append(state.Folders, var2)
		}
	}

	state.Id = types.StringPointerValue(ans.Id)

	var3, var4 := types.ListValueFrom(ctx, types.StringType, ans.Labels)
	state.Labels = var3
	resp.Diagnostics.Append(var4.Errors()...)

	state.LastUpdate = types.StringPointerValue(ans.LastUpdate)

	state.Name = types.StringValue(ans.Name)

	state.SharedIn = types.StringPointerValue(ans.SharedIn)

	state.Type = types.StringPointerValue(ans.Type)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read performs Read for the struct.
func (r *snippetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var savestate, state snippetRsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tfid := savestate.Tfid.ValueString()
	tokens := strings.Split(tfid, IdSeparator)
	if len(tokens) != 1 {
		resp.Diagnostics.AddError("Error in resource ID format", "Expected 1 tokens")
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource read", map[string]any{
		"terraform_provider_function": "Read",
		"resource_name":               "scm_snippet",
		"locMap":                      map[string]int{"id": 0},
		"tokens":                      tokens,
	})

	// Prepare to read the config.
	svc := zAHtTyI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := zAHtTyI.ReadInput{}

	input.Id = tokens[0]

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
	state.Tfid = savestate.Tfid

	state.CreatedIn = types.StringPointerValue(ans.CreatedIn)

	state.Description = types.StringPointerValue(ans.Description)

	state.DisplayName = types.StringPointerValue(ans.DisplayName)

	if len(ans.Folders) == 0 {
		state.Folders = nil
	} else {
		state.Folders = make([]snippetRsModel_ujXZojh_FolderObject, 0, len(ans.Folders))
		for _, var0 := range ans.Folders {
			var1 := snippetRsModel_ujXZojh_FolderObject{}

			var1.Id = types.StringPointerValue(var0.Id)

			var1.Name = types.StringPointerValue(var0.Name)
			state.Folders = append(state.Folders, var1)
		}
	}

	state.Id = types.StringPointerValue(ans.Id)

	var2, var3 := types.ListValueFrom(ctx, types.StringType, ans.Labels)
	state.Labels = var2
	resp.Diagnostics.Append(var3.Errors()...)

	state.LastUpdate = types.StringPointerValue(ans.LastUpdate)

	state.Name = types.StringValue(ans.Name)

	state.SharedIn = types.StringPointerValue(ans.SharedIn)

	state.Type = types.StringPointerValue(ans.Type)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update performs the Update for the struct.
func (r *snippetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state snippetRsModel
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
	if len(tokens) != 1 {
		resp.Diagnostics.AddError("Error in resource ID format", "Expected 1 tokens")
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource update", map[string]any{
		"terraform_provider_function": "Update",
		"resource_name":               "scm_snippet",
		"tfid":                        state.Tfid.ValueString(),
	})

	// Prepare to update the config.
	svc := zAHtTyI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := zAHtTyI.UpdateInput{}

	if tokens[0] != "" {
		input.Id = tokens[0]
	}
	input.Request = &ujXZojh.Config{}

	input.Request.Description = plan.Description.ValueStringPointer()

	resp.Diagnostics.Append(plan.Labels.ElementsAs(ctx, &input.Request.Labels, false)...)
	//if len(plan.Labels) != 0 {
	//    input.Request.Labels = make([]string, 0, len(plan.Labels))
	//    for _, var0 := range plan.Labels {
	//        input.Request.Labels = append(input.Request.Labels, var0.ValueString())
	//    }
	//}

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

	state.CreatedIn = types.StringPointerValue(ans.CreatedIn)

	state.Description = types.StringPointerValue(ans.Description)

	state.DisplayName = types.StringPointerValue(ans.DisplayName)

	if len(ans.Folders) == 0 {
		state.Folders = nil
	} else {
		state.Folders = make([]snippetRsModel_ujXZojh_FolderObject, 0, len(ans.Folders))
		for _, var1 := range ans.Folders {
			var2 := snippetRsModel_ujXZojh_FolderObject{}

			var2.Id = types.StringPointerValue(var1.Id)

			var2.Name = types.StringPointerValue(var1.Name)
			state.Folders = append(state.Folders, var2)
		}
	}

	state.Id = types.StringPointerValue(ans.Id)

	var3, var4 := types.ListValueFrom(ctx, types.StringType, ans.Labels)
	state.Labels = var3
	resp.Diagnostics.Append(var4.Errors()...)

	state.LastUpdate = types.StringPointerValue(ans.LastUpdate)

	state.Name = types.StringValue(ans.Name)

	state.SharedIn = types.StringPointerValue(ans.SharedIn)

	state.Type = types.StringPointerValue(ans.Type)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete performs delete for the struct.
func (r *snippetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var idType types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("tfid"), &idType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tfid := idType.ValueString()
	tokens := strings.Split(tfid, IdSeparator)
	if len(tokens) != 1 {
		resp.Diagnostics.AddError("Error in resource ID format", "Expected 1 tokens")
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource delete", map[string]any{
		"terraform_provider_function": "Delete",
		"resource_name":               "scm_snippet",
		"locMap":                      map[string]int{"id": 0},
		"tokens":                      tokens,
	})

	svc := zAHtTyI.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := zAHtTyI.DeleteInput{}

	input.Id = tokens[0]

	// Perform the operation.
	if _, err := svc.Delete(ctx, input); err != nil && !IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error in delete", err.Error())
	}
}

func (r *snippetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}