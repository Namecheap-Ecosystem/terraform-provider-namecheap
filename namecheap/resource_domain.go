package namecheap

import (
	"context"

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
			"registrant": {
				Description: "Registrant configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: getDomainRegistrantSchema(),
				},
			},
			"tech": {
				Description: "Tech user configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: getDomainTechSchema(),
				},
			},
			"admin": {
				Description: "Admin user configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: getDomainAdminSchema(),
				},
			},
			"aux_billing": {
				Description: "Aux Billing configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: getDomainAuxBillingSchema(),
				},
			},
			"billing": {
				Description: "Billing configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: getDomainBillingSchema(),
				},
			},
			"extended_attribute": {
				Description: "Extended attribute",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		UpdateContext: resourceDomainUpdate,
		DeleteContext: resourceDomainDelete,
	}
}

func getDomainRegistrantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Type:        schema.TypeString,
			Description: "First name of the registrant",
			Required:    true,
		},
		"last_name": {
			Type:        schema.TypeString,
			Description: "Last name of the registrant",
			Required:    true,
		},
		"city": {
			Type:        schema.TypeString,
			Description: "City of the registrant address",
			Required:    true,
		},
		"state": {
			Type:        schema.TypeString,
			Description: "State province of the registrant address",
			Required:    true,
		},
		"postal_code": {
			Type:        schema.TypeString,
			Description: "Postal code of the registrant address",
			Required:    true,
		},
		"address_1": {
			Type:        schema.TypeString,
			Description: "Address 1 of the registrant address",
			Required:    true,
		},
		"country": {
			Type:        schema.TypeString,
			Description: "Country of the registrant address",
			Required:    true,
		},
		"phone": {
			Type:        schema.TypeString,
			Description: "Phone number",
			Required:    true,
		},
		"email": {
			Type:        schema.TypeString,
			Description: "Email address of the registrant user",
			Required:    true,
		},
	}
}

func getDomainTechSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Type:        schema.TypeString,
			Description: "First name of the registrant",
			Required:    true,
		},
		"last_name": {
			Type:        schema.TypeString,
			Description: "Last name of the registrant",
			Required:    true,
		},
		"city": {
			Type:        schema.TypeString,
			Description: "City of the registrant address",
			Required:    true,
		},
		"state": {
			Type:        schema.TypeString,
			Description: "State province of the registrant address",
			Required:    true,
		},
		"postal_code": {
			Type:        schema.TypeString,
			Description: "Postal code of the registrant address",
			Required:    true,
		},
		"address_1": {
			Type:        schema.TypeString,
			Description: "Address 1 of the registrant address",
			Required:    true,
		},
		"country": {
			Type:        schema.TypeString,
			Description: "Country of the registrant address",
			Required:    true,
		},
		"phone": {
			Type:        schema.TypeString,
			Description: "Phone number",
			Required:    true,
		},
		"email": {
			Type:        schema.TypeString,
			Description: "Email address of the registrant user",
			Required:    true,
		},
	}
}

func getDomainAdminSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Type:        schema.TypeString,
			Description: "First name of the registrant",
			Required:    true,
		},
		"last_name": {
			Type:        schema.TypeString,
			Description: "Last name of the registrant",
			Required:    true,
		},
		"city": {
			Type:        schema.TypeString,
			Description: "City of the registrant address",
			Required:    true,
		},
		"state": {
			Type:        schema.TypeString,
			Description: "State province of the registrant address",
			Required:    true,
		},
		"postal_code": {
			Type:        schema.TypeString,
			Description: "Postal code of the registrant address",
			Required:    true,
		},
		"address_1": {
			Type:        schema.TypeString,
			Description: "Address 1 of the registrant address",
			Required:    true,
		},
		"country": {
			Type:        schema.TypeString,
			Description: "Country of the registrant address",
			Required:    true,
		},
		"phone": {
			Type:        schema.TypeString,
			Description: "Phone number",
			Required:    true,
		},
		"email": {
			Type:        schema.TypeString,
			Description: "Email address of the registrant user",
			Required:    true,
		},
	}
}

func getDomainAuxBillingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Type:        schema.TypeString,
			Description: "First name of the registrant",
			Required:    true,
		},
		"last_name": {
			Type:        schema.TypeString,
			Description: "Last name of the registrant",
			Required:    true,
		},
		"city": {
			Type:        schema.TypeString,
			Description: "City of the registrant address",
			Required:    true,
		},
		"state": {
			Type:        schema.TypeString,
			Description: "State province of the registrant address",
			Required:    true,
		},
		"postal_code": {
			Type:        schema.TypeString,
			Description: "Postal code of the registrant address",
			Required:    true,
		},
		"address_1": {
			Type:        schema.TypeString,
			Description: "Address 1 of the registrant address",
			Required:    true,
		},
		"country": {
			Type:        schema.TypeString,
			Description: "Country of the registrant address",
			Required:    true,
		},
		"phone": {
			Type:        schema.TypeString,
			Description: "Phone number",
			Required:    true,
		},
		"email": {
			Type:        schema.TypeString,
			Description: "Email address of the registrant user",
			Required:    true,
		},
	}
}

func getDomainBillingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Type:        schema.TypeString,
			Description: "First name of the registrant",
			Required:    true,
		},
		"last_name": {
			Type:        schema.TypeString,
			Description: "Last name of the registrant",
			Required:    true,
		},
		"city": {
			Type:        schema.TypeString,
			Description: "City of the registrant address",
			Required:    true,
		},
		"state": {
			Type:        schema.TypeString,
			Description: "State province of the registrant address",
			Required:    true,
		},
		"postal_code": {
			Type:        schema.TypeString,
			Description: "Postal code of the registrant address",
			Required:    true,
		},
		"address_1": {
			Type:        schema.TypeString,
			Description: "Address 1 of the registrant address",
			Required:    true,
		},
		"country": {
			Type:        schema.TypeString,
			Description: "Country of the registrant address",
			Required:    true,
		},
		"phone": {
			Type:        schema.TypeString,
			Description: "Phone number",
			Required:    true,
		},
		"email": {
			Type:        schema.TypeString,
			Description: "Email address of the registrant user",
			Required:    true,
		},
	}
}

func resourceDomainRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics

	// board, err := c.Boards.Get(ctx, data.Id())
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	// if board == nil {
	// 	data.SetId("")
	// 	return diags
	// }

	// if err := data.Set("boards", board); err != nil {
	// 	return diag.FromErr(err)
	// }

	// data.SetId(board.ID)
	_ = c
	return diags
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	name := d.Get("name").(string)
	years := d.Get("years").(int)

	result, err := c.DomainCreate(name, years)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.Domain)
	return resourceDomainRead(ctx, d, meta)
}

func resourceDomainUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	// id := data.Id()
	// name := data.Get("name").(string)
	// description := data.Get("description").(string)

	// req := &miro.UpdateBoardRequest{
	// 	Name:        name,
	// 	Description: description,
	// }

	// _, err := c.Boards.Update(ctx, id, req)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	_ = c
	return resourceDomainRead(ctx, data, meta)
}

func resourceDomainDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	var diags diag.Diagnostics
	// if err := c.Boards.Delete(ctx, data.Id()); err != nil {
	// 	return diag.FromErr(err)
	// }

	// data.SetId("")
	_ = c
	return diags
}
