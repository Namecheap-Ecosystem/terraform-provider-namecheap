package namecheap

import (
	"context"
	"errors"
	"fmt"

	namecheap "github.com/billputer/go-namecheap"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// API doc: https://www.namecheap.com/support/api/methods/domains/create/
func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the domain",
				Type:        schema.TypeString,
				Required:    true,
			},
			"years": {
				Description: "Number of years to register",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
			},
			"nameservers": {
				Description: "List of custom nameservers to be associated with the domain name",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"add_free_who_isguard": {
				Description: "Adds free WhoisGuard for the domain",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"wg_enabled": {
				Description: "Enables free WhoisGuard for the domain",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		UpdateContext: resourceDomainUpdate,
		DeleteContext: resourceDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceDomainImportState,
		},
	}
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	domain, err := c.DomainGetInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("nameservers", domain.DNSDetails.Nameservers); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("add_free_who_isguard", domain.Whoisguard); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("nameservers", domain.DNSDetails.Nameservers); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(domain.Name)
	return diags
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	name := d.Get("name").(string)
	years := d.Get("years").(int)
	nameservers := expandStringListFromSetSchema(d.Get("nameservers").(*schema.Set))

	opts := namecheap.DomainCreateOption{}
	if len(nameservers) > 0 {
		opts.Nameservers = nameservers
	}

	if v, ok := d.GetOk("add_free_who_isguard"); ok {
		opts.AddFreeWhoisguard = v.(bool)
	}

	if v, ok := d.GetOk("wg_enabled"); ok {
		opts.AddFreeWhoisguard = v.(bool)
	}

	result, err := c.DomainCreate(name, years, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.Domain)
	return resourceDomainRead(ctx, d, meta)
}

func resourceDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)

	domain, err := c.DomainGetInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	var years int
	if d.HasChange("years") {
		years = d.Get("years").(int)
	}

	if _, err := c.DomainRenew(domain.Name, years); err != nil {
		return diag.FromErr(err)
	}

	return resourceDomainRead(ctx, d, meta)
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(errors.New("this resource can't be deleted because the Namecheap API does not provide this operation. Please delete the actual resource and remove from the state"))
}

func resourceDomainImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if diags := resourceDomainRead(ctx, d, meta); diags.HasError() {
		return nil, fmt.Errorf("failed to import domain")
	}

	return []*schema.ResourceData{d}, nil
}

func expandStringListFromSetSchema(list *schema.Set) []string {
	res := make([]string, list.Len())
	for i, v := range list.List() {
		res[i] = v.(string)
	}

	return res
}
