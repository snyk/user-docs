---
nav_context: agnostic
description: >-
  How to configure Snyk SCM and CI/CD integrations for your Organization,
  including recommended settings
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Configure integrations

Set up the SCM or CI/CD integrations you identified in the previous phase.

## SCM integrations

For SCM integrations, Snyk recommends the following settings.

To disable gating initially, use the daily monitoring that Snyk configures automatically when you onboard a Project, and disable the following in the integration settings:

* Automatic fixes
  * Automatic fix PRs
  * Automatic dependency upgrade PRs
  * Snyk vulnerability patches
* Pull request status checks
  * Open Source security and licenses
  * Code analysis

## CI/CD integrations

Add Snyk as a step in your build pipeline to keep your applications secure by preventing deployment of vulnerable applications or components (registries).

The CLI provides:

* Better resolution of dependencies for a number of specific package managers (such as Scala, Gradle, and Go)
* Support for private packages without configuring an additional integration, if your build environment has access to your private packages
* Visibility into components pushed to production, either by breaking builds and reporting to Snyk or by only reporting to Snyk

Several [CI/CD integrations](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations) are available. Alternatively, use the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli) in your pipeline for more flexibility in the tests you run.

In the initial phase, Snyk recommends using the `monitor` command to import information into Snyk so you can see any discovered issues (unless you are already importing your repos using a source control integration to achieve this). Later, when you want to start gating new vulnerabilities, introduce the `test` command. Initially, fail builds on critical issues, and then gradually adapt the fail criteria over time.

{% hint style="info" %}
For `snyk iac test --report`, if Snyk finds issues, the command returns a non-zero exit code, which can stop the build.

If you want to test passively with the `--report` option, either set the build step to always continue or use an alternative such as appending logic equivalent to "or true" (for example, `snyk iac test --report || true`). The exact syntax depends on the ecosystem the CLI runs in.
{% endhint %}

When configuring pipelines, you can use popular plugins like [`snyk-filter`](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter) for advanced filtering.

{% hint style="info" %}
Some plugins that require API access might not be usable on the Team plan, because API access is available on the Snyk Enterprise plan.
{% endhint %}

To see demonstrations of pipeline integrations, visit [Snyk Labs CI/CD integration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).
