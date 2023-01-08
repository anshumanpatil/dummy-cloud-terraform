package bucket

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
	_ datasource.DataSource              = &bucketDataSource{}
	_ datasource.DataSourceWithConfigure = &bucketDataSource{}
)

// NewBucketDataSource is a helper function to simplify the provider implementation.
func NewBucketDataSource() datasource.DataSource {
	return &bucketDataSource{}
}

// bucketDataSource is the data source implementation.
type bucketDataSource struct {
	client *dummycloudclient.Client
}

// bucketDataSourceModel maps the data source schema data.
type bucketDataSourceModel struct {
	Instances []bucketModel `tfsdk:"buckets"`
	ID        types.String  `tfsdk:"id"`
}

// bucketModel maps bucket schema data.
type bucketModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Size   types.String `tfsdk:"size"`
	Region types.String `tfsdk:"region"`
}

// Metadata returns the data source type name.
func (d *bucketDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bucket"
}

// Schema defines the schema for the data source.
func (d *bucketDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches the list of buckets.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"buckets": schema.ListNestedAttribute{
				Description: "List of buckets.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID identifier of the bucket.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Product name of the bucket.",
							Computed:    true,
						},
						"size": schema.StringAttribute{
							Description: "Size of the bucket.",
							Computed:    true,
						},
						"region": schema.StringAttribute{
							Description: "Region of the bucket.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *bucketDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*dummycloudclient.Client)
}

// Read refreshes the Terraform state with the latest data.
func (d *bucketDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state bucketDataSourceModel

	buckets, err := d.client.GetBuckets("")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Dummyclod Instances",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, bucket := range buckets {
		bucketState := bucketModel{
			ID:     types.StringValue(bucket.ID),
			Name:   types.StringValue(bucket.Name),
			Size:   types.StringValue(bucket.Size),
			Region: types.StringValue(bucket.Region),
		}
		state.Instances = append(state.Instances, bucketState)
	}

	state.ID = types.StringValue(uuid.New().String())

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
