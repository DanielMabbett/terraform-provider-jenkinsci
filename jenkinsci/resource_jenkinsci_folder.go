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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_folder": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFolderCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	if _, ok := d.GetOk("parent_folder"); ok {
		// Create a nested folder
		parentFolder := d.Get("parent_folder").(string)
		nFolder, err := client.CreateFolder(name, parentFolder)
		if err != nil {
			panic(err)
		}
		d.SetId(nFolder.GetName())
	} else {
		// Create folder normally
		pFolder, err := client.CreateFolder(name)
		if err != nil {
			panic(err)
		}
		d.SetId(pFolder.GetName())
	}

	return resourceFolderRead(d, meta)
}

func resourceFolderRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFolderUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceFolderRead(d, meta)
}

func resourceFolderDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	if _, ok := d.GetOk("parent_folder"); ok {
		// Delete a nested folder
		parentFolder := d.Get("parent_folder").(string)
		fullPath := parentFolder + "/" + name
		_, err := client.DeleteJob(fullPath)
		if err != nil {
			panic(err)
		}
	} else {
		// Delete a standard folder
		_, err := client.DeleteJob(name)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
