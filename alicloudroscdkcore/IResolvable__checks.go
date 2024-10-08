//go:build !no_runtime_type_checking

package alicloudroscdkcore

import (
	"fmt"
)

func (i *jsiiProxy_IResolvable) validateResolveParameters(context IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

