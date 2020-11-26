package xoa

import (
	"github.com/ddelnano/terraform-provider-xenorchestra/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("XOA_URL", nil),
				Description: "Hostname of the xoa router",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("XOA_USER", nil),
				Description: "User account for xoa api",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("XOA_PASSWORD", nil),
				Description: "Password for xoa api",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"xenorchestra_acl":          resourceAcl(),
			"xenorchestra_cloud_config": resourceCloudConfigRecord(),
			"xenorchestra_vm":           resourceRecord(),
			"xenorchestra_resource_set": resourceResourceSet(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"xenorchestra_network":      dataSourceXoaNetwork(),
			"xenorchestra_pif":          dataSourceXoaPIF(),
			"xenorchestra_pool":         dataSourceXoaPool(),
			"xenorchestra_template":     dataSourceXoaTemplate(),
			"xenorchestra_resource_set": dataSourceXoaResourceSet(),
			"xenorchestra_sr":           dataSourceXoaStorageRepository(),
		},
		ConfigureFunc: xoaConfigure,
	}
}

func xoaConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	config := client.Config{
		Url:      url,
		Username: username,
		Password: password,
	}
	return client.NewClient(config)
}
