package utils_test

import (
	"testing"

	utils "github.com/armakuni/go-terratest-helper/utils"
	"github.com/stretchr/testify/assert"
)

func TestIfInputIsNilShouldReturnError(t *testing.T) {
	_, error := utils.InterfaceSliceToStringSliceE(nil)
	assert.ErrorContains(t, error, "Invalid input interface, cannot be nil")
}

func TestWhenSliceInterfaceIsEmptyReturnsEmptyStringSlice(t *testing.T) {
	input := []interface{}{}

	stringSlice, _ := utils.InterfaceSliceToStringSliceE(input)

	expectedSlice := []string{}
	assert.Equal(t, expectedSlice, stringSlice)
}

func TestWhenSliceInterfaceIsPassedReturnsStringSlice(t *testing.T) {
	input := []interface{}{"mello", "Hello, World!", 3.14}

	stringSlice, _ := utils.InterfaceSliceToStringSliceE(input)

	expectedSlice := []string{"mello", "Hello, World!", "3.14"}
	assert.Equal(t, expectedSlice, stringSlice)
}
