package datasource


// Properties for defining a `WorkspaceResourceMaxComputes`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/datasource-pai-workspaceresourcemaxcomputes
type WorkspaceResourceMaxComputesProps struct {
	// Property workspaceId: The ID of the workspace to which the workspace belongs.
	WorkspaceId interface{} `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Property groupName: Resource group name.
	//
	// If you want to obtain a resource group name.
	GroupName interface{} `field:"optional" json:"groupName" yaml:"groupName"`
	// Property refreshOptions: The refresh strategy for the datasource resource when the stack is updated.
	//
	// Valid values:
	// - Never: Never refresh the datasource resource when the stack is updated.
	// - Always: Always refresh the datasource resource when the stack is updated.
	// Default is Never.
	RefreshOptions interface{} `field:"optional" json:"refreshOptions" yaml:"refreshOptions"`
}

