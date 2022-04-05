# Getting started with Snyk Code

This section describes how to get started with Snyk Code using the Snyk Web UI Environment.

You can work with Snyk Code using the following Snyk Environments:

* &#x20;Snyk Web UI (including the PR Checks Git Environment)
* IDE
* CLI
* API

The prerequisites for using Snyk Code are the same for all Snyk Environments:

* If you are using the Snyk Web UI - follow the step-by-step instructions in this section to activate and use Snyk Code.
* If you are using the Snyk IDE, CLI, or API Environments – use the Snyk Web UI to [enable the **Snyk Code** option](activating-snyk-code-using-the-web-ui/step-1-enabling-the-snyk-code-option.md). Then, continue to test your code with Snyk Code in your selected Environment.

**What is required for getting started with Snyk Code?**

To get started with Snyk Code, all you need is a Snyk Account, a supported SCM, and source code that is written in one of the supported platforms and languages. &#x20;

**Note**: For more information, see [Prerequisites for Snyk Code](prerequisites-for-snyk-code.md).

**What is the activation process of Snyk Code?**

First, you need to enable the **Snyk Code** option in your Snyk Org Settings via the Web UI. If you intend on using the IDE, CLI, or API Environments, this is the only step you need to do in order to start working with Snyk Code.

**Note**: If you are working with the API Environment WITHOUT the CLI, you also need to [integrate your SCM via the Web UI](activating-snyk-code-using-the-web-ui/step-2-integrating-your-source-control-system-with-snyk-code.md). &#x20;

When working with the Snyk Web UI, the next step in the activation process is integrating your Snyk Account with the SCM that contains the repositories you want to test. Then, you can import to your Snyk Account the required repositories, and Snyk Code will automatically analyze them and present to you its analysis results.

**Note**: For more information, see [Activating Snyk Code using the Web UI](activating-snyk-code-using-the-web-ui/).

**Where are the Snyk Code Results?**

During the import process of your selected repositories, Snyk Code automatically analyzes your imported code in search of potential vulnerabilities. All the vulnerability findings that Snyk Code detects in the code of one imported repository, are aggregated in one Snyk Project, called **Code analysis**:

![](<../../../.gitbook/assets/Snyk Code - Getting Started - Projects - Code analysis.png>)

**Note**: Unlike other Snyk products, which create a separate Snyk Project for each imported file, Snyk Code creates one Snyk Project for all the imported files of one repository. This way, all the vulnerabilities that were detected in the repository code are aggregated in one Project, and the Snyk Code results can present the data flow of a vulnerability across multiple files.  &#x20;

To view all the security vulnerabilities that Snyk Code detected in your imported code, click the **Code analysis** Project, and explore the details of each vulnerability:

![](<../../../.gitbook/assets/Snyk Code - Getting Started - Code Analysis Project.png>)

If your Snyk Account is already integrated with your SCM, and it already contains imported repositories, Snyk Code may be already active and running in your Org Settings. In this case, you can check if your existing repositories are already being tested by Snyk Code, by searching for the **Code analysis** Project in your Target folders. If a **Code analysis** Project exists in your imported repositories, you can skip this _Getting Started_ section and move to the [Exploring and working with the Snyk Code results](../exploring-and-working-with-the-snyk-code-results.md) section. However, you may want to learn more about the following topics:

* [Importing additional repositories to Snyk](activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/importing-additional-repositories-to-snyk.md).
* [Excluding directories and files from the import process](activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/excluding-directories-and-files-from-the-import-process.md).
* [Removing imported repositories from the Snyk Code test](activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/removing-imported-repositories-from-the-snyk-code-test.md).







Get started with [Snyk Code](https://snyk.io/product/snyk-code/) to find, prioritize and fix potential vulnerabilities in your proprietary code.

{% hint style="info" %}
This page describes getting started with Snyk Code using the Web UI (see [using-snyk-code-web.md](../using-snyk-code-web.md "mention")). You can also use Snyk Code in your IDE (see [using-snyk-code-via-ide.md](../using-snyk-code-via-ide.md "mention")), or in the CLI (see [cli-for-snyk-code](../cli-for-snyk-code/ "mention")).
{% endhint %}

## Prerequisites

* A Snyk account.
* Projects with code in a supported language (see [snyk-code-language-and-framework-support.md](../snyk-code-language-and-framework-support.md "mention")).
* A supported Source Code Management (SCM) system: GitHub, GitHub Enterprise, Bitbucket Cloud, Bitbucket Server, GitLab Cloud, Azure Repos.

## Stage 1: Enable Snyk Code

Snyk Code is disabled by default, so you must enable it for each organization:

1. Log in to [Snyk.io](http://snyk.io).
2. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **Snyk Code**.
3. Under **Enable Snyk Code**, change **Disabled** to **Enabled.**
4. Click **Save changes**.

## Stage 2: Add source control integration

{% hint style="info" %}
If you already have an integration set up, you can skip this step
{% endhint %}

Choose a source code integration, to allow Snyk to work on a project:

1. Log in to [Snyk.io](http://snyk.io).
2. Select **Integrations > Source control**.
3. Click the source control system (for example, GitHub) to integrate with Snyk:
4. Fill in the account credentials as prompted (or authenticate with your account in GitHub), to grant Snyk access permissions for integration.

Snyk Code temporarily clones your repository for code analysis, this requires appropriate permissions and HTTPS access to your SCM.

For more information about how data is stored, see [How Snyk handles your data](../../../more-info/how-snyk-handles-your-data.md). For more details about integrations, see [DevOps integrations & languages](https://docs.snyk.io/introducing-snyk/introduction-to-snyk/integrations-and-languages).

## Stage 3: Add projects

{% hint style="info" %}
If you already have projects added, you can skip this step.
{% endhint %}

Add projects to test with Snyk, by choosing repositories for Snyk to test and monitor:

1. Select **Projects** from [snyk.io](http://snyk.io).
2. Select the tool to add the project from (for example GitHub):
3. In **Personal and Organization repositories**, select the repositories to use:
4. Click **Add selected repositories** to import the selected repositories into your projects. This sets Snyk to run a regular check (daily by default) for your proprietary code vulnerabilities.
5. A progress bar appears: click **View log** to see log results.
6. Project import completes.

{% hint style="info" %}
Currently Snyk Code does not support the **Exclude folders** option during import. Please contact us if you need more information.
{% endhint %}

See [Snyk projects](https://support.snyk.io/hc/en-us/sections/360004724958-Snyk-projects) for more details.

##
