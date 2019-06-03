package jenkinsci

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
			"disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"false", "true"}, true),
			},
			// "build_step": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	MaxItems: 1,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"build_step_type": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 			},
			// 			"command": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 				// ValidateFunc: validation.ValidateJsonString,
			// 				// StateFunc: func(v interface{}) string {
			// 				// 	json, _ := structure.NormalizeJsonString(v)
			// 				// 	return json
			// 				// },
			// 			},
			// 			"redirect_all_requests_to": {
			// 				Type: schema.TypeString,
			// 				ConflictsWith: []string{
			// 					"website.0.build_step_type",
			// 					"website.0.routing_rules",
			// 				},
			// 				Optional: true,
			// 			},
			// 		},
			// 	},
			// },
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
