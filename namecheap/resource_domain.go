package namecheap

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	namecheap "github.com/Namecheap-Ecosystem/go-namecheap"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// API doc: https://www.namecheap.com/support/api/methods/
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
			"add_free_who_isguard": {
				Description: "Adds free WhoisGuard for the domain",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
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

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	name := d.Get("name").(string)
	years := d.Get("years").(int)

	opts := namecheap.DomainCreateOption{}
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

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	domain, err := c.DomainGetInfo(d.Id())
	if err != nil {
		log.Fatal(err)
		return diag.FromErr(err)
	}

	if err := d.Set("name", domain.Name); err != nil {
		return diag.FromErr(err)
	}

	years, err := getDomainYears(domain)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("years", years); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("add_free_who_isguard", domain.Whoisguard.Enabled); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(domain.Name)
	return diags
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
		return nil, fmt.Errorf("failed to import domain, error:%s ", diags[0].Summary)
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

func getDomainYears(domain *namecheap.DomainInfo) (int, error) {
	// Time format is quite unique and there are no pre-defined formats
	// Ref: https://www.namecheap.com/support/api/methods/domains/get-info/
	timeFormat := "01/02/2006"
	createdAt, err := time.Parse(timeFormat, domain.Created)
	if err != nil {
		return -1, err
	}

	expiresAt, err := time.Parse(timeFormat, domain.Expires)
	if err != nil {
		return -1, err
	}

	return int(expiresAt.Sub(createdAt).Seconds() / 31207680), nil
}

func parseDomain(domain string) (string, string) {
	elements := strings.Split(domain, ".")
	return elements[0], elements[1]
}
