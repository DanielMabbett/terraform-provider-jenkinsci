provider "jenkinsci" {
  jenkins_endpoint       = "http://localhost:8080"
  jenkins_admin_username = "admin"
  jenkins_admin_password = "547b55dbeb9240d5b345a772d8905325"
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
