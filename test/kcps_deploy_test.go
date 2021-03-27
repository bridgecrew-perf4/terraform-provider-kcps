package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestTerraformKcpsDeploy(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform/",
	}
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	publicip := terraform.OutputMap(t, terraformOptions, "publicip")["value"]
	vm_pass := terraform.OutputMap(t, terraformOptions, "vm_pass")["value"]
	vm_port := terraform.OutputMap(t, terraformOptions, "vm_port")["value"]

	assert.True(t, isConnect(publicip, vm_port, "root", vm_pass))
}
