package utils

import (
	"fmt"
	"strings"
)

func JoinErrors(errors []error) error {
	errorStrings := make([]string, 0, len(errors))
	for _, err := range errors {
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
		}
	}
	if len(errorStrings) > 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

