package photoncpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AgentBlobstore struct {

	/*Provider - Descr: Provider type for the blobstore used by deployed BOSH agents (e.g. dav, s3) Default: dav
*/
	Provider interface{} `yaml:"provider,omitempty"`

	/*Options - Descr: Options for the blobstore used by deployed BOSH agents Default: map[]
*/
	Options interface{} `yaml:"options,omitempty"`

}