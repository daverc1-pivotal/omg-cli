package cloudfoundry

import (
	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/pluginlib/product"
	"github.com/enaml-ops/omg-cli/pluginlib/util"
	"github.com/xchapter7x/lo"
	"gopkg.in/yaml.v2"
)

func init() {
	RegisterInstanceGrouperFactory(NewConsulPartition)
	RegisterInstanceGrouperFactory(NewNatsPartition)
	RegisterInstanceGrouperFactory(NewEtcdPartition)
	//diego_database-partition
	RegisterInstanceGrouperFactory(NewNFSPartition)
	RegisterInstanceGrouperFactory(NewGoRouterPartition)
	RegisterInstanceGrouperFactory(NewMySQLProxyPartition)
	RegisterInstanceGrouperFactory(NewMySQLPartition)
	RegisterInstanceGrouperFactory(NewCloudControllerPartition)
	//ha_proxy-partition
	//clock_global-partition
	RegisterInstanceGrouperFactory(NewCloudControllerWorkerPartition)
	//uaa-partition
	RegisterInstanceGrouperFactory(NewDiegoBrainPartition)
	//diego_cell-partition
	//doppler-partition
	//loggregator_trafficcontroller-partition
}

//GetFlags -
func (s *Plugin) GetFlags() (flags []cli.Flag) {
	return []cli.Flag{
		// shared for all instance groups:
		cli.StringFlag{Name: "stemcell-name", Usage: "the name of your desired stemcell"},
		cli.StringSliceFlag{Name: "az", Usage: "list of AZ names to use"},
		cli.StringFlag{Name: "network", Usage: "the name of the network to use"},
		cli.StringFlag{Name: "system-domain", Usage: "System Domain"},
		cli.StringSliceFlag{Name: "app-domain", Usage: "Applications Domain"},
		cli.BoolFlag{Name: "allow-app-ssh-access", Usage: "Allow SSH access for CF applications"},

		cli.StringSliceFlag{Name: "router-ip", Usage: "a list of the router ips you wish to use"},
		cli.StringFlag{Name: "router-vm-type", Usage: "the name of your desired vm size"},
		cli.StringFlag{Name: "router-ssl-cert", Usage: "the go router ssl cert, or a filename preceded by '@'"},
		cli.StringFlag{Name: "router-ssl-key", Usage: "the go router ssl key, or a filename preceded by '@'"},
		cli.StringFlag{Name: "router-user", Value: "router_status", Usage: "the username of the go-routers"},
		cli.StringFlag{Name: "router-pass", Usage: "the password of the go-routers"},
		cli.BoolFlag{Name: "router-enable-ssl", Usage: "enable or disable ssl on your routers"},

		cli.StringSliceFlag{Name: "haproxy-ip", Usage: "a list of the haproxy ips you wish to use"},
		cli.StringFlag{Name: "haproxy-vm-type", Usage: "the name of your desired vm size"},

		cli.StringFlag{Name: "nats-vm-type", Usage: "the name of your desired vm size for NATS"},
		cli.StringFlag{Name: "nats-user", Value: "nats", Usage: "username for your nats pool"},
		cli.StringFlag{Name: "nats-pass", Value: "nats-password", Usage: "password for your nats pool"},
		cli.IntFlag{Name: "nats-port", Usage: "the port for the NATS server to listen on"},
		cli.StringSliceFlag{Name: "nats-machine-ip", Usage: "ip of a nats node vm"},

		cli.StringFlag{Name: "metron-zone", Usage: "zone guid for the metron agent"},
		cli.StringFlag{Name: "metron-secret", Usage: "shared secret for the metron agent endpoint"},
		cli.IntFlag{Name: "metron-port", Usage: "local metron agent's port"},

		cli.StringSliceFlag{Name: "consul-ip", Usage: "a list of the consul ips you wish to use"},
		cli.StringFlag{Name: "consul-vm-type", Usage: "the name of your desired vm size for consul"},
		cli.StringSliceFlag{Name: "consul-encryption-key", Usage: "encryption key for consul"},
		cli.StringFlag{Name: "consul-ca-cert", Usage: "ca cert for consul, or a filename preceded by '@'"},
		cli.StringFlag{Name: "consul-agent-cert", Usage: "agent cert for consul, or a filename preceded by '@'"},
		cli.StringFlag{Name: "consul-agent-key", Usage: "agent key for consul, or a filename preceded by '@'"},
		cli.StringFlag{Name: "consul-server-cert", Usage: "server cert for consul, or a filename preceded by '@'"},
		cli.StringFlag{Name: "consul-server-key", Usage: "server key for consul, or a filename preceded by '@'"},

		cli.StringFlag{Name: "syslog-address", Usage: "address of syslog server"},
		cli.IntFlag{Name: "syslog-port", Usage: "port of syslog server"},
		cli.StringFlag{Name: "syslog-transport", Usage: "transport to syslog server"},

		cli.StringSliceFlag{Name: "etcd-machine-ip", Usage: "ip of a etcd node vm"},
		cli.StringFlag{Name: "etcd-vm-type", Usage: "the name of your desired vm size for etcd"},
		cli.StringFlag{Name: "etcd-disk-type", Usage: "the name of your desired persistent disk type for etcd"},

		cli.StringSliceFlag{Name: "nfs-ip", Usage: "a list of the nfs ips you wish to use"},
		cli.StringFlag{Name: "nfs-vm-type", Usage: "the name of your desired vm size for nfs"},
		cli.StringFlag{Name: "nfs-disk-type", Usage: "the name of your desired persistent disk type for nfs"},
		cli.StringFlag{Name: "nfs-server-address", Usage: "NFS Server address"},
		cli.StringFlag{Name: "nfs-share-path", Usage: "NFS Share Path"},
		cli.StringSliceFlag{Name: "nfs-allow-from-network-cidr", Usage: "the network cidr you wish to allow connections to nfs from"},

		//Mysql Flags
		cli.StringSliceFlag{Name: "mysql-ip", Usage: "a list of the mysql ips you wish to use"},
		cli.StringFlag{Name: "mysql-vm-type", Usage: "the name of your desired vm size for mysql"},
		cli.StringFlag{Name: "mysql-disk-type", Usage: "the name of your desired persistent disk type for mysql"},
		cli.StringFlag{Name: "mysql-admin-password", Usage: "admin password for mysql"},
		cli.StringFlag{Name: "mysql-bootstrap-username", Usage: "bootstrap username for mysql"},
		cli.StringFlag{Name: "mysql-bootstrap-password", Usage: "bootstrap password for mysql"},

		//MySQL proxy flags
		cli.StringSliceFlag{Name: "mysql-proxy-ip", Usage: "a list of -mysql proxy ips you wish to use"},
		cli.StringFlag{Name: "mysql-proxy-vm-type", Usage: "the name of your desired vm size for mysql proxy"},
		cli.StringFlag{Name: "mysql-proxy-external-host", Usage: "Host name of MySQL proxy"},
		cli.StringFlag{Name: "mysql-proxy-api-username", Usage: "Proxy API user name"},
		cli.StringFlag{Name: "mysql-proxy-api-password", Usage: "Proxy API password"},

		//CC Worker Partition Flags
		cli.StringFlag{Name: "cc-worker-vm-type", Usage: "the name of the desired vm type for cc worker"},
		cli.StringFlag{Name: "cc-staging-upload-user", Usage: "user name for staging upload"},
		cli.StringFlag{Name: "cc-staging-upload-password", Usage: "password for staging upload"},
		cli.StringFlag{Name: "cc-bulk-api-user", Usage: "user name for bulk api calls"},
		cli.StringFlag{Name: "cc-bulk-api-password", Usage: "password for bulk api calls"},
		cli.IntFlag{Name: "cc-bulk-batch-size", Usage: "number of apps to fetch at once from bulk API"},
		cli.StringFlag{Name: "cc-db-encryption-key", Usage: "Cloud Controller DB encryption key"},
		cli.StringFlag{Name: "cc-internal-api-user", Usage: "user name for Internal API calls"},
		cli.StringFlag{Name: "cc-internal-api-password", Usage: "password for Internal API calls"},
		cli.IntFlag{Name: "cc-uploader-poll-interval", Usage: "CC uploader job polling interval, in seconds"},
		cli.IntFlag{Name: "cc-fetch-timeout", Usage: "how long to wait for completion of requests to CC, in seconds"},
		cli.StringFlag{Name: "cc-vm-type", Usage: "Cloud Controller VM Type"},
		cli.StringFlag{Name: "host-key-fingerprint", Usage: "Host Key Fingerprint"},
		cli.StringFlag{Name: "support-address", Usage: "Support URL"},
		cli.StringFlag{Name: "min-cli-version", Usage: "Min CF CLI Version supported"},

		cli.StringFlag{Name: "db-uaa-username", Usage: "uaa db username"},
		cli.StringFlag{Name: "db-uaa-password", Usage: "uaa db password"},
		cli.StringFlag{Name: "db-ccdb-username", Usage: "ccdb db username"},
		cli.StringFlag{Name: "db-ccdb-password", Usage: "ccdb db password"},
		cli.StringFlag{Name: "db-console-username", Usage: "console db username"},
		cli.StringFlag{Name: "db-console-password", Usage: "console db password"},

		// Diego Cell
		cli.StringSliceFlag{Name: "diego-cell-ip", Usage: "a list of static IPs for the diego brain"},
		cli.StringFlag{Name: "diego-cell-vm-type", Usage: "the name of the desired vm type for the diego cell"},
		cli.StringFlag{Name: "diego-cell-disk-type", Usage: "the name of your desired persistent disk type for the diego cell"},

		// Diego Brain
		cli.StringSliceFlag{Name: "diego-brain-ip", Usage: "a list of static IPs for the diego brain"},
		cli.StringFlag{Name: "diego-brain-vm-type", Usage: "the name of the desired vm type for the diego brain"},
		cli.StringFlag{Name: "diego-brain-disk-type", Usage: "the name of your desired persistent disk type for the diego brain"},

		cli.StringFlag{Name: "bbs-ca-cert", Usage: "BBS CA SSL cert (or a file containing it)"},
		cli.StringFlag{Name: "bbs-client-cert", Usage: "BBS client SSL cert (or a file containing it)"},
		cli.StringFlag{Name: "bbs-client-key", Usage: "BBS client SSL key (or a file containing it)"},
		cli.StringFlag{Name: "bbs-api", Usage: "location of the bbs api"},
		cli.BoolTFlag{Name: "bbs-require-ssl", Usage: "enable SSL for all communications with the BBS"},

		cli.BoolTFlag{Name: "skip-cert-verify", Usage: "ignore bad SSL certificates when connecting over HTTPS"},

		cli.StringFlag{Name: "fs-listen-addr", Usage: "address of interface on which to serve files"},
		cli.StringFlag{Name: "fs-static-dir", Usage: "fully qualified path to the doc root for the file server's static files"},
		cli.StringFlag{Name: "fs-debug-addr", Usage: "address at which to serve debug info"},
		cli.StringFlag{Name: "fs-log-level", Usage: "file server log level"},

		cli.IntFlag{Name: "cc-external-port", Usage: "external port of the Cloud Controller API"},
		cli.StringFlag{Name: "ssh-proxy-uaa-secret", Usage: "the OAuth client secret used to authenticate the SSH proxy"},
		cli.StringFlag{Name: "traffic-controller-url", Usage: "the URL of the traffic controller"},

		//Doppler
		cli.StringSliceFlag{Name: "doppler-ip", Usage: "a list of the doppler ips you wish to use"},
		cli.StringFlag{Name: "doppler-vm-type", Usage: "the name of your desired vm size for doppler"},
		cli.StringFlag{Name: "doppler-zone", Usage: "the name zone for doppler"},
		cli.IntFlag{Name: "doppler-drain-buffer-size", Usage: "message drain buffer size"},
		cli.StringFlag{Name: "doppler-shared-secret", Usage: "doppler shared secret"},

		//UAA
		cli.StringFlag{Name: "uaa-vm-type", Usage: "the name of your desired vm size for uaa"},
		cli.IntFlag{Name: "uaa-instances", Usage: "the number of your desired vms for uaa"},

		cli.StringFlag{Name: "uaa-company-name", Usage: "name of company for UAA branding"},
		cli.StringFlag{Name: "uaa-product-logo", Usage: "product logo for UAA branding"},
		cli.StringFlag{Name: "uaa-square-logo", Usage: "square logo for UAA branding"},
		cli.StringFlag{Name: "uaa-footer-legal-txt", Usage: "legal text for UAA branding"},
		cli.BoolTFlag{Name: "uaa-enable-selfservice-links", Usage: "enable self service links"},
		cli.BoolTFlag{Name: "uaa-signups-enabled", Usage: "enable signups"},
		cli.StringFlag{Name: "uaa-login-protocol", Usage: "uaa login protocol, default https"},
		cli.StringFlag{Name: "uaa-saml-service-provider-key", Usage: "saml service provider key for uaa"},
		cli.StringFlag{Name: "uaa-saml-service-provider-certificate", Usage: "saml service provider certificate for uaa"},
		cli.StringFlag{Name: "uaa-jwt-signing-key", Usage: "signing key for jwt used by UAA"},
		cli.StringFlag{Name: "uaa-jwt-verification-key", Usage: "verification key for jwt used by UAA"},
		cli.BoolFlag{Name: "uaa-ldap-enabled", Usage: "is ldap enabled for UAA"},
		cli.StringFlag{Name: "uaa-ldap-url", Usage: "url for ldap server"},
		cli.StringFlag{Name: "uaa-ldap-user-dn", Usage: "userDN to bind to ldap with"},
		cli.StringFlag{Name: "uaa-ldap-user-password", Usage: "bind password for ldap user"},
		cli.StringFlag{Name: "uaa-ldap-search-filter", Usage: "search filter for users"},
		cli.StringFlag{Name: "uaa-ldap-search-base", Usage: "search base for users"},
		cli.StringFlag{Name: "uaa-ldap-mail-attributename", Usage: "attribute name for mail"},
		cli.StringFlag{Name: "uaa-admin-secret", Usage: "admin account client secret"},

		//User accounts
		cli.StringFlag{Name: "admin-password", Usage: "password for admin account"},
		cli.StringFlag{Name: "push-apps-manager-password", Usage: "password for push_apps_manager account"},
		cli.StringFlag{Name: "smoke-tests-password", Usage: "password for smoke_tests account"},
		cli.StringFlag{Name: "system-services-password", Usage: "password for system_services account"},
		cli.StringFlag{Name: "system-verification-password", Usage: "password for system_verification account"},

		//Client secrets
		cli.StringFlag{Name: "opentsdb-firehose-nozzle-client-secret", Usage: "client-secret for opentsdb firehose nozzle"},
		cli.StringFlag{Name: "identity-client-secret", Usage: "client-secret for identity"},
		cli.StringFlag{Name: "login-client-secret", Usage: "client-secret for login"},
		cli.StringFlag{Name: "portal-client-secret", Usage: "client-secret for portal"},
		cli.StringFlag{Name: "autoscaling-service-client-secret", Usage: "client-secret for autoscaling service"},
		cli.StringFlag{Name: "system-passwords-client-secret", Usage: "client-secret for system-passwords"},
		cli.StringFlag{Name: "cc-service-dashboards-client-secret", Usage: "client-secret for cc-service-dashboards"},
		cli.StringFlag{Name: "doppler-client-secret", Usage: "client-secret for doppler"},
		cli.StringFlag{Name: "gorouter-client-secret", Usage: "client-secret for gorouter"},
		cli.StringFlag{Name: "notifications-client-secret", Usage: "client-secret for notifications"},
		cli.StringFlag{Name: "notifications-ui-client-secret", Usage: "client-secret for notification-ui"},
		cli.StringFlag{Name: "cloud-controller-username-lookup-client-secret", Usage: "client-secret for cloud controller username lookup"},
		cli.StringFlag{Name: "cc-routing-client-secret", Usage: "client-secret for cc routing"},
		cli.StringFlag{Name: "ssh-proxy-client-secret", Usage: "client-secret for ssh proxy"},
		cli.StringFlag{Name: "apps-metrics-client-secret", Usage: "client-secret for apps metrics "},
		cli.StringFlag{Name: "apps-metrics-processing-client-secret", Usage: "client-secret for apps metrics processing"},

		cli.StringFlag{Name: "errand-vm-type", Usage: "vm type to be used for running errands"},
	}
}

