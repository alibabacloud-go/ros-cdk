package datasource

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkpai/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkpai/datasource/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `DATASOURCE::PAI::WorkspaceResourceFlinks`.
type WorkspaceResourceFlinks interface {
	alicloudroscdkcore.Resource
	// Attribute GroupNames: The list of group names.
	AttrGroupNames() alicloudroscdkcore.IResolvable
	// Attribute WorkspaceResourceFlinks: The list of workspace resource flinks.
	AttrWorkspaceResourceFlinks() alicloudroscdkcore.IResolvable
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
	Props() *WorkspaceResourceFlinksProps
	SetProps(val *WorkspaceResourceFlinksProps)
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

// The jsii proxy struct for WorkspaceResourceFlinks
type jsiiProxy_WorkspaceResourceFlinks struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_WorkspaceResourceFlinks) AttrGroupNames() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) AttrWorkspaceResourceFlinks() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrWorkspaceResourceFlinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Props() *WorkspaceResourceFlinksProps {
	var returns *WorkspaceResourceFlinksProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceResourceFlinks) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewWorkspaceResourceFlinks(scope alicloudroscdkcore.Construct, id *string, props *WorkspaceResourceFlinksProps, enableResourcePropertyConstraint *bool) WorkspaceResourceFlinks {
	_init_.Initialize()

	if err := validateNewWorkspaceResourceFlinksParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceResourceFlinks{}

	_jsii_.Create(
		"@alicloud/ros-cdk-pai.datasource.WorkspaceResourceFlinks",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewWorkspaceResourceFlinks_Override(w WorkspaceResourceFlinks, scope alicloudroscdkcore.Construct, id *string, props *WorkspaceResourceFlinksProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-pai.datasource.WorkspaceResourceFlinks",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		w,
	)
}

func (j *jsiiProxy_WorkspaceResourceFlinks)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_WorkspaceResourceFlinks)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkspaceResourceFlinks)SetProps(val *WorkspaceResourceFlinksProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_WorkspaceResourceFlinks)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceResourceFlinks)SetScope(val alicloudroscdkcore.Construct) {
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
func WorkspaceResourceFlinks_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceResourceFlinks_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-pai.datasource.WorkspaceResourceFlinks",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceResourceFlinks) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := w.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addCondition",
		[]interface{}{condition},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) AddCount(count interface{}) {
	if err := w.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addCount",
		[]interface{}{count},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := w.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addDependency",
		[]interface{}{resource},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) AddResourceDesc(desc *string) {
	if err := w.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := w.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceResourceFlinks) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := w.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		w,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceResourceFlinks) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) OnSynthesize(session constructs.ISynthesisSession) {
	if err := w.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceResourceFlinks) Prepare() {
	_jsii_.InvokeVoid(
		w,
		"prepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) SetMetadata(key *string, value interface{}) {
	if err := w.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := w.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"synthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_WorkspaceResourceFlinks) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceResourceFlinks) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

