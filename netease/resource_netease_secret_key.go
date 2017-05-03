package netease

import (
	"fmt"
	"github.com/bingohuang/cloudcomb-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceNeteaseSecretKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretKeyCreate,
		Read:   resourceSecretKeyRead,
		Update: resourceSecretKeyUpdate,
		Delete: resourceSecretKeyDelete,

		Schema: map[string]*schema.Schema{
			"key_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSecretKeyCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cloudcomb.CloudComb)

	keyName := d.Get("key_name").(string)
	params := `{
				"key_name": "%s"
			}`
	params = fmt.Sprintf(params, keyName)
	id, name, err := client.CreateSecretKey(params)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprint(id))
	d.Set("name", name)

	log.Printf("[INFO] Create Secret Key ID: %v", id)

	return resourceSecretKeyRead(d, m)
}

func resourceSecretKeyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cloudcomb.CloudComb)

	content, err := client.GetSecretKey(d.Id())
	if err != nil {
		return err
	}
	d.Set("content", content)

	return nil
}

func resourceSecretKeyUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceSecretKeyRead(d, m)
}

func resourceSecretKeyDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cloudcomb.CloudComb)

	log.Printf("[INFO] Deleting Secret Key: %s", d.Id())

	err := client.DeleteSecretKey(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting secret key: %s", err)
	}

	d.SetId("")

	return nil
}
