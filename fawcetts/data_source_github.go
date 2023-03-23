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
	owner := d.Get("owner").(string)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/repos", "https://api.github.com/users", owner), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	repos := make([]map[string]interface{}, 0)
	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &repos)
	if err != nil {
		return diag.FromErr(err)
	}

	Items := flatten(&repos)

	if err := d.Set("repositories", Items); err != nil {
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
			oi := make(map[string]interface{})

			// Pick the bones out of the JSON
			oi["id"] = m["id"]
			oi["name"] = m["name"]

			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}
