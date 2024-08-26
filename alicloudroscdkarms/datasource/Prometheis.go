package datasource

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkarms/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkarms/datasource/internal"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `DATASOURCE::ARMS::Prometheis`, which is used to query Prometheus instances.
type Prometheis interface {
	alicloudroscdkcore.Resource
	// Attribute ClusterIds: The list of cluster IDs.
	AttrClusterIds() alicloudroscdkcore.IResolvable
	// Attribute Prometheis: The list of Prometheis.
	AttrPrometheis() alicloudroscdkcore.IResolvable
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
	Props() *PrometheisProps
	SetProps(val *PrometheisProps)
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

// The jsii proxy struct for Prometheis
type jsiiProxy_Prometheis struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_Prometheis) AttrClusterIds() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrClusterIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) AttrPrometheis() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPrometheis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Props() *PrometheisProps {
	var returns *PrometheisProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheis) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewPrometheis(scope alicloudroscdkcore.Construct, id *string, props *PrometheisProps, enableResourcePropertyConstraint *bool) Prometheis {
	_init_.Initialize()

	if err := validateNewPrometheisParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Prometheis{}

	_jsii_.Create(
		"@alicloud/ros-cdk-arms.datasource.Prometheis",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewPrometheis_Override(p Prometheis, scope alicloudroscdkcore.Construct, id *string, props *PrometheisProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-arms.datasource.Prometheis",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		p,
	)
}

func (j *jsiiProxy_Prometheis)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_Prometheis)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Prometheis)SetProps(val *PrometheisProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_Prometheis)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_Prometheis)SetScope(val alicloudroscdkcore.Construct) {
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
func Prometheis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrometheis_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-arms.datasource.Prometheis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheis) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := p.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCondition",
		[]interface{}{condition},
	)
}

func (p *jsiiProxy_Prometheis) AddCount(count interface{}) {
	if err := p.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCount",
		[]interface{}{count},
	)
}

func (p *jsiiProxy_Prometheis) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := p.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{resource},
	)
}

func (p *jsiiProxy_Prometheis) AddResourceDesc(desc *string) {
	if err := p.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (p *jsiiProxy_Prometheis) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_Prometheis) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheis) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := p.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		p,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheis) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prometheis) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Prometheis) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheis) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prometheis) SetMetadata(key *string, value interface{}) {
	if err := p.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (p *jsiiProxy_Prometheis) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Prometheis) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheis) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

