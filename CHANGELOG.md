# CHANGELOG

## Version 0.2.0

Some improvements including:

* Added Makefile

Fixes:

* Now supports CSRF security

Added features:

* New Resource: Create a credentials secret text with `jenkinsci_credential_secret_text`
* New Resource: Can now create a basic Jenkins pipeline item `jenkinsci_pipeline`
* Update: `jenkinsci_view` Can now delete a Jenkins View
* Update: `jenkinsci_view` Views can now have multiple `assigned_projects` as an argument

## Version 0.1.0

Initial Draft. Can:

* Create a jenkins folder
* Create / Install a jenkins plugin
* Create a jenkins project
* Create a jenkins view

There are some caveats at the moment:

* Can't delete jenkins views / folders
* Must destroy and recreate the projects
* One project per view (this shall be worked on for 0.2)
