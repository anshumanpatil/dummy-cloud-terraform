package bucket

import (
	"context"
	"time"

	"dummy-cloud/dummycloudclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &bucketResource{}
	_ resource.ResourceWithConfigure   = &bucketResource{}
	_ resource.ResourceWithImportState = &bucketResource{}
)

// NewBucketResource is a helper function to simplify the provider implementation.
func NewBucketResource() resource.Resource {
	return &bucketResource{}
}

// bucketResource is the resource implementation.
type bucketResource struct {
	client *dummycloudclient.Client
}

// bucketSchemaModel maps coffee order item data.
type bucketSchemaModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Size        types.String `tfsdk:"size"`
	Region      types.String `tfsdk:"region"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

// Metadata returns the data source type name.
func (r *bucketResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bucket"
}

// Schema defines the schema for the data source.
func (r *bucketResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an bucket.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Numeric identifier of the order.",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Numeric identifier of the order.",
				Required:    true,
			},
			"size": schema.StringAttribute{
				Description: "Numeric identifier of the order.",
				Required:    true,
			},
			"region": schema.StringAttribute{
				Description: "Numeric identifier of the order.",
				Required:    true,
			},
			"last_updated": schema.StringAttribute{
				Description: "Timestamp of the last Terraform update of the order.",
				Computed:    true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *bucketResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*dummycloudclient.Client)
}

// Create creates the resource and sets the initial Terraform state.
func (r *bucketResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan bucketSchemaModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bucketToCreate := dummycloudclient.Bucket{
		Name:   plan.Name.ValueString(),
		Size:   plan.Size.ValueString(),
		Region: plan.Region.ValueString(),
	}

	// Create new order
	order, err := r.client.CreateBucket(bucketToCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(order.ID)
	plan.Name = types.StringValue(order.Name)
	plan.Size = types.StringValue(order.Size)
	plan.Region = types.StringValue(order.Region)

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *bucketResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state bucketSchemaModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from DummyCloud
	order, err := r.client.GetBuckets(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading DummyCloud Order",
			"Could not read DummyCloud order ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	if len(order) <= 0 {
		resp.Diagnostics.AddError(
			"Error Reading DummyCloud Order",
			"Could not read DummyCloud order ID "+state.ID.ValueString(),
		)
		return
	}

	singleBucket := order[0]

	state.ID = types.StringValue(singleBucket.ID)
	state.Name = types.StringValue(singleBucket.Name)
	state.Size = types.StringValue(singleBucket.Size)
	state.Region = types.StringValue(singleBucket.Region)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *bucketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan bucketSchemaModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := dummycloudclient.Bucket{
		ID:     plan.ID.ValueString(),
		Name:   plan.Name.ValueString(),
		Size:   plan.Size.ValueString(),
		Region: plan.Region.ValueString(),
	}

	// Update existing order
	order, err := r.client.UpdateBucket(input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating DummyCloud Order",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(order.ID)
	plan.Name = types.StringValue(order.Name)
	plan.Size = types.StringValue(order.Size)
	plan.Region = types.StringValue(order.Region)

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *bucketResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state bucketSchemaModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	_, err := r.client.DeleteBucket(state.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting DummyCloud Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *bucketResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
