# Terraform Cloud integration

## Terraform Cloud overview

[Terraform Cloud](https://www.terraform.io/cloud) (TFC) is a SaaS paid platform provided by HashiCorp that provides production-ready state management and continuous delivery for teams using Terraform. It enables teams managing their cloud infrastructure with Terraform to:

* Have Terraform state management on the cloud with versioning and out of the box
* Have a centralized place for the team to collaborate on their infrastructure, reviewing and approving changes to infrastructure
* Have Terraform Cloud manage the remote operations against the cloud providers in an automated way, similar to a CI/CD pipeline, for applying changes to cloud infrastructure using Terraform

{% hint style="info" %}
This feature is available and enabled for users of the Hashicorp **Business tier** plan\*\*.\*\* If you are on a lower Hashicorp plan, you must sign up for access to the private beta feature flag in the [**Terraform Cloud Beta Sign Up form**](http://hashi.co/tfc-beta).
{% endhint %}

## **Snyk integration with Terraform Cloud overview**

Terraform Cloud introduced a new feature called Run Checks. A “Run” in TFC represents a unit of execution in TFC that eventually generates a Terraform plan, to eventually be reviewed, approved, and applied.

This Run-check beta feature allows external integrations to connect to these “Run” events and interact with them, providing a status to determine if this run should pass or fail.

Snyk [introduced support in May 2021](https://snyk.io/blog/prevent-cloud-misconfigurations-hashicorp-terraform-snyk-iac/) for Terraform users to scan their Terraform plan json output against Snyk security policies for all major cloud providers.

The Snyk integration connects the “Run” workflow of Terraform Cloud with Snyk Terraform plan scanning, meaning that for each “Run” generated, Snyk scans the Terraform plan artifact for misconfigurations.

If you are a user of Terraform Cloud, you can sign up with Snyk to set the integration and connect it with your workspaces in Terraform Cloud. Then you'll be able track, manage, and resolve security misconfigurations as part of your software development lifecycle with Terraform Cloud.

## **How to set up and use the integration between Snyk and Terraform Cloud**

{% hint style="warning" %}
You must be an administrator on the Snyk organization to configure the Terraform Cloud integration.
{% endhint %}

Navigate to the dedicated Terraform Cloud integration settings page, under the **Integrations** page in the Snyk Web UI, then follow these steps:

### Set up Terraform plan scanning

1. Copy the provided URL and HMAC Key from the integration setting page in Snyk.
2. Navigate to [Terraform Cloud](https://app.terraform.io) and go into the organization global settings.
3. Go to the Run Tasks settings, for example,\
   `https://app.terraform.io/app/{YOUR_TFC_ORG}/settings/tasks`
4. Create a new Run Task for Snyk with the URL and HMAC key values.\
   The HMAC key is mandatory for the Snyk integration to work, even though it is identified as optional on Terraform Cloud.
5. To connect this Run Task to your workspace in Terraform Cloud, navigate to the Run Tasks settings of a specific workspace and add the newly created run task for Snyk. For example:\
   `https://app.terraform.io/app/{YOUR_TFC_ORG}/workspaces/{YOUR_WORKSPACE}/settings/tasks`
6. Choose the enforcement level (Advisory or Mandatory) and click Create.

Once your integration is set up Snyk scans Terraform plans for each run triggered in your workspace.

### View Terraform plan scanning results

1. For each run triggered in the Terraform Cloud workspace, the result of the Snyk Terraform plan scanning appears under the \`Run tasks\` step, which triggers after the Plan stage finishes.
2. The scan results in either a \`passed\` or a \`failed\` status. If Snyk finds issues in your Terraform plan file, the scan results in a failure.
3. Click on the Details link of the run task results in Terraform Cloud to view further details in Snyk.
   1. You can also find the results under the Projects tab in Snyk by searching for terraform-plan.json which will be under a Target named by `{YOUR_TFC_ORG_NAME}/{YOUR_TFC_WORKSPACE_NAME}`
   2. You can also use the filter in the left pane to show only Terraform Cloud projects
4. A single project in Snyk (`terraform-plan.json`) is created per workspace which uses the Snyk integration. Every project page shows the latest scanning results.
5. To see historical scan results, navigate into the History tab under the relevant project and choose the historic snapshot you wish to view.

### Customize Terraform plan scanning

Snyk Terraform Cloud integration provides the following levels of customization:

* Severity Threshold: Set the minimum level of severity for failure. This can be set on the integration page in Snyk.
* Custom Severities: Set custom severities for issues which overwrite the defaults (for example, [SNYK-CC-TF-63](https://snyk.io/security-rules/SNYK-CC-TF-63)).
* Enforcement Level: Determine whether a failure blocks the apply or not. This setting is controlled via Terraform Cloud. For example, the `Advisory` level does not block the apply even if Snyk finds issues within the minimum severity threshold.

### Notes and limitations

* Snyk receives an event from Terraform Cloud for each "plan" stage finished within the latest run in Terraform Cloud.
* The only way to trigger a scan is through Terraform Cloud by triggering a new "Run".
* You cannot trigger a re-scan of the Terraform plan file through the Snyk UI.
* If you customize the Snyk integration (for example, change severity threshold or customise policy severities), you must trigger a new run in Terraform Cloud for the changes to take effect in Snyk.
