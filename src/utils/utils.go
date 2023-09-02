package utils

import (
	"fmt"
)

// 拼接错误
func AppendError(prev, next error) error {
	if prev == nil {
		return next
	}
	return fmt.Errorf("%v,%w", prev, next)
}
