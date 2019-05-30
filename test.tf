provider "jenkinsci" {
  jenkins_endpoint         = "http://localhost:8080"
  jenkins_admin_username   = "admin"
  jenkins_admin_password   = "547b55dbeb9240d5b345a772d8905325"
}

resource "jenkinsci_project" "test" {
  name = "mytestproj2"
}

resource "jenkinsci_folder" "name" {
  name = "test"
}

resource "jenkinsci_view" "name" {
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
