package network

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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &networkResource{}
	_ resource.ResourceWithConfigure   = &networkResource{}
	_ resource.ResourceWithImportState = &networkResource{}
)

// NewNetworkResource is a helper function to simplify the provider implementation.
func NewNetworkResource() resource.Resource {
	return &networkResource{}
}

// networkResource is the resource implementation.
type networkResource struct {
	client *dummycloudclient.Client
}

// networkSchemaModel maps coffee order item data.
type networkSchemaModel struct {
	ID           types.String      `tfsdk:"id"`
	Name         types.String      `tfsdk:"name"`
	IPList       []types.String    `tfsdk:"iplist"`
	InstanceList []NetworkInstance `tfsdk:"instancelist"`
	IsActive     types.Bool        `tfsdk:"isactive"`
	LastUpdated  types.String      `tfsdk:"last_updated"`
}

type NetworkInstance struct {
	Name   types.String `tfsdk:"name"`
	Region types.String `tfsdk:"region"`
}

// Metadata returns the data source type name.
func (r *networkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

// Schema defines the schema for the data source.
func (r *networkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an network.",
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
			"iplist": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Size of the network.",
				Optional:    true,
				Computed:    true,
			},
			"instancelist": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"region": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"isactive": schema.BoolAttribute{
				Description: "Region of the network.",
				Optional:    true,
				Computed:    true,
			},
			"last_updated": schema.StringAttribute{
				Description: "Timestamp of the last Terraform update of the order.",
				Computed:    true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *networkResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*dummycloudclient.Client)
}

// Create creates the resource and sets the initial Terraform state.
func (r *networkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan networkSchemaModel
	diags := req.Plan.Get(ctx, &plan)

	tflog.Info(ctx, "[Anshuman] plan "+r.client.PrettyJson(plan.IPList))
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "[Anshuman] "+r.client.PrettyJson(plan.IPList))

	newIPList := []string{}
	for _, v := range plan.IPList {
		newIPList = append(newIPList, v.ValueString())
	}

	newInstanceList := []dummycloudclient.NetworkRegion{}
	for _, v := range plan.InstanceList {
		newInstanceList = append(newInstanceList, dummycloudclient.NetworkRegion{
			Name:   v.Name.ValueString(),
			Region: v.Region.ValueString(),
		})
	}

	networkToCreate := dummycloudclient.Network{
		Name:         plan.Name.ValueString(),
		IsActive:     plan.IsActive.ValueBool(),
		IPList:       newIPList,
		InstanceList: newInstanceList,
	}

	// Create new order
	order, err := r.client.CreateNetwork(networkToCreate)
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

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *networkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state networkSchemaModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from DummyCloud
	order, err := r.client.GetNetworks(state.ID.ValueString())
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

	singleNetwork := order[0]

	state.ID = types.StringValue(singleNetwork.ID)
	state.Name = types.StringValue(singleNetwork.Name)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *networkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan networkSchemaModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newIPList := []string{}
	for _, v := range plan.IPList {
		newIPList = append(newIPList, v.ValueString())
	}

	newInstanceList := []dummycloudclient.NetworkRegion{}
	for _, v := range plan.InstanceList {
		newInstanceList = append(newInstanceList, dummycloudclient.NetworkRegion{
			Name:   v.Name.ValueString(),
			Region: v.Region.ValueString(),
		})
	}

	networkToUpdate := dummycloudclient.Network{
		ID:           plan.ID.ValueString(),
		Name:         plan.Name.ValueString(),
		IsActive:     plan.IsActive.ValueBool(),
		IPList:       newIPList,
		InstanceList: newInstanceList,
	}

	// Update existing order
	order, err := r.client.UpdateNetwork(networkToUpdate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating DummyCloud Order",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(order.ID)
	plan.Name = types.StringValue(order.Name)

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *networkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state networkSchemaModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	_, err := r.client.DeleteNetwork(state.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting DummyCloud Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *networkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
