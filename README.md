# Sonic

Server Inventory and management tooling - built around AWS SSM

## Planned Features

* [x] CLI interface integrated with AWS-SDK
* [ ] `sonic search [identifier]` Locate servers by either: [instance-id, ip-address, name, fqdn]
* [ ] `sonic connect [identifier]` connect to a session with AWS SSM to an instance - either AWS EC2, or non-aws supported by AWS SSM advanced subscription
* `sonic issues` List issues via AWS OpsItems (like scheduled AWS instance retirements)
* [x] `sonic info [identifier]` Show information about an instance - often email notifications regarding instance replacements are tedious to locate and summarise info like: account, environment, deployment group to ascertain whether any action will be needed to handle replacement.
* [ ] Support multiple accounts via STS Session profiles

## TODO

* For non-ec2:
  - Info for known managed-instance-id: `aws ssm describe-instance-information --filters "Key=InstanceIds,Values=mi-XXXXXXXXXXXX"`
  - Search by fields Name, ComputerName (hostname, hostname -fqdn typically) cannot be filtered via this API. Query full output and cache/jq parse?
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
```
