# Upgrading to the Enterprise plan

The purpose of this page is to help you if you are currently using Snyk on either the Free or Team plan and are interested in upgrading to the Enterprise plan or ready to upgrade. For information about plan options, see the Snyk [Plans and pricing page](https://snyk.io/plans/?utm\_medium=Paid-Search\&utm\_source=google\&utm\_campaign=GS\_SN:\_Brand\&utm\_content=BR\_Pricing\&utm\_term=synk%20price\&gclid=Cj0KCQiA\_bieBhDSARIsADU4zLf8Dvv9pa39ofNqQLTd35KUjmfOTABUAGFImlAYnn2P\_f\_HcJtD4ksaAvsgEALw\_wcB).

For information about how to implement Snyk on the Enterprise plan, see:

* [Getting started with the Snyk Enterprise plan](getting-started-with-the-snyk-enterprise-plan.md)
* [Preparing for Implementation: Enterprise Plan](preparing-for-implementation-of-the-enterprise-plan.md)

## Overview of features of the Snyk Enterprise plan

Snyk Enterprise features include:

* [Snyk Groups with multiple Snyk Organizations](../snyk-admin/manage-groups-and-organizations/introduction-to-groups.md)
* [Single Sign-On](using-single-sign-on-sso-for-authentication/)
* [Service accounts](service-accounts/)
* [Snyk API](../snyk-api/)
* [GitHub Enterprise integration](../integrations/git-repository-scm-integrations/github-enterprise-integration.md)
* [Reports](../manage-issues/reporting/legacy-reports/)
* [Security policy management](../manage-issues/policies/security-policies/)
* [Snyk Broker](snyk-broker/)

## Checklist for upgrading to the Snyk Enterprise plan

As you work through the rest of this page, use the following checklist to ensure you have completed the steps to upgrade to the Enterprise plan. There are links to the relevant documentation throughout the page.

* [ ] Set up your Snyk environment:
  * [ ] Set your Snyk Group settings.
  * [ ] Structure your account with multiple Snyk Organizations if applicable.
* [ ] Define your Snyk Project structure.
* [ ] Use the Snyk API.
* [ ] Set up SSO if applicable.
* [ ] If you are using SSO, delete social logins.
* [ ] Implement service accounts.
* [ ] Consider setting up security policies.
* [ ] Use Snyk Reports.
* [ ] Consider using GitHub Enterprise integration.
* [ ] Consider setting up Snyk Broker for on-premise or private cloud code repositories, container registries, Jira instances, and other special integrations.

## Set up your Snyk environment

A notable difference between the Free/Team plan and the Enterprise plan is having a [Snyk Group](../snyk-admin/manage-groups-and-organizations/introduction-to-groups.md) and the ability to create multiple [Snyk Organizations](../snyk-admin/manage-groups-and-organizations/whats-a-snyk-organization.md).

### Set your Snyk Group settings

{% hint style="warning" %}
Only [Group Administrators](../snyk-admin/manage-permissions-and-roles/permissions-associated-with-each-pre-defined-role.md) can edit Snyk Group settings.
{% endhint %}

Confirm your Snyk Group name. It should reflect your company name. If it needs to be updated, go to **Group Settings > General**.

Set [session expiration](../snyk-admin/manage-users-in-organizations-and-groups/configure-session-length-for-a-snyk-group.md) within the Group. This will be the default for all Snyk Organizations within the Group.

### Structure your account with multiple Snyk Organizations

When you plan to use multiple Organizations in Snyk, decide how to structure your account before you invite members to the account or scan a large number of Snyk Projects. To make this decision, consider the following:

* Which team members can access specific Snyk Projects?
* How do you want to apply policies to Projects for prioritizing and automating tests?
* How do you want to report on Projects?
* Specifically, what level of granularity would you like to see in the reports?

{% hint style="info" %}
See the [Snyk account structure](https://training.snyk.io/courses/snyk-account-structure) training course for more help with structuring your account.
{% endhint %}

### Create additional Snyk Organizations

To create a new Organization, open the Organization switcher in the navigation panel and select **Create new Organization**. See [Manage Snyk organizations](../snyk-admin/manage-groups-and-organizations/manage-organizations.md) for more details.

You can **use an existing Organization as a template.**

After you configure your first Snyk Organization, you can use it as a template for creating additional Organizations. Make sure you have **completed the configurations for the Organization before copying it**.

Select the Organization you want to use as a template from your list of Organizations. The following settings will be copied from the selected Organization:

* Source control integrations (GitHub, GitLab, BitBucket)
* Container registry integrations (ACR, Docker Hub, ECR, GCR)
* Container orchestrator integrations (Kubernetes)
* PaaS and Serverless integrations (Heroku, AWS Lambda)
* Notification integrations (Slack, Jira)
* Policies
* Ignore settings
* Language settings
* Snyk Code settings
* Infrastructure as Code settings

The new Organization will not use the same settings from the copied Organization for the following:

* Service accounts
* Members
* Projects
* Notification preferences

Any of the Organization settings that you configured for the first Organization can then be customized for the new Organization.

### Enable Snyk Code for each Snyk Organization

When you create a new Organization, Snyk Code (SAST) scanning is disabled by default. Snyk recommends enabling Snyk Code before you import your first Projects to Snyk. If Snyk Code is enabled after a Project is imported, Snyk will not detect Snyk Code files.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to enable Snyk Code, then click **Save changes**

## Define your Snyk Project structure

You may want to house different Projects in different Organizations.

{% hint style="warning" %}
You cannot move Projects between Organizations directly.
{% endhint %}

To re-house Projects, you must re-import your Projects into your new Organizations.

If you are working with a large number of Projects, you should use the API to re-import Projects. For details, see [Use the Snyk API](upgrading-to-the-enterprise-plan.md#use-the-snyk-api).

## Use the Snyk API

Access to the [Snyk API](../snyk-api/) allows you to scale and automate different processes, including importing Projects. You will want to use this strategy if you have a large number of Projects that need to be moved to your new Snyk Organizations.

You can use the [Snyk API import tool](../snyk-api-info/other-tools/tool-snyk-api-import/) to import Projects into Snyk at a controlled pace using available Snyk APIs.

{% hint style="info" %}
If you plan to run SAST scans with Snyk Code, check that Snyk Code is enabled within the related Snyk Organization settings before importing Snyk Code Projects.
{% endhint %}

## Set up Single Sign-On (SSO)

Users can authenticate with their Snyk accounts in several ways, such as with a GitHub or Google account. However, now that you have Enterprise access, you may want to set up Single Sign-On (SSO) via your existing identity provider to streamline sign-ins and provisioning of new users.

{% hint style="info" %}
If you are using SSO, after you set it up, you must remove any duplicate users from social login accounts. These are users who have logged in with other methods than SSO.
{% endhint %}

See [Setting up Single Sign-On (SSO) for authentication](using-single-sign-on-sso-for-authentication/) for more details on steps for using your identity provider to authenticate and provision Snyk to your teams.

See the [SSO, authentication, and user provisions](https://training.snyk.io/courses/sso) training course.

## Implement Service accounts

[Service accounts](service-accounts/) allow you to better scale and automate.

* If you use CI/CD, either by using a [CI/CD integration](../integrations/snyk-ci-cd-integrations/) or the [Snyk CLI](../snyk-cli/), Snyk recommends that you use a service account.
* If you use an IDE plugin or the CLI to test in your local environment, Snyk recommends that you use personal access tokens.

## Consider setting up security policies

Security policies allow you to customize the prioritization of speciFor more details, see issues from the default and create rules. This is particularly helpful for changing the severities of issues that are not relevant to a specific Project or environment.

For more information, see [Security policies](../manage-issues/policies/security-policies/).

## Use Snyk Reports

Snyk Reports offer data and analytics across all of your Projects, displaying historical and aggregated data about Projects, issues, dependencies, and licenses.

See [Getting started with Snyk Reports](../manage-issues/reporting/getting-started-with-snyk-reports.md).

## Consider using GitHub Enterprise integration

If you use GitHub as your SCM with the Snyk Enterprise plan, you now have access to GitHub Enterprise integration. The GitHub Enterprise integration allows the use of a single personal access token across a Snyk Organization rather than a personal access token tied to an individual user account. Because of this, Snyk recommends switching from GitHub integration to GitHub Enterprise integration.

{% hint style="info" %}
You do not need a GitHub Enterprise license or subscription to use the GitHub Enterprise integration within Snyk.
{% endhint %}

See [Using Github or Github Enterprise integration](../integrations/git-repository-scm-integrations/using-github-or-github-enterprise-integration.md) for steps in migrating from Github to Github Enterprise integration.

## Consider using Snyk Broker

If you have on-premise repositories, you can scan them with Snyk using [Snyk Broker](snyk-broker/). Consider setting up Snyk Broker for on-premise or private cloud code repositories, container registries, and on-premises Jira, JFrog Artifactory, and Nexus instances.

Using Snyk Broker you. can scan Snyk Projects in:

* An SCM that is not internet reachable
* A publicly-accessible SCM, allowing you to view and control Snyk activity for increased data security

For details, see the [Snyk Broker](snyk-broker/) documentation.

