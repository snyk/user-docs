# Integrating Snyk with Terraform Cloud

## Terraform Cloud overview

[Terraform Cloud](https://www.terraform.io/cloud) (TFC) is SaaS paid platform provided by HashiCorp, providing production ready State Management & Continuous Delivery for teams using Terraform. It enables teams managing their Cloud Infrastructure with Terraform to:

* Have Terraform state management on the cloud with versioning and out of the box.
* Have a centralised place for the team to collaborate on their infrastructure, reviewing and approving changes to infrastructure.
* Have Terraform Cloud manage the remote operations against the cloud providers in an automated way, similarly to a CI/CD pipeline for applying changes to cloud infrastructure using Terraform.

{% hint style="info" %}
Since this feature is behind a private beta feature flag on Terraform Cloud, in order to get it enabled, you will have to **sign up for access in the **[**Terraform Cloud Beta Sign Up form**](http://hashi.co/tfc-beta).
{% endhint %}

## **Snyk Integration with Terraform Cloud overview**

Terraform Cloud introduced a new feature called Run Checks. A “Run” in TFC represents a unit of execution in TFC that eventually generates a Terraform plan, to eventually be reviewed, approved, and applied. 

This Run-check beta feature allows external integrations to connect to these “Run” events and interact with them, providing a status to determine if this run should pass or fail.

Snyk [introduced support in May 2021](https://snyk.io/blog/prevent-cloud-misconfigurations-hashicorp-terraform-snyk-iac/) for Terraform users to scan their terraform-plan json output against our security policies for all major cloud providers.

The Snyk integration connects the “Run” workflow of Terraform Cloud with Snyk’s Terraform Plan scanning, meaning that for each “Run” generated, Snyk scans the Terraform Plan artifact for misconfigurations.

Users of Terraform Cloud should sign up to Snyk to set the integration, and connect it with their workspaces in Terraform Cloud. Then, you'll be able track, manage, and resolve security misconfigurations as part of your software development lifecycle with Terraform Cloud.

## **How to setup the integration between Snyk and Terraform Cloud**

Navigate into the dedicated Terraform Cloud integration settings page, under the Integrations page in the Snyk app, then follow the next steps:

### Setup Terraform plan scanning:

1. Copy the provided URL and HMAC Key from the integration setting page in Snyk.
2. Navigate to [Terraform Cloud](https://app.terraform.io) and go into the organization global settings.
3. Go to the Task Event Hooks settings. For example:\
   `https://app.terraform.io/app/{YOUR_TFC_ORG}/settings/event-hooks`
4. Create a new Task Event Hook for Snyk with the URL and HMAC key values.\
   (The HMAC key is mandatory for the Snyk integration to work, even though it's stated as optional on Terraform Cloud)
5. To connect this Task Event Hook to your workspace in Terraform Cloud, navigate to the Tasks settings of a specific workspace and add the newly created Task-Event Hook for Snyk. For example:\
   `https://app.terraform.io/app/{YOUR_TFC_ORG}/workspaces/{YOUR_WORKSPACE}/settings/tasks`
6. Choose the enforcement level (Advisory or Mandatory) and click Create.

Now your integration is set up, you will get Terraform plans scanned by Snyk for each run triggered in your workspace.

### Viewing Terraform Plan scanning results:

1. For each run triggered in the Terraform Cloud workspace, the result of Snyk’s Terraform plan scanning appears under the \`Run tasks\` step, which triggers after the Plan stage finishes.
2. The scan results in either \`passed\` or a \`failed\` status. If Snyk finds issues in your Terraform plan file, it results in a failure.
3. Clicking on the Details link of the run-task results in Terraform Cloud to view further details in Snyk.
   1. You can also find the results under the Projects tab in Snyk by searching for terraform-plan.json which will be under a Target named by `{YOUR_TFC_ORG_NAME}/{YOUR_TFC_WORKSPACE_NAME}`, you can also use the filter in the left pane to show only Terraform Cloud projects
4. A single project in Snyk (`terraform-plan.json`) is created per workspace which uses the Snyk integration. Every project page shows the latest scanning results.
5. To see historical scan results, navigate into the History tab under the relevant project and choose the historic snapshot you wish to view.

### Customizing Terraform Plan scanning:

Snyk’s integration provides the following levels of customization:

* Severity Threshold: decides the minimum level of severity for failure. This can be set on the integration page in Snyk.
* Custom Severities: set custom severities for issues which overwrite the defaults. (for example,  [SNYK-CC-TF-63](https://snyk.io/security-rules/SNYK-CC-TF-63))
* Enforcement Level: decides on whether a failure will block the apply or not. This setting is controlled via Terraform Cloud.\
  (for example, the `Advisory` level will not block the apply even if Snyk founds issues within the minimum severity threshold)

### Notes and limitations

* Snyk receives an event from Terraform Cloud for each "Plan" stage finished within the latest run in Terraform Cloud.
* The only way to trigger a scan is via Terraform Cloud by triggering a new "Run".
* There's no way to trigger a re-scan to the Terraform Plan file via the Snyk UI.
* If you customise the Snyk integration (for example, change severity threshold or customise policy severities), you must trigger a new run in Terraform Cloud for the changes to take effect in Snyk.
