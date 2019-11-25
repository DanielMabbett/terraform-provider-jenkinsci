# Views

resource "jenkinsci_view" "view" {
  name             = "1st-view"
}

resource "jenkinsci_view" "test2" {
  name             = "2nd-view"
  assigned_projects = [
    jenkinsci_project.test.name,
    jenkinsci_project.test2.name
  ]
}


# Folders

resource "jenkinsci_folder" "test" {
  name = "folder"
}

resource "jenkinsci_folder" "test-nested-folder" {
  name          = "nestedfolder"
  parent_folder = jenkinsci_folder.test.name
}

# Projects

## Simple Empty Project with nothing in it
resource "jenkinsci_project" "test" {
  name = "test-project-1a"
}

## A project with disabled features in and description added 
resource "jenkinsci_project" "test2" {
  name          = "test-project-2a"
  description   = "my test project - version 2"
  disabled      = true
  assigned_node = "terraform-pod"
}

## A project with additional config added 
resource "jenkinsci_project" "test3" {
  name          = "test-project-3a"
  description   = "my test project - version 3a"
  assigned_node = "terraform-pod"

  parameter {
    type  = "string"
    value = "tp-value"
    key   = "tp-key"
  }

  ### If you want to add additional configuration from things such as installed plugins then you can add them as xml
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

## Simple Project in a folder
resource "jenkinsci_project" "test-in-folder" {
  name = "testprojinfolder"
  folder = jenkinsci_folder.test.name
}

# Pipelines

resource "jenkinsci_pipeline" "test" {
  name = "pipelinejob"
  disabled = true
  pipeline_script = "${file("${path.root}/jenkinsfile")}"
}

# Credentials 

resource "jenkinsci_credential_secret_text" "test" {
  name        = "test"
  domain      = "_"
  scope       = "GLOBAL"
  secret      = "thevalue"
  alias_id    = "test"
  description = "some description now"
}

# Plugins

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