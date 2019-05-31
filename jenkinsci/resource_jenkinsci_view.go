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

	d.SetId(view.GetName())
	return resourceViewRead(d, meta)
}

func resourceViewRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceViewUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceViewRead(d, meta)
}

func resourceViewDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
