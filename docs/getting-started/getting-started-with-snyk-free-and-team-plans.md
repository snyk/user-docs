# Getting started with Snyk: Free and Team plans

These pages explain how you can try a few scans to see the results.

Snyk has a number of tools and processes that help secure your entire software development lifecycle. With Snyk, you can validate your code while you are coding. You can also monitor code when you are not working on it. Snyk also provides visibility into issues across your Projects with a Git repository integration and can work with your CI/CD through integrations, the CLI, or curated containers. It is common practice to integrate Snyk into several points of your development process for enabling your developers, for visibility, and for gating your applications.

If this is your first time performing a scan, or you are interested in the results for a single application while you are working on it, scanning in your local environment is a great place to start, and that is covered in this guide.&#x20;

If you have a set of applications you are responsible for, as an individual or a team, Snyk recommends configuring the Git repository integration to start getting visibility for the issues in your repositories in a few clicks.

{% hint style="info" %}
The tool or tools that best serve your tech stack, environment, and workflow depend on your individual circumstances. See the language pages for more information.
{% endhint %}

To learn more about choosing the integration points in the software development lifecycle that work best for you and your current level of security maturity, see the [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) course in Snyk Training.

## Try out a Project

This guide explains how to test a sample or single Project in your local development environment or by using the Snyk CLI.

{% hint style="info" %}
The Snyk free plan provides limited free tests per month. For unlimited tests, consider a paid plan.
{% endhint %}

### Create or log into your account

You need a Snyk account to use Snyk functionality, even within your local environment. [Create a free account](quickstart/create-or-log-in-to-a-snyk-account.md) to try out a Project. If your Organization is already using Snyk, you may be able to log in using single sign-on to be provisioned with a Snyk account. For more information, see [Create or log in to a Snyk account](quickstart/create-or-log-in-to-a-snyk-account.md).

### Test a Project in your local development environment

To scan a single Project in your local development environment, you must use a Snyk plugin or extension with your IDE. You must also authenticate the plugin or extension with your Snyk account, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/bva0d7k46b" %}
Install IDE and authenticate to Snyk
{% endembed %}

{% hint style="warning" %}
When authenticating the IDE, you may see a warning about scanning folders you trust. Because Snyk is executing code, such as invoking the package manager to get dependency information, you must trust the folder you are scanning to continue.
{% endhint %}

A scan with the Snyk IDE plugin or extension in a local Project provides information about open-source package issues, including fix advice.

{% embed url="https://thoughtindustries-1.wistia.com/medias/bny3j0ywui" %}
Review open source dependency issues video
{% endembed %}

Scanning with the Snyk IDE plugin or extension in a local Project also provides information about code issues, including example fixes.

{% embed url="https://thoughtindustries-1.wistia.com/medias/fazwzk6oi3" %}
Review code issues video
{% endembed %}

### Test a Project with the Snyk CLI

Some package managers rely on context from the local environment, so testing in the local environment or as part of the CI/CD pipeline provides the most accurate results.

First, [install the Snyk CLI](../snyk-cli/install-or-update-the-snyk-cli/). After installation, you must authenticate it to your Snyk account, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
Authenticate CLI video
{% endembed %}

A scan with [**Snyk test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/use-snyk-open-source-from-the-cli/) provides information about open-source package issues, including fix advice, as demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test video
{% endembed %}

A scan with [**Snyk code test** ](../snyk-cli/scan-and-maintain-projects-using-the-cli/using-snyk-code-from-the-cli/)runs a Static Code Analysis test on the code in that Project, and returns the list of detected vulnerability issues, general information about the test, and a summary of the test findings.

A scan with [**Snyk container test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/use-snyk-container-from-the-cli/) returns a list of vulnerabilities in the container image, along with recommendations for upgrading the base image to one that is more secure.

A scan with [**Snyk iac test**](../scan-infrastructure/snyk-cli-for-iac/) returns advice on how to resolve discovered issues in your infrastructure as code files.

## What’s next?

* When you are ready to start scanning more applications, see [Preparing for Implementation:  Free and Team plans](preparing-for-implementation-free-and-team-plans.md).
* To get specific recommendations for your tech stack, see the pages specific to your language.
* If you decide you want to expand the use of Snyk throughout your business and involve more teams in Snyk, read the [Enterprise implementation guide](../enterprise-setup/enterprise-implementation-guide/).
