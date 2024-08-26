package alicloudroscdkconfig

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkconfig/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkconfig/internal"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class is a base encapsulation around the ROS resource type `ALIYUN::Config::Rule`, which is used to create or modify a rule.
type RosRule interface {
	alicloudroscdkcore.RosResource
	AttrConfigRuleArn() alicloudroscdkcore.IResolvable
	AttrConfigRuleId() alicloudroscdkcore.IResolvable
	AttrConfigRuleTriggerTypes() alicloudroscdkcore.IResolvable
	AttrDescription() alicloudroscdkcore.IResolvable
	AttrEventSource() alicloudroscdkcore.IResolvable
	AttrExcludeResourceIdsScope() alicloudroscdkcore.IResolvable
	AttrInputParameters() alicloudroscdkcore.IResolvable
	AttrMaximumExecutionFrequency() alicloudroscdkcore.IResolvable
	AttrRegionIdsScope() alicloudroscdkcore.IResolvable
	AttrResourceGroupIdsScope() alicloudroscdkcore.IResolvable
	AttrResourceTypesScope() alicloudroscdkcore.IResolvable
	AttrRiskLevel() alicloudroscdkcore.IResolvable
	AttrRuleName() alicloudroscdkcore.IResolvable
	AttrSourceIdentifier() alicloudroscdkcore.IResolvable
	AttrSourceOwner() alicloudroscdkcore.IResolvable
	AttrTagKeyScope() alicloudroscdkcore.IResolvable
	AttrTagValueScope() alicloudroscdkcore.IResolvable
	ConfigRuleTriggerTypes() interface{}
	SetConfigRuleTriggerTypes(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aliyun:ros:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Description() interface{}
	SetDescription(val interface{})
	EnableResourcePropertyConstraint() *bool
	SetEnableResourcePropertyConstraint(val *bool)
	ExcludeResourceIdsScope() interface{}
	SetExcludeResourceIdsScope(val interface{})
	InputParameters() interface{}
	SetInputParameters(val interface{})
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
	MaximumExecutionFrequency() interface{}
	SetMaximumExecutionFrequency(val interface{})
	// The construct tree node associated with this construct.
	Node() alicloudroscdkcore.ConstructNode
	// Return a string that will be resolved to a RosTemplate `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	RegionIdsScope() interface{}
	SetRegionIdsScope(val interface{})
	ResourceGroupIdsScope() interface{}
	SetResourceGroupIdsScope(val interface{})
	ResourceTypesScope() interface{}
	SetResourceTypesScope(val interface{})
	RiskLevel() interface{}
	SetRiskLevel(val interface{})
	// Options for this resource, such as condition, update policy etc.
	RosOptions() alicloudroscdkcore.IRosResourceOptions
	RosProperties() *map[string]interface{}
	// ROS resource type.
	RosResourceType() *string
	RuleName() interface{}
	SetRuleName(val interface{})
	SourceIdentifier() interface{}
	SetSourceIdentifier(val interface{})
	SourceOwner() interface{}
	SetSourceOwner(val interface{})
	// The stack in which this element is defined.
	//
	// RosElements must be defined within a stack scope (directly or indirectly).
	Stack() alicloudroscdkcore.Stack
	TagKeyLogicScope() interface{}
	SetTagKeyLogicScope(val interface{})
	TagKeyScope() interface{}
	SetTagKeyScope(val interface{})
	TagValueScope() interface{}
	SetTagValueScope(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for RosRule
type jsiiProxy_RosRule struct {
	internal.Type__alicloudroscdkcoreRosResource
}

func (j *jsiiProxy_RosRule) AttrConfigRuleArn() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrConfigRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrConfigRuleId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrConfigRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrConfigRuleTriggerTypes() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrConfigRuleTriggerTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrDescription() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrEventSource() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrExcludeResourceIdsScope() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrExcludeResourceIdsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrInputParameters() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrInputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrMaximumExecutionFrequency() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrMaximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrRegionIdsScope() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRegionIdsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrResourceGroupIdsScope() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResourceGroupIdsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrResourceTypesScope() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrRiskLevel() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRiskLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrRuleName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrSourceIdentifier() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrSourceOwner() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrTagKeyScope() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) AttrTagValueScope() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) ConfigRuleTriggerTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configRuleTriggerTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) Description() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) ExcludeResourceIdsScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceIdsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) InputParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) MaximumExecutionFrequency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) RegionIdsScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionIdsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) ResourceGroupIdsScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceGroupIdsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) ResourceTypesScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) RiskLevel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"riskLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) RosOptions() alicloudroscdkcore.IRosResourceOptions {
	var returns alicloudroscdkcore.IRosResourceOptions
	_jsii_.Get(
		j,
		"rosOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) RosProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"rosProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) RosResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rosResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) RuleName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) SourceIdentifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) SourceOwner() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) TagKeyLogicScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagKeyLogicScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) TagKeyScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) TagValueScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RosRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


