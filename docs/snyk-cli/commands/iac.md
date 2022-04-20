# IAC

## Usage

`snyk iac <SUBCOMMAND> [<OPTIONS>] [<PATH>]`

## Description

The `snyk iac` subcommands find and report security issues in Infrastructure as Code files; detect, track, and alert on infrastructure drift and unmanaged resources; and create a .driftigore file.

For more information see [Snyk CLI for Infrastructure as Code](https://docs.snyk.io/products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code)

## Subcommands with their options

Help docs for all the `iac` commands are listed here:

* [iac describe](iac-describe.md): detects infrastructure drift and unmanaged cloud resources
  * Example: `snyk iac describe --only-unmanaged`&#x20;
* [iac update-exclude-policy](iac-update-exclude-policy.md): auto-generates .snyk exclusions for cloud resources
  * Example: `snyk iac describe --json --all | snyk iac update-exclude-policy`
