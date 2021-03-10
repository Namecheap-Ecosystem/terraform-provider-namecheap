package namecheap

import (
	"context"
	"fmt"
	"strings"

	namecheap "github.com/Namecheap-Ecosystem/go-namecheap"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"nameservers": {
				Description: "List of custom nameservers to be associated with the domain name",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
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

func resourceDomainDNSCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)

	sld, tld := parseDomain(d.Get("domain").(string))
	nameservers := strings.Join(expandStringListFromSetSchema(d.Get("nameservers").(*schema.Set)), ",")
	res, err := c.DomainDNSSetCustom(sld, tld, nameservers)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.Domain)
	return resourceDomainDNSRead(ctx, d, meta)
}

func resourceDomainDNSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	sld, tld := parseDomain(d.Id())
	res, err := c.DomainDNSGetList(sld, tld)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("nameservers", res.Nameservers); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.Domain)
	return diags
}

func resourceDomainDNSUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)

	sld, tld := parseDomain(d.Id())
	nameservers := strings.Join(expandStringListFromSetSchema(d.Get("nameservers").(*schema.Set)), ",")
	_, err := c.DomainDNSSetCustom(sld, tld, nameservers)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDomainDNSRead(ctx, d, meta)
}

func resourceDomainDNSDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	sld, tld := parseDomain(d.Id())
	if _, err := c.DomainDNSSetDefault(sld, tld, []namecheap.DomainDNSHost{}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func resourceDomainDNSImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if diags := resourceDomainDNSRead(ctx, d, meta); diags.HasError() {
		return nil, fmt.Errorf("failed to import domain dns")
	}

	return []*schema.ResourceData{d}, nil
}
