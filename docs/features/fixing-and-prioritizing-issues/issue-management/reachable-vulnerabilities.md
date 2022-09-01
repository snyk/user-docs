# Reachable vulnerabilities

## Introduction

The first step of scanning apps for open source vulnerabilities is identifying the apps' packages.

The next step is to see which of the identified packages contains a vulnerability and whether or not the vulnerability affects the apps’ code. This process can easily lead to hundreds of vulnerabilities at the organizational level.

There are multiple ways to [prioritize vulnerabilities](https://docs.snyk.io/features/fixing-and-prioritizing-issues/prioritizing-and-managing-issues/evaluating-and-prioritizing-vulnerabilities), and one of the ways is to look for reachable vulnerabilities. A **reachable** vulnerability has a path from the code to the vulnerable function.

By looking deeper into how the app uses the open-source dependencies and how the open-source dependencies interact with each other, we can add the needed context around the found vulnerabilities. Looking at the code reachability helps you decide how to prioritize which vulnerabilities to fix first.

### How it works

To provide as accurate results as possible, we use multiple algorithms to build a call graph from your app to the open-source dependencies you use. After that call graph, we can identify which vulnerabilities have a path from the app’s code down to the vulnerable function or package.

We split the results into the following areas:

1. **Reachable** - A clear path from the code to vulnerable functions was found.&#x20;
2. **No path found** - Direct calls from the code to vulnerable functions could not be found.

### Supported languages and prerequisites

Reachable Vulnerabilities analysis is available for Java (Maven and Gradle) in the [supported versions](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven) using GitHub integration. Kotlin is currently not available.

### Scanning for Reachable Vulnerabilities using Git Integrations

{% hint style="info" %}
To provide this feature, Snyk takes a temporary copy of your Git repository contents.

For more information see [how-snyk-handles-your-data.md](../../../more-info/how-snyk-handles-your-data.md "mention")
{% endhint %}

1. Set up your [GitHub integration](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration).
2. Opt-in to Reachable Vulnerabilities analysis
   * In **Organization** settings, go to the **Languages** section.
   * Go to the **Reachable vulnerabilities analysis** section.
   * Opt-in to the **Reachable vulnerabilities analysis** and save your changes.
3. Go to the **import projects** page and choose the repositories to import to Snyk.
4. Selected projects are imported and analyzed for Reachable Vulnerabilities issues.

### Project page

After running `snyk monitor` in the CLI, or importing a project via Snyk UI, the project is monitored by Snyk, and the results of the Reachable Vulnerabilities analysis appear in the Project page in the following places:

1. **Filters** - Allows you to focus on reachable vulnerabilities first by filtering results based on reachability.
2. **Reachability badge** - Allows you to quickly understand the reachability level of vulnerabilities.
3. **Call path** - Allows you to see the path from your code to the vulnerable function to validate the result.
4. **Priority Score** - We will increase the Priority Score for reachable vulnerabilities, to allow you to see these issues first.

![Reachability Call Path](<../../../.gitbook/assets/image (222).png>)

![Priorty Score](<../../../.gitbook/assets/image (126).png>)

### Reports

You can filter by reachability status to quickly show the reachable issues.

![](<../../../.gitbook/assets/image (153).png>)

### Scanning for Reachable Vulnerabilities using Snyk CLI

1. Using the Snyk CLI workflow for reachable vulnerabilities scanning is currently not available.
2. We are developing new and improved capabilities for prioritizing vulnerabilities to upgrade this functionality.&#x20;
