package utils

import "fmt"

func AppendError(preError, currentError error) error {
	if preError == nil {
		return currentError
	}
	return fmt.Errorf("%v,%w", preError, currentError)
}
