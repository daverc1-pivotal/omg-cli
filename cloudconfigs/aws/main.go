package main

import (
	"github.com/enaml-ops/omg-cli/cloudconfigs/aws/plugin"
	"github.com/enaml-ops/omg-cli/pluginlib/cloudconfig"
)

func main() {
	cloudconfig.Run(new(awsccplugin.AWSCloudConfig))
}
