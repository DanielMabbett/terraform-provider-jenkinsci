
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