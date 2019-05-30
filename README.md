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
# Note that this currently only supports http connections
provider "jenkinsci" {
  jenkins_endpoint         = "..."
  jenkins_admin_username   = "..."
  jenkins_admin_password   = "..."
}

# Create a basic Jenkins Project
resource "jenkinsci_project" "test" {
  name = "mytestproj2"
}

# Create a Folder
resource "jenkinsci_folder" "name" {
  name = "test"
}

# Create a jenkins view tab
resource "jenkinsci_view" "name" {
  name = "view"
}

# Install a plugin
resource "jenkinsci_plugin" "ccm" {
  name    = "CCM"
  version = "3.2"
}

```