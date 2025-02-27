package fmc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fmc_access_policies",
			Category:         "Data Sources",
			ShortDescription: `Data source for Access Policies in FMC An example is shown below: hcl data "fmc_access_policies" "acp" { name = "FTD ACP" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_devices",
			Category:         "Data Sources",
			ShortDescription: `Data source for FTD Devices in FMC An example is shown below: hcl data "fmc_devices" "device" { name = "ftd.adyah.cisco" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_dynamic_objects",
			Category:         "Data Sources",
			ShortDescription: `Data source for Dynamic Objects in FMC An example is shown below: hcl data "fmc_dynamic_object" "dyobj" { name = "Dynamic Object 1" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_file_policies",
			Category:         "Data Sources",
			ShortDescription: `Data source for File Policies in FMC An example is shown below: hcl data "fmc_file_policies" "file_policy" { name = "AMP Policy" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_host_objects",
			Category:         "Data Sources",
			ShortDescription: `Data source for Host Objects in FMC An example is shown below: hcl data "fmc_host_objects" "existing_host" { name = "CUCM-Pub" } Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_ips_policies",
			Category:         "Data Sources",
			ShortDescription: `Data source for IPS Policy in FMC An example is shown below: hcl data "fmc_ips_policies" "ips_policy" { name = "Connectivity Over Security" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_network_objects",
			Category:         "Data Sources",
			ShortDescription: `Data source for Network Objects in FMC An example is shown below: hcl data "fmc_network_objects" "PrivateVLAN" { name = "VLAN825-Private" } Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_port_objects",
			Category:         "Data Sources",
			ShortDescription: `Data source for Port Objects in FMC An example is shown below: hcl data "fmc_port_objects" "existing" { name = "DNS_over_TCP" } Any one of the id, name or port can be specified. The first filter in the order of id, name and port will be used, and the rest will be ignored if multiple are specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_security_zones",
			Category:         "Data Sources",
			ShortDescription: `Data source for Security Zones in FMC An example is shown below: hcl data "fmc_security_zones" "inside" { name = "inside" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_syslog_alerts",
			Category:         "Data Sources",
			ShortDescription: `Data source for Syslog Alert Configuration in FMC An example is shown below: hcl data "fmc_security_zones" "inside" { name = "inside" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_url_objects",
			Category:         "Data Sources",
			ShortDescription: `Data source for URL Objects in FMC An example is shown below: hcl data "fmc_url_objects" "existing" { name = "DNAC" } Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"fmc_access_policies": 0,
		"fmc_devices":         1,
		"fmc_dynamic_objects": 2,
		"fmc_file_policies":   3,
		"fmc_host_objects":    4,
		"fmc_ips_policies":    5,
		"fmc_network_objects": 6,
		"fmc_port_objects":    7,
		"fmc_security_zones":  8,
		"fmc_syslog_alerts":   9,
		"fmc_url_objects":     10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
