package repository

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

// Whenever  terraform is required to READ the datasource ( data_source.go ) then this is called.
// which when completed should populate the structure.

func dataSourceReposRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var err error
	var req *http.Request

	// Get any parameters
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
		// If our URL does not return 200 as expected then lets play with formating a nice message
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

	repos := make([]map[string]interface{}, 0) // create a dummy MAP to hold JSON
	body, err := ioutil.ReadAll(r.Body)        // Read the request body into a string

	err = json.Unmarshal(body, &repos) //extract the JSON from the request Body string
	if err != nil {                    // The extract failed, good chance it didnt return an array so try again but in a flatter format
		temp := make(map[string]interface{}, 0)
		err = json.Unmarshal([]byte(body), &temp)
		if err != nil {
			return diag.FromErr(err)
		}
		err = d.Set("repositories", extract(temp)) // Populate the repositories element
	} else {
		err = d.Set("repositories", flatten(&repos)) // Populate the repositories element
	}

	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
