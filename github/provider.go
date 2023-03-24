package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-fawcetts/github/repository"
)

// Provider -
// provide the datasource .. as fawcetts_repositories - This is in the repository subdirectory
// provide a schema  and add an entity ... tokem
// provide a provider configuration in file configure.go

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"fawcetts_repositories": repository.DataSourceRepos(),
		},
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("token", nil),
			},
		},
		ConfigureContextFunc: providerConfigure, // Call the provider config routine.
	}
}
