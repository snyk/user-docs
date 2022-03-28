# Share CLI results with the Snyk App

By using the `snyk iac test` command, you are already set up to address known configuration issues.&#x20;

## See results in the Snyk App

To see these in the Snyk App, run the following:

* `snyk iac test myproject --report` , or
* `snyk iac report myproject`

Example:

```
> snyk iac test myproject --report

Testing arm-file.tf...


Infrastructure as code issues:
  ✗ VM Agent is not provisioned automatically for Windows [Low Severity] [SNYK-CC-AZURE-667] in Compute
    introduced by resource > azurerm_virtual_machine[my_terraformvm] > os_profile_windows_config > provision_vm_agent


Organization:      my.org
Type:              Terraform
Target file:       arm-file.tf
Project name:      myproject
Open source:       no
Project path:      myproject

Tested arm-file.tf for known issues, found 1 issues

Your test results are available at: https://app.snyk.io/org/my.org/projects under the name myproject
```

This sends a snapshot of your current configuration issues to your Snyk dashboard. \
Log in to Snyk App ([https://app.snyk.io](https://app.snyk.io/login)) and go to your organization's project page, to see the specific snapshot of your scanned project.

![Newly scanned project in the projects page](<../../.gitbook/assets/image (88).png>)

You can also open the project itself and see the usual individual project view:

![individual project view](<../../.gitbook/assets/image (86).png>)

## **Ignores**

You can ignore issues using the Snyk App, or by creating a `.snyk` policy file along with your project when scanning. For more information, see [Iac Ignores using the .snyk policy file](snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file.md).

**Note:** issues that are ignored by using the `.snyk` policy file can not be unignored in the Snyk App.

## **Notes**

Using [Custom rules](custom-rules/) and the Share Results functionality together is not currently supported.

Snyk does not share any file contents over the network, only the required metadata for the configuration issues that were just scanned.

For a full CLI commands, see the [CLI Commands help page](../../snyk-cli/commands/).

For additional information see the [CLI reference](../../snyk-cli/cli-reference/).
