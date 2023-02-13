# Reachable vulnerabilities

## Introduction

There are multiple ways to [prioritize vulnerabilities](evaluating-and-prioritizing-vulnerabilities.md), one of which is to look for **reachable vulnerabilities**.

A reachable vulnerability has a path from your code to the vulnerable function in a dependency.

By analyzing how your code uses open source dependencies, and also how the open source dependencies interact with each other, Snyk can provide powerful insight into which vulnerabilities are reachable:

* Directly from your own code, _and also_
* Indirectly via transitive dependencies

## How finding reachable vulnerabilities works

First Snyk identifies the vulnerable functions associated with specific vulnerabilities.

Next Snyk builds a call graph that is the union of the calls within the application and the calls within open source libraries scanned by Snyk. If a vulnerable function is reachable using this unified call graph from the application code, Snyk identifies the vulnerability as reachable.

Each vulnerability then receives a reachability status:

**Reachable** - A direct or indirect path from your code to vulnerable functions was found OR

**No path found** - No path found from your code to vulnerable functions

Note that when a "no path found" status is given, it should _not_ be inferred that the vulnerability is definitely unreachable.

Producing a precise call graph may require that the analysis sometimes under-approximates the set of reachable functions.

For example, although the Snyk solution does support it, it is not always possible to give a precise answer if reflection is used. In such a case, there is a tradeoff between returning a large set of false positives, or possibly returning that no path was found even if a call path is theoretically possible.

Therefore, Snyk recommends that you use reachability as one signal in a broader strategy for [prioritizing](./) issues. You may want to give reachable vulnerabilities higher priority as there is more evidence it is exploitable, but do not ignore the rest.

## Supported languages and integrations

Reachable vulnerabilities analysis is available for **Java** (Maven and Gradle) in the [supported versions](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven) using the **GitHub** integration.

Reachable vulnerability analysis via the Snyk CLI, other Git integrations and other languages are not currently supported.

## Scanning for Reachable Vulnerabilities using Git Integrations

{% hint style="info" %}
To provide this feature, Snyk takes a temporary copy of your Git repository contents.

For more information see [How Snyk handles your data](../../snyk-processes/how-snyk-handles-your-data.md).
{% endhint %}

1. Set up your [GitHub integration](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration).
2. Opt in to Reachable Vulnerabilities analysis
   * In the **Organization** settings, go to the **Languages** section.
   * Go to the **Reachable vulnerabilities analysis** section.
   * Opt in to the **Reachable vulnerabilities analysis** and save your changes.
3. Go to the **import projects** page and choose the repositories to import to Snyk.
4. Selected projects are imported and analyzed for reachable vulnerabilities issues.

If you have existing projects you want to add reachability analysis to, opt in via the Organization settings and retest the projects manually or wait for the daily recurring test.

## Project page

After import or testing of a project via Snyk UI, the project is monitored by Snyk and the results of the reachable vulnerabilities analysis appear in the Project page in the following places:

1. **Filters** - Allows you to focus on reachable vulnerabilities first by filtering results based on reachability.
2. **Reachability badge** - Allows you to quickly understand the reachability level of vulnerabilities.
3. **Call path** - Allows you to see the path from your code to the vulnerable function to validate the result.
4. **Priority Score** - Snyk will increase the Priority Score for reachable vulnerabilities, to allow you to see these issues first.

<figure><img src="../../.gitbook/assets/image (124) (1) (1) (1) (2) (1).png" alt="Reachability Call Path"><figcaption><p>Reachability Call Path</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (126) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Priority Score"><figcaption><p>Priority Score</p></figcaption></figure>

## Reports

You can filter by reachability status to quickly show the reachable issues.

<figure><img src="../../.gitbook/assets/image (137) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Issue Filters"><figcaption><p>Issue Filters</p></figcaption></figure>
