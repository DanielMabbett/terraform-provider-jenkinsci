package jenkinsci

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	jenkins "github.com/DanielMabbett/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

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
			"disabled": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"pipeline_script": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

// FlowDefinition - The struct for defining a pipeline
type FlowDefinition struct {
	XMLName xml.Name `xml:"flow-definition"`
	Text    string   `xml:",chardata"`
	Plugin  string   `xml:"plugin,attr"`
	Actions struct {
		Text                                                                  string `xml:",chardata"`
		OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"`
	} `xml:"actions"`
	Description      string     `xml:"description"`
	KeepDependencies string     `xml:"keepDependencies"`
	Properties       string     `xml:"properties"`
	Definition       Definition `xml:"definition"`
	Triggers         string     `xml:"triggers"`
	Disabled         string     `xml:"disabled"`
}

type Definition struct {
	Text    string `xml:",chardata"`
	Class   string `xml:"class,attr"`
	Plugin  string `xml:"plugin,attr"`
	Script  string `xml:"script"`
	Sandbox string `xml:"sandbox"`
}

func resourcePipelineCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	pipelineScript := d.Get("pipeline_script").(string)
	disabled := strconv.FormatBool(d.Get("disabled").(bool))

	parameters := FlowDefinition{
		Plugin:           "workflow-job@2.36",
		KeepDependencies: "false",
		Disabled:         disabled,
		Definition: Definition{
			Class:   "org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition",
			Plugin:  "workflow-cps@2.76",
			Sandbox: "true",
			Script:  pipelineScript,
		},
	}

	str, err := xml.Marshal(&parameters)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("Compact XML: %s\n\n", string(str))

	_, err = client.CreateJob(string(str), name)
	if err != nil {
		return fmt.Errorf("Error creating the Jenkins Pipeline: %s,%s", err, string(str))
	}

	d.SetId(name)
	return resourcePipelineRead(d, meta)
}

func resourcePipelineRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePipelineDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Deleting Jenkins Pipeline %s", d.Id())
	_, err := client.DeleteJob(name)
	if err != nil {
		return fmt.Errorf("Error deleting the Jenkins Pipeline: %s", err)
	}

	d.SetId("")
	return nil
}