func NewRosRule(scope alicloudroscdkcore.Construct, id *string, props *RosRuleProps, enableResourcePropertyConstraint *bool) RosRule {
	_init_.Initialize()

	if err := validateNewRosRuleParameters(scope, id, props, enableResourcePropertyConstraint); err != nil {
		panic(err)
	}
	j := jsiiProxy_RosRule{}

	_jsii_.Create(
		"@alicloud/ros-cdk-config.RosRule",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

func NewRosRule_Override(r RosRule, scope alicloudroscdkcore.Construct, id *string, props *RosRuleProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-config.RosRule",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		r,
	)
}

func (j *jsiiProxy_RosRule)SetConfigRuleTriggerTypes(val interface{}) {
	if err := j.validateSetConfigRuleTriggerTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configRuleTriggerTypes",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetDescription(val interface{}) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetExcludeResourceIdsScope(val interface{}) {
	if err := j.validateSetExcludeResourceIdsScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeResourceIdsScope",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetInputParameters(val interface{}) {
	if err := j.validateSetInputParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetMaximumExecutionFrequency(val interface{}) {
	if err := j.validateSetMaximumExecutionFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetRegionIdsScope(val interface{}) {
	if err := j.validateSetRegionIdsScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionIdsScope",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetResourceGroupIdsScope(val interface{}) {
	if err := j.validateSetResourceGroupIdsScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupIdsScope",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetResourceTypesScope(val interface{}) {
	if err := j.validateSetResourceTypesScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetRiskLevel(val interface{}) {
	if err := j.validateSetRiskLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"riskLevel",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetRuleName(val interface{}) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetSourceIdentifier(val interface{}) {
	if err := j.validateSetSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetSourceOwner(val interface{}) {
	if err := j.validateSetSourceOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceOwner",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetTagKeyLogicScope(val interface{}) {
	if err := j.validateSetTagKeyLogicScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyLogicScope",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetTagKeyScope(val interface{}) {
	if err := j.validateSetTagKeyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_RosRule)SetTagValueScope(val interface{}) {
	if err := j.validateSetTagValueScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

// Return whether the given object is a Construct.
func RosRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRosRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-config.RosRule",
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
func RosRule_IsRosElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRosRule_IsRosElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-config.RosRule",
		"isRosElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a RosResource.
func RosRule_IsRosResource(construct alicloudroscdkcore.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRosRule_IsRosResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-config.RosRule",
		"isRosResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func RosRule_ROS_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@alicloud/ros-cdk-config.RosRule",
		"ROS_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RosRule) AddCondition(con alicloudroscdkcore.RosCondition) {
	if err := r.validateAddConditionParameters(con); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addCondition",
		[]interface{}{con},
	)
}

func (r *jsiiProxy_RosRule) AddCount(count interface{}) {
	if err := r.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addCount",
		[]interface{}{count},
	)
}

func (r *jsiiProxy_RosRule) AddDeletionOverride(path *string) {
	if err := r.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (r *jsiiProxy_RosRule) AddDependsOn(target alicloudroscdkcore.RosResource) {
	if err := r.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (r *jsiiProxy_RosRule) AddDesc(desc *string) {
	if err := r.validateAddDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addDesc",
		[]interface{}{desc},
	)
}

func (r *jsiiProxy_RosRule) AddMetaData(key *string, value interface{}) {
	if err := r.validateAddMetaDataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMetaData",
		[]interface{}{key, value},
	)
}

func (r *jsiiProxy_RosRule) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RosRule) AddPropertyDeletionOverride(propertyPath *string) {
	if err := r.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (r *jsiiProxy_RosRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := r.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (r *jsiiProxy_RosRule) AddRosDependency(target *string) {
	if err := r.validateAddRosDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addRosDependency",
		[]interface{}{target},
	)
}

func (r *jsiiProxy_RosRule) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy, options *alicloudroscdkcore.RemovalPolicyOptions) {
	if err := r.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (r *jsiiProxy_RosRule) GetAtt(attributeName *string) alicloudroscdkcore.Reference {
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

func (r *jsiiProxy_RosRule) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RosRule) OnSynthesize(session constructs.ISynthesisSession) {
	if err := r.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RosRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosRule) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RosRule) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RosRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (r *jsiiProxy_RosRule) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := r.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RosRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RosRule) ValidateProperties(_properties interface{}) {
	if err := r.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"validateProperties",
		[]interface{}{_properties},
	)
}

