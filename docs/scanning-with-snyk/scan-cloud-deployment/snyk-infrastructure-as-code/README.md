# Snyk Infrastructure as Code

Snyk Infrastructure as Code (IaC) helps developers to write secure application configurations.  Snyk IaC provides fix advice so you can make changes directly to code, before applications reach production.

Snyk IaC helps developers write secure configurations for [HashiCorp Terraform](scan-terraform-files/), [AWS CloudFormation](scan-cloudformation-files/), [Kubernetes](scan-kubernetes-configuration-files/), and [Azure Resource Manager (ARM)](scan-arm-configuration-files.md).&#x20;

### Find and fix misconfigurations

Knowing how to securely deploy to Kubernetes, or how to securely provision infrastructure with Terraform, can be easy to get wrong, causing configuration errors (**misconfigurations**). These misconfigurations can cause security problems; for example, studies show that misconfigurations such as insecure cloud storage are the second most common error leading to breaches in the finance and insurance sector.&#x20;

#### Snyk IaC and misconfigurations

Snyk IaC integrates security checks for misconfigurations into your development lifecycle:

* The [Snyk CLI for Infrastructure as Code](snyk-cli-for-infrastructure-as-code/) provides immediate local feedback as you write configurations, so you can fix issues before you commit.
* Integrate Snyk into your CI/CD processes to automate security checks.
* Import your source repositories into Snyk for ongoing monitoring and analysis.
* Integrate with Hashicorp Terraform Cloud to scan as part of your deployment pipeline

### Snyk IaC security rules

Snyk IaC has a comprehensive set of [predefined security rules](https://snyk.io/security-rules), based on industry benchmarks, cloud-provider best practices, and threat model research from Snyk’s security intelligence team.&#x20;

You can also [build your own custom rules](custom-rules/), leveraging Open Policy Agent (OPA).
