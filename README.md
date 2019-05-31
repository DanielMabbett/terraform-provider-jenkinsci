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

resource "jenkinsci_project" "test" {
  name = "testproj"
}

resource "jenkinsci_project" "test2" {
  name          = "testproj2"
  description   = "my test project - version 2"
  disabled      = "true"
  assigned_node = "terraform-pod"
}

resource "jenkinsci_project" "test-in-folder" {
  name   = "testprojinfolder"
  folder = "${jenkinsci_folder.test.name}"
}

resource "jenkinsci_folder" "test" {
  name = "folder"
}

resource "jenkinsci_view" "test" {
  name = "view"
}

resource "jenkinsci_plugin" "terraform" {
  name    = "Terraform"
  version = "1.0.9"
}

resource "jenkinsci_plugin" "ccm" {
  name    = "CCM"
  version = "3.2"
}

```

Known Issues
---------------------
Due to some of the limitations of gojenkins, we presently: 
* Cannot delete views
* Cannot delete folders