package jenkinsci

import (
	"fmt"
	jenkins "github.com/DanielMabbett/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
	"regexp"
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

	client.InstallPlugin(name, version)
	d.SetId(name)
	return resourcePluginRead(d, meta)
}

func resourcePluginRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePluginUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	if d.HasChange("version") {
		client.InstallPlugin(d.Get("name").(string), d.Get("version").(string))
	}
	return resourcePluginRead(d, meta)
}

func resourcePluginDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	client.UninstallPlugin(name)

	return nil
}

func validatePluginVersion(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[0-9.]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("numbers and periods are only are allowed in %q: %q", k, value))
	}

	return warnings, errors
}
