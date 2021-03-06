package health_monitor 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ConsulEventForwarder struct {

	/*HeartbeatsAsAlerts - Descr: Should we treat all heartbeats as alerts as well? Default: false
*/
	HeartbeatsAsAlerts interface{} `yaml:"heartbeats_as_alerts,omitempty"`

	/*Host - Descr: Location of Consul Cluster or agent Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

	/*Namespace - Descr: A namespace for handling multiples of the same release Default: <nil>
*/
	Namespace interface{} `yaml:"namespace,omitempty"`

	/*TtlNote - Descr: A note for ttl checks Default: Automatically Registered by Bosh-Monitor
*/
	TtlNote interface{} `yaml:"ttl_note,omitempty"`

	/*Ttl - Descr: A ttl time for ttl checks, if set ttl checks will be used Default: <nil>
*/
	Ttl interface{} `yaml:"ttl,omitempty"`

	/*Protocol - Descr: http/https Default: http
*/
	Protocol interface{} `yaml:"protocol,omitempty"`

	/*Events - Descr: Whether or not to use the events api Default: false
*/
	Events interface{} `yaml:"events,omitempty"`

	/*Params - Descr: Params for url can be used for passing ACL token Default: <nil>
*/
	Params interface{} `yaml:"params,omitempty"`

	/*Port - Descr: Consul Port Default: 8500
*/
	Port interface{} `yaml:"port,omitempty"`

}