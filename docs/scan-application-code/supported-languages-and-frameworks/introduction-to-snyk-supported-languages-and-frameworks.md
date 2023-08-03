# Introduction to Snyk supported languages and frameworks

## How Snyk works for code analysis

Snyk scans your codebase following this sequence

1. The source code is analyzed, generating an event graph. The event graph is similar to a code map that helps Snyk understand how different parts of the code are related. There are two node types, each node in the graph representing something that happens in the code. Some represent parts of the code, and others represent how the code is used.
2. Rules are run against the event graph to find matches. The rules act as a checklist of known vulnerabilities that Snyk looks for in the event graph.
3. If a match is found, Snyk looks for a vulnerability in the event graph, identifying where problems might be hiding in the code.&#x20;

:link: See [Snyk Code AI Engine](../snyk-code/snyk-code-key-features/snyk-code-ai-engine.md).

## How Snyk works for open source and licensing

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any packages anywhere in that tree.

The way Snyk analyzes and builds the graph varies depending on the language and package manager of the Project, as well as the location of your Project.

\
