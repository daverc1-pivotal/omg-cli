package broker_deregistrar 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BrokerDeregistrarJob struct {

	/*Cf - Descr: Password of the admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*Broker - Descr: Name of the service broker Default: <nil>
*/
	Broker *Broker `yaml:"broker,omitempty"`

	/*Redis - Descr: Service name Default: p-redis
*/
	Redis *Redis `yaml:"redis,omitempty"`

}