package alicloudroscdkalb


// Properties for defining a `RosHealthCheckTemplate`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/aliyun-alb-healthchecktemplate
type RosHealthCheckTemplateProps struct {
	HealthCheckTemplateName interface{} `field:"required" json:"healthCheckTemplateName" yaml:"healthCheckTemplateName"`
	HealthCheckCodes interface{} `field:"optional" json:"healthCheckCodes" yaml:"healthCheckCodes"`
	HealthCheckConnectPort interface{} `field:"optional" json:"healthCheckConnectPort" yaml:"healthCheckConnectPort"`
	HealthCheckHost interface{} `field:"optional" json:"healthCheckHost" yaml:"healthCheckHost"`
	HealthCheckInterval interface{} `field:"optional" json:"healthCheckInterval" yaml:"healthCheckInterval"`
	HealthCheckMethod interface{} `field:"optional" json:"healthCheckMethod" yaml:"healthCheckMethod"`
	HealthCheckPath interface{} `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	HealthCheckProtocol interface{} `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	HealthCheckTimeout interface{} `field:"optional" json:"healthCheckTimeout" yaml:"healthCheckTimeout"`
	HealthyThreshold interface{} `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	UnhealthyThreshold interface{} `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

