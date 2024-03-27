# Configure integrations

## Configure your SDLC Integration Points (IDE, Git repo, CI/CD, Container Registry)

{% hint style="info" %}
For information about importing Projects, see [Import Projects](../../phase-3-gain-visibility/import-projects.md) and [Rollout](../../phase-5-initial-rollout-to-team/).
{% endhint %}

When you create your Organizations, you must decide where to add Snyk in your Software Development Lifecycle (SDLC).&#x20;

Snyk recommends that you first gain visibility for your vulnerabilities, with features related to gating disabled initially.

The first step is to import your Projects using one of the available methods, and then after you have full visibility for any existing issues and vulnerabilities, you can start to gradually introduce gating methods to prevent new issues from being added to your codebase.

{% hint style="info" %}
**Notification settings (email notifications)**

* Snyk suggests that you initially disable all email notifications so that users do not receive a large number of notifications while Projects are being imported.
* You would disable this at the Group level for new Organizations and at the Organization level for all existing Organizations.
{% endhint %}

## Integration tools

The most common tools to integrate Snyk into your SDLC to gain visibility are described here.

### Git repository

Snyk can integrate with multiple Git repositories to help you track, monitor, and fix the issues and vulnerabilities in your code.

If you have on-premise Git repositories, you must configure and run [Snyk Broker](../../../../enterprise-configuration/snyk-broker/) to enable the integration.

When importing a repository, Snyk will scan and monitor the default branch for vulnerabilities.&#x20;

Monitoring takes the form of

* daily scanning of a specified branch, even when you are not working on it,&#x20;
* Pull Request and Merge Checks

Source control scanning of your Git repositories is suitable for the majority of supported languages, but note that if you use a private package manager, such as Artifactory, this must be integrated with Snyk to scan your private packages.

Importing Projects through a Git repository integration can be done manually through the browser, or you can leverage the API, either by using the [snyk-api-import](../../../../snyk-api/snyk-tools/tool-snyk-api-import/) tool to import repositories in bulk, or the [Import Targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets) API v1 endpoint to import specific repositories, which can be inserted into a pipeline.

{% hint style="info" %}
For Snyk Enterprise customers, it is strongly suggested to use the **GitHub Enterprise integration card** on the Snyk Integrations page. You do not need to be a GitHub Enterprise customer to use this option; however, using this option allows a Personal Access Token (PAT) to be used, whereas OAUTH, provided through the GitHub Integration card, provides an inconsistent experience in terms of access in the interface.
{% endhint %}

If you import your Project from a Git repository, you can also configure Snyk [PR Checks](../../../../scan-with-snyk/run-pr-checks/) and auto-fix PRs. These can prevent new security issues from entering your codebase by automatically scanning code changes in real-time whenever you submit a PR in your Git repository.

This allows scanning and visibility earlier in the software development lifecycle by checking all submitted PRs for security issues.

{% hint style="info" %}
To disable gating initially, use the daily monitoring that is automatically configured when a Project is onboarded in Snyk, and disable PR/MR Checks in the configuration.

* Automatic fixes Snyk can provide
  * Automatic fix PRs
  * Automatic dependency upgrade PRs
  * Snyk vulnerability patches
* PR Checks are available for
  * Open Source security and licenses
  * Code analysis (Beta)
{% endhint %}

&#x20;Similarly, you may want to disable fix and upgrade PR features.

### CI/CD (Build Pipelines)

Keep your applications secure by preventing deployment of vulnerable applications or components (registries), by adding Snyk in the build as a step of the pipeline.

The CLI offers

* Optimum accuracy for a number of specific package managers, such as Scala, Gradle, and GO
* Support for private packages without the need to configure an additional integration, given that your build environment will have access to your private packages.
* Visibility to components that are pushed to production by either breaking builds and reporting to Snyk or only by reporting to Snyk.

There are a number of [CI/CD integrations](../../../../integrate-with-snyk/snyk-ci-cd-integrations/) that you can use, or you can use the [Snyk CLI](../../../../snyk-cli/) as part of your pipeline to have more flexibility in the tests you are running.

In the initial phase, Snyk recommends using the `monitor` feature to import information into Snyk so you can see any discovered issues, unless you are already importing your repos using a source control integration to achieve this. Later, when you want to start gating and blocking new vulnerabilities from being added, you can introduce `test` features, initially failing builds on critical issues, and then gradually adapting the fail criteria over time.

{% hint style="info" %}
For `snyk iac test --report` and `snyk code test --report` (beta), finding issues will result in the build possibly stopping with a non-zero response code. \
\
If you want to passively test this, including the `--report` option requires either setting the build step to always continue or an alternative like concatenating logic equating to `or true`, that is,  `snyk code test --report || true.` The exact syntax will depend on the ecosystem the CLI is running in.&#x20;
{% endhint %}

Plugins like [Snyk Filter](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-filter) for advanced filtering and [Snyk Delta](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-delta) for highlighting new issues are quite popular for configuring pipelines.

Demonstrations of various pipeline integrations can be found on [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)





