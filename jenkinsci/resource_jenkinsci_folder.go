package jenkinsci

import (
	jenkins "github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceFolderCreate,
		Read:   resourceFolderRead,
		Update: resourceFolderUpdate,
		Delete: resourceFolderDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFolderCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	// Create folder
	pFolder, err := client.CreateFolder(name)
	if err != nil {
		panic(err)
	}

	d.SetId(pFolder.GetName())
	return resourceFolderRead(d, meta)
}

func resourceFolderRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFolderUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceFolderRead(d, meta)
}

func resourceFolderDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
