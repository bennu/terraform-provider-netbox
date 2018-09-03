package netbox

import (
	//"log"
	//"github.com/digitalocean/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIPAddress() *schema.Resource{

	return &schema.Resource{
		Create: resourceIPAddressCreate,
		Read:   resourceIPAddressRead,
		Update: resourceIPAddressUpdate,
		Delete: resourceIPAddressDelete,

		Schema: map[string]*schema.Schema{
			"pool": &schema.Schema{
				Type: 	  schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type: 	  schema.TypeString,
				Required: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func resourceIPAddressCreate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceIPAddressRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceIPAddressUpdate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceIPAddressDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}