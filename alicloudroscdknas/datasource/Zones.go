package datasource

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdknas/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdknas/datasource/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `DATASOURCE::NAS::Zones`, which is used to query zones.
type Zones interface {
	alicloudroscdkcore.Resource
	// Attribute ZoneIds: The list of zone IDs.
	AttrZoneIds() alicloudroscdkcore.IResolvable
	// Attribute Zones: The list of zones.
	AttrZones() alicloudroscdkcore.IResolvable
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
	Props() *ZonesProps
	SetProps(val *ZonesProps)
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

// The jsii proxy struct for Zones
type jsiiProxy_Zones struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_Zones) AttrZoneIds() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrZoneIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) AttrZones() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Props() *ZonesProps {
	var returns *ZonesProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zones) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewZones(scope alicloudroscdkcore.Construct, id *string, props *ZonesProps, enableResourcePropertyConstraint *bool) Zones {
	_init_.Initialize()

	if err := validateNewZonesParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Zones{}

	_jsii_.Create(
		"@alicloud/ros-cdk-nas.datasource.Zones",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewZones_Override(z Zones, scope alicloudroscdkcore.Construct, id *string, props *ZonesProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-nas.datasource.Zones",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		z,
	)
}

func (j *jsiiProxy_Zones)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_Zones)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Zones)SetProps(val *ZonesProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_Zones)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_Zones)SetScope(val alicloudroscdkcore.Construct) {
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
func Zones_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZones_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-nas.datasource.Zones",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_Zones) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := z.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addCondition",
		[]interface{}{condition},
	)
}

func (z *jsiiProxy_Zones) AddCount(count interface{}) {
	if err := z.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addCount",
		[]interface{}{count},
	)
}

func (z *jsiiProxy_Zones) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := z.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addDependency",
		[]interface{}{resource},
	)
}

func (z *jsiiProxy_Zones) AddResourceDesc(desc *string) {
	if err := z.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (z *jsiiProxy_Zones) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := z.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (z *jsiiProxy_Zones) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_Zones) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := z.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		z,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_Zones) OnPrepare() {
	_jsii_.InvokeVoid(
		z,
		"onPrepare",
		nil, // no parameters
	)
}

func (z *jsiiProxy_Zones) OnSynthesize(session constructs.ISynthesisSession) {
	if err := z.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (z *jsiiProxy_Zones) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_Zones) Prepare() {
	_jsii_.InvokeVoid(
		z,
		"prepare",
		nil, // no parameters
	)
}

func (z *jsiiProxy_Zones) SetMetadata(key *string, value interface{}) {
	if err := z.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (z *jsiiProxy_Zones) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := z.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"synthesize",
		[]interface{}{session},
	)
}

func (z *jsiiProxy_Zones) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_Zones) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

