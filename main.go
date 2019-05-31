package main

import (
	// "github.com/DanielMabbett/terraform-provider-jenkinsci/jenkins"
	"github.com/DanielMabbett/terraform-provider-jenkinsci/jenkinsci"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return jenkinsci.Provider()
		},
	})
}
