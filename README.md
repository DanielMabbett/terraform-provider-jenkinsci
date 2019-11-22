terraform-provider-jenkinsci
==================

[![Go Report Card](https://goreportcard.com/badge/github.com/DanielMabbett/terraform-provider-jenkinsci)](https://goreportcard.com/report/github.com/DanielMabbett/terraform-provider-jenkinsci)
[![Build Status](https://travis-ci.org/DanielMabbett/terraform-provider-jenkinsci.svg?branch=master)](https://travis-ci.org/DanielMabbett/terraform-provider-jenkinsci)

(Older CI)
[![CircleCI](https://circleci.com/gh/DanielMabbett/terraform-provider-jenkinsci.svg?style=svg)](https://circleci.com/gh/DanielMabbett/terraform-provider-jenkinsci)

Building The Provider
---------------------

Clone the repository.

```bash
mkdir -p $GOPATH/src/github.com/terraform;
cd $GOPATH/src/github.com/terraform
git clone https://github.com/DanielMabbett/terraform-provider-jenkinsci
```

Enter the provider directory and build the provider. Run `make` or:

```bash
make build
```

Using The Provider
---------------------

```hcl
provider "jenkinsci" {
  jenkins_endpoint         = "..."
  jenkins_admin_username   = "..."
  jenkins_admin_password   = "..."
}


# Views

resource "jenkinsci_view" "view" {
  name             = "1st-view"
}

# A view with an assigned project in the view. Only works with 1 project assigned so far
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


# Nested Folder
resource "jenkinsci_folder" "test-nested-folder" {
  name          = "nestedfolder"
  parent_folder = jenkinsci_folder.test.name
}

# Projects

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
  folder = jenkinsci_folder.test.name
}

# Pipelines

resource "jenkinsci_pipeline" "test" {
  name = "pipelinejob"
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

```

Developing the Provider
----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is **required**). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

First clone the repository to: `$GOPATH/src/github.com/danielmabbett/terraform-provider-jenkinsci`

```sh
mkdir -p $GOPATH/src/github.com/danielmabbett; cd $GOPATH/src/github.com/danielmabbett
git clone git@github.com:danielmabbett/terraform-provider-jenkinsci
cd $GOPATH/src/github.com/danielmabbett/terraform-provider-jenkinsci
```

Once inside the provider directory, you can run `make tools` to install the dependent tooling required to compile the provider.

At this point you can compile the provider by running `make build`, which will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-jenkinsci
...
```

You can also cross-compile if necessary:

```sh
GOOS=windows GOARCH=amd64 make build
```

In order to run the Unit Tests for the provider, you can run:

```sh
make test
```

The majority of tests in the provider are Acceptance Tests. It's possible to run the entire acceptance test suite by running `make testacc` - however it's likely you'll want to run a subset, which you can do using a prefix, by running:

```sh
make testacc TESTARGS='-run=mytest'
```

The following Environment Variables must be set in your shell prior to running acceptance tests:

- `JENKINS_ENDPOINT`
- `JENKINS_ADMIN_USERNAME`
- `JENKINS_ADMIN_PASSWORD`

Known Issues
---------------------

Due to some of the limitations of gojenkins, we presently:

- Cannot delete folders that are greater than 1 layer deep (more than a folder in a folder at root)

Contributors
---------------------

Contributors are welcome! If you have any problems/ideas, please post these into the issues page.
