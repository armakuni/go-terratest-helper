package utils

import (
	"errors"
	"fmt"
)

func InterfaceSliceToStringSliceE(input []interface{}) ([]string, error) {
	if input == nil {
		return nil, errors.New("Invalid input interface, cannot be nil")
	}
	result := make([]string, len(input))
	for i, val := range input {
		result[i] = fmt.Sprintf("%v", val)
	}
	return result, nil
}
