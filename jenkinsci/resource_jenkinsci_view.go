package jenkinsci

import (
	"github.com/bndr/gojenkins"
	jenkins "github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceView() *schema.Resource {
	return &schema.Resource{
		Create: resourceViewCreate,
		Read:   resourceViewRead,
		Update: resourceViewUpdate,
		Delete: resourceViewDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"assigned_project": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			// "view_type": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			// "description": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
		},
	}
}

func resourceViewCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	view, err := client.CreateView(name, gojenkins.LIST_VIEW)
	if err != nil {
		panic(err)
	}

	if _, ok := d.GetOk("assigned_project"); ok {
		assignedProject := d.Get("assigned_project").(string)
		view.AddJob(assignedProject)
	}

	d.SetId(view.GetName())
	return resourceViewRead(d, meta)
}

func resourceViewRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceViewUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("assigned_project") {
		client := meta.(*jenkins.Jenkins)
		oldProj, newProj := d.GetChange("assigned_project")

		v, err := client.GetView(d.Get("assigned_project").(string))
		if err != nil {
			panic(err)
		}

		v.DeleteJob(oldProj.(string))
		v.AddJob(newProj.(string))

		d.SetPartial("assigned_project")
	}
	return resourceViewRead(d, meta)
}

func resourceViewDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}