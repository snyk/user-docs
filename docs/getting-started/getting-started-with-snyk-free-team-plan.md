# Getting started with Snyk: Free / Team plan

## Introduction

In this guide, we’ll look at how you can try a few scans to see the results.

Snyk has a number of tools and processes that help secure your entire software development lifecycle. With Snyk, you can validate your code while you are coding. You can also monitor code when you’re not working on it. Snyk also provides visibility into issues across your projects with a git repository integration. And Snyk can integrate into CI/CD through integrations, the CLI, or curated containers. It's quite common to have Snyk integrated into several points of your development process for enabling your developers, for visibility and for gating of your applications.

If this is your first time performing testing, or you are interested in the results for a single application while you're working on it, scanning in your local environment is a great place to start, and that will be covered in this guide.&#x20;

If you have a set of applications you're responsible for, as an individual or a team, we recommend configuring the Git repository integration to start getting visibility of the issues on your repositories in a few clicks, this will be covered in the following Implementation and Getting Started guides for each plan.

{% hint style="info" %}
The tool(s) that best serves your tech stack, environment, and workflow will depend on your individual circumstances. See the tech stack implementation guides for more details.
{% endhint %}

To learn more about choosing the integration points within the software development lifecycle that work best for you and your current level of security maturity, see [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) course in Snyk Training.

### Try out a project

This guide explains how to test a sample or single project in your local development environment or by using the Snyk CLI.

{% hint style="info" %}
Snyk free plan provides limited free tests per month. For unlimited tests, consider a paid plan.
{% endhint %}

#### Create or log into your account

You need a Snyk account to use Snyk functionality, even within your local environment. [Create a free account](quickstart/create-a-snyk-account/) to try out a project. If your organization is already using Snyk, you may be able to log in via single sign-on to be provisioned with a Snyk account (see [Logging in to an existing account](quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md)).

#### Test a project in your local development environment

To scan a single project in your local development environment, you need to use a Snyk plugin or extension with your IDE. You also need to authenticate the plugin or extension with your Snyk account, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/bva0d7k46b" %}
Install IDE and authenticate to Snyk
{% endembed %}

{% hint style="warning" %}
When authenticating the IDE, you may see a warning about scanning folders you trust. Because Snyk is executing code, such as invoking the package manager to get dependency information, you’ll need to trust the folder you’re scanning to continue.
{% endhint %}

A scan with the Snyk IDE plugin or extension in a local project surfaces information about open source package issues, including fix advice.

{% embed url="https://thoughtindustries-1.wistia.com/medias/bny3j0ywui" %}
Review open source dependency issues video
{% endembed %}

Scanning with the Snyk IDE plugin or extension in a local project also surfaces information about code issues, including example fixes.

{% embed url="https://thoughtindustries-1.wistia.com/medias/fazwzk6oi3" %}
Review code issues video
{% endembed %}

#### Test a project with the Snyk CLI

Some package managers rely on context from the local environment, so testing in the local environment or as part of the CI/CD pipeline provides the most accurate results.

You need to [install the Snyk CLI](../snyk-cli/install-or-update-the-snyk-cli/). Once installed, you need to authenticate it to your Snyk account, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
Authenticate CLI video
{% endembed %}

A scan with [**Snyk test**](../scan-application-code/snyk-open-source/use-snyk-open-source-from-the-cli/) surfaces information about open source package issues, including fix advice, demonstrated in this video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test video
{% endembed %}

A scan with [**Snyk code test** ](../scan-application-code/snyk-code/using-snyk-code-from-the-cli/)runs a Static Code Analysis test on the code in that project, and returns the list of detected vulnerability issues, general information about the test, and a summary of the test findings.

A scan with [**Snyk container test**](../scan-applications/snyk-container/use-snyk-container-from-the-cli/) returns a list of vulnerabilities in the container image, along with recommendations for upgrading the base image for one that is more secure.

A scan with [**Snyk iac test**](../scan-infrastructure/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/) returns advice on how to resolve discovered issues in your Infrastructure as Code files.

### What’s next?

* When you are ready to start scanning more applications, read the [Preparing for implementation guide: Free / Team plan](preparing-for-implementation-free-team-plan.md).
* To get specific recommendations for your tech stack, read the guide specific to your language.
* If you decide you want to expand Snyk usage throughout your business, and involve more teams in Snyk, read the [Getting started with Snyk: Enterprise plan](../enterprise-setup/getting-started-with-the-snyk-enterprise-plan.md).
