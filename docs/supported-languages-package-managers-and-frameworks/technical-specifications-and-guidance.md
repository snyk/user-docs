# Technical specifications and guidance

## Snyk Open Source

### How Snyk for Open Source and licensing works

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must build your Project. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk builds a dependency graph and (dependencies tree) and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

Snyk analyzes and builds the dependencies tree depending on the language and package manager for the Project, as well as the location of the Project.

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of Projects that have a package manager, this means a release of the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the GitHub repo.
{% endhint %}

### Snyk policies in Open Source

For information on managing dependencies and vulnerabilities from your developer workflows through the use of policies, see:

* [Defining a secure open-source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

### Open Source license compliance

To check compliance for open source licenses, see [Snyk License Compliance Management](../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md).

## Snyk Code

### Supported file extensions for Snyk Code

<table><thead><tr><th>Language</th><th width="215">Interfile Support</th><th>Supported Extension</th></tr></thead><tbody><tr><td>Apex</td><td>Yes</td><td>.cls, .trigger, .tgr</td></tr><tr><td>C/C++</td><td>Yes</td><td>.c, .cc, .cpp, .cxx, .h, .hpp, .hxx</td></tr><tr><td>CSharp</td><td>Yes</td><td>.aspx, .cs</td></tr><tr><td>Go</td><td>Yes</td><td>.go</td></tr><tr><td>Java</td><td>Yes</td><td>.java, .jsp, jspx</td></tr><tr><td>JavaScript/TypeScript</td><td>Yes</td><td>.ejs, .es, .es6, .htm, .html, .js, .jsx, .ts, .cts, .mts, .tsx, .vue, .mjs, .cjs</td></tr><tr><td>Kotlin</td><td>Yes</td><td>.kt</td></tr><tr><td>PHP</td><td>Yes</td><td>.php, .phtml, .module, .inc, .install, .theme, .profile</td></tr><tr><td>Python</td><td>Yes</td><td>.py</td></tr><tr><td>Ruby</td><td>No</td><td>.erb, .haml, .rb, .rhtml, .slim</td></tr><tr><td>Scala</td><td>Yes</td><td>.scala</td></tr><tr><td>Swift</td><td>Yes</td><td>.swift</td></tr><tr><td>Visual Basic</td><td>Yes</td><td>.vb</td></tr></tbody></table>

### File size limit for Snyk Code analysis

Snyk Code automatically excludes the following files from analysis:

* On the Web UI - files that are larger than 1MB.
* On the CLI and IDE - files that are larger than 1MB.
* Minified JS files with 3 or fewer lines.

### Filename length limitation

The analysis is available only for files with names shorter than or equal to 255 characters. If the filename exceeds this limit, you receive an error. To ensure that all files are being analyzed, Snyk recommends shortening long filenames.

### Unicode character encoding

Snyk Code only accepts source code files in UTF-8 encoding. Consider converting source files to this encoding type before importing them into Snyk.

### Framework support

To support a specific framework, Snyk Code must both support the relevant language and be trained on Projects using the framework. The found patterns are then annotated by the security team and extended by curated content.

Most frameworks are partially supported out of the box, as Snyk Code needs only to parse the code to analyze it. In some cases, frameworks may require specific rules, or require specific program analysis engine updates, or both.

If you notice any gaps in support for a specific framework, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

### Framework support levels

Snyk categorizes framework support into two levels: Comprehensive and Partial.

Comprehensive support indicates the following:

* Sources and sinks: Snyk has thoroughly identified and included all relevant sources and sinks.
* Data flow testing: Extensive testing has been conducted to ensure comprehensive data flow coverage.
* Engine support: The Snyk Code engine is fully optimized for this framework.
* Limitations: Snyk is unaware of any limitations. If you encounter a false negative, please report it to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

Partial support indicates the following:

* Sources and sinks: Snyk coverage is limited, and some sources, sanitizers, or sinks may be missing.
* Data flow testing:  Snyk has conducted some testing.
* Engine support: Compatibility of the engine with this framework is limited, potentially affecting analysis accuracy.
* Limitations: False negatives in taint analysis or source and sink identification may still occur.

Partial support for a framework typically involves a mix of these factors. For instance, some sources or sinks may be missing, and while the engine might offer better support, further data flow testing has not been conducted to ensure the analysis is fully reliable.

Snyk continually expands framework coverage and improves the accuracy of analysis.

### How Snyk Code analysis works

Snyk scans your codebase following this sequence:

1. The source code is analyzed, generating an event graph. The event graph is similar to a code map that helps Snyk understand how different parts of the code are related. There are two node types, each node in the graph representing something that happens in the code. Some represent parts of the code, and others represent how the code is used.
2. Rules are run against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If a match is found, Snyk looks for a vulnerability in the event graph, identifying where problems might be hiding in the code.

For more information, see [Snyk Code AI Engine](../scan-with-snyk/snyk-code/#ai-engine). For more information about Snyk Code language support, see [Supported languages, package managers, and frameworks (Snyk Code)](./#code-analysis-snyk-code).

## Language support and CLI, CI/CD, and SCM integrations

### CLI for Snyk Code

To start testing your code using Snyk Code through the CLI, open your repository in a terminal and run `snyk code test`.

For information about customizing test options, running other commands, excluding directories and files, and viewing and exploring the results in different formats, see the following:

* [Available commands](../snyk-cli/commands/#available-commands)
* [Exclude directories and files from Snyk Code CLI tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Output test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Export test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)
* [snyk-to-html](../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)

After you have run `snyk code test`, you can:

* [Open a Fix PR ](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/)
* [Configure PR Checks](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### CLI for Snyk Open Source

Ensure you have installed the relevant package manager and you have included the relevant manifest files supported by Snyk before testing.

To test your Open Source Project for vulnerabilities, run the `snyk test` command.

### Steps to start using SCM integrations

* [Set up an integration](../getting-started/#set-up-a-snyk-integration).
* For details, see [Snyk SCM integrations](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/).
* For language-specific information, see [Git repositories with Maven and Gradle](java-and-kotlin/git-repositories-with-maven-and-gradle.md), [Git repositories and JavaScript](javascript/git-repositories-and-javascript.md), and [Git repositories and Python](python/git-repositories-and-python.md).
