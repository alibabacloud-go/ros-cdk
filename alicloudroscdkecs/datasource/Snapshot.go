package datasource

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkecs/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkecs/datasource/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `DATASOURCE::ECS::Snapshot`.
type Snapshot interface {
	alicloudroscdkcore.Resource
	// Attribute CreateTime: The create time.
	AttrCreateTime() alicloudroscdkcore.IResolvable
	// Attribute Description: The description.
	AttrDescription() alicloudroscdkcore.IResolvable
	// Attribute DiskId: The source disk id.
	AttrDiskId() alicloudroscdkcore.IResolvable
	// Attribute Encrypted: Whether the snapshot is encrypted.
	AttrEncrypted() alicloudroscdkcore.IResolvable
	// Attribute InstantAccess: Whether snapshot speed availability is enabled.
	AttrInstantAccess() alicloudroscdkcore.IResolvable
	// Attribute InstantAccessRetentionDays: InstantAccessRetentionDays.
	AttrInstantAccessRetentionDays() alicloudroscdkcore.IResolvable
	// Attribute ProductCode: The product number inherited from the mirror market.
	AttrProductCode() alicloudroscdkcore.IResolvable
	// Attribute Progress: Snapshot creation progress, in percentage.
	AttrProgress() alicloudroscdkcore.IResolvable
	// Attribute RemainTime: Remaining completion time for the snapshot being created.
	AttrRemainTime() alicloudroscdkcore.IResolvable
	// Attribute ResourceGroupId: The resource group id.
	AttrResourceGroupId() alicloudroscdkcore.IResolvable
	// Attribute RetentionDays: Automatic snapshot retention days.
	AttrRetentionDays() alicloudroscdkcore.IResolvable
	// Attribute SnapshotId: The snapshot id.
	AttrSnapshotId() alicloudroscdkcore.IResolvable
	// Attribute SnapshotName: Snapshot Display Name.
	AttrSnapshotName() alicloudroscdkcore.IResolvable
	// Attribute SnapshotSN: Snapshot serial number.
	AttrSnapshotSn() alicloudroscdkcore.IResolvable
	// Attribute SnapshotType: Snapshot creation type.
	AttrSnapshotType() alicloudroscdkcore.IResolvable
	// Attribute SourceDiskSize: Source disk capacity.
	AttrSourceDiskSize() alicloudroscdkcore.IResolvable
	// Attribute SourceDiskType: Source disk attributes.
	AttrSourceDiskType() alicloudroscdkcore.IResolvable
	// Attribute SourceStorageType: Original disk type.
	AttrSourceStorageType() alicloudroscdkcore.IResolvable
	// Attribute Tags: The tags.
	AttrTags() alicloudroscdkcore.IResolvable
	// Attribute Usage: A resource type that has a reference relationship.
	AttrUsage() alicloudroscdkcore.IResolvable
	EnableResourcePropertyConstraint() *bool
	SetEnableResourcePropertyConstraint(val *bool)
	Id() *string
	SetId(val *string)
	// The construct tree node associated with this construct.
	Node() alicloudroscdkcore.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the ROS resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by ROS
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	Props() *SnapshotProps
	SetProps(val *SnapshotProps)
	Ref() *string
	Resource() alicloudroscdkcore.RosResource
	SetResource(val alicloudroscdkcore.RosResource)
	Scope() alicloudroscdkcore.Construct
	SetScope(val alicloudroscdkcore.Construct)
	// The stack in which this resource is defined.
	Stack() alicloudroscdkcore.Stack
	AddCondition(condition alicloudroscdkcore.RosCondition)
	AddCount(count interface{})
	AddDependency(resource alicloudroscdkcore.Resource)
	AddResourceDesc(desc *string)
	ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy)
	GeneratePhysicalName() *string
	GetAtt(name *string) alicloudroscdkcore.IResolvable
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	Prepare()
	SetMetadata(key *string, value interface{})
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	Synthesize(session alicloudroscdkcore.ISynthesisSession)
	// Returns a string representation of this construct.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	Validate() *[]*string
}

// The jsii proxy struct for Snapshot
type jsiiProxy_Snapshot struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_Snapshot) AttrCreateTime() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrDescription() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrDiskId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrEncrypted() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrInstantAccess() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrInstantAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrInstantAccessRetentionDays() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrInstantAccessRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrProductCode() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrProductCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrProgress() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrRemainTime() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRemainTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrResourceGroupId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrRetentionDays() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSnapshotId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSnapshotName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSnapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSnapshotSn() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSnapshotSn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSnapshotType() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSnapshotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSourceDiskSize() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSourceDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSourceDiskType() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSourceDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrSourceStorageType() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSourceStorageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrTags() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) AttrUsage() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Props() *SnapshotProps {
	var returns *SnapshotProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Snapshot) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewSnapshot(scope alicloudroscdkcore.Construct, id *string, props *SnapshotProps, enableResourcePropertyConstraint *bool) Snapshot {
	_init_.Initialize()

	if err := validateNewSnapshotParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Snapshot{}

	_jsii_.Create(
		"@alicloud/ros-cdk-ecs.datasource.Snapshot",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewSnapshot_Override(s Snapshot, scope alicloudroscdkcore.Construct, id *string, props *SnapshotProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-ecs.datasource.Snapshot",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		s,
	)
}

func (j *jsiiProxy_Snapshot)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_Snapshot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Snapshot)SetProps(val *SnapshotProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_Snapshot)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_Snapshot)SetScope(val alicloudroscdkcore.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

// Return whether the given object is a Construct.
func Snapshot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnapshot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-ecs.datasource.Snapshot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Snapshot) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := s.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addCondition",
		[]interface{}{condition},
	)
}

func (s *jsiiProxy_Snapshot) AddCount(count interface{}) {
	if err := s.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addCount",
		[]interface{}{count},
	)
}

func (s *jsiiProxy_Snapshot) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := s.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDependency",
		[]interface{}{resource},
	)
}

func (s *jsiiProxy_Snapshot) AddResourceDesc(desc *string) {
	if err := s.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (s *jsiiProxy_Snapshot) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Snapshot) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Snapshot) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := s.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		s,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Snapshot) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Snapshot) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Snapshot) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Snapshot) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Snapshot) SetMetadata(key *string, value interface{}) {
	if err := s.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_Snapshot) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Snapshot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Snapshot) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

