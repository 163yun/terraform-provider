// bingo@2017-05-02

package netease

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"app_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETEASE_ACCESS_KEY", nil),
				Description: "App key of 163yun",
			},
			"app_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETEASE_ACCESS_SECRET", nil),
				Description: "App secret of 163yun",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"netease_secret_key": resourceNeteaseSecretKey(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AppKey:    d.Get("app_key").(string),
		AppSecret: d.Get("app_secret").(string),
	}

	client, err := config.Client()
	if err != nil {
		log.Println("[INFO] Initializing 163yun client fail.")
		return nil, err
	}

	log.Println("[INFO] Initializing 163yun client success.")

	return client, nil
}
