package datasource


// Properties for defining a `Namespaces`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/datasource-cr-namespaces
type NamespacesProps struct {
	// Property refreshOptions: The refresh strategy for the datasource resource when the stack is updated.
	//
	// Valid values:
	// - Never: Never refresh the datasource resource when the stack is updated.
	// - Always: Always refresh the datasource resource when the stack is updated.
	// Default is Never.
	RefreshOptions interface{} `field:"optional" json:"refreshOptions" yaml:"refreshOptions"`
	// Property status: The status of namespace.
	Status interface{} `field:"optional" json:"status" yaml:"status"`
}

