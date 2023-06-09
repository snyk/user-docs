# Detect and report infrastructure drift

Changes to cloud workloads happen all the time. How do you know what resources are not yet managed by IaC in your cloud? And do the managed resources remain the same in the cloud as when you defined them in code?&#x20;

The `snyk iac describe` command helps report changes on managed resources and list resources that are not yet under IaC control.

You can run the `iac describe` in different modes:

* `--only-managed` / `--drift` to scan only managed resources from the Terraform states
* `--only-unmanaged` to report only unmanaged resources
* `--all` to scan both managed and unmanaged resources

See the [`snyk iac describe` command](../../../snyk-cli/commands/iac-describe.md) help for details.
