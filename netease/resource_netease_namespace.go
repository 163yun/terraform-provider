package netease

import (
	"fmt"
	"github.com/bingohuang/163yun-go-sdk/cloudcomb"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceNeteaseNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceNamespaceCreate,
		Read:   resourceNamespaceRead,
		Update: resourceNamespaceUpdate,
		Delete: resourceNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceNamespaceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cloudcomb.CloudComb)

	keyName := d.Get("name").(string)
	params := `{
				"name": "%s"
			}`
	params = fmt.Sprintf(params, keyName)

	id, err := client.CreateNamespace(params)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprint(id))

	log.Printf("[INFO] Create Namespace ID: %v", id)

	return resourceNamespaceRead(d, m)

	return nil
}

func resourceNamespaceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cloudcomb.CloudComb)

	result, err := client.GetNamespaceServices(d.Id(), -1, -1)
	if err != nil {
		return err
	}
	d.Set("result", result)

	return nil
}

func resourceNamespaceUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceNamespaceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cloudcomb.CloudComb)

	log.Printf("[INFO] Deleting Namespace: %s", d.Id())

	err := client.DeleteNamespace(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting namespace: %s", err)
	}

	d.SetId("")

	return nil
}
