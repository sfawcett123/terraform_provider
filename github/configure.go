package github

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "terraform-provider-fawcetts/github/client"
)

// Simply configure the Provider, in this case we will only set up a HTTP client
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string) // Get the token value from the input

	c := client.NewClient(&token) // Connect see client.client.go
	return c, c.Diags
}
