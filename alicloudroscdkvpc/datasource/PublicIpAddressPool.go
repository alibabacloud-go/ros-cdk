package datasource

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkvpc/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkvpc/datasource/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `DATASOURCE::VPC::PublicIpAddressPool`.
type PublicIpAddressPool interface {
	alicloudroscdkcore.Resource
	// Attribute CreateTime: The creation time of the resource.
	AttrCreateTime() alicloudroscdkcore.IResolvable
	// Attribute Description: Description.
	AttrDescription() alicloudroscdkcore.IResolvable
	// Attribute IpAddressRemaining: Whether there is a free IP address.
	AttrIpAddressRemaining() alicloudroscdkcore.IResolvable
	// Attribute Isp: Service providers.
	AttrIsp() alicloudroscdkcore.IResolvable
	// Attribute PublicIpAddressPoolId: The first ID of the resource.
	AttrPublicIpAddressPoolId() alicloudroscdkcore.IResolvable
	// Attribute PublicIpAddressPoolName: The name of the resource.
	AttrPublicIpAddressPoolName() alicloudroscdkcore.IResolvable
	// Attribute ResourceGroupId: The ID of the resource group to which the VPC belongs.
	AttrResourceGroupId() alicloudroscdkcore.IResolvable
	// Attribute Tags: The tags of PrefixList.
	AttrTags() alicloudroscdkcore.IResolvable
	// Attribute TotalIpNum: Total ip number of PublicIpAddressPool.
	AttrTotalIpNum() alicloudroscdkcore.IResolvable
	// Attribute UsedIpNum: Used ip number of PublicIpAddressPool.
	AttrUsedIpNum() alicloudroscdkcore.IResolvable
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
	Props() *PublicIpAddressPoolProps
	SetProps(val *PublicIpAddressPoolProps)
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

// The jsii proxy struct for PublicIpAddressPool
type jsiiProxy_PublicIpAddressPool struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_PublicIpAddressPool) AttrCreateTime() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrDescription() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrIpAddressRemaining() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrIpAddressRemaining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrIsp() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrIsp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrPublicIpAddressPoolId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPublicIpAddressPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrPublicIpAddressPoolName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPublicIpAddressPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrResourceGroupId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrTags() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrTotalIpNum() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTotalIpNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) AttrUsedIpNum() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrUsedIpNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Props() *PublicIpAddressPoolProps {
	var returns *PublicIpAddressPoolProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicIpAddressPool) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewPublicIpAddressPool(scope alicloudroscdkcore.Construct, id *string, props *PublicIpAddressPoolProps, enableResourcePropertyConstraint *bool) PublicIpAddressPool {
	_init_.Initialize()

	if err := validateNewPublicIpAddressPoolParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PublicIpAddressPool{}

	_jsii_.Create(
		"@alicloud/ros-cdk-vpc.datasource.PublicIpAddressPool",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewPublicIpAddressPool_Override(p PublicIpAddressPool, scope alicloudroscdkcore.Construct, id *string, props *PublicIpAddressPoolProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-vpc.datasource.PublicIpAddressPool",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		p,
	)
}

func (j *jsiiProxy_PublicIpAddressPool)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_PublicIpAddressPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PublicIpAddressPool)SetProps(val *PublicIpAddressPoolProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_PublicIpAddressPool)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_PublicIpAddressPool)SetScope(val alicloudroscdkcore.Construct) {
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
func PublicIpAddressPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublicIpAddressPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-vpc.datasource.PublicIpAddressPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicIpAddressPool) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := p.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCondition",
		[]interface{}{condition},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) AddCount(count interface{}) {
	if err := p.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCount",
		[]interface{}{count},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := p.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{resource},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) AddResourceDesc(desc *string) {
	if err := p.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicIpAddressPool) GetAtt(name *string) alicloudroscdkcore.IResolvable {
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

func (p *jsiiProxy_PublicIpAddressPool) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublicIpAddressPool) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicIpAddressPool) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublicIpAddressPool) SetMetadata(key *string, value interface{}) {
	if err := p.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PublicIpAddressPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicIpAddressPool) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

