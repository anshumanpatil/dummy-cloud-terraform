package datasources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"terraform-provider-pfxm/helper"
	schemastructures "terraform-provider-pfxm/schema-structures"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSdcs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSdcsRead,
		Schema:      schemastructures.SDC_DATA_RESOURCE_SCHEMA,
	}
}

func dataSourceSdcsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	path := fmt.Sprintf("%s/sdc/readall", helper.ENV.Host)
	if d.Get("id") != "" {
		path = fmt.Sprintf("%s/sdc/read/%s", helper.ENV.Host, d.Get("sdcid"))
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// sdcs := make([]map[string]interface{}, 0)

	var returnsdcs []interface{} //= []interface{}{}
	err = json.NewDecoder(r.Body).Decode(&returnsdcs)

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Error,
	// 	Summary:  "sdcs got",
	// 	Detail:   "id is null all sdc want ?",
	// })

	// return diags
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("sdcs", returnsdcs); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
