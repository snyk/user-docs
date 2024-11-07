# Snyk Code language support technical specifications and guidance

## Supported file extensions for Snyk Code

<table><thead><tr><th>Language</th><th width="215">Interfile Support</th><th>Supported Extension</th></tr></thead><tbody><tr><td>Apex</td><td>Yes</td><td>.cls, .trigger, .tgr</td></tr><tr><td>C/C++</td><td>Yes</td><td>.c, .cc, .cpp, .cxx, .h, .hpp, .hxx</td></tr><tr><td>CSharp</td><td>Yes</td><td>.aspx, .cs</td></tr><tr><td>Go</td><td>Yes</td><td>.go</td></tr><tr><td>Java</td><td>Yes</td><td>.java, .jsp, jspx</td></tr><tr><td>JavaScript/TypeScript</td><td>Yes</td><td>.ejs, .es, .es6, .htm, .html, .js, .jsx, .ts, .cts, .mts, .tsx, .vue, .mjs, .cjs</td></tr><tr><td>Kotlin</td><td>Yes</td><td>.kt</td></tr><tr><td>PHP</td><td>Yes</td><td>.php, .phtml, .module, .inc, .install, .theme, .profile</td></tr><tr><td>Python</td><td>Yes</td><td>.py</td></tr><tr><td>Ruby</td><td>No</td><td>.erb, .haml, .rb, .rhtml, .slm</td></tr><tr><td>Scala</td><td>Yes</td><td>.scala</td></tr><tr><td>Swift</td><td>Yes</td><td>.swift</td></tr><tr><td>Visual Basic</td><td>Yes</td><td>.vb</td></tr></tbody></table>

## File size limit for Snyk Code analysis

Snyk Code automatically excludes the following files from analysis:

* On the Web UI - files that are larger than 1MB.
* On the CLI and IDE - files that are larger than 1MB.
* Minified JS files with 3 or fewer lines.

## Filename length limitation

The analysis is available only for files with names shorter than or equal to 255 characters. If the filename exceeds this limit, you receive an error. To ensure that all files are being analyzed, Snyk recommends shortening long filenames.

## Unicode character encoding

Snyk Code only accepts source code files in UTF-8 encoding. Consider converting source files to this encoding type before importing them into Snyk.

## Framework support

To support a specific framework, Snyk Code must both support the relevant language and be trained on Projects using the framework. The found patterns are then annotated by the security team and extended by curated content.

Most frameworks are partially supported out of the box, as Snyk Code needs only to parse the code to analyze it. In some cases, frameworks may require specific rules, or require specific program analysis engine updates, or both.

If you notice any gaps in support for a specific framework, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

## Framework support levels

Snyk categorizes framework support into two levels: Comprehensive and Partial.

Comprehensive Support indicates the following:

* Sources and Sinks: Snyk has thoroughly identified and included all relevant sources and sinks.
* Data Flow Testing: Extensive testing has been conducted to ensure comprehensive data flow coverage.
* Engine Support: The Snyk Code engine is fully optimized for this framework.
* Limitations: Snyk is unaware of any limitations. If you encounter a false negative, please report it to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

Partial Support indicates the following:

* Sources and Sinks: Snyk coverage is limited, and some sources, sanitizers, or sinks may be missing.
* Data Flow Testing:  Snyk has conducted some testing.
* Engine Support: Compatibility of the engine with this framework is limited, potentially affecting analysis accuracy.
* Limitations: False negatives in taint analysis or source and sink identification may still occur.

Partial support for a framework typically involves a mix of these factors. For instance, some sources or sinks may be missing, and while the engine might offer better support, further data flow testing has not been conducted to ensure the analysis is fully reliable.

Snyk continually expands framework coverage and improves the accuracy of analysis.

## How Snyk Code analysis works

Snyk scans your codebase following this sequence:

1. The source code is analyzed, generating an event graph. The event graph is similar to a code map that helps Snyk understand how different parts of the code are related. There are two node types, each node in the graph representing something that happens in the code. Some represent parts of the code, and others represent how the code is used.
2. Rules are run against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If a match is found, Snyk looks for a vulnerability in the event graph, identifying where problems might be hiding in the code.

For more information, see [Snyk Code AI Engine](../scan-with-snyk/snyk-code/#ai-engine). For more information about Snyk Code language support, see [Supported languages, package managers, and frameworks (Snyk Code)](./#code-analysis-snyk-code).
