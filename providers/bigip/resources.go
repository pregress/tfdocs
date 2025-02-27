package bigip

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_as3",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_as3 resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"as3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "as3_json",
					Description: `(Required) Path/Filename of Declarative AS3 JSON which is a json file used with builtin ` + "`" + `` + "`" + `` + "`" + `file` + "`" + `` + "`" + `` + "`" + ` function`,
				},
				resource.Attribute{
					Name:        "tenant_filter",
					Description: `(Optional) If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.`,
				},
				resource.Attribute{
					Name:        "ignore_metadata",
					Description: `(Optional) Set True if you want to ignore metadata changes during update. By default it is set to false`,
				},
				resource.Attribute{
					Name:        "as3_example1.json",
					Description: `Example AS3 Declarative JSON file with single tenant ` + "`" + `` + "`" + `` + "`" + `json { "class": "AS3", "action": "deploy", "persist": true, "declaration": { "class": "ADC", "schemaVersion": "3.0.0", "id": "example-declaration-01", "label": "Sample 1", "remark": "Simple HTTP application with round robin pool", "Sample_01": { "class": "Tenant", "defaultRouteDomain": 0, "Application_1": { "class": "Application", "template": "http", "serviceMain": { "class": "Service_HTTP", "virtualAddresses": [ "10.0.2.10" ], "pool": "web_pool" }, "web_pool": { "class": "Pool", "monitors": [ "http" ], "members": [ { "servicePort": 80, "serverAddresses": [ "192.0.1.100", "192.0.1.110" ] } ] } } } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "as3_example2.json",
					Description: `Example AS3 Declarative JSON file with multiple tenants ` + "`" + `` + "`" + `` + "`" + `json { "class": "AS3", "action": "deploy", "persist": true, "declaration": { "class": "ADC", "schemaVersion": "3.0.0", "id": "example-declaration-01", "label": "Sample 1", "remark": "Simple HTTP application with round robin pool", "Sample_02": { "class": "Tenant", "defaultRouteDomain": 0, "Application_2": { "class": "Application", "template": "http", "serviceMain": { "class": "Service_HTTP", "virtualAddresses": [ "10.2.2.10" ], "pool": "web_pool2" }, "web_pool2": { "class": "Pool", "monitors": [ "http" ], "members": [ { "servicePort": 80, "serverAddresses": [ "192.2.1.100", "192.2.1.110" ] } ] } } }, "Sample_03": { "class": "Tenant", "defaultRouteDomain": 0, "Application_3": { "class": "Application", "template": "http", "serviceMain": { "class": "Service_HTTP", "virtualAddresses": [ "10.1.2.10" ], "pool": "web_pool3" }, "web_pool3": { "class": "Pool", "monitors": [ "http" ], "members": [ { "servicePort": 80, "serverAddresses": [ "192.3.1.100", "192.3.1.110" ] } ] } } } } } ` + "`" + `` + "`" + `` + "`" + ` ## Import As3 resources can be imported using the partition name, e.g., ( use comma separated partition names if there are multiple partitions in as3 deployments ) ` + "`" + `` + "`" + `` + "`" + ` terraform import bigip_as3.test Sample_http_01 terraform import bigip_as3.test Sample_http_01,Sample_non_http_01 ` + "`" + `` + "`" + `` + "`" + ` #### Import examples ( single and multiple partitions ) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import bigip_as3.test Sample_http_01 bigip_as3.test: Importing from ID "Sample_http_01"... bigip_as3.test: Import prepared! Prepared bigip_as3 for import bigip_as3.test: Refreshing state... [id=Sample_http_01] Import successful! The resources that were imported are shown above. These resources are now in your Terraform state and will henceforth be managed by Terraform. $ terraform show # bigip_as3.test: resource "bigip_as3" "test" { as3_json = jsonencode( { action = "deploy" class = "AS3" declaration = { Sample_http_01 = { A1 = { class = "Application" jsessionid = { class = "Persist" cookieMethod = "hash" cookieName = "JSESSIONID" persistenceMethod = "cookie" } service = { class = "Service_HTTP" persistenceMethods = [ { use = "jsessionid" }, ] pool = "web_pool" virtualAddresses = [ "10.0.2.10", ] } web_pool = { class = "Pool" members = [ { serverAddresses = [ "192.0.2.10", "192.0.2.11", ] servicePort = 80 }, ] monitors = [ "http", ] } } class = "Tenant" } class = "ADC" id = "UDP_DNS_Sample" label = "UDP_DNS_Sample" remark = "Sample of a UDP DNS Load Balancer Service" schemaVersion = "3.0.0" } persist = true } ) id = "Sample_http_01" tenant_filter = "Sample_http_01" tenant_list = "Sample_http_01" } $ terraform import bigip_as3.test Sample_http_01,Sample_non_http_01 bigip_as3.test: Importing from ID "Sample_http_01,Sample_non_http_01"... bigip_as3.test: Import prepared! Prepared bigip_as3 for import bigip_as3.test: Refreshing state... [id=Sample_http_01,Sample_non_http_01] Import successful! The resources that were imported are shown above. These resources are now in your Terraform state and will henceforth be managed by Terraform. $ terraform show # bigip_as3.test: resource "bigip_as3" "test" { as3_json = jsonencode( { action = "deploy" class = "AS3" declaration = { Sample_http_01 = { A1 = { class = "Application" jsessionid = { class = "Persist" cookieMethod = "hash" cookieName = "JSESSIONID" persistenceMethod = "cookie" } service = { class = "Service_HTTP" persistenceMethods = [ { use = "jsessionid" }, ] pool = "web_pool" virtualAddresses = [ "10.0.2.10", ] } web_pool = { class = "Pool" members = [ { serverAddresses = [ "192.0.2.10", "192.0.2.11", ] servicePort = 80 }, ] monitors = [ "http", ] } } class = "Tenant" } Sample_non_http_01 = { DNS_Service = { Pool1 = { class = "Pool" members = [ { serverAddresses = [ "10.1.10.100", ] servicePort = 53 }, { serverAddresses = [ "10.1.10.101", ] servicePort = 53 }, ] monitors = [ "icmp", ] } class = "Application" service = { class = "Service_UDP" pool = "Pool1" virtualAddresses = [ "10.1.20.121", ] virtualPort = 53 } } class = "Tenant" } class = "ADC" id = "UDP_DNS_Sample" label = "UDP_DNS_Sample" remark = "Sample of a UDP DNS Load Balancer Service" schemaVersion = "3.0.0" } persist = true } ) id = "Sample_http_01,Sample_non_http_01" tenant_filter = "Sample_http_01,Sample_non_http_01" tenant_list = "Sample_http_01,Sample_non_http_01" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "AS3 documentation",
					Description: `https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/composing-a-declaration.html`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_bigiq_as3",
			Category:         "BIG-IQ",
			ShortDescription: `Provides details about bigip_bigiq_as3 resource`,
			Description:      ``,
			Keywords: []string{
				"big",
				"iq",
				"bigiq",
				"as3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bigiq_address",
					Description: `(Required, type ` + "`" + `string` + "`" + `) Address of the BIG-IQ to which your targer BIG-IP is attached`,
				},
				resource.Attribute{
					Name:        "bigiq_user",
					Description: `(Required, type ` + "`" + `string` + "`" + `) User name of the BIG-IQ to which your targer BIG-IP is attached`,
				},
				resource.Attribute{
					Name:        "bigiq_password",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Password of the BIG-IQ to which your targer BIG-IP is attached`,
				},
				resource.Attribute{
					Name:        "bigiq_port",
					Description: `(Optional) type ` + "`" + `int` + "`" + `, BIGIQ License Manager Port number, specify if port is other than ` + "`" + `443` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_token_auth",
					Description: `(Optional) type ` + "`" + `bool` + "`" + `, if set to ` + "`" + `true` + "`" + ` enables Token based Authentication,default is ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_login_ref",
					Description: `(Optional) BIGIQ Login reference for token authentication`,
				},
				resource.Attribute{
					Name:        "as3_json",
					Description: `(Required) Path/Filename of Declarative AS3 JSON which is a json file used with builtin ` + "`" + `` + "`" + `` + "`" + `file` + "`" + `` + "`" + `` + "`" + ` function`,
				},
				resource.Attribute{
					Name:        "ignore_metadata",
					Description: `(Optional) Set True if you want to ignore metadata changes during update. By default it is set to ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_example.json",
					Description: `Example AS3 Declarative JSON file ` + "`" + `` + "`" + `` + "`" + `json { "class": "AS3", "action": "deploy", "persist": true, "declaration": { "class": "ADC", "schemaVersion": "3.7.0", "id": "example-declaration-01", "label": "Task1", "remark": "Task 1 - HTTP Application Service", "target": { "address": "xx.xxx.xx.xxx" }, "Task1": { "class": "Tenant", "MyWebApp1http": { "class": "Application", "template": "http", "serviceMain": { "class": "Service_HTTP", "virtualAddresses": [ "10.1.2.10" ], "pool": "web_pool" }, "web_pool": { "class": "Pool", "monitors": [ "http" ], "members": [ { "servicePort": 80, "serverAddresses": [ "192.0.2.33", "192.0.2.13" ], "shareNodes": true } ] } } } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "AS3 documentation",
					Description: `https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_cm_device",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_cm_devicegroup",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip device_name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bigip_cm_devicegroup",
					Description: `Is the resource used to configure new device group on the BIG-IP.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Is the name of the device Group`,
				},
				resource.Attribute{
					Name:        "auto_sync",
					Description: `Specifies if the device-group will automatically sync configuration data to its members`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies if the device-group will be used for failover or resource syncing`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Name of the device to be included in device group, this need to be configured before using devicegroup resource`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_command",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "commands",
					Description: `(Required) The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to ` + "`" + `command_result` + "`" + ``,
				},
				resource.Attribute{
					Name:        "when",
					Description: `(Optional,possible values:` + "`" + `apply` + "`" + ` or ` + "`" + `destroy` + "`" + `) default value will be ` + "`" + `apply` + "`" + `,can be set to ` + "`" + `destroy` + "`" + ` for terraform destroy call. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "command_result",
					Description: `The resulting output from the ` + "`" + `commands` + "`" + ` executed`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "command_result",
					Description: `The resulting output from the ` + "`" + `commands` + "`" + ` executed`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_common_license_manage_bigiq",
			Category:         "BIG-IQ",
			ShortDescription: `Provides details about bigip_common_license_manage_bigiq resource`,
			Description:      ``,
			Keywords: []string{
				"big",
				"iq",
				"common",
				"license",
				"manage",
				"bigiq",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bigiq_address",
					Description: `(Required) BIGIQ License Manager IP Address, variable type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_user",
					Description: `(Required) BIGIQ License Manager username, variable type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_password",
					Description: `(Required) BIGIQ License Manager password. variable type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_port",
					Description: `(Optional) type ` + "`" + `int` + "`" + `, BIGIQ License Manager Port number, specify if port is other than ` + "`" + `443` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_token_auth",
					Description: `(Optional) type ` + "`" + `bool` + "`" + `, if set to ` + "`" + `true` + "`" + ` enables Token based Authentication,default is ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bigiq_login_ref",
					Description: `(Optional) BIGIQ Login reference for token authentication`,
				},
				resource.Attribute{
					Name:        "assignment_type",
					Description: `(Required) The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.`,
				},
				resource.Attribute{
					Name:        "license_poolname",
					Description: `(Required) A name given to the license pool. type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "unit_of_measure",
					Description: `(Optional, Required for ` + "`" + `Utility` + "`" + ` licenseType) The units used to measure billing. For example, “hourly” or “daily”. Type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "skukeyword1",
					Description: `(Optional) An optional offering name. type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "skukeyword2",
					Description: `(Optional) An optional offering name. type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional, Required Only for ` + "`" + `unreachable BIG-IP` + "`" + `) MAC address of the BIG-IP. type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `(Optional,Required Only for ` + "`" + `unreachable BIG-IP` + "`" + `) Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type ` + "`" + `string` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Optional) For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) License Assignment is done with specified ` + "`" + `key` + "`" + `, supported only with RegKeypool type License assignement. type ` + "`" + `string` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_do",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_do resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"do",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "do_json",
					Description: `(Required) Name of the of the Declarative DO JSON file`,
				},
				resource.Attribute{
					Name:        "bigip_address",
					Description: `(optional) IP Address of BIGIP Host to be used for this resource,this is optional parameter. whenever we specify this parameter it gets overwrite provider configuration`,
				},
				resource.Attribute{
					Name:        "bigip_user",
					Description: `(optional) UserName of BIGIP host to be used for this resource,this is optional parameter. whenever we specify this parameter it gets overwrite provider configuration`,
				},
				resource.Attribute{
					Name:        "bigip_port",
					Description: `(optional) Port number of BIGIP host to be used for this resource,this is optional parameter. whenever we specify this parameter it gets overwrite provider configuration`,
				},
				resource.Attribute{
					Name:        "bigip_password",
					Description: `(optional) Password of BIGIP host to be used for this resource,this is optional parameter. whenever we specify this parameter it gets overwrite provider configuration`,
				},
				resource.Attribute{
					Name:        "timeout(minutes)",
					Description: `(optional) timeout to keep polling DO endpoint until Bigip is provisioned by DO.( Default timeout is 20 minutes ) ~>`,
				},
				resource.Attribute{
					Name:        "example.json",
					Description: `Example of DO Declarative JSON ` + "`" + `` + "`" + `` + "`" + `json { "schemaVersion": "1.0.0", "class": "Device", "async": true, "label": "my BIG-IP declaration for declarative onboarding", "Common": { "class": "Tenant", "hostname": "bigip.example.com", "myLicense": { "class": "License", "licenseType": "regKey", "regKey": "xxxx" }, "admin": { "class": "User", "userType": "regular", "password": "xxxx", "shell": "bash" }, "myProvisioning": { "class": "Provision", "ltm": "nominal", "gtm": "minimum" }, "external": { "class": "VLAN", "tag": 4093, "mtu": 1500, "interfaces": [ { "name": "1.1", "tagged": true } ], "cmpHash": "dst-ip" }, "external-self": { "class": "SelfIp", "address": "x.x.x.x", "vlan": "external", "allowService": "default", "trafficGroup": "traffic-group-local-only" } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "DO documentation",
					Description: `https://clouddocs.f5.com/products/extensions/f5-declarative-onboarding/latest/composing-a-declaration.html#sample-declaration-for-a-standalone-big-ip`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_event_service_discovery",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_event_service_discovery resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"event",
				"service",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "taskid",
					Description: `(Required) servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Required) Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port) For more information, please refer below document https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0. With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes. When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes. For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool First we show the initial declaration to POST to the BIG-IP system. { "class": "ADC", "schemaVersion": "3.9.0", "id": "Pool", "Sample_event_sd": { "class": "Tenant", "My_app": { "class": "Application", "My_pool": { "class": "Pool", "members": [ { "servicePort": 8080, "addressDiscovery": "static", "serverAddresses": [ "192.0.2.2" ] }, { "servicePort": 8080, "addressDiscovery": "event" } ] } } } } Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_application",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_fast_application resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fast_json",
					Description: `(Required) Path/Filename of Declarative FAST JSON which is a json file used with builtin ` + "`" + `` + "`" + `` + "`" + `file` + "`" + `` + "`" + `` + "`" + ` function`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.`,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Optional) A FAST tenant name on which you want to manage application.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Optional) A FAST application name.`,
				},
				resource.Attribute{
					Name:        "FAST documentation",
					Description: `https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_http_app",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_fast_http_app resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"http",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required, ` + "`" + `string` + "`" + `) Name of the FAST HTTPS application tenant.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required ,` + "`" + `string` + "`" + `) Name of the FAST HTTPS application.`,
				},
				resource.Attribute{
					Name:        "virtual_server",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `virtual_server` + "`" + ` block will provide ` + "`" + `ip` + "`" + ` and ` + "`" + `port` + "`" + ` options to be used for virtual server. See [virtual server](#virtual-server) below for more details.`,
				},
				resource.Attribute{
					Name:        "existing_snat_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP SNAT pool.`,
				},
				resource.Attribute{
					Name:        "snat_pool_address",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of address to be used for FAST-Generated SNAT Pool.`,
				},
				resource.Attribute{
					Name:        "exist_pool_name",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP pool.`,
				},
				resource.Attribute{
					Name:        "pool_members",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `pool_members` + "`" + ` block takes input for FAST-Generated Pool. See [Pool Members](#pool-members) below for more details.`,
				},
				resource.Attribute{
					Name:        "service_discovery",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of different cloud service discovery config provided as string, provided ` + "`" + `service_discovery` + "`" + ` block to Automatically Discover Pool Members with Service Discovery on different clouds.`,
				},
				resource.Attribute{
					Name:        "load_balancing_mode",
					Description: `(Optional,` + "`" + `string` + "`" + `) A ` + "`" + `load balancing method` + "`" + ` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method`,
				},
				resource.Attribute{
					Name:        "slow_ramp_time",
					Description: `(Optional,` + "`" + `int` + "`" + `) Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds`,
				},
				resource.Attribute{
					Name:        "existing_monitor",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `monitor` + "`" + ` block takes input for FAST-Generated Pool Monitor. See [Pool Monitor](#pool-monitor) below for more details.`,
				},
				resource.Attribute{
					Name:        "existing_waf_security_policy",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing WAF Security policy.`,
				},
				resource.Attribute{
					Name:        "endpoint_ltm_policy",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of LTM Policies to be applied FAST HTTP Application.`,
				},
				resource.Attribute{
					Name:        "waf_security_policy",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `waf_security_policy` + "`" + ` block takes input for FAST-Generated WAF Security Policy. See [WAF Security Policy](#waf-security-policy) below for more details.`,
				},
				resource.Attribute{
					Name:        "security_log_profiles",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of security log profiles to be used for FAST application ### virtual server This IP address, combined with the port you specify below, becomes the BIG-IP virtual server address and port, which clients use to access the application The ` + "`" + `virtual_server` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional , ` + "`" + `string` + "`" + `) IP4/IPv6 address to be used for virtual server ex: ` + "`" + `10.1.1.1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Optional , ` + "`" + `list` + "`" + `) List of server address to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional , ` + "`" + `int` + "`" + `) port number of serviceport to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional , ` + "`" + `int` + "`" + `) connectionLimit value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "priority_group",
					Description: `(Optional , ` + "`" + `int` + "`" + `) priorityGroup value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "share_nodes",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) shareNodes value to be used for FAST-Generated Pool. ### Pool Monitor Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated Pool Monitor. The ` + "`" + `monitor` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "monitor_auth",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) set ` + "`" + `true` + "`" + ` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional , ` + "`" + `string` + "`" + `) username for web access on FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional , ` + "`" + `string` + "`" + `) password for web access on FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional , ` + "`" + `int` + "`" + `) Set the time between health checks,in seconds for FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "send_string",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Specify data to be sent during each health check for FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional , ` + "`" + `string` + "`" + `) The presence of this string anywhere in the HTTP response implies availability. ### WAF Security policy Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated WAF Security Policy The ` + "`" + `waf_security_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) Setting ` + "`" + `true` + "`" + ` will enable FAST to create WAF Security Policy.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_https_app",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_fast_https_app resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"https",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required, ` + "`" + `string` + "`" + `) Name of the FAST HTTPS application tenant.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required ,` + "`" + `string` + "`" + `) Name of the FAST HTTPS application.`,
				},
				resource.Attribute{
					Name:        "virtual_server",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `virtual_server` + "`" + ` block will provide ` + "`" + `ip` + "`" + ` and ` + "`" + `port` + "`" + ` options to be used for virtual server. See [virtual server](#virtual-server) below for more details.`,
				},
				resource.Attribute{
					Name:        "existing_snat_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP SNAT pool.`,
				},
				resource.Attribute{
					Name:        "snat_pool_address",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of address to be used for FAST-Generated SNAT Pool.`,
				},
				resource.Attribute{
					Name:        "existing_tls_server_profile",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing TLS server profile.`,
				},
				resource.Attribute{
					Name:        "tls_server_profile",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `tls_server_profile` + "`" + ` block takes input for FAST-Generated TLS Server Profile. See [TLS Server Profile](#tls-server-profile) below for more details. ~>`,
				},
				resource.Attribute{
					Name:        "existing_tls_client_profile",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing TLS client profile.`,
				},
				resource.Attribute{
					Name:        "tls_client_profile",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `tls_client_profile` + "`" + ` block takes input for FAST-Generated TLS client Profile. See [TLS Client Profile](#tls-client-profile) below for more details. ~>`,
				},
				resource.Attribute{
					Name:        "endpoint_ltm_policy",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of LTM Policies to be applied FAST HTTPS Application.`,
				},
				resource.Attribute{
					Name:        "existing_waf_security_policy",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing WAF Security policy.`,
				},
				resource.Attribute{
					Name:        "waf_security_policy",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `waf_security_policy` + "`" + ` block takes input for FAST-Generated WAF Security Policy. See [WAF Security Policy](#waf-security-policy) below for more details.`,
				},
				resource.Attribute{
					Name:        "existing_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP pool.`,
				},
				resource.Attribute{
					Name:        "pool_members",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `pool_members` + "`" + ` block takes input for FAST-Generated Pool. See [Pool Members](#pool-members) below for more details.`,
				},
				resource.Attribute{
					Name:        "service_discovery",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of different cloud service discovery config provided as string, provided ` + "`" + `service_discovery` + "`" + ` block to Automatically Discover Pool Members with Service Discovery on different clouds.`,
				},
				resource.Attribute{
					Name:        "load_balancing_mode",
					Description: `(Optional,` + "`" + `string` + "`" + `) A ` + "`" + `load balancing method` + "`" + ` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method`,
				},
				resource.Attribute{
					Name:        "slow_ramp_time",
					Description: `(Optional,` + "`" + `int` + "`" + `) Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds`,
				},
				resource.Attribute{
					Name:        "existing_monitor",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `monitor` + "`" + ` block takes input for FAST-Generated Pool Monitor. See [Pool Monitor](#pool-monitor) below for more details.`,
				},
				resource.Attribute{
					Name:        "security_log_profiles",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of security log profiles to be used for FAST application ### virtual server This IP address, combined with the port you specify below, becomes the BIG-IP virtual server address and port, which clients use to access the application The ` + "`" + `virtual_server` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional , ` + "`" + `string` + "`" + `) IP4/IPv6 address to be used for virtual server ex: ` + "`" + `10.1.1.1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tls_cert_name",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.`,
				},
				resource.Attribute{
					Name:        "tls_key_name",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile. ### TLS Client Profile Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated TLS Client Profile. The ` + "`" + `tls_client_profile` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "tls_cert_name",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.`,
				},
				resource.Attribute{
					Name:        "tls_key_name",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile. ### WAF Security policy Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated WAF Security Policy The ` + "`" + `waf_security_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) Setting ` + "`" + `true` + "`" + ` will enable FAST to create WAF Security Policy. ### Pool Members Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated Pool. The ` + "`" + `pool_members` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Optional , ` + "`" + `list` + "`" + `) List of server address to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional , ` + "`" + `int` + "`" + `) port number of serviceport to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional , ` + "`" + `int` + "`" + `) connectionLimit value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "priority_group",
					Description: `(Optional , ` + "`" + `int` + "`" + `) priorityGroup value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "share_nodes",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) shareNodes value to be used for FAST-Generated Pool. ### Pool Monitor Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated Pool Monitor. The ` + "`" + `monitor` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "monitor_auth",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) set ` + "`" + `true` + "`" + ` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional , ` + "`" + `string` + "`" + `) username for web access on FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional , ` + "`" + `string` + "`" + `) password for web access on FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional , ` + "`" + `int` + "`" + `) Set the time between health checks,in seconds for FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "send_string",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Specify data to be sent during each health check for FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional , ` + "`" + `string` + "`" + `) The presence of this string anywhere in the HTTP response implies availability.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_tcp_app",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_fast_tcp_app resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"tcp",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Name of the FAST TCP application.`,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required) Name of the FAST TCP application tenant.`,
				},
				resource.Attribute{
					Name:        "virtual_server",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `virtual_server` + "`" + ` block will provide ` + "`" + `ip` + "`" + ` and ` + "`" + `port` + "`" + ` options to be used for virtual server. See [virtual server](#virtual-server) below for more details.`,
				},
				resource.Attribute{
					Name:        "existing_snat_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP SNAT pool.`,
				},
				resource.Attribute{
					Name:        "snat_pool_address",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of address to be used for FAST-Generated SNAT Pool.`,
				},
				resource.Attribute{
					Name:        "existing_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP pool.`,
				},
				resource.Attribute{
					Name:        "pool_members",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `pool_members` + "`" + ` block takes input for FAST-Generated Pool. See [Pool Members](#pool-members) below for more details.`,
				},
				resource.Attribute{
					Name:        "load_balancing_mode",
					Description: `(Optional,` + "`" + `string` + "`" + `) A ` + "`" + `load balancing method` + "`" + ` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method`,
				},
				resource.Attribute{
					Name:        "slow_ramp_time",
					Description: `(Optional,` + "`" + `int` + "`" + `) Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds`,
				},
				resource.Attribute{
					Name:        "existing_monitor",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `monitor` + "`" + ` block takes input for FAST-Generated Pool Monitor. See [Pool Monitor](#pool-monitor) below for more details. ### virtual server This IP address, combined with the port you specify below, becomes the BIG-IP virtual server address and port, which clients use to access the application The ` + "`" + `virtual_server` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional , ` + "`" + `string` + "`" + `) IP4/IPv6 address to be used for virtual server ex: ` + "`" + `10.1.1.1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Optional , ` + "`" + `list` + "`" + `) List of server address to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional , ` + "`" + `int` + "`" + `) port number of serviceport to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional , ` + "`" + `int` + "`" + `) connectionLimit value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "priority_group",
					Description: `(Optional , ` + "`" + `int` + "`" + `) priorityGroup value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "share_nodes",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) shareNodes value to be used for FAST-Generated Pool. ### Pool Monitor Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated Pool Monitor. Unlike FAST HTTP and HTTPS apps, TCP only has the ` + "`" + `interval` + "`" + ` option here. The ` + "`" + `monitor` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional , ` + "`" + `int` + "`" + `) Set the time between health checks,in seconds for FAST-Generated Pool Monitor.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_template",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_fast_template resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Path to the zip archive file containing FAST template set on Local Disk`,
				},
				resource.Attribute{
					Name:        "md5_hash",
					Description: `(Required) MD5 hash of the zip archive file containing FAST template`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_udp_app",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details about bigip_fast_udp_app resource`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"udp",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Name of the FAST UDP application.`,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required) Name of the FAST UDP application tenant.`,
				},
				resource.Attribute{
					Name:        "virtual_server",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `virtual_server` + "`" + ` block will provide ` + "`" + `ip` + "`" + ` and ` + "`" + `port` + "`" + ` options to be used for virtual server. See [virtual server](#virtual-server) below for more details.`,
				},
				resource.Attribute{
					Name:        "enable_fastl4",
					Description: `(Optional,` + "`" + `string` + "`" + `) Enables use of FastL4 profiles.`,
				},
				resource.Attribute{
					Name:        "existing_profile",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP FastL4 or UDP profile.`,
				},
				resource.Attribute{
					Name:        "existing_snat_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP SNAT pool.`,
				},
				resource.Attribute{
					Name:        "snat_pool_address",
					Description: `(Optional,` + "`" + `list` + "`" + `) List of address to be used for FAST-Generated SNAT Pool.`,
				},
				resource.Attribute{
					Name:        "persistence_profile",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP persistence profile to be used.`,
				},
				resource.Attribute{
					Name:        "persistence_type",
					Description: `(Optional,` + "`" + `string` + "`" + `) Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.`,
				},
				resource.Attribute{
					Name:        "fallback_persistence",
					Description: `(Optional,` + "`" + `string` + "`" + `) Type of fallback persistence record to be created for each new client connection.`,
				},
				resource.Attribute{
					Name:        "existing_pool",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP pool.`,
				},
				resource.Attribute{
					Name:        "pool_members",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `pool_members` + "`" + ` block takes input for FAST-Generated Pool. See [Pool Members](#pool-members) below for more details.`,
				},
				resource.Attribute{
					Name:        "load_balancing_mode",
					Description: `(Optional,` + "`" + `string` + "`" + `) A ` + "`" + `load balancing method` + "`" + ` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method`,
				},
				resource.Attribute{
					Name:        "slow_ramp_time",
					Description: `(Optional,` + "`" + `int` + "`" + `) Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds`,
				},
				resource.Attribute{
					Name:        "existing_monitor",
					Description: `(Optional,` + "`" + `string` + "`" + `) Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `monitor` + "`" + ` block takes input for FAST-Generated Pool Monitor. See [Pool Monitor](#pool-monitor) below for more details.`,
				},
				resource.Attribute{
					Name:        "irules",
					Description: `(Optional,` + "`" + `list` + "`" + `) Irules to attach to Virtual Server.`,
				},
				resource.Attribute{
					Name:        "vlans_allowed",
					Description: `(Optional,` + "`" + `list` + "`" + `) Names of existing VLANs to allow.`,
				},
				resource.Attribute{
					Name:        "vlans_rejected",
					Description: `(Optional,` + "`" + `list` + "`" + `) Names of existing VLANs to reject.`,
				},
				resource.Attribute{
					Name:        "security_log_profiles",
					Description: `(Optional,` + "`" + `list` + "`" + `) Existing security log profiles to enable. ### virtual server This IP address, combined with the port you specify below, becomes the BIG-IP virtual server address and port, which clients use to access the application The ` + "`" + `virtual_server` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional , ` + "`" + `string` + "`" + `) IP4/IPv6 address to be used for virtual server ex: ` + "`" + `10.1.1.1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Optional , ` + "`" + `list` + "`" + `) List of server address to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional , ` + "`" + `int` + "`" + `) port number of serviceport to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional , ` + "`" + `int` + "`" + `) connectionLimit value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "priority_group",
					Description: `(Optional , ` + "`" + `int` + "`" + `) priorityGroup value to be used for FAST-Generated Pool.`,
				},
				resource.Attribute{
					Name:        "share_nodes",
					Description: `(Optional , ` + "`" + `bool` + "`" + `) shareNodes value to be used for FAST-Generated Pool. ### Pool Monitor Using this block will ` + "`" + `enable` + "`" + ` for FAST-Generated Pool Monitor. The ` + "`" + `monitor` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional , ` + "`" + `int` + "`" + `) Set the time between health checks,in seconds for FAST-Generated Pool Monitor.`,
				},
				resource.Attribute{
					Name:        "send_string",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Optional data to be sent during each health check.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Optional , ` + "`" + `string` + "`" + `) The presence of this optional string is required in the response, if specified it confirms availability.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ipsec_policy",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ipsec_policy resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example ` + "`" + `/Common/test-policy` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Description of the IPSec policy.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the IPsec protocol. Valid choices are: ` + "`" + `ah, esp` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the processing mode. Valid choices are: ` + "`" + `transport, interface, isession, tunnel` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tunnel_local_address",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnel_remote_address",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.`,
				},
				resource.Attribute{
					Name:        "encrypt_algorithm",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the algorithm to use for IKE encryption. Valid choices are: ` + "`" + `null, 3des, aes128, aes192, aes256, aes-gmac256, aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128` + "`" + ``,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the algorithm to use for IKE authentication. Valid choices are: ` + "`" + `sha1, sha256, sha384, sha512, aes-gcm128, aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256` + "`" + ``,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the length of time before the IKE security association expires, in minutes.`,
				},
				resource.Attribute{
					Name:        "kb_lifetime",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the length of time before the IKE security association expires, in kilobytes.`,
				},
				resource.Attribute{
					Name:        "perfect_forward_secrecy",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: ` + "`" + `none, modp768, modp1024, modp1536, modp2048, modp3072, modp4096, modp6144, modp8192` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ipcomp",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies whether to use IPComp encapsulation. Valid choices are: ` + "`" + `none", null", deflate` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ipsec_profile",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_ipsec_profile resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"ipsec",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example ` + "`" + `/Common/test-profile` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies descriptive text that identifies the IPsec interface tunnel profile.`,
				},
				resource.Attribute{
					Name:        "parent_profile",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the profile from which this profile inherits settings. The default is the system-supplied ` + "`" + `/Common/ipsec` + "`" + ` profile`,
				},
				resource.Attribute{
					Name:        "traffic_selector",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_datagroup",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_datagroup resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"datagroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the datagroup`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) datagroup type (applies to the ` + "`" + `name` + "`" + ` field of the record), supports: ` + "`" + `string` + "`" + `, ` + "`" + `ip` + "`" + ` or ` + "`" + `integer` + "`" + ``,
				},
				resource.Attribute{
					Name:        "records_src",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, ` + "`" + `foo := bar` + "`" + `. This should be used in conjunction with ` + "`" + `internal` + "`" + ` attribute set ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `(Optional,` + "`" + `bool` + "`" + `) Set ` + "`" + `false` + "`" + ` if you want to Create External Datagroups. default is ` + "`" + `true` + "`" + `,means creates internal datagroup.`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `(Optional) a set of ` + "`" + `name` + "`" + ` and ` + "`" + `data` + "`" + ` attributes, name must be of type specified by the ` + "`" + `type` + "`" + ` attributed (` + "`" + `string` + "`" + `, ` + "`" + `ip` + "`" + ` and ` + "`" + `integer` + "`" + `), data is optional and can take any value, multiple ` + "`" + `record` + "`" + ` sets can be specified as needed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required if ` + "`" + `record` + "`" + ` defined), sets the value of the record's ` + "`" + `name` + "`" + ` attribute, must be of type defined in ` + "`" + `type` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional if ` + "`" + `record` + "`" + ` defined), sets the value of the record's ` + "`" + `data` + "`" + ` attribute, specifying a value here will create a record in the form of ` + "`" + `name := data` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_irule",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_irule resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"irule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the iRule`,
				},
				resource.Attribute{
					Name:        "irule",
					Description: `(Required) Body of the iRule`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_monitor",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_monitor resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Parent monitor for the system to use for setting initial values for the new monitor.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of ` + "`" + `interval` + "`" + ` should be always less than ` + "`" + `timeout` + "`" + `. Default is ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "up_interval",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the interval for the system to use to perform the health check when a resource is up. The default is ` + "`" + `0(Disabled)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of seconds the target has in which to respond to the monitor request. The default is ` + "`" + `16` + "`" + ` seconds`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the text string that the monitor sends to the target object.`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the regular expression representing the text string that the monitor looks for in the returned resource.`,
				},
				resource.Attribute{
					Name:        "receive_disable",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.`,
				},
				resource.Attribute{
					Name:        "transparent",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether the monitor operates in transparent mode.`,
				},
				resource.Attribute{
					Name:        "manual_resume",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.`,
				},
				resource.Attribute{
					Name:        "ip_dscp",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Displays the differentiated services code point (DSCP).The default is ` + "`" + `0 (zero)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_until_up",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) Specifies the database in which the user is created`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specify an alias address for monitoring`,
				},
				resource.Attribute{
					Name:        "adaptive",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether adaptive response time monitoring is enabled for this monitor. The default is ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adaptive_limit",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the user name if the monitored target requires authentication`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the password if the monitored target requires authentication`,
				},
				resource.Attribute{
					Name:        "compatibility",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).`,
				},
				resource.Attribute{
					Name:        "ssl_profile",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the ssl profile for the monitor. It only makes sense when the parent is ` + "`" + `/Common/https` + "`" + ` ## Importing An existing monitor can be imported into this resource by supplying monitor Name in ` + "`" + `full path` + "`" + ` as ` + "`" + `id` + "`" + `. An example is below: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import bigip_ltm_monitor.monitor /Common/terraform_monitor ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_node",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_node resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required , type ` + "`" + `string` + "`" + `) Name of the node`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required, type ` + "`" + `string` + "`" + `) IP or hostname of the node`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) User-defined description give ltm_node`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the maximum number of connections allowed for the node or node address.`,
				},
				resource.Attribute{
					Name:        "dynamic_ratio",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the fixed ratio value used for a node during ratio load balancing.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional) specifies the name of the monitor or monitor rule that you want to associate with the node.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Default is "user-up" you can set to "user-down" if you want to disable ~>`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name). ## Importing An existing Node can be imported into this resource by supplying Node Name in ` + "`" + `full path` + "`" + ` as ` + "`" + `id` + "`" + `. An example is below: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import bigip_ltm_node.site2_node "/TEST/testnode" (or) $ terraform import bigip_ltm_node.site2_node "/Common/3.3.3.3" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_persistence_profile_cookie",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_cookie resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"persistence",
				"profile",
				"cookie",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_persistence_profile_dstaddr",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_dstaddr resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"persistence",
				"profile",
				"dstaddr",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_persistence_profile_srcaddr",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_srcaddr resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"persistence",
				"profile",
				"srcaddr",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_persistence_profile_ssl",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_ssl resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"persistence",
				"profile",
				"ssl",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_policy",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_policy resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Specifies the match strategy`,
				},
				resource.Attribute{
					Name:        "requires",
					Description: `(Optional) Specifies the protocol`,
				},
				resource.Attribute{
					Name:        "published_copy",
					Description: `(Optional) If you want to publish the policy else it will be deployed in Drafts mode.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) List of Rules can be applied using the policy. Each rule is block type with following arguments.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Name of Rule to be applied in policy.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional,type ` + "`" + `set` + "`" + `) Block type. See [condition](#condition) block for more details.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional,type ` + "`" + `set` + "`" + `) Block type. See [action](#action) block for more details.`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Optional) This action will affect forwarding.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional ) This action will direct the stream to this pool.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) This action is set to ` + "`" + `true` + "`" + ` by default, it needs to be explicitly set to ` + "`" + `false` + "`" + ` for actions it conflicts with. ### condition`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_pool",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_pool resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Name of the pool,it should be ` + "`" + `full path` + "`" + `.The full path is the combination of the ` + "`" + `partition + name` + "`" + ` of the pool.(For example ` + "`" + `/Common/my-pool` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "monitors",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) List of monitor names to associate with the pool`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies descriptive text that identifies the pool.`,
				},
				resource.Attribute{
					Name:        "allow_nat",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : ` + "`" + `yes` + "`" + `, Possible Values ` + "`" + `yes` + "`" + ` or ` + "`" + `no` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "allow_snat",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : ` + "`" + `yes` + "`" + `, Possible Values ` + "`" + `yes` + "`" + ` or ` + "`" + `no` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "load_balancing_mode",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the load balancing method. The default is Round Robin.`,
				},
				resource.Attribute{
					Name:        "minimum_active_members",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is ` + "`" + `0` + "`" + ` meaning ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slow_ramp_time",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the duration during which the system sends less traffic to a newly-enabled pool member.`,
				},
				resource.Attribute{
					Name:        "service_down_action",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies how the system should respond when the target pool member becomes unavailable. The default is ` + "`" + `None` + "`" + `, Possible values: ` + "`" + `[none, reset, reselect, drop]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reselect_tries",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the number of times the system tries to contact a new pool member after a passive failure. ## Importing An existing pool can be imported into this resource by supplying pool Name in ` + "`" + `full path` + "`" + ` as ` + "`" + `id` + "`" + `. An example is below: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import bigip_ltm_pool.k8s_prod_import /Common/k8prod_Pool ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_pool_attachment",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_pool_attachment resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"pool",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example ` + "`" + `/Common/my-pool` + "`" + `) or partition + directory + name of the pool (For example ` + "`" + `/Common/test/my-pool` + "`" + `).When including directory in fullpath we have to make sure it is created in the given partition before using it.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Required) Pool member address/fqdn with service port, (ex: ` + "`" + `1.1.1.1:80/www.google.com:80` + "`" + `). (Note: Member will be in same partition of Pool)`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.`,
				},
				resource.Attribute{
					Name:        "connection_rate_limit",
					Description: `(Optional) Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.`,
				},
				resource.Attribute{
					Name:        "dynamic_ratio",
					Description: `(Optional) Specifies the fixed ratio value used for a node during ratio load balancing.`,
				},
				resource.Attribute{
					Name:        "priority_group",
					Description: `(Optional) Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority`,
				},
				resource.Attribute{
					Name:        "fqdn_autopopulate",
					Description: `(Optional) Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled ## Importing An existing pool attachment (i.e. pool membership) can be imported into this resource by supplying both the pool full path, and the node full path with the relevant port. If the pool or node membership is not found, an error will be returned. An example is below: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import bigip_ltm_pool_attachment.node-pool-attach \ '{"pool": "/Common/terraform-pool", "node": "/Common/node1:80"}' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_client_ssl",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_client_ssl resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"client",
				"ssl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Parent profile for this clientssl profile.Once this value has been set, it cannot be changed. Default value is ` + "`" + `/Common/clientssl` + "`" + `. It Should Full path ` + "`" + `/partition/profile_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "allow_non_ssl",
					Description: `(Optional) Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile`,
				},
				resource.Attribute{
					Name:        "authenticate",
					Description: `(Optional) Specifies the frequency of client authentication for an SSL session.When ` + "`" + `once` + "`" + `,specifies that the system authenticates the client once for an SSL session. When ` + "`" + `always` + "`" + `, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.`,
				},
				resource.Attribute{
					Name:        "tm_options",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) List of Enabled selection from a set of industry standard options for handling SSL processing.By default, Don't insert empty fragments and No TLSv1.3 are listed as Enabled Options. ` + "`" + `Usage` + "`" + ` : tm_options = ["dont-insert-empty-fragments","no-tlsv1.3"]`,
				},
				resource.Attribute{
					Name:        "authenticate_depth",
					Description: `(Optional) Specifies the maximum number of certificates to be traversed in a client certificate chain`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional) Specifies a cert name for use.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Contains a key name`,
				},
				resource.Attribute{
					Name:        "chain",
					Description: `(Optional) Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.`,
				},
				resource.Attribute{
					Name:        "cipher_group",
					Description: `(Optional) Specifies the cipher group for the SSL server profile. It is mutually exclusive with the argument, ` + "`" + `ciphers` + "`" + `. The default value is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_cert_mode",
					Description: `(Optional) Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.`,
				},
				resource.Attribute{
					Name:        "renegotiation",
					Description: `(Optional) Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile`,
				},
				resource.Attribute{
					Name:        "retain_certificate",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, client certificate is retained in SSL session.`,
				},
				resource.Attribute{
					Name:        "secure_renegotiation",
					Description: `(Optional) Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile. When ` + "`" + `request` + "`" + ` is set the system request secure renegotation of SSL connections. ` + "`" + `require` + "`" + ` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients. The ` + "`" + `require-strict` + "`" + ` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Optional) Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sni_default",
					Description: `(Optional) Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile. There can be only one SSL profile with this setting enabled.`,
				},
				resource.Attribute{
					Name:        "sni_require",
					Description: `(Optional) Requires that the network peers also provide SNI support, this setting only takes effect when ` + "`" + `sni_default` + "`" + ` is set to ` + "`" + `true` + "`" + `.When creating a new profile, the setting is provided by the parent profile`,
				},
				resource.Attribute{
					Name:        "strict_resume",
					Description: `(Optional) Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.`,
				},
				resource.Attribute{
					Name:        "ssl_forward_proxy",
					Description: `(Optional) Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "ssl_forward_proxy_bypass",
					Description: `(Optional) Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_fasthttp",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_fasthttp resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"fasthttp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "connpoolidle_timeoutoverride",
					Description: `(Optional) Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.`,
				},
				resource.Attribute{
					Name:        "connpool_maxreuse",
					Description: `(Optional) Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).`,
				},
				resource.Attribute{
					Name:        "connpool_maxsize",
					Description: `(Optional) Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.`,
				},
				resource.Attribute{
					Name:        "connpool_replenish",
					Description: `(Optional) The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.`,
				},
				resource.Attribute{
					Name:        "connpool_minsize",
					Description: `(Optional) Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.`,
				},
				resource.Attribute{
					Name:        "forcehttp_10response",
					Description: `(Optional) Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "maxheader_size",
					Description: `(Optional) Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_fastl4",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_fastl4 resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"fastl4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "late_binding",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.`,
				},
				resource.Attribute{
					Name:        "explicitflow_migration",
					Description: `(Optional,type ` + "`" + `string` + "`" + `)Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.`,
				},
				resource.Attribute{
					Name:        "hardware_syncookie",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.`,
				},
				resource.Attribute{
					Name:        "iptos_toclient",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.`,
				},
				resource.Attribute{
					Name:        "keepalive_interval",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).`,
				},
				resource.Attribute{
					Name:        "tcp_handshake_timeout",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is ` + "`" + `5 seconds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "loose_initiation",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.`,
				},
				resource.Attribute{
					Name:        "loose_close",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.`,
				},
				resource.Attribute{
					Name:        "receive_windowsize",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero). ## Import BIG-IP LTM fastl4 profiles can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import bigip_ltm_profile_fastl4.test-fastl4 /Common/test-fastl4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_ftp",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_ftp resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"ftp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Displays the administrative partition within which this profile resides`,
				},
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. ### Arguments which are updated/available in Bigip versions (14.x - 16.x) for FTP profile.For more information refer below KB article https://support.f5.com/csp/article/K08859735`,
				},
				resource.Attribute{
					Name:        "ftps_mode",
					Description: `(Optional) Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow`,
				},
				resource.Attribute{
					Name:        "enforce_tlssession_reuse",
					Description: `(Optional) Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)`,
				},
				resource.Attribute{
					Name:        "allow_active_mode",
					Description: `(Optional)Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled ### Arguments which are updated/available in Bigip versions (12.x - 13.x) for FTP profile.For more information refer below KB article https://support.f5.com/csp/article/K13044205`,
				},
				resource.Attribute{
					Name:        "allow_ftps",
					Description: `(Optional)Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.`,
				},
				resource.Attribute{
					Name:        "translate_extended",
					Description: `(Optional)Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled). ## Common Arguments for all versions`,
				},
				resource.Attribute{
					Name:        "security",
					Description: `(Optional)Specifies, when checked (enabled), that the system inspects FTP traffic for security vulnerabilities using an FTP security profile. This option is available only on systems licensed for BIG-IP ASM.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)Allows you to configure the FTP service to run on an alternate port. The default is 20.`,
				},
				resource.Attribute{
					Name:        "log_profile",
					Description: `(Optional)Configures the ALG log profile that controls logging`,
				},
				resource.Attribute{
					Name:        "log_publisher",
					Description: `(Optional)Configures the log publisher that handles events logging for this profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)User defined description for FTP profile`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_http",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_http resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"http",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "proxy_type",
					Description: `(optional,type ` + "`" + `string` + "`" + `) Specifies the proxy mode for this profile: reverse, explicit, or transparent. The default is ` + "`" + `reverse` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(optional,type ` + "`" + `string` + "`" + `) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(optional,type ` + "`" + `string` + "`" + `) Specifies user-defined description.`,
				},
				resource.Attribute{
					Name:        "basic_auth_realm",
					Description: `(Optional) Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is ` + "`" + `none` + "`" + ``,
				},
				resource.Attribute{
					Name:        "fallback_host",
					Description: `(Optional) Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number`,
				},
				resource.Attribute{
					Name:        "fallback_status_codes",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) Specifies one or more three-digit status codes that can be returned by an HTTP server,that should trigger a redirection to the fallback host.`,
				},
				resource.Attribute{
					Name:        "head_erase",
					Description: `(Optional) Specifies the header string that you want to erase from an HTTP request. Default is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "head_insert",
					Description: `(Optional) Specifies a quoted header string that you want to insert into an HTTP request.Default is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "insert_xforwarded_for",
					Description: `(Optional) When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address`,
				},
				resource.Attribute{
					Name:        "response_headers_permitted",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) Specifies headers that the BIG-IP system allows in an HTTP response.If you are specifying more than one header, separate the headers with a blank space.`,
				},
				resource.Attribute{
					Name:        "request_chunking",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies how the system handles HTTP content that is chunked by a client. The default is ` + "`" + `preserve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_chunking",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies how the system handles HTTP content that is chunked by a server. The default is ` + "`" + `selective` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oneconnect_transformations",
					Description: `(Optional) Enables the system to perform HTTP header transformations for the purpose of keeping server-side connections open. This feature requires configuration of a OneConnect profile`,
				},
				resource.Attribute{
					Name:        "redirect_rewrite",
					Description: `(Optional) Specifies whether the system rewrites the URIs that are part of HTTP redirect (3XX) responses. The default is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_chunking",
					Description: `(Optional) Specifies how the system handles HTTP content that is chunked by a client. The default is ` + "`" + `preserve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encrypt_cookies",
					Description: `(Optional) Type the cookie names for the system to encrypt.`,
				},
				resource.Attribute{
					Name:        "encrypt_cookie_secret",
					Description: `(Optional) Type a passphrase for cookie encryption.`,
				},
				resource.Attribute{
					Name:        "insert_xforwarded_for",
					Description: `(Optional) Specifies, when enabled, that the system inserts an X-Forwarded-For header in an HTTP request with the client IP address, to use with connection pooling. The default is ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lws_width",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the maximum column width for any given line, when inserting an HTTP header in an HTTP request. The default is ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lws_width",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the linear white space (LWS) separator that the system inserts when a header exceeds the maximum width you specify in the LWS Maximum Columns setting.`,
				},
				resource.Attribute{
					Name:        "accept_xff",
					Description: `(Optional) Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's XFF (X-forwarded-for) headers, if they exist.`,
				},
				resource.Attribute{
					Name:        "xff_alternative_names",
					Description: `(Optional) Specifies alternative XFF headers instead of the default X-forwarded-for header. ## Import BIG-IP LTM http profiles can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `bash terraform import bigip_ltm_profile_http.test-http /Common/test-http ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_http2",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_http2 resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"http2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "concurrent_streams_per_connection",
					Description: `(Optional,` + "`" + `type int` + "`" + `) Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.`,
				},
				resource.Attribute{
					Name:        "connection_idle_timeout",
					Description: `(Optional,` + "`" + `type int` + "`" + `) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion.`,
				},
				resource.Attribute{
					Name:        "insert_header",
					Description: `(Optional,` + "`" + `type string` + "`" + `) This setting specifies whether the BIG-IP system should add an HTTP header to the HTTP request to show that the request was received over HTTP/2, Allowed Values : ` + "`" + `"enabled"/"disabled"` + "`" + ` [ Default: ` + "`" + `"disabled"` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "insert_header_name",
					Description: `(Optional,` + "`" + `type string` + "`" + `) This setting specifies the name of the header that the BIG-IP system will add to the HTTP request when the Insert Header is enabled.`,
				},
				resource.Attribute{
					Name:        "header_table_size",
					Description: `(Optional) The size of the header table, in KB, for the HTTP headers that the HTTP/2 protocol compresses to save bandwidth.`,
				},
				resource.Attribute{
					Name:        "enforce_tls_requirements",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Enable or disable enforcement of TLS requirements,Allowed Values : ` + "`" + `"enabled"/"disabled"` + "`" + ` [Default:` + "`" + `"enabled"` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "frame_size",
					Description: `(Optional,` + "`" + `type int` + "`" + `) The size of the data frames, in bytes, that the HTTP/2 protocol sends to the client. ` + "`" + `Default: 2048` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receive_window",
					Description: `(Optional,` + "`" + `type int` + "`" + `) The flow-control size for upload streams, in KB. ` + "`" + `Default: 32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "write_size",
					Description: `(Optional,` + "`" + `type int` + "`" + `) The total size of combined data frames, in bytes, that the HTTP/2 protocol sends in a single write function. ` + "`" + `Default: 16384` + "`" + `".`,
				},
				resource.Attribute{
					Name:        "activation_modes",
					Description: `(Optional) This setting specifies the condition that will cause the BIG-IP system to handle an incoming connection as an HTTP/2 connection, Allowed values : ` + "`" + `[“alpn”]` + "`" + ` (or) ` + "`" + `[“always”]` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_httpcompress",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_httpcompress resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"httpcompress",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "content_type_include",
					Description: `(Optional,type ` + "`" + `set` + "`" + `) Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.`,
				},
				resource.Attribute{
					Name:        "content_type_exclude",
					Description: `(Optional,type ` + "`" + `set` + "`" + `) Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.`,
				},
				resource.Attribute{
					Name:        "compression_buffersize",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the maximum number of compressed bytes that the system buffers before inserting a Content-Length header (which specifies the compressed size) into the response. The default is ` + "`" + `4096` + "`" + ` bytes.`,
				},
				resource.Attribute{
					Name:        "gzip_compression_level",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the degree to which the system compresses the content. Higher compression levels cause the compression process to be slower. The default is 1 - Least Compression (Fastest)`,
				},
				resource.Attribute{
					Name:        "gzip_memory_level",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of bytes of memory that the system uses for internal compression buffers when compressing a server response. The default is ` + "`" + `8 kilobytes/8192 bytes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_window_size",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The default is ` + "`" + `16` + "`" + ` kilobytes`,
				},
				resource.Attribute{
					Name:        "keep_accept_encoding",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system does not remove the Accept-Encoding: header from an HTTP request. The default is ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vary_header",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system inserts a Vary header into cacheable server responses. The default is ` + "`" + `enabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_saver",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system monitors the percent CPU usage and adjusts compression rates automatically when the CPU usage reaches either the CPU Saver High Threshold or the CPU Saver Low Threshold. The default is ` + "`" + `enabled` + "`" + `. ## Import BIG-IP LTM HTTP Compress profiles can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import bigip_ltm_profile_httpcompress.test-httpcomprs_import /Common/test-httpcomprs ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_oneconnect",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_oneconnect resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"oneconnect",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Displays the administrative partition within which this profile resides`,
				},
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_override",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are ` + "`" + `disabled` + "`" + `, ` + "`" + `indefinite` + "`" + `, or a numeric value that you specify. The default value is ` + "`" + `disabled` + "`" + ``,
				},
				resource.Attribute{
					Name:        "limit_type",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Controls how connection limits are enforced in conjunction with OneConnect. The default is ` + "`" + `None` + "`" + `. Supported Values: ` + "`" + `[None,idle,strict]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "share_pools",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Specify if you want to share the pool, default value is ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional,` + "`" + `type int` + "`" + `) Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is ` + "`" + `86400` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_reuse",
					Description: `(Optional,` + "`" + `type int` + "`" + `) Specifies the maximum number of times that a server-side connection can be reused. The default value is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional,` + "`" + `type int` + "`" + `) Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is ` + "`" + `10000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_mask",
					Description: `(Optional,` + "`" + `type string` + "`" + `) Specifies a source IP mask. The default value is ` + "`" + `0.0.0.0` + "`" + `. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address. ## Import BIG-IP LTM oneconnect profiles can be imported using the ` + "`" + `name` + "`" + ` , e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import bigip_ltm_profile_oneconnect.test-oneconnect /Common/test-oneconnect ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_server_ssl",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_server_ssl resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"server",
				"ssl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is ` + "`" + `/Common/serverssl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional) Specifies the name of the certificate that the system uses for server-side SSL processing.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Specifies the file name of the SSL key.`,
				},
				resource.Attribute{
					Name:        "chain",
					Description: `(Optional) Specifies the certificates-key chain to associate with the SSL profile`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.`,
				},
				resource.Attribute{
					Name:        "cipher_group",
					Description: `(Optional) Specifies the cipher group for the SSL server profile. It is mutually exclusive with the argument, ` + "`" + `ciphers` + "`" + `. The default value is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_cert_mode",
					Description: `(Optional) Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.`,
				},
				resource.Attribute{
					Name:        "authenticate",
					Description: `(Optional) Specifies the frequency of server authentication for an SSL session.When ` + "`" + `once` + "`" + `,specifies that the system authenticates the server once for an SSL session. When ` + "`" + `always` + "`" + `, specifies that the system authenticates the server once for an SSL session and also upon reuse of that session.`,
				},
				resource.Attribute{
					Name:        "tm_options",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) List of Enabled selection from a set of industry standard options for handling SSL processing.By default, Don't insert empty fragments and No TLSv1.3 are listed as Enabled Options. ` + "`" + `Usage` + "`" + ` : tm_options = ["dont-insert-empty-fragments","no-tlsv1.3"]`,
				},
				resource.Attribute{
					Name:        "renegotiation",
					Description: `(Optional) Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile`,
				},
				resource.Attribute{
					Name:        "retain_certificate",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, client certificate is retained in SSL session.`,
				},
				resource.Attribute{
					Name:        "secure_renegotiation",
					Description: `(Optional) Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile. When ` + "`" + `request` + "`" + ` is set the system request secure renegotation of SSL connections. ` + "`" + `require` + "`" + ` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients. The ` + "`" + `require-strict` + "`" + ` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Optional) Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sni_default",
					Description: `(Optional) Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile. There can be only one SSL profile with this setting enabled.`,
				},
				resource.Attribute{
					Name:        "sni_require",
					Description: `(Optional) Requires that the network peers also provide SNI support, this setting only takes effect when ` + "`" + `sni_default` + "`" + ` is set to ` + "`" + `true` + "`" + `.When creating a new profile, the setting is provided by the parent profile`,
				},
				resource.Attribute{
					Name:        "strict_resume",
					Description: `(Optional) Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.`,
				},
				resource.Attribute{
					Name:        "ssl_forward_proxy",
					Description: `(Optional) Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "ssl_forward_proxy_bypass",
					Description: `(Optional) Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_profile_tcp",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_profile_tcp resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"profile",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.`,
				},
				resource.Attribute{
					Name:        "close_wait_timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.`,
				},
				resource.Attribute{
					Name:        "finwait_timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.`,
				},
				resource.Attribute{
					Name:        "finwait_2timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).`,
				},
				resource.Attribute{
					Name:        "keepalive_interval",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.`,
				},
				resource.Attribute{
					Name:        "zerowindow_timeout",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.`,
				},
				resource.Attribute{
					Name:        "send_buffersize",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the SEND window size. The default is 131072 bytes.`,
				},
				resource.Attribute{
					Name:        "receive_windowsize",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.`,
				},
				resource.Attribute{
					Name:        "proxybuffer_high",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the proxy buffer level, in bytes, at which the receive window is closed.`,
				},
				resource.Attribute{
					Name:        "congestion_control",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.`,
				},
				resource.Attribute{
					Name:        "initial_congestion_windowsize",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.`,
				},
				resource.Attribute{
					Name:        "delayed_acks",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.`,
				},
				resource.Attribute{
					Name:        "nagle",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.`,
				},
				resource.Attribute{
					Name:        "early_retransmit",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.`,
				},
				resource.Attribute{
					Name:        "tailloss_probe",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.`,
				},
				resource.Attribute{
					Name:        "timewait_recycle",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.`,
				},
				resource.Attribute{
					Name:        "fast_open",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is ` + "`" + `enabled` + "`" + `. If ` + "`" + `fast_open` + "`" + ` set to ` + "`" + `enabled` + "`" + `, argument ` + "`" + `verified_accept` + "`" + ` can't be set to ` + "`" + `enabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verified_accept",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deferred_accept",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled. ## Importing An existing tcp profile can be imported into this resource by supplying tcp profile Name in ` + "`" + `full path` + "`" + ` as ` + "`" + `id` + "`" + `. An example is below: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import bigip_ltm_profile_tcp.tcp-lan-profile-import /Common/test-tcp-lan-profile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_snat",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_snat resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SNAT, name of SNAT should be full path. Full path is the combination of the ` + "`" + `partition + SNAT name` + "`" + `,For example ` + "`" + `/Common/test-snat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origins",
					Description: `(Required) Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports`,
				},
				resource.Attribute{
					Name:        "translation",
					Description: `(Optional) Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when ` + "`" + `automap` + "`" + ` and ` + "`" + `snatpool` + "`" + ` are not used.`,
				},
				resource.Attribute{
					Name:        "snatpool",
					Description: `(Optional) Specifies the name of a SNAT pool. You can only use this option when ` + "`" + `automap` + "`" + ` and ` + "`" + `translation` + "`" + ` are not used.`,
				},
				resource.Attribute{
					Name:        "mirror",
					Description: `(Optional) Enables or disables mirroring of SNAT connections.`,
				},
				resource.Attribute{
					Name:        "sourceport",
					Description: `(Optional) Specifies how the SNAT object handles the client's source port. The default is ` + "`" + `preserve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vlansdisabled",
					Description: `(Optional,bool) Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is ` + "`" + `true` + "`" + `, vlandisabled on VLANS specified by ` + "`" + `vlans` + "`" + `,if set to ` + "`" + `false` + "`" + ` vlanEnabled set on VLANS specified by ` + "`" + `vlans` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional) Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_snatpool",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_snatpool resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"snatpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the snatpool`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_virtual_address",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_virtual_address resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"virtual",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the virtual address`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the virtual address`,
				},
				resource.Attribute{
					Name:        "advertize_route",
					Description: `(Optional) Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `(Optional, Default=0) Max number of connections for virtual address`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, Default=true) Enable or disable the virtual address`,
				},
				resource.Attribute{
					Name:        "arp",
					Description: `(Optional, Default=true) Enable or disable ARP for the virtual address`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional, Default=true) Automatically delete the virtual address with the virtual server`,
				},
				resource.Attribute{
					Name:        "icmp_echo",
					Description: `(Optional, Default=enabled) Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `(Optional, Default=/Common/traffic-group-1) Specify the partition and traffic group`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_virtual_server",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_virtual_server resource`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Listen port for the virtual server`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination IP`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of Virtual server`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) Default pool name`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `(Optional) Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0`,
				},
				resource.Attribute{
					Name:        "source_address_translation",
					Description: `(Optional) Can be either omitted for ` + "`" + `none` + "`" + ` or the values ` + "`" + `automap` + "`" + ` options : [` + "`" + `snat` + "`" + `,` + "`" + `automap` + "`" + `,` + "`" + `none` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "translate_address",
					Description: `Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.`,
				},
				resource.Attribute{
					Name:        "translate_port",
					Description: `Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(Optional) List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.`,
				},
				resource.Attribute{
					Name:        "client_profiles",
					Description: `(Optional) List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles`,
				},
				resource.Attribute{
					Name:        "server_profiles",
					Description: `(Optional) List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Specifies an IP address or network from which the virtual server will accept traffic.`,
				},
				resource.Attribute{
					Name:        "irules",
					Description: `(Optional) The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.`,
				},
				resource.Attribute{
					Name:        "snatpool",
					Description: `(Optional) Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional) The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute ` + "`" + `vlan_enabled` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vlans_enabled",
					Description: `(Optional Bool) Enables the virtual server on the VLANs specified by the ` + "`" + `vlans` + "`" + ` option. By default it is ` + "`" + `false` + "`" + ` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by ` + "`" + `vlans` + "`" + `, mark this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence_profiles",
					Description: `(Optional) List of persistence profiles associated with the Virtual Server.`,
				},
				resource.Attribute{
					Name:        "fallback_persistence_profile",
					Description: `(Optional) Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.`,
				},
				resource.Attribute{
					Name:        "security_log_profiles",
					Description: `(Optional) Specifies the log profile applied to the virtual server.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies whether the system preserves the source port of the connection. The default is ` + "`" + `preserve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firewall_enforced_policy",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: ` + "`" + `/Common/afm-test-policy` + "`" + `. ## Importing An existing virtual-server can be imported into this resource by supplying virtual-server Name in ` + "`" + `full path` + "`" + ` as ` + "`" + `id` + "`" + `. An example is below: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import bigip_ltm_virtual_server.http /Common/terraform_vs_http ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_net_ike_peer",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_net_ike_peer resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"net",
				"ike",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ike_peer`,
				},
				resource.Attribute{
					Name:        "app_service",
					Description: `(Optional)The application service that the object belongs to`,
				},
				resource.Attribute{
					Name:        "ca_cert_file",
					Description: `(Optional)the trusted root and intermediate certificate authorities`,
				},
				resource.Attribute{
					Name:        "crl_file",
					Description: `(Optional)Specifies the file name of the Certificate Revocation List. Only supported in IKEv1`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)User defined description`,
				},
				resource.Attribute{
					Name:        "generate_policy",
					Description: `(Optional)Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional)Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder`,
				},
				resource.Attribute{
					Name:        "my_cert_file",
					Description: `(Optional)Specifies the name of the certificate file object`,
				},
				resource.Attribute{
					Name:        "my_cert_key_file",
					Description: `(Optional)Specifies the name of the certificate key file object`,
				},
				resource.Attribute{
					Name:        "my_cert_key_passphrase",
					Description: `(Optional)Specifies the passphrase of the key used for my-cert-key-file`,
				},
				resource.Attribute{
					Name:        "my_id_type",
					Description: `(Optional)Specifies the identifier type sent to the remote host to use in the phase 1 negotiation`,
				},
				resource.Attribute{
					Name:        "my_id_value",
					Description: `(Optional)Specifies the identifier value sent to the remote host in the phase 1 negotiation`,
				},
				resource.Attribute{
					Name:        "nat_traversal",
					Description: `(Optional)Enables use of the NAT-Traversal IPsec extension`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `(Optional)Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer`,
				},
				resource.Attribute{
					Name:        "peers_cert_file",
					Description: `(Optional)Specifies the peer’s certificate for authentication`,
				},
				resource.Attribute{
					Name:        "peers_cert_type",
					Description: `(Optional)Specifies that the only peers-cert-type supported is certfile`,
				},
				resource.Attribute{
					Name:        "peers_id_type",
					Description: `(Optional)Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type`,
				},
				resource.Attribute{
					Name:        "peers_id_value",
					Description: `(Optional)Specifies the peer’s identifier to be received`,
				},
				resource.Attribute{
					Name:        "phase1_auth_method",
					Description: `(Optional)Specifies the authentication method used for phase 1 negotiation`,
				},
				resource.Attribute{
					Name:        "phase1_encrypt_algorithm",
					Description: `(Optional)Specifies the encryption algorithm used for the isakmp phase 1 negotiation`,
				},
				resource.Attribute{
					Name:        "phase1_hash_algorithm",
					Description: `(Optional)Defines the hash algorithm used for the isakmp phase 1 negotiation`,
				},
				resource.Attribute{
					Name:        "phase1_perfect_forward_secrecy",
					Description: `(Optional)Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy`,
				},
				resource.Attribute{
					Name:        "preshared_key",
					Description: `(Optional)Specifies the preshared key for ISAKMP SAs`,
				},
				resource.Attribute{
					Name:        "preshared_key_encrypted",
					Description: `(Optional)Display the encrypted preshared-key for the IKE remote node`,
				},
				resource.Attribute{
					Name:        "prf",
					Description: `(Optional) Specifies the pseudo-random function used to derive keying material for all cryptographic operations`,
				},
				resource.Attribute{
					Name:        "proxy_support",
					Description: `(Optional)If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs`,
				},
				resource.Attribute{
					Name:        "remote_address",
					Description: `(Required)Specifies the IP address of the IKE remote node`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional)Enables or disables this IKE remote node`,
				},
				resource.Attribute{
					Name:        "traffic_selector",
					Description: `(Optional)Specifies the names of the traffic-selector objects associated with this ike-peer`,
				},
				resource.Attribute{
					Name:        "verify_cert",
					Description: `(Optional)Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional)Specifies which version of IKE to be used`,
				},
				resource.Attribute{
					Name:        "dpd_delay",
					Description: `(Optional)Specifies the number of seconds between Dead Peer Detection messages`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional)Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations`,
				},
				resource.Attribute{
					Name:        "replay_window_size",
					Description: `(Optional)Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_net_route",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_net_route resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"net",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the route.Name of Route should be full path,full path is the combination of the ` + "`" + `partition + route name` + "`" + `,For ex: ` + "`" + `/Common/test-net-route` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The destination subnet and netmask for the route.`,
				},
				resource.Attribute{
					Name:        "gw",
					Description: `(Optional) Specifies a gateway address for the route.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_net_selfip",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_net_selfip resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"net",
				"selfip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the selfip`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The Self IP's address and netmask. The IP address could also contain the route domain, e.g. ` + "`" + `10.12.13.14%4/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `(Optional) Specifies the traffic group, defaults to ` + "`" + `traffic-group-local-only` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "port_lockdown",
					Description: `(Optional) Specifies the port lockdown, defaults to ` + "`" + `Allow None` + "`" + ` if not specified.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_net_tunnel",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_net_tunnel resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"net",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the tunnel`,
				},
				resource.Attribute{
					Name:        "local_address",
					Description: `(Required) Specifies a local IP address. This option is required`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Required) Specifies the profile that you want to associate with the tunnel`,
				},
				resource.Attribute{
					Name:        "app_service",
					Description: `(Optional) The application service that the object belongs to`,
				},
				resource.Attribute{
					Name:        "auto_last_hop",
					Description: `(Optional) Specifies whether auto lasthop is enabled or not`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) User defined description`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Specifies how the tunnel carries traffic`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Displays the admin-partition within which this component resides`,
				},
				resource.Attribute{
					Name:        "remote_address",
					Description: `(Optional) Specifies a remote IP address`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `(Optional) Specifies a secondary non-floating IP address when the local-address is set to a floating address`,
				},
				resource.Attribute{
					Name:        "tos",
					Description: `(Optional) Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `(Optional) Specifies a traffic-group for use with the tunnel`,
				},
				resource.Attribute{
					Name:        "transparent",
					Description: `(Optional) Enables or disables the tunnel to be transparent`,
				},
				resource.Attribute{
					Name:        "use_pmtu",
					Description: `(Optional) Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Specifies an idle timeout for wildcard tunnels in seconds`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key field may represent different values depending on the type of the tunnel`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Specifies the maximum transmission unit (MTU) of the tunnel`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_net_vlan",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_net_vlan resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"net",
				"vlan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the vlan`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Specifies a number that the system adds into the header of any frame passing through the VLAN.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Specifies which interfaces you want this VLAN to use for traffic management.`,
				},
				resource.Attribute{
					Name:        "vlanport",
					Description: `Physical or virtual port used for traffic`,
				},
				resource.Attribute{
					Name:        "cmp_hash",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies how the traffic on the VLAN will be disaggregated. The value selected determines the traffic disaggregation method. possible options: [` + "`" + `default` + "`" + `, ` + "`" + `src-ip` + "`" + `, ` + "`" + `dst-ip` + "`" + `]`,
				},
				resource.Attribute{
					Name:        "tagged",
					Description: `Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ssl_certificate",
			Category:         "System",
			ShortDescription: `Provides details about bigip_ssl_certificate resource`,
			Description:      ``,
			Keywords: []string{
				"system",
				"ssl",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Content of certificate on Local Disk,path of SSL certificate will be provided to terraform ` + "`" + `file` + "`" + ` function`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Partition on to SSL Certificate to be imported. The parameter is not required when running terraform import operation. In such case the name must be provided in full_path format.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ssl_key",
			Category:         "System",
			ShortDescription: `Provides details about bigip_ssl_key resource`,
			Description:      ``,
			Keywords: []string{
				"system",
				"ssl",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Content of certificate key on Local Disk,path of SSL certificate key will be provided to terraform ` + "`" + `file` + "`" + ` function`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Partition on to SSL Certificate key to be imported. The parameter is not required when running terraform import operation. In such case the name must be provided in ` + "`" + `full_path` + "`" + ` format.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_sys_dns",
			Category:         "System",
			ShortDescription: `Provides details about bigip_sys_dns resource`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sys",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Required,type ` + "`" + `list` + "`" + ` ) Specifies the name servers that the system uses to validate DNS lookups, and resolve host names.`,
				},
				resource.Attribute{
					Name:        "number_of_dots",
					Description: `(Optional,type ` + "`" + `int` + "`" + ` ) Configures the number of dots needed in a name before an initial absolute query will be made.`,
				},
				resource.Attribute{
					Name:        "search",
					Description: `(Optional,type ` + "`" + `list` + "`" + ` ) Specifies the domains that the system searches for local domain lookups, to resolve local host names.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_sys_iapp",
			Category:         "System",
			ShortDescription: `Provides details about bigip_sys_iapp resource`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sys",
				"iapp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the iApp.`,
				},
				resource.Attribute{
					Name:        "jsonfile",
					Description: `Refer to the Json file which will be deployed on F5 BIG-IP. ## Example Usage of Json file ` + "`" + `` + "`" + `` + "`" + ` { "fullPath":"/Common/simplehttp.app/simplehttp", "generation":222, "inheritedDevicegroup":"true", "inheritedTrafficGroup":"true", "kind":"tm:sys:application:service:servicestate", "name":"simplehttp", "partition":"Common", "selfLink":"https://localhost/mgmt/tm/sys/application/service/~Common~simplehttp.app~simplehttp?ver=13.0.0", "strictUpdates":"enabled", "subPath":"simplehttp.app", "tables":[ { "name":"basic__snatpool_members" }, { "name":"net__snatpool_members" }, { "name":"optimizations__hosts" }, { "columnNames":[ "name" ], "name":"pool__hosts", "rows":[ { "row":[ "f5.cisco.com" ] } ] }, { "columnNames":[ "addr", "port", "connection_limit" ], "name":"pool__members", "rows":[ { "row":[ "10.0.2.167", "80", "0" ] }, { "row":[ "10.0.2.168", "80", "0" ] } ] }, { "name":"server_pools__servers" } ], "template":"/Common/f5.http", "templateModified":"no", "templateReference":{ "link":"https://localhost/mgmt/tm/sys/application/template/~Common~f5.http?ver=13.0.0" }, "trafficGroup":"/Common/traffic-group-1", "trafficGroupReference":{ "link":"https://localhost/mgmt/tm/cm/traffic-group/~Common~traffic-group-1?ver=13.0.0" }, "variables":[ { "encrypted":"no", "name":"client__http_compression", "value":"/#create_new#" }, { "encrypted":"no", "name":"monitor__monitor", "value":"/Common/http" }, { "encrypted":"no", "name":"net__client_mode", "value":"wan" }, { "encrypted":"no", "name":"net__server_mode", "value":"lan" }, { "encrypted":"no", "name":"net__v13_tcp", "value":"warn" }, { "encrypted":"no", "name":"pool__addr", "value":"10.0.1.100" }, { "encrypted":"no", "name":"pool__pool_to_use", "value":"/#create_new#" }, { "encrypted":"no", "name":"pool__port", "value":"80" }, { "encrypted":"no", "name":"ssl__mode", "value":"no_ssl" }, { "encrypted":"no", "name":"ssl_encryption_questions__advanced", "value":"no" }, { "encrypted":"no", "name":"ssl_encryption_questions__help", "value":"hide" } ] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `User defined description.`,
				},
				resource.Attribute{
					Name:        "deviceGroup",
					Description: `The name of the device group that the application service is assigned to.`,
				},
				resource.Attribute{
					Name:        "executeAction",
					Description: `Run the specified template action associated with the application.`,
				},
				resource.Attribute{
					Name:        "inheritedTrafficGroup",
					Description: `Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Displays the administrative partition within which the application resides.`,
				},
				resource.Attribute{
					Name:        "strictUpdates",
					Description: `Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.`,
				},
				resource.Attribute{
					Name:        "templateModified",
					Description: `Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.`,
				},
				resource.Attribute{
					Name:        "templatePrerequisiteErrors",
					Description: `Indicates any missing prerequisites associated with the template that defines this application.`,
				},
				resource.Attribute{
					Name:        "trafficGroup",
					Description: `The name of the traffic group that the application service is assigned to.`,
				},
				resource.Attribute{
					Name:        "lists",
					Description: `string values`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `User defined generic data for the application service. It is a name and value pair.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `Values provided like pool name, nodes etc.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `Name, values, encrypted or not`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_sys_ntp",
			Category:         "System",
			ShortDescription: `Provides details about bigip_sys_ntp resource`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sys",
				"ntp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required,type ` + "`" + `string` + "`" + `) User defined description.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Required,type ` + "`" + `list` + "`" + `) Specifies the time servers that the system uses to update the system time.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the time zone that you want to use for the system time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_sys_provision",
			Category:         "System",
			ShortDescription: `Provides details about module provision resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sys",
				"provision",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Name of module to provision in BIG-IP. possible options:`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to ` + "`" + `dedicated` + "`" + ` requires setting all others to ` + "`" + `none` + "`" + `. Setting the level of a module to ` + "`" + `none` + "`" + ` means the module is not activated. default is ` + "`" + `nominal` + "`" + ` possible options:`,
				},
				resource.Attribute{
					Name:        "cpu_ratio",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none`,
				},
				resource.Attribute{
					Name:        "disk_ratio",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none`,
				},
				resource.Attribute{
					Name:        "memory_ratio",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_sys_snmp",
			Category:         "System",
			ShortDescription: `Provides details about SNMP resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sys",
				"snmp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sys_contact",
					Description: `(Optional) Specifies the contact information for the system administrator.`,
				},
				resource.Attribute{
					Name:        "sys_location",
					Description: `Describes the system's physical location.`,
				},
				resource.Attribute{
					Name:        "allowedaddresses",
					Description: `Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_sys_snmp_traps",
			Category:         "System",
			ShortDescription: `Provides details about snmp_traps resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sys",
				"snmp",
				"traps",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the snmp trap.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Optional) Specifies the community string used for this trap.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host the trap will be sent to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The port that the trap will be sent to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) User defined description.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_traffic_selector",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_traffic_selector resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"traffic",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example ` + "`" + `/Common/test-selector` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Description of the traffic selector.`,
				},
				resource.Attribute{
					Name:        "destination_address",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional,type ` + "`" + `int` + "`" + `) Specifies the IP port used by the application. The default value is ` + "`" + `All Ports (0)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the IP port used by the application. The default value is ` + "`" + `All Ports (0)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is ` + "`" + `Both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_policy",
					Description: `(Optional, type ` + "`" + `string` + "`" + `) Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is ` + "`" + `default-ipsec-policy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number). When creating a new traffic selector, if this parameter is not specified, the default is ` + "`" + `last` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional, type ` + "`" + `int` + "`" + `) Specifies the network protocol to use for this traffic. The default value is ` + "`" + `All Protocols (255)` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_vcmp_guest",
			Category:         "Network",
			ShortDescription: `Provides details about bigip_vcmp_guest resource`,
			Description:      ``,
			Keywords: []string{
				"network",
				"vcmp",
				"guest",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the vCMP guest`,
				},
				resource.Attribute{
					Name:        "initial_image",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Specifies the base software release ISO image file for installing the TMOS hypervisor instance.`,
				},
				resource.Attribute{
					Name:        "initial_hotfix",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Specifies the hotfix ISO image file which is applied on top of the base image.`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional, ` + "`" + `list` + "`" + `) Specifies the list of VLANs the vCMP guest uses to communicate with other guests, the host, and with the external network. The naming format must be the combination of the partition + name. For example /Common/my-vlan`,
				},
				resource.Attribute{
					Name:        "mgmt_network",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Specifies the method by which the management address is used in the vCMP guest. options : [` + "`" + `bridged` + "`" + `,` + "`" + `isolated` + "`" + `,` + "`" + `host-only` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "mgmt_address",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Specifies the IP address and subnet or subnet mask you use to access the guest when you want to manage a module running within the guest.`,
				},
				resource.Attribute{
					Name:        "mgmt_route",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Specifies the gateway address for the ` + "`" + `mgmt_address` + "`" + `. Can be set to ` + "`" + `none` + "`" + ` to remove the value from the configuration.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional, ` + "`" + `string` + "`" + `) Specifies the state of the vCMP guest on the system. options : [` + "`" + `configured` + "`" + `,` + "`" + `provisioned` + "`" + `,` + "`" + `deployed` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "cores_per_slot",
					Description: `(Optional, ` + "`" + `int` + "`" + `) Specifies the number of cores the system allocates to the guest.`,
				},
				resource.Attribute{
					Name:        "number_of_slots",
					Description: `(Optional, ` + "`" + `int` + "`" + `) Specifies the number of slots for the system to use when creating the guest.`,
				},
				resource.Attribute{
					Name:        "min_number_of_slots",
					Description: `(Optional, ` + "`" + `int` + "`" + `) Specifies the minimum number of slots the guest must be assigned to in order to deploy.`,
				},
				resource.Attribute{
					Name:        "delete_virtual_disk",
					Description: `(Optional, ` + "`" + `bool` + "`" + `) Indicates if virtual disk associated with vCMP guest should be removed during remove operation. The default is ` + "`" + `true` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_waf_policy",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides details about bigip_waf_policy resource`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required,type ` + "`" + `string` + "`" + `) The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Specifies the name of the template used for the policy creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the description of the policy.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) Specifies the partition of the policy. Default is ` + "`" + `Common` + "`" + ``,
				},
				resource.Attribute{
					Name:        "application_language",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is ` + "`" + `utf-8` + "`" + ``,
				},
				resource.Attribute{
					Name:        "case_insensitive",
					Description: `(Optional,type ` + "`" + `bool` + "`" + `) Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration`,
				},
				resource.Attribute{
					Name:        "enable_passivemode",
					Description: `(Optional,type ` + "`" + `bool` + "`" + `) Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.`,
				},
				resource.Attribute{
					Name:        "protocol_independent",
					Description: `(Optional,type ` + "`" + `bool` + "`" + `) When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.`,
				},
				resource.Attribute{
					Name:        "enforcement_mode",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) How the system processes a request that triggers a security policy violation`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) The type of policy you want to create. The default policy type is ` + "`" + `security` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_technologies",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) This section defines parameters that the security policy permits in requests.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.`,
				},
				resource.Attribute{
					Name:        "signature_sets",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.`,
				},
				resource.Attribute{
					Name:        "signatures",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) This section defines the properties of a signature on the policy.`,
				},
				resource.Attribute{
					Name:        "policy_builder",
					Description: `(Optional,` + "`" + `set` + "`" + `) ` + "`" + `policy_builder` + "`" + ` block will provide ` + "`" + `learning_mode` + "`" + ` options to be used for policy builder. See [policy builder](#policy-builder) below for more details.`,
				},
				resource.Attribute{
					Name:        "graphql_profiles",
					Description: `(Optional,` + "`" + `list of set` + "`" + `) ` + "`" + `graphql_profiles` + "`" + ` takes list of graphql profile options to be used for policy builder. See [graphql profiles](#graphql-profiles) below for more details.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(Optional,` + "`" + `list of set` + "`" + `) ` + "`" + `file_types` + "`" + ` takes list of file-types options to be used for policy builder. See [file types](#file-types) below for more details.`,
				},
				resource.Attribute{
					Name:        "open_api_files",
					Description: `(Optional,type ` + "`" + `list` + "`" + `) This section defines the Link for open api files on the policy.`,
				},
				resource.Attribute{
					Name:        "policy_import_json",
					Description: `(Optional,type ` + "`" + `string` + "`" + `) The payload of the WAF Policy to be used for IMPORT on to BIG-IP. ### policy builder The ` + "`" + `policy_builder` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "learning_mode",
					Description: `(Optional , ` + "`" + `string` + "`" + `) learning mode setting for policy-builder, possible options: [` + "`" + `automatic` + "`" + `,` + "`" + `disabled` + "`" + `, ` + "`" + `manual` + "`" + `] ### graphql profiles The ` + "`" + `graphql_profile` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional , ` + "`" + `string` + "`" + `) name of graphql profile to be used for policy config. ### file types The ` + "`" + `file_types` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Specifies the file type name as appearing in the URL extension.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional , ` + "`" + `string` + "`" + `) Determines the type of the name attribute. Only when setting the type to ` + "`" + `wildcard` + "`" + ` will the special wildcard characters in the name be interpreted as such ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The id of the A.WAF Policy as it would be calculated on the BIG-IP.`,
				},
				resource.Attribute{
					Name:        "policy_export_json",
					Description: `Exported WAF policy deployed on BIGIP. ## Import An existing WAF Policy or if the WAF Policy has been manually created or modified on the BIG-IP WebUI, it can be imported using its ` + "`" + `id` + "`" + `. e.g: ` + "`" + `` + "`" + `` + "`" + ` terraform import bigip_waf_policy.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `The id of the A.WAF Policy as it would be calculated on the BIG-IP.`,
				},
				resource.Attribute{
					Name:        "policy_export_json",
					Description: `Exported WAF policy deployed on BIGIP. ## Import An existing WAF Policy or if the WAF Policy has been manually created or modified on the BIG-IP WebUI, it can be imported using its ` + "`" + `id` + "`" + `. e.g: ` + "`" + `` + "`" + `` + "`" + ` terraform import bigip_waf_policy.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"bigip_bigip_as3":                             0,
		"bigip_bigip_bigiq_as3":                       1,
		"bigip_bigip_cm_device":                       2,
		"bigip_bigip_cm_devicegroup":                  3,
		"bigip_bigip_command":                         4,
		"bigip_bigip_common_license_manage_bigiq":     5,
		"bigip_bigip_do":                              6,
		"bigip_bigip_event_service_discovery":         7,
		"bigip_bigip_fast_application":                8,
		"bigip_bigip_fast_http_app":                   9,
		"bigip_bigip_fast_https_app":                  10,
		"bigip_bigip_fast_tcp_app":                    11,
		"bigip_bigip_fast_template":                   12,
		"bigip_bigip_fast_udp_app":                    13,
		"bigip_bigip_ipsec_policy":                    14,
		"bigip_bigip_ipsec_profile":                   15,
		"bigip_bigip_ltm_datagroup":                   16,
		"bigip_bigip_ltm_irule":                       17,
		"bigip_bigip_ltm_monitor":                     18,
		"bigip_bigip_ltm_node":                        19,
		"bigip_bigip_ltm_persistence_profile_cookie":  20,
		"bigip_bigip_ltm_persistence_profile_dstaddr": 21,
		"bigip_bigip_ltm_persistence_profile_srcaddr": 22,
		"bigip_bigip_ltm_persistence_profile_ssl":     23,
		"bigip_bigip_ltm_policy":                      24,
		"bigip_bigip_ltm_pool":                        25,
		"bigip_bigip_ltm_pool_attachment":             26,
		"bigip_bigip_ltm_profile_client_ssl":          27,
		"bigip_bigip_ltm_profile_fasthttp":            28,
		"bigip_bigip_ltm_profile_fastl4":              29,
		"bigip_bigip_ltm_profile_ftp":                 30,
		"bigip_bigip_ltm_profile_http":                31,
		"bigip_bigip_ltm_profile_http2":               32,
		"bigip_bigip_ltm_profile_httpcompress":        33,
		"bigip_bigip_ltm_profile_oneconnect":          34,
		"bigip_bigip_ltm_profile_server_ssl":          35,
		"bigip_bigip_ltm_profile_tcp":                 36,
		"bigip_bigip_ltm_snat":                        37,
		"bigip_bigip_ltm_snatpool":                    38,
		"bigip_bigip_ltm_virtual_address":             39,
		"bigip_bigip_ltm_virtual_server":              40,
		"bigip_bigip_net_ike_peer":                    41,
		"bigip_bigip_net_route":                       42,
		"bigip_bigip_net_selfip":                      43,
		"bigip_bigip_net_tunnel":                      44,
		"bigip_bigip_net_vlan":                        45,
		"bigip_bigip_ssl_certificate":                 46,
		"bigip_bigip_ssl_key":                         47,
		"bigip_bigip_sys_dns":                         48,
		"bigip_bigip_sys_iapp":                        49,
		"bigip_bigip_sys_ntp":                         50,
		"bigip_bigip_sys_provision":                   51,
		"bigip_bigip_sys_snmp":                        52,
		"bigip_bigip_sys_snmp_traps":                  53,
		"bigip_bigip_traffic_selector":                54,
		"bigip_bigip_vcmp_guest":                      55,
		"bigip_bigip_waf_policy":                      56,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
