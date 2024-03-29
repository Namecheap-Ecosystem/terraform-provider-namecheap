package namecheap

import (
	"context"

	namecheap "github.com/Namecheap-Ecosystem/go-namecheap"

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
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_USERNAME", nil),
			},
			"api_user": {
				Type:        schema.TypeString,
				Description: "User of the API token",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_API_USER", nil),
			},
			"api_token": {
				Type:        schema.TypeString,
				Description: "Token for the API",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_API_TOKEN", nil),
			},
			"url": {
				Type:        schema.TypeString,
				Description: "URL of the API endpoint",
				Optional:    true,
				Default:     "https://api.namecheap.com/xml.response",
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"namecheap_domain":     resourceDomain(),
			"namecheap_domain_dns": resourceDomainDNS(),
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
