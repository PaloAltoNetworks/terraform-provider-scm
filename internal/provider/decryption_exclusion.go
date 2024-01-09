package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	xdEvbZX "github.com/paloaltonetworks/scm-go/netsec/schemas/decryption/exclusions"
	wRodOhd "github.com/paloaltonetworks/scm-go/netsec/services/decryptionexclusions"

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

// Data source.
var (
	_ datasource.DataSource              = &decryptionExclusionDataSource{}
	_ datasource.DataSourceWithConfigure = &decryptionExclusionDataSource{}
)

func NewDecryptionExclusionDataSource() datasource.DataSource {
	return &decryptionExclusionDataSource{}
}

type decryptionExclusionDataSource struct {
	client *scm.Client
}

// decryptionExclusionDsModel is the model.
type decryptionExclusionDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Id types.String `tfsdk:"id"`

	// Output.
	Description types.String `tfsdk:"description"`
	// omit input: id
	Name types.String `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (d *decryptionExclusionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_decryption_exclusion"
}

// Schema defines the schema for this data source.
func (d *decryptionExclusionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"id":true} outputs:map[string]bool{"description":true, "id":true, "name":true, "tfid":true} forceNew:map[string]bool{"id":true}
			"description": dsschema.StringAttribute{
				Description: "The Description param.",
				Computed:    true,
			},
			"id": dsschema.StringAttribute{
				Description: "The Id param.",
				Required:    true,
			},
			"name": dsschema.StringAttribute{
				Description: "The Name param.",
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
func (d *decryptionExclusionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *decryptionExclusionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state decryptionExclusionDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source read", map[string]any{
		"data_source_name":            "scm_decryption_exclusion",
		"terraform_provider_function": "Read",
		"id":                          state.Id.ValueString(),
	})

	// Prepare to run the command.
	svc := wRodOhd.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := wRodOhd.ReadInput{}

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

	state.Description = types.StringPointerValue(ans.Description)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Resource.
var (
	_ resource.Resource                = &decryptionExclusionResource{}
	_ resource.ResourceWithConfigure   = &decryptionExclusionResource{}
	_ resource.ResourceWithImportState = &decryptionExclusionResource{}
)

func NewDecryptionExclusionResource() resource.Resource {
	return &decryptionExclusionResource{}
}

type decryptionExclusionResource struct {
	client *scm.Client
}

// decryptionExclusionRsModel is the model.
type decryptionExclusionRsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Description types.String `tfsdk:"description"`
	Device      types.String `tfsdk:"device"`
	Folder      types.String `tfsdk:"folder"`
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Snippet     types.String `tfsdk:"snippet"`

	// Output.
	// omit input: description
	// omit input: id
	// omit input: name
}

// Metadata returns the data source type name.
func (r *decryptionExclusionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_decryption_exclusion"
}

// Schema defines the schema for this data source.
func (r *decryptionExclusionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rsschema.Schema{
		Description: "Retrieves a config item.",

		Attributes: map[string]rsschema.Attribute{
			// inputs:map[string]bool{"description":true, "device":true, "folder":true, "id":true, "name":true, "snippet":true} outputs:map[string]bool{"description":true, "id":true, "name":true, "tfid":true} forceNew:map[string]bool{"device":true, "folder":true, "snippet":true}
			"description": rsschema.StringAttribute{
				Description: "The Description param.",
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
			"id": rsschema.StringAttribute{
				Description: "UUID of the resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": rsschema.StringAttribute{
				Description: "The Name param.",
				Required:    true,
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
func (r *decryptionExclusionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*scm.Client)
}

// Create resource.
func (r *decryptionExclusionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state decryptionExclusionRsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name":               "scm_decryption_exclusion",
		"terraform_provider_function": "Create",
		"folder":                      state.Folder.ValueString(),
		"snippet":                     state.Snippet.ValueString(),
		"device":                      state.Device.ValueString(),
	})

	// Prepare to create the config.
	svc := wRodOhd.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := wRodOhd.CreateInput{}

	input.Folder = state.Folder.ValueStringPointer()

	input.Snippet = state.Snippet.ValueStringPointer()

	input.Device = state.Device.ValueStringPointer()
	input.Request = &xdEvbZX.Config{}

	input.Request.Description = state.Description.ValueStringPointer()

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

	state.Description = types.StringPointerValue(ans.Description)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read performs Read for the struct.
func (r *decryptionExclusionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var savestate, state decryptionExclusionRsModel
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
		"resource_name":               "scm_decryption_exclusion",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	// Prepare to read the config.
	svc := wRodOhd.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := wRodOhd.ReadInput{}

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

	state.Description = types.StringPointerValue(ans.Description)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update performs the Update for the struct.
func (r *decryptionExclusionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state decryptionExclusionRsModel
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
		"resource_name":               "scm_decryption_exclusion",
		"tfid":                        state.Tfid.ValueString(),
	})

	// Prepare to update the config.
	svc := wRodOhd.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := wRodOhd.UpdateInput{}

	if tokens[3] != "" {
		input.Id = tokens[3]
	}
	input.Request = &xdEvbZX.Config{}

	input.Request.Description = plan.Description.ValueStringPointer()

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

	state.Description = types.StringPointerValue(ans.Description)

	state.Id = types.StringPointerValue(ans.Id)

	state.Name = types.StringValue(ans.Name)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete performs delete for the struct.
func (r *decryptionExclusionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
		"resource_name":               "scm_decryption_exclusion",
		"locMap":                      map[string]int{"device": 2, "folder": 0, "id": 3, "snippet": 1},
		"tokens":                      tokens,
	})

	svc := wRodOhd.NewClient(r.client)

	// Prepare input for the API endpoint.
	input := wRodOhd.DeleteInput{}

	input.Id = tokens[3]

	// Perform the operation.
	if _, err := svc.Delete(ctx, input); err != nil && !IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error in delete", err.Error())
	}
}

func (r *decryptionExclusionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}