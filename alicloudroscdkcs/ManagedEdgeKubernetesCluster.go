package alicloudroscdkcs

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkcs/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcs/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `ALIYUN::CS::ManagedEdgeKubernetesCluster`, which is used to create a Container Service for Kubernetes (ACK) edge cluster.
type ManagedEdgeKubernetesCluster interface {
	alicloudroscdkcore.Resource
	// Attribute APIServerSLBId: The id of API server SLB.
	AttrApiServerSlbId() alicloudroscdkcore.IResolvable
	// Attribute ClusterId: Cluster instance ID.
	AttrClusterId() alicloudroscdkcore.IResolvable
	// Attribute DefaultUserKubeConfig: Default user kubernetes config which is used for configuring cluster credentials.
	AttrDefaultUserKubeConfig() alicloudroscdkcore.IResolvable
	// Attribute IngressSLBId: The id of ingress SLB.
	AttrIngressSlbId() alicloudroscdkcore.IResolvable
	// Attribute Nodes: The list of cluster nodes.
	AttrNodes() alicloudroscdkcore.IResolvable
	// Attribute PrivateUserKubConfig: Private user kubernetes config which is used for configuring cluster credentials.
	AttrPrivateUserKubConfig() alicloudroscdkcore.IResolvable
	// Attribute ScalingConfigurationId: Scaling configuration id.
	AttrScalingConfigurationId() alicloudroscdkcore.IResolvable
	// Attribute ScalingGroupId: Scaling group id.
	AttrScalingGroupId() alicloudroscdkcore.IResolvable
	// Attribute ScalingRuleId: Scaling rule id.
	AttrScalingRuleId() alicloudroscdkcore.IResolvable
	// Attribute TaskId: Task ID.
	//
	// Automatically assigned by the system, the user queries the task status.
	AttrTaskId() alicloudroscdkcore.IResolvable
	// Attribute WorkerRamRoleName: Worker ram role name.
	AttrWorkerRamRoleName() alicloudroscdkcore.IResolvable
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
	Props() *ManagedEdgeKubernetesClusterProps
	SetProps(val *ManagedEdgeKubernetesClusterProps)
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

// The jsii proxy struct for ManagedEdgeKubernetesCluster
type jsiiProxy_ManagedEdgeKubernetesCluster struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrApiServerSlbId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrApiServerSlbId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrClusterId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrDefaultUserKubeConfig() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrDefaultUserKubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrIngressSlbId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrIngressSlbId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrNodes() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrPrivateUserKubConfig() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPrivateUserKubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrScalingConfigurationId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrScalingConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrScalingGroupId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrScalingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrScalingRuleId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrScalingRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrTaskId() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTaskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) AttrWorkerRamRoleName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrWorkerRamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Props() *ManagedEdgeKubernetesClusterProps {
	var returns *ManagedEdgeKubernetesClusterProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewManagedEdgeKubernetesCluster(scope alicloudroscdkcore.Construct, id *string, props *ManagedEdgeKubernetesClusterProps, enableResourcePropertyConstraint *bool) ManagedEdgeKubernetesCluster {
	_init_.Initialize()

	if err := validateNewManagedEdgeKubernetesClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedEdgeKubernetesCluster{}

	_jsii_.Create(
		"@alicloud/ros-cdk-cs.ManagedEdgeKubernetesCluster",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewManagedEdgeKubernetesCluster_Override(m ManagedEdgeKubernetesCluster, scope alicloudroscdkcore.Construct, id *string, props *ManagedEdgeKubernetesClusterProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-cs.ManagedEdgeKubernetesCluster",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		m,
	)
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster)SetProps(val *ManagedEdgeKubernetesClusterProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_ManagedEdgeKubernetesCluster)SetScope(val alicloudroscdkcore.Construct) {
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
func ManagedEdgeKubernetesCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedEdgeKubernetesCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-cs.ManagedEdgeKubernetesCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := m.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addCondition",
		[]interface{}{condition},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) AddCount(count interface{}) {
	if err := m.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addCount",
		[]interface{}{count},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := m.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addDependency",
		[]interface{}{resource},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) AddResourceDesc(desc *string) {
	if err := m.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := m.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		m,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := m.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) Prepare() {
	_jsii_.InvokeVoid(
		m,
		"prepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) SetMetadata(key *string, value interface{}) {
	if err := m.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := m.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEdgeKubernetesCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

