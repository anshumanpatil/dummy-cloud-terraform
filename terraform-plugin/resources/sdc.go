package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"terraform-provider-pfxm/helper"
	schemastructures "terraform-provider-pfxm/schema-structures"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", tab)

	err := encoder.Encode(data)
	if err != nil {
		return "err"
	}
	return buffer.String()
}

func ResourceSdc() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSdcCreate,
		ReadContext:   resourceSdcRead,
		UpdateContext: resourceSdcUpdate,
		DeleteContext: resourceSdcDelete,
		Schema:        schemastructures.SDC_RESOURCE_SCHEMA,
	}
}

func resourceSdcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "anshu in resourceSdcCreate ")

	client := &http.Client{Timeout: 10 * time.Second}

	// tflog.Debug(ctx, "anshu HasChange "+strconv.FormatBool(d.HasChange("name")))
	if !d.HasChange("name") {
		return resourceSdcRead(ctx, d, m)
	}
	// d.Get("items").
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics

	// path := fmt.Sprintf("%s/sdc/readall", helper.ENV.Host)
	// if d.Get("id") != "" {
	path := fmt.Sprintf("%s/sdc/update/%s", helper.ENV.Host, d.Get("sdcid").(string))
	// }

	tflog.Debug(ctx, "anshu in resourceSdcCreate path "+path)
	// var jsonData = []byte(`{
	// 	"name": ` + d.Get("name").(string) + `
	// }`)
	type Student struct {
		Name string `json:"name"`
	}

	jbody := &Student{
		Name: d.Get("name").(string),
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(jbody)
	req, err := http.NewRequest("PUT", path, payloadBuf)
	// req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		tflog.Debug(ctx, "anshu in resourceSdcCreate err "+PrettyJson(err))
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// sdcs := make([]map[string]interface{}, 0)

	// body, _ := ioutil.ReadAll(r.Body)

	sdc := make(map[string]interface{})
	sdcs := make([]map[string]interface{}, 0)

	err = json.NewDecoder(r.Body).Decode(&sdc)
	sdcs = append(sdcs, sdc)
	tflog.Debug(ctx, "anshu resourceSdcCreate response "+PrettyJson(sdcs))

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Error,
	// 	Summary:  "sdcs got",
	// 	Detail:   "id is null all sdc want ?",
	// })

	// return diags
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("sdcs", sdcs); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(d.Get("name").(string))
	// d.Set("last_updated", time.Now().Format(time.RFC850))
	return resourceSdcRead(ctx, d, m)
}

func resourceSdcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	tflog.Debug(ctx, "anshu HasChange "+strconv.FormatBool(d.HasChange("name")))
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	path := fmt.Sprintf("%s/sdc/read/%s", helper.ENV.Host, d.Get("sdcid").(string))
	tflog.Debug(ctx, "anshu path "+path)
	tflog.Debug(ctx, "anshu sdcid "+d.Get("sdcid").(string))
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
	tflog.Debug(ctx, "anshu returnsdcs "+PrettyJson(returnsdcs))
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
	// d.SetId(d.Get("name").(string))

	return diags
}

func resourceSdcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return renameSDC(ctx, d, m)
}

func resourceSdcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	d.SetId("")
	// return resourceSdcRead(ctx, d, m)
	return diags
}

func renameSDC(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "anshu in resourceSdcCreate ")

	client := &http.Client{Timeout: 10 * time.Second}

	// tflog.Debug(ctx, "anshu HasChange "+strconv.FormatBool(d.HasChange("name")))
	if !d.HasChange("name") {
		return resourceSdcRead(ctx, d, m)
	}
	// d.Get("items").
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics

	// path := fmt.Sprintf("%s/sdc/readall", helper.ENV.Host)
	// if d.Get("id") != "" {
	path := fmt.Sprintf("%s/sdc/update/%s", helper.ENV.Host, d.Get("sdcid").(string))
	// }

	tflog.Debug(ctx, "anshu in resourceSdcCreate path "+path)
	// var jsonData = []byte(`{
	// 	"name": ` + d.Get("name").(string) + `
	// }`)
	type Student struct {
		Name string `json:"name"`
	}

	jbody := &Student{
		Name: d.Get("name").(string),
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(jbody)
	req, err := http.NewRequest("PUT", path, payloadBuf)
	// req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		tflog.Debug(ctx, "anshu in resourceSdcCreate err "+PrettyJson(err))
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// sdcs := make([]map[string]interface{}, 0)

	// body, _ := ioutil.ReadAll(r.Body)

	sdc := make(map[string]interface{})
	sdcs := make([]map[string]interface{}, 0)

	err = json.NewDecoder(r.Body).Decode(&sdc)
	sdcs = append(sdcs, sdc)
	tflog.Debug(ctx, "anshu resourceSdcCreate response "+PrettyJson(sdcs))

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Error,
	// 	Summary:  "sdcs got",
	// 	Detail:   "id is null all sdc want ?",
	// })

	// return diags
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("sdcs", sdcs); err != nil {
		return diag.FromErr(err)
	}

	// always run
	// d.SetId(d.Get("name").(string))
	// d.Set("last_updated", time.Now().Format(time.RFC850))
	return resourceSdcRead(ctx, d, m)
}
