package netbox

import (
	"github.com/digitalocean/go-netbox/netbox/client"
	"github.com/digitalocean/go-netbox/netbox"
)

type Config struct {
	Endpoint string
	Token string
}
func (c *Config) Client() (*client.NetBox, error) {
	cli := netbox.NewNetboxWithAPIKey(c.Endpoint, c.Token)
	return cli, nil
}	