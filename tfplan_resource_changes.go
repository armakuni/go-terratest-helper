package tfplanstruct

import (
	"errors"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// Retrieves the "After" state of a resource change within a Terraform plan by matching it to a provided resource address.
func GetResourceChangeAfterByAddressE(address string, plan *terraform.PlanStruct) (map[string]interface{}, error) {
	if address == "" || len(address) == 0 {
		return nil, errors.New("Address cannot be empty")
	} else if plan == nil {
		return nil, errors.New("PlanStruct cannot be empty or nil")
	}
	for _, value := range plan.ResourceChangesMap {
		if value.Address == address {
			return value.Change.After.(map[string]interface{}), nil
		}
	}
	return nil, errors.New("No matching Address found for " + address)
}
