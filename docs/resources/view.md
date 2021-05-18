# jenkinsci_view

Manages a view in Jenkins.

## Example Usage

```hcl
resource "jenkinsci_view" "view" {
  name             = "1st-view"
}

resource "jenkinsci_project" "test" {
  name = "test-project-1a"
}

# A view with an assigned project in the view. Only works with 1 project assigned so far
resource "jenkinsci_view" "test2" {
  name             = "2nd-view"
  assigned_projects = [
    jenkinsci_project.test.name
  ]
}

```

## Argument Reference

The following arguments are supported:

- `name` `(string: <required>)` - A unique name for the view.
- `assigned_projects` `(list(string): <optional>)` - A list of assigned projects to the view
