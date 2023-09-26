# Go Terratest Helper
Repository with Utility Functions for Enhancing Efficiency in Terratest Integration and Unit Testing

## Usage
### GetResourceChangeAfterByAddressE
The GetResourceChangeAfterByAddressE function allows you to retrieve the "After" state of a resource change within a Terraform plan struct by specifying the resource address.
#### Pre Requisite
We must have a local temp copy of the tfplan in json format for retrieving the exact `ResourceAddress`.

1. Add below snippet in your test code to generate a TFPlan file.
```go
// Path to where we would like to create our TF Plan file
terraformOptions.PlanFilePath = "./tfplan"
// Using InitAndPlanAndShow  func from "github.com/gruntwork-io/terratest/modules/terraform" module to generate the plan file
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
resourceAfterState, _ := tfplanstruct.GetResourceChangeAfterByAddressE(resourceAddress, plan)
// Use the resourceAfterState as needed.
```

#### Parameters:
address (string): The resource address you want to retrieve the "After" state for.
plan (*terraform.PlanStruct): The Terraform plan to analyze.
#### Returns:
map[string]interface{}: A map representing the "After" state of the resource change.
error: An error if the resource address is not found in the plan.

#### Real Time Example:
We are using above function in testing our terraform modules as below:
- [terraform-aws-route53-public-hosted-zone unit test](https://github.com/armakuni/terraform-aws-route53-public-hosted-zone/blob/3f025a61b44823f3165781e516bf67a61bf7df05/test/unit/route53_zone_records_test.go#L63)
- [terraform-aws-static-website-bucket testing multiple resource creations](https://github.com/armakuni/terraform-aws-static-website-bucket/blob/main/test/unit/website_bucket_test.go)

