# namecheap_domain Resource

Provides a Namecheap domain resource.

## Example Usage

```hcl
# Create a new Namecheap domain
resource "namecheap_domain" "mydomain" {
    name        = "mydomain.com"
    // TODO
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Pipeline name.
* `description` - (Optional) Pipeline JSON content.

## Import

Namecheap domain can be imported by the domain name.

```
$ terraform import namecheap_domain.mydomain mydomain.com
```
