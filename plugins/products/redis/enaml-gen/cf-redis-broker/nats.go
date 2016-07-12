package cf_redis_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*Username - Descr: The user to use when authenticating with NATS Default: <nil>
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Port - Descr: Port that NATS listens on Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Host - Descr: Hostname/IP of NATS Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

	/*Password - Descr: The password to use when authenticating with NATS Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

}