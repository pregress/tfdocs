package transloadit

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "transloadit_template",
			Category:         "Resources",
			ShortDescription: `Manages Transloadit template`,
			Description: `

Manages Transloadit Templates. 
For additional details please refer to the [Transloadit documentation](https://transloadit.com/docs/)

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "transloadit_template_credential",
			Category:         "Resources",
			ShortDescription: `Manages Transloadit Template Credential`,
			Description: `

Manages Transloadit Templates Credential
For additional details please refer to the [Transloadit documentation](https://transloadit.com/docs/topics/template-credentials/).

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"transloadit_template":            0,
		"transloadit_template_credential": 1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
