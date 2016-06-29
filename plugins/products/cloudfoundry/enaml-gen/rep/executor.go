package rep 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Executor struct {

	/*MemoryCapacityMb - Descr: the memory capacity the executor should manage.  this should not be greater than the actual memory on the VM Default: auto
*/
	MemoryCapacityMb interface{} `yaml:"memory_capacity_mb,omitempty"`

	/*DeleteWorkPoolSize - Descr: Maximum number of concurrent delete container operations. Default: 32
*/
	DeleteWorkPoolSize interface{} `yaml:"delete_work_pool_size,omitempty"`

	/*MaxConcurrentDownloads - Descr: the max concurrent download steps that can be active Default: 5
*/
	MaxConcurrentDownloads interface{} `yaml:"max_concurrent_downloads,omitempty"`

	/*CachePath - Descr: path to the executor's cache Default: /var/vcap/data/executor_cache
*/
	CachePath interface{} `yaml:"cache_path,omitempty"`

	/*HealthyMonitoringIntervalInSeconds - Descr: Interval to check healthy containers in seconds. Default: 30
*/
	HealthyMonitoringIntervalInSeconds interface{} `yaml:"healthy_monitoring_interval_in_seconds,omitempty"`

	/*PostSetupHook - Descr: Experimental: arbitrary command to run after setup action Default: <nil>
*/
	PostSetupHook interface{} `yaml:"post_setup_hook,omitempty"`

	/*ContainerInodeLimit - Descr: the inode limit enforced on each garden container. Default: 200000
*/
	ContainerInodeLimit interface{} `yaml:"container_inode_limit,omitempty"`

	/*MaxCacheSizeInBytes - Descr: maximum size of the cache in bytes - this should leave a healthy overhead for temporary items, etc. Default: 10000000000
*/
	MaxCacheSizeInBytes interface{} `yaml:"max_cache_size_in_bytes,omitempty"`

	/*ExportNetworkEnvVars - Descr: Export network environment variables into container (e.g. CF_INSTANCE_IP, CF_INSTANCE_PORT). Default: true
*/
	ExportNetworkEnvVars interface{} `yaml:"export_network_env_vars,omitempty"`

	/*UnhealthyMonitoringIntervalInSeconds - Descr: Interval to check unhealthy containers in seconds. Default: 0.5
*/
	UnhealthyMonitoringIntervalInSeconds interface{} `yaml:"unhealthy_monitoring_interval_in_seconds,omitempty"`

	/*ContainerMaxCpuShares - Descr: the maximum number of cpu shares for a container. Default: 1024
*/
	ContainerMaxCpuShares interface{} `yaml:"container_max_cpu_shares,omitempty"`

	/*PostSetupUser - Descr: Experimental: user to run post setup hook command Default: <nil>
*/
	PostSetupUser interface{} `yaml:"post_setup_user,omitempty"`

	/*HealthcheckWorkPoolSize - Descr: Maximum number of concurrent health check operations. Default: 64
*/
	HealthcheckWorkPoolSize interface{} `yaml:"healthcheck_work_pool_size,omitempty"`

	/*Garden - Descr: Network type for the garden server connection (tcp or unix). Default: unix
*/
	Garden *Garden `yaml:"garden,omitempty"`

	/*CaCertsForDownloads - Descr: Concatenation of trusted CA certificates to be used when downloading assets. Default: <nil>
*/
	CaCertsForDownloads interface{} `yaml:"ca_certs_for_downloads,omitempty"`

	/*DiskCapacityMb - Descr: the container disk capacity the executor should manage.  this should not be greater than the actual disk quota on the VM Default: auto
*/
	DiskCapacityMb interface{} `yaml:"disk_capacity_mb,omitempty"`

	/*ReadWorkPoolSize - Descr: Maximum number of concurrent get container info operations. Default: 64
*/
	ReadWorkPoolSize interface{} `yaml:"read_work_pool_size,omitempty"`

	/*CreateWorkPoolSize - Descr: Maximum number of concurrent create container operations. Default: 32
*/
	CreateWorkPoolSize interface{} `yaml:"create_work_pool_size,omitempty"`

	/*GardenHealthcheck - Descr: Frequency for healtchecking garden Default: 10m
*/
	GardenHealthcheck *GardenHealthcheck `yaml:"garden_healthcheck,omitempty"`

	/*MetricsWorkPoolSize - Descr: Maximum number of concurrent get container metrics operations. Default: 8
*/
	MetricsWorkPoolSize interface{} `yaml:"metrics_work_pool_size,omitempty"`

}