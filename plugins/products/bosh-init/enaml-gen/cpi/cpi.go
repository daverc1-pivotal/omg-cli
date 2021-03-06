package cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cpi struct {

	/*Agent - Descr: Options for the blobstore used by deployed BOSH agents Default: map[]
*/
	Agent *CpiAgent `yaml:"agent,omitempty"`

	/*Actions - Descr: Port of registry Default: 6301
*/
	Actions *Actions `yaml:"actions,omitempty"`

	/*HostIp - Descr: IP address of the host that will be used by containers, must be the same as mbus IP Default: <nil>
*/
	HostIp interface{} `yaml:"host_ip,omitempty"`

}