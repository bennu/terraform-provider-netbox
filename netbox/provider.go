package netbox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)


func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:         providerSchema(),
		ResourcesMap:   providerResourceMap(),
		ConfigureFunc:  providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema {
		"api_endpoint": &schema.Schema {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("NETBOX_ENDPOINT", nil),
			Description: "API endpoint",
		},
		"api_token": &schema.Schema {
			Type:		 schema.TypeString,
			Required:	 true,
			DefaultFunc: schema.EnvDefaultFunc("NETBOX_TOKEN", nil),
			Description: "API TOKEN",
		},
	}
}

func providerResourceMap() map[string]*schema.Resource{
	return map[string]*schema.Resource{
		"ip_address": resourceIPAddress(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Endpoint: d.Get("api_endpoint").(string),
		Token:    d.Get("api_token").(string),
	}
	return config.Client()
}