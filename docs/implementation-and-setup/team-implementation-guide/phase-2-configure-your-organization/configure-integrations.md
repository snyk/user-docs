# Configure integrations

Integrate Git or CI/CD integrations as identified in the previous phase.

## Git repositories

For Git integration, Snyk suggests the following settings:

To disable gating initially, use the daily monitoring that is automatically configured when a Project is onboarded in Snyk, and disable the following PR/MR Checks in the configuration settings:

* Automatic fixes Snyk can position
  * Automatic fix PRs
  * Automatic dependency upgrade PRs
  * Snyk vulnerability patches.
* Pull request status checks for
  * Open Source security and licenses
  * Code analysis

## CI/CD (Build pipelines)

Keep your applications secure by preventing deployment of vulnerable applications or components (registries), adding Snyk in the build as a step of the pipeline.

The CLI provides:

* Better resolution of dependencies for a number of specific package managers (such as Scala, Gradle, and GO).
* Supports private packages without having to configure an additional integration, providing that your build environment will have access to your private packages.
* Give visibility to components that are pushed to production by either breaking builds and reporting to Snyk or only reporting to Snyk.

There are several [CI/CD integrations](../../../developer-tools/snyk-ci-cd-integrations/) available, or you can use the [Snyk CLI](../../../developer-tools/snyk-cli/) as part of your pipeline, in order to have more flexibility in the tests you are running.

In the initial phase, Snyk recommends using the “monitor” feature to import information into Snyk so you can see any discovered issues (unless you are already importing your repos using a source control integration to achieve this). Later, when you want to start gating/blocking new vulnerabilities from being added, you can introduce “test” features - initially failing builds on critical issues and then gradually adapting the fail criteria over time.

{% hint style="info" %}
For `snyk iac test --report`, finding issues will result in the build possibly stopping with a non-zero response code. \
\
If you want to test passively,  the inclusion of the `--report` argument requires either setting the build step to always continue or an alternative like concatenating logic equating to "or true" (for example `snyk iac test --report || true`). The exact syntax depends on the ecosystem the CLI is run in.&#x20;
{% endhint %}

When configuring pipelines, you can use popular plugins like [`snyk-filter`](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md) for advanced filtering.&#x20;

{% hint style="info" %}
Some plugins that require API access may not be usable on the Team plan, as API access is available on the Snyk Enterprise plan.
{% endhint %}

To see demonstrations of pipeline integrations, see [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples).
