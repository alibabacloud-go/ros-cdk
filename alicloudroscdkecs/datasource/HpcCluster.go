package datasource

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkecs/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkecs/datasource/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `DATASOURCE::ECS::HpcCluster`.
type HpcCluster interface {
	alicloudroscdkcore.Resource
	// Attribute Description: The description of the HPC cluster.
	AttrDescription() alicloudroscdkcore.IResolvable
	// Attribute HpcClusterId: The ID of cluster.
	AttrHpcClusterId() alicloudroscdkcore.IResolvable
	// Attribute HpcClusterName: The name of the HPC cluster.
	AttrHpcClusterName() alicloudroscdkcore.IResolvable
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
	Props() *HpcClusterProps
	SetProps(val *HpcClusterProps)
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

// The jsii proxy struct for HpcCluster
type jsiiProxy_HpcCluster struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_HpcCluster) AttrDescription() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) AttrHpcClusterId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrHpcClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) AttrHpcClusterName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrHpcClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Props() *HpcClusterProps {
	var returns *HpcClusterProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCluster) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewHpcCluster(scope alicloudroscdkcore.Construct, id *string, props *HpcClusterProps, enableResourcePropertyConstraint *bool) HpcCluster {
	_init_.Initialize()

	if err := validateNewHpcClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HpcCluster{}

	_jsii_.Create(
		"@alicloud/ros-cdk-ecs.datasource.HpcCluster",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewHpcCluster_Override(h HpcCluster, scope alicloudroscdkcore.Construct, id *string, props *HpcClusterProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-ecs.datasource.HpcCluster",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		h,
	)
}

func (j *jsiiProxy_HpcCluster)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_HpcCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HpcCluster)SetProps(val *HpcClusterProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_HpcCluster)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_HpcCluster)SetScope(val alicloudroscdkcore.Construct) {
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
func HpcCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-ecs.datasource.HpcCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCluster) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := h.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addCondition",
		[]interface{}{condition},
	)
}

func (h *jsiiProxy_HpcCluster) AddCount(count interface{}) {
	if err := h.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addCount",
		[]interface{}{count},
	)
}

func (h *jsiiProxy_HpcCluster) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := h.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addDependency",
		[]interface{}{resource},
	)
}

func (h *jsiiProxy_HpcCluster) AddResourceDesc(desc *string) {
	if err := h.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (h *jsiiProxy_HpcCluster) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := h.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (h *jsiiProxy_HpcCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCluster) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := h.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		h,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := h.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_HpcCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCluster) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCluster) SetMetadata(key *string, value interface{}) {
	if err := h.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (h *jsiiProxy_HpcCluster) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := h.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_HpcCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

