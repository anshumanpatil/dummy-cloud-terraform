package pfxm

import (
	datasources "terraform-provider-pfxm/data-sources"
	resources "terraform-provider-pfxm/resources"
	schemastructures "terraform-provider-pfxm/schema-structures"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: schemastructures.AUTH_SCHEMA,
		ResourcesMap: map[string]*schema.Resource{
			"pfxm_sdcs": resources.ResourceSdc(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"pfxm_sdcs": datasources.DataSourceSdcs(),
		},
		ConfigureContextFunc: schemastructures.AuthConfigure,
	}
}
