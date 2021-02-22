# namecheap_domain_dns Resource

Provides a Namecheap domain DNS resource.

## Example Usage

```hcl
# Create a new Namecheap domain DNS
resource "namecheap_domain_dns" "mydns" {
    domain = "mydomain.com"
    
    hosts = {
        ttl     = 300
        type    = "CNAME"
        address = "namecheap.com"
    }
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The domain name.
* `hosts` - (Required) List of the custom hosts to configure.


* `hosts`: Argument for hosts
    * `ttl` - (Required) TTL for the host record.
    * `type` - (Required) Type of the record .Possible values are `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `MXE`, `NS`, `TXT`, `URL`, `URL301`, `FRAME`.
    * `address` - (Required) Possible values are URL or IP address.

## Import

Namecheap domain can be imported by the domain name.

```
$ terraform import namecheap_domain.mydomain mydomain.com
```
