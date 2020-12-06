# Contributing

Interested in contributing to Namecheap? Please help us!

## Setup

### Go

[Install Go](https://golang.org/doc/install). 

### Go modules

Clone the repository to a directory outside of your GOPATH:

```bash
$ git clone https://github.com/Namecheap-Ecosystem/terraform-provider-namecheap
```

Afterward, use `go build` to build the program. This will automatically fetch dependencies.

```bash
$ go build
```

Upon first build, you may see output while the `go` tool fetches dependencies.
To verify dependencies match checksums under go.sum, run `go mod verify`.
To clean up any old, unused go.mod or go.sum lines, run `go mod tidy`.

## Running Miro Provider

Create a `provider.tf`.

```hcl
provider "namecheap" {
  username = "<MASKED>"
  api_user  = "<MASKED>"
  api_token = "<MASKED>"
}
```

Build this provider.

```console
$ go build
```

Then run the Terraform operations.

```console
$ terraform init
```

## Running tests

Test the provider by

```bash
go test -v ./...
```
