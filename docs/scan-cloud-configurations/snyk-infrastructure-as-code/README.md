# Snyk Infrastructure as Code

Snyk Infrastructure as Code (IaC) helps developers to write secure application configurations. Snyk IaC provides fix advice so you can make changes directly to code before applications reach production.

Snyk IaC helps developers write secure configurations for [HashiCorp Terraform](scan-terraform-files/), [AWS CloudFormation](scan-cloudformation-files/), [Kubernetes](scan-kubernetes-configuration-files/), and [Azure Resource Manager (ARM)](scan-arm-configuration-files.md).

## Find and fix misconfigurations

Securely deploying to Kubernetes or provisioning infrastructure with Terraform can be easy to get wrong, causing configuration errors (**misconfigurations**). These misconfigurations can cause security problems. Studies show that misconfigurations, such as insecure cloud storage, are the second most common error leading to breaches in the finance and insurance sector.

## Snyk IaC and misconfigurations

To integrate security checks for misconfigurations into your development lifecycle, you can use Snyk IaC in the following ways:

* Use the [Snyk CLI for Infrastructure as Code](snyk-cli-for-infrastructure-as-code/) to get immediate local feedback as you write configurations, so you can fix issues before you commit.
* Integrate Snyk into your CI/CD processes to automate security checks.
* Import your source repositories into Snyk for ongoing monitoring and analysis.
* Integrate with HashiCorp Terraform Cloud to scan as part of your deployment pipeline.

## Snyk IaC security rules

Snyk IaC has a comprehensive set of predefined [security rules](https://security.snyk.io/rules/cloud/) based on industry benchmarks, cloud provider best practices, and threat model research from Snyk’s security intelligence team.

You can also [build your own custom rules](custom-rules/), leveraging Open Policy Agent (OPA).
