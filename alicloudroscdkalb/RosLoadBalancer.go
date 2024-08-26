package alicloudroscdkalb

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkalb/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkalb/internal"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class is a base encapsulation around the ROS resource type `ALIYUN::ALB::LoadBalancer`, which is used to create an Application Load Balancer (ALB) instance.
type RosLoadBalancer interface {
	alicloudroscdkcore.RosResource
	AccessLogConfig() interface{}
	SetAccessLogConfig(val interface{})
	AddressAllocatedMode() interface{}
	SetAddressAllocatedMode(val interface{})
	AddressIpVersion() interface{}
	SetAddressIpVersion(val interface{})
	AddressType() interface{}
	SetAddressType(val interface{})
	AttrAddressType() alicloudroscdkcore.IResolvable
	AttrDnsName() alicloudroscdkcore.IResolvable
	AttrLoadBalancerEdition() alicloudroscdkcore.IResolvable
	AttrLoadBalancerId() alicloudroscdkcore.IResolvable
	AttrVpcId() alicloudroscdkcore.IResolvable
	AttrZoneMappings() alicloudroscdkcore.IResolvable
	BandwidthPackageId() interface{}
	SetBandwidthPackageId(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aliyun:ros:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	EnableResourcePropertyConstraint() *bool
	SetEnableResourcePropertyConstraint(val *bool)
	LoadBalancerBillingConfig() interface{}
	SetLoadBalancerBillingConfig(val interface{})
	LoadBalancerEdition() interface{}
	SetLoadBalancerEdition(val interface{})
	LoadBalancerName() interface{}
	SetLoadBalancerName(val interface{})
	// The logical ID for this stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	ModificationProtectionConfig() interface{}
	SetModificationProtectionConfig(val interface{})
	// The construct tree node associated with this construct.
	Node() alicloudroscdkcore.ConstructNode
	// Return a string that will be resolved to a RosTemplate `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	ResourceGroupId() interface{}
	SetResourceGroupId(val interface{})
	// Options for this resource, such as condition, update policy etc.
	RosOptions() alicloudroscdkcore.IRosResourceOptions
	RosProperties() *map[string]interface{}
	// ROS resource type.
	RosResourceType() *string
	SecurityGroupIds() interface{}
	SetSecurityGroupIds(val interface{})
	// The stack in which this element is defined.
	//
	// RosElements must be defined within a stack scope (directly or indirectly).
	Stack() alicloudroscdkcore.Stack
	Tags() *[]*RosLoadBalancer_TagsProperty
	SetTags(val *[]*RosLoadBalancer_TagsProperty)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	VpcId() interface{}
	SetVpcId(val interface{})
	ZoneMappings() interface{}
	SetZoneMappings(val interface{})
	AddCondition(con alicloudroscdkcore.RosCondition)
	AddCount(count interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target alicloudroscdkcore.RosResource)
	AddDesc(desc *string)
	AddMetaData(key *string, value interface{})
	// Adds an override to the synthesized ROS resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// For example,
	// ```typescript
	// addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute'])
	// addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE')
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	AddRosDependency(target *string)
	// Sets the deletion policy of the resource based on the removal policy specified.
	ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy, options *alicloudroscdkcore.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) alicloudroscdkcore.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	Synthesize(session alicloudroscdkcore.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for RosLoadBalancer
type jsiiProxy_RosLoadBalancer struct {
	internal.Type__alicloudroscdkcoreRosResource
}

func (j *jsiiProxy_RosLoadBalancer) AccessLogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AddressAllocatedMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressAllocatedMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AddressIpVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AddressType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AttrAddressType() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AttrDnsName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AttrLoadBalancerEdition() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrLoadBalancerEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AttrLoadBalancerId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrLoadBalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AttrVpcId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) AttrZoneMappings() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrZoneMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) BandwidthPackageId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bandwidthPackageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) LoadBalancerBillingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerBillingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) LoadBalancerEdition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) LoadBalancerName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) ModificationProtectionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modificationProtectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) ResourceGroupId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) RosOptions() alicloudroscdkcore.IRosResourceOptions {
	var returns alicloudroscdkcore.IRosResourceOptions
	_jsii_.Get(
		j,
		"rosOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) RosProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"rosProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) RosResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rosResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) SecurityGroupIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) Tags() *[]*RosLoadBalancer_TagsProperty {
	var returns *[]*RosLoadBalancer_TagsProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) VpcId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosLoadBalancer) ZoneMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneMappings",
		&returns,
	)
	return returns
}


