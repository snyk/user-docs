# Test your IaC files with Snyk CLI (Integrated IaC)

You can test your IaC files with [Integrated IaC](./) by leveraging the functionalities of the Snyk CLI, much like you would with [Current IaC](../). There are some differences between Integrated and Current IaC, which are summarized in the following table.

{% hint style="info" %}
To access the latest Integrated IaC functionalities, you should use Snyk CLI v1.1022.0 or later.
{% endhint %}

|                             | **Current IaC support** | **Integrated IaC support** |
| --------------------------- | ----------------------- | -------------------------- |
| **Terraform (single file)** | Yes                     | Yes                        |
| **Terraform (modules)**     | No                      | Yes                        |
| **Terraform (plan)**        | Yes                     | Yes                        |
| **CloudFormation**          | Yes                     | Yes                        |
| **AWS CDK**                 | Yes                     | Yes                        |
| **Kubernetes**              | Yes                     | Yes                        |
| **ARM**                     | Yes                     | Yes                        |

While Integrated IaC provides better support for Terraform than Current IaC, the following limitations apply:

* When scanning a Terraform plan, line number information is not provided in the scan results.
* Terraform modules are scanned by default, but only public modules are currently supported.
