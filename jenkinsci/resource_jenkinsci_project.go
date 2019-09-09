package jenkinsci

import (
	"log"
	"strings"

	"github.com/beevik/etree"
	jenkins "github.com/DanielMabbett/gojenkins"
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"folder": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"assigned_node": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"additional_config": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"false", "true"}, true),
			},
			"parameter": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
							ForceNew: true,
						},
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, meta interface{}) error {
	// vars
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	desc := d.Get("description").(string)
	assNode := d.Get("assigned_node").(string)
	disab := d.Get("disabled").(string)

	// Start by Creating a new document with xml layer0 as "project", and include encoding params
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	project := doc.CreateElement("project")

	project.CreateElement("actions")
	description := project.CreateElement("description")
	description.CreateText(desc)
	keepDependencies := project.CreateElement("keepDependencies")
	keepDependencies.CreateText("false")

	scmclass := project.CreateElement("scm")
	scmclass.CreateAttr("class", "hudson.scm.NullSCM")

	// Assigned Node to the project
	assignedNode := project.CreateElement("assignedNode")
	assignedNode.CreateText(assNode)

	// Define canRoam settings
	canRoam := project.CreateElement("canRoam")
	canRoam.CreateText("true")

	// Define Disabled Settings
	disabled := project.CreateElement("disabled")
	disabled.CreateText(disab)

	// Define blockBuildWhenDownstreamBuilding
	blockBuildWhenDownstreamBuilding := project.CreateElement("blockBuildWhenDownstreamBuilding")
	blockBuildWhenDownstreamBuilding.CreateText("false")

	// Define blockBuildWhenUpstreamBuilding
	blockBuildWhenUpstreamBuilding := project.CreateElement("blockBuildWhenUpstreamBuilding")
	blockBuildWhenUpstreamBuilding.CreateText("false")

	// Define triggers
	triggers := project.CreateElement("triggers")
	triggers.CreateAttr("class", "vector")

	// Define Concurrenty builds
	concurrentBuild := project.CreateElement("concurrentBuild")
	concurrentBuild.CreateText("false")

	// Create other elements that weren't filled
	project.CreateElement("builders")
	project.CreateElement("publishers")
	project.CreateElement("buildWrappers")

	project.CreateElement("properties")
	// Param Definitions Section
	// If Parameter Block Specified then add that
	if _, ok := d.GetOk("parameter"); ok {
		properties := project.SelectElement("properties")
		p := properties.CreateElement("hudson.model.ParametersDefinitionProperty")
		p0 := p.CreateElement("parameterDefinitions")
		p1 := p0.CreateElement("hudson.model.StringParameterDefinition")

		p2a := p1.CreateElement("name")
		// parameter := d.Get("parameter").(*schema.Set).List()
		// config := parameter[0].(map[string]interface{})
		// key := config["key"].(string)
		p2a.CreateText("test")

		p2b := p1.CreateElement("description")
		p2b.CreateText("the description for param")

		p2c := p1.CreateElement("defaultValue")
		p2c.CreateText("the default value")

		p2d := p1.CreateElement("trim")
		p2d.CreateText("false")
	}

	str, err := doc.WriteToString()
	if err != nil {
		panic(err)
	}

	if _, ok := d.GetOk("additional_config"); ok {
		replacement := d.Get("additional_config").(string)
		fullReplacement := replacement + "</project>"
		newstr := strings.Replace(str, "</project>", fullReplacement, -1)
		if _, ok := d.GetOk("folder"); ok {
			folder := d.Get("folder").(string)
			client.CreateJobInFolder(newstr, name, folder)
		} else {
			_, err := client.CreateJob(newstr, name)
			if err != nil {
				panic(err)
			}
		}
	} else {
		// This was the original output method
		if _, ok := d.GetOk("folder"); ok {
			folder := d.Get("folder").(string)
			client.CreateJobInFolder(str, name, folder)
		} else {
			_, err := client.CreateJob(str, name)
			if err != nil {
				panic(err)
			}
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
