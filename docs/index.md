# Namecheap Provider

The [Namecheap](https://www.namecheap.com/) provider is used to interact with the
Namecheap resources. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Namecheap provider
provider "namecheap" {
  username  = "${var.username}"
  api_user  = "${var.api_user}"
  api_token = "${var.api_token}"
}
```

## Argument Reference

The following arguments are supported:

* `username`  - (Required) Username of the user.
* `api_user`  - (Required) User of the API token.
* `api_token` - (Required) Credential.
* `url`       - (Optional) The URL of the endpoint. Defaults to Production API URL. You can specify the sandbox API.
