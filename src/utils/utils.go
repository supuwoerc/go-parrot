package utils

import "fmt"

func AppendError(prev, next error) error {
	if prev == nil {
		return next
	}
	return fmt.Errorf("%v,%w", prev, next)
}
