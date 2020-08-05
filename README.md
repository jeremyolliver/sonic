# Sonic

![Build Passing](https://github.com/jeremyolliver/sonic/workflows/Build%20Passing/badge.svg)

Server Inventory and management tooling - built around AWS EC2 and SSM

## Features

* Detects AWS credentials via all official sources (config file, ENV variables, or assumed role with STS tokens) via official aws golang sdk.
  * To run under different roles, either export `AWS_PROFILE` or exec sonic inside an already assumed role using something like `aws-vault`
* `sonic info [identifier]` with either ec2 instance id, or SSM managed instance id. Displays information about a server.

## Planned Features

* [ ] `sonic search [identifier]` Locate servers by either: [instance-id, ip-address, name, fqdn]
* [ ] `sonic connect [identifier]` connect to a session with AWS SSM to an instance - either AWS EC2, or non-aws supported by AWS SSM advanced subscription
* [ ] `sonic issues` List issues via AWS OpsItems (like scheduled AWS instance retirements)
* [ ] Support discovery across multiple regions/accounts. For now users need to exec within a region/account where they know resources are located.

## TODO

* For non-ec2:
  - Info for known managed-instance-id: `aws ssm describe-instance-information --filters "Key=InstanceIds,Values=mi-XXXXXXXXXXXX"`
  - Search by fields such as tags `Name`, SSM's `ComputerName` (hostname, hostname -fqdn typically) cannot be filtered via this API. Query full output and cache/jq parse?
* General
  - profiles, and search order?
  - config file to enable usage of SSM?
  - config file for profiles to use, search in order?

## Installation

```
brew install go@1.13
go install github.com/jeremyolliver/sonic
```

## Development

Requirements go 1.11+ (uses https://blog.golang.org/using-go-modules)

```
make build
make test
make install # install a binary compiled from local source
```
