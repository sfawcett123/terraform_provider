package client

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
)

type Response http.Response // Define a local type, so we can add a new method.

// add the new method
func (resp *Response) HandleStatusCode(msg string) diag.Diagnostics {

	var diags diag.Diagnostics

	if resp.StatusCode != 200 {
		// If our URL does not return 200 as expected then lets play with formating a nice message
		summary := fmt.Sprintf("Unable to find repository %s", msg)

		detail := fmt.Sprintf("Return code %d using URL %s", resp.StatusCode, resp.Request.URL)

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  summary,
			Detail:   detail,
		})

		return diags
	}

	return nil

}
