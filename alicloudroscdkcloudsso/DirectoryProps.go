package alicloudroscdkcloudsso


// Properties for defining a `Directory`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/aliyun-cloudsso-directory
type DirectoryProps struct {
	// Property directoryName: The name of the directory.
	//
	// The name must be globally unique.
	// The name can contain lowercase letters, digits, or hyphens (-). The name cannot start or end with a hyphen (-) and cannot contain two consecutive hyphens (-). The name cannot start with d-.
	// The name must be 2 to 64 characters in length.
	// **Note**: If you do not specify this parameter, the value of this parameter is automatically generated by the system.
	DirectoryName interface{} `field:"optional" json:"directoryName" yaml:"directoryName"`
}

