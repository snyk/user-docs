# Enterprise setup

Enterprise configuration focuses on helping understand, plan, and roll out your Enterprise Snyk solution. This includes:

* [Auto-provisioning guide](auto-provisioning-guide.md)
* [Authenticate for third-party tools](authentication-for-third-party-tools.md)
* [Set up SSO](single-sign-on-sso-for-authentication-to-snyk/)
* [Snyk Broker](snyk-broker/)
* [Service accounts](service-accounts/)

These pages provide instructions on how to authenticate and connect while using the Snyk Enterprise plan.

The [Enterprise implementation guide](../enterprise-implementation-guide/) focuses on open-source and code analysis. Updates for container and infrastructure as code are planned.

This page looks at how you can get several members of your team involved in trying Snyk to get feedback and buy-in.

{% hint style="info" %}
If you have been using Snyk on the Free or Team Plan and are looking for guidance on upgrading to the Enterprise plan, see the [Enterprise implementation guide](../enterprise-implementation-guide/).
{% endhint %}

Snyk has a number of tools and processes that help secure your entire software development lifecycle. With Snyk, you can scan while you are coding or monitor code when you are not working on it. Snyk can also provide visibility into issues across your Projects with a Git repository integration or integrate into your CI/CD through plugins, the CLI, or curated containers.

For users who are evaluating Snyk or planning an enterprise deployment, and for most programming languages, Snyk recommends integrating with a Git repository to get started.

{% hint style="info" %}
The tool that best serves your tech stack, environment, and workflow depends on your individual circumstances. See [Supported languages and frameworks ](../../supported-languages/supported-languages-package-managers-and-frameworks.md)for the guide specific to your language for more details.
{% endhint %}

To learn more about choosing the best integration points within the software development lifecycle for you and for your team, at their current level of security maturity, see [Integrating Snyk at your company](https://learn.snyk.io/lesson/integrate-snyk-at-your-company/).

The rest of this page explains how to connect a Git repository to Snyk and how to show results from scanning Snyk Projects that are in that repository. The focus is on helping you understand how Snyk works before you start the process of implementing Snyk with your team or teams.

{% hint style="info" %}
Snyk provides limited free tests per month for each type of scan, Snyk Open Source, Code, Container, or IaC. For unlimited tests, ensure you have a paid plan for the type of testing you need to do, and that the setting for that type has been enabled.
{% endhint %}

## Create or log in to your Snyk account

You must have a Snyk account to use Snyk functionality, even within your local environment. Create a free account to try out Snyk for a Project. If your enterprise is already using Snyk, you may be able to log in using single sign-on to be provisioned with a Snyk account.

## **Enable Snyk Code**

When you create a new Organization in Snyk, Snyk Code (SAST) scanning is disabled by default. Snyk recommends enabling the Snyk Code product before you import your first Projects into Snyk.

1. Select the **Settings** > **Snyk Code** option.
2. Click the toggle to enable Snyk Code; then click **Save changes**.

## **Add a Snyk Project**

Add a Snyk Project to connect a Git repository integration. This video demonstrates the process.

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
Adding projects video
{% endembed %}

## **Configure initial Snyk integration settings**

After your Git repository is connected (see [See Git repository integrations (SCMs)](../../developer-tools/scm-integrations/organization-level-integrations/) for details), you have automated processes available to automatically check pull requests for vulnerabilities, automatically generate pull requests, and automatically generate dependency upgrade pull requests. Snyk recommends you disable these options initially.

The settings for each Snyk Project are inherited from the Snyk Organization integration settings. Follow these steps to ensure these settings are disabled: Default Snyk test for pull requests, Automatic fix pull requests, Automatic dependency upgrade pull requests, and Automatic updates to Dockerfile base images. You can go back and enable these settings when your teams are ready to implement these options.

1. Select the **Integrations** page from the left-hand navigation menu.
2. Select the settings cog icon for your Git repository integration.
3. In the **Default Snyk test for pull requests** section, ensure the following are disabled:
   1. **Open Source Security & Licenses** (default checks when PRs are opened)
   2. **Automatic fix pull requests**: both **New vulnerabilities** and **Known vulnerabilities (backlog)**
   3. **Automatically update Dockerfile base images**
   4. **Automatic dependency upgrade pull requests**

{% hint style="info" %}
Snyk recommends defining standards for these options as well as notification defaults before you add more than a few Projects. When your teams are ready for a broader implementation, Snyk recommends defining standards for these options according to your security maturity. For more information, see [Configure integrations](../enterprise-implementation-guide/phase-2-configure-account/set-visibility-and-configure-an-organization-template/configure-integrations.md).
{% endhint %}

## **Review the Snyk scan result**

For Open Source Projects, Snyk provides visibility into dependency and license component issues. Some package managers also provide a link to open a pull request to fix a specific issue, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/z0lcfht0na" %}
Fixing issues video
{% endembed %}

Snyk also provides information about potential security and quality issues in your proprietary code, and provides recommendations and some remediation examples, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/p32o09thet" %}
Review code issues video
{% endembed %}

If Snyk identifies a Dockerfile in the repository, Snyk provides information about the base image, including detected vulnerabilities. Snyk also provides the option to open a pull request to replace the base image with one that is more secure, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/3s7h3r9o8h" %}
Replacing base image video
{% endembed %}

If Snyk identifies a Kubernetes or Terraform file in the repository, Snyk provides information about the configuration. The following video demonstrates the information that Snyk provides.

{% embed url="https://thoughtindustries-1.wistia.com/medias/l29m2dwrk8" %}
Reviewing infrastructure issues video
{% endembed %}

## Scan with the Snyk CLI

Some package managers rely on context from the local environment. With these package managers, scanning in the local environment or as part of the CI/CD pipeline gives the most accurate results.

To start using the Snyk CLI or a CI/CD plugin, [install the Snyk CLI](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/). After you have installed it, you must [authenticate your machine to associate the CLI with your Snyk account](../../developer-tools/snyk-cli/authenticate-to-use-the-cli.md), as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
Authenticate CLI video
{% endembed %}

A scan with [Snyk test](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/) provides information about open-source package issues, including fix advice, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test video
{% endembed %}

A scan with [`snyk code test` ](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)runs a Static Code Analysis test on the code in that Project, and returns the list of detected vulnerability issues, general information about the test, and a summary of the test findings.

A scan with [`snyk container test`](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) returns a list of vulnerabilities in the container image, along with recommendations for upgrading the base image to one that is more secure.

A scan with [`snyk iac test`](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/) returns advice on how to resolve discovered issues in your Infrastructure as code files.

## Next steps in implementing the Snyk Enterprise plan

* If you want developers to try Snyk in their local environment using the IDE or CLI, review [Snyk API](../../snyk-api/snyk-api.md) and [Snyk CLI](../../developer-tools/snyk-cli/).
* To get specific recommendations for your tech stack, visit the guide specific to your language.
* When you are ready to plan a Snyk rollout to more teams, review the [Enterprise implementation guide](../enterprise-implementation-guide/) for more information.
* See the [Developer launch package](https://assets.ctfassets.net/4un77bcsnjzw/2YfaqJNMsogGNJM6BBQz4p/8f5ca77b9c40a1bbe14cc9fb0aa05462/Snyk-developer-launch-package.pdf) for additional strategies, communication templates, and checklists for rolling Snyk out to a wider audience.
