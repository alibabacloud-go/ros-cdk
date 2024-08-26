package datasource


// Properties for defining a `Instances`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/datasource-cr-instances
type InstancesProps struct {
	// Property instanceName: InstanceName.
	InstanceName interface{} `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Property refreshOptions: The refresh strategy for the datasource resource when the stack is updated.
	//
	// Valid values:
	// - Never: Never refresh the datasource resource when the stack is updated.
	// - Always: Always refresh the datasource resource when the stack is updated.
	// Default is Never.
	RefreshOptions interface{} `field:"optional" json:"refreshOptions" yaml:"refreshOptions"`
	// Property resourceGroupId: The ID of the resource group.
	ResourceGroupId interface{} `field:"optional" json:"resourceGroupId" yaml:"resourceGroupId"`
}

