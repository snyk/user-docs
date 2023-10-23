# Getting started with the Snyk Enterprise plan

This page looks at how you can get several members of your team involved in trying Snyk to get feedback and buy-in.

{% hint style="info" %}
If you have been using Snyk on the Free or Team Plan and are looking for guidance on upgrading to the Enterprise plan, see [Upgrading to the Enterprise Plan](upgrading-to-the-enterprise-plan.md).
{% endhint %}

Snyk has a number of tools and processes that help secure your entire software development lifecycle. With Snyk, you can scan while you are coding or monitor code when you are not working on it. Snyk can also provide visibility into issues across your Projects with a Git repository integration or integrate into your CI/CD through plugins, the CLI, or curated containers.

For users who are evaluating Snyk or planning an enterprise deployment, and for most programming languages, Snyk recommends integrating with a Git repository to get started.

{% hint style="info" %}
The tool that best serves your tech stack, environment, and workflow depends on your individual circumstances. Visit the [guide](broken-reference) specific to your language for more details.
{% endhint %}

To learn more about choosing the best integration points within the software development lifecycle for you and for your team, at their current level of security maturity, see the [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) course.

To see what Snyk can do for you, **try out a Project**

The rest of this page explains how to connect a Git repository to Snyk and how to show results from scanning Snyk Projects that are in that repository. The focus is on helping you understand how Snyk works before you start the process of implementing Snyk with your team or teams.

{% hint style="warning" %}
Snyk provides limited free tests per month for each type of scan, Snyk Open Source, Code, Container, or IaC. For unlimited tests, ensure you have a paid plan for the type of testing you need to do, and that the setting for that type has been enabled.
{% endhint %}

## Create or log into your Snyk account

You need a Snyk account to use Snyk functionality, even within your local environment. [Create a free account](../getting-started/quickstart/create-a-snyk-account.md) to try out a Project. If your enterprise is already using Snyk, you may be able to log in using single sign-on to be provisioned with a Snyk account. For details, see [Logging in to an existing account](broken-reference).

## **Enable Snyk Code**

When you create a new Organization in Snyk, Snyk Code (SAST) scanning is disabled by default. Snyk recommends enabling the Snyk Code product before you import your first Projects into Snyk.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to enable Snyk Code; then click **Save changes**.

## **Add a Snyk Project**

Add a Snyk Project to connect a Git repository integration. This video demonstrates the process.

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
Adding projects video
{% endembed %}

See [Preparing for implementation of the Enterprise plan](preparing-for-implementation-of-the-enterprise-plan.md) for more information on configuring the Git repository integration for your Organization.

## **Configure initial Snyk integration settings**

After your Git repository is connected (see [See Git repository integrations (SCMs)](../integrations/git-repository-scm-integrations/) for details), you have automated processes available to automatically check pull requests for vulnerabilities, automatically generate pull requests, and automatically generate dependency upgrade pull requests. Snyk recommends you disable these options initially.

The settings for each Snyk Project are inherited from the Snyk Organization integration settings. Follow these steps to ensure these settings are disabled: Default Snyk test for pull requests, Automatic fix pull requests, Automatic dependency upgrade pull requests, and Automatic updates to Dockerfile base images. You can go back and enable these settings when your teams are ready to implement these options.

1. Select the **Integrations** page from the left-hand navigation menu.
2. Select the **settings cog icon** for your Git repository integration.
3. In the **Default Snyk test for pull requests** section, ensure the following are disabled:
   1. **Open Source Security & Licenses** (default checks when PRs are opened)
   2. **Automatic fix pull requests**: both **New vulnerabilities** and **Known vulnerabilities (backlog)**
   3. **Automatically update Dockerfile base images**
   4. **Automatic dependency upgrade pull requests**

{% hint style="info" %}
Snyk recommends defining standards for these options as well as notification defaults before you add more than a few Projects. When your teams are ready for a broader implementation, Snyk recommends defining standards for these options according to your security maturity. Review [Preparing for implementation of the Enterprise plan](preparing-for-implementation-of-the-enterprise-plan.md) for more information
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

Some package managers **rely on context from the local environment**. With these package managers, **scanning in the local environment or as part of the CI/CD pipeline gives the most accurate results**.

To start using the Snyk CLI or a CI/CD plugin, [install the Snyk CLI](../snyk-cli/install-or-update-the-snyk-cli/). After you have installed it, you must [authenticate your machine to associate the CLI with your Snyk account](../snyk-cli/authenticate-the-cli-with-your-account.md), as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
Authenticate CLI video
{% endembed %}

A scan with [**Snyk test**](../scan-applications/snyk-open-source/use-snyk-open-source-from-the-cli/) provides information about open-source package issues, including fix advice, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test video
{% endembed %}

A scan with [**Snyk code test** ](../scan-applications/snyk-code/using-snyk-code-from-the-cli/)runs a Static Code Analysis test on the code in that Project, and returns the list of detected vulnerability issues, general information about the test, and a summary of the test findings.

A scan with [**Snyk container test**](../scan-applications/snyk-container/use-snyk-container-from-the-cli/) returns a list of vulnerabilities in the container image, along with recommendations for upgrading the base image to one that is more secure.

A scan with [**Snyk iac test**](../scan-infrastructure/snyk-cli-for-iac/) returns advice on how to resolve discovered issues in your Infrastructure as Code files.

## Next steps in implementing the Snyk Enterprise plan

* If you want developers to try Snyk in their local environment using the IDE or CLI, review the [Getting started guide for individuals and small teams](../getting-started/getting-started-with-snyk-free-team-plan.md).
* To get specific recommendations for your tech stack, visit the guide specific to your language.
* When you are ready to plan a Snyk rollout to more teams, review [Preparing for implementation of the Enterprise plan](preparing-for-implementation-of-the-enterprise-plan.md) for more information.
* See the [Launch Snyk to your teams](https://training.snyk.io/courses/launch-snyk-to-your-teams) training course for additional strategies, communication templates, and checklists for rolling Snyk out to a wider audience.
