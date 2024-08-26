package alicloudroscdkpai

import (
	_init_ "github.com/alibabacloud-go/ros-cdk/alicloudroscdkpai/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkcore"
	"github.com/alibabacloud-go/ros-cdk/alicloudroscdkpai/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// This class encapsulates and extends the ROS resource type `ALIYUN::PAI::Service`, which is used to create an Elastic Algorithm Service (EAS) service in Machine Learning Platform for AI (PAI).
type Service interface {
	alicloudroscdkcore.Resource
	// Attribute AccessToken: Service Request authentication token.
	AttrAccessToken() alicloudroscdkcore.IResolvable
	// Attribute CallerUid: The ID of the service creator, which can be the ID of the RAM account.
	AttrCallerUid() alicloudroscdkcore.IResolvable
	// Attribute Cpu: Number of service CPU cores.
	AttrCpu() alicloudroscdkcore.IResolvable
	// Attribute CreateTime: Creation time of the service.
	AttrCreateTime() alicloudroscdkcore.IResolvable
	// Attribute CurrentVersion: Current running version of the service.
	AttrCurrentVersion() alicloudroscdkcore.IResolvable
	// Attribute ExtraData: Service Extra Information.
	AttrExtraData() alicloudroscdkcore.IResolvable
	// Attribute Gpu: Number of service GPU cards.
	AttrGpu() alicloudroscdkcore.IResolvable
	// Attribute Image: Service Deployment Mirroring.
	AttrImage() alicloudroscdkcore.IResolvable
	// Attribute InternetEndpoint: Public network Endpoint of the service.
	AttrInternetEndpoint() alicloudroscdkcore.IResolvable
	// Attribute IntranetEndpoint: The intranet Endpoint of the service.
	AttrIntranetEndpoint() alicloudroscdkcore.IResolvable
	// Attribute Labels: Service Tag.
	AttrLabels() alicloudroscdkcore.IResolvable
	// Attribute LatestVersion: The latest version of the service.
	AttrLatestVersion() alicloudroscdkcore.IResolvable
	// Attribute Memory: Memory of service (MB).
	AttrMemory() alicloudroscdkcore.IResolvable
	// Attribute Message: Latest information on services.
	AttrMessage() alicloudroscdkcore.IResolvable
	// Attribute Namespace: The namespace to which the service belongs.
	AttrNamespace() alicloudroscdkcore.IResolvable
	// Attribute ParentUid: Primary account ID of the creator.
	AttrParentUid() alicloudroscdkcore.IResolvable
	// Attribute PendingInstance: Number of instances where the service is not currently ready.
	AttrPendingInstance() alicloudroscdkcore.IResolvable
	// Attribute Reason: Service deployment failure reason.
	AttrReason() alicloudroscdkcore.IResolvable
	// Attribute Resource: The ID of the resource group to which the service belongs.
	AttrResource() alicloudroscdkcore.IResolvable
	// Attribute ResourceAlias: Name of the resource group where the service resides.
	AttrResourceAlias() alicloudroscdkcore.IResolvable
	// Attribute Role: Grouping Service Role.
	AttrRole() alicloudroscdkcore.IResolvable
	// Attribute RoleAttrs: Grouping Service Role Properties.
	AttrRoleAttrs() alicloudroscdkcore.IResolvable
	// Attribute RunningInstance: Number of instances in service running.
	AttrRunningInstance() alicloudroscdkcore.IResolvable
	// Attribute SafetyLock: Service Security Lock Status.
	AttrSafetyLock() alicloudroscdkcore.IResolvable
	// Attribute ServiceConfig: Service configuration information.
	AttrServiceConfig() alicloudroscdkcore.IResolvable
	// Attribute ServiceGroup: Group to which the service belongs.
	AttrServiceGroup() alicloudroscdkcore.IResolvable
	// Attribute ServiceName: Service Name.
	AttrServiceName() alicloudroscdkcore.IResolvable
	// Attribute ServiceUid: Unique Service ID.
	AttrServiceUid() alicloudroscdkcore.IResolvable
	// Attribute TotalInstance: Total number of instances required by the service.
	AttrTotalInstance() alicloudroscdkcore.IResolvable
	// Attribute UpdateTime: Service Last Updated.
	AttrUpdateTime() alicloudroscdkcore.IResolvable
	// Attribute Weight: Packet Service Traffic Weight.
	AttrWeight() alicloudroscdkcore.IResolvable
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
	Props() *ServiceProps
	SetProps(val *ServiceProps)
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

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	internal.Type__alicloudroscdkcoreResource
}

func (j *jsiiProxy_Service) AttrAccessToken() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrCallerUid() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCallerUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrCpu() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrCreateTime() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrCurrentVersion() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrCurrentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrExtraData() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrExtraData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrGpu() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrImage() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrInternetEndpoint() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrIntranetEndpoint() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrIntranetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrLabels() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrLatestVersion() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrLatestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrMemory() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrMessage() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrNamespace() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrParentUid() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrParentUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrPendingInstance() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrPendingInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrReason() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrResource() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrResourceAlias() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrResourceAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrRole() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrRoleAttrs() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRoleAttrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrRunningInstance() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrRunningInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrSafetyLock() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrSafetyLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrServiceConfig() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrServiceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrServiceGroup() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrServiceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrServiceName() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrServiceUid() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrServiceUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrTotalInstance() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrTotalInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrUpdateTime() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrUpdateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AttrWeight() alicloudroscdkcore.IResolvable {
	var returns alicloudroscdkcore.IResolvable
	_jsii_.Get(
		j,
		"attrWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) EnableResourcePropertyConstraint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourcePropertyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() alicloudroscdkcore.ConstructNode {
	var returns alicloudroscdkcore.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Props() *ServiceProps {
	var returns *ServiceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Resource() alicloudroscdkcore.RosResource {
	var returns alicloudroscdkcore.RosResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Scope() alicloudroscdkcore.Construct {
	var returns alicloudroscdkcore.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Stack() alicloudroscdkcore.Stack {
	var returns alicloudroscdkcore.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewService(scope alicloudroscdkcore.Construct, id *string, props *ServiceProps, enableResourcePropertyConstraint *bool) Service {
	_init_.Initialize()

	if err := validateNewServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Service{}

	_jsii_.Create(
		"@alicloud/ros-cdk-pai.Service",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		&j,
	)

	return &j
}

// Param scope - scope in which this resource is defined Param id    - scoped id of the resource Param props - resource properties.
func NewService_Override(s Service, scope alicloudroscdkcore.Construct, id *string, props *ServiceProps, enableResourcePropertyConstraint *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@alicloud/ros-cdk-pai.Service",
		[]interface{}{scope, id, props, enableResourcePropertyConstraint},
		s,
	)
}

func (j *jsiiProxy_Service)SetEnableResourcePropertyConstraint(val *bool) {
	if err := j.validateSetEnableResourcePropertyConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourcePropertyConstraint",
		val,
	)
}

func (j *jsiiProxy_Service)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Service)SetProps(val *ServiceProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_Service)SetResource(val alicloudroscdkcore.RosResource) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_Service)SetScope(val alicloudroscdkcore.Construct) {
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
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alicloud/ros-cdk-pai.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) AddCondition(condition alicloudroscdkcore.RosCondition) {
	if err := s.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addCondition",
		[]interface{}{condition},
	)
}

func (s *jsiiProxy_Service) AddCount(count interface{}) {
	if err := s.validateAddCountParameters(count); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addCount",
		[]interface{}{count},
	)
}

func (s *jsiiProxy_Service) AddDependency(resource alicloudroscdkcore.Resource) {
	if err := s.validateAddDependencyParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDependency",
		[]interface{}{resource},
	)
}

func (s *jsiiProxy_Service) AddResourceDesc(desc *string) {
	if err := s.validateAddResourceDescParameters(desc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addResourceDesc",
		[]interface{}{desc},
	)
}

func (s *jsiiProxy_Service) ApplyRemovalPolicy(policy alicloudroscdkcore.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Service) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetAtt(name *string) alicloudroscdkcore.IResolvable {
	if err := s.validateGetAttParameters(name); err != nil {
		panic(err)
	}
	var returns alicloudroscdkcore.IResolvable

	_jsii_.Invoke(
		s,
		"getAtt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Service) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) SetMetadata(key *string, value interface{}) {
	if err := s.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_Service) Synthesize(session alicloudroscdkcore.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

