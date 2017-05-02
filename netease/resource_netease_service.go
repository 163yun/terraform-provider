// bingo@2017-05-02

package netease

import "github.com/hashicorp/terraform/helper/schema"

func resourceNeteaseService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address + "!")
	return nil
}

func resourceServiceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("address") {
		// Try updating the address
		if err := updateAddress(d, m); err != nil {
			return err
		}

		d.SetPartial("address")
	}

	// If we were to return here, before disabling partial mode below,
	// then only the "address" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// save all fields again.
	d.Partial(false)

	return nil
}

func updateAddress(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address + "!!")
	return nil
}

func resourceServiceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
