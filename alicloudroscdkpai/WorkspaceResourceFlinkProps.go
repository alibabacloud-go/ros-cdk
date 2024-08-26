package alicloudroscdkpai


// Properties for defining a `WorkspaceResourceFlink`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/aliyun-pai-workspaceresourceflink
type WorkspaceResourceFlinkProps struct {
	// Property groupName: Resource group name.
	//
	// If you want to obtain a resource group name, see [ListResources].
	GroupName interface{} `field:"required" json:"groupName" yaml:"groupName"`
	// Property resources: Resource List.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// Property workspaceId: The ID of the workspace to which the workspace belongs.
	WorkspaceId interface{} `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Property isDefault: Indicates whether it is the default resource.
	//
	// Currently, this parameter only supports the input of true and does not support false.
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Property option: Create behavior that supports the following values: - CreateAndAttach: Create resource and bind to workspace - Attach: bind resource to workspace.
	Option interface{} `field:"optional" json:"option" yaml:"option"`
}

