
# Plugins Examples
resource "jenkinsci_plugin" "terraform" {
  name    = "Terraform"
  version = "1.0.9"
}

resource "jenkinsci_plugin" "ccm" {
  name    = "CCM"
  version = "3.2"
}

resource "jenkinsci_plugin" "ccm" {
  name    = "AnsiColor"
  version = "0.6.2"
}