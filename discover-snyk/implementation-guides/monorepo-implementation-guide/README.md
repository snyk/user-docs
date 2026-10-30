---
description: Guide to scanning monorepos with Snyk using strategies and best practices tailored to your repository structure
nav_context: agnostic
---

# Monorepo implementation guide

A monorepo is a single repository containing multiple distinct projects with well-defined relationships. The defining characteristic is unified versioning, which makes changes atomic across the entire monorepo. Other common benefits include extensive code sharing, simplified dependency management, large-scale refactoring, collaboration across teams, flexible code ownership, and code visibility.

## Choose your scanning approach

{% hint style="success" %}
**Key decision**: Choose a primary scanning approach, SCM integration or Snyk CLI.
{% endhint %}

When scanning a monorepo, Snyk's recommended approach is to use the [SCM integration](https://docs.snyk.io/developer-tools/integrations/scm-integrations), as long as:

* The monorepo can be scanned within a single Snyk Organization.
* The SCM import completes without timing out.

Because the SCM integration requires Snyk's servers to clone the repository and process a large volume of data, monorepos with many manifest files are more likely to exceed the connection timeout before the scan finishes.

This approach also lets you view code snippets for SAST issues in the Snyk user interface and use Snyk's pull request checks.

If the SCM integration isn't an option, use the Snyk CLI. The CLI gives you flexibility in directory navigation, works with the "affected" logic of your build tooling, and avoids the size and timeout limitations of the SCM integration. To get the best results, follow these steps:

1. [Apply CLI best practices:](cli-best-practices.md) Configure ignores, reporting, Organizations, and pipeline scans.
2. [Tailor your strategy to your monorepo structure:](monorepo-structures/) Follow examples for Functional/Layered, Domain-centric, App-centric, and Flat/Namespace monorepos.
