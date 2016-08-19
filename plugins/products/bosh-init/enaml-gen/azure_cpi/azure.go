package azure_cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Azure struct {

	/*DebugMode - Descr: Enable debug mode to log all raw HTTP requests/responses Default: false
*/
	DebugMode interface{} `yaml:"debug_mode,omitempty"`

	/*ClientId - Descr: The client id for your service principal Default: <nil>
*/
	ClientId interface{} `yaml:"client_id,omitempty"`

	/*ResourceGroupName - Descr: Resource group name to use when spinning up new vms Default: <nil>
*/
	ResourceGroupName interface{} `yaml:"resource_group_name,omitempty"`

	/*StorageAccountName - Descr: Azure storage account name Default: <nil>
*/
	StorageAccountName interface{} `yaml:"storage_account_name,omitempty"`

	/*ParallelUploadThreadNum - Descr: The number of threads to upload stemcells in parallel Default: 16
*/
	ParallelUploadThreadNum interface{} `yaml:"parallel_upload_thread_num,omitempty"`

	/*SshUser - Descr: Default ssh user to use when spinning up new vms Default: vcap
*/
	SshUser interface{} `yaml:"ssh_user,omitempty"`

	/*AzureStackDomain - Descr: The domain for your AzureStack deployment Default: <nil>
*/
	AzureStackDomain interface{} `yaml:"azure_stack_domain,omitempty"`

	/*Environment - Descr: The environment for Azure Management Service, AzureCloud, AzureChinaCloud or AzureStack Default: AzureCloud
*/
	Environment interface{} `yaml:"environment,omitempty"`

	/*SubscriptionId - Descr: Azure Subscription Id Default: <nil>
*/
	SubscriptionId interface{} `yaml:"subscription_id,omitempty"`

	/*TenantId - Descr: The tenant id for your service principal Default: <nil>
*/
	TenantId interface{} `yaml:"tenant_id,omitempty"`

	/*SshPublicKey - Descr: The content of the SSH public key to use when spinning up new vms Default: <nil>
*/
	SshPublicKey interface{} `yaml:"ssh_public_key,omitempty"`

	/*AzureStackAuthentication - Descr: The authentication type for your AzureStack deployment. AzureAD, AzureStackAD or AzureStack Default: <nil>
*/
	AzureStackAuthentication interface{} `yaml:"azure_stack_authentication,omitempty"`

	/*ClientSecret - Descr: The client secret for your service principal Default: <nil>
*/
	ClientSecret interface{} `yaml:"client_secret,omitempty"`

	/*DefaultSecurityGroup - Descr: The name of the default security group that will be applied to all created VMs Default: <nil>
*/
	DefaultSecurityGroup interface{} `yaml:"default_security_group,omitempty"`

}