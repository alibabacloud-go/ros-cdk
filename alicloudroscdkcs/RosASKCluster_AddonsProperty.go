package alicloudroscdkcs


type RosASKCluster_AddonsProperty struct {
	Name interface{} `field:"required" json:"name" yaml:"name"`
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