//GetMeta -
func (s *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "cloudfoundry",
	}
}

//GetProduct -
func (s *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte) {
	c := pluginutil.NewContext(args, s.GetFlags())
	dm := enaml.NewDeploymentManifest([]byte(``))
	dm.SetName(DeploymentName)

	dm.AddRelease(enaml.Release{Name: CFReleaseName, Version: CFReleaseVersion})
	dm.AddRelease(enaml.Release{Name: CFMysqlReleaseName, Version: CFMysqlReleaseVersion})
	dm.AddRelease(enaml.Release{Name: DiegoReleaseName, Version: DiegoReleaseVersion})
	dm.AddRelease(enaml.Release{Name: GardenReleaseName, Version: GardenReleaseVersion})
	dm.AddRelease(enaml.Release{Name: CFLinuxReleaseName, Version: CFLinuxReleaseVersion})
	dm.AddRelease(enaml.Release{Name: EtcdReleaseName, Version: EtcdReleaseVersion})
	dm.AddRelease(enaml.Release{Name: PushAppsReleaseName, Version: PushAppsReleaseVersion})
	dm.AddRelease(enaml.Release{Name: NotificationsReleaseName, Version: NotificationsReleaseVersion})
	dm.AddRelease(enaml.Release{Name: NotificationsUIReleaseName, Version: NotificationsUIReleaseVersion})
	dm.AddRelease(enaml.Release{Name: CFAutoscalingReleaseName, Version: CFAutoscalingReleaseVersion})

	dm.AddStemcell(enaml.Stemcell{OS: StemcellName, Version: StemcellVersion, Alias: StemcellAlias})

	for _, factory := range factories {
		// create and validate all registered instance groups
		grouper := factory(c)
		if grouper.HasValidValues() {
			ig := grouper.ToInstanceGroup()
			lo.G.Debug("instance-group: ", ig)
			dm.AddInstanceGroup(ig)
		} else {
			b, _ := yaml.Marshal(grouper)
			lo.G.Panic("invalid values in instance group: ", string(b))
		}
	}

	return dm.Bytes()
}

//GetContext -
func (s *Plugin) GetContext(args []string) (c *cli.Context) {
	c = pluginutil.NewContext(args, s.GetFlags())
	return
}
