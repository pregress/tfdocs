package pagerduty

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_addon",
			Category:         "Resources",
			ShortDescription: `Creates and manages an add-on in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"addon",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the add-on.`,
				},
				resource.Attribute{
					Name:        "src",
					Description: `(Required) The source URL to display in a frame in the PagerDuty UI. ` + "`" + `HTTPS` + "`" + ` is required. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on. ## Import Add-ons can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_addon.example P3DH5M6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on. ## Import Add-ons can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_addon.example P3DH5M6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_business_service",
			Category:         "Resources",
			ShortDescription: `Creates and manages a business service in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"business",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the business service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the service. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "point_of_contact",
					Description: `(Optional) The owner of the business service.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) ID of the team that owns the business service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_escalation_policy",
			Category:         "Resources",
			ShortDescription: `Creates and manages an escalation policy in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"escalation",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the escalation policy.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) Team associated with the policy (Only 1 team can be assigned to an Escalation Policy). Account must have the ` + "`" + `teams` + "`" + ` ability to use this parameter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the escalation policy. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "num_loops",
					Description: `(Optional) The number of times the escalation policy will repeat after reaching the end of its escalation.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) An Escalation rule block. Escalation rules documented below. Escalation rules (` + "`" + `rule` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "escalation_delay_in_minutes",
					Description: `(Required) The number of minutes before an unacknowledged incident escalates away from this rule.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) A target block. Target blocks documented below. Targets (` + "`" + `target` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Can be ` + "`" + `user_reference` + "`" + ` or ` + "`" + `schedule_reference` + "`" + `. Defaults to ` + "`" + `user_reference` + "`" + `. For multiple users as example, repeat the target.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) A target ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Escalation policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_escalation_policy.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Escalation policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_escalation_policy.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_event_rule",
			Category:         "Resources",
			ShortDescription: `Creates and manages an event rule in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"event",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action_json",
					Description: `(Required) A list of one or more actions for each rule. Each action within the list is itself a list.`,
				},
				resource.Attribute{
					Name:        "condition_json",
					Description: `(Required) Contains a list of conditions. The first field in the list is ` + "`" + `and` + "`" + ` or ` + "`" + `or` + "`" + `, followed by a list of operators and values.`,
				},
				resource.Attribute{
					Name:        "advanced_condition_json",
					Description: `(Optional) Contains a list of specific conditions including ` + "`" + `active-between` + "`" + `,` + "`" + `scheduled-weekly` + "`" + `, and ` + "`" + `frequency-over` + "`" + `. The first element in the list is the label for the condition, followed by a list of values for the specific condition. For more details on these conditions see [Advanced Condition](https://developer.pagerduty.com/docs/rest-api-v2/global-event-rules-api/#advanced-condition-parameter) in the PagerDuty API documentation.`,
				},
				resource.Attribute{
					Name:        "depends_on",
					Description: `(Optional) A [Terraform meta-parameter](https://www.terraform.io/docs/configuration-0-11/resources.html#depends_on) that ensures that the ` + "`" + `event_rule` + "`" + ` specified is created before the current rule. This is important because Event Rules in PagerDuty are executed in order. ` + "`" + `depends_on` + "`" + ` ensures that the rules are created in the order specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the event rule.`,
				},
				resource.Attribute{
					Name:        "catch_all",
					Description: `A boolean that indicates whether the rule is a catch-all for the account. This field is read-only through the PagerDuty API. ## Import Event rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_event_rule.main 19acac92-027a-4ea0-b06c-bbf516519601 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the event rule.`,
				},
				resource.Attribute{
					Name:        "catch_all",
					Description: `A boolean that indicates whether the rule is a catch-all for the account. This field is read-only through the PagerDuty API. ## Import Event rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_event_rule.main 19acac92-027a-4ea0-b06c-bbf516519601 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_extension",
			Category:         "Resources",
			ShortDescription: `Creates and manages a service extension in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"extension",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service extension.`,
				},
				resource.Attribute{
					Name:        "endpoint_url",
					Description: `(Required|Optional) The url of the extension.`,
				},
				resource.Attribute{
					Name:        "extension_schema",
					Description: `(Required) This is the schema for this extension.`,
				},
				resource.Attribute{
					Name:        "extension_objects",
					Description: `(Required) This is the objects for which the extension applies (An array of service ids).`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The configuration of the service extension as string containing plain JSON-encoded data.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the extension.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app ## Import Extensions can be imported using the id.e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_extension.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the extension.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app ## Import Extensions can be imported using the id.e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_extension.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_extension_servicenow",
			Category:         "Resources",
			ShortDescription: `Creates and manages a ServiceNow service extension in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"extension",
				"servicenow",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service extension.`,
				},
				resource.Attribute{
					Name:        "extension_schema",
					Description: `(Required) This is the schema for this extension.`,
				},
				resource.Attribute{
					Name:        "extension_objects",
					Description: `(Required) This is the objects for which the extension applies (An array of service ids).`,
				},
				resource.Attribute{
					Name:        "snow_user",
					Description: `(Required) The ServiceNow username.`,
				},
				resource.Attribute{
					Name:        "snow_password",
					Description: `(Required) The ServiceNow password.`,
				},
				resource.Attribute{
					Name:        "sync_options",
					Description: `(Required) The ServiceNow sync option.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Target Webhook URL.`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The ServiceNow task type, typically ` + "`" + `incident` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "referer",
					Description: `(Required) The ServiceNow referer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the extension.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app. ## Import Extensions can be imported using the id.e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_extension_servicenow.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the extension.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app. ## Import Extensions can be imported using the id.e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_extension_servicenow.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_maintenance_window",
			Category:         "Resources",
			ShortDescription: `Creates and manages a maintenance window in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"maintenance",
				"window",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the maintenance window. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window. ## Import Maintenance windows can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_maintenance_window.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window. ## Import Maintenance windows can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_maintenance_window.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_response_play",
			Category:         "Resources",
			ShortDescription: `Creates and manages a response play in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"response",
				"play",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the response play.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Required) The email of the user attributed to the request. Needs to be a valid email address of a user in the PagerDuty account.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the response play. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) A string that determines the schema of the object. If not set, the default value is "response_play".`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) The ID of the team associated with the response play.`,
				},
				resource.Attribute{
					Name:        "subscriber",
					Description: `(Required) A user and/or team to be added as a subscriber to any incident on which this response play is run. There can be multiple subscribers defined on a single response play.`,
				},
				resource.Attribute{
					Name:        "subscribers_message",
					Description: `(Optional) The content of the notification that will be sent to all incident subscribers upon the running of this response play. Note that this includes any users who may have already been subscribed to the incident prior to the running of this response play. If empty, no notifications will be sent.`,
				},
				resource.Attribute{
					Name:        "responder",
					Description: `(Required) A user and/or escalation policy to be requested as a responder to any incident on which this response play is run. There can be multiple responders defined on a single response play.`,
				},
				resource.Attribute{
					Name:        "responders_message",
					Description: `(Optional) The message body of the notification that will be sent to this response play's set of responders. If empty, a default response request notification will be sent.`,
				},
				resource.Attribute{
					Name:        "runnability",
					Description: `(Optional) String representing how this response play is allowed to be run. Valid options are:`,
				},
				resource.Attribute{
					Name:        "conference_number",
					Description: `(Optional) The telephone number that will be set as the conference number for any incident on which this response play is run.`,
				},
				resource.Attribute{
					Name:        "conference_url",
					Description: `(Optional) The URL that will be set as the conference URL for any incident on which this response play is run. ### Responders (` + "`" + `responder` + "`" + `) can have two different objects and supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user defined as the responder`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Should be set as ` + "`" + `user_reference` + "`" + ` for user responders. ` + "`" + `escalation_policy` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user defined as the responder`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Should be set as ` + "`" + `escalation_policy` + "`" + ` for escalation policy responders.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the escalation policy`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of escalation policy`,
				},
				resource.Attribute{
					Name:        "num_loops",
					Description: `The number of times the escalation policy will repeat after reaching the end of its escalation.`,
				},
				resource.Attribute{
					Name:        "on_call_handoff_notifications",
					Description: `Determines how on call handoff notifications will be sent for users on the escalation policy. Defaults to "if_has_services". Could be "if_has_services", "always`,
				},
				resource.Attribute{
					Name:        "escalation_rule",
					Description: `The escalation rules`,
				},
				resource.Attribute{
					Name:        "escalation_delay_in_minutes",
					Description: `The number of minutes before an unacknowledged incident escalates away from this rule.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The targets an incident should be assigned to upon reaching this rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of object of the target. Supported types are ` + "`" + `user_reference` + "`" + `, ` + "`" + `schedule_reference` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `There can be multiple services associated with a policy.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) Teams associated with the policy. Account must have the ` + "`" + `teams` + "`" + ` ability to use this parameter. There can be multiple teams associated with a policy. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the response play. ## Import Response Plays can be imported using the ` + "`" + `id.from(email)` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_response_play.main 16208303-022b-f745-f2f5-560e537a2a74.user@email.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the response play. ## Import Response Plays can be imported using the ` + "`" + `id.from(email)` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_response_play.main 16208303-022b-f745-f2f5-560e537a2a74.user@email.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_ruleset",
			Category:         "Resources",
			ShortDescription: `Creates and manages an ruleset in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"ruleset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ruleset.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) Reference to the team that owns the ruleset. If none is specified, only admins have access. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ruleset.`,
				},
				resource.Attribute{
					Name:        "routing_keys",
					Description: `Routing keys routed to this ruleset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of ruleset. Currently, only sets to ` + "`" + `global` + "`" + `. ## Import Rulesets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_ruleset.main 19acac92-027a-4ea0-b06c-bbf516519601 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ruleset.`,
				},
				resource.Attribute{
					Name:        "routing_keys",
					Description: `Routing keys routed to this ruleset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of ruleset. Currently, only sets to ` + "`" + `global` + "`" + `. ## Import Rulesets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_ruleset.main 19acac92-027a-4ea0-b06c-bbf516519601 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_ruleset_rule",
			Category:         "Resources",
			ShortDescription: `Creates and manages a ruleset rule in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"ruleset",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ruleset",
					Description: `(Required) The ID of the ruleset that the rule belongs to.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) Conditions evaluated to check if an event matches this event rule. Is always empty for the catch-all rule, though.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) Position/index of the rule within the ruleset.`,
				},
				resource.Attribute{
					Name:        "catch_all",
					Description: `(Optional) Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates whether the rule is disabled and would therefore not be evaluated.`,
				},
				resource.Attribute{
					Name:        "time_frame",
					Description: `(Optional) Settings for [scheduling the rule](https://support.pagerduty.com/docs/rulesets#section-scheduled-event-rules).`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Optional) Actions to apply to an event if the conditions match.`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) Populate variables from event payloads and use those variables in other event actions.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator to combine sub-conditions. Can be ` + "`" + `and` + "`" + ` or ` + "`" + `or` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subconditions",
					Description: `List of sub-conditions that define the condition. ### Sub-Conditions (` + "`" + `subconditions` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Type of operator to apply to the sub-condition. Can be ` + "`" + `exists` + "`" + `,` + "`" + `nexists` + "`" + `,` + "`" + `equals` + "`" + `,` + "`" + `nequals` + "`" + `,` + "`" + `contains` + "`" + `,` + "`" + `ncontains` + "`" + `,` + "`" + `matches` + "`" + `, or ` + "`" + `nmatches` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parameter",
					Description: `Parameter for the sub-condition. It requires both a ` + "`" + `path` + "`" + ` and ` + "`" + `value` + "`" + ` to be set. ### Action (` + "`" + `actions` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Field where the data is being copied from. Must be a [PagerDuty Common Event Format (PD-CEF)](https://support.pagerduty.com/docs/pd-cef) field.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Field where the data is being copied to. Must be a [PagerDuty Common Event Format (PD-CEF)](https://support.pagerduty.com/docs/pd-cef) field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `The conditions that need to be met for the extraction to happen. Must use valid [RE2 regular expression syntax](https://github.com/google/re2/wiki/Syntax).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `A customized field message. This can also include variables extracted from the payload by using string interpolation.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Field where the data is being copied to. Must be a [PagerDuty Common Event Format (PD-CEF)](https://support.pagerduty.com/docs/pd-cef) field.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.`,
				},
				resource.Attribute{
					Name:        "weekdays",
					Description: `An integer array representing which days during the week the rule executes. For example ` + "`" + `weekdays = [1,3,7]` + "`" + ` would execute on Monday, Wednesday and Sunday.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `[The name of the timezone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for the given schedule, which will be used to determine UTC offset including adjustment for daylight saving time. For example: ` + "`" + `timezone = "America/Toronto"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `A Unix timestamp in milliseconds which is combined with the ` + "`" + `timezone` + "`" + ` to determine the time this rule will start on each specified ` + "`" + `weekday` + "`" + `. Note that the _date_ of the timestamp you specify does`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Length of time the schedule will be active in milliseconds. For example ` + "`" + `duration = 2`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the rule.`,
				},
				resource.Attribute{
					Name:        "catch_all",
					Description: `Indicates whether the rule is the last rule of the ruleset that serves as a catch-all. It has limited functionality compared to other rules. ## Import Ruleset rules can be imported using the related ` + "`" + `ruleset` + "`" + ` ID and the ` + "`" + `ruleset_rule` + "`" + ` ID separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_ruleset_rule.main a19cdca1-3d5e-4b52-bfea-8c8de04da243.19acac92-027a-4ea0-b06c-bbf516519601 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the rule.`,
				},
				resource.Attribute{
					Name:        "catch_all",
					Description: `Indicates whether the rule is the last rule of the ruleset that serves as a catch-all. It has limited functionality compared to other rules. ## Import Ruleset rules can be imported using the related ` + "`" + `ruleset` + "`" + ` ID and the ` + "`" + `ruleset_rule` + "`" + ` ID separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_ruleset_rule.main a19cdca1-3d5e-4b52-bfea-8c8de04da243.19acac92-027a-4ea0-b06c-bbf516519601 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_schedule",
			Category:         "Resources",
			ShortDescription: `Creates and manages a schedule in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the schedule.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Required) The time zone of the schedule (e.g. ` + "`" + `Europe/Berlin` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the schedule.`,
				},
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) A schedule layer block. Schedule layers documented below.`,
				},
				resource.Attribute{
					Name:        "overflow",
					Description: `(Optional) Any on-call schedule entries that pass the date range bounds will be truncated at the bounds, unless the parameter ` + "`" + `overflow` + "`" + ` is passed. For instance, if your schedule is a rotation that changes daily at midnight UTC, and your date range is from ` + "`" + `2011-06-01T10:00:00Z` + "`" + ` to ` + "`" + `2011-06-01T14:00:00Z` + "`" + `: If you don't pass the overflow=true parameter, you will get one schedule entry returned with a start of ` + "`" + `2011-06-01T10:00:00Z` + "`" + ` and end of ` + "`" + `2011-06-01T14:00:00Z` + "`" + `. If you do pass the ` + "`" + `overflow` + "`" + ` parameter, you will get one schedule entry returned with a start of ` + "`" + `2011-06-01T00:00:00Z` + "`" + ` and end of ` + "`" + `2011-06-02T00:00:00Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) Teams associated with the schedule. Schedule layers (` + "`" + `layer` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the schedule layer.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The start time of the schedule layer.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) The end time of the schedule layer. If not specified, the layer does not end.`,
				},
				resource.Attribute{
					Name:        "rotation_virtual_start",
					Description: `(Required) The effective start time of the schedule layer. This can be before the start time of the schedule.`,
				},
				resource.Attribute{
					Name:        "rotation_turn_length_seconds",
					Description: `(Required) The duration of each on-call shift in ` + "`" + `seconds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) The ordered list of users on this layer. The position of the user on the list determines their order in the layer.`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Optional) A schedule layer restriction block. Restriction blocks documented below. Restriction blocks (` + "`" + `restriction` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be ` + "`" + `daily_restriction` + "`" + ` or ` + "`" + `weekly_restriction` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time_of_day",
					Description: `(Required) The start time in ` + "`" + `HH:mm:ss` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "duration_seconds",
					Description: `(Required) The duration of the restriction in ` + "`" + `seconds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_day_of_week",
					Description: `(Required for ` + "`" + `weekly_restriction` + "`" + `) Number of the day when restriction starts. From 1 to 7 where 1 is Monday and 7 is Sunday. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule. ## Import Schedules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_schedule.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule. ## Import Schedules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_schedule.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service",
			Category:         "Resources",
			ShortDescription: `Creates and manages a service in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the service. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "auto_resolve_timeout",
					Description: `(Optional) Time in seconds that an incident is automatically resolved if left open for that long. Disabled if set to the ` + "`" + `"null"` + "`" + ` string.`,
				},
				resource.Attribute{
					Name:        "acknowledgement_timeout",
					Description: `(Optional) Time in seconds that an incident changes to the Triggered State after being Acknowledged. Disabled if set to the ` + "`" + `"null"` + "`" + ` string. If not passed in, will default to '"1800"'.`,
				},
				resource.Attribute{
					Name:        "escalation_policy",
					Description: `(Required) The escalation policy used by this service.`,
				},
				resource.Attribute{
					Name:        "response_play",
					Description: `(Optional) The response play used by this service.`,
				},
				resource.Attribute{
					Name:        "alert_creation",
					Description: `(Optional) Must be one of two values. PagerDuty receives events from your monitoring systems and can then create incidents in different ways. Value "create_incidents" is default: events will create an incident that cannot be merged. Value "create_alerts_and_incidents" is the alternative: events will create an alert and then add it to a new incident, these incidents can be merged. This option is recommended.`,
				},
				resource.Attribute{
					Name:        "alert_grouping",
					Description: `(Optional) (Deprecated) Defines how alerts on this service will be automatically grouped into incidents. Note that the alert grouping features are available only on certain plans. If not set, each alert will create a separate incident; If value is set to ` + "`" + `time` + "`" + `: All alerts within a specified duration will be grouped into the same incident. This duration is set in the ` + "`" + `alert_grouping_timeout` + "`" + ` setting (described below). Available on Standard, Enterprise, and Event Intelligence plans; If value is set to ` + "`" + `intelligent` + "`" + ` - Alerts will be intelligently grouped based on a machine learning model that looks at the alert summary, timing, and the history of grouped alerts. Available on Enterprise and Event Intelligence plan. This field is deprecated, use ` + "`" + `alert_grouping_parameters.type` + "`" + ` instead,`,
				},
				resource.Attribute{
					Name:        "alert_grouping_timeout",
					Description: `(Optional) (Deprecated) The duration in minutes within which to automatically group incoming alerts. This setting applies only when ` + "`" + `alert_grouping` + "`" + ` is set to ` + "`" + `time` + "`" + `. To continue grouping alerts until the incident is resolved, set this value to ` + "`" + `0` + "`" + `. This field is deprecated, use ` + "`" + `alert_grouping_parameters.config.timeout` + "`" + ` instead,`,
				},
				resource.Attribute{
					Name:        "alert_grouping_parameters",
					Description: `(Optional) Defines how alerts on this service will be automatically grouped into incidents. Note that the alert grouping features are available only on certain plans. If not set, each alert will create a separate incident.`,
				},
				resource.Attribute{
					Name:        "auto_pause_notifications_parameters",
					Description: `(Optional) Defines how alerts on this service are automatically suspended for a period of time before triggering, when identified as likely being transient. Note that automatically pausing notifications is only available on certain plans as mentioned [here](https://support.pagerduty.com/docs/auto-pause-incident-notifications). The ` + "`" + `alert_grouping_parameters` + "`" + ` block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The duration in minutes within which to automatically group incoming alerts. This setting applies only when ` + "`" + `type` + "`" + ` is set to ` + "`" + `time` + "`" + `. To continue grouping alerts until the incident is resolved, set this value to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) One of ` + "`" + `any` + "`" + ` or ` + "`" + `all` + "`" + `. This setting applies only when ` + "`" + `type` + "`" + ` is set to ` + "`" + `content_based` + "`" + `. Group alerts based on one or all of ` + "`" + `fields` + "`" + ` value(s).`,
				},
				resource.Attribute{
					Name:        "fields",
					Description: `(Optional) Alerts will be grouped together if the content of these fields match. This setting applies only when ` + "`" + `type` + "`" + ` is set to ` + "`" + `content_based` + "`" + `. The ` + "`" + `auto_pause_notifications_parameters` + "`" + ` block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of incident urgency: ` + "`" + `constant` + "`" + ` or ` + "`" + `use_support_hours` + "`" + ` (when depending on specific support hours; see ` + "`" + `support_hours` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "urgency",
					Description: `The urgency: ` + "`" + `low` + "`" + ` Notify responders (does not escalate), ` + "`" + `high` + "`" + ` (follows escalation rules) or ` + "`" + `severity_based` + "`" + ` Set's the urgency of the incident based on the severity set by the triggering monitoring tool.`,
				},
				resource.Attribute{
					Name:        "during_support_hours",
					Description: `(Optional) Incidents' urgency during support hours.`,
				},
				resource.Attribute{
					Name:        "outside_support_hours",
					Description: `(Optional) Incidents' urgency outside support hours. When using ` + "`" + `type = "use_support_hours"` + "`" + ` in ` + "`" + `incident_urgency_rule` + "`" + ` you must specify exactly one (otherwise optional) ` + "`" + `support_hours` + "`" + ` block. Your PagerDuty account must have the ` + "`" + `service_support_hours` + "`" + ` ability to assign support hours. The block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of support hours. Can be ` + "`" + `fixed_time_per_day` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `The time zone for the support hours.`,
				},
				resource.Attribute{
					Name:        "days_of_week",
					Description: `Array of days of week as integers. ` + "`" + `1` + "`" + ` to ` + "`" + `7` + "`" + `, ` + "`" + `1` + "`" + ` being Monday and ` + "`" + `7` + "`" + ` being Sunday.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The support hours' starting time of day.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The support hours' ending time of day. A ` + "`" + `scheduled_actions` + "`" + ` block is required when using ` + "`" + `type = "use_support_hours"` + "`" + ` in ` + "`" + `incident_urgency_rule` + "`" + `. The block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of scheduled action. Currently, this must be set to ` + "`" + `urgency_change` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "to_urgency",
					Description: `The urgency to change to: ` + "`" + `low` + "`" + ` (does not escalate), or ` + "`" + `high` + "`" + ` (follows escalation rules).`,
				},
				resource.Attribute{
					Name:        "at",
					Description: `A block representing when the scheduled action will occur. The ` + "`" + `at` + "`" + ` block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of time specification. Currently, this must be set to ` + "`" + `named_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Designates either the start or the end of the scheduled action. Can be ` + "`" + `support_hours_start` + "`" + ` or ` + "`" + `support_hours_end` + "`" + `. Note that it is currently only possible to define the scheduled action when urgency is set to ` + "`" + `high` + "`" + ` for ` + "`" + `during_support_hours` + "`" + ` and to ` + "`" + `low` + "`" + ` for ` + "`" + `outside_support_hours` + "`" + ` in ` + "`" + `incident_urgency_rule` + "`" + `. Below is an example for a ` + "`" + `pagerduty_service` + "`" + ` resource with ` + "`" + `incident_urgency_rules` + "`" + ` with ` + "`" + `type = "use_support_hours"` + "`" + `, ` + "`" + `support_hours` + "`" + ` and a default ` + "`" + `scheduled_action` + "`" + ` as well. ` + "`" + `` + "`" + `` + "`" + `hcl resource "pagerduty_service" "foo" { name = "bar" description = "bar bar bar" auto_resolve_timeout = 3600 acknowledgement_timeout = 3600 escalation_policy = pagerduty_escalation_policy.foo.id incident_urgency_rule { type = "use_support_hours" during_support_hours { type = "constant" urgency = "high" } outside_support_hours { type = "constant" urgency = "low" } } support_hours { type = "fixed_time_per_day" time_zone = "America/Lima" start_time = "09:00:00" end_time = "17:00:00" days_of_week = [1, 2, 3, 4, 5] } scheduled_actions { type = "urgency_change" to_urgency = "high" at { type = "named_time" name = "support_hours_start" } } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of object. The value returned will be ` + "`" + `service` + "`" + `. Can be used for passing to a service dependency. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_service.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of object. The value returned will be ` + "`" + `service` + "`" + `. Can be used for passing to a service dependency. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_service.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service_dependency",
			Category:         "Resources",
			ShortDescription: `Creates and manages a business service dependency in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"dependency",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dependency",
					Description: `(Required) The relationship between the ` + "`" + `supporting_service` + "`" + ` and ` + "`" + `dependent_service` + "`" + `. One and only one dependency block must be defined.`,
				},
				resource.Attribute{
					Name:        "supporting_service",
					Description: `(Required) The service that supports the dependent service. Dependency supporting service documented below.`,
				},
				resource.Attribute{
					Name:        "dependent_service",
					Description: `(Required) The service that dependents on the supporting service. Dependency dependent service documented below. Dependency supporting and dependent service supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the service dependency.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be ` + "`" + `business_service` + "`" + `, ` + "`" + `service` + "`" + `, ` + "`" + `business_service_reference` + "`" + ` or ` + "`" + `technical_service_reference` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service dependency.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service dependency.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service_integration",
			Category:         "Resources",
			ShortDescription: `Creates and manages a service integration in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) The ID of the service the integration should belong to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The service type. Can be: ` + "`" + `aws_cloudwatch_inbound_integration` + "`" + `, ` + "`" + `cloudkick_inbound_integration` + "`" + `, ` + "`" + `event_transformer_api_inbound_integration` + "`" + `, ` + "`" + `events_api_v2_inbound_integration` + "`" + ` (requires service ` + "`" + `alert_creation` + "`" + ` to be ` + "`" + `create_alerts_and_incidents` + "`" + `), ` + "`" + `generic_email_inbound_integration` + "`" + `, ` + "`" + `generic_events_api_inbound_integration` + "`" + `, ` + "`" + `keynote_inbound_integration` + "`" + `, ` + "`" + `nagios_inbound_integration` + "`" + `, ` + "`" + `pingdom_inbound_integration` + "`" + `or ` + "`" + `sql_monitor_inbound_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `(Optional) The ID of the vendor the integration should integrate with (e.g. Datadog or Amazon Cloudwatch).`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Optional) This is the unique key used to route events to this integration when received via the PagerDuty Events API.`,
				},
				resource.Attribute{
					Name:        "integration_email",
					Description: `(Optional) This is the unique fully-qualified email address used for routing emails to this integration for processing.`,
				},
				resource.Attribute{
					Name:        "email_incident_creation",
					Description: `(Optional) Behaviour of Email Management feature ([explained in PD docs](https://support.pagerduty.com/docs/email-management-filters-and-rules#control-when-a-new-incident-or-alert-is-triggered)). Can be ` + "`" + `on_new_email` + "`" + `, ` + "`" + `on_new_email_subject` + "`" + `, ` + "`" + `only_if_no_open_incidents` + "`" + ` or ` + "`" + `use_rules` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "email_filter_mode",
					Description: `(Optional) Mode of Emails Filters feature ([explained in PD docs](https://support.pagerduty.com/docs/email-management-filters-and-rules#configure-a-regex-filter)). Can be ` + "`" + `all-email` + "`" + `, ` + "`" + `or-rules-email` + "`" + ` or ` + "`" + `and-rules-email` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "email_parsing_fallback",
					Description: `(Optional) Can be ` + "`" + `open_new_incident` + "`" + ` or ` + "`" + `discard` + "`" + `. Email filters (` + "`" + `email_filter` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "body_mode",
					Description: `(Required) Can be ` + "`" + `always` + "`" + ` or ` + "`" + `match` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "body_regex",
					Description: `(Optional) Should be a valid regex or ` + "`" + `null` + "`" + ``,
				},
				resource.Attribute{
					Name:        "from_email_mode",
					Description: `(Required) Can be ` + "`" + `always` + "`" + ` or ` + "`" + `match` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "from_email_regex",
					Description: `(Optional) Should be a valid regex or ` + "`" + `null` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subject_mode",
					Description: `(Required) Can be ` + "`" + `always` + "`" + ` or ` + "`" + `match` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subject_regex",
					Description: `(Optional) Should be a valid regex or ` + "`" + `null` + "`" + ` Email parsers (` + "`" + `email_parser` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Can be ` + "`" + `resolve` + "`" + ` or ` + "`" + `trigger` + "`" + `. Match predicate (` + "`" + `match_predicate` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be ` + "`" + `any` + "`" + ` or ` + "`" + `all` + "`" + `. Predicates (` + "`" + `predicate` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be ` + "`" + `contains` + "`" + `, ` + "`" + `exactly` + "`" + `, ` + "`" + `regex` + "`" + ` or ` + "`" + `not` + "`" + `. If type is ` + "`" + `not` + "`" + ` predicate should contain child predicate with all parameters.`,
				},
				resource.Attribute{
					Name:        "matcher",
					Description: `(Optional) Predicate value or valid regex.`,
				},
				resource.Attribute{
					Name:        "part",
					Description: `(Optional) Can be ` + "`" + `subject` + "`" + `, ` + "`" + `body` + "`" + ` or ` + "`" + `from_addresses` + "`" + `. Value extractors (` + "`" + `value_extractor` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be ` + "`" + `between` + "`" + `, ` + "`" + `entire` + "`" + ` or ` + "`" + `regex` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "part",
					Description: `(Required) Can be ` + "`" + `subject` + "`" + ` or ` + "`" + `body` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value_name",
					Description: `(Required) First value extractor should have name ` + "`" + `incident_key` + "`" + ` other value extractors should contain custom names.`,
				},
				resource.Attribute{
					Name:        "ends_before",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "starts_after",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) If ` + "`" + `type` + "`" + ` has value ` + "`" + `regex` + "`" + ` this value should contain valid regex.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service integration.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `This is the unique key used to route events to this integration when received via the PagerDuty Events API.`,
				},
				resource.Attribute{
					Name:        "integration_email",
					Description: `This is the unique fully-qualified email address used for routing emails to this integration for processing.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app. To configure an event, please use the ` + "`" + `integration_key` + "`" + ` in the following interpolation: ` + "`" + `` + "`" + `` + "`" + `hcl https://events.pagerduty.com/integration/${pagerduty_service_integration.slack.integration_key}/enqueue ` + "`" + `` + "`" + `` + "`" + ` ## Import Services can be imported using their related ` + "`" + `service` + "`" + ` id and service integration ` + "`" + `id` + "`" + ` separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_service_integration.main PLSSSSS.PLIIIII ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service integration.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `This is the unique key used to route events to this integration when received via the PagerDuty Events API.`,
				},
				resource.Attribute{
					Name:        "integration_email",
					Description: `This is the unique fully-qualified email address used for routing emails to this integration for processing.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app. To configure an event, please use the ` + "`" + `integration_key` + "`" + ` in the following interpolation: ` + "`" + `` + "`" + `` + "`" + `hcl https://events.pagerduty.com/integration/${pagerduty_service_integration.slack.integration_key}/enqueue ` + "`" + `` + "`" + `` + "`" + ` ## Import Services can be imported using their related ` + "`" + `service` + "`" + ` id and service integration ` + "`" + `id` + "`" + ` separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_service_integration.main PLSSSSS.PLIIIII ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_slack_connection",
			Category:         "Resources",
			ShortDescription: `Creates and manages a slack connection in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"slack",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_id",
					Description: `(Required) The ID of the source in PagerDuty. Valid sources are services or teams.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) The type of the source. Either ` + "`" + `team_reference` + "`" + ` or ` + "`" + `service_reference` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Required) The ID of the connected Slack workspace. Can also be defined by the ` + "`" + `SLACK_CONNECTION_WORKSPACE_ID` + "`" + ` environment variable.`,
				},
				resource.Attribute{
					Name:        "channel_id",
					Description: `(Required) The ID of a Slack channel in the workspace.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Configuration options for the Slack connection that provide options to filter events.`,
				},
				resource.Attribute{
					Name:        "notification_type",
					Description: `(Required) Type of notification. Either ` + "`" + `responder` + "`" + ` or ` + "`" + `stakeholder` + "`" + `. ### Connection Config (` + "`" + `config` + "`" + `) Supports the following:`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Required) A list of strings to filter events by PagerDuty event type. ` + "`" + `"incident.triggered"` + "`" + ` is required. The follow event types are also possible: - ` + "`" + `incident.acknowledged` + "`" + ` - ` + "`" + `incident.escalated` + "`" + ` - ` + "`" + `incident.resolved` + "`" + ` - ` + "`" + `incident.reassigned` + "`" + ` - ` + "`" + `incident.annotated` + "`" + ` - ` + "`" + `incident.unacknowledged` + "`" + ` - ` + "`" + `incident.delegated` + "`" + ` - ` + "`" + `incident.priority_updated` + "`" + ` - ` + "`" + `incident.responder.added` + "`" + ` - ` + "`" + `incident.responder.replied` + "`" + ` - ` + "`" + `incident.status_update_published` + "`" + ` - ` + "`" + `incident.reopened` + "`" + ``,
				},
				resource.Attribute{
					Name:        "priorities",
					Description: `(Optional) Allows you to filter events by priority. Needs to be an array of PagerDuty priority IDs. Available through [pagerduty_priority](https://registry.terraform.io/providers/PagerDuty/pagerduty/latest/docs/data-sources/priority) data source. - When omitted or set to an empty array (` + "`" + `[]` + "`" + `) in the configuration for a Slack Connection, its default behaviour is to set ` + "`" + `priorities` + "`" + ` to ` + "`" + `No Priority` + "`" + ` value. - When set to ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "urgency",
					Description: `(Optional) Allows you to filter events by urgency. Either ` + "`" + `high` + "`" + ` or ` + "`" + `low` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the slack connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the slack connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_tag",
			Category:         "Resources",
			ShortDescription: `Creates and manages a tag in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label of the tag. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the tag.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app. ## Import Tags can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_tag.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the tag.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app. ## Import Tags can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_tag.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_tag_assignment",
			Category:         "Resources",
			ShortDescription: `Creates and manages a tag assignment in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"tag",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tag_id",
					Description: `(Required) The ID of the tag.`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) Type of entity in the tag assignment. Possible values can be ` + "`" + `users` + "`" + `, ` + "`" + `teams` + "`" + `, and ` + "`" + `escalation_policies` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Required) The ID of the entity. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the tag assignment. ## Import Tag assignments can be imported using the ` + "`" + `id` + "`" + ` which is constructed by taking the ` + "`" + `entity` + "`" + ` Type, ` + "`" + `entity` + "`" + ` ID and the ` + "`" + `tag` + "`" + ` ID separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_tag_assignment.main users.P7HHMVK.PYC7IQQ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the tag assignment. ## Import Tag assignments can be imported using the ` + "`" + `id` + "`" + ` which is constructed by taking the ` + "`" + `entity` + "`" + ` Type, ` + "`" + `entity` + "`" + ` ID and the ` + "`" + `tag` + "`" + ` ID separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_tag_assignment.main users.P7HHMVK.PYC7IQQ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_team",
			Category:         "Resources",
			ShortDescription: `Creates and manages a team in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the team. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Optional) ID of the parent team. This is available to accounts with the Team Hierarchy feature enabled. Please contact your account manager for more information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_team.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_team.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_team_membership",
			Category:         "Resources",
			ShortDescription: `Creates and manages a team membership in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user to add to the team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The ID of the team in which the user will belong.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user belonging to the team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The team ID the user belongs to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user belonging to the team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The team ID the user belongs to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages a user in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The user's email address.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) The schedule color for the user. Valid options are purple, red, green, blue, teal, orange, brown, turquoise, dark-slate-blue, cayenne, orange-red, dark-orchid, dark-slate-grey, lime, dark-magenta, lime-green, midnight-blue, deep-pink, dark-green, dark-orange, dark-cyan, darkolive-green, dark-slate-gray, grey20, firebrick, maroon, crimson, dark-red, dark-goldenrod, chocolate, medium-violet-red, sea-green, olivedrab, forest-green, dark-olive-green, blue-violet, royal-blue, indigo, slate-blue, saddle-brown, or steel-blue.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The user role. Can be ` + "`" + `admin` + "`" + `, ` + "`" + `limited_user` + "`" + `, ` + "`" + `observer` + "`" + `, ` + "`" + `owner` + "`" + `, ` + "`" + `read_only_user` + "`" + `, ` + "`" + `read_only_limited_user` + "`" + `, ` + "`" + `restricted_access` + "`" + `, or ` + "`" + `user` + "`" + `. Notes:`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `(Optional) The user's title.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) The time zone of the user. Default is account default timezone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the user. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "license",
					Description: `(Optional) The license id assigned to the user. If provided the user's role must exist in the assigned license's ` + "`" + `valid_roles` + "`" + ` list. To reference purchased licenses' ids see data source ` + "`" + `pagerduty_licenses` + "`" + ` [data source][1]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The URL of the user's avatar.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `The timezone of the user.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app`,
				},
				resource.Attribute{
					Name:        "invitation_sent",
					Description: `If true, the user has an outstanding invitation. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ` [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzNA-create-a-user [2]: https://registry.terraform.io/providers/PagerDuty/pagerduty/latest/docs/data-sources/pagerduty_license`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The URL of the user's avatar.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `The timezone of the user.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app`,
				},
				resource.Attribute{
					Name:        "invitation_sent",
					Description: `If true, the user has an outstanding invitation. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ` [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzNA-create-a-user [2]: https://registry.terraform.io/providers/PagerDuty/pagerduty/latest/docs/data-sources/pagerduty_license`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user_contact_method",
			Category:         "Resources",
			ShortDescription: `Creates and manages contact methods for a user in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"contact",
				"method",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The contact method type. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_short_email",
					Description: `(Optional) Send an abbreviated email message instead of the standard email output.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) The 1-to-3 digit country calling code. Required when using ` + "`" + `phone_contact_method` + "`" + ` or ` + "`" + `sms_contact_method` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label (e.g., "Work", "Mobile", etc.).`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The "address" to deliver to: ` + "`" + `email` + "`" + `, ` + "`" + `phone number` + "`" + `, etc., depending on the type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the contact method.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. ## Import Contact methods can be imported using the ` + "`" + `user_id` + "`" + ` and the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user_contact_method.main PLBP09X:PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the contact method.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. ## Import Contact methods can be imported using the ` + "`" + `user_id` + "`" + ` and the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user_contact_method.main PLBP09X:PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user_notification_rule",
			Category:         "Resources",
			ShortDescription: `Creates and manages notification rules for a user in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"notification",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user.`,
				},
				resource.Attribute{
					Name:        "start_delay_in_minutes",
					Description: `(Required) The delay before firing the rule, in minutes.`,
				},
				resource.Attribute{
					Name:        "urgency",
					Description: `(Required) Which incident urgency this rule is used for. Account must have the ` + "`" + `urgencies` + "`" + ` ability to have a low urgency notification rule. Can be ` + "`" + `high` + "`" + ` or ` + "`" + `low` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "contact_method",
					Description: `(Required) A contact method block, configured as a block described below. Contact methods (` + "`" + `contact_method` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the referenced contact method.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of contact method. Can be ` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + ` or ` + "`" + `sms_contact_method` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user notification rule. ## Import User notification rules can be imported using the ` + "`" + `user_id` + "`" + ` and the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user_notification_rule.main PXPGF42:PPSCXAN ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user notification rule. ## Import User notification rules can be imported using the ` + "`" + `user_id` + "`" + ` and the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user_notification_rule.main PXPGF42:PPSCXAN ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_webhook_subscription",
			Category:         "Resources",
			ShortDescription: `Creates and manages a webhook subscription in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type indicating the schema of the object. The provider sets this as ` + "`" + `webhook_subscription` + "`" + `, which is currently the only acceptable value.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) Determines whether the subscription will produce webhook events.`,
				},
				resource.Attribute{
					Name:        "delivery_method",
					Description: `(Required) The object describing where to send the webhooks.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A short description of the webhook subscription`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Required) A set of outbound event types the webhook will receive. The follow event types are possible:`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) determines which events will match and produce a webhook. There are currently three types of filters that can be applied to webhook subscriptions: ` + "`" + `service_reference` + "`" + `, ` + "`" + `team_reference` + "`" + ` and ` + "`" + `account_reference` + "`" + `. ### Webhook delivery method (` + "`" + `delivery_method` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "temporarily_disabled",
					Description: `(Required) Whether this webhook subscription is temporarily disabled. Becomes true if the delivery method URL is repeatedly rejected by the server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Indicates the type of the delivery method. Allowed and default value: ` + "`" + `http_delivery_method` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The destination URL for webhook delivery.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) The custom_header of a webhook subscription define any optional headers that will be passed along with the payload to the destination URL. ### Webhook filter (` + "`" + `filter` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the object being used as the filter. This field is required for all filter types except account_reference.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of object being used as the filter. Allowed values are ` + "`" + `account_reference` + "`" + `, ` + "`" + `service_reference` + "`" + `, and ` + "`" + `team_reference` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the slack connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the slack connection.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"pagerduty_addon":                  0,
		"pagerduty_business_service":       1,
		"pagerduty_escalation_policy":      2,
		"pagerduty_event_rule":             3,
		"pagerduty_extension":              4,
		"pagerduty_extension_servicenow":   5,
		"pagerduty_maintenance_window":     6,
		"pagerduty_response_play":          7,
		"pagerduty_ruleset":                8,
		"pagerduty_ruleset_rule":           9,
		"pagerduty_schedule":               10,
		"pagerduty_service":                11,
		"pagerduty_service_dependency":     12,
		"pagerduty_service_integration":    13,
		"pagerduty_slack_connection":       14,
		"pagerduty_tag":                    15,
		"pagerduty_tag_assignment":         16,
		"pagerduty_team":                   17,
		"pagerduty_team_membership":        18,
		"pagerduty_user":                   19,
		"pagerduty_user_contact_method":    20,
		"pagerduty_user_notification_rule": 21,
		"pagerduty_webhook_subscription":   22,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
