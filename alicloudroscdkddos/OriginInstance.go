package alicloudroscdkddos

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkddos/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkddos/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `ALIYUN::DDoS::OriginInstance`, which is used to create an Anti-DDoS Origin instance.
type OriginInstance interface {
	alicloudroscdkcore.Resource
	// Attribute InstanceId: The ID of the instance.
	AttrInstanceId() alicloudroscdkcore.IResolvable
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
	Props() *OriginInstanceProps
	SetProps(val *OriginInstanceProps)
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

// The jsii proxy struct for OriginInstance
type jsiiProxy_OriginInstance struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_OriginInstance) AttrInstanceId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Props() *OriginInstanceProps {
	var returns *OriginInstanceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginInstance) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewOriginInstance(scope alicloudroscdkcore.Construct, id *string, props *OriginInstanceProps, enableResourcePropertyConstraint *bool) OriginInstance {
	_init_.Initialize()

	if err := validateNewOriginInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OriginInstance{}

	_jsii_.Create(
		"@alicloud/ros-cdk-ddos.OriginInstance",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewOriginInstance_Override(o OriginInstance, scope alicloudroscdkcore.Construct, id *string, props *OriginInstanceProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-ddos.OriginInstance",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		o,
	)
}

func (j *jsiiProxy_OriginInstance)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_OriginInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OriginInstance)SetProps(val *OriginInstanceProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_OriginInstance)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_OriginInstance)SetScope(val alicloudroscdkcore.Construct) {
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
func OriginInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-ddos.OriginInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginInstance) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := o.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addCondition",
		[]interface{}{condition},
	)
}

func (o *jsiiProxy_OriginInstance) AddCount(count interface{}) {
	if err := o.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addCount",
		[]interface{}{count},
	)
}

func (o *jsiiProxy_OriginInstance) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := o.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addDependency",
		[]interface{}{resource},
	)
}

func (o *jsiiProxy_OriginInstance) AddResourceDesc(desc *string) {
	if err := o.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (o *jsiiProxy_OriginInstance) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := o.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OriginInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginInstance) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := o.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		o,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		o,
		"onPrepare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginInstance) OnSynthesize(session constructs.ISynthesisSession) {
	if err := o.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (o *jsiiProxy_OriginInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginInstance) Prepare() {
	_jsii_.InvokeVoid(
		o,
		"prepare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginInstance) SetMetadata(key *string, value interface{}) {
	if err := o.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (o *jsiiProxy_OriginInstance) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := o.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"synthesize",
		[]interface{}{session},
	)
}

func (o *jsiiProxy_OriginInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

