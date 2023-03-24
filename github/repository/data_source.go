package repository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Define the Data Source Structure for Repositories,
// INPUTS:
//    owner - Required String
//    name  - Optional String
// OUTPUTS:
//    id - Computed Int
//    name - Computed String
// READ CONTEXT:
//    dataSourceReposRead = In file read_datasource.go

func DataSourceRepos() *schema.Resource {
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
