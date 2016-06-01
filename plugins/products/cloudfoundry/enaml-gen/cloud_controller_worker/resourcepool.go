package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ResourcePool struct {

	/*ResourceDirectoryKey - Descr: Directory (bucket) used store app resources.  It does not have be pre-created. Default: cc-resources
*/
	ResourceDirectoryKey interface{} `yaml:"resource_directory_key,omitempty"`

	/*Cdn - Descr: Key pair name for signed download URIs Default: 
*/
	Cdn *Cdn `yaml:"cdn,omitempty"`

	/*MinimumSize - Descr: Minimum size of a resource to add to the pool Default: 65536
*/
	MinimumSize interface{} `yaml:"minimum_size,omitempty"`

	/*MaximumSize - Descr: Maximum size of a resource to add to the pool Default: 536870912
*/
	MaximumSize interface{} `yaml:"maximum_size,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*WebdavConfig - Descr: The location of the webdav server eg: https://blobstore.com Default: 
*/
	WebdavConfig *WebdavConfig `yaml:"webdav_config,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

}