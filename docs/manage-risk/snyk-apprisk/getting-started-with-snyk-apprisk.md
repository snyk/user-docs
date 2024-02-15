# Getting started with Snyk AppRisk

## Prerequisites

Before you get started, ensure that you meet the following prerequisites:&#x20;

* You are a Snyk Enterprise customer.
* Your account is entitled with access for Snyk AppRisk Essentials.
* You are a Group Administrator for the Group associated with Snyk AppRisk, or you are assigned a Group level role with permissions to View Group and Edit AppRisk.&#x20;
* The Group associated with Snyk AppRisk includes organizations that have onboarded Snyk application security products.
* You have the necessary permissions and authority to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk AppRisk for repository asset discovery.

{% hint style="info" %}
When you integrate a Git code repository with Snyk AppRisk, you should use a secondary token with a broad, complete view of the code repository, not only of what you imported into Snyk. \
Use a secondary token to have a counterview of everything onboarded using Snyk. \
Using the secondary token reduces the likelihood of introducing a blindspot from a limited token at the Organization level configuration. \
The first import, synchronization, can take up to 24 hours to complete.
{% endhint %}

## Permissions

You can access Snyk AppRisk with one of the Group level roles permissions described below. To access the permissions, navigate to **View groups**, then select the **Snyk AppRisk permissions** option.

1. View AppRisk - Grants you a read-only access to AppRisk.
2. Edit AppRisk - Grants you edit access to AppRisk, for example, edit policies, edit asset classification, and add the integration.

A Group Administrator has the **Edit AppRisk** permission assigned by default, and a Group Viewer has the **View AppRisk** permission assigned by default.

{% hint style="info" %}
For more information on default user roles and permissions, see [Default user roles](../../snyk-admin/user-roles/pre-defined-roles.md).
{% endhint %}

## Login and Authentication

Login and authenticate to Snyk using existing mechanisms (SSO, Google SAML, and so on).

## Accessing Snyk AppRisk

Select the Group with Snyk AppRisk enabled. A link to Snyk AppRisk appears in the left navigation menu. Click this link to open the application in a new browser window.

{% hint style="info" %}
Keep both Snyk Web UI and Snyk AppRisk tabs open to ensure optimal functionality.
{% endhint %}

## Key Concepts

**Asset**: An asset is an identifiable entity that is part of an application, and relevant for security and developers.&#x20;

**Controls**: The security controls associated with the asset. Navigate to the [Coverage controls](policies-for-snyk-apprisk/use-cases-for-policies/coverage-control-policy-use-case.md) section to see all available statuses for security controls.

**Coverage**: An assessment of whether applicable assets are scanned and tested by security tools (like Snyk Open Source, for instance), as it relates to an application security program.  A type of policy that allows you to specify what controls should be applied and, optionally, how often it needs to be run.

**Tags**: A way to categorize assets. Helps you recognize or handle assets differently according to mutual properties. Assets can be filtered by their tags in the inventory or when creating policy rules. A tag can be automatically assigned to an asset, or the asset can be tagged by a policy you created. GitHub and GitLab topics are treated as asset tags and you can use them for creating policies.&#x20;

**Class**: A way to assign business context to assets and categorize an asset based on the business criticality. Assets can be assigned Classes A, B, C, or D, where Class A (assets that are business critical, deal with sensitive data, subject to compliance, and so on) is the most important and Class D (test apps, sandbox environments, and so on) the least important. Assets are assigned Class C by default. A class can be used in policies as well as defined in a policy.

**Policy**: A way to automate actions in certain conditions, like classifying and tagging assets with business context. You can also use a policy to configure actions like sending a message or setting the coverage gap control using a Policy builder UI.

## Scanning methods

You can initiate a scan from the Web UI, the CLI, the API, or with PR Checks. See the [Start scanning using the CLI, Web UI, or API](../../scan-with-snyk/start-scanning-using-the-cli-web-ui-or-api.md) page for more details.

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





\


\
