package hcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcloud_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Certificate to represent a TLS certificate in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_firewall",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Firewall to represent a Firewall in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Floating IP to represent a publicly-accessible static IP address that can be mapped to one of your servers.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip_assignment",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Floating IP Assignment to assign a Floating IP to a Hetzner Cloud Server.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
				"assignment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Load Balancer to represent a Load Balancer in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer_network",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Load Balancer Network to represent a private network on a Load Balancer in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer_service",
			Category:         "Resources",
			ShortDescription: `Define services for Hetzner Cloud Load Balancers.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"service",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Network to represent a Network in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network_route",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Network Route to represent a Network route in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"route",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network_subnet",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Network Subnet to represent a Subnet in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"subnet",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_placement_group",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Placement Group to represent a Placement Group in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"placement",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_primary_ip",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Primary IP to represent a publicly-accessible static IP address that can be mapped to one of your servers.`,
			Description:      ``,
			Keywords: []string{
				"primary",
				"ip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_rdns",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Reverse DNS Entry to create, modify and reset reverse dns entries for Hetzner Cloud Servers, Primary IPs, Floating IPs or Load Balancers.`,
			Description:      ``,
			Keywords: []string{
				"rdns",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server",
			Category:         "Resources",
			ShortDescription: `Provides an Hetzner Cloud server resource. This can be used to create, modify, and delete servers. Servers also support provisioning.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server_network",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Server Network to represent a private network on a server in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud snapshot to represent an image with type snapshot in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud SSH key resource to manage SSH keys for server access.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud volume resource to manage volumes.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Volume attachment to attach a Volume to a Hetzner Cloud Server.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"attachment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"hcloud_certificate":            0,
		"hcloud_firewall":               1,
		"hcloud_floating_ip":            2,
		"hcloud_floating_ip_assignment": 3,
		"hcloud_load_balancer":          4,
		"hcloud_load_balancer_network":  5,
		"hcloud_load_balancer_service":  6,
		"hcloud_network":                7,
		"hcloud_network_route":          8,
		"hcloud_network_subnet":         9,
		"hcloud_placement_group":        10,
		"hcloud_primary_ip":             11,
		"hcloud_rdns":                   12,
		"hcloud_server":                 13,
		"hcloud_server_network":         14,
		"hcloud_snapshot":               15,
		"hcloud_ssh_key":                16,
		"hcloud_volume":                 17,
		"hcloud_volume_attachment":      18,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
