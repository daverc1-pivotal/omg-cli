package metron_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Tls struct {

	/*ClientKey - Descr: TLS client key Default: 
*/
	ClientKey interface{} `yaml:"client_key,omitempty"`

	/*CaCert - Descr: CA root required for key/cert verification Default: 
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*ClientCert - Descr: TLS client certificate Default: 
*/
	ClientCert interface{} `yaml:"client_cert,omitempty"`

}