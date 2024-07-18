# Using Snyk AppRisk

## Prerequisites

Before you get started, ensure that you meet the following prerequisites:&#x20;

* You are a Snyk Enterprise customer.
* Your account is entitled with access for Snyk AppRisk Essentials or Snyk AppRisk Pro.
* You are a Group Administrator for the Group associated with Snyk AppRisk, or you are assigned a Group level role with permissions to View Group and Edit AppRisk.&#x20;
* The Group associated with Snyk AppRisk includes organizations that have onboarded Snyk application security products.
* You have the necessary permissions and authority to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk AppRisk for repository asset discovery.

{% hint style="info" %}
When you integrate a Git code repository with Snyk AppRisk, you should use a secondary token with a broad, complete view of the code repository, not only of what you imported into Snyk. \
Use a secondary token to have a counterview of everything onboarded using Snyk. \
Using the secondary token reduces the likelihood of introducing a blindspot from a limited token at the Organization level configuration. \
The first import, synchronization, can take up to 24 hours to complete.
{% endhint %}

The following video is an overview of what Snyk AppRisk can offer:

{% embed url="https://www.youtube.com/embed/23y6pEnTfqQ" %}

## Permissions

You can access Snyk AppRisk with one of the Group level roles permissions described below. To access the permissions, navigate to **View groups**, then select the **Snyk AppRisk permissions** option.

1. View AppRisk - Grants you a read-only access to AppRisk.
2. Edit AppRisk - Grants you edit access to AppRisk, for example, edit policies, edit asset classification, and add the integration.

A Group Administrator has the **Edit AppRisk** permission assigned by default, and a Group Viewer has the **View AppRisk** permission assigned by default.

{% hint style="info" %}
For more information on default user roles and permissions, see [Default user roles](../../snyk-admin/manage-permissions-and-roles/pre-defined-roles.md).
{% endhint %}

## Login and Authentication

Login and authenticate to Snyk using existing mechanisms (SSO, Google SAML, and so on).

## Accessing Snyk AppRisk

Ensure you are at the Group level to access the Snyk AppRisk options. From the Group level you have a centralized security management that enhances security and simplifies security procedures for projects.

The following video presents an overview of the Snyk AppRisk interfaces:

{% embed url="https://www.youtube.com/watch?v=OUoukT5518c" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)!
{% endembed %}

## Key Concepts

**Asset**: meaningful, real-world components in an application’s SDLC, where meaningful means either carries a risk or aggregates risk of other components (for example, repositories that contain packages) and real-world means that the concept exists outside of Snyk, for example, repository (which is a generally applicable term).

**Controls**: The security controls associated with the asset. Navigate to the [Coverage controls](../../manage-risk/policies/assets-policies/use-cases-for-policies/coverage-control-policy-use-case.md) section to see all available statuses for security controls.

**Coverage**: An assessment of whether applicable assets are scanned and tested by security tools (like Snyk Open Source, for instance), as it relates to an application security program.  A type of policy that allows you to specify what controls should be applied and, optionally, how often it needs to be run.

**Tags**: A way to categorize assets. Helps you recognize or handle assets differently according to mutual properties. Assets can be filtered by their tags in the inventory or when creating policy rules. A tag can be automatically assigned to an asset, or the asset can be tagged by a policy you created. GitHub and GitLab topics are treated as asset tags and you can use them for creating policies.&#x20;

**Class**: A way to assign business context to assets and categorize an asset based on the business criticality. Assets can be assigned Classes A, B, C, or D, where Class A (assets that are business critical, deal with sensitive data, subject to compliance, and so on) is the most important and Class D (test apps, sandbox environments, and so on) the least important. Assets are assigned Class C by default. A class can be used in policies as well as defined in a policy.

**Policy**: A way to automate actions in certain conditions, like classifying and tagging assets with business context. You can also use a policy to configure actions like sending a message or setting the coverage gap control using a Policy builder UI.

## Scanning methods

You can initiate a scan from the Web UI, the CLI, the API, or with PR Checks. See [Start scanning](../../scan-with-snyk/start-scanning-using-the-cli-web-ui-or-api.md) for more details.

If you initiate your scans using the CLI, you might encounter one of the following situations:

1. If you have a `.git` folder available in the directory that the CLI is scanning, then the `git remoteurl` is picked up automatically for Snyk Open Source, Snyk Container, and Snyk IaC.&#x20;

{% hint style="info" %}
Snyk Code does not automatically pick up the `git remoteurl`, even if the `.git` folder is available in the directory scanned by the CLI.
{% endhint %}

2. If you do not have a `.git` folder available in the directory that the CLI is scanning, you can use different test or monitor commands to achieve the same result:
   * [`snyk monitor`](../../snyk-cli/commands/monitor.md#remote-repo-url-less-than-url-greater-than), for Snyk Open Source
   * [`snyk iac test`](../../snyk-cli/commands/iac-test.md#remote-repo-url-less-than-url-greater-than) - also requires the `--report` command
   * `snyk container monitor` - momentarily, no options are available.
   * `snyk code test` - momentarily, no options are available.



{% hint style="info" %}
The Asset Dashboard menu option is available only for Snyk AppRisk Essentials users.&#x20;

If you are using Snyk AppRisk Pro, navigate to the [Application Analytics ](../../manage-risk/enterprise-analytics/aspm-analytics/application-analytics.md)page.
{% endhint %}

## Asset dashboard

On the Snyk AppRisk dashboard page, you can add widgets that display an overview of your application and security controls.

You can customize the dashboard widgets as desired. Choose to rename or modify display configurations, or add multiple widget instances.

### Configure a widget

Customize your dashboard with the available widgets. You can change the settings of an existing widget or change the way it is displayed. You have the ability to move a widget around the dashboard, rename it, display or hide the legend, view it in full screen, export or download it. See the [Repositories assets](../../manage-assets/#repository-assets) section for additional details regarding archived or deleted repositories and the information presented in the dashboard widgets.

### Settings menu

You can make several changes to a widget. All widgets allow you to change the name. Other particular settings are available for each widget. You can access the Settings menu by following these steps:

1. Select a widget and click the **Setting** menu.
2. &#x20;Customize the widget by changing its name or other specific details.&#x20;
3. After all changes are done, click **Apply**.

### Action menu

Access the full list of general options from the action menu. You can access the Action menu by following these steps:

1. Select a widget and click the **Action** menu.
2. Select one of the following actions:
   * View in full-screen
   * CSV export
   * XLS export
   * Download PNG
   * Download PDF



\


\


\
