# namecheap_domain_dns Resource

Provides a Namecheap domain DNS resource.

## Example Usage

```hcl
# Create a new Namecheap domain DNS
resource "namecheap_domain_dns" "mydns" {
    domain = "mydomain.com"
    
    nameservers = [
        "aws.com",
        "google.com",
        "azure.com",
    ]
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The domain name.
* `nameservers` - (Required) The custom nameservers

## Import

Namecheap domain can be imported by the domain name

```console
$ terraform import namecheap_domain_dns.mydns mydomain.com
```
