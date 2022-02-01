package test

import (
	//"fmt"
	"testing"
	// "time"

	//http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestS3FileUpload(t *testing.T) {

	// The values to pass into the Terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the example Terraform code is located
		TerraformDir: "..",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"region": "us-west-2",
			"label":  "terratest",
			"app":    "cx-sg1",
			"user":   "circleci-test",
		},
	})

	// Run a Terraform init and apply with the Terraform options
	terraform.InitAndApply(t, terraformOptions)

	// Run a Terraform Destroy at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Retrieve the Endpoint using Terraform Show
	endpoint := terraform.Output(t, terraformOptions, "Endpoint")
	bucket_id := terraform.Output(t, terraformOptions, "bucket_id")

	assert.True(t, len(bucket_id) > 0)
	assert.True(t, len(endpoint) > 0)

}
