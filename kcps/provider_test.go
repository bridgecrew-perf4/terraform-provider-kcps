package kcps

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"kcps": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	requiredEnvs := []string{
		"KCPS_API_URL",
		"KCPS_API_KEY",
		"KCPS_SECRET_KEY",
	}

	for _, env := range requiredEnvs {
		if v := os.Getenv(env); v == "" {
			t.Fatal(fmt.Sprintf("%s must be set for acceptance tests", env))
		}
	}
}
