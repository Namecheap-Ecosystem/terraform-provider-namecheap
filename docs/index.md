# Namecheap Provider

The [Namecheap](https://www.namecheap.com/) provider is used to interact with the
Namecheap resources. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Namecheap provider
provider "namecheap" {
  username  = var.namecheap_username
  api_user  = var.namecheap_api_user
  api_token = var.namecheap_api_token
}
```

## Argument Reference

The following arguments are supported:

* `username`  - (Required) Username of the user. This can also be set via the `NAMECHEAP_USERNAME` environment variable.
* `api_user`  - (Required) User of the API token. This can also be set via the `NAMECHEAP_API_USER` environment variable.
* `api_token` - (Required) Credential. This can also be set via the `NAMECHEAP_API_TOKEN` environment variable.
* `url`       - (Optional) The URL of the endpoint. Defaults to Production API URL. You can specify the sandbox API. This can also be set via the `NAMECHEAP_URL` environment variable.
