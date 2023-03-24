package github

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "terraform-provider-fawcetts/github/client"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string) // Get the token value from the input

	var diags diag.Diagnostics

	c, err := client.NewClient(&token) // Connect see client.client.go
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
