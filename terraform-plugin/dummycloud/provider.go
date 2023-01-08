package dummycloud

import (
	"context"
	"os"

	"dummy-cloud/bucket"
	"dummy-cloud/dummycloudclient"
	"dummy-cloud/instance"
	"dummy-cloud/network"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &dummycloudProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &dummycloudProvider{}
}

// dummycloudProvider is the provider implementation.
type dummycloudProvider struct{}

// dummycloudProviderModel maps provider schema data to a Go type.
type dummycloudProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// Metadata returns the provider type name.
func (p *dummycloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dummycloud"
}

// Schema defines the provider-level schema for configuration data.
func (p *dummycloudProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with DummyCloud.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "URI for DummyCloud API. May also be provided via DUMMYCLOUD_HOST environment variable.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for DummyCloud API. May also be provided via DUMMYCLOUD_USERNAME environment variable.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for DummyCloud API. May also be provided via DUMMYCLOUD_PASSWORD environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares a DummyCloud API client for data sources and resources.
func (p *dummycloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring DummyCloud client")

	// Retrieve provider data from configuration
	var config dummycloudProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown DummyCloud API Host",
			"The provider cannot create the DummyCloud API client as there is an unknown configuration value for the DummyCloud API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DUMMYCLOUD_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown DummyCloud API Username",
			"The provider cannot create the DummyCloud API client as there is an unknown configuration value for the DummyCloud API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DUMMYCLOUD_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown DummyCloud API Password",
			"The provider cannot create the DummyCloud API client as there is an unknown configuration value for the DummyCloud API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DUMMYCLOUD_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("DUMMYCLOUD_HOST")
	username := os.Getenv("DUMMYCLOUD_USERNAME")
	password := os.Getenv("DUMMYCLOUD_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing DummyCloud API Host",
			"The provider cannot create the DummyCloud API client as there is a missing or empty value for the DummyCloud API host. "+
				"Set the host value in the configuration or use the DUMMYCLOUD_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing DummyCloud API Username",
			"The provider cannot create the DummyCloud API client as there is a missing or empty value for the DummyCloud API username. "+
				"Set the username value in the configuration or use the DUMMYCLOUD_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing DummyCloud API Password",
			"The provider cannot create the DummyCloud API client as there is a missing or empty value for the DummyCloud API password. "+
				"Set the password value in the configuration or use the DUMMYCLOUD_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "dummycloud_host", host)
	ctx = tflog.SetField(ctx, "dummycloud_username", username)
	ctx = tflog.SetField(ctx, "dummycloud_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "dummycloud_password")

	tflog.Debug(ctx, "Creating DummyCloud client")

	// Create a new DummyCloud client using the configuration values
	client, err := dummycloudclient.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create DummyCloud API Client",
			"An unexpected error occurred when creating the DummyCloud API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"DummyCloud Client Error: "+err.Error(),
		)
		return
	}

	// Make the DummyCloud client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured DummyCloud client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *dummycloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		instance.NewInstanceDataSource,
		bucket.NewBucketDataSource,
		network.NewNetworkDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *dummycloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		instance.NewInstanceResource,
		bucket.NewBucketResource,
	}
}
