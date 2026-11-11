package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/cloud_ngfw"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/cloud_ngfw"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &NgfirewallListDataSource{}
	_ datasource.DataSourceWithConfigure = &NgfirewallListDataSource{}
)

func NewNgfirewallListDataSource() datasource.DataSource {
	return &NgfirewallListDataSource{}
}

// NgfirewallListDataSource defines the data source implementation.
type NgfirewallListDataSource struct {
	client *cloud_ngfw.APIClient
}

func (d *NgfirewallListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "scm_cngfw_ngfirewall_list"
}

func (d *NgfirewallListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.NgfirewallsListDataSourceSchema
}

func (d *NgfirewallListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["cloud_ngfw"].(*cloud_ngfw.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *cloud_ngfw.APIClient for 'cloud_ngfw' client."))
		return
	}
	d.client = client
}

func (d *NgfirewallListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.NgfirewallsListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API request.
	// For Cloud NGFW, ListCloudNGFirewalls fetches all pages transparently.
	listResponse, _, err := d.client.NgfirewallsAPI.ListCloudNGFirewalls(ctx,
		data.Region.ValueString(),
		data.Globalrulestackname.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError("Error Listing Ngfirewallss", fmt.Sprintf("Could not list Ngfirewallss: %s", err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"Resource Listing Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// NON-STANDARD ENVELOPE LOGIC (Cloud NGFW):
	// The API wraps items inside Response.FirewallsDescribe with NextToken pagination.
	if listResponse == nil {
		return // Nothing to do.
	}

	innerResponse := listResponse.GetResponse()

	firewallDescribes := innerResponse.GetFirewallsDescribe()

	totalCount := int64(len(firewallDescribes))
	data.Total = types.Int64PointerValue(&totalCount)

	packedList, diags := packNgfirewallResponseBodyListFromSdk(ctx, firewallDescribes)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedList.ElementsAs(ctx, &data.FirewallsDescribe, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create a unique TFID for this data source.
	var idBuilder strings.Builder

	// Folder, snippet params based TFID is not applicable to CNGFW resources. So, we will use the query parameters \
	// to create a unique TFID for this data source.
	idBuilder.WriteString(data.Region.ValueString())
	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Globalrulestackname.ValueString())
	h := sha256.Sum256([]byte(idBuilder.String()))
	data.Tfid = types.StringValue("cngfw:" + hex.EncodeToString(h[:]))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
