package alicloudroscdkcen

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkcen/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcen/internal"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `ALIYUN::CEN::TransitRouterPeerAttachment`, which is used to create a cross-region connection for an Enterprise Edition transit router.
type TransitRouterPeerAttachment interface {
	alicloudroscdkcore.Resource
	// Attribute AutoPublishRouteEnabled: AutoPublishRouteEnabled.
	AttrAutoPublishRouteEnabled() alicloudroscdkcore.IResolvable
	// Attribute Bandwidth: Bandwidth.
	AttrBandwidth() alicloudroscdkcore.IResolvable
	// Attribute CenBandwidthPackageId: BandwidthPackageId.
	AttrCenBandwidthPackageId() alicloudroscdkcore.IResolvable
	// Attribute CenId: CenId.
	AttrCenId() alicloudroscdkcore.IResolvable
	// Attribute ClientToken: ClientToken.
	AttrClientToken() alicloudroscdkcore.IResolvable
	// Attribute GeographicSpanId: GeographicSpanId.
	AttrGeographicSpanId() alicloudroscdkcore.IResolvable
	// Attribute PeerTransitRouterId: PeerTransitRouterId.
	AttrPeerTransitRouterId() alicloudroscdkcore.IResolvable
	// Attribute PeerTransitRouterOwnerId: PeerTransitRouterOwnerId.
	AttrPeerTransitRouterOwnerId() alicloudroscdkcore.IResolvable
	// Attribute PeerTransitRouterRegionId: PeerTransitRouterRegionId.
	AttrPeerTransitRouterRegionId() alicloudroscdkcore.IResolvable
	// Attribute ResourceType: ResourceType.
	AttrResourceType() alicloudroscdkcore.IResolvable
	// Attribute TransitRouterAttachmentDescription: TransitRouterAttachmentDescription.
	AttrTransitRouterAttachmentDescription() alicloudroscdkcore.IResolvable
	// Attribute TransitRouterAttachmentId: The first ID of the resource.
	AttrTransitRouterAttachmentId() alicloudroscdkcore.IResolvable
	// Attribute TransitRouterAttachmentName: TransitRouterAttachmentName.
	AttrTransitRouterAttachmentName() alicloudroscdkcore.IResolvable
	// Attribute TransitRouterId: TransitRouterId.
	AttrTransitRouterId() alicloudroscdkcore.IResolvable
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
	Props() *TransitRouterPeerAttachmentProps
	SetProps(val *TransitRouterPeerAttachmentProps)
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

// The jsii proxy struct for TransitRouterPeerAttachment
type jsiiProxy_TransitRouterPeerAttachment struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrAutoPublishRouteEnabled() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrAutoPublishRouteEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrBandwidth() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrCenBandwidthPackageId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCenBandwidthPackageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrCenId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrClientToken() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrClientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrGeographicSpanId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrGeographicSpanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrPeerTransitRouterId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPeerTransitRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrPeerTransitRouterOwnerId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPeerTransitRouterOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrPeerTransitRouterRegionId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPeerTransitRouterRegionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrResourceType() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrTransitRouterAttachmentDescription() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTransitRouterAttachmentDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrTransitRouterAttachmentId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTransitRouterAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrTransitRouterAttachmentName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTransitRouterAttachmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) AttrTransitRouterId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTransitRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Props() *TransitRouterPeerAttachmentProps {
	var returns *TransitRouterPeerAttachmentProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitRouterPeerAttachment) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewTransitRouterPeerAttachment(scope alicloudroscdkcore.Construct, id *string, props *TransitRouterPeerAttachmentProps, enableResourcePropertyConstraint *bool) TransitRouterPeerAttachment {
	_init_.Initialize()

	if err := validateNewTransitRouterPeerAttachmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransitRouterPeerAttachment{}

	_jsii_.Create(
		"@alicloud/ros-cdk-cen.TransitRouterPeerAttachment",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewTransitRouterPeerAttachment_Override(t TransitRouterPeerAttachment, scope alicloudroscdkcore.Construct, id *string, props *TransitRouterPeerAttachmentProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-cen.TransitRouterPeerAttachment",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		t,
	)
}

func (j *jsiiProxy_TransitRouterPeerAttachment)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_TransitRouterPeerAttachment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TransitRouterPeerAttachment)SetProps(val *TransitRouterPeerAttachmentProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_TransitRouterPeerAttachment)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_TransitRouterPeerAttachment)SetScope(val alicloudroscdkcore.Construct) {
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
func TransitRouterPeerAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitRouterPeerAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-cen.TransitRouterPeerAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitRouterPeerAttachment) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := t.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addCondition",
		[]interface{}{condition},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) AddCount(count interface{}) {
	if err := t.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addCount",
		[]interface{}{count},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := t.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addDependency",
		[]interface{}{resource},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) AddResourceDesc(desc *string) {
	if err := t.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitRouterPeerAttachment) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := t.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		t,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitRouterPeerAttachment) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) OnSynthesize(session constructs.ISynthesisSession) {
	if err := t.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitRouterPeerAttachment) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) SetMetadata(key *string, value interface{}) {
	if err := t.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := t.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TransitRouterPeerAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitRouterPeerAttachment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

