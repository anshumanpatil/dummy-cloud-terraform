package schemastructures

import (
	"context"
	"terraform-provider-pfxm/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var AUTH_SCHEMA map[string]*schema.Schema = map[string]*schema.Schema{
	"host": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		DefaultFunc: schema.EnvDefaultFunc("PFxM_HOST", nil),
	},
	"username": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		DefaultFunc: schema.EnvDefaultFunc("PFxM_USERNAME", nil),
	},
	"password": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		DefaultFunc: schema.EnvDefaultFunc("PFxM_PASSWORD", nil),
	},
}

func AuthConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	host := d.Get("host").(string)

	helper.ENV.Username = username
	helper.ENV.Password = password
	helper.ENV.Host = host

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Warning,
	// 	Summary:  "Unable to create Goscaleio client",
	// 	Detail:   "Unable to authenticate user for authenticated Goscaleio client",
	// })

	return nil, diags
}
