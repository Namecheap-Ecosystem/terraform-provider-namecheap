# namecheap_domain Resource

Provides a Namecheap domain resource.

## Example Usage

```hcl
# Create a new Namecheap domain
resource "namecheap_domain" "mydomain" {
    name = "mydomain.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The domain name.
* `years` - (Optional) Number of years to register, defaults to `1` years.
* `nameservers` - (Optional) List of custom nameservers to be associated with the domain name.
* `add_free_who_isguard` - (Optional) Adds free WhoisGuard for the domain.
* `wg_enabled` - (Optional) Enables free WhoisGuard for the domain.

## Import

Namecheap domain can be imported by the domain name.

```
$ terraform import namecheap_domain.mydomain mydomain.com
```
