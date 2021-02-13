package namecheap

import (
	"context"
	"fmt"
	"strings"

	namecheap "github.com/Namecheap-Ecosystem/go-namecheap"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

// API doc: https://www.namecheap.com/support/api/methods/
func resourceDomainDNS() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain": {
				Description: "Name of the domain",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hosts": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "ALIAS", "CAA", "CNAME", "MX", "MXE", "NS", "TXT", "URL", "URL301", "FRAME"}, false),
						},
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ttl": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
		CreateContext: resourceDomainDNSCreate,
		ReadContext:   resourceDomainDNSRead,
		UpdateContext: resourceDomainDNSUpdate,
		DeleteContext: resourceDomainDNSDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceDomainDNSImportState,
		},
	}
}

func resourceDomainDNSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	sld, tld := parseDomain(d.Id())
	res, err := c.DomainsDNSGetHosts(sld, tld)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("hosts", flattenHosts(res.Hosts)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.Domain)
	return diags
}

func resourceDomainDNSCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)

	sld, tld := parseDomain(d.Get("domain").(string))
	hosts := expandHosts(d)
	result, err := c.DomainDNSSetHosts(sld, tld, hosts)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.Domain)
	return resourceDomainDNSRead(ctx, d, meta)
}

func resourceDomainDNSUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)

	sld, tld := parseDomain(d.Id())
	res, err := c.DomainsDNSGetHosts(sld, tld)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("roles") {
		res.Hosts = expandHosts(d)
	}

	_, err = c.DomainDNSSetHosts(sld, tld, res.Hosts)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDomainDNSRead(ctx, d, meta)
}

func resourceDomainDNSDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	sld, tld := parseDomain(d.Id())
	_, err := c.DomainsDNSGetHosts(sld, tld)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func resourceDomainDNSImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if diags := resourceDomainDNSRead(ctx, d, meta); diags.HasError() {
		return nil, fmt.Errorf("failed to import domain nameserver")
	}

	return []*schema.ResourceData{d}, nil
}

func parseDomain(domain string) (string, string) {
	elements := strings.Split(domain, ".")
	return elements[0], elements[1]
}

func expandHosts(d *schema.ResourceData) []namecheap.DomainDNSHost {
	var hosts []namecheap.DomainDNSHost

	if v, ok := d.GetOk("hosts"); ok {
		if hs := v.(*schema.Set); hs.Len() > 0 {
			hosts = make([]namecheap.DomainDNSHost, hs.Len())

			for k, r := range hs.List() {
				hostMap := r.(map[string]interface{})
				hosts[k] = namecheap.DomainDNSHost{
					Type:    hostMap["type"].(string),
					TTL:     hostMap["ttl"].(int),
					Address: hostMap["address"].(string),
				}
			}
		}
	}

	return hosts
}

func flattenHosts(hosts []namecheap.DomainDNSHost) interface{} {
	hostList := make([]interface{}, len(hosts))
	for i, v := range hosts {
		hostList[i] = map[string]interface{}{
			"type":    v.TTL,
			"ttl":     v.TTL,
			"address": v.Address,
		}
	}

	return hostList
}
