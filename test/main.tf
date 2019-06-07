# ----------------------------{Views}---------------------------- #
# A simple view
resource "jenkinsci_view" "test" {
  name             = "1st-view"
}

# A view with an assigned project in the view. Only works with 1 project assigned so far
resource "jenkinsci_view" "test2" {
  name             = "2nd-view"
  assigned_project = "${jenkinsci_project.test2.name}"
}


# ----------------------------{Folders}---------------------------- #

# Simple folder
resource "jenkinsci_folder" "test" {
  name = "folder"
}

# Nested Folder
resource "jenkinsci_folder" "nested-folder" {
  name          = "nestedfolder"
  parent_folder = "${jenkinsci_folder.test.name}"
}


# ----------------------------{Projects}---------------------------- #

# Simple Empty Project with nothing in it
resource "jenkinsci_project" "test" {
  name = "test-project-1a"
}

# A project with disabled features in and description added 
resource "jenkinsci_project" "test2" {
  name          = "test-project-2a"
  description   = "my test project - version 2"
  disabled      = "true"
  assigned_node = "terraform-pod"
}

# A project with additional config added 
resource "jenkinsci_project" "test3" {
  name          = "test-project-3a"
  description   = "my test project - version 3a"
  assigned_node = "terraform-pod"

  parameter {
    type  = "string"
    value = "tp-value"
    key   = "tp-key"
  }

  # If you want to add additional configuration from things such as installed plugins then you can add them as xml
  additional_config = <<XML
    <builders>
      <hudson.tasks.Shell>
        <command>
          hostname; echo "hello world";
        </command>
      </hudson.tasks.Shell>
    </builders>
    <authToken>asdadadadadasd</authToken>
  XML
}

data "template_file" "cloud-config" {
  template = "${file("${path.module}/project-template.xml")}"
  vars     = {
    authToken = "anauthtoken"
  }
}

# Simple Project in a folder
resource "jenkinsci_project" "test-in-folder" {
  name = "testprojinfolder"
  folder = "${jenkinsci_folder.test.name}"
}

# ----------------------------{Plugins}---------------------------- #
# Plugins Examples
resource "jenkinsci_plugin" "terraform" {
  name = "Terraform"
  version = "1.0.9"
}

resource "jenkinsci_plugin" "ccm" {
  name = "CCM"
  version = "3.2"
}

resource "jenkinsci_plugin" "ansicolor" {
  name = "AnsiColor"
  version = "0.6.2"
}