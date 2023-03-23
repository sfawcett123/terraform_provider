package fawcetts

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRepos() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceReposRead,
		Schema: map[string]*schema.Schema{
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repositories": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceReposRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var err error
	var req *http.Request

	owner := d.Get("owner").(string)
	name := d.Get("name").(string)

	if len(name) == 0 {
		req, err = http.NewRequest("GET", fmt.Sprintf("%s/%s/repos", "https://api.github.com/users", owner), nil)

	} else {
		req, err = http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", "https://api.github.com/repos", owner, name), nil)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	if r.StatusCode != 200 {
		summary := fmt.Sprintf("Unable to find repository %s/%s", owner, name)
		detail := fmt.Sprintf("Return code %d using URL %s", r.StatusCode, r.Request.URL)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  summary,
			Detail:   detail,
		})

		return diags
	}

	defer r.Body.Close()

	repos := make([]map[string]interface{}, 0)
	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &repos)
	if err != nil {
		temp := make(map[string]interface{}, 0)
		err = json.Unmarshal([]byte(body), &temp)
		if err != nil {
			return diag.FromErr(err)
		}
		err = d.Set("repositories", extract(temp))
	} else {
		err = d.Set("repositories", flatten(&repos))
	}

	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flatten(repos *[]map[string]interface{}) []interface{} {

	if repos != nil {
		ois := make([]interface{}, len(*repos), len(*repos))

		for i, m := range *repos {
			ois[i] = pick_bones(m)
		}
		return ois
	}

	return make([]interface{}, 0)
}

func extract(repos map[string]interface{}) interface{} {

	if repos != nil {
		ois := make([]interface{}, 1, 1)
		ois[0] = pick_bones(repos)
		return ois
	}

	return make([]interface{}, 0)
}

func pick_bones(data map[string]interface{}) map[string]interface{} {
	oi := make(map[string]interface{})
	oi["id"] = data["id"]
	oi["name"] = data["name"]

	return oi
}