func NewRosLoadBalancer(scope alicloudroscdkcore.Construct, id *string, props *RosLoadBalancerProps, enableResourcePropertyConstraint *bool) RosLoadBalancer {
	_init_.Initialize()

	if err := validateNewRosLoadBalancerParameters(scope, id, props, enableResourcePropertyConstraint); err != nil {
		panic(err)
	}
	j := jsiiProxy_RosLoadBalancer{}

	_jsii_.Create(
		"@alicloud/ros-cdk-alb.RosLoadBalancer",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

func NewRosLoadBalancer_Override(r RosLoadBalancer, scope alicloudroscdkcore.Construct, id *string, props *RosLoadBalancerProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-alb.RosLoadBalancer",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		r,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetAccessLogConfig(val interface{}) {
	if err := j.validateSetAccessLogConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLogConfig",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetAddressAllocatedMode(val interface{}) {
	if err := j.validateSetAddressAllocatedModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressAllocatedMode",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetAddressIpVersion(val interface{}) {
	if err := j.validateSetAddressIpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressIpVersion",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetAddressType(val interface{}) {
	if err := j.validateSetAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressType",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetBandwidthPackageId(val interface{}) {
	if err := j.validateSetBandwidthPackageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthPackageId",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetLoadBalancerBillingConfig(val interface{}) {
	if err := j.validateSetLoadBalancerBillingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerBillingConfig",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetLoadBalancerEdition(val interface{}) {
	if err := j.validateSetLoadBalancerEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerEdition",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetLoadBalancerName(val interface{}) {
	if err := j.validateSetLoadBalancerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetModificationProtectionConfig(val interface{}) {
	if err := j.validateSetModificationProtectionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modificationProtectionConfig",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetResourceGroupId(val interface{}) {
	if err := j.validateSetResourceGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupId",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetSecurityGroupIds(val interface{}) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetTags(val *[]*RosLoadBalancer_TagsProperty) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetVpcId(val interface{}) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_RosLoadBalancer)SetZoneMappings(val interface{}) {
	if err := j.validateSetZoneMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneMappings",
		val,
	)
}

// Return whether the given object is a Construct.
func RosLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRosLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-alb.RosLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func RosLoadBalancer_IsRosElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRosLoadBalancer_IsRosElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-alb.RosLoadBalancer",
		"isRosElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a RosResource.
func RosLoadBalancer_IsRosResource(construct alicloudroscdkcore.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRosLoadBalancer_IsRosResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-alb.RosLoadBalancer",
		"isRosResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func RosLoadBalancer_ROS_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@alicloud/ros-cdk-alb.RosLoadBalancer",
		"ROS_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RosLoadBalancer) AddCondition(con alicloudroscdkcore.RosCondition) {
	if err := r.validateAddConditionParameters(con); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addCondition",
		[]interface{}{con},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddCount(count interface{}) {
	if err := r.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addCount",
		[]interface{}{count},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddDeletionOverride(path *string) {
	if err := r.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddDependsOn(target alicloudroscdkcore.RosResource) {
	if err := r.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddDesc(desc *string) {
	if err := r.validateAddDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addDesc",
		[]interface{}{desc},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddMetaData(key *string, value interface{}) {
	if err := r.validateAddMetaDataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMetaData",
		[]interface{}{key, value},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddPropertyDeletionOverride(propertyPath *string) {
	if err := r.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := r.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (r *jsiiProxy_RosLoadBalancer) AddRosDependency(target *string) {
	if err := r.validateAddRosDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addRosDependency",
		[]interface{}{target},
	)
}

func (r *jsiiProxy_RosLoadBalancer) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy, options *alicloudroscdkcore.RemovalPolicyOptions) {
	if err := r.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (r *jsiiProxy_RosLoadBalancer) GetAtt(attributeName *string) alicloudroscdkcore.Reference {
	if err := r.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.Reference

	_jsii_.Invoke(
		r,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosLoadBalancer) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RosLoadBalancer) OnSynthesize(session constructs.ISynthesisSession) {
	if err := r.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RosLoadBalancer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosLoadBalancer) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RosLoadBalancer) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RosLoadBalancer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := r.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosLoadBalancer) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := r.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RosLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosLoadBalancer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosLoadBalancer) ValidateProperties(_properties interface{}) {
	if err := r.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"validateProperties",
		[]interface{}{_properties},
	)
}

