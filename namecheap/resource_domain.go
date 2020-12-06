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
			},
		},
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		UpdateContext: resourceDomainUpdate,
		DeleteContext: resourceDomainDelete,
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

func resourceDomainCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*namecheap.Client)
	// name := data.Get("name").(string)
	// description := data.Get("description").(string)

	// req := &miro.CreateBoardRequest{
	// 	Name:        name,
	// 	Description: description,
	// }

	// board, err := c.Boards.Create(ctx, req)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	// data.SetId(board.ID)
	_ = c
	return resourceDomainRead(ctx, data, meta)
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
