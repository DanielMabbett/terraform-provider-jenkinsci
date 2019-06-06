terraform-provider-jenkinsci
==================

[![pipeline status](https://gitlab.com/daniel.mabbett/terraform-provider-jenkinsci/badges/master/pipeline.svg)](https://gitlab.com/daniel.mabbett/terraform-provider-jenkinsci/commits/master)

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



# Simple Empty Project with nothing in it
resource "jenkinsci_project" "test" {
  name = "testproj"
}


# A view with an assigned project in the view. Only works with 1 project assigned so far
resource "jenkinsci_view" "test" {
  name = "view2"
  assigned_project = "${jenkinsci_project.test2.name}"
}

# A test project that is inside a folder
resource "jenkinsci_project" "test2" {
  name          = "test-project-2a"
  description   = "my test project - version 2"
  disabled      = "true"
  assigned_node = "terraform-pod"
}

# A test project that is inside a folder
resource "jenkinsci_project" "test3" {
  name          = "test-project-3a"
  description   = "my test project - version 3a"
  assigned_node = "terraform-pod"

  parameter {
    value = "tp-value"
    type = "tp-string"
    key = "tp-key"
  }

  # If you want to add additional configuration from things such as installed plugins then you can add them as xml
  additional_config = <<XML
    <authToken>asdadadadadasd</authToken>
  XML
}

# Simple folder
resource "jenkinsci_folder" "test" {
  name = "folder"
}

# Nested Folder
resource "jenkinsci_folder" "nested-folder" {
  name = "nestedfolder"
  parent_folder = "${jenkinsci_folder.test.name}"
}

# Simple Project in a folder
resource "jenkinsci_project" "test-in-folder" {
  name   = "testprojinfolder"
  folder = "${jenkinsci_folder.test.name}"
}

# Plugins Examples
resource "jenkinsci_plugin" "terraform" {
  name    = "Terraform"
  version = "1.0.9"
}

resource "jenkinsci_plugin" "ccm" {
  name    = "CCM"
  version = "3.2"
}

resource "jenkinsci_plugin" "ansicolor" {
  name    = "AnsiColor"
  version = "0.6.2"
}

```

Known Issues
---------------------
Due to some of the limitations of gojenkins, we presently: 
* Cannot delete views
* Cannot delete folders