package jenkinsci

import (
	"fmt"
	"regexp"

	jenkins "github.com/DanielMabbett/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePlugin() *schema.Resource {
	return &schema.Resource{
		Create: resourcePluginCreate,
		Read:   resourcePluginRead,
		Update: resourcePluginUpdate,
		Delete: resourcePluginDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validatePluginVersion,
			},
		},
	}
}

func resourcePluginCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	version := d.Get("version").(string)

	err := client.InstallPlugin(name, version)
	if err != nil {
		return fmt.Errorf("Error installing the Jenkins plugin: %s", err)
	}

	d.SetId(name)
	return resourcePluginRead(d, meta)
}

func resourcePluginRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePluginUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)

	if d.HasChange("version") {
		err := client.InstallPlugin(d.Get("name").(string), d.Get("version").(string))
		if err != nil {
			return fmt.Errorf("Error installing the Jenkins plugin: %s", err)
		}
	}

	return resourcePluginRead(d, meta)
}

func resourcePluginDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	err := client.UninstallPlugin(name)
	if err != nil {
		return fmt.Errorf("Error installing the Jenkins plugin: %s", err)
	}

	return nil
}

func validatePluginVersion(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if !regexp.MustCompile(`^[0-9.]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("numbers and periods are only are allowed in %q: %q", k, value))
	}

	return warnings, errors
}
