provider "jenkinsci" {
  jenkins_endpoint       = "http://localhost:8080"
  jenkins_admin_username = "admin"
  jenkins_admin_password = "547b55dbeb9240d5b345a772d8905325"
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
  name          = "testproj2"
  description   = "my test project - version 2"
  disabled      = "true"
  assigned_node = "terraform-pod"
}

# Simple folder
resource "jenkinsci_folder" "test" {
  name = "folder"
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
