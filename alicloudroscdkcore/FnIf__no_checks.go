//go:build no_runtime_type_checking

package alicloudroscdkcore

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FnIf) validateNewErrorParameters(message *string) error {
	return nil
}

func (f *jsiiProxy_FnIf) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func validateNewFnIfParameters(condition interface{}, valueIfTrue interface{}, valueIfFalse interface{}) error {
	return nil
}

