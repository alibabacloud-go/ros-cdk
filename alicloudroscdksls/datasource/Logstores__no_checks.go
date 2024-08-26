//go:build no_runtime_type_checking

package datasource

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Logstores) validateAddConditionParameters(condition alicloudroscdkcore.RosCondition) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateAddCountParameters(count interface{}) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateAddDependencyParameters(resource alicloudroscdkcore.Resource) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateAddResourceDescParameters(desc *string) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateApplyRemovalPolicyParameters(policy alicloudroscdkcore.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateGetAttParameters(name *string) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateSetMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Logstores) validateSynthesizeParameters(session alicloudroscdkcore.ISynthesisSession) error {
	return nil
}

func validateLogstores_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Logstores) validateSetEnableResourcePropertyConstraintParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_Logstores) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Logstores) validateSetPropsParameters(val *LogstoresProps) error {
	return nil
}

func (j *jsiiProxy_Logstores) validateSetScopeParameters(val alicloudroscdkcore.Construct) error {
	return nil
}

func validateNewLogstoresParameters(scope alicloudroscdkcore.Construct, id *string, props *LogstoresProps) error {
	return nil
}

