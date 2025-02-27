package mso

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mso_label",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Label`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) name of the label. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) type of the label.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) type of the label.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_role",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Role`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the schema. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Name displayed associated to this Role.`,
				},
				resource.Attribute{
					Name:        "read_permissions",
					Description: `(Optional) Read permissions assigned to the role. Choices for read_permissions: "view-sites", "view-tenants", "view-schemas", "view-tenant-schemas", "view-users", "view-roles", "view-all-audit-records", "view-backup", "view-labels"`,
				},
				resource.Attribute{
					Name:        "write_permissions",
					Description: `(Optional) Write permissions assigned to the role. Choices for write_permissions: "manage-sites", "manage-tenants", "manage-labels", "manage-schemas", "manage-tenant-schemas", "manage-users", "manage-roles", "manage-audit-records", "manage-backup", "manage-labels"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for this role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Name displayed associated to this Role.`,
				},
				resource.Attribute{
					Name:        "read_permissions",
					Description: `(Optional) Read permissions assigned to the role. Choices for read_permissions: "view-sites", "view-tenants", "view-schemas", "view-tenant-schemas", "view-users", "view-roles", "view-all-audit-records", "view-backup", "view-labels"`,
				},
				resource.Attribute{
					Name:        "write_permissions",
					Description: `(Optional) Write permissions assigned to the role. Choices for write_permissions: "manage-sites", "manage-tenants", "manage-labels", "manage-schemas", "manage-tenant-schemas", "manage-users", "manage-roles", "manage-audit-records", "manage-backup", "manage-labels"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for this role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the schema. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) A block that represents the template associated with the schema. Type - Block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of template.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the template.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `tenant_id for the template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) A block that represents the template associated with the schema. Type - Block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of template.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the template.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `tenant_id for the template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the site to fetch.`,
				},
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where site is associated. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) The name of the template deployed to the site.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) Site id is set to the MSO site UUID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) The name of the template deployed to the site.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) Site id is set to the MSO site UUID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Site Application Network Profile(ANP) Data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Site Anp.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Site Anp. The name of the ANP should be present in the ANP list of the given ` + "`" + `schema_id` + "`" + ` and ` + "`" + `template_name` + "`" + ` ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Site Anp to be created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Site Anp to be created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Application Network Profiles Endpoint Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp Epg.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group to manage. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Anp Epg to be created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Anp Epg to be created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Application Network Profiles Endpoint Groups Domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Domain.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp Epg Domain.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group to manage.`,
				},
				resource.Attribute{
					Name:        "dn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "domain_dn",
					Description: `(Optional) The dn of domain. This is required when ` + "`" + `domain_name` + "`" + ` and ` + "`" + `domain_type` + "`" + ` are not specified.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) The domain profile name. This is required when ` + "`" + `domain_dn` + "`" + ` is not used. This attribute requires ` + "`" + `domain_type` + "`" + ` and ` + "`" + `vmm_domain_type` + "`" + ` (when it is applicable) to be set.`,
				},
				resource.Attribute{
					Name:        "domain_type",
					Description: `(Optional) The type of domain to associate. This is required when ` + "`" + `domain_dn` + "`" + ` is not used. Choices: [ vmmDomain, l3ExtDomain, l2ExtDomain, physicalDomain, fibreChannelDomain ]`,
				},
				resource.Attribute{
					Name:        "vmm_domain_type",
					Description: `(Optional) The vmm domain type. This is required when ` + "`" + `domain_type` + "`" + ` is vmmDomain and ` + "`" + `domain_dn` + "`" + ` is not used. Choices: [ VMware, Microsoft, Redhat ] ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Anp Epg Domain to be created.`,
				},
				resource.Attribute{
					Name:        "deploy_immediacy",
					Description: `(Optional) The deployment immediacy of the domain. choices: [ immediate, lazy ]`,
				},
				resource.Attribute{
					Name:        "resolution_immediacy",
					Description: `(Optional) Determines when the policies should be resolved and available. choices: [ immediate, lazy, pre-provision ]`,
				},
				resource.Attribute{
					Name:        "vlan_encap_mode",
					Description: `(Optional) Which VLAN enacap mode to use. This attribute can only be used with vmmDomain domain association. choices: [ static, dynamic ]`,
				},
				resource.Attribute{
					Name:        "allow_micro_segmentation",
					Description: `(Optional) Specifies microsegmentation is enabled or not. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Which switching mode to use with this domain association. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "switch_type",
					Description: `(Optional) Which switch type to use with this domain association. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan_type",
					Description: `(Optional) Virtual LAN type for microsegmentation. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan",
					Description: `(Optional) Virtual LAN for microsegmentation. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan_type",
					Description: `(Optional) Virtual LAN type for port encap. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan",
					Description: `(Optional) Virtual LAN for port encap. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy_name",
					Description: `(Optional) EPG enhanced lagpolicy name. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy_dn",
					Description: `(Optional) Distinguished name of EPG lagpolicy. This attribute can only be used with vmmDomain domain association.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Anp Epg Domain to be created.`,
				},
				resource.Attribute{
					Name:        "deploy_immediacy",
					Description: `(Optional) The deployment immediacy of the domain. choices: [ immediate, lazy ]`,
				},
				resource.Attribute{
					Name:        "resolution_immediacy",
					Description: `(Optional) Determines when the policies should be resolved and available. choices: [ immediate, lazy, pre-provision ]`,
				},
				resource.Attribute{
					Name:        "vlan_encap_mode",
					Description: `(Optional) Which VLAN enacap mode to use. This attribute can only be used with vmmDomain domain association. choices: [ static, dynamic ]`,
				},
				resource.Attribute{
					Name:        "allow_micro_segmentation",
					Description: `(Optional) Specifies microsegmentation is enabled or not. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Which switching mode to use with this domain association. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "switch_type",
					Description: `(Optional) Which switch type to use with this domain association. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan_type",
					Description: `(Optional) Virtual LAN type for microsegmentation. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan",
					Description: `(Optional) Virtual LAN for microsegmentation. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan_type",
					Description: `(Optional) Virtual LAN type for port encap. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan",
					Description: `(Optional) Virtual LAN for port encap. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy_name",
					Description: `(Optional) EPG enhanced lagpolicy name. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy_dn",
					Description: `(Optional) Distinguished name of EPG lagpolicy. This attribute can only be used with vmmDomain domain association.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Application Network Profiles Endpoint Groups Selector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Selector.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) site ID under which you want to deploy Anp Epg Selector.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template under above site id where Anp Epg Selector to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the selector. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) expressions of Selector.`,
				},
				resource.Attribute{
					Name:        "expressions.key",
					Description: `(Optional) expression key for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.operator",
					Description: `(Optional) expression operator for the selector. value should be from "equals", "notEquals", "in", "notIn", "keyExist", "keyNotExist".`,
				},
				resource.Attribute{
					Name:        "expressions.value",
					Description: `(Optional) expression value for the selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) expressions of Selector.`,
				},
				resource.Attribute{
					Name:        "expressions.key",
					Description: `(Optional) expression key for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.operator",
					Description: `(Optional) expression operator for the selector. value should be from "equals", "notEquals", "in", "notIn", "keyExist", "keyNotExist".`,
				},
				resource.Attribute{
					Name:        "expressions.value",
					Description: `(Optional) expression value for the selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_static_leaf",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Application Network Profiles Endpoint Groups StaticLeaf.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg StaticLeaf.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp Epg StaticLeaf.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group to manage.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path Given to the StaticLeaf. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Anp Epg StaticLeaf to be created.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan",
					Description: `(Optional) The VLAN id of the static leaf.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Anp Epg StaticLeaf to be created.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan",
					Description: `(Optional) The VLAN id of the static leaf.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_static_port",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site ANP EPG Static Port.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name under which Static Port is deployed.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) ANP name under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) EPG name under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "path_type",
					Description: `(Required) The type of the static port.`,
				},
				resource.Attribute{
					Name:        "pod",
					Description: `(Required) The pod of the static port.`,
				},
				resource.Attribute{
					Name:        "leaf",
					Description: `(Required) The leaf of the static port.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the static port.`,
				},
				resource.Attribute{
					Name:        "fex",
					Description: `(Optional) Fex-id to be used. This parameter will work only with the ` + "`" + `path_type` + "`" + ` as ` + "`" + `port` + "`" + `. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan",
					Description: `(Optional) The microsegmentation VLAN id of the static port.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode of the static port.`,
				},
				resource.Attribute{
					Name:        "deployment_immediacy",
					Description: `(Optional) The deployment immediacy of the static port.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) The port encap VLAN id of the static port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "micro_seg_vlan",
					Description: `(Optional) The microsegmentation VLAN id of the static port.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode of the static port.`,
				},
				resource.Attribute{
					Name:        "deployment_immediacy",
					Description: `(Optional) The deployment immediacy of the static port.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) The port encap VLAN id of the static port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site ANP EPG Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) ANP name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) EPG name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP of the Subnet. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of this subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of this subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_bd",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Site Bridge Domain(BD) Data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Site Bd.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Bd.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Site Bd. The name of the Bd should be present in the Bd list of the given ` + "`" + `schema_id` + "`" + ` and ` + "`" + `template_name` + "`" + ` ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Site Bd to be created.`,
				},
				resource.Attribute{
					Name:        "host_route",
					Description: `(Optional) Value to check whether host-based routing is enabled. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Site Bd to be created.`,
				},
				resource.Attribute{
					Name:        "host_route",
					Description: `(Optional) Value to check whether host-based routing is enabled. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_bd_l3out",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Bridge Domain L3out.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Bd L3out.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Bd L3out.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "l3out_name",
					Description: `(Required) Name of L3out to manage. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Bd L3out to be created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Bd L3out to be created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_bd_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Bd Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Bd name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP of the Subnet. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of this subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of this subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_external_epg_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site external Endpoint Groups Selector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy external Epg Selector.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) site ID under which you want to deploy external Epg Selector.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template under above site id where external Epg Selector to be created.`,
				},
				resource.Attribute{
					Name:        "external_epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the selector. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Ip address associated with the selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Ip address associated with the selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Site VRF Data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Site Vrf.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Site Vrf. The name of the VRF should be present in the VRF list of the given ` + "`" + `schema_id` + "`" + ` and ` + "`" + `template_name` + "`" + ` ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Site Vrf to be created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Site Vrf to be created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf_region",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Vrf Region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Name of Region to manage. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Vrf Region to be created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR to set into region`,
				},
				resource.Attribute{
					Name:        "cidr.cidr_ip",
					Description: `(Optional) IP address for CIDR.`,
				},
				resource.Attribute{
					Name:        "cidr.primary",
					Description: `(Optional) primary flag to set CIDR as primary. Only one CIDR can be set as primary.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet",
					Description: `(Optional) subnets to associate with CIDR.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.ip",
					Description: `(Optional) IP address for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.name",
					Description: `(Optional) Name for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.zone",
					Description: `(Optional) zone for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.usage",
					Description: `(Optional) usage information of particular subnet.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway",
					Description: `(Optional) VPN gateway flag.`,
				},
				resource.Attribute{
					Name:        "hub_network_enable",
					Description: `(Optional) Hub Network enable flag. To set hub network in region, this attribute should be true. this parameter is supported in MSO v3.0 or higher with Cloud APIC version 5.0 or higher.`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `(Optional) Hub Network to set into the region. this parameter is supported in MSO v3.0 or higher with Cloud APIC version 5.0 or higher.`,
				},
				resource.Attribute{
					Name:        "hub_network.name",
					Description: `(Optional) name of the hub network.`,
				},
				resource.Attribute{
					Name:        "hub_network.tenant_name",
					Description: `(Optional) Tenant name for the hub network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Vrf Region to be created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR to set into region`,
				},
				resource.Attribute{
					Name:        "cidr.cidr_ip",
					Description: `(Optional) IP address for CIDR.`,
				},
				resource.Attribute{
					Name:        "cidr.primary",
					Description: `(Optional) primary flag to set CIDR as primary. Only one CIDR can be set as primary.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet",
					Description: `(Optional) subnets to associate with CIDR.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.ip",
					Description: `(Optional) IP address for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.name",
					Description: `(Optional) Name for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.zone",
					Description: `(Optional) zone for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr.subnet.usage",
					Description: `(Optional) usage information of particular subnet.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway",
					Description: `(Optional) VPN gateway flag.`,
				},
				resource.Attribute{
					Name:        "hub_network_enable",
					Description: `(Optional) Hub Network enable flag. To set hub network in region, this attribute should be true. this parameter is supported in MSO v3.0 or higher with Cloud APIC version 5.0 or higher.`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `(Optional) Hub Network to set into the region. this parameter is supported in MSO v3.0 or higher with Cloud APIC version 5.0 or higher.`,
				},
				resource.Attribute{
					Name:        "hub_network.name",
					Description: `(Optional) name of the hub network.`,
				},
				resource.Attribute{
					Name:        "hub_network.tenant_name",
					Description: `(Optional) Tenant name for the hub network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf_region_cidr",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Vrf Region Cidr.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Name of Region to manage.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The name of the region CIDR to manage. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Vrf Region to be created.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) Whether this is the primary CIDR.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Vrf Region to be created.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) Whether this is the primary CIDR.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf_region_cidr_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Site Vrf Region Cidr Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Vrf Region Cidr Subnet.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf Region Cidr Subnet.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Name of Region to manage.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Required) The IP range of for the region CIDR where Vrf Region Cidr Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP subnet of this region CIDR. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Vrf Region Cidr Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The name of the zone for the region CIDR subnet.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage for the region CIDR Subnet. ## Note ## Multiple Subnets with same Ip are allowed, but the operations will take place on the first found Subnet with the given Ip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) Template where Vrf Region Cidr Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The name of the zone for the region CIDR subnet.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage for the region CIDR Subnet. ## Note ## Multiple Subnets with same Ip are allowed, but the operations will take place on the first found Subnet with the given Ip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the template to fetch.`,
				},
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where template is associated. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant id is set to the MSO template UUID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the template deployed to the site`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant id is set to the MSO template UUID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the template deployed to the site`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Anp`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where anp is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the anp to add.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) template associated with the anp.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name as displayed on the MSO web interface. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Application Network Profiles Endpoint Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Endpoint Group to manage. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Optional) Name of Bridge Domain to associate with.`,
				},
				resource.Attribute{
					Name:        "bd_schema_id",
					Description: `(Opional) The schemaID that defines the referenced BD.`,
				},
				resource.Attribute{
					Name:        "bd_template_name",
					Description: `(Optional) The template that defines the referenced BD.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) The schemaID that defines the referenced VRF.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) The template that defines the referenced VRF.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name as displayed on the MSO web interface.`,
				},
				resource.Attribute{
					Name:        "useg_epg",
					Description: `(Optional) Boolean flag to enable or disable whether this is a USEG EPG.`,
				},
				resource.Attribute{
					Name:        "intra_epg",
					Description: `(Optional) Whether intra EPG isolation is enforced. choices: [ enforced, unenforced ]`,
				},
				resource.Attribute{
					Name:        "intersite_multicast_source",
					Description: `(Optional) Whether intersite multicast source is enabled. Default to false.`,
				},
				resource.Attribute{
					Name:        "proxy_arp",
					Description: `(Optional) Whether to enable Proxy ARP or not. (For Forwarding control) Default to false.`,
				},
				resource.Attribute{
					Name:        "preferred_group",
					Description: `(Optional) Boolean flag to enable or disable whether this EPG is added to preferred group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Optional) Name of Bridge Domain to associate with.`,
				},
				resource.Attribute{
					Name:        "bd_schema_id",
					Description: `(Opional) The schemaID that defines the referenced BD.`,
				},
				resource.Attribute{
					Name:        "bd_template_name",
					Description: `(Optional) The template that defines the referenced BD.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) The schemaID that defines the referenced VRF.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) The template that defines the referenced VRF.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name as displayed on the MSO web interface.`,
				},
				resource.Attribute{
					Name:        "useg_epg",
					Description: `(Optional) Boolean flag to enable or disable whether this is a USEG EPG.`,
				},
				resource.Attribute{
					Name:        "intra_epg",
					Description: `(Optional) Whether intra EPG isolation is enforced. choices: [ enforced, unenforced ]`,
				},
				resource.Attribute{
					Name:        "intersite_multicast_source",
					Description: `(Optional) Whether intersite multicast source is enabled. Default to false.`,
				},
				resource.Attribute{
					Name:        "proxy_arp",
					Description: `(Optional) Whether to enable Proxy ARP or not. (For Forwarding control) Default to false.`,
				},
				resource.Attribute{
					Name:        "preferred_group",
					Description: `(Optional) Boolean flag to enable or disable whether this EPG is added to preferred group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg_contract",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Template Application Network Profile(ANP) Endpoint Group(EPG) Contract Data Source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID of Application Network Profile(ANP) EPG.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name of ANP EPG .`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of the Application Network Profile.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of the Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "relationship_type",
					Description: `(Optional) Relationship Type of the ANP EPG Contract on MSO UI.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) Name of the contract. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaID associated with the contract.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Template Name associated with the contract.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaID associated with the contract.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Template Name associated with the contract.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Application Network Profiles Endpoint Groups Selector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Subnet. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) expressions of Selector.`,
				},
				resource.Attribute{
					Name:        "expressions.key",
					Description: `(Optional) expression key for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.operator",
					Description: `(Optional) expression operator for the selector. value should be from "equals", "notEquals", "in", "notIn", "keyExist", "keyNotExist".`,
				},
				resource.Attribute{
					Name:        "expressions.value",
					Description: `(Optional) expression value for the selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) expressions of Selector.`,
				},
				resource.Attribute{
					Name:        "expressions.key",
					Description: `(Optional) expression key for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.operator",
					Description: `(Optional) expression operator for the selector. value should be from "equals", "notEquals", "in", "notIn", "keyExist", "keyNotExist".`,
				},
				resource.Attribute{
					Name:        "expressions.value",
					Description: `(Optional) expression value for the selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Application Network Profiles Endpoint Groups Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Subnet.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) Template where Anp Epg Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Ip Address of Subnet. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Scope of Subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether the subnet should be shared or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Scope of Subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether the subnet should be shared or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg_useg_attr",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Application Network Profiles Endpoint Groups Useg Attributes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to create Anp Epg Useg Attributes .`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg Useg Attributes to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Useg Attributes. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "useg_type",
					Description: `(Optional) Type of Useg Attribute.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) String which describes this Useg Attribute.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Comparison Operator used in the Useg Attribute.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Classifier Category. It's used with useg_type ` + "`" + `tag` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of Useg-Attribute.`,
				},
				resource.Attribute{
					Name:        "useg_subnet",
					Description: `(Optional) Whether the Useg Subnet is enabled or not. This field only works with the ` + "`" + `useg_type` + "`" + ` Ip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "useg_type",
					Description: `(Optional) Type of Useg Attribute.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) String which describes this Useg Attribute.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Comparison Operator used in the Useg Attribute.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Classifier Category. It's used with useg_type ` + "`" + `tag` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of Useg-Attribute.`,
				},
				resource.Attribute{
					Name:        "useg_subnet",
					Description: `(Optional) Whether the Useg Subnet is enabled or not. This field only works with the ` + "`" + `useg_type` + "`" + ` Ip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_bd",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Template Bridge Domain Data Source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Bridge Domain. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display Name of the Bridge Domain on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of VRF attached with Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF.`,
				},
				resource.Attribute{
					Name:        "layer2_unknown_unicast",
					Description: `(Optional) Type of layer 2 unknown unicast.`,
				},
				resource.Attribute{
					Name:        "intersite_bum_traffic",
					Description: `(Optional) Boolean Flag to enable or disable intersite bum traffic.`,
				},
				resource.Attribute{
					Name:        "optimize_wan_bandwidth",
					Description: `(Optional) Boolean flag to enable or disable the wan bandwidth optimization.`,
				},
				resource.Attribute{
					Name:        "layer2_stretch",
					Description: `(Optional) Boolean flag to enable or disable the layer-2 stretch.`,
				},
				resource.Attribute{
					Name:        "layer3_multicast",
					Description: `(Optional) Boolean flag to enable or disable layer 3 multicast traffic.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies",
					Description: `(Optional) Block to provide dhcp_policy configurations.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.name",
					Description: `(Optional) Dhcp_policy name. Required if you specify the dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.version",
					Description: `(Optional) Version of dhcp_policy. Required if you specify the dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.dhcp_option_policy_name",
					Description: `(Optional) Name of dhcp_option_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.dhcp_option_policy_version",
					Description: `(Optional) Version of dhcp_option_policy. ### Deprecation warning: do not use 'dhcp_policy' map below in combination with NDO releases 3.2 and higher, use above 'dhcp_policies' block instead.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy",
					Description: `(Optional) Map to provide dhcp_policy configurations.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.name",
					Description: `(Optional) dhcp_policy name.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.version",
					Description: `(Optional) Version of dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.dhcp_option_policy_name",
					Description: `(Optional) Name of dhcp_option_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.dhcp_option_policy_version",
					Description: `(Optional) Version of dhcp_option_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display Name of the Bridge Domain on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of VRF attached with Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF.`,
				},
				resource.Attribute{
					Name:        "layer2_unknown_unicast",
					Description: `(Optional) Type of layer 2 unknown unicast.`,
				},
				resource.Attribute{
					Name:        "intersite_bum_traffic",
					Description: `(Optional) Boolean Flag to enable or disable intersite bum traffic.`,
				},
				resource.Attribute{
					Name:        "optimize_wan_bandwidth",
					Description: `(Optional) Boolean flag to enable or disable the wan bandwidth optimization.`,
				},
				resource.Attribute{
					Name:        "layer2_stretch",
					Description: `(Optional) Boolean flag to enable or disable the layer-2 stretch.`,
				},
				resource.Attribute{
					Name:        "layer3_multicast",
					Description: `(Optional) Boolean flag to enable or disable layer 3 multicast traffic.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies",
					Description: `(Optional) Block to provide dhcp_policy configurations.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.name",
					Description: `(Optional) Dhcp_policy name. Required if you specify the dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.version",
					Description: `(Optional) Version of dhcp_policy. Required if you specify the dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.dhcp_option_policy_name",
					Description: `(Optional) Name of dhcp_option_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policies.dhcp_option_policy_version",
					Description: `(Optional) Version of dhcp_option_policy. ### Deprecation warning: do not use 'dhcp_policy' map below in combination with NDO releases 3.2 and higher, use above 'dhcp_policies' block instead.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy",
					Description: `(Optional) Map to provide dhcp_policy configurations.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.name",
					Description: `(Optional) dhcp_policy name.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.version",
					Description: `(Optional) Version of dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.dhcp_option_policy_name",
					Description: `(Optional) Name of dhcp_option_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.dhcp_option_policy_version",
					Description: `(Optional) Version of dhcp_option_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_bd_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Bridge Domain Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Bridge Domain to be created.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP range in CIDR notation. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the subnet.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the subnet.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_contract",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Contract.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Contract.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Contract to be created.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) The name of the contract to manage. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the contract on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Optional) The type of filters defined in this contract. Allowed values are ` + "`" + `bothWay` + "`" + ` and ` + "`" + `oneWay` + "`" + `. Default to ` + "`" + `bothWay` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the contract.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_schema_id",
					Description: `(Optional) The schemaId in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_template_name",
					Description: `(Optional) The template name in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_name",
					Description: `(Required) The filter to associate with this contract.`,
				},
				resource.Attribute{
					Name:        "filter_relationship",
					Description: `(Required if filter_relationships is not used) Map to provide Filter Relationships.`,
				},
				resource.Attribute{
					Name:        "filter_relationship.filter_schema_id",
					Description: `(Optional) The schemaId in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationship.filter_template_name",
					Description: `(Optional) The template name in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationship.filter_name",
					Description: `(Required) The filter to associate with this contract.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) A list of filter directives. Allowed values are ` + "`" + `log` + "`" + ` and ` + "`" + `none` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the contract on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Optional) The type of filters defined in this contract. Allowed values are ` + "`" + `bothWay` + "`" + ` and ` + "`" + `oneWay` + "`" + `. Default to ` + "`" + `bothWay` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the contract.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_schema_id",
					Description: `(Optional) The schemaId in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_template_name",
					Description: `(Optional) The template name in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_name",
					Description: `(Required) The filter to associate with this contract.`,
				},
				resource.Attribute{
					Name:        "filter_relationship",
					Description: `(Required if filter_relationships is not used) Map to provide Filter Relationships.`,
				},
				resource.Attribute{
					Name:        "filter_relationship.filter_schema_id",
					Description: `(Optional) The schemaId in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationship.filter_template_name",
					Description: `(Optional) The template name in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationship.filter_name",
					Description: `(Required) The filter to associate with this contract.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) A list of filter directives. Allowed values are ` + "`" + `log` + "`" + ` and ` + "`" + `none` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_contract_filter",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Contract Filter.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Contract Filter.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Contract Filter to be created.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) The name of the contract to manage. There should be an existing contract with this name.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Required) The type of filters defined in this contract. Allowed values are ` + "`" + `bothWay` + "`" + `, ` + "`" + `provider_to_consumer` + "`" + ` and ` + "`" + `consumer_to_provider` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_schema_id",
					Description: `(Required) The schemaId in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_template_name",
					Description: `(Required) The template name in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_name",
					Description: `(Required) The filter name to associate with this contract. Filter must exist with the given ` + "`" + `filter_name` + "`" + `, ` + "`" + `filter_schema_id` + "`" + ` and ` + "`" + `filter_template_name` + "`" + ` Force New set to ` + "`" + `true` + "`" + `. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) A list of filter directives. Allowed values are ` + "`" + `log` + "`" + ` and ` + "`" + `none` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) A list of filter directives. Allowed values are ` + "`" + `log` + "`" + ` and ` + "`" + `none` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_contract_service_graph",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Contract Service Graph.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Contract Service Graph.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Contract Service Graph.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Contract Service Graph to be created.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) The name of the contract to manage. There should be an existing contract with this name.`,
				},
				resource.Attribute{
					Name:        "service_graph_name",
					Description: `(Required) The name of service graph. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "node_relationship",
					Description: `(Optional) Service graph node relationship information.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_bd_name",
					Description: `(Optional) bd name for provider connector at template level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_bd_name",
					Description: `(Optional) bd name for consumer connector at template level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_cluster_interface",
					Description: `(Optional) cluster interface for provider connector to attach with node at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_cluster_interface",
					Description: `(Optional) cluster interface for consumer connector to attach with node at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_redirect_policy_tenant",
					Description: `(Optional) tenant for redirection policy for provider connector at site level. It is required to set redirection policy for provider connector.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_redirect_policy",
					Description: `(Optional) redirection policy for provider connector at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_redirect_policy_tenant",
					Description: `(Optional) tenant for redirection policy for consumer connector at site level. It is required to set redirection policy for consumer connector.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_redirect_policy",
					Description: `(Optional) redirection policy for consumer connector at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_subnet_ips",
					Description: `(Optional) subnet ips which will be associated with provider connector at site level. It should be in CIDR format.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_subnet_ips",
					Description: `(Optional) subnet ips which will be associated with consumer connector at site level. It should be in CIDR format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "node_relationship",
					Description: `(Optional) Service graph node relationship information.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_bd_name",
					Description: `(Optional) bd name for provider connector at template level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_bd_name",
					Description: `(Optional) bd name for consumer connector at template level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_cluster_interface",
					Description: `(Optional) cluster interface for provider connector to attach with node at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_cluster_interface",
					Description: `(Optional) cluster interface for consumer connector to attach with node at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_redirect_policy_tenant",
					Description: `(Optional) tenant for redirection policy for provider connector at site level. It is required to set redirection policy for provider connector.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_connector_redirect_policy",
					Description: `(Optional) redirection policy for provider connector at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_redirect_policy_tenant",
					Description: `(Optional) tenant for redirection policy for consumer connector at site level. It is required to set redirection policy for consumer connector.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_connector_redirect_policy",
					Description: `(Optional) redirection policy for consumer connector at site level.`,
				},
				resource.Attribute{
					Name:        "node_relationship.provider_subnet_ips",
					Description: `(Optional) subnet ips which will be associated with provider connector at site level. It should be in CIDR format.`,
				},
				resource.Attribute{
					Name:        "node_relationship.consumer_subnet_ips",
					Description: `(Optional) subnet ips which will be associated with consumer connector at site level. It should be in CIDR format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_external_epg",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Template External Endpoint Group Data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy External-epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where External-epg to be created.`,
				},
				resource.Attribute{
					Name:        "external_epg_name",
					Description: `(Required) Name of External-epg. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the External-epg on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) The VRF associated to this External-epg. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of External-epg will be used if not provided. Should use this parameter when VRF is in different schema than external-epg.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of External-epg will be used if not provided.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Optional) Name of anp to attach.`,
				},
				resource.Attribute{
					Name:        "anp_schema_id",
					Description: `(Optional) SchemaId of anp. ` + "`" + `schema_id` + "`" + ` will be used if not provided.`,
				},
				resource.Attribute{
					Name:        "anp_template_name",
					Description: `(Optional) Template name of anp. ` + "`" + `template_name` + "`" + ` will be used if not provided.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) List of ids of sites associated with the schema. Required when ` + "`" + `external_epg_type` + "`" + ` is "cloud".`,
				},
				resource.Attribute{
					Name:        "selector_name",
					Description: `(Optional) name of the selector for external epg. Required when ` + "`" + `external_epg_type` + "`" + ` is "cloud".`,
				},
				resource.Attribute{
					Name:        "selector_ip",
					Description: `(Optional) ip address for expression in selector. Required when ` + "`" + `external_epg_type` + "`" + ` is "cloud".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the External-epg on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) The VRF associated to this External-epg. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of External-epg will be used if not provided. Should use this parameter when VRF is in different schema than external-epg.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of External-epg will be used if not provided.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Optional) Name of anp to attach.`,
				},
				resource.Attribute{
					Name:        "anp_schema_id",
					Description: `(Optional) SchemaId of anp. ` + "`" + `schema_id` + "`" + ` will be used if not provided.`,
				},
				resource.Attribute{
					Name:        "anp_template_name",
					Description: `(Optional) Template name of anp. ` + "`" + `template_name` + "`" + ` will be used if not provided.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) List of ids of sites associated with the schema. Required when ` + "`" + `external_epg_type` + "`" + ` is "cloud".`,
				},
				resource.Attribute{
					Name:        "selector_name",
					Description: `(Optional) name of the selector for external epg. Required when ` + "`" + `external_epg_type` + "`" + ` is "cloud".`,
				},
				resource.Attribute{
					Name:        "selector_ip",
					Description: `(Optional) ip address for expression in selector. Required when ` + "`" + `external_epg_type` + "`" + ` is "cloud".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_external_epg_contract",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Template External Endpoint Group Contract Data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy External-epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where External-epg to be created.`,
				},
				resource.Attribute{
					Name:        "external_epg_name",
					Description: `(Required) Name of External-epg.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) Name of Contract. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "relationship_type",
					Description: `(Optional) RelationType of the Contract. Values that can be used is provider or consumer`,
				},
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaID of Contract. schema_id of External-epg will be used if not provided. Should use this parameter when Contract is in different schema than external-epg.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Template Name of Contract. template_name of External-epg will be used if not provided.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "relationship_type",
					Description: `(Optional) RelationType of the Contract. Values that can be used is provider or consumer`,
				},
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaID of Contract. schema_id of External-epg will be used if not provided. Should use this parameter when Contract is in different schema than external-epg.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Template Name of Contract. template_name of External-epg will be used if not provided.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_external_epg_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template External Endpoint Groups Selector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Subnet. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) expressions of Selector.`,
				},
				resource.Attribute{
					Name:        "expressions.key",
					Description: `(Optional) expression key for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.operator",
					Description: `(Optional) expression operator for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.value",
					Description: `(Optional) expression value for the selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) expressions of Selector.`,
				},
				resource.Attribute{
					Name:        "expressions.key",
					Description: `(Optional) expression key for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.operator",
					Description: `(Optional) expression operator for the selector.`,
				},
				resource.Attribute{
					Name:        "expressions.value",
					Description: `(Optional) expression value for the selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_external_epg_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template External EPG Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy External EPG Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where External EPG Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "external_epg_name",
					Description: `(Required) Name of External EPG.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP range in CIDR notation. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet. Allowed values are ` + "`" + `shared-rtctrl` + "`" + `, ` + "`" + `export-rtctrl` + "`" + `, ` + "`" + `shared-security` + "`" + `, ` + "`" + `import-rtctrl` + "`" + `, ` + "`" + `import-security` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) The aggregate of the subnet. Allowed values are ` + "`" + `shared-rtctrl` + "`" + `, ` + "`" + `export-rtctrl` + "`" + `, ` + "`" + `shared-security` + "`" + `, ` + "`" + `import-rtctrl` + "`" + `. Aggregate should be enabled only if shared-rtctrl is enabled in Scope.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet. Allowed values are ` + "`" + `shared-rtctrl` + "`" + `, ` + "`" + `export-rtctrl` + "`" + `, ` + "`" + `shared-security` + "`" + `, ` + "`" + `import-rtctrl` + "`" + `, ` + "`" + `import-security` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) The aggregate of the subnet. Allowed values are ` + "`" + `shared-rtctrl` + "`" + `, ` + "`" + `export-rtctrl` + "`" + `, ` + "`" + `shared-security` + "`" + `, ` + "`" + `import-rtctrl` + "`" + `. Aggregate should be enabled only if shared-rtctrl is enabled in Scope.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_filter_entry",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Template Filter Entry Data Source..`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID of Entry.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name of Entry.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Filter name of the Entry.`,
				},
				resource.Attribute{
					Name:        "entry_name",
					Description: `(Required) Name of the entry. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name of the filter as displayed on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "entry_display_name",
					Description: `(Optional) The name of the entry as displayed on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "entry_description",
					Description: `(Optional) The description of entry.`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `(Optional) The ethernet type to use for this filter entry.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol to use for this filter entry.`,
				},
				resource.Attribute{
					Name:        "tcp_session_rules",
					Description: `(Optional) A list of TCP session rules.`,
				},
				resource.Attribute{
					Name:        "source_from",
					Description: `(Optional) The source port range from.`,
				},
				resource.Attribute{
					Name:        "source_to",
					Description: `(Optional) The source port range to.`,
				},
				resource.Attribute{
					Name:        "destination_from",
					Description: `(Optional) The destination port range from.`,
				},
				resource.Attribute{
					Name:        "destination_to",
					Description: `(Optional) The destination port range to.`,
				},
				resource.Attribute{
					Name:        "arp_flag",
					Description: `(Optional) The ARP flag to use for this filter entry.`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Whether this filter entry is stateful.`,
				},
				resource.Attribute{
					Name:        "match_only_fragments",
					Description: `(Optional) Whether this filter entry only matches fragments.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name of the filter as displayed on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "entry_display_name",
					Description: `(Optional) The name of the entry as displayed on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "entry_description",
					Description: `(Optional) The description of entry.`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `(Optional) The ethernet type to use for this filter entry.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol to use for this filter entry.`,
				},
				resource.Attribute{
					Name:        "tcp_session_rules",
					Description: `(Optional) A list of TCP session rules.`,
				},
				resource.Attribute{
					Name:        "source_from",
					Description: `(Optional) The source port range from.`,
				},
				resource.Attribute{
					Name:        "source_to",
					Description: `(Optional) The source port range to.`,
				},
				resource.Attribute{
					Name:        "destination_from",
					Description: `(Optional) The destination port range from.`,
				},
				resource.Attribute{
					Name:        "destination_to",
					Description: `(Optional) The destination port range to.`,
				},
				resource.Attribute{
					Name:        "arp_flag",
					Description: `(Optional) The ARP flag to use for this filter entry.`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Whether this filter entry is stateful.`,
				},
				resource.Attribute{
					Name:        "match_only_fragments",
					Description: `(Optional) Whether this filter entry only matches fragments.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_l3out",
			Category:         "Data Sources",
			ShortDescription: `MSO Schema Template L3Out Data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy L3Out.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where L3Out to be created.`,
				},
				resource.Attribute{
					Name:        "l3out_name",
					Description: `(Required) Name of L3Out. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the L3Out on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) The VRF associated to this L3out. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of L3Out will be used if not provided. Should use this parameter when VRF is in different schema than l3out.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of L3Out will be used if not provided.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the L3Out on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) The VRF associated to this L3out. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of L3Out will be used if not provided. Should use this parameter when VRF is in different schema than l3out.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of L3Out will be used if not provided.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_service_graph",
			Category:         "Data Sources",
			ShortDescription: `Data Source for MSO Schema Template Service Graph`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) Schema ID where Service Graph is created.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template Name where Service Graph is created.`,
				},
				resource.Attribute{
					Name:        "service_graph_name",
					Description: `(Required) Name of Service Graph.`,
				},
				resource.Attribute{
					Name:        "node_index",
					Description: `(Required) Integer node index of service nodes. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "service_node_type",
					Description: `(Optional) Type of Service Node attached to this Graph at index provided by ` + "`" + `node_index` + "`" + `. Allowed values are ` + "`" + `firewall` + "`" + `, ` + "`" + `load-balancer` + "`" + ` , ` + "`" + `other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "site_nodes",
					Description: `(Optional) List of maps to Hold Site level Node association.`,
				},
				resource.Attribute{
					Name:        "site_nodes.site_id",
					Description: `(Optional) Site-Id Attached with the template. Where Service Graph is created.`,
				},
				resource.Attribute{
					Name:        "site_nodes.tenant_name",
					Description: `(Optional) Name of Tenant holding the Service Node.`,
				},
				resource.Attribute{
					Name:        "site_nodes.node_name",
					Description: `(Optional) Name of Site level Service Node/Device Name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_node_type",
					Description: `(Optional) Type of Service Node attached to this Graph at index provided by ` + "`" + `node_index` + "`" + `. Allowed values are ` + "`" + `firewall` + "`" + `, ` + "`" + `load-balancer` + "`" + ` , ` + "`" + `other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "site_nodes",
					Description: `(Optional) List of maps to Hold Site level Node association.`,
				},
				resource.Attribute{
					Name:        "site_nodes.site_id",
					Description: `(Optional) Site-Id Attached with the template. Where Service Graph is created.`,
				},
				resource.Attribute{
					Name:        "site_nodes.tenant_name",
					Description: `(Optional) Name of Tenant holding the Service Node.`,
				},
				resource.Attribute{
					Name:        "site_nodes.node_name",
					Description: `(Optional) Name of Site level Service Node/Device Name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_vrf",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Schema Template Vrf`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where vrf is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the vrf to add.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) template associated with the vrf. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name as displayed on the MSO web interface.`,
				},
				resource.Attribute{
					Name:        "layer3_multicast",
					Description: `(Optional) Whether to enable L3 multicast.`,
				},
				resource.Attribute{
					Name:        "vzany",
					Description: `(Optional) Whether to enable vzany.`,
				},
				resource.Attribute{
					Name:        "ip_data_plane_learning",
					Description: `(Optional) Whether IP data plane learning is enabled or disabled. Allowed values are ` + "`" + `disabled` + "`" + `and ` + "`" + `enabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preferred_group",
					Description: `(Optional) Whether to enable preferred Endpoint Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name as displayed on the MSO web interface.`,
				},
				resource.Attribute{
					Name:        "layer3_multicast",
					Description: `(Optional) Whether to enable L3 multicast.`,
				},
				resource.Attribute{
					Name:        "vzany",
					Description: `(Optional) Whether to enable vzany.`,
				},
				resource.Attribute{
					Name:        "ip_data_plane_learning",
					Description: `(Optional) Whether IP data plane learning is enabled or disabled. Allowed values are ` + "`" + `disabled` + "`" + `and ` + "`" + `enabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preferred_group",
					Description: `(Optional) Whether to enable preferred Endpoint Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_vrf_contract",
			Category:         "Data Sources",
			ShortDescription: `Data Source for MSO Schema Template Vrf Contract.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where vrf is associated.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) template associated with the vrf.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) name of the vrf with contract to be attached.`,
				},
				resource.Attribute{
					Name:        "relationship_type",
					Description: `(Required) Type of relation between VRF and Contract. Allowed values are ` + "`" + `provider` + "`" + ` and ` + "`" + `consumer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) Name of contract to be attached with the VRF. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaId of contract.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Name of template where contract is residing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaId of contract.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Name of template where contract is residing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_service_node_type",
			Category:         "Data Sources",
			ShortDescription: `Data Source for MSO Service Node Type`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Service Node Type. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of Service Node Type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of Service Node Type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_site",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Site`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the site. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username for the APICs.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the APICs.`,
				},
				resource.Attribute{
					Name:        "apic_site_id",
					Description: `(Optional) The site ID of the APICs.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Optional) A list of URLs to reference the APICs.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The labels for this site.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Location of the site.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username for the APICs.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the APICs.`,
				},
				resource.Attribute{
					Name:        "apic_site_id",
					Description: `(Optional) The site ID of the APICs.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Optional) A list of URLs to reference the APICs.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The labels for this site.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Location of the site.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_tenant",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO Tenant`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tenant.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name of the tenant to be displayed in the web UI. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for this tenant.`,
				},
				resource.Attribute{
					Name:        "user_associations",
					Description: `(Optional) A list of associated users for this tenant.`,
				},
				resource.Attribute{
					Name:        "site_association",
					Description: `(Optional) A list of associated sites for this tenant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for this tenant.`,
				},
				resource.Attribute{
					Name:        "user_associations",
					Description: `(Optional) A list of associated users for this tenant.`,
				},
				resource.Attribute{
					Name:        "site_association",
					Description: `(Optional) A list of associated sites for this tenant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_user",
			Category:         "Data Sources",
			ShortDescription: `Data source for MSO User`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) username of the user. It must contain at least 1 character in length. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `(Optional) password of the user. It must contain at least 8 characters in length.`,
				},
				resource.Attribute{
					Name:        "roles.roleid",
					Description: `(Optional) id of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "roles.access_type",
					Description: `(Optional) access_type of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "user_rbac",
					Description: `(Optional) roles given to the user.`,
				},
				resource.Attribute{
					Name:        "user_rbac.name",
					Description: `(Optional) name of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "user_rbac.user_priv",
					Description: `(Optional) Privilege access given to users (WritePriv, ReadPriv)`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) firstname of the user.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) lastname of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email of the user.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional) phone of the user.`,
				},
				resource.Attribute{
					Name:        "account-status",
					Description: `(Optional) account status of the user.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) domain status of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_password",
					Description: `(Optional) password of the user. It must contain at least 8 characters in length.`,
				},
				resource.Attribute{
					Name:        "roles.roleid",
					Description: `(Optional) id of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "roles.access_type",
					Description: `(Optional) access_type of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "user_rbac",
					Description: `(Optional) roles given to the user.`,
				},
				resource.Attribute{
					Name:        "user_rbac.name",
					Description: `(Optional) name of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "user_rbac.user_priv",
					Description: `(Optional) Privilege access given to users (WritePriv, ReadPriv)`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) firstname of the user.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) lastname of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email of the user.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional) phone of the user.`,
				},
				resource.Attribute{
					Name:        "account-status",
					Description: `(Optional) account status of the user.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) domain status of the user.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"mso_label":                                  0,
		"mso_role":                                   1,
		"mso_schema":                                 2,
		"mso_schema_site":                            3,
		"mso_schema_site_anp":                        4,
		"mso_schema_site_anp_epg":                    5,
		"mso_schema_site_anp_epg_domain":             6,
		"mso_schema_site_anp_epg_selector":           7,
		"mso_schema_site_anp_epg_static_leaf":        8,
		"mso_schema_site_anp_epg_static_port":        9,
		"mso_schema_site_anp_epg_subnet":             10,
		"mso_schema_site_bd":                         11,
		"mso_schema_site_bd_l3out":                   12,
		"mso_schema_site_bd_subnet":                  13,
		"mso_schema_site_external_epg_selector":      14,
		"mso_schema_site_vrf":                        15,
		"mso_schema_site_vrf_region":                 16,
		"mso_schema_site_vrf_region_cidr":            17,
		"mso_schema_site_vrf_region_cidr_subnet":     18,
		"mso_schema_template":                        19,
		"mso_schema_template_anp":                    20,
		"mso_schema_template_anp_epg":                21,
		"mso_schema_template_anp_epg_contract":       22,
		"mso_schema_template_anp_epg_selector":       23,
		"mso_schema_template_anp_epg_subnet":         24,
		"mso_schema_template_anp_epg_useg_attr":      25,
		"mso_schema_template_bd":                     26,
		"mso_schema_template_bd_subnet":              27,
		"mso_schema_template_contract":               28,
		"mso_schema_template_contract_filter":        29,
		"mso_schema_template_contract_service_graph": 30,
		"mso_schema_template_external_epg":           31,
		"mso_schema_template_external_epg_contract":  32,
		"mso_schema_template_external_epg_selector":  33,
		"mso_schema_template_external_epg_subnet":    34,
		"mso_schema_template_filter_entry":           35,
		"mso_schema_template_l3out":                  36,
		"mso_schema_template_service_graph":          37,
		"mso_schema_template_vrf":                    38,
		"mso_schema_template_vrf_contract":           39,
		"mso_service_node_type":                      40,
		"mso_site":                                   41,
		"mso_tenant":                                 42,
		"mso_user":                                   43,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
