package cc_uploader 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CcUploader struct {

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

	/*ListenAddr - Descr: Address of interface on which to serve files Default: 0.0.0.0:9090
*/
	ListenAddr interface{} `yaml:"listen_addr,omitempty"`

	/*DebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17018
*/
	DebugAddr interface{} `yaml:"debug_addr,omitempty"`

	/*Cc - Descr: the interval between job polling requests Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*Diego - Descr: Log level Default: info
*/
	Diego *Diego `yaml:"diego,omitempty"`

}