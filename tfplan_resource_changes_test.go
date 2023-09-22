package resourcechanges_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/armakuni/go-terratest-helper"
)

func TestGetResourceChangeAfterByAddressReturnsErrorWhenAddressStringIsEmpty(t *testing.T) {
	tfStruct := getMockTFPlanStruct("./mocktfplan.json")

	_, err := GetResourceChangeAfterByAddressE("", tfStruct)

	assert.ErrorContains(t, err, `Address cannot be empty`)
}

func TestGetResourceChangeAfterByAddressReturnsMatchingAddress(t *testing.T) {
	tfStruct := getMockTFPlanStruct("./mocktfplan.json")

	module, _ := GetResourceChangeAfterByAddressE("module.test_website_bucket.module.bucket.aws_s3_bucket_public_access_block.this[0]", tfStruct)

	assert.NotEmpty(t, module)
}

func TestGetResourceChangeAfterByAddressReturnsErrorWhenPlanIsNil(t *testing.T) {
	_, err := GetResourceChangeAfterByAddressE("module.test_website_bucket.module.bucket.aws_s3_bucket_public_access_block.this[0]", nil)

	assert.ErrorContains(t, err, `PlanStruct cannot be empty or nil`)
}

// func TestGetResourceChangeAfterByAddressReturnsErrorWhenPlanIsEmpty(t *testing.T) {
// 	_, err := GetResourceChangeAfterByAddressE("module.test_website_bucket.module.bucket.aws_s3_bucket_public_access_block.this[0]", &terraform.PlanStruct{})

// 	assert.ErrorContains(t, err, `PlanStruct cannot be empty or nil`)
// }

func TestGetResourceChangeAfterByAddressReturnsErrorWhenNoMatchingAddressFound(t *testing.T) {
	tfStruct := getMockTFPlanStruct("./mocktfplan.json")
	address := "module.test_website_bucket.module.bucket.aws_s3_bucket_public_access_block"

	_, err := GetResourceChangeAfterByAddressE(address, tfStruct)

	assert.ErrorContains(t, err, "No matching Address found for "+address)
}

func getMockTFPlanStruct(filePath string) *terraform.PlanStruct {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	planStruct, err := terraform.ParsePlanJSON(string(data))
	if err != nil {
		fmt.Println("Error Parsing Mock Plan JSON file:", err)
	}
	return planStruct
}
