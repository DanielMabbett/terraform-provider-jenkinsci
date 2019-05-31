package main

import (
	"log"

	"github.com/beevik/etree"
	jenkins "github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"folder": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				// description: "The folder you wish to place the Project within.",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validation.StringInSlice([]string{"false", "true"}, true),
			},
			"assigned_node": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	desc := d.Get("description").(string)
	assNode := d.Get("assigned_node").(string)
	disab := d.Get("disabled").(string)

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	project := doc.CreateElement("project")

	project.CreateElement("actions")
	description := project.CreateElement("description")
	description.CreateText(desc)
	keepDependencies := project.CreateElement("keepDependencies")
	keepDependencies.CreateText("false")

	// Param Definitions Section
	// parameterDefinitions := project.CreateElement("parameterDefinitions")
	// parameterDefinitions.CreateElement

	project.CreateElement("properties")
	scmclass := project.CreateElement("scm")
	scmclass.CreateAttr("class", "hudson.scm.NullSCM")

	// Assigned Node to the project
	assignedNode := project.CreateElement("assignedNode")
	assignedNode.CreateText(assNode)

	// Can roam?
	canRoam := project.CreateElement("canRoam")
	canRoam.CreateText("true")

	// Define if disabled or not
	disabled := project.CreateElement("disabled")
	disabled.CreateText(disab)

	blockBuildWhenDownstreamBuilding := project.CreateElement("blockBuildWhenDownstreamBuilding")
	blockBuildWhenDownstreamBuilding.CreateText("false")
	blockBuildWhenUpstreamBuilding := project.CreateElement("blockBuildWhenUpstreamBuilding")
	blockBuildWhenUpstreamBuilding.CreateText("false")
	triggers := project.CreateElement("triggers")
	triggers.CreateAttr("class", "vector")
	concurrentBuild := project.CreateElement("concurrentBuild")
	concurrentBuild.CreateText("false")
	project.CreateElement("builders")
	project.CreateElement("publishers")
	project.CreateElement("buildWrappers")

	str, err := doc.WriteToString()
	if err != nil {
		panic(err)
	}

	if _, ok := d.GetOk("folder"); ok {
		folder := d.Get("folder").(string)
		client.CreateJobInFolder(str, name, folder)
	} else {
		_, err := client.CreateJob(str, name)
		if err != nil {
			panic(err)
		}
	}

	d.SetId(name)
	return resourceProjectRead(d, meta)
}

func resourceProjectRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)

	// Change the name for the project
	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		client.RenameJob(oldName.(string), newName.(string))
		d.SetPartial("name")
	}

	return resourceProjectRead(d, meta)
}

func resourceProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Delete Jenkins Project %s", d.Id())
	_, err := client.DeleteJob(name)
	if err != nil {
		panic(err)
	}

	d.SetId("")
	return nil
}