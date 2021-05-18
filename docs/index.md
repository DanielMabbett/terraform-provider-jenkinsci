---
layout: "jenkinsci"
page_title: "Provider: jenkinsci"
sidebar_current: "docs-jenkinsci-index"
description: |-
  Jenkins provider
---

# jenkinsci Provider

[jenkinsci](https://www.jenkins.io/) is an open source continous integration tool. The
jenkinsci provider exposes resources to interact with a jenkins instance.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Using the Provider
```hcl
provider "jenkinsci" {
  jenkins_endpoint         = "..."
  jenkins_admin_username   = "..."
  jenkins_admin_password   = "..."
}
```

## Argument Reference

The following arguments are supported:

- `jenkins_endpoint` `(string: "http://127.0.0.1:4646")` - The HTTP(S) API address of the
  jenkinsci agent. This must include the leading protocol (e.g. `https://`).

- `jenkins_admin_username` `(string)` - The Username for the Admin account associated with the Jenkins instance.

- `jenkins_admin_password` `(string)` - The Password for the Admin account associated with the Jenkins instance.
