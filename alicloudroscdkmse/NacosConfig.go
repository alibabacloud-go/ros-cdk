package alicloudroscdkmse

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkmse/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkmse/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `ALIYUN::MSE::NacosConfig`, which is used to create a Nacos configuration.
type NacosConfig interface {
	alicloudroscdkcore.Resource
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
	Props() *NacosConfigProps
	SetProps(val *NacosConfigProps)
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

// The jsii proxy struct for NacosConfig
type jsiiProxy_NacosConfig struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_NacosConfig) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Props() *NacosConfigProps {
	var returns *NacosConfigProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NacosConfig) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewNacosConfig(scope alicloudroscdkcore.Construct, id *string, props *NacosConfigProps, enableResourcePropertyConstraint *bool) NacosConfig {
	_init_.Initialize()

	if err := validateNewNacosConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NacosConfig{}

	_jsii_.Create(
		"@alicloud/ros-cdk-mse.NacosConfig",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewNacosConfig_Override(n NacosConfig, scope alicloudroscdkcore.Construct, id *string, props *NacosConfigProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-mse.NacosConfig",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		n,
	)
}

func (j *jsiiProxy_NacosConfig)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_NacosConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NacosConfig)SetProps(val *NacosConfigProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_NacosConfig)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_NacosConfig)SetScope(val alicloudroscdkcore.Construct) {
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
func NacosConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNacosConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-mse.NacosConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NacosConfig) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := n.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addCondition",
		[]interface{}{condition},
	)
}

func (n *jsiiProxy_NacosConfig) AddCount(count interface{}) {
	if err := n.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addCount",
		[]interface{}{count},
	)
}

func (n *jsiiProxy_NacosConfig) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := n.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addDependency",
		[]interface{}{resource},
	)
}

func (n *jsiiProxy_NacosConfig) AddResourceDesc(desc *string) {
	if err := n.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (n *jsiiProxy_NacosConfig) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NacosConfig) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NacosConfig) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := n.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		n,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NacosConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NacosConfig) OnSynthesize(session constructs.ISynthesisSession) {
	if err := n.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NacosConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NacosConfig) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NacosConfig) SetMetadata(key *string, value interface{}) {
	if err := n.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (n *jsiiProxy_NacosConfig) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := n.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NacosConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NacosConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

