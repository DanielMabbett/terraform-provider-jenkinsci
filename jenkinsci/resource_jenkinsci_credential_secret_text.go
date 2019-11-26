package jenkinsci

import (
	"fmt"

	jenkins "github.com/DanielMabbett/gojenkins"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCredentialSecretText() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialSecretTextCreate,
		Read:   resourceCredentialSecretTextRead,
		Update: resourceCredentialSecretTextUpdate,
		Delete: resourceCredentialSecretTextDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alias_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCredentialSecretTextCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	domain := d.Get("domain").(string)
	scope := d.Get("scope").(string)
	secret := d.Get("secret").(string)
	id := d.Get("alias_id").(string)
	description := d.Get("description").(string)

	cm := &jenkins.CredentialsManager{
		J:      client,
		Folder: "",
	}

	cred := &jenkins.StringCredentials{
		Scope:       scope,
		Secret:      secret,
		ID:          id,
		Description: description,
	}

	err := cm.Add(domain, cred)
	if err != nil {
		return fmt.Errorf("Error creating the secret text credentials: %s", err)
	}

	d.SetId(name)
	return resourceCredentialSecretTextRead(d, meta)
}

func resourceCredentialSecretTextRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCredentialSecretTextUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	domain := d.Get("domain").(string)
	scope := d.Get("scope").(string)
	secret := d.Get("secret").(string)
	id := d.Get("alias_id").(string)
	description := d.Get("description").(string)

	cm := &jenkins.CredentialsManager{
		J:      client,
		Folder: "",
	}

	cred := &jenkins.StringCredentials{
		Scope:       scope,
		Secret:      secret,
		ID:          id,
		Description: description,
	}

	err := cm.Update(domain, name, cred)
	if err != nil {
		return fmt.Errorf("Error updating the secret text credentials: %s", err)
	}

	d.SetId(name)
	return resourceCredentialSecretTextRead(d, meta)
}

func resourceCredentialSecretTextDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*jenkins.Jenkins)
	name := d.Get("name").(string)
	domain := d.Get("domain").(string)

	cm := &jenkins.CredentialsManager{
		J:      client,
		Folder: "",
	}

	err := cm.Delete(domain, name)
	if err != nil {
		return fmt.Errorf("Error deleting the secret text credentials: %s", err)
	}

	return nil
}
