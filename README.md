# Go Terratest Helper
Repository with Utility Functions for Enhancing Efficiency in Terratest Integration and Unit Testing

## Usage
### GetResourceChangeAfterByAddressE
The `GetResourceChangeAfterByAddressE` function allows you to retrieve the "After" state of a resource change within a Terraform plan struct by specifying the resource address.

- **Input Parameters:**
  - `address (string)`: The resource address you want to retrieve the "After" state for.
plan (*terraform.PlanStruct): The Terraform plan to analyze.

- **Returns:**
  - `map[string]interface{}`: A map representing the "After" state of the resource change.
error: An error if the resource address is not found in the plan.

<details>
  <summary>Click here to see detailed Implementation:</summary>

#### Pre Requisite
We must have a local temp copy of the tfplan in json format for retrieving the exact `ResourceAddress`.

1. Add below snippet in your test code to generate a TFPlan file.
    ```go
    // Path to where we would like to create our TF Plan file
    terraformOptions.PlanFilePath = "./tfplan"
    // Using InitAndPlanAndShow  func from "github.com/gruntwork-io/terratest/modules/terraform"    module to generate the plan file
    planJson := terraform.InitAndPlanAndShow(t, terraformOptions)
    ```
2. Once the test is executed it'll generate a `tfplan` file which we need to convert into a readable json file by executing below command:
    ```ssh
    terraform show -json tfplan | jq > tfplan.json
    ```
3. From the Json file you need to pick the `resource_changes.address` value for each individual resource that you would like to test.

[Click here](mocktfplan.json) to review a sample TFPlan.json file.

> Note: Plan files shouldn't be committed to the repository or printed in the log files. Once you have got the test plan file generated ensure the above code snippet is removed.

#### Here's how to use it:
```go
import (
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/gruntwork-io/terratest/modules/tfplanstruct"
)

// Define your Terraform plan.
plan := terraform.InitAndPlan(t, terraformOptions)

// Resource Address retrieved from the tfplan.json for provisioning a S3 Bucket
resourceAddress := "module.test_website_bucket.module.bucket.aws_s3_bucket.this[0]"

// Get the "After" state of the resource change.
resourceAfterState, _ := tfplanstruct.GetResourceChangeAfterByAddressE(resourceAddress,plan)
// Use the resourceAfterState as needed.
```

####  Example:
We are using above function in testing our terraform modules as below:
- [terraform-aws-route53-public-hosted-zone unit test](https://github.com/armakuni/terraform-aws-route53-public-hosted-zone/blob/3f025a61b44823f3165781e516bf67a61bf7df05/test/unit/route53_zone_records_test.go#L63)
- [terraform-aws-static-website-bucket testing multiple resource creations](https://github.com/armakuni/terraform-aws-static-website-bucket/blob/main/test/unit/website_bucket_test.go)

</details>

### InterfaceSliceToStringSliceE

The `InterfaceSliceToStringSliceE` function is a Go function that converts a slice of interface values into a slice of strings. Here's a quick brief:

- **Parameters:**
  - `input []interface{}`: A slice of interface values that you want to convert to strings.

- **Return Values:**
  - `[]string`: A slice of strings containing the string representations of the input values.
  - `error`: An error, which is non-nil if the input is `nil`.

<details>
  <summary>Click here to see detailed Implementation:</summary>
    
  - **Usage Example:**
    You can use this function to convert a slice of arbitrary interface values into a slice of  strings, which can be helpful when you need to display or manipulate the string  representations of these values.

    Here's an example of how you might use this function:

    ```go
    inputSlice := []interface{}{42, "Hello, World!", 3.14}
    stringSlice, err := InterfaceSliceToStringSliceE(inputSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(stringSlice) // Output: ["42" "Hello, World!" "3.14"]
    }
    ```

 - **Error Handling:**
  
   - The function checks for a `nil` input and returns an error if the input is `nil`.
   - Other error cases, such as invalid input types, are not explicitly handled in thi    function.
</details>

### InitAndPlanAndShowWithStructNoLogTempPlanFileE

This Go function facilitates Terratest testing of Terraform configurations.

The Terratest lib do not have function which returns error when TF Plan fails. Example when we would like to test the code for validation or an expected error scenario.

This function allows you to initialize Terraform, create a plan, and show the plan with options for suppressing log output. It also generates a temporary plan file to store the plan. Here's a quick brief:

- **Parameters:**
  - `t testing.TestingT`: The test context for Terratest.
  - `options *terraform.Options`: Terraform configuration options.

- **Return Values:**
  - `*terraform.PlanStruct`: The generated Terraform plan.
  - `error`: Any encountered errors during execution. Example: A TF Variable validation failures.


<details>
  <summary>Click here to see detailed Implementation:</summary>
    
  - **Usage Example:**
    Conveniently run Terraform operations with log suppression and plan storage in Terratest tests.

    ```go
    plan, err := InitAndPlanAndShowWithStructNoLogTempPlanFileE(t, options)
    ```

 - **Notes:**
  
   - Designed for Terratest.
   - Ensure correct test setup.
   - Does not handle Terraform configuration errors.
</details>
