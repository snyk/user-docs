---
cover: .gitbook/assets/Snyk General Banner.webp
coverY: 0
---

# What is new?

The most recent updates include significant changes to the user docs, such as features added or removed, structure changes that affect how you find relevant information, and other improvements aimed at enhancing your interaction with the Snyk knowledge base.

## April 2025

### **Snyk API**&#x20;

* Several APIs have been updated; see the [Changelog](snyk-api/changelog.md).
* The navigation of the API section now reflects the use of Authentication and the Changelog for both the V1 and REST APIs.
* The [Authentication for API](snyk-api/authentication-for-api/) page has been updated with region information and clarity on using the bearer token.
* The [API endpoints index and tips](snyk-api/api-endpoints-index-and-tips/) page now has a note about how to find your `org_id`.

### **Snyk AppRisk**&#x20;

* [The Inventory Overview tab](manage-assets/assets-inventory-layouts.md) is now available to provide insights and prescriptive guidance to improve your application security.
* [The Visibility column](manage-assets/assets-inventory-components.md#visibility) has been added to show the visibility status of your repositories.

### **Snyk Broker**

Additional updates have been made to the [Universal Broker](snyk-api/reference/universal-broker.md) documentation to clarify the instructions and add details about the use of the APIs.

### **Snyk CLI**

Information has been added about Snyk support for the Model Context Protocol (MCP) through the [`snyk mcp` experimental CLI command](snyk-cli/snyk-mcp-experimental.md).

### **Snyk Code**

* [Consistent Ignores ](manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code-early-access/)is now available in Early Access. Your development teams can create ignores that are consistently respected regardless of how and where the test is run and what branch is being tested.&#x20;
* Snyk Code supports gRPC libraries.

### **Snyk Container**

* [Using Custom Base Image Recommendation ](scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/)has been updated with clarifications on how Snyk recommends images.
* The list of [Operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated to include Alpine Linux 3.21, Ubuntu 25.04 - Plucky Puffin, and Ubuntu 24.10 - Oracular Oriole.
* The section describing the automated integration process for Amazon Elastic Container Registry (ECR) has been removed, as Snyk no longer supports this method.

### **Snyk Integrations**

* For the [Jira integration](integrate-with-snyk/jira-and-slack-integrations/jira-integration.md#prerequisites-for-jira-integration-with-snyk), Snyk now supports Jira versions 5 to 10.
* For [SCM integrations with Gradle](supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md), Snyk now supports `allprojects` and `subprojects` blocks, as well as Spring Boot plugins BOMs.

### **Other updates**

* DAST scanning is now available with [Snyk API & Web](https://docs.snyk.io/scan-with-snyk#select-scanning-methods), enabling users to discover and test the security of their APIs and web apps, including AI-generated ones.
* PR Checks is now available with a General Availability release status.

## March 2025

### **Snyk Broker**

* The Snyk Broker section has been divided into [Universal Broker](enterprise-setup/snyk-broker/universal-broker/) and [Classic Broker](enterprise-setup/snyk-broker/classic-broker/) documentation and the [main page](enterprise-setup/snyk-broker/) has been updated.
* The Classic Broker installation instructions now include the command to set the `BROKER_SERVER_URL` for [Docker](enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md) and the `brokerServerUrl` for [Helm](enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md).

### **Snyk API**

* The [V1 API overview](snyk-api/v1-api.md) and [reference](snyk-api/reference/) are now on the user docs site only. Additional details from Apiary have been added to the V1 reference on the user docs site. The API reference has been removed from the V1 API Apiary site.
* A section has been added for [pages that explain how to use specific APIs in depth](snyk-api/using-specific-snyk-apis/).

### **Snyk CLI, CI/CD, IDE**

* [Advanced use of Snyk Container CLI](snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/advanced-use-of-snyk-container-cli.md) now includes support for scanning Kaniko image archives.
* The [support policy for the CI/CD plugins](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations#support-policy) was updated to align with the CLI support policy.
* The Net new issues feature was added to the IDE documentation for [Eclipse](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects.md#net-new-issues-versus-all-issues), [JetBrains](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/run-an-analysis-with-the-jetbrains-plugin.md#net-new-issues-versus-all-issues), [Visual Studio](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/view-analysis-results-from-visual-studio-extension.md#net-new-issues-versus-all-issues), and [Visual Studio Code](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/view-analysis-results-from-visual-studio-code-extension/#net-new-issues-versus-all-issues), and [troubleshooting information](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/troubleshooting-ides/net-new-issues-delta-scan-troubleshooting.md) was added.

### **Snyk Code**

* The Generated Pull Requests report is now available in Early Access. This report provides an overview of how Fix, Backlog, and Upgrade PRs are used and highlights the efficiency of PR merges.
* The Pull Request Experience now supports GitLab and Azure Repos SCM integrations, with a few [limitations](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md#gitlab-and-azure-repos).
* New Snyk Code filters and columns were added to [Snyk Reports](manage-risk/reporting/issue-columns-dictionary.md#issue-characteristics) and [Snowflake Data Share](manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary.md): File Path,  Code Region, and Asset Finding ID.
* Snyk Code now supports [Rust](supported-languages-package-managers-and-frameworks/rust.md) and [Groovy](supported-languages-package-managers-and-frameworks/groovy.md) available in Early Access and accessible from Snyk Preview.

### Snyk Essentials

* A new feature is now available in Snyk Essentials, introducing a new type of [asset tag](manage-risk/policies/assets-policies/#asset-tagging) known as GitHub custom properties.
* [Asset tags](manage-assets/assets-inventory-components.md#tags) have been redefined and are now clearly separated into system tags and user-defined tags.

### **Snyk  Integrations**

* The [GitHub Server App](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-server-app.md) has moved into General Availability.
* The [Jira integration documentation](integrate-with-snyk/jira-and-slack-integrations/jira-integration.md#prerequisites-for-jira-integration-with-snyk) has been updated to state that Snyk supports version 5 to version 9.

### **Other updates**

* The [PCI-DSS v4.0.1 report](manage-issues/reporting/available-snyk-reports.md#pci-dss-v4.0.1-report) is now available in Early Access. This report leverages Snyk scan results to assess, prove, and improve readiness for PCI-DSS AppSec compliance regarding SCA and SAST vulnerabilities.
* The [Repositories Tested in CI/CD report](manage-issues/reporting/available-snyk-reports.md#repositories-tested-in-ci-cd-report) is available in Early Access. This report tracks Snyk CI/CD testing to prevent vulnerable production deployments.
* [Severity levels](manage-risk/prioritize-issues-for-fixing/severity-levels.md#why-are-there-multiple-cvss-scores-for-the-same-vulnerability) now provide more details about the CVSS v4.0.

## February 2025

### Snyk AppRisk

* A new third-party integration, the [Google Cloud Security Command Center](integrate-with-snyk/event-forwarding/google-security-command-center.md) (SCC) integration, is now available in Early Access. This integration sends Snyk issues to SCC, enabling you to view and manage Snyk issues alongside cloud security findings from Google Cloud in a single pane of glass.
* [Snyk Runtime Sensor](manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md) has been updated and now clearly separates the DaemonSet and Deployment helm installation steps.

### Snyk Essentials

* The Integrations UI at the Group level has been enhanced to improve readability and actionability and provide inline instructions and inline profile helpers.
* Group-level [Integrations documentation](integrate-with-snyk/#integrations-syncing-time) has been updated with new, more accurate sync times.
* The [asset filter](manage-risk/policies/assets-policies/create-policies.md) documentation has been consolidated into one section, and it now links to all relevant areas, such as Inventory and Asset Policy filters.
* [Issues Analytics](manage-risk/enterprise-analytics/issues-analytics.md) is now available with a General Availability release status.

### Other updates

* A new [Automated Provisioning guide](enterprise-setup/auto-provisioning-guide.md) has been created for **Pilot** and **Enterprise** **users**, detailing the steps of the auto-provisioning process for new and existing user accounts.
* [Snyk Code PR Checks](scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md#configure-for-code-analysis-click-to-expand) are in General Availability.

## December 2024 and January 2025

### **Snyk Container**

* [Integrate with Docker Desktop Extension](scan-with-snyk/snyk-container/container-registry-integrations/integrate-with-docker-desktop-extension.md) has been updated to include an end-of-support notice. Effective June 20, 2025, the integration with Docker Desktop will no longer receive updates or technical support.

### **Snyk CLI and IDEs**

* [Eclipse IDE](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/) major update
* [Visual Studio IDE](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/) major update
* Region configuration update for [IDEs](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)
* [Snyk images EOL policy updated](scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-images-and-eol-image-policy.md)
* [`snyk container test`](snyk-cli/commands/container-test.md) and [`snyk container monitor`](snyk-cli/commands/container-monitor.md) option `--exclude-node-modules` added

### **Other updates**

* [Snyk Admin](https://docs.snyk.io/snyk-admin) pages have been updated to reflect the addition of [Tenants](https://docs.snyk.io/snyk-admin/groups-and-organizations/tenant) in the Snyk hierarchy, including a new infographic to illustrate the Tenant position in the [hierarchy](https://docs.snyk.io/snyk-admin/groups-and-organizations#the-snyk-hierarchy).

## November 2024

### &#x20;**Snyk AppRisk**&#x20;

* The [Asset Dashboard](manage-issues/reporting/available-snyk-reports.md#asset-dashboard) has been redesigned and is included in the Reports section. It now features several enhancements, including a Global filter bar, new data widgets, and the option to export the dashboard as a PDF.
* The [Snyk Broker - AppRisk](enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md#scm-integrations) documentation has been updated to outline the specific steps necessary to configure each SCM integration with a Snyk Broker token.

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
* The [API Authentication](snyk-api/authentication-for-api/) pages have been updated to include when to use a personal token and when to use a service account, and to improve the flow of information.

### **Snyk AppRisk**

* [Risk factors on assets](manage-assets/#inventory-overview) is now in [EA](getting-started/snyk-release-process.md#early-access) for Snyk AppRisk customers.
* [Asset inventory components](manage-assets/assets-inventory-components.md#clusters) has been updated to include details on clusters.
* Snyk Broker - AppRisk has been updated to include [SonarQube installation steps](enterprise-setup/snyk-broker/using-snyk-apprisk-with-snyk-broker.md#sonarqube-sast-integration).

### **Snyk CLI and IDEs**

* The [CLI authentication page](snyk-cli/authenticate-to-use-the-cli.md) has been updated for the OAuth 2.0 protocol.
* The page [Debugging the Snyk CLI](snyk-cli/debugging-the-snyk-cli.md) has been added.
* [CLI standalone executables](snyk-cli/install-or-update-the-snyk-cli/#install-with-standalone-executables) have been updated to include Alpine Arm64.
* IDE Eclipse[ plugin](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/) and [JetBrains plugin](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/) documentation pages have been updated.
* [Authentication information](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/) have been updated for all IDEs.

### **Snyk Integrations**

* [Snowflake Data Share](manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/) is now in [GA](getting-started/snyk-release-process.md).
* [Snyk SCM integrations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/) has been updated with additional notices relating to repository retrieval and permission or scope modifications after initial configuration.
* GitHub Cloud App has been added to feature support notices for Fix, Backlog, and Upgrade PRs.
* Snyk SCM integrations has been updated to include a table detailing the [permissions and scopes](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#github-cloud-app-permission-requirements) required for the GitHub Cloud App.

### **Other updates**

* [Getting started](getting-started/) has been updated to centralize content related to everything you need to know before using Snyk.&#x20;
* The [Snyk Jumpstart Services Description](working-with-snyk/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-services-description.md) and [Customer Prerequisites](working-with-snyk/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites.md) have been updated for Snyk AppRisk.
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
  * [Application Analytics](manage-risk/enterprise-analytics/application-analytics.md) - available only for Snyk AppRisk customers.
* The [Manage Assets](manage-assets/) documentation has been updated to reflect the addition of Quick filters. Users are only shown quick filters relevant to their entitlement.

### **Snyk CLI**

* [`snyk auth`](snyk-cli/commands/auth.md) command help updated to reflect OAuth default.
* [CLI authentication](snyk-cli/authenticate-to-use-the-cli.md) instructions updated for OAuth default and improved flow.
* [`snyk config environment`](snyk-cli/commands/config-environment.md) command help has been added.
* CLI [support for pnpm added](supported-languages-package-managers-and-frameworks/javascript/javascript-for-open-source.md#pnpm).

### Snyk IDE

* [CLI authentication](snyk-cli/authenticate-to-use-the-cli.md) instructions updated for IDE.
* IDE authentication instructions updated: [Eclipse](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/authentication-for-the-eclipse-plugin.md), [Jetbrains](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/authentication-for-the-jetbrains-plugin.md), [VS extension](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-authentication.md), [VS Code extension](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-authentication.md)

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
  * Integrations - [Snyk SCM integrations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-essentials-scm-integrations) and [Integrate with Snyk](integrate-with-snyk/) sections
  * Snyk AppRisk - [Scan with Snyk](scan-with-snyk/snyk-apprisk.md) section
  * Analytics - [Manage risk](manage-risk/enterprise-analytics/application-analytics.md) section

### Snyk Integrations

* A comparison of the GitHub and GitHub Enterprise integrations functions now resides on the [SCM, IDE, and CI/CD integrations](scm-ide-and-ci-cd-integrations/#github-vs-github-enterprise) page.
* Steps for [migrating from the GitHub integration to the GitHub Enterprise integration](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#migrate-to-the-github-enterprise-integration) now reside on the GitHub integration page.
* The [Snyk SCM Integrations](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations) page now contains information critical to using these integrations successfully in your SDLC. This includes:
  * Organizing integrations by [Group](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-essentials-scm-integrations) and [Organization](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#organization-level-snyk-scm-integrations) level in line with Snyk AppRisk functionality
  * [Git repository cloning](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-git-repository-cloning) details
  * [Deployment recommendations](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#deployment-order-recommendations) for Enterprise customers
  * [User permissions and access scope requirements](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#user-permissions-and-access-scope-requirements) for each SCM integration
  * Instructions on how to generate [integrated SCM tokens for Snyk Broker](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#integrated-scm-tokens-for-snyk-broker)

### **Other updates**

* **Snyk Reports:** The [issue column dictionary](manage-risk/reporting/issue-columns-dictionary.md#issue-vulnerability-details) includes new filters and columns for Jira (JIRA ISSUES LIST, LATEST JIRA ISSUE) and EPSS (EPSS SCORE, EPSS PERCENTILE). This allows you to manage your work with Jira and to include EPSS in your prioritization steps.
* **Snyk Security:** Snyk has improved the prioritization workflow and risk assessment by adopting [CVSS V4.0](manage-risk/prioritize-issues-for-fixing/severity-levels.md#severity-levels-and-cvss) as the default evaluation for new vulnerabilities.
* **Fix code vulnerabilities automatically:** [DeepCode AI Fix](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#deepcode-ai-fix-language-support) is now available in AWS Environments and JetBrains IDEs. If you use AWS multi-tenant environments, turn on the Snyk Preview [Snyk Code Fix Suggestions](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#enable-deepcode-ai-fix) and retest with Snyk in your IDE.









