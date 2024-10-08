package alicloudroscdkcms


// Properties for defining a `GroupMetricRule`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/aliyun-cms-groupmetricrule
type GroupMetricRuleProps struct {
	// Property category: The abbreviation of the service name.
	//
	// Valid values:
	// ECS (including Alibaba Cloud and non-Alibaba Cloud hosts)
	// RDS (ApsaraDB for RDS)
	// ADS (AnalyticDB)
	// SLB (Server Load Balancer)
	// VPC (Virtual Private Cloud)
	// APIGATEWAY (API Gateway)
	// CDN
	// CS (Container Service for Swarm)
	// DCDN (Dynamic Route for CDN)
	// DDoS (distributed denial of service)
	// EIP (Elastic IP)
	// ELASTICSEARCH (Elasticsearch)
	// EMR (E-MapReduce)
	// ESS (Auto Scaling)
	// HBASE (ApsaraDB for HBase)
	// IOT_EDGE (IoT Edge)
	// K8S_POD (k8s pod)
	// KVSTORE_SHARDING (ApsaraDB for Redis cluster version)
	// KVSTORE_SPLITRW (ApsaraDB for Redis read\/write splitting version)
	// KVSTORE_STANDARD (ApsaraDB for Redis standard version)
	// MEMCACHE (ApsaraDB for Memcache)
	// MNS (Message Service)
	// MONGODB (ApsaraDB for MongoDB replica set instances)
	// MONGODB_CLUSTER (ApsaraDB for MongoDB cluster version)
	// MONGODB_SHARDING (ApsaraDB for MongoDB sharded clusters)
	// MQ_TOPIC (Message Service topic)
	// OCS (original version of ApsaraDB for Memcache)
	// OPENSEARCH (Open Search)
	// OSS (Object Storage Service)
	// POLARDB (ApsaraDB for POLARDB)
	// PETADATA (HybridDB for MySQL)
	// SCDN (Secure Content Delivery Network)
	// SHAREBANDWIDTHPACKAGES (shared bandwidth package)
	// SLS (Log Service)
	// VPN (VPN Gateway).
	Category interface{} `field:"required" json:"category" yaml:"category"`
	// Property escalations: undefined.
	Escalations interface{} `field:"required" json:"escalations" yaml:"escalations"`
	// Property groupId: The ID of application group.
	GroupId interface{} `field:"required" json:"groupId" yaml:"groupId"`
	// Property metricName: The name of the metric.
	//
	// For more information, call DescribeMetricMetaList or see Preset metrics reference.
	MetricName interface{} `field:"required" json:"metricName" yaml:"metricName"`
	// Property namespace: The data namespace of the service.
	//
	// For more information, call DescribeMetricMetaList
	// or see Preset metrics reference.
	Namespace interface{} `field:"required" json:"namespace" yaml:"namespace"`
	// Property ruleId: The ID of the alert rule.
	//
	// The IDs of alert rules are generated by callers to ensure
	// uniqueness.
	RuleId interface{} `field:"required" json:"ruleId" yaml:"ruleId"`
	// Property ruleName: The name of the alert rule.
	RuleName interface{} `field:"required" json:"ruleName" yaml:"ruleName"`
	// Property dimensions: The expended resource dimensions.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Property effectiveInterval: The period when the alert rule is effective.
	EffectiveInterval interface{} `field:"optional" json:"effectiveInterval" yaml:"effectiveInterval"`
	// Property emailSubject: The subject of the alert notification email.
	EmailSubject interface{} `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// Property interval: The detection period of alerts.
	Interval interface{} `field:"optional" json:"interval" yaml:"interval"`
	// Property noEffectiveInterval: The period when the alert rule is ineffective.
	NoEffectiveInterval interface{} `field:"optional" json:"noEffectiveInterval" yaml:"noEffectiveInterval"`
	// Property period: The aggregation period.
	//
	// Unite: second.
	Period interface{} `field:"optional" json:"period" yaml:"period"`
	// Property silenceTime: The duration of the mute period during which new alerts are not sent even if the trigger conditions are met.
	//
	// Unit: second. Default value: 86400. Minimum value: 60.
	SilenceTime interface{} `field:"optional" json:"silenceTime" yaml:"silenceTime"`
	// Property webhook: The URL of the callback triggered when an alert occurs.
	Webhook interface{} `field:"optional" json:"webhook" yaml:"webhook"`
}

