# Snyk Release Process

Snyk features are provided to users in the following types of releases.

{% hint style="info" %}
Not all features follow all these stages, and timelines for each feature to move stages vary depending on the feature.
{% endhint %}

## Alpha

<table><thead><tr><th width="240">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>Internal release only.</td><td>Snyk internal users, potentially some design partners.</td><td>Controlled </td><td>No docs provided.</td></tr></tbody></table>

## Closed Beta

<table><thead><tr><th width="243">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>The first customer-facing rollout of a feature.</td><td>A preselected group of users.</td><td>On invitation by Snyk.</td><td>Provided but may not be public.</td></tr></tbody></table>

**Snyk features in Closed Beta**

* [Configure PR Checks](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)
* [Publish Snyk Code CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/publish-snyk-code-cli-results-and-ignore-issues.md)
* Third-party integrations for Snyk AppRisk:
  * [Nightfall](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#nightfall-setup-guide)
  * [Dynatrace](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#dynatrace-setup-guide)
  * [Sysdig](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#sysdig-setup-guide)
  * [Orca Security](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#orca-security-setup-guide)

## Early Access

<table><thead><tr><th width="246">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>Feature is tested and ready for use, but not available by default</td><td>All users on an opt-in basis. This may include some additional purchase costs.</td><td>Opt-in: on request via Snyk account team, or via Snyk Preview</td><td>Public docs.</td></tr></tbody></table>

**Snyk features in Early Access**

* [Snyk GitHub Cloud App](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app.md)
* Projects
  * [Group Projects by branch or version for monitoring](../snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring.md)
  * [Automatically created Project collections](../snyk-admin/introduction-to-snyk-projects/automatically-created-project-collections.md)
* Snyk Code
  * [Snyk Code custom rules](../scan-with-snyk/snyk-code/snyk-code-custom-rules/)
  * [Fix code vulnerabilities automatically](../scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md)
* Risk Management
  * [Risk Score](../manage-risk/prioritize-issues-for-fixing/risk-score.md)
  * [Reachabilty analysis](../manage-risk/prioritize-issues-for-fixing/reachability-analysis.md)
* Snyk Broker
  * [Snyk Broker commit signing](../enterprise-setup/snyk-broker/snyk-broker-commit-signing.md)
  * [Snyk Code - Clone capability with Broker for Docker](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md)
  * [Universal Broker](../enterprise-setup/snyk-broker/universal-broker/)
* Language support
  * [Improved .NET scanning](../supported-languages-package-managers-and-frameworks/.net/improved-.net-scanning.md)
  * [Snyk CLI pnpm support](../supported-languages-package-managers-and-frameworks/javascript/javascript-for-open-source.md#pnpm)
  * [Improved Gradle SCM scanning](../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning-early-access)
* Third-party integrations for Snyk AppRisk:&#x20;
  * [Veracode](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#veracode-setup-guide)
  * [Checkmarx](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#checkmarx-setup-guide)
  * [SonarQube](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#sonarqube-setup-guide)
  * [GitGuardian](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#gitguardian-setup-guide)
  * [Jira](../integrate-with-snyk/third-party-integrations-for-snyk-apprisk.md#jira-setup-guide)

## General Availability

<table><thead><tr><th width="249">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>Feature is fully enabled.</td><td>All users, subject to standard feature availability.</td><td>Available by default.</td><td>Full public docs.</td></tr></tbody></table>

## Deprecated

<table><thead><tr><th width="256">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>The feature is available, but use is discouraged. </td><td>Current users only</td><td>Available by default.</td><td>Full public docs, marked-up appropriately.</td></tr></tbody></table>

## Maintenance mode

<table><thead><tr><th width="256">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>No new development or updates will be made to the feature. </td><td>Current users only</td><td>Available by default.</td><td>Full public docs, marked-up appropriately.</td></tr></tbody></table>

## End of support

<table><thead><tr><th width="256">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>No new support tickets will be answered. </td><td>Current users only</td><td>Available by default.</td><td>Full public docs, marked-up appropriately.</td></tr></tbody></table>

## End of life

<table><thead><tr><th width="256">Description</th><th>Available to</th><th>Access</th><th>Docs</th></tr></thead><tbody><tr><td>The feature is no longer available. </td><td>No users</td><td>Not available</td><td>No docs available</td></tr></tbody></table>
