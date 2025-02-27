package gridscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gridscale_backupschedule",
			Category:         "Resources",
			ShortDescription: `Manages a storage backup schedule.`,
			Description:      ``,
			Keywords: []string{
				"backupschedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) UUID of the backup schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage that the backup schedule belongs to.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) The status of the schedule active or not.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `(Required) The date and time that the backup schedule will be run.`,
				},
				resource.Attribute{
					Name:        "keep_backups",
					Description: `(Required) The amount of Snapshots to keep before overwriting the last created Snapshot (>=1).`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `(Required) The interval at which the schedule will run (in minutes, >=60).`,
				},
				resource.Attribute{
					Name:        "backup_location_uuid",
					Description: `(Optional, ForceNew) UUID of the location where your backup is stored. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the backup schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the backup schedule.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_location_name",
					Description: `The human-readable name of backup location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime_computed",
					Description: `The date and time that the backup schedule will be run. This date and time is computed by gridscale's server.`,
				},
				resource.Attribute{
					Name:        "keep_backups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the backup schedule was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last backup schedule change.`,
				},
				resource.Attribute{
					Name:        "storage_backups",
					Description: `Related backups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the backup.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `UUID of the backup.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the backup was initially created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the backup schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the backup schedule.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_location_name",
					Description: `The human-readable name of backup location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime_computed",
					Description: `The date and time that the backup schedule will be run. This date and time is computed by gridscale's server.`,
				},
				resource.Attribute{
					Name:        "keep_backups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the backup schedule was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last backup schedule change.`,
				},
				resource.Attribute{
					Name:        "storage_backups",
					Description: `Related backups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the backup.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `UUID of the backup.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the backup was initially created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_object_storage_bucket",
			Category:         "Resources",
			ShortDescription: `Manages an object storage bucket in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required, Force New) Access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required, Force New) Secret key.`,
				},
				resource.Attribute{
					Name:        "s3_host",
					Description: `(Required, Force New) Host of the s3. Default: "gos3.io".`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required, Force New) Name of the bucket.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_filesystem",
			Category:         "Resources",
			ShortDescription: `Manage a Filesystem service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"filesystem",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The filesystem service release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available filesystem service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of filesystem service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "root_squash",
					Description: `(Optional) Map root user/group ownership to anon_uid/anon_gid.`,
				},
				resource.Attribute{
					Name:        "allowed_ip_ranges",
					Description: `(Optional) Allowed CIDR block or IP address in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "anon_uid",
					Description: `(Optional) Target user id when root squash is active.`,
				},
				resource.Attribute{
					Name:        "anon_gid",
					Description: `(Optional) Target group id when root squash is active. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "root_squash",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allowed_ip_ranges",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "anon_uid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "anon_gid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where the filesystem service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that filesystem service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_firewall",
			Category:         "Resources",
			ShortDescription: `Manages a firewall in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the firewall.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The information about networks which are related to this firewall.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the firewall.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the firewall.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the firewall.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the firewall.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The information about networks which are related to this firewall.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the firewall.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the firewall.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the firewall.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ipv4",
			Category:         "Resources",
			ShortDescription: `Manages an IPv4 address in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Defines the reverse DNS entry for the IP address (PTR Resource Record).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The network address and the subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the object is administratively blocked. If true, it can not be deleted by the user.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ipv6",
			Category:         "Resources",
			ShortDescription: `Manages an IPv6 address in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Defines the reverse DNS entry for the IP address (PTR Resource Record).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this resource is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The network address and the subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the object is administratively blocked. If true, it can not be deleted by the user.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_isoimage",
			Category:         "Resources",
			ShortDescription: `Manages an ISO Image in Gridscale.`,
			Description:      ``,
			Keywords: []string{
				"isoimage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Required) Contains the source URL of the ISO Image that it was originally fetched from.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `Contains the source URL of the ISO Image that it was originally fetched from.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The information about servers which are related to this ISO Image.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the server.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the server.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `True if the ISO Image is a boot device of this server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/ISO Image/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `Contains the source URL of the ISO Image that it was originally fetched from.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The information about servers which are related to this ISO Image.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the server.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the server.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `True if the ISO Image is a boot device of this server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/ISO Image/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_k8s",
			Category:         "Resources",
			ShortDescription: `Manages a k8s cluster in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"k8s",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "gsk_version",
					Description: `(Optional) The gridscale's Kubernetes version of this instance (e.g. "1.21.14-gs1"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available gridscale k8s version.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Optional) The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available releases.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `(Required) Node pool's specification.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Immutable) Name of the node pool.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of worker nodes.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Immutable) Cores per worker node.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Immutable) Memory per worker node (in GiB).`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Immutable) Storage per worker node (in GiB).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Immutable) Storage type (one of storage, storage_high, storage_insane).`,
				},
				resource.Attribute{
					Name:        "surge_node",
					Description: `Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the ` + "`" + `cluster_cidr` + "`" + ` in the k8s resource is empty). ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "45m" - 45 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "45m" - 45 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "45m" - 45 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gsk_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that k8s service uses. The ` + "`" + `service_template_uuid` + "`" + ` may not relate to ` + "`" + `release` + "`" + `, if ` + "`" + `service_template_uuid` + "`" + `/` + "`" + `release` + "`" + ` is updated outside of terraform (e.g. the k8s service is upgraded by gridscale staffs).`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The kubeconfig file content of the k8s cluster.`,
				},
				resource.Attribute{
					Name:        "k8s_private_network_uuid",
					Description: `Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "surge_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Manage a loadbalancer in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "redirect_http_to_https",
					Description: `(Required) Whether the load balancer is forced to redirect requests from HTTP to HTTPS.`,
				},
				resource.Attribute{
					Name:        "listen_ipv4_uuid",
					Description: `(Required) The UUID of the IPv4 address the load balancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "listen_ipv6_uuid",
					Description: `(Required) The UUID of the IPv6 address the load balancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) The algorithm used to process requests. Accepted values: roundrobin/leastconn.`,
				},
				resource.Attribute{
					Name:        "backend_server",
					Description: `(Required) The servers that the load balancer can communicate with.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A valid domain or an IP address of a server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The backend host weight. Default: 100.`,
				},
				resource.Attribute{
					Name:        "forwarding_rule",
					Description: `(Required) The forwarding rules of the load balancer.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this load balancer is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm used to process requests.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "redirect_http_to_https",
					Description: `Whether the Load balancer is forced to redirect requests from HTTP to HTTPS.`,
				},
				resource.Attribute{
					Name:        "listen_ipv4_uuid",
					Description: `The UUID of the IPv4 address the load balancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "listen_ipv6_uuid",
					Description: `The UUID of the IPv6 address the load balancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "backend_server",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "forwarding_rule",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "letsencrypt_ssl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_mariadb",
			Category:         "Resources",
			ShortDescription: `Manage a MariaDB service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"mariadb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The MariaDB release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available MariaDB service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of MariaDB service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mariadb_log_bin",
					Description: `(Optional) MariaDB parameter: Binary Logging. Default: false.`,
				},
				resource.Attribute{
					Name:        "mariadb_sql_mode",
					Description: `(Optional) MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".`,
				},
				resource.Attribute{
					Name:        "mariadb_server_id",
					Description: `(Optional) MariaDB parameter: Server Id. Default: 1.`,
				},
				resource.Attribute{
					Name:        "mariadb_query_cache",
					Description: `(Optional) MariaDB parameter: Enable query cache. Default: true.`,
				},
				resource.Attribute{
					Name:        "mariadb_binlog_format",
					Description: `(Optional) MariaDB parameter: Binary Logging Format. Default: "MIXED".`,
				},
				resource.Attribute{
					Name:        "mariadb_max_connections",
					Description: `(Optional) MariaDB parameter: Max Connections. Default: 4000.`,
				},
				resource.Attribute{
					Name:        "mariadb_query_cache_size",
					Description: `(Optional) MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.`,
				},
				resource.Attribute{
					Name:        "mariadb_default_time_zone",
					Description: `(Optional) MariaDB parameter: Server Timezone. Default: UTC.`,
				},
				resource.Attribute{
					Name:        "mariadb_query_cache_limit",
					Description: `(Optional) MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.`,
				},
				resource.Attribute{
					Name:        "mariadb_max_allowed_packet",
					Description: `(Optional) MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `(Optional) Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and ` + "`" + `max_core_count` + "`" + `. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_log_bin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_sql_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_server_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_query_cache",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_binlog_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_max_connections",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_query_cache_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_default_time_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_query_cache_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mariadb_max_allowed_packet",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the MariaDB instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the MariaDB instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this MariaDB service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that MariaDB service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_marketplace_application",
			Category:         "Resources",
			ShortDescription: `Manages marketplace applications in Gridscale.`,
			Description:      ``,
			Keywords: []string{
				"marketplace",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "object_storage_path",
					Description: `(Required) Path to the images for the application, must be in .gz format and started with s3//.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving".`,
				},
				resource.Attribute{
					Name:        "setup_cores",
					Description: `(Required) Number of server's cores.`,
				},
				resource.Attribute{
					Name:        "setup_memory",
					Description: `(Required) The capacity of server's memory in GB.`,
				},
				resource.Attribute{
					Name:        "setup_storage_capacity",
					Description: `(Required) The capacity of server's storage in GB.`,
				},
				resource.Attribute{
					Name:        "meta_license",
					Description: `(Optional) License number.`,
				},
				resource.Attribute{
					Name:        "meta_os",
					Description: `(Optional) Operating system.`,
				},
				resource.Attribute{
					Name:        "meta_components",
					Description: `(Optional) Components (e.g: MySql, Apache, etc.).`,
				},
				resource.Attribute{
					Name:        "meta_overview",
					Description: `(Optional) Describes the main function of the application.`,
				},
				resource.Attribute{
					Name:        "meta_hints",
					Description: `(Optional) Hints.`,
				},
				resource.Attribute{
					Name:        "meta_terms_of_use",
					Description: `(Optional) Terms of use.`,
				},
				resource.Attribute{
					Name:        "meta_icon",
					Description: `(Optional) base64 encoded image of the icon.`,
				},
				resource.Attribute{
					Name:        "meta_features",
					Description: `(Optional) List of functions.`,
				},
				resource.Attribute{
					Name:        "meta_author",
					Description: `(Optional) Author. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object_storage_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "setup_cores",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "setup_memory",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "setup_storage_capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_license",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_os",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_components",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_overview",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_hints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_terms_of_use",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_icon",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_features",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_author",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_advices",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unique_hash",
					Description: `Unique hash to allow user to import the self-created marketplace application.`,
				},
				resource.Attribute{
					Name:        "is_application_owner",
					Description: `Whether the you are the owner of application or not.`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `Whether the template is published by the partner to their tenant".`,
				},
				resource.Attribute{
					Name:        "published_date",
					Description: `The date when the template is published into other tenant in the same partner.`,
				},
				resource.Attribute{
					Name:        "is_publish_requested",
					Description: `Whether the tenants want their template to be published or not.`,
				},
				resource.Attribute{
					Name:        "publish_requested_date",
					Description: `The date when the tenant requested their template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global_requested",
					Description: `Whether a partner wants their tenant template published to other partners.`,
				},
				resource.Attribute{
					Name:        "publish_global_requested_date",
					Description: `The date when a partner requested their tenants template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global",
					Description: `Whether a template is published to other partner or not.`,
				},
				resource.Attribute{
					Name:        "published_global_date",
					Description: `The date when a template is published to other partner.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of template.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the marketplace application was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last marketplace application change.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object_storage_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "setup_cores",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "setup_memory",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "setup_storage_capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_license",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_os",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_components",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_overview",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_hints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_terms_of_use",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_icon",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_features",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_author",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "meta_advices",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unique_hash",
					Description: `Unique hash to allow user to import the self-created marketplace application.`,
				},
				resource.Attribute{
					Name:        "is_application_owner",
					Description: `Whether the you are the owner of application or not.`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `Whether the template is published by the partner to their tenant".`,
				},
				resource.Attribute{
					Name:        "published_date",
					Description: `The date when the template is published into other tenant in the same partner.`,
				},
				resource.Attribute{
					Name:        "is_publish_requested",
					Description: `Whether the tenants want their template to be published or not.`,
				},
				resource.Attribute{
					Name:        "publish_requested_date",
					Description: `The date when the tenant requested their template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global_requested",
					Description: `Whether a partner wants their tenant template published to other partners.`,
				},
				resource.Attribute{
					Name:        "publish_global_requested_date",
					Description: `The date when a partner requested their tenants template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global",
					Description: `Whether a template is published to other partner or not.`,
				},
				resource.Attribute{
					Name:        "published_global_date",
					Description: `The date when a template is published to other partner.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of template.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the marketplace application was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last marketplace application change.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_marketplace_application_import",
			Category:         "Resources",
			ShortDescription: `Manages imported marketplace applications in Gridscale.`,
			Description:      ``,
			Keywords: []string{
				"marketplace",
				"application",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "import_unique_hash",
					Description: `(Required) Hash of a specific marketplace application that you want to import. A change of this argument necessitates the re-creation of the resource. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of marketplace application.`,
				},
				resource.Attribute{
					Name:        "object_storage_path",
					Description: `Path to the images for the application.`,
				},
				resource.Attribute{
					Name:        "setup_cores",
					Description: `Number of server's cores.`,
				},
				resource.Attribute{
					Name:        "setup_memory",
					Description: `The capacity of server's memory in GB.`,
				},
				resource.Attribute{
					Name:        "setup_storage_capacity",
					Description: `The capacity of server's storage in GB.`,
				},
				resource.Attribute{
					Name:        "meta_license",
					Description: `License number.`,
				},
				resource.Attribute{
					Name:        "meta_os",
					Description: `Operating system.`,
				},
				resource.Attribute{
					Name:        "meta_components",
					Description: `Components (e.g: MySql, Apache, etc.).`,
				},
				resource.Attribute{
					Name:        "meta_overview",
					Description: `Describes the main function of the application.`,
				},
				resource.Attribute{
					Name:        "meta_hints",
					Description: `Hints.`,
				},
				resource.Attribute{
					Name:        "meta_terms_of_use",
					Description: `Terms of use.`,
				},
				resource.Attribute{
					Name:        "meta_icon",
					Description: `base64 encoded image of the icon.`,
				},
				resource.Attribute{
					Name:        "meta_features",
					Description: `List of functions.`,
				},
				resource.Attribute{
					Name:        "meta_author",
					Description: `Author.`,
				},
				resource.Attribute{
					Name:        "meta_advices",
					Description: `User manual; Wiki URL; ...`,
				},
				resource.Attribute{
					Name:        "unique_hash",
					Description: `Unique hash to allow user to import the self-created marketplace application.`,
				},
				resource.Attribute{
					Name:        "is_application_owner",
					Description: `Whether the you are the owner of application or not.`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `Whether the template is published by the partner to their tenant".`,
				},
				resource.Attribute{
					Name:        "published_date",
					Description: `The date when the template is published into other tenant in the same partner.`,
				},
				resource.Attribute{
					Name:        "is_publish_requested",
					Description: `Whether the tenants want their template to be published or not.`,
				},
				resource.Attribute{
					Name:        "publish_requested_date",
					Description: `The date when the tenant requested their template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global_requested",
					Description: `Whether a partner wants their tenant template published to other partners.`,
				},
				resource.Attribute{
					Name:        "publish_global_requested_date",
					Description: `The date when a partner requested their tenants template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global",
					Description: `Whether a template is published to other partner or not.`,
				},
				resource.Attribute{
					Name:        "published_global_date",
					Description: `The date when a template is published to other partner.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of template.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the marketplace application was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last marketplace application change.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of marketplace application.`,
				},
				resource.Attribute{
					Name:        "object_storage_path",
					Description: `Path to the images for the application.`,
				},
				resource.Attribute{
					Name:        "setup_cores",
					Description: `Number of server's cores.`,
				},
				resource.Attribute{
					Name:        "setup_memory",
					Description: `The capacity of server's memory in GB.`,
				},
				resource.Attribute{
					Name:        "setup_storage_capacity",
					Description: `The capacity of server's storage in GB.`,
				},
				resource.Attribute{
					Name:        "meta_license",
					Description: `License number.`,
				},
				resource.Attribute{
					Name:        "meta_os",
					Description: `Operating system.`,
				},
				resource.Attribute{
					Name:        "meta_components",
					Description: `Components (e.g: MySql, Apache, etc.).`,
				},
				resource.Attribute{
					Name:        "meta_overview",
					Description: `Describes the main function of the application.`,
				},
				resource.Attribute{
					Name:        "meta_hints",
					Description: `Hints.`,
				},
				resource.Attribute{
					Name:        "meta_terms_of_use",
					Description: `Terms of use.`,
				},
				resource.Attribute{
					Name:        "meta_icon",
					Description: `base64 encoded image of the icon.`,
				},
				resource.Attribute{
					Name:        "meta_features",
					Description: `List of functions.`,
				},
				resource.Attribute{
					Name:        "meta_author",
					Description: `Author.`,
				},
				resource.Attribute{
					Name:        "meta_advices",
					Description: `User manual; Wiki URL; ...`,
				},
				resource.Attribute{
					Name:        "unique_hash",
					Description: `Unique hash to allow user to import the self-created marketplace application.`,
				},
				resource.Attribute{
					Name:        "is_application_owner",
					Description: `Whether the you are the owner of application or not.`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `Whether the template is published by the partner to their tenant".`,
				},
				resource.Attribute{
					Name:        "published_date",
					Description: `The date when the template is published into other tenant in the same partner.`,
				},
				resource.Attribute{
					Name:        "is_publish_requested",
					Description: `Whether the tenants want their template to be published or not.`,
				},
				resource.Attribute{
					Name:        "publish_requested_date",
					Description: `The date when the tenant requested their template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global_requested",
					Description: `Whether a partner wants their tenant template published to other partners.`,
				},
				resource.Attribute{
					Name:        "publish_global_requested_date",
					Description: `The date when a partner requested their tenants template to be published.`,
				},
				resource.Attribute{
					Name:        "is_publish_global",
					Description: `Whether a template is published to other partner or not.`,
				},
				resource.Attribute{
					Name:        "published_global_date",
					Description: `The date when a template is published to other partner.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of template.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the marketplace application.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the marketplace application was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last marketplace application change.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_memcached",
			Category:         "Resources",
			ShortDescription: `Manage a Memcached service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"memcached",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The Memcached release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available Memcached service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of Memcached service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `(Optional) Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and ` + "`" + `max_core_count` + "`" + `. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the Memcached instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the Memcached instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this Memcached service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that Memcached service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_mysql",
			Category:         "Resources",
			ShortDescription: `Manage a MySQL service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The mysql release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available mysql service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of mysql service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mysql_log_bin",
					Description: `(Optional) mysql parameter: Binary Logging. Default: false.`,
				},
				resource.Attribute{
					Name:        "mysql_sql_mode",
					Description: `(Optional) mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".`,
				},
				resource.Attribute{
					Name:        "mysql_server_id",
					Description: `(Optional) mysql parameter: Server Id. Default: 1.`,
				},
				resource.Attribute{
					Name:        "mysql_query_cache",
					Description: `(Optional) mysql parameter: Enable query cache. Default: true.`,
				},
				resource.Attribute{
					Name:        "mysql_binlog_format",
					Description: `(Optional) mysql parameter: Binary Logging Format. Default: "ROW".`,
				},
				resource.Attribute{
					Name:        "mysql_max_connections",
					Description: `(Optional) mysql parameter: Max Connections. Default: 4000.`,
				},
				resource.Attribute{
					Name:        "mysql_query_cache_size",
					Description: `(Optional) mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.`,
				},
				resource.Attribute{
					Name:        "mysql_default_time_zone",
					Description: `(Optional) mysql parameter: Server Timezone. Default: UTC.`,
				},
				resource.Attribute{
					Name:        "mysql_query_cache_limit",
					Description: `(Optional) mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.`,
				},
				resource.Attribute{
					Name:        "mysql_max_allowed_packet",
					Description: `(Optional) mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `(Optional) Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and ` + "`" + `max_core_count` + "`" + `. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_log_bin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_sql_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_server_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_query_cache",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_binlog_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_max_connections",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_query_cache_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_default_time_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_query_cache_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mysql_max_allowed_packet",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the mysql instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the mysql instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this mysql service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that mysql service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_network",
			Category:         "Resources",
			ShortDescription: `Manages a network in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `(Optional) Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "dhcp_active",
					Description: `(Optional) Enable DHCP.`,
				},
				resource.Attribute{
					Name:        "dhcp_gateway",
					Description: `(Optional) The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.`,
				},
				resource.Attribute{
					Name:        "dhcp_dns",
					Description: `(Optional) The IP address reserved and communicated by the dhcp service to be the default gateway.`,
				},
				resource.Attribute{
					Name:        "dhcp_range",
					Description: `(Optional, Computed) DHCP DNS. If it is not set and DHCP is enabled, ` + "`" + `dhcp_range` + "`" + ` will be set by the backend automatically.`,
				},
				resource.Attribute{
					Name:        "dhcp_reserved_subnet",
					Description: `(Optional) Subrange within the IP range. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this network is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_active",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_range",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_reserved_subnet",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auto_assigned_servers",
					Description: `A list of server UUIDs with the corresponding IPs that are designated by the DHCP server.`,
				},
				resource.Attribute{
					Name:        "server_uuid",
					Description: `UUID of the server.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP which is assigned to the server.`,
				},
				resource.Attribute{
					Name:        "pinned_servers",
					Description: `A list of server UUIDs with the corresponding IPs that are designated by the user.`,
				},
				resource.Attribute{
					Name:        "server_uuid",
					Description: `UUID of the server.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP which is assigned to the server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of this network, can be mpls, breakout or network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "public_net",
					Description: `Is the network public or not.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `If deleting this network is allowed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_object_storage_accesskey",
			Category:         "Resources",
			ShortDescription: `Manages an access key of an object storage in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"accesskey",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The access key of the object storage.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key of an object storage.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key of an object storage.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_paas",
			Category:         "Resources",
			ShortDescription: `Manages a PaaS in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"paas",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `(Required) The template used to create the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Contains the service parameters for the service.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `(Required) Name of parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the corresponding parameter.`,
				},
				resource.Attribute{
					Name:        "resource_limit",
					Description: `(Optional) A list of service resource limits..`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) The name of the resource you would like to cap.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Required) The maximum number of the specific resource your service can use.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Primitive type of the parameter: bool, int (better use float for int case), float, string. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Ports that PaaS service listens to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid_computed",
					Description: `Template that PaaS service uses. The ` + "`" + `service_template_uuid_computed` + "`" + ` will be different from ` + "`" + `service_template_uuid` + "`" + `, when ` + "`" + `service_template_uuid` + "`" + ` is updated outside of terraform.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minute",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Current price of PaaS service.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of the creation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "parameter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "resource_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_postgres",
			Category:         "Resources",
			ShortDescription: `Manage a PostgreSQL service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"postgres",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The PostgreSQL release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available PostgreSQL service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of PostgreSQL service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `(Optional) Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and ` + "`" + `max_core_count` + "`" + `. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the PostgreSQL instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the PostgreSQL instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this PostgreSQL service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that PostgreSQL service uses.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "max_core_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_redis_cache",
			Category:         "Resources",
			ShortDescription: `Manage a Redis cache service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"cache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The Redis cache release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available Redis cache service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of Redis cache service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the Redis cache instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the Redis cache instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this Redis cache service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that Redis cache service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_redis_store",
			Category:         "Resources",
			ShortDescription: `Manage a Redis store service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The Redis store release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available Redis store service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of Redis store service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the Redis store instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the Redis store instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this Redis store service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that Redis store service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_paas_securityzone",
			Category:         "Resources",
			ShortDescription: `Manages a security zone in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"securityzone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) The location this object is placed. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the security zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the location's country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
				resource.Attribute{
					Name:        "relations",
					Description: `List of PaaS services' UUIDs relating to the security zone.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_server",
			Category:         "Resources",
			ShortDescription: `Manages a server in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) The number of server cores.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) The amount of server memory in GB.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `(Optional) If the server should be auto-started in case of a failure (default=true).`,
				},
				resource.Attribute{
					Name:        "hardware_profile",
					Description: `(Optional, Computed) The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check [the official docs](https://gridscale.io/en/api-documentation/index.html#operation/createServer).`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) The UUID of the IPv4 address of the server. (`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) The UUID of the IPv6 address of the server. (`,
				},
				resource.Attribute{
					Name:        "isoimage",
					Description: `(Optional) The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in [the expert panel](https://my.gridscale.io/Expert/ISOImage).`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `(Optional, Computed) The power state of the server. Set this to true to will boot the server, false will shut it down.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, Computed) Defines which Availability-Zone the Server is placed.`,
				},
				resource.Attribute{
					Name:        "hardware_profile_config",
					Description: `(Optional, Computed) Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time.`,
				},
				resource.Attribute{
					Name:        "machinetype",
					Description: `(Optional, Computed) Allowed values: ` + "`" + `"i440fx"` + "`" + `, ` + "`" + `"q35_bios"` + "`" + `, ` + "`" + `"q35_uefi"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_device",
					Description: `(Optional, Computed) Allowed values: ` + "`" + `"ide"` + "`" + `, ` + "`" + `"sata"` + "`" + `, ` + "`" + `"virtio_scsi"` + "`" + `, ` + "`" + `"virtio_block"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "usb_controller",
					Description: `(Optional, Computed) Allowed values: ` + "`" + `"nec_xhci"` + "`" + `, ` + "`" + `"piix3_uhci"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `(Optional, Computed) Boolean.`,
				},
				resource.Attribute{
					Name:        "hyperv_extensions",
					Description: `(Optional, Computed) Boolean.`,
				},
				resource.Attribute{
					Name:        "network_model",
					Description: `(Optional, Computed) Allowed values: ` + "`" + `"e1000"` + "`" + `, ` + "`" + `"e1000e"` + "`" + `, ` + "`" + `"virtio"` + "`" + `, ` + "`" + `"vmxnet3"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "serial_interface",
					Description: `(Optional, Computed) Boolean.`,
				},
				resource.Attribute{
					Name:        "server_renice",
					Description: `(Optional, Computed) Boolean.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Connects a storage to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `(Required) The object UUID or id of the storage.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `(Required) The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `(Optional, Computed) Make this network the boot device. This can only be set for one network.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Manually assign DHCP IP to the server (if applicable).`,
				},
				resource.Attribute{
					Name:        "firewall_template_uuid",
					Description: `(Optional) The UUID of firewall template.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `(Optional) Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `(Optional) Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `(Optional) Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `(Optional) Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of server cores.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of server memory in GB.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this server is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "hardware_profile",
					Description: `The hardware profile of the server.`,
				},
				resource.Attribute{
					Name:        "hardware_profile_config",
					Description: `(See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "machinetype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_device",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "usb_controller",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "hyperv_extensions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "serial_interface",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_renice",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Connects a storage to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the storage.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Indicates the speed of the storage. This may be (storage, storage_high or storage_insane).`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `True is the storage is a boot device.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the storage.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Capacity of the storage (GB).`,
				},
				resource.Attribute{
					Name:        "controller",
					Description: `Defines the SCSI controller id. The SCSI defines transmission routes such as Serial Attached SCSI (SAS), Fibre Channel and iSCSI.`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The SCSI bus id. The SCSI defines transmission routes like Serial Attached SCSI (SAS), Fibre Channel and iSCSI. Each SCSI device is addressed via a specific number. Each SCSI bus can have multiple SCSI devices connected to it.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Defines the SCSI target ID. The target ID is a device (e.g. disk).`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `Is the common SCSI abbreviation of the Logical Unit Number. A lun is a unique identifier for a single disk or a composite of disks.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Connects a network to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `Make this network the boot device. This can only be set for one network.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `DHCP IP which is manually assigned to the server (if applicable).`,
				},
				resource.Attribute{
					Name:        "auto_assigned_ip",
					Description: `DHCP IP which is automatically assigned to the server (if applicable).`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of network, network_high, network_insane.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `network_mac defines the MAC address of the network interface.`,
				},
				resource.Attribute{
					Name:        "firewall_template_uuid",
					Description: `The UUID of firewall template.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `Firewall template rules for inbound traffic - covers IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `Firewall template rules for outbound traffic - covers IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `Firewall template rules for inbound traffic - covers IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `Firewall template rules for outbound traffic - covers IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The UUID of the IPv4 address of the server.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The UUID of the IPv6 address of the server.`,
				},
				resource.Attribute{
					Name:        "isoimage",
					Description: `The UUID of an ISO image in gridscale.`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `The power state of the server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Defines which Availability-Zone the Server is placed.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `If the server should be auto-started in case of a failure.`,
				},
				resource.Attribute{
					Name:        "console_token",
					Description: `The token used by the panel to open the websocket VNC connection to the server console.`,
				},
				resource.Attribute{
					Name:        "legacy",
					Description: `Legacy-Hardware emulation instead of virtio hardware. If enabled, hot-plugging cores, memory, storage, network, etc. will not work, but the server will most likely run every x86 compatible operating system. This mode comes with a performance penalty, as emulated hardware does not benefit from the virtio driver infrastructure.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes_memory",
					Description: `Total minutes of memory used.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes_cores",
					Description: `Total minutes of cores used.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_snapshot",
			Category:         "Resources",
			ShortDescription: `Manages a storage snapshot in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage used to create this snapshot.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The list of labels.`,
				},
				resource.Attribute{
					Name:        "object_storage_export",
					Description: `(Optional) Export snapshot to a object storage.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Host of object storage. Must be of URL type, e.g., https://gos3.io`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) Access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Secret key.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Bucket name.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required) Name of file (include file path).`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `(Required) Privacy.`,
				},
				resource.Attribute{
					Name:        "rollback",
					Description: `(Optional) Returns a storage to the state of the selected Snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the rollback request. It can be any string value. Each rollback request has to have a UNIQUE id. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot change.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the ip has been running.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "object_storage_export",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the export request.`,
				},
				resource.Attribute{
					Name:        "rollback",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rollback_time",
					Description: `The time when rollback request is fulfilled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rollback request.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot change.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the ip has been running.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "object_storage_export",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the export request.`,
				},
				resource.Attribute{
					Name:        "rollback",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rollback_time",
					Description: `The time when rollback request is fulfilled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rollback request.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_snapshotschedule",
			Category:         "Resources",
			ShortDescription: `Manages a storage snapshot schedule.`,
			Description:      ``,
			Keywords: []string{
				"snapshotschedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage that the snapshot schedule belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The list of labels.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `(Optional) The date and time that the snapshot schedule will be run.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `(Required) The amount of Snapshots to keep before overwriting the last created Snapshot (>=1).`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `(Required) The interval at which the schedule will run (in minutes, >=60). ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime_computed",
					Description: `The date and time that the snapshot schedule will be run. This date and time is computed by gridscale's server.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot schedule was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot schedule change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `Related snapshots.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot was initially created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime_computed",
					Description: `The date and time that the snapshot schedule will be run. This date and time is computed by gridscale's server.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot schedule was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot schedule change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `Related snapshots.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot was initially created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_sqlserver",
			Category:         "Resources",
			ShortDescription: `Manage a MS SQL server service in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Required) The MS SQL server release of this instance. For convenience, please use [gscloud](https://github.com/gridscale/gscloud) to get the list of available MS SQL server service releases.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `(Required) Performance class of MS SQL server service. Available performance classes at the time of writing: ` + "`" + `standard` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `insane` + "`" + `, ` + "`" + `ultra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The UUID of the network that the service is attached to.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: ``,
				},
				resource.Attribute{
					Name:        "s3_backup",
					Description: `(Optional) Allow backup/restore MS SQL server to/from a S3 bucket.`,
				},
				resource.Attribute{
					Name:        "backup_bucket",
					Description: `(Required) Object Storage bucket to upload backups to and restore backups from.`,
				},
				resource.Attribute{
					Name:        "backup_retention",
					Description: `(Optional) Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).`,
				},
				resource.Attribute{
					Name:        "backup_access_key",
					Description: `(Required) Access key used to authenticate against Object Storage server.`,
				},
				resource.Attribute{
					Name:        "backup_secret_key",
					Description: `(Required) Secret key used to authenticate against Object Storage server.`,
				},
				resource.Attribute{
					Name:        "backup_server_url",
					Description: `(Optional, Default: "https://gos3.io/") Object Storage server URL the bucket is located on.`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "15m" - 15 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "15m" - 15 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "15m" - 15 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "performance_class",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service. It is used to connect to the MS SQL server instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service. It is used to connect to the MS SQL server instance.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `The port numbers where this MS SQL server service accepts connections.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host address.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "s3_backup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_bucket",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_access_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_secret_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_server_url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The UUID of the network that the service is attached to or network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `PaaS service template that MS SQL server service uses.`,
				},
				resource.Attribute{
					Name:        "service_template_category",
					Description: `The template service's category used to create the service.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Date time this service has been created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_sshkey",
			Category:         "Resources",
			ShortDescription: `Manages an SSH public key in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"sshkey",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `(Required) This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ssl_certificate",
			Category:         "Resources",
			ShortDescription: `Manages a TLS/SSL certificate resource in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Force New) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required, Force New) The PEM-formatted private-key of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "leaf_certificate",
					Description: `(Required, Force New) The PEM-formatted public SSL of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional, Force New) The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, Force New) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `The common domain name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "leaf_certificate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprints",
					Description: `Defines a list of unique identifiers generated from the MD5, SHA-1, and SHA-256 fingerprints of the certificate.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "not_valid_after",
					Description: `Defines the date after which the certificate is not valid.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage",
			Category:         "Resources",
			ShortDescription: `Manages a storage in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) required (integer - minimum: 1 - maximum: 4096).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) (one of storage, storage_high, storage_insane).`,
				},
				resource.Attribute{
					Name:        "storage_variant",
					Description: `(Optional) Storage variant (one of local or distributed). Default: "distributed".`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "rollback_from_backup_uuid",
					Description: `(Optional) Rollback the storage from a specific storage backup.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `(Required) The UUID of a template. This can be found in the [the page Template](https://my.gridscale.io/Template) by clicking more on the template or by using a gridscale_template datasource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plain-text or a crypt string (modular crypt format - MCF).`,
				},
				resource.Attribute{
					Name:        "password_type",
					Description: `(Optional) (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates).`,
				},
				resource.Attribute{
					Name:        "sshkeys",
					Description: `(Optional) (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of the installed server (ignored for private templates and public windows templates). ~>`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_variant",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this storage is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rollback_from_backup_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage_clone",
			Category:         "Resources",
			ShortDescription: `Make a clone of an existing storage instance.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"clone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_storage_id",
					Description: `(Required) The ID of a storage instance which will be cloned.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this resource is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage_import",
			Category:         "Resources",
			ShortDescription: `Make a clone of an existing storage instance.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_backup_id",
					Description: `(Required) ID of the storage backup that will be used to create a new storage from.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this resource is placed. The location of a resource is determined by it's project.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_template",
			Category:         "Resources",
			ShortDescription: `Manages a template in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The exact name of the template as show in [the page Template](https://my.gridscale.io/Template).`,
				},
				resource.Attribute{
					Name:        "snapshot_uuid",
					Description: `(Required) Snapshot uuid for template.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels. ## Timeouts Timeouts configuration options (in seconds): More info: [terraform.io/docs/configuration/resources.html#operation-timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts)`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default value is "5m" - 5 minutes) Used for creating a resource.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default value is "5m" - 5 minutes) Used for updating a resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default value is "5m" - 5 minutes) Used for deleting a resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the template.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "ostype",
					Description: `The operating system installed in the template.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the template.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `The OS distribution that the template contains.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/template/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the template.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The location this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Two digit country code (ISO 3166-2) of the location where this object is placed.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "ostype",
					Description: `The operating system installed in the template.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the template.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `The OS distribution that the template contains.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/template/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"gridscale_backupschedule":                 0,
		"gridscale_object_storage_bucket":          1,
		"gridscale_filesystem":                     2,
		"gridscale_firewall":                       3,
		"gridscale_ipv4":                           4,
		"gridscale_ipv6":                           5,
		"gridscale_isoimage":                       6,
		"gridscale_k8s":                            7,
		"gridscale_loadbalancer":                   8,
		"gridscale_mariadb":                        9,
		"gridscale_marketplace_application":        10,
		"gridscale_marketplace_application_import": 11,
		"gridscale_memcached":                      12,
		"gridscale_mysql":                          13,
		"gridscale_network":                        14,
		"gridscale_object_storage_accesskey":       15,
		"gridscale_paas":                           16,
		"gridscale_postgres":                       17,
		"gridscale_redis_cache":                    18,
		"gridscale_redis_store":                    19,
		"gridscale_paas_securityzone":              20,
		"gridscale_server":                         21,
		"gridscale_snapshot":                       22,
		"gridscale_snapshotschedule":               23,
		"gridscale_sqlserver":                      24,
		"gridscale_sshkey":                         25,
		"gridscale_ssl_certificate":                26,
		"gridscale_storage":                        27,
		"gridscale_storage_clone":                  28,
		"gridscale_storage_import":                 29,
		"gridscale_template":                       30,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
