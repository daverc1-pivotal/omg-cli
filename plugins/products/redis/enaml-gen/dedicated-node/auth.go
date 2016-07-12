package dedicated_node 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Auth struct {

	/*Username - Descr: The username for HTTP Basic Auth on the agent, also used for the broker Default: admin
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Password - Descr: The password for HTTP Basic Auth on the agent, also used for the broker Default: admin
*/
	Password interface{} `yaml:"password,omitempty"`

}