package tfplan

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestInitAndPlanAndShowWithStructNoLogTempPlanFileEReturnsPlan(t *testing.T) {
	inputVariables := map[string]interface{}{
		"prefix": "test",
	}
	options := toTerraformOptions("../testresources/tfcode", inputVariables)

	plan, err := InitAndPlanAndShowWithStructNoLogTempPlanFileE(t, &options)

	assert.Empty(t, err)
	assert.NotEmpty(t, plan)
}

func TestInitAndPlanAndShowWithStructNoLogTempPlanFileEReturnsError(t *testing.T) {
	inputVariables := map[string]interface{}{
		"prefix": "123test",
	}
	options := toTerraformOptions("../testresources/tfcode", inputVariables)

	plan, err := InitAndPlanAndShowWithStructNoLogTempPlanFileE(t, &options)

	assert.ErrorContains(t, err, "Prefix can contain only alphabetic characters.")
	assert.Empty(t, plan)
}

func toTerraformOptions(path string, vars map[string]interface{}) terraform.Options {
	return terraform.Options{
		TerraformDir: path,
		Vars:         vars,
	}
}
