package alicloudroscdkddospro


// Properties for defining a `SchedulerRule`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/aliyun-ddospro-schedulerrule
type SchedulerRuleProps struct {
	// Property ruleName: The name of the rule.
	RuleName interface{} `field:"required" json:"ruleName" yaml:"ruleName"`
	// Property rules: The details of the scheduling rule.
	//
	// This parameter is a JSON string. The following list describes the fields in the value of the parameter:
	// Type: the address type of the interaction resource that you want to use in the scheduling rule. This field is required and must be of the string type. Valid values: A: IP address. CNAME: domain name
	// Value: the address of the interaction resource that you want to use in the scheduling rule. This field is required and must be of the string type.
	// Priority: the priority of the scheduling rule. This field is required and must be of the integer type. Valid values: 0 to 100. A larger value indicates a higher priority.
	// ValueType: the type of the interaction resource that you want to use in the scheduling rule. This field is required and must be of the integer type. Valid values: 1: the IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance. 2: the IP address of the interaction resource in the tiered protection scenario. 3: the IP address that is used to accelerate access in the network acceleration scenario. 5: the domain name that is configured in Alibaba Cloud CDN (CDN) in the CDN interaction scenario. 6 the IP address of the interaction resource in the cloud service interaction scenario
	// RegionId: the region where the interaction resource is deployed. This parameter must be specified when ValueType is set to 2. The value must be of the string type.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Property ruleType: The type of the custom defense rule.
	//
	// Valid values:
	// 2: tiered protection
	// 3: network acceleration
	// 5: CDN interaction
	// 6: cloud service interaction
	// 8: secure acceleration.
	RuleType interface{} `field:"required" json:"ruleType" yaml:"ruleType"`
	// Property param: The details of the CDN interaction rule.
	//
	// This parameter is a JSON string. The following list describes the fields in the value of the parameter:
	// ParamType: the type of the scheduling rule. This field is required and must be of the string type. Set the value to cdn. This indicates that you want to modify a CDN interaction rule.
	// ParamData: the values of parameters that you want to modify for the CDN interaction rule. This field is required and must be of the map type. ParamData contains the following parameters: Domain: the accelerated domain in CDN. This parameter is required and must be of the string type; Cname: the CNAME that is assigned to the accelerated domain. This parameter is required and must be of the string type; AccessQps: the queries per second (QPS) threshold that is used to switch service traffic to Anti-DDoS Pro or Anti-DDoS Premium. This parameter is required and must be of the integer type; UpstreamQps: the QPS threshold that is used to switch service traffic to CDN. This parameter is optional and must be of the integer type.
	Param interface{} `field:"optional" json:"param" yaml:"param"`
	// Property resourceGroupId: The ID of the resource group to which the instance belongs in Resource Management.
	//
	// This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	ResourceGroupId interface{} `field:"optional" json:"resourceGroupId" yaml:"resourceGroupId"`
}

