package jenkinsci

import (
	"encoding/xml"
	"fmt"
	"log"

	jenkins "github.com/DanielMabbett/gojenkins"
	//"github.com/beevik/etree"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
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
	XMLName          xml.Name `xml:"flow-definition"`
	Text             string   `xml:",chardata"`
	Plugin           string   `xml:"plugin,attr"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       string   `xml:"properties"`
	Triggers         string   `xml:"triggers"`
	Disabled         string   `xml:"disabled"`
}

func resourcePipelineCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)

	parameters := FlowDefinition{
		Plugin:           "workflow-job@2.36",
		KeepDependencies: "false",
		Disabled:         "false",
	}

	str, err := xml.Marshal(&parameters)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	// fmt.Printf("Compact XML: %s\n\n", string(str))

	_, err = client.CreateJob(string(str), name)
	if err != nil {
		panic(err)
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
