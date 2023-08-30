# Detect and report infrastructure drift

{% hint style="info" %}
**Deprecation notice: Drift detection of managed resources**\
Drift detection of managed resources, including `snyk iac describe --only-managed and snyk iac describe --drift` has been deprecated. The end-of-life date for drift detection of managed resources is September 30. 2023.
{% endhint %}

Changes to cloud workloads happen all the time. How do you know what resources are not yet managed by IaC in your cloud? And do the managed resources remain the same in the cloud as when you defined them in code?&#x20;

The `snyk iac describe` command helps report changes on managed resources and list resources that are not yet under IaC control.

You can run the `iac describe` in different modes:

* `--only-managed` / `--drift` to scan only managed resources from the Terraform states
* `--only-unmanaged` to report only unmanaged resources
* `--all` to scan both managed and unmanaged resources

See the [`snyk iac describe` command](../../../snyk-cli/commands/iac-describe.md) help for details.
