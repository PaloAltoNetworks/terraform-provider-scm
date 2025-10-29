package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/identity_services"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &LdapServerProfileListDataSource{}
	_ datasource.DataSourceWithConfigure = &LdapServerProfileListDataSource{}
)

func NewLdapServerProfileListDataSource() datasource.DataSource {
	return &LdapServerProfileListDataSource{}
}

// LdapServerProfileListDataSource defines the data source implementation.
type LdapServerProfileListDataSource struct {
	client *identity_services.APIClient
}

func (d *LdapServerProfileListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ldap_server_profile_list"
}

func (d *LdapServerProfileListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.LdapServerProfilesListDataSourceSchema
}

func (d *LdapServerProfileListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["identity_services"].(*identity_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *identity_services.APIClient for 'identity_services' client."))
		return
	}
	d.client = client
}

func (d *LdapServerProfileListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.LdapServerProfilesListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API request.
	listReq := d.client.LDAPServerProfilesAPI.ListLDAPServerProfiles(ctx)

	// Apply filters from the configuration.

	v := reflect.ValueOf(data)

	if f := v.FieldByName("Folder"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			listReq = listReq.Folder(val.ValueString())
		}
	}

	if f := v.FieldByName("Snippet"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			listReq = listReq.Snippet(val.ValueString())
		}
	}

	if f := v.FieldByName("Device"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			listReq = listReq.Device(val.ValueString())
		}
	}

	if !data.Limit.IsNull() {
		listReq = listReq.Limit(int32(data.Limit.ValueInt64()))
	}
	if !data.Offset.IsNull() {
		listReq = listReq.Offset(int32(data.Offset.ValueInt64()))
	}

	// Execute the request.
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Listing LdapServerProfiless", fmt.Sprintf("Could not list LdapServerProfiless: %s", err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"Tag Listing Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// Convert the response to the Terraform model.
	if listResponse == nil || listResponse.GetData() == nil {
		return // Nothing to do.
	}

	total := int64(listResponse.GetTotal())
	data.Total = types.Int64PointerValue(&total)
	data.Limit = types.Int64Value(int64(listResponse.GetLimit()))
	data.Offset = types.Int64Value(int64(listResponse.GetOffset()))

	// =================== START: THE IMPROVEMENT ===================
	// Use the generated list packer to pack the SCM items into a TF list.
	packedList, diags := packLdapServerProfilesListFromSdk(ctx, listResponse.GetData())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert the TF list to a Go slice of the model and set it to the data.
	resp.Diagnostics.Append(packedList.ElementsAs(ctx, &data.Data, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// ===================  END: THE IMPROVEMENT  ===================

	// Create a unique TFID for this data source.
	var idBuilder strings.Builder

	// Use reflection again for Tfid creation to ensure safety

	v = reflect.ValueOf(data)

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
	if !data.Name.IsNull() {
		idBuilder.WriteString(data.Name.ValueString())
	}
	idBuilder.WriteString(":")
	if !data.Limit.IsNull() {
		idBuilder.WriteString(strconv.FormatInt(data.Limit.ValueInt64(), 10))
	}
	idBuilder.WriteString(":")
	if !data.Offset.IsNull() {
		idBuilder.WriteString(strconv.FormatInt(data.Offset.ValueInt64(), 10))
	}
	data.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
