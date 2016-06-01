package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Uaa struct {

	/*Clients - Descr: Used for generating SSO clients for service brokers. Default: <nil>
*/
	Clients *Clients `yaml:"clients,omitempty"`

	/*Jwt - Descr: ssl cert defined in the manifest by the UAA, required by the cc to communicate with UAA Default: 
*/
	Jwt *Jwt `yaml:"jwt,omitempty"`

	/*Url - Descr: URL of the UAA server Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*Cc - Descr: Symmetric secret used to decode uaa tokens. Used for testing. Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

}