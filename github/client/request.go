package client

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Return the overriden type, so we can call the new method later
func (c Client) Get(url string) (*Response, diag.Diagnostics) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, url), nil)
	if err != nil {
		if err != nil {
			return nil, diag.FromErr(err)
		}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return (*Response)(resp), nil // cast to the new type.
}
