package test

import (
	"testing"
	"time"

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

	publicip := terraform.Output(t, terraformOptions, "publicip")
	vm_pass := terraform.Output(t, terraformOptions, "vm_pass")
	vm_port := terraform.Output(t, terraformOptions, "vm_port")

	time.Sleep(15 * time.Second)
	assert.True(t, isConnect(publicip, vm_port, "root", vm_pass))
}
