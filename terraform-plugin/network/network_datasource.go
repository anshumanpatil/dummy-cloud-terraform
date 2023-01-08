package network

import (
	"context"

	"dummy-cloud/dummycloudclient"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &networkDataSource{}
	_ datasource.DataSourceWithConfigure = &networkDataSource{}
)

// NewNetworkDataSource is a helper function to simplify the provider implementation.
func NewNetworkDataSource() datasource.DataSource {
	return &networkDataSource{}
}

// networkDataSource is the data source implementation.
type networkDataSource struct {
	client *dummycloudclient.Client
}

// networkDataSourceModel maps the data source schema data.
type networkDataSourceModel struct {
	Networks []networkModel `tfsdk:"networks"`
	ID       types.String   `tfsdk:"id"`
}

// networkModel maps network schema data.
type networkModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	IPList       types.List   `tfsdk:"iplist"`
	InstanceList types.List   `tfsdk:"instancelist"`
	IsActive     types.Bool   `tfsdk:"isactive"`
}

// Metadata returns the data source type name.
func (d *networkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

// Schema defines the schema for the data source.
func (d *networkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches the list of networks.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"networks": schema.ListNestedAttribute{
				Description: "List of networks.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID identifier of the network.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Product name of the network.",
							Computed:    true,
						},
						"iplist": schema.ListAttribute{
							ElementType: types.StringType,
							Description: "Size of the network.",
							Computed:    true,
						},
						"instancelist": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed: true,
									},
									"region": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"isactive": schema.BoolAttribute{
							Description: "Region of the network.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *networkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*dummycloudclient.Client)
}

// Read refreshes the Terraform state with the latest data.
func (d *networkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state networkDataSourceModel

	networks, err := d.client.GetNetworks("")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Dummyclod Networks",
			err.Error(),
		)
		return
	}

	tflog.Info(ctx, "[Anshuman] "+d.client.PrettyJson(networks))
	// Map response body to model
	for _, network := range networks {

		newiplist, err := getIpList(network)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read IPLIST Dummyclod Networks",
				err.Error(),
			)
			return
		}

		newinstancelist, err := getInstanceList(network)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read InstanceList Dummyclod Networks",
				err.Error(),
			)
			return
		}

		networkState := networkModel{
			ID:           types.StringValue(network.ID),
			Name:         types.StringValue(network.Name),
			IsActive:     types.BoolValue(network.IsActive),
			IPList:       *newiplist,
			InstanceList: *newinstancelist,
		}
		state.Networks = append(state.Networks, networkState)
	}

	state.ID = types.StringValue(uuid.New().String())

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
