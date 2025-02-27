package bridgecrew

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_apitokens",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform apitokens. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/apitokenlistbyuserid>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_apitokens_customer",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all of your Bridgecrew platform apitokens. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/apitokenlistbycustomername>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_authors",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform authors. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getgitblameauthors>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_code_reviews",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform code reviews. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getcodereviewdata>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_enforcement_accounts",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all the Bridgecrew platform accounts managed with enforcement rules. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getaccounts>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_enforcement_rule",
			Category:         "Data Sources",
			ShortDescription: `Get a Bridgecrew platform enforcement rule by AccountID. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getschemeforaccount>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_enforcement_rules",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform enforcement rules. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getschemeforcustomer>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_incidents",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform incidents. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getincidents>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_incidents_info",
			Category:         "Data Sources",
			ShortDescription: `Gets all the info and counters of the incidents and violations. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getinfo>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_incidents_preset",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform incidents presets and counters. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getpresets>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_integrations",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform integrations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_justifications",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform suppression justifications. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getsuppressionsjustifications>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_mappings",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew check mappings. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getcheckovmappings>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_organisation",
			Category:         "Data Sources",
			ShortDescription: `Get your Bridgecrew Organisation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_policies",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all of your custom policies in the Bridgecrew platform. Details on this API call are available here <https://docs.bridgecrew.io/reference/getcustompoliciestable>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_repositories",
			Category:         "Data Sources",
			ShortDescription: `Gets a list of all your repositories managed by the Bridgecrew platform <https://docs.bridgecrew.io/reference/getrepositories>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_repository_branches",
			Category:         "Data Sources",
			ShortDescription: `Supplies the details on the branches under analysis for a given repository <https://docs.bridgecrew.io/reference/getbranches>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_suppressions",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your policies suppressions in the Bridgecrew platform <https://docs.bridgecrew.io/reference/getsuppressions>`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_tag",
			Category:         "Data Sources",
			ShortDescription: `Get a tag rule. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/gettag>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_tags",
			Category:         "Data Sources",
			ShortDescription: `Get all the tag rules. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/gettag>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_users",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"bridgecrew_apitokens":            0,
		"bridgecrew_apitokens_customer":   1,
		"bridgecrew_authors":              2,
		"bridgecrew_code_reviews":         3,
		"bridgecrew_enforcement_accounts": 4,
		"bridgecrew_enforcement_rule":     5,
		"bridgecrew_enforcement_rules":    6,
		"bridgecrew_incidents":            7,
		"bridgecrew_incidents_info":       8,
		"bridgecrew_incidents_preset":     9,
		"bridgecrew_integrations":         10,
		"bridgecrew_justifications":       11,
		"bridgecrew_mappings":             12,
		"bridgecrew_organisation":         13,
		"bridgecrew_policies":             14,
		"bridgecrew_repositories":         15,
		"bridgecrew_repository_branches":  16,
		"bridgecrew_suppressions":         17,
		"bridgecrew_tag":                  18,
		"bridgecrew_tags":                 19,
		"bridgecrew_users":                20,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
