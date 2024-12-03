---
cover: .gitbook/assets/Snyk General Banner.webp
coverY: 0
---

# What's new in the Snyk docs?

The most recent updates include significant changes to the user docs, such as features added or removed, structure changes that affect how you find relevant information, and other improvements aimed at enhancing your interaction with the Snyk knowledge base.

## November 2024

### &#x20;**Snyk AppRisk**&#x20;

* The [Asset Dashboard](manage-issues/reporting/available-snyk-reports.md#asset-dashboard) has been redesigned and is included in the Reports section. It now features several enhancements, including a Global filter bar, new data widgets, and the option to export the dashboard as a PDF.
* The [Snyk Broker - AppRisk](enterprise-setup/snyk-broker/snyk-broker-apprisk.md#scm-integrations) documentation has been updated to outline the specific steps necessary to configure each SCM integration with a Snyk Broker token.

### **Snyk Container**

* The list of [operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated to include Ubuntu 24.10 - Oracular Oriole and Ubuntu 24.04 - Noble Numbat 04.
* [How Snyk Container works](scan-with-snyk/snyk-container/how-snyk-container-works/) has been updated with details on the logic Snyk applies when providing public base image recommendations.

### **Snyk Integrations**

* The GitLab page has been updated to state that for the Snyk AppRisk level integration [PAT creation](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab#gitlab-access-tokens), the user generating the GitLab group requires a minimum GitLab permission level of Guest.

### **Other updates**

* The Pull Request Checks section has been updated to include the new [Pull Request Experience for PR Checks](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience).
* The [Supported Languages](supported-languages-package-managers-and-frameworks/) page has been reorganized to provide detailed information about language availability for each Snyk product. Additionally, it provides a list of package managers, frameworks, and features for each supported language.
* A service account using OAuth 2.0 can now be [created through the Snyk Web UI](enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md#create-oauth-service-accounts-through-the-ui).
* The [API index ](snyk-api/api-endpoints-index-and-tips/)now includes entries for each endpoint mentioned in the Snyk user docs.
* The[ Developer IDE and CLI usage](manage-issues/reporting/available-snyk-reports.md#developer-ide-and-cli-usage) report has been enhanced with additional functionalities: **Developer email address** and **PDF export**.
* The [Vulnerabilities Detail](manage-issues/reporting/available-snyk-reports.md#vulnerabilities-detail-report) report has been enhanced with additional functionalities, such as **Target indication** and **Column picker**.

## October 2024

### **Snyk API**&#x20;

* Links to APIs have been updated throughout the Snyk User Docs site.
* The [API](snyk-api/) pages have been extensively updated to reflect changes in how APIs are published.
* The [API Authentication](snyk-api/rest-api/authentication-for-api/) pages have been updated to include when to use a personal token and when to use a service account, and to improve the flow of information.

### &#x20;**Snyk AppRisk**&#x20;

* [Risk factors on assets](manage-assets/#inventory-overview) is now in [EA](getting-started/snyk-release-process.md#early-access) for Snyk AppRisk Pro customers.
* [Asset inventory components](manage-assets/assets-inventory-components.md#clusters) has been updated to include details on clusters.
* Snyk Broker - AppRisk has been updated to include [SonarQube installation steps](enterprise-setup/snyk-broker/snyk-broker-apprisk.md#sonarqube-sast-integration).

### **Snyk CLI and IDEs**

* The [CLI authentication page](snyk-cli/authenticate-to-use-the-cli.md) has been updated for the OAuth 2.0 protocol.
* The page [Debugging the Snyk CLI](snyk-cli/debugging-the-snyk-cli.md) has been added.
* [CLI standalone executables](snyk-cli/install-or-update-the-snyk-cli/#install-with-standalone-executables) have been updated to include Alpine Arm64.
* IDE Eclipse[ plugin](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/) and [JetBrains plugin ](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/)documentation pages have been updated.
* [Authentication information](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/) have been updated for all IDEs.

### **Snyk Integrations**

* [Snowflake Data Share](manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/) is now in [GA](getting-started/snyk-release-process.md#general-availability).
* [Snyk SCM integrations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/) has been updated with additional notices relating to repository retrieval and permission or scope modifications after initial configuration.
* The **GitHub and GitHub Enterprise permission requirements** section title has been changed to [Personal Access Token (PAT) scopes and repository permission requirements for SCMs](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#personal-access-token-pat-scopes-and-repository-permission-requirements-for-scms) to capture PAT and repository permissions for multiple SCM integrations, including the GitHub Cloud App.
* GitHub Cloud App has been added to feature support notices for Fix, Backlog, and Upgrade PRs.
* Snyk SCM integrations has been updated to include a table detailing the [permissions and scopes](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#github-cloud-app-permission-requirements) required for the GitHub Cloud App.

### **Other updates**

* [Getting started](getting-started/) has been updated to centralize content related to everything you need to know before using Snyk.&#x20;
* The [Snyk Jumpstart Services Description](working-with-snyk/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-services-description.md) and [Customer Prerequisite](working-with-snyk/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites.md)s have been updated for AppRisk Pro.
* Scanning methods have been added for the [Dart and Flutter](supported-languages-package-managers-and-frameworks/dart-and-flutter.md) languages.

## September 2024

### **Snyk API**&#x20;

* Most links to APIs have been updated on the docs site.
* Additional entries were made on the [API docs index page](snyk-api/api-endpoints-index-and-tips/).
* [Regional URLs for APIs](snyk-api/rest-api/about-the-rest-api.md#api-url) were updated.

### &#x20;**Snyk AppRisk**&#x20;

* A prerequisites section has been added to the Group level of [GitHub integration](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-enterprise.md#prerequisites), and more details about the [Pull personal repositories](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-enterprise.md#github-integrate-using-snyk-apprisk) option have been added to the same documentation page.&#x20;
* The [Set up Insights for Snyk AppRisk ](manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/)section was updated to emphasize the risk factors availability for each integration option.
* The [Snyk Runtime Sensor](manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md) has been updated to reflect the importance of adopting it to achieve the most effective integration and to access its continuously expanded set of features.&#x20;

### Snyk Broker

The Universal Broker feature is now available in Early Access. The Universal Broker separates deployment and container concerns from connection concerns. It allows for a smaller or a single deployment to support numerous connections of varied types.

### **Snyk CLI**

* The [CLI commands and options summary](snyk-cli/cli-commands-and-options-summary.md) was updated.
* [Authentication](snyk-cli/authenticate-to-use-the-cli.md) has been updated.
* Configuration has been updated: Environment variables for Snyk CLI, [`snyk config`](snyk-cli/commands/config.md) help, [`snyk config environment`](snyk-cli/commands/config-environment.md) help.

### **Snyk  Integrations**

The Snowflake Data Share section has been updated to include a [Data Share Dictionary](manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary.md), designed to help you navigate and build your dataset.

### **Other updates**

* The updated [Regional hosting and data residency](working-with-snyk/regional-hosting-and-data-residency.md) page was published.
* [Glossary](getting-started/glossary.md) terms were updated for SCA, SAST, DAST, and IAST as well as Software Composition Analysis.
* [Early Access](getting-started/snyk-release-process.md#early-access) release status notices were updated.

## August 2024

### **Snyk API**&#x20;

* Links in the API reference docs have been updated.
* The [API endpoints index and notes](snyk-api/api-endpoints-index-and-tips/) have been updated.

### **Snyk AppRisk**&#x20;

* The **Manage Risk > Analytics** pages have been consolidated to better reflect the two Snyk offerings:
  * [Issues Analytics](manage-risk/enterprise-analytics/issues-analytics.md) - now in Early Access for Snyk Enterprise customers.
  * [Application Analytics](manage-risk/enterprise-analytics/application-analytics.md) - available only for Snyk AppRisk Pro customers.
* The [Manage Assets](manage-assets/) documentation has been updated to reflect the addition of Quick filters. Users are only shown quick filters relevant to their entitlement: AppRisk Essentials or AppRisk Pro.

### **Snyk CLI**

* [`snyk auth`](snyk-cli/commands/auth.md) command help updated to reflect OAuth default.
* &#x20; [CLI authentication](snyk-cli/authenticate-to-use-the-cli.md) instructions updated for OAuth default and improved flow.
* [`snyk config environment`](snyk-cli/commands/config-environment.md) command help has been added.
* CLI [support for pnpm added](supported-languages-package-managers-and-frameworks/javascript/javascript-for-open-source.md#pnpm).

### Snyk IDE

* &#x20;[CLI authentication](snyk-cli/authenticate-to-use-the-cli.md) instructions updated for IDE.
* IDE authentication instructions updated: [Eclipse](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/authentication-for-the-eclipse-plugin.md), [Jetbrains](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/authentication-for-the-jetbrains-plugins.md), [VS extension](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-authentication.md), [VS Code extension](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-authentication.md)

### **Snyk Integrations**

* Git repository cloning has been renamed [Workspaces for SCM integrations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/introduction-to-git-repository-integrations/workspaces-for-scm-integrations.md) to better reflect its functionality. Additional detail on [enablement](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/introduction-to-git-repository-integrations/workspaces-for-scm-integrations.md#manage-workspaces) has been added.
* The [relationship](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app.md#how-to-set-up-the-github-cloud-app) between GitHub organizations and Snyk Organizations when integrating with the GitHub Cloud App has been clarified.
* The [deployment stages table](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/introduction-to-git-repository-integrations/deployment-recommendations-for-scm-integrations.md#recommended-deployment-order) for implementing SCM integrations has been updated to reflect AppRisk functionality.

## July 2024

### **Snyk API**

* The API documentation now provides the API Reference and explanatory documentation in the [API section](snyk-api/).
* The [API End of Life (EOL) process and migration guides](api-end-of-life-eol-process-and-migration-guides/) are now published and updated to support the process, which began in July.

### Snyk AppRisk

* [Asset inventory filtering](manage-assets/assets-inventory-components.md#asset-tabs) describes the new, simplified view that provides an improved experience of filtering the assets.
* The [Asset inventory layouts](manage-assets/assets-inventory-layouts.md#inventory-layouts) have been renamed to better reflect their functionality.
* Four new SCM integrations are now available for Snyk AppRisk:&#x20;
  * [Atlassian Compass](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#atlassian-compass)
  * [Harness](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#harness)
  * [OpsLevel](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#opslevel)
  * [Datadog Service Catalog](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#datadog-service-catalog)
* The Snyk AppRisk documentation section has been reorganized to enhance visibility and simplify the adoption of Snyk AppRisk. Here is where you can find the main features:
  * Inventory - [Manage assets](manage-assets/) section
  * Issues - [Prioritize issues for fixing](manage-risk/prioritize-issues-for-fixing/#prioritization-based-on-risk) section
  * Policies - [Assets policies](manage-risk/policies/assets-policies/) section
  * Integrations - [Snyk SCM integrations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) and [Integrate with Snyk](integrate-with-snyk/#integrations-for-snyk-apprisk) sections
  * Snyk AppRisk - [Scan with Snyk](scan-with-snyk/snyk-apprisk/) section
  * Analytics - [Manage risk](manage-risk/enterprise-analytics/application-analytics.md) section

### Snyk Integrations

* A comparison of the GitHub and GitHub Enterprise integrations functions now resides on the [SCM, IDE, and CI/CD integrations](scm-ide-and-ci-cd-integrations/#github-vs-github-enterprise) page.
* Steps for [migrating from the GitHub integration to the GitHub Enterprise integration](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#migrate-to-the-github-enterprise-integration) now reside on the GitHub integration page.
* The [Snyk SCM Integrations](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations) page now contains information critical to using these integrations successfully in your SDLC. This includes:
  * Organizing integrations by [Group](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) and [Organization](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#organization-level-snyk-scm-integrations) level in line with Snyk AppRisk functionality
  * [Git repository cloning](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-git-repository-cloning) details
  * [Deployment recommendations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#deployment-order-recommendations) for Enterprise customers
  * [User permissions and access scope requirements](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#user-permissions-and-access-scope-requirements) for each SCM integration
  * Instructions on how to generate [integrated SCM tokens for Snyk Broker](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#integrated-scm-tokens-for-snyk-broker)

### **Other updates**

* **Snyk Reports:** The [issue column dictionary](manage-risk/reporting/issue-columns-dictionary.md#issue-vulnerability-details) includes new filters and columns for Jira (JIRA ISSUES LIST, LATEST JIRA ISSUE) and EPSS (EPSS SCORE, EPSS PERCENTILE). This allows you to manage your work with Jira and to include EPSS in your prioritization steps.
* **Snyk Security:** Snyk has improved the prioritization workflow and risk assessment by adopting [CVSS V4.0](manage-risk/prioritize-issues-for-fixing/severity-levels.md#severity-levels-and-cvss) as the default evaluation for new vulnerabilities.
* **Fix code vulnerabilities automatically:** [DeepCode AI Fix](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#deepcode-ai-fix-language-support) is now available in AWS Environments and JetBrains IDEs. If you use AWS multi-tenant environments, turn on the Snyk Preview [Snyk Code Fix Suggestions](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#enable-deepcode-ai-fix) and retest with Snyk in your IDE.









