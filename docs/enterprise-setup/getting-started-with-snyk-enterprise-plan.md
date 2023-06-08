# Getting started with Snyk: Enterprise plan

## Introduction

In this guide, we’ll look at how you can get several team members involved in trying Snyk to get feedback and buy-in.

{% hint style="info" %}
If you have been using Snyk on the Free/Team Plan and are looking for guidance on upgrading to the Enterprise plan, read the [Upgrading to Enterprise Plan](upgrading-to-enterprise-plan.md) guide.
{% endhint %}

Snyk has a number of tools and processes that help secure your entire software development lifecycle. With Snyk, you can scan while you’re coding. You can also monitor code when you’re not working on it. Snyk also provides visibility into issues across your projects with a git repository integration. And Snyk can integrate into CI/CD through integrations, the CLI, or curated containers.

For users evaluating Snyk or planning an enterprise deployment, for most programming languages, we recommend integrating with a Git repository to get started.

{% hint style="info" %}
The tool that best serves your tech stack, environment, and workflow will depend on your individual circumstances. Visit the guide specific to your language for more details.
{% endhint %}

To learn more about choosing the integration points within the software development lifecycle that work best for you and your teams’ current level of security maturity, see the [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) training course.

### Try out a project

This guide explains how to connect a Git repository to Snyk, and show results from scanning Snyk Projects in that repository. The focus is to help you understand how Snyk works, before you start the process of implementing Snyk with your team(s).

{% hint style="warning" %}
Snyk provides limited free tests per month for each type of scan (open source, code, container, IaC). For unlimited tests, ensure you have a paid plan for that type of testing and that the setting for that type has been enabled.
{% endhint %}

#### Create or log into account

You need a Snyk account to use Snyk functionality, even within your local environment. [Create a free account](../getting-started/quickstart/create-a-snyk-account/) to try out a project. If your organization is already using Snyk, you may be able to log in via single sign-on to be provisioned with a Snyk account (see [Logging in to an existing account](../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md)).

**Enable Snyk Code**

When you create a new Organization, Snyk Code (SAST) scanning is disabled by default. We recommend enabling this product before you import your first projects to Snyk.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to enable Snyk Code, then click **Save changes**.

**Add Project**

Add a Snyk Project to connect a Git repository integration. This video demonstrates the process.

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
Adding projects video
{% endembed %}

Read the [Preparing for implementation: Enterprise plan](preparing-for-implementation-enterprise-plan.md) guide for more information on configuring the Git repository integration for your Organization.

**Configure initial integration settings**

After your Git repository is connected (see [See Git repository integrations (SCMs)](../integrations/git-repository-scm-integrations/) for details), there are automated processes available to automatically check pulls requests for vulnerabilities, automatically generate pull requests, and automatically generate dependency upgrade pull requests. We recommend you disable these options initially.

The settings for each Snyk Project are inherited from the Organization integration settings. Use the following steps to ensure the following settings are disabled: Default Snyk test for pull requests, Automatic fix pull requests, Automatic dependency upgrade pull requests, and Automatic updates to Dockerfile base images. You can revisit these settings when your teams are ready to implement these options.

1. Select the **Integrations** page on the left-hand navigation menu.
2. Select the settings ‘cog’ icon for your Git repository integration.
3. In the **Default Snyk test for pull requests** section, ensure the following are disabled:
   1. **Open Source Security & Licenses** (default checks when PRs are opened)
   2. **Automatic fix pull requests**: both **New vulnerabilities** and **Known vulnerabilities (backlog)**
   3. **Automatically update Dockerfile base images**
   4. **Automatic dependency upgrade pull requests**

{% hint style="info" %}
We recommend defining standards for these options as well as notification defaults before you add more than a few Projects. When your teams are ready for a broader implementation, we recommend defining standards for these options according to your security maturity. Review the Preparing for implementation guide for more information
{% endhint %}

**Review scan result**

For Open Source projects, Snyk provides visibility into dependency and license component issues. Some package managers also provide a link to open a pull request to fix a specific issue, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/z0lcfht0na" %}
Fixing issues video
{% endembed %}

Snyk also provides information about potential security and quality issues in your proprietary code, and provides recommendations and some remediation examples, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/p32o09thet" %}
Review code issues video
{% endembed %}

If Snyk identified a Dockerfile in the repository, it provides information about the base image, including detected vulnerabilities. Snyk also provides the option to open a pull request to replace the base image with one that is more secure, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/3s7h3r9o8h" %}
Replacing base image video
{% endembed %}

If Snyk identified a Kubernetes or Terraform file in the repository, it provides information about the configuration. The following video demonstrates the information provided.

{% embed url="https://thoughtindustries-1.wistia.com/medias/l29m2dwrk8" %}
Reviewing configuration issues video
{% endembed %}

#### Scan with Snyk CLI

Some package managers rely on context from the local environment, so scanning in the local environment or as part of the CI/CD pipeline gives the most accurate results.

You need to [install the Snyk CLI](../snyk-cli/install-the-snyk-cli/). Once installed, you need to authenticate it to your Snyk account, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
Authenticate CLI video
{% endembed %}

A scan with [**Snyk test**](../scan-application-code/snyk-open-source/use-snyk-open-source-from-the-cli/) surfaces information about open source package issues, including fix advice, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test video
{% endembed %}

A scan with [**Snyk code test** ](../scan-application-code/snyk-code/cli-for-snyk-code/)runs a Static Code Analysis test on the code in that Project, and returns the list of detected vulnerability issues, general information about the test, and a summary of the test findings.

A scan with [**Snyk container test**](../scan-containers/snyk-cli-for-container-security/) returns a list of vulnerabilities in the container image, along with recommendations for upgrading the base image for one that is more secure.

A scan with [**Snyk iac test**](../scan-cloud-configurations/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/) returns advice on how to resolve discovered issues in your Infrastructure as Code files.

### What’s next?

* If you want developers to try Snyk in their local environment using the IDE or CLI, review the [Getting started guide for individuals and small teams](../getting-started/getting-started-with-snyk-free-team-plan.md).
* To get specific recommendations for your tech stack, visit the guide specific to your language.
* When you are ready to plan a Snyk rollout to more teams, review the [Preparing for implementation guide: Enterprise plan](preparing-for-implementation-enterprise-plan.md).
* See the [Launch Snyk to your teams](https://training.snyk.io/courses/launch-snyk-to-your-teams) training course for additional strategies, communication templates, and checklists for rolling Snyk out to a wider audience.
