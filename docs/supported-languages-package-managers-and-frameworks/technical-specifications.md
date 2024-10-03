# Technical specifications

## Snyk Code

### Supported file extensions for Snyk Code

<table><thead><tr><th>Language</th><th width="215">Interfile Support</th><th>Supported Extension</th></tr></thead><tbody><tr><td>Apex</td><td>Yes</td><td>.cls, .trigger, .tgr</td></tr><tr><td>C/C++</td><td>Yes</td><td>.c, .cc, .cpp, .cxx, .h, .hpp, .hxx</td></tr><tr><td>CSharp</td><td>Yes</td><td>.aspx, .cs</td></tr><tr><td>Go</td><td>Yes</td><td>.go</td></tr><tr><td>Java</td><td>Yes</td><td>.java, .jsp, jspx</td></tr><tr><td>JavaScript/TypeScript</td><td>Yes</td><td>.ejs, .es, .es6, .htm, .html, .js, .jsx, .ts, .cts, .mts, .tsx, .vue, .mjs, .cjs</td></tr><tr><td>Kotlin</td><td>Yes</td><td>.kt</td></tr><tr><td>PHP</td><td>Yes</td><td>.php, .phtml, .module, .inc, .install, .theme, .profile</td></tr><tr><td>Python</td><td>Yes</td><td>.py</td></tr><tr><td>Ruby</td><td>No</td><td>.erb, .haml, .rb, .rhtml, .slm</td></tr><tr><td>Scala</td><td>Yes</td><td>.scala</td></tr><tr><td>Swift</td><td>Yes</td><td>.swift</td></tr><tr><td>Visual Basic</td><td>Yes</td><td>.vb</td></tr></tbody></table>

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

Most frameworks are supported out of the box, as Snyk Code needs only to parse the code to analyze it. In some cases, they might require specific rules, or it might require specific program analysis engine updates or both.

If you notice any gaps in support for a specific framework, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

### How Snyk Code analysis works

Snyk scans your codebase following this sequence:

1. The source code is analyzed, generating an event graph. The event graph is similar to a code map that helps Snyk understand how different parts of the code are related. There are two node types, each node in the graph representing something that happens in the code. Some represent parts of the code, and others represent how the code is used.
2. Rules are run against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If a match is found, Snyk looks for a vulnerability in the event graph, identifying where problems might be hiding in the code.

For more information, see [Snyk Code AI Engine](../scan-with-snyk/snyk-code/#ai-engine). For more information about Snyk Code language support, see [Supported languages, package managers, and frameworks (Snyk Code)](./#code-analysis-snyk-code).

## Snyk Open Source

Snyk for Open Source and licensing works as follows.

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any packages anywhere in that tree.

After Snyk has built the dependencies tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

The way Snyk analyzes and builds the dependencies tree varies depending on the language and package manager of the Project, as well as the location of your Project.

For more information about language support for Snyk Open Source, see [Supported languages, package managers, and frameworks (Snyk Open Source)](./#open-source-and-licensing-snyk-open-source).

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of projects that have a package manager, this means a release to the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the Github repo.
{% endhint %}
