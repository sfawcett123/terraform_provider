package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-fawcetts/github/repository"
)

// Provider -
// provide the datasource .. as fawcetts_repositories - This is in the repository subdirectory

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"fawcetts_repositories": repository.DataSourceRepos(),
		},
	}
}
