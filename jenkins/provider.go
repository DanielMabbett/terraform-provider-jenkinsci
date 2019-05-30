package main

import (
	// "fmt"
	// "strings"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider Jenkinsci
func Provider() terraform.ResourceProvider {

	// The Provider for Jenkins
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"jenkins_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ENDPOINT", nil),
				Description: "The endpoint (URL) for the Jenkins CI Server.",
			},
			"jenkins_admin_username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ADMIN_USERNAME", nil),
				Description: "The Admin Username for the Jenkins CI Server.",
			},
			"jenkins_admin_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ADMIN_PASSWORD", nil),
				Description: "The Admin Password for the Jenkins CI Server.",
			},
			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_INSECURE", nil),
				Description: "Specify if the connection is insecure or not, default is true (http).",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"jenkinsci_project": resourceProject(),
			"jenkinsci_folder":  resourceFolder(),
			"jenkinsci_view":    resourceView(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		jenkinsEndpoint:      d.Get("jenkins_endpoint").(string),
		jenkinsAdminUsername: d.Get("jenkins_admin_username").(string),
		jenkinsAdminPassword: d.Get("jenkins_admin_password").(string),
		insecure:             d.Get("insecure").(bool),
	}

	// client, err := config.Client()
	// if err != nil {
	// 	return nil, err
	// }

	//return client, nil
	return config.Client()
}
