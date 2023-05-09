# Test your IaC files with Snyk CLI (Integrated IaC)

{% hint style="info" %}
To access the latest Integrated IaC, you should use Snyk CLI v1.1022.0 or later.
{% endhint %}

You can test your IaC files with [Integrated IaC](./) by leveraging the functionalities of the Snyk CLI, much like you would with [Current IaC](../../snyk-infrastructure-as-code/). There are some differences between Integrated and Current IaC, which are summarized in the following table.

|                             | **Current IaC support** | **Integrated IaC support** |
| --------------------------- | ----------------------- | -------------------------- |
| **Terraform (single file)** | Yes                     | Yes                        |
| **Terraform (modules)**     | No                      | Yes                        |
| **Terraform (plan)**        | Yes                     | Yes                        |
| **CloudFormation**          | Yes                     | Yes                        |
| **AWS CDK**                 | Yes                     | Yes                        |
| **Azure Resource Manager**  | Yes                     | Yes                        |
| **Kubernetes**              | Yes                     | Coming soon                |

{% hint style="warning" %}
Although Integrated IaC provides better support for Terraform than Current IaC, line number information is not provided in the scan results.
{% endhint %}

## Stage 1: Test IaC files

Snyk Infrastructure as Code allows you to test your configuration files with the CLI. For information on how to use the `snyk iac test` command, see the [CLI documentation](../../../snyk-cli/commands/iac-test.md).

```
Snyk Infrastructure as Code

✔ Test completed.

Issues

Low Severity Issues: 1

  [Low] API Gateway access logging disabled
  Info:    Amazon Api Gateway access logging is not enabled. Audit records may not be available during investigation
  Rule:    https://snyk.io/security-rules/SNYK-CC-TF-138
  Path:    resource > aws_api_gateway_stage[denied] > access_log_settings
  File:    aws_api_gateway_stage_logging.tf
  Resolve: Set `access_log_settings` attribute

-------------------------------------------------------

Test Summary

  Organization: demo-org

✔ Files without issues: 0
✗ Files with issues: 1
  Invalid files: 0
  Ignored issues: 0
  Total issues: 1 [ 0 critical, 0 high, 0 medium, 1 low ]
```

{% hint style="info" %}
The CLI for Integrated IaC is able to scan Terraform modules too, regardless of whether they're public or private modules. Just run `terraform init` before running the `snyk iac test` command and the CLI will read the generated `.terraform files.`
{% endhint %}

## Stage 2: View IaC issues in the Snyk Web UI

You can use the [CLI](../../../snyk-cli/) `snyk iac test` command to address known configuration issues.

To see these issues displayed in the Snyk Web UI, run the following CLI command:

`snyk iac test myproject --report`

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

Your test results are available at: https://snyk.io/org/my.org/cloud/issues?environment_name=my.org
```

If you click on the link in the CLI output, you will be shown your issues in the Snyk Web UI. To learn more about the Cloud Issues view, see [View cloud issues in the Snyk Web UI](../snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md).

## Stage 3: Fix IaC Issues

Act on the recommendations produced by Snyk IaC.

1. After you have run a test scan, you will find all the relevant details about where that issue exists as well as advice on how to remediate that issue.
2. Fix the issue based on the recommended advice.
3. Run another test scan to see if the issue has been resolved.
4. (Optional) View a list of all Integrated IaC and Snyk Cloud rules and adjust rule severity as needed. See [Managing Snyk Cloud rules](../managing-snyk-cloud-rules.md).

## Get help

To get help, run `snyk iac test --help` or visit the [online documentation](../../../snyk-cli/commands/iac-test.md).

## Get started with Terraform Cloud

To use Integrated IaC with Terraform Cloud, see [Terraform Cloud integration for IaC](https://docs.snyk.io/integrations/ci-cd-integrations/integrating-snyk-with-terraform-cloud).

