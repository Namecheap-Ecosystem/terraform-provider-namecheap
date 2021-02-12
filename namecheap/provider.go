package namecheap

import (
	"context"

	namecheap "github.com/billputer/go-namecheap"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Description: "Name of the user",
				Required:    true,
			},
			"api_user": {
				Type:        schema.TypeString,
				Description: "User of the API token",
				Required:    true,
			},
			"api_token": {
				Type:        schema.TypeString,
				Description: "Token for the API",
				Required:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "URL of the API endpoint",
				Optional:    true,
				Default:     "https://api.namecheap.com/xml.response",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"namecheap_domain": resourceDomain(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigureFunc,
	}
}

func providerConfigureFunc(_ context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	username := data.Get("username").(string)
	apiUser := data.Get("api_user").(string)
	apiToken := data.Get("api_token").(string)

	c := namecheap.NewClient(apiUser, apiToken, username)
	if v, ok := data.GetOk("url"); ok {
		c.BaseURL = v.(string)
	}

	return c, diags
}
