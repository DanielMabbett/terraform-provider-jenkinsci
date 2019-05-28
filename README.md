terraform-provider-jenkinsci
==================

Building The Provider
---------------------
Clone the repository.

```bash
mkdir -p $GOPATH/src/github.com/terraform; 
cd $GOPATH/src/github.com/terraform
git clone https://github.com/DanielMabbett/terraform-provider-jenkinsci
```

Enter the provider directory and build the provider
```bash
sh build.sh
```

Using The Provider
---------------------
```hcl
provider "jenkinsci" {
  jenkins_endpoint         = "..."
  jenkins_admin_username   = "..."
  jenkins_admin_password   = "..."
}

# Create a Jenkins Project
resource "jenkins_project" "example" {
  name     = "test"
}
```