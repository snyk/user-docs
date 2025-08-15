# Terraform Cloud integration for Snyk IaC using Run Tasks

{% hint style="info" %}
Terraform Cloud run tasks are available in the Terraform Cloud Team & Governance tier. The Terraform Cloud Free tier does not support run tasks.
{% endhint %}

## Terraform Cloud overview

[Terraform Cloud](https://www.terraform.io/cloud) (TFC) is a SaaS paid platform provided by HashiCorp that provides production-ready state management and continuous delivery for teams using Terraform. It enables teams managing their cloud infrastructure with Terraform to:

* Have Terraform state management on the cloud with versioning and out-of-the box
* Have a centralized place for the team to collaborate on their infrastructure, reviewing and approving changes to infrastructure
* Have Terraform Cloud manage the remote operations against the cloud providers in an automated way, similar to a CI/CD pipeline, for applying changes to cloud infrastructure using Terraform

## **Snyk integration with Terraform Cloud overview**

Terraform Cloud has a Run Tasks feature available to customers with Run Task entitlement on their Terraform plan. This entitlement is an optional add-on to the Team plan and above. A `run` in TFC represents a unit of execution in TFC that eventually generates a Terraform plan to be reviewed, approved, and applied.

The Run Tasks feature allows external integrations to connect to `run` events and interact with them, providing a status to determine if this run should pass or fail.

Snyk [introduced support in May 2021](https://snyk.io/blog/prevent-cloud-misconfigurations-hashicorp-terraform-snyk-iac/) for Terraform users to scan their Terraform plan JSON output against Snyk security policies for all major cloud providers.

The Snyk integration connects the `run` workflow of Terraform Cloud with Snyk Terraform plan scanning, meaning that for each `run` enerated, Snyk scans the Terraform plan artifact for misconfigurations.

If you are a user of Terraform Cloud ([Team & Governance tier](https://www.hashicorp.com/products/terraform/pricing)), you can sign up with Snyk to set the integration and connect it with your workspaces in Terraform Cloud. If you have the Run Task entitlement add-on, you can track, manage, and resolve security misconfigurations as part of your software development lifecycle with Terraform Cloud.

For setup and use details, see the following pages:

* [Set up the Terraform Cloud integration for IaC](set-up-the-terraform-cloud-integration-for-iac.md)
* [How to use the Terraform Cloud integration for IaC](how-to-use-the-terraform-cloud-integration-for-iac.md)
