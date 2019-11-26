package jenkinsci

import (
	"fmt"
	jenkins "github.com/DanielMabbett/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceView() *schema.Resource {
	return &schema.Resource{
		Create: resourceViewCreate,
		Read:   resourceViewRead,
		Delete: resourceViewDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"assigned_projects": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceViewCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	view, err := client.CreateView(name, jenkins.LIST_VIEW)
	if err != nil {
		return fmt.Errorf("Error creating the Jenkins View: %s", err)
	}

	assigedProjects := d.Get("assigned_projects").([]interface{})
	for _, project := range assigedProjects {
		view.AddJob(project.(string))
	}

	d.SetId(view.GetName())
	return resourceViewRead(d, meta)
}

func resourceViewRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceViewDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	_, err := client.DeleteView(name)
	if err != nil {
		return fmt.Errorf("Error deleting the Jenkins View: %s", err)
	}

	return nil
}
