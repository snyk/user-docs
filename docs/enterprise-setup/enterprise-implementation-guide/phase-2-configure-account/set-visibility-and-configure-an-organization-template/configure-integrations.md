# Configure Integrations

### Configure your SDLC Integration Points (IDE, Git repo, CI/CD, Container Registry)

{% hint style="info" %}
To see information about importing Projects, see [Import Projects](../../phase-3-gain-visibility/import-projects.md) and [Rollout](../../phase-5-initial-rollout-to-team/).
{% endhint %}

When you create your Organizations, you must decide where to add Snyk in your Software Development Lifecycle (SDLC).&#x20;

#### Visibility then gating

We recommend you first gain visibility of your vulnerabilities, with features related to gating disabled initially.

The first step is to import your Projects using one of the available methods, and then after you have full oversight of any existing issues/vulnerabilities, you can start to gradually introduce gating methods to prevent new issues from being added to your codebase.

{% hint style="info" %}
Notification settings (email notifications)

* Snyk suggests that you initially disable all email notifications so that users do not receive a large number of notifications while projects are imported.
* You would disable this at the ‘Group’ level (for new organizations), and at the ‘Organization’ level for all existing Organizations.&#x20;
{% endhint %}

### Integration tools

The most common tools to integrate Snyk into your SDLC to gain visibility are:

#### Git Repository

Snyk can integrate with multiple Git repositories to help you track, monitor, and fix the issues and vulnerabilities in your code. When importing a repository, Snyk will scan and monitor the ‘default’ branch for vulnerabilities.&#x20;

Monitoring takes the form of&#x20;

* daily scanning of a specified branch, even when you are not working on it,&#x20;
* Pull Request/Merge Checks

Source control scanning of your Git repositories is suitable for the majority of supported languages, but please note that if you use a private package manager (such as Artifactory), this must be integrated with Snyk to scan your private packages.

Importing Projects through a Git repository integration can be done manually through the browser, or you can leverage the API, either by using the [snyk-api-import](../../../../snyk-api-info/other-tools/tool-snyk-api-import/) tool to import repositories in bulk, or the [API endpoint](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization) to import specific repositories (which can be inserted into a pipeline).

{% hint style="info" %}
For Snyk Enterprise customers, it's strongly suggested to use the **GitHub Enterprise integration card** on the Integrations screen. You do not need to be a GitHub Enterprise customer to use this option, however by using this option, it allows a Personal Access Token (PAT) to be used, whereas OAUTH, provided via the GitHub Integration card, provides an inconsistent experience in terms of access in the interface.
{% endhint %}

If you import your Project from a Git repository, you can also configure Snyk [PR Checks](../../../../scan-using-snyk/run-pr-checks/) and auto-fix PRs. These can prevent new security issues from entering your codebase by automatically scanning code changes in real-time whenever you submit a PR in your Git repository.

* Allows scanning and visibility earlier in the software development lifecycle by checking all submitted PRs for security issues.

{% hint style="info" %}
To disable gating initially, use the daily monitoring that's automatically configured when a project is onboarded in Snyk, and disable PR/MR Checks in the configuration.

* Automatic fixes Snyk can position
  * Automatic fix PRs
  * Automatic dependency upgrade PRs
  * Snyk vulnerability patches
* PR Checks
  * Open Source security and licenses
  * Code analysis (Beta)
{% endhint %}

&#x20;Similarly you may want to disable fix and upgrade PR features.

#### CI/CD (Build Pipelines)

Keep your applications secure by preventing deployment of vulnerable applications or components (registries), adding Snyk in the build as a step of the pipeline.

The CLI offers

* More accurate for a number of specific package managers, (such as Scala, Gradle, and GO).
* Supports private packages without needing to configure an additional integration, given that your build environment will have access to your private packages.
* Give visibility to components that are pushed to production by either breaking builds and reporting to Snyk or only reporting to Snyk.

There are a number of [CI/CD integrations](../../../../integrations/snyk-ci-cd-integrations/) available, or you can utilise the [Snyk CLI](../../../../snyk-cli/) as part of your pipeline to have more flexibility in the tests you are running.

In the initial phase, we recommend using the “monitor” feature to import information into Snyk so you can see any discovered issues (unless you are already importing your repos using a source control integration to achieve this). Later, when you want to start gating/blocking new vulnerabilities from being added, you can introduce “test” features - initially failing builds on critical issues, and then gradually adapting the fail criteria over time.

{% hint style="info" %}
For snyk iac test --report and snyk code test --report (beta), finding issues will result in build possibly stopping with a non-zero response code. \
\
If you want to passively test,  the inclusion of the --report argument requires either setting the build step to always continue or an alternative like concatenating logic equating to "or true" (i.e. snyk code test --report || true). The exact syntax will depend on the ecosystem the CLI is run in.&#x20;
{% endhint %}

Plugins like [Snyk Filter](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-filter) for advanced filtering and [Snyk Delta](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-delta) for highlighting new issues are quite popular when configuring pipelines.

Demonstrations of various pipeline integrations can be found on [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)





