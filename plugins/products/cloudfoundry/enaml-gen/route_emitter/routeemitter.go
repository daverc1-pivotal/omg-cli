package route_emitter 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type RouteEmitter struct {

	/*Diego - Descr: local metron agent's port Default: 3457
*/
	Diego *Diego `yaml:"diego,omitempty"`

	/*Nats - Descr: Password for server authentication. Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Bbs - Descr: PEM-encoded client certificate Default: <nil>
*/
	Bbs *Bbs `yaml:"bbs,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*DebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17009
*/
	DebugAddr interface{} `yaml:"debug_addr,omitempty"`

	/*SyncIntervalInSeconds - Descr: Interval to sync routes to the router in seconds. Default: 60
*/
	SyncIntervalInSeconds interface{} `yaml:"sync_interval_in_seconds,omitempty"`

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

}