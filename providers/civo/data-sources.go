package civo

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "civo_database",
			Category:         "Data Sources",
			ShortDescription: `Get information of an Database for use in other resources. This data source provides all of the Database's properties as configured on your Civo account. Note: This data source returns a single Database. When specifying a name, an error will be raised if more than one Databases with the same name found.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_disk_image",
			Category:         "Data Sources",
			ShortDescription: `Get information on an disk image for use in other resources (e.g. creating a instance) with the ability to filter the results.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_name",
			Category:         "Data Sources",
			ShortDescription: `Get information on a domain. This data source provides the name and the id. An error will be raised if the provided domain name is not in your Civo account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_record",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DNS record. This data source provides the name, TTL, and zone file as configured on your Civo account. An error will be raised if the provided domain name or record are not in your Civo account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_firewall",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a firewall for use in other resources. This data source provides all of the firewall's properties as configured on your Civo account. Firewalls may be looked up by id or name, and you can optionally pass region if you want to make a lookup for an expecific firewall inside that region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information on an instance for use in other resources. This data source provides all of the instance's properties as configured on your Civo account. Note: This data source returns a single instance. When specifying a hostname, an error will be raised if more than one instances found.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instances",
			Category:         "Data Sources",
			ShortDescription: `Get information on instances for use in other resources, with the ability to filter and sort the results. If no filters are specified, all instances will be returned. Note: You can use the civo_instance data source to obtain metadata about a single instance if you already know the id, unique hostname, or unique tag to retrieve.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instances_size",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about the instance sizes that Civo supports, with the ability to filter the results.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides a Civo Kubernetes cluster data source. Note: This data source returns a single Kubernetes cluster. When specifying a name, an error will be raised if more than one Kubernetes cluster found.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_version",
			Category:         "Data Sources",
			ShortDescription: `Provides access to the available Civo Kubernetes versions, with the ability to filter the results.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_loadbalancer",
			Category:         "Data Sources",
			ShortDescription: `Get information on a load balancer for use in other resources. This data source provides all of the load balancers properties as configured on your Civo account. An error will be raised if the provided load balancer name does not exist in your Civo account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_network",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a network for use in other resources. This data source provides all of the network's properties as configured on your Civo account. Networks may be looked up by id or label, and you can optionally pass region if you want to make a lookup for an expecific network inside that region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_object_store",
			Category:         "Data Sources",
			ShortDescription: `Get information of an Object Store for use in other resources. This data source provides all of the Object Store's properties as configured on your Civo account. Note: This data source returns a single Object Store. When specifying a name, an error will be raised if more than one Object Stores with the same name found.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_object_store_credential",
			Category:         "Data Sources",
			ShortDescription: `Get information of an Object Store Credential for use in other resources. This data source provides all of the Object Store Credential's properties as configured on your Civo account. Note: This data source returns a single Object Store Credential. When specifying a name, an error will be raised if more than one Object Store Credentials with the same name found.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_region",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about the region that Civo supports, with the ability to filter the results.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_reserved_ip",
			Category:         "Data Sources",
			ShortDescription: `Get information on a reserved IP. This data source provides the region and Instance id as configured on your Civo account. This is useful if the reserved IP in question is not managed by Terraform or you need to find the instance the IP is attached to. An error will be raised if the provided domain name is not in your Civo account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_size",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about the sizes that Civo supports, with the ability to filter the results.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a SSH key. This data source provides the name, and fingerprint as configured on your Civo account. An error will be raised if the provided SSH key name does not exist in your Civo account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_volume",
			Category:         "Data Sources",
			ShortDescription: `Get information on a volume for use in other resources. This data source provides all of the volumes properties as configured on your Civo account. An error will be raised if the provided volume name does not exist in your Civo account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"civo_database":                0,
		"civo_disk_image":              1,
		"civo_dns_domain_name":         2,
		"civo_dns_domain_record":       3,
		"civo_firewall":                4,
		"civo_instance":                5,
		"civo_instances":               6,
		"civo_instances_size":          7,
		"civo_kubernetes_cluster":      8,
		"civo_kubernetes_version":      9,
		"civo_loadbalancer":            10,
		"civo_network":                 11,
		"civo_object_store":            12,
		"civo_object_store_credential": 13,
		"civo_region":                  14,
		"civo_reserved_ip":             15,
		"civo_size":                    16,
		"civo_ssh_key":                 17,
		"civo_volume":                  18,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
