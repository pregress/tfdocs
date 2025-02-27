package astra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "astra_access_list",
			Category:         "Resources",
			ShortDescription: `astra_access_list resource represents a database access list, used to limit the ip's / CIDR groups that have access to a database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_cdc",
			Category:         "Resources",
			ShortDescription: `astra_cdc enables cdc for an Astra Serverless table.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_database",
			Category:         "Resources",
			ShortDescription: `astra_database provides an Astra Serverless Database resource. You can create and delete databases. Note: Classic Tier databases are not supported by the Terraform provider. (see https://docs.datastax.com/en/astra/docs/index.html for more about Astra DB)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_keyspace",
			Category:         "Resources",
			ShortDescription: `astra_keyspace provides a keyspace resource. Keyspaces are groupings of tables for Cassandra. astra_keyspace resources are associated with a database id. You can have multiple keyspaces per DB in addition to the default keyspace provided in the astra_database resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_private_link",
			Category:         "Resources",
			ShortDescription: `astra_private_link provides a private link resource. Private Link is a private network endpoint that can be created to connect from your vpc to Astra without using a publicly routable IP address. astra_private_link resources are associated with a database id. Once the private_link resource is created in Astra it must be linked to an endpoint within your vpc, use astra_private_link_endpoint to do this.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_private_link_endpoint",
			Category:         "Resources",
			ShortDescription: `astra_private_link_endpoint completes the creation of a private link endpoint by associating it with your endpoint.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_role",
			Category:         "Resources",
			ShortDescription: `astra_role resource represents custom roles for a particular Astra Org. Custom roles can be assigned to an Astra user is to grant them granular permissions when the default roles in the UI are not specific enough. Roles are composed of policies which are granted to resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_streaming_sink",
			Category:         "Resources",
			ShortDescription: `astra_streaming_sink creates a streaming sink which sends data from a topic to a target system.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_streaming_tenant",
			Category:         "Resources",
			ShortDescription: `astra_streaming_tenant creates an Astra Streaming tenant.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_streaming_topic",
			Category:         "Resources",
			ShortDescription: `astra_streaming_topic creates an Astra Streaming topic.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_table",
			Category:         "Resources",
			ShortDescription: `astra_table provides a table resource which represents a table in cassandra.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_token",
			Category:         "Resources",
			ShortDescription: `astra_token resource represents a token with a specific role assigned.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"astra_access_list":           0,
		"astra_cdc":                   1,
		"astra_database":              2,
		"astra_keyspace":              3,
		"astra_private_link":          4,
		"astra_private_link_endpoint": 5,
		"astra_role":                  6,
		"astra_streaming_sink":        7,
		"astra_streaming_tenant":      8,
		"astra_streaming_topic":       9,
		"astra_table":                 10,
		"astra_token":                 11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
