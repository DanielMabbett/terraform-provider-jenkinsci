package jenkinsci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider Jenkinsci
func Provider() terraform.ResourceProvider {

	// The Provider for Jenkins
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"jenkins_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ENDPOINT", nil),
				Description: "The endpoint (URL) for the Jenkins CI Server.",
			},
			"jenkins_admin_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ADMIN_USERNAME", nil),
				Description: "The Admin Username for the Jenkins CI Server.",
			},
			"jenkins_admin_password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ADMIN_PASSWORD", nil),
				Description: "The Admin Password for the Jenkins CI Server.",
			},
			"insecure": {
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
			"jenkinsci_plugin":  resourcePlugin(),
		},

		ConfigureFunc: providerConfigure,
	}
}

// Provider - Configure it
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		jenkinsEndpoint:      d.Get("jenkins_endpoint").(string),
		jenkinsAdminUsername: d.Get("jenkins_admin_username").(string),
		jenkinsAdminPassword: d.Get("jenkins_admin_password").(string),
		insecure:             d.Get("insecure").(bool),
	}

	//return client, nil
	return config.Client()
}
