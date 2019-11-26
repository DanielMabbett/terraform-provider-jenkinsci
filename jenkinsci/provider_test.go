package jenkinsci

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"jenkins": testAccProvider,
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
	if v := os.Getenv("JENKINS_ENDPOINT"); v == "" {
		t.Fatal("JENKINS_ENDPOINT must be set for acceptance tests")
	}
	if v := os.Getenv("JENKINS_ADMIN_USERNAME"); v == "" {
		t.Fatal("JENKINS_ADMIN_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("JENKINS_ADMIN_PASSWORD"); v == "" {
		t.Fatal("JENKINS_ADMIN_PASSWORD must be set for acceptance tests")
	}
}
