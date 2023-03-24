package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	Client "terraform-provider-fawcetts/github/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Whenever  terraform is required to READ the datasource ( data_source.go ) then this is called.
// which when completed should populate the structure.

func dataSourceReposRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get the client information from the resource data
	client := m.(*Client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var resp *Client.Response

	// Get any parameters
	owner := d.Get("owner").(string)
	name := d.Get("name").(string)

	if len(name) == 0 {
		resp, diags = client.Get(fmt.Sprintf("users/%s/repos", owner))

	} else {
		resp, diags = client.Get(fmt.Sprintf("repos/%s/%s", owner, name))
	}

	if diags != nil {
		return diags
	}

	// Check for response codes that are not 200 and issue an error
	diags = resp.HandleStatusCode(fmt.Sprintf("%s:%s", owner, name))
	if diags != nil {
		return diags
	}

	defer resp.Body.Close()

	repos := make([]map[string]interface{}, 0) // create a dummy MAP to hold JSON
	body, err := ioutil.ReadAll(resp.Body)     // Read the request body into a string

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
