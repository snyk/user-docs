---
nav_context: agnostic
description: >-
  Technical requirements for Snyk Code and Snyk Open Source, including file
  encoding and how Snyk builds the dependency tree
last_reviewed: 2026-10-27
last_reviewed_source: focused-edit
---

# Technical specifications and guidance

## Unicode character encoding

Both Snyk Code and Snyk Open Source accept source code files in UTF-8 encoding. Convert source files to this encoding type before importing them into Snyk.

## Snyk Open Source

Snyk analyzes and builds the dependency tree depending on the language and package manager for the Project and the location of the Project.

### How Snyk Open Source and licensing works

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must build your Project. For details, visit [Open Source Projects that must be built before testing with the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli).
{% endhint %}

Snyk builds a dependency graph (dependency tree) and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

### Snyk policies in Open Source

For information on managing dependencies and vulnerabilities from your developer workflows using policies, visit the following resources:

* [Defining a secure open-source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

### Open Source license compliance

To check compliance for open source licenses, visit [Snyk License Compliance Management](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management).

## Snyk Code

### File size limit for Snyk Code analysis

Snyk Code automatically excludes the following files from analysis:

* In the Web UI: files larger than 1MB
* In the CLI and IDE: files larger than 1MB
* Minified JavaScript files with three or fewer lines

### Filename length limitation

Snyk Code analyzes only files with names of 255 characters or fewer. If a filename exceeds this limit, you receive an error. To ensure that Snyk Code analyzes all files, Snyk recommends shortening long filenames.

### Framework support

To support a specific framework, Snyk Code must both support the relevant language and be trained on Projects using the framework. The Snyk security team then annotates the patterns found and extends them with curated content.

Most frameworks are partially supported out of the box, because Snyk Code needs only to parse the code to analyze it. In some cases, frameworks require specific rules, specific program analysis engine updates, or both.

If you notice any gaps in support for a specific framework, [contact Snyk Support](https://support.snyk.io).

Snyk categorizes framework support into two levels: Comprehensive and Partial.

Comprehensive support includes:

* Sources and sinks: All relevant sources and sinks are identified.
* Data flow testing: Extensive testing provides thorough data flow coverage.
* Engine support: The Snyk Code engine is fully optimized for this framework.
* Limitations: No known limitations. Report any false negatives to [Snyk Support](https://support.snyk.io).

Partial support includes:

* Sources and sinks: Coverage can be limited, with some missing elements.
* Data flow testing: Snyk has conducted limited testing.
* Engine support: Compatibility is restricted, which can affect analysis accuracy.
* Limitations: Potential for false negatives in taint analysis or source/sink identification.

Snyk continuously expands its framework coverage and improves analysis accuracy.

### Template file analysis

Snyk Code analyzes template files together with the application code that renders them. Snyk Code follows data that reaches a template from your application code into the template, so it reports a cross-site scripting vulnerability that becomes exploitable only where the template writes its output. The reported data flow spans both the application code and the template.

Snyk Code takes the escaping behavior of the template engine into account. Snyk Code does not report output written through the default escaping of the engine. Snyk Code treats output written through a construct that bypasses escaping, such as a raw or unescaped directive, as a sink.

The pairing of a template engine with the language that renders the template defines support. Snyk Code supports the following pairs. Snyk Code analyzes template files in addition to the source file formats listed for each language.

| Template engine | File extensions | Languages |
| :--- | :--- | :--- |
| EJS | `.ejs` | JavaScript |
| FreeMarker | `.ftl`, `.ftlh`, `.ftlx` | Java, Kotlin, Groovy |
| Handlebars | `.hbs`, `.handlebars` | JavaScript |
| Jinja2 | `.j2`, `.jinja`, `.jinja2` | Python |
| Mustache | `.mustache`, `.mu` | JavaScript, C# |
| Pug | `.pug`, `.jade` | JavaScript |
| Razor | `.cshtml`, `.razor` | C# |
| Thymeleaf | `.html` | Java, Kotlin, Groovy |
| Twig | `.twig` | PHP |
| Velocity | `.vm`, `.vtl` | Java, Kotlin, Groovy |

Snyk is expanding this coverage. If a template engine or a pair you use is not listed, [contact Snyk Support](https://support.snyk.io).

### How Snyk Code analysis works

Snyk scans your codebase following this sequence:

1. Snyk analyzes the source code to generate an event graph. The event graph is similar to a code map, helping Snyk understand how different parts of the code are related. Each node in the graph represents something that happens in the code. There are two node types: some nodes represent parts of the code, and others represent how the code is used.
2. Snyk runs rules against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If Snyk finds a match, Snyk looks for a vulnerability in the event graph, identifying where problems are hiding in the code.

For more information, visit [Snyk Code AI Engine](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code#ai-engine). For more information about Snyk Code language support, visit [Supported languages, package managers, and frameworks](supported-languages-package-managers-and-frameworks.md).

## Language support and CLI, CI/CD, and SCM integrations

Snyk supports a variety of programming languages. You can integrate Snyk into your development workflow through CLI commands, continuous integration/continuous delivery (CI/CD) pipelines, and source code management (SCM) integrations.

You can use these tools to automatically check your code for security vulnerabilities as you develop your software. This ensures that strong security practices are part of your development process.

For more details, visit the following pages:

* CLI for [Snyk Open Source](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source) and [Snyk Code](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code)
* CI/CD for [Snyk Open Source](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-open-source-specific-ci-cd-strategies) and [Snyk Code](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations/use-snyk-code-in-the-ci-cd-pipeline)
* [SCM integrations](https://docs.snyk.io/developer-tools/integrations/scm-integrations/organization-level-integrations) for Snyk Open Source and Snyk Code
