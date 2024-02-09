# Introduction to Snyk-supported languages and frameworks

## Snyk Code

### Supported file extensions

<details>

<summary>Snyk Code supports the following file extensions</summary>

* apex

<!---->

* ASPX

<!---->

* c

<!---->

* cc

<!---->

* cjs

<!---->

* cls

<!---->

* cpp

<!---->

* CS

<!---->

* ejs

<!---->

* erb

<!---->

* es

<!---->

* es6

<!---->

* go

<!---->

* h

<!---->

* haml

<!---->

* hpp

<!---->

* htm

<!---->

* html

<!---->

* hxx

<!---->

* java

<!---->

* js

<!---->

* jspx

<!---->

* jsx

<!---->

* jsp

<!---->

* kt

<!---->

* mjs

<!---->

* php

<!---->

* py

<!---->

* rb

<!---->

* rhtml

<!---->

* scala

<!---->

* slim

<!---->

* swift

<!---->

* ts

<!---->

* tsx

<!---->

* trigger

<!---->

* vb

<!---->

* vue

<!---->

* xml

</details>

### File size limit for Snyk Code analysis

Snyk Code automatically excludes the following files from analysis:

* On the Web UI - files that are larger than 1MB.
* On the CLI and IDE - files that are larger than 1MB.
* Minified JS files with 3 or fewer lines.

### Filename length limitation

The analysis is available only for files with names shorter or equal to 255 characters. You receive an error if the filename exceeds this limit. To ensure that all files are being analyzed, consider shortening long filenames.

### Framework support

To support a specific framework, Snyk Code must both support the relevant language and be trained on Projects using the framework. The found patterns are then annotated by the security team and extended by curated content.

Most frameworks are supported out of the box, as Snyk Code needs only to parse the code to analyze it. In some cases, they might require specific rules, or it might require specific program analysis engine updates or both.&#x20;

If you notice any gaps in support for a specific framework, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

### How Snyk works for code analysis&#x20;

Snyk scans your codebase following this sequence

1. The source code is analyzed, generating an event graph. The event graph is similar to a code map that helps Snyk understand how different parts of the code are related. There are two node types, each node in the graph representing something that happens in the code. Some represent parts of the code, and others represent how the code is used.
2. Rules are run against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If a match is found, Snyk looks for a vulnerability in the event graph, identifying where problems might be hiding in the code.&#x20;

:link: See [Snyk Code AI Engine](broken-reference).

### More information about language support for Snyk Code

For more information, see [Supported languages, frameworks, and feature availability overview, Snyk Code section](supported-languages-frameworks-and-feature-availability-overview.md#code-analysis-snyk-code).

### Code Quality

{% hint style="warning" %}
**Release status**

Code Quality is an experimental feature. If you are interested in using the feature, contact your Snyk team.
{% endhint %}

## Snyk Open Source

### How Snyk works for open source and licensing

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any packages anywhere in that tree.

The way Snyk analyzes and builds the graph varies depending on the language and package manager of the Project, as well as the location of your Project.

### More information about language support for Snyk Open Source

For more information, see [Supported languages, frameworks, and feature availability overview, Open Source section](supported-languages-frameworks-and-feature-availability-overview.md#open-source-and-licensing-snyk-open-source).

\
