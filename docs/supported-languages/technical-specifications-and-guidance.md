# Technical specifications and guidance

## Unicode character encoding

Both Snyk Code and Snyk Open Source accept source code files in UTF-8 encoding. Consider converting source files to this encoding type before importing them into Snyk.

## Snyk Open Source

Snyk analyzes and builds the dependencies tree depending on the language and package manager for the Project, as well as the location of the Project.

### How Snyk for Open Source and licensing works

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must build your Project. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk builds a dependency graph and (dependencies tree) and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

### Snyk policies in Open Source

For information on managing dependencies and vulnerabilities from your developer workflows through the use of policies, see:

* [Defining a secure open-source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

### Open Source license compliance

To check compliance for open source licenses, see [Snyk License Compliance Management](../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md).

## Snyk Code

### File size limit for Snyk Code analysis

Snyk Code automatically excludes the following files from analysis:

* On the Web UI - files that are larger than 1MB.
* On the CLI and IDE - files that are larger than 1MB.
* Minified JS files with 3 or fewer lines.

### Filename length limitation

The analysis is available only for files with names shorter than or equal to 255 characters. If the filename exceeds this limit, you receive an error. To ensure that all files are being analyzed, Snyk recommends shortening long filenames.

### Framework support&#x20;

To support a specific framework, Snyk Code must both support the relevant language and be trained on Projects using the framework. The found patterns are then annotated by the security team and extended by curated content.

Most frameworks are partially supported out of the box, as Snyk Code needs only to parse the code to analyze it. In some cases, frameworks may require specific rules, or require specific program analysis engine updates, or both.

If you notice any gaps in support for a specific framework, [contact Snyk Support](https://support.snyk.io).

Snyk categorizes framework support into two levels: Comprehensive and Partial.

Comprehensive support includes:

* Sources and sinks: All relevant sources and sinks are identified.
* Data flow testing: Extensive testing provides thorough data flow coverage.
* Engine support: The Snyk Code engine is fully optimized for this framework.
* Limitations: No known limitations; report any false negatives to [Snyk Support](https://support.snyk.io).

Partial support includes:

* Sources and sinks: Coverage may be limited, with some missing elements.
* Data flow testing: Limited testing has been conducted.
* Engine support: Compatibility is restricted, which may impact analysis accuracy.
* Limitations: Potential for false negatives in taint analysis or source/sink identification.

Snyk continuously expands its framework coverage and improves analysis accuracy.

### How Snyk Code analysis works

Snyk scans your codebase following this sequence:

1. The source code is analyzed, generating an event graph. The event graph is similar to a code map that helps Snyk understand how different parts of the code are related. There are two node types, each node in the graph representing something that happens in the code. Some represent parts of the code, and others represent how the code is used.
2. Rules are run against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If a match is found, Snyk looks for a vulnerability in the event graph, identifying where problems might be hiding in the code.

For more information, see [Snyk Code AI Engine](../scan-with-snyk/snyk-code/#ai-engine). For more information about Snyk Code language support, see [Supported languages, package managers, and frameworks](supported-languages-package-managers-and-frameworks.md).

## Language support and CLI, CI/CD, and SCM integrations

Snyk supports a variety of programming languages, enabling seamless integration into your development workflow through CLI commands, CI/CD pipelines, and SCM integrations.&#x20;

You can use these tools to automatically check your code for security issues as you develop your software. This ensures that strong security practices are part of your development process.

Navigate to the following pages for more details:

* CLI for [Snyk Open Source](../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/) and [Snyk Code](../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)
* CI/CD for [Snyk Open Source](../developer-tools/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-open-source-specific-ci-cd-strategies.md) and [Snyk Code](../developer-tools/snyk-ci-cd-integrations/use-snyk-code-in-the-ci-cd-pipeline.md)
* [SCM integrations](../developer-tools/scm-integrations/organization-level-integrations/) for Snyk Open Source and Snyk Code
