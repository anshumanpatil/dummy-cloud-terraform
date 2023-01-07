package schemastructures

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var SDC_DATA_RESOURCE_SCHEMA map[string]*schema.Schema = map[string]*schema.Schema{
	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"sdcs": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"sdcid": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	},
}

var SDC_RESOURCE_SCHEMA map[string]*schema.Schema = map[string]*schema.Schema{
	"sdcid": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		// Computed: true,
	},
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"sdcs": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	},
}
