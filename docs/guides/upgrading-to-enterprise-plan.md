# Upgrading to Enterprise plan

Read this guide if you are currently using Snyk on either the Free or Team plan, and are ready to upgrade to the Enterprise plan (see [Plans & Pricing page](https://snyk.io/plans/?utm\_medium=Paid-Search\&utm\_source=google\&utm\_campaign=GS\_SN:\_Brand\&utm\_content=BR\_Pricing\&utm\_term=synk%20price\&gclid=Cj0KCQiA\_bieBhDSARIsADU4zLf8Dvv9pa39ofNqQLTd35KUjmfOTABUAGFImlAYnn2P\_f\_HcJtD4ksaAvsgEALw\_wcB) for plan options).

For full information about how to implement Snyk on the Enterprise plan, read:

* [Getting started with Snyk: Enterprise plan](getting-started-with-snyk-enterprise-plan.md)
* [Preparing for Implementation: Enterprise Plan](preparing-for-implementation-enterprise-plan.md)

#### Enterprise features overview

Snyk Enterprise features include:

* [Snyk Group with multiple orgs](../snyk-admin/manage-groups-and-organizations/whats-a-snyk-group.md)
* [Single Sign-On](../snyk-admin/set-up-single-sign-on-sso-for-authentication/)
* [Service accounts](../snyk-admin/service-accounts.md)
* [Rich API](../snyk-api-info/)
* [Reports](../manage-issues/snyk-reports/)
* [Security policy management](../manage-issues/policies/security-policies/)
* [Snyk Broker](../snyk-admin/snyk-broker/)
* [GitHub Enterprise integration](../integrations/git-repository-scm-integrations/github-enterprise-integration.md)

### Upgrade to Enterprise checklist

As you work through the rest of this guide, use the following checklist to ensure you have completed the steps to upgrade to the Enterprise plan.

* [ ] Set up your Snyk environment
  * [ ] Set your Snyk Group settings
  * [ ] Structure your account with multiple Snyk organizations (if applicable)
* [ ] Define project structure
* [ ] Set up SSO (if applicable).
  * [ ] Delete social logins
* [ ] Consider setting up security policies
* [ ] Consider using Github Enterprise integration
* [ ] Consider setting up Snyk Broker for onprem or private cloud code repositories, container registries, and Jira instances

### Set up your Snyk environment

A notable difference between the Free/Team plan and the Enterprise plan is having a [Snyk Group](../snyk-admin/manage-groups-and-organizations/whats-a-snyk-group.md) and the ability to create multiple [Snyk Organizations](../snyk-admin/manage-groups-and-organizations/whats-a-snyk-organization.md).

#### Set your Snyk Group settings

{% hint style="warning" %}
Only [Group Administrators](../snyk-admin/manage-users-and-permissions/managing-permissions.md) can edit Snyk Group settings
{% endhint %}

Confirm your Snyk Group name. It should reflect your company name. If it needs to be updated go to **Group Settings > General**.

Set [session expiration](../snyk-admin/manage-users-and-permissions/session-length.md) within the Group. This will be the default for all Snyk Organizations within the Group.

#### Structure your account with multiple Snyk organizations

When you plan to use multiple Organizations in Snyk, decide how to structure your account before you invite members to the account or scan a large number of projects. To make this decision, consider the following:

* Which team members can access specific projects?
* How do you want to apply policies to projects for prioritizing and automating tests?
* How do you want to report on projects?
  * What is the level of granularity you would like to see in the reports section?

{% hint style="info" %}
See the [Snyk account structure](https://training.snyk.io/courses/snyk-account-structure) training course for more help with structuring your account.
{% endhint %}

### Create additional Organizations

To create a new Organization, open the Organization switcher in the navigation panel and select **Create new Organization**. See [Manage Snyk organizations](../snyk-admin/manage-groups-and-organizations/manage-snyk-organizations.md) for more details.

#### Using an existing Organization as a template

After you configure your first Organization, you can use it as a template for creating additional Organizations. Make sure you have completed the configurations for the Organization before copying it.

Select the Organization from the list. The following settings will be copied from the selected Organization:

* Source control integrations (GitHub, GitLab, BitBucket)
* Container registry integrations (ACR, Docker Hub, ECR, GCR)
* Container orchestrator integrations (Kubernetes)
* PaaS and Serverless integrations (Heroku, AWS Lambda)
* Notification integrations (Slack, Jira)
* Policies
* Ignore settings
* Language settings
* Infrastructure as Code settings
* Snyk Code settings

The new Organization will not use the same settings from the copied Organization for the following:

* Service accounts
* Members
* Projects
* Notification preferences

Any of the Organization settings that you configured for the first Organization can then be customized for the new Organization.

#### Enable Snyk Code

When you create a new Organization, Snyk Code (SAST) scanning is disabled by default. We recommend enabling Snyk Code before you import your first projects to Snyk. If Snyk Code is enabled after a project is imported, it won't detect Snyk Code files.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to enable Snyk Code, then click **Save changes**

#### Define Project structure

You may want to house different Projects in different Organizations.

{% hint style="warning" %}
You cannot move Projects between Organizations directly.
{% endhint %}

To re-house Projects, you must re-import your Projects into your new Organizations.

If you are working with a large number of Projects, you should use the API to re-import. See **importing Projects at scale** below for more detail.

### Set up Single Sign-On (SSO)

Users can authenticate into their Snyk accounts in several ways, such as with a GitHub or Google account. However, now that you have Enterprise access, you may want to set up Single Sign-On (SSO) via your existing identity provider to streamline sign-ins and new user provisioning.

{% hint style="info" %}
If using SSO, after you set it up, you'll need to remove any duplicate users from social login accounts (users who have logged in with other methods than SSO).
{% endhint %}

See [Setting up Single Sign-On (SSO) for authentication](../snyk-admin/set-up-single-sign-on-sso-for-authentication/) for more details on steps for using your identity provider to authenticate and provision Snyk to your teams.

See the [SSO, authentication and user provisions](https://training.snyk.io/courses/sso) training course for more details.

### Using Service accounts

[Service accounts](../snyk-admin/service-accounts.md) allow you to better scale and automate.

* If you use CI/CD (using a [CI/CD integration](../integrations/ci-cd-integrations/) or the [Snyk CLI](../snyk-cli/)), we recommend you use a service account.
* If you use an IDE plugin or the CLI to test in your local environment, we recommend you use personal access tokens.

### Rich API

Access to the [Snyk API](../snyk-api-info/) allows you to scale and automate different processes, including importing projects. You will want to use this strategy if you have a large number of projects that need to be moved to your new Snyk Organizations.

**Importing Projects at scale**

{% hint style="info" %}
If you plan to run SAST scans with Snyk Code, check that Snyk Code is enabled within the related Snyk org settings before importing.
{% endhint %}

Use the [Snyk API import tool](../snyk-api-info/other-tools/tool-snyk-api-import/) to import projects into Snyk with a controlled pace using available Snyk APIs.

### Reports

Snyk reports offer data and analytics across all of your projects, displaying historical and aggregated data about projects, issues, dependencies, and licenses.

See [Getting started with Snyk Reports](../manage-issues/snyk-reports/reporting-beta-2022/getting-started-with-snyk-reports.md).

### Security policy management

Security policies allow you to customize the prioritization of specific issues from the default and create rules. This is particularly helpful for changing the severities of issues that are not relevant to a specific project or environment.

See [Getting started with security policies](broken-reference).

### Remote code repositories

If you have onprem repositories, you can scan them with Snyk using [Snyk Broker](../snyk-admin/snyk-broker/).

### GitHub Enterprise integration

If you use GitHub as your SCM, with the Enterprise plan, you now have access to Github Enterprise integration. The GitHub Enterprise integration allows the use of a single personal access token across a Snyk Organization rather than a personal access token tied to an individual user account. Because of this, we recommend making the switch.

{% hint style="info" %}
You do not need a GitHub Enterprise license or subscription to use the GitHub Enterprise integration within Snyk.
{% endhint %}

See [Using Github or Github Enterprise integration](../integrations/git-repository-scm-integrations/using-github-or-github-enterprise-integration.md) for steps on migrating from Github to Github Enterprise integration.
