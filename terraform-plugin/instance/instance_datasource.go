package instance

import (
	"context"

	"dummy-cloud/dummycloudclient"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &instanceDataSource{}
	_ datasource.DataSourceWithConfigure = &instanceDataSource{}
)

// NewInstanceDataSource is a helper function to simplify the provider implementation.
func NewInstanceDataSource() datasource.DataSource {
	return &instanceDataSource{}
}

// instanceDataSource is the data source implementation.
type instanceDataSource struct {
	client *dummycloudclient.Client
}

// instanceDataSourceModel maps the data source schema data.
type instanceDataSourceModel struct {
	Instances []instanceModel `tfsdk:"instances"`
	ID        types.String    `tfsdk:"id"`
}

// instanceModel maps instance schema data.
type instanceModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Size   types.String `tfsdk:"size"`
	Region types.String `tfsdk:"region"`
	Ram    types.String `tfsdk:"ram"`
	OS     types.String `tfsdk:"os"`
}

// Metadata returns the data source type name.
func (d *instanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance"
}

// Schema defines the schema for the data source.
func (d *instanceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches the list of instances.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"instances": schema.ListNestedAttribute{
				Description: "List of instances.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID identifier of the instance.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Product name of the instance.",
							Computed:    true,
						},
						"size": schema.StringAttribute{
							Description: "Size of the instance.",
							Computed:    true,
						},
						"region": schema.StringAttribute{
							Description: "Region of the instance.",
							Computed:    true,
						},
						"ram": schema.StringAttribute{
							Description: "RAM of the instance.",
							Computed:    true,
						},
						"os": schema.StringAttribute{
							Description: "OS of the instance.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *instanceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*dummycloudclient.Client)
}

// Read refreshes the Terraform state with the latest data.
func (d *instanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state instanceDataSourceModel

	instances, err := d.client.GetInstances("")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Dummyclod Instances",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, instance := range instances {
		instanceState := instanceModel{
			ID:     types.StringValue(instance.ID),
			Name:   types.StringValue(instance.Name),
			Size:   types.StringValue(instance.Size),
			Region: types.StringValue(instance.Region),
			Ram:    types.StringValue(instance.Ram),
			OS:     types.StringValue(instance.OS),
		}
		state.Instances = append(state.Instances, instanceState)
	}

	state.ID = types.StringValue(uuid.New().String())

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
