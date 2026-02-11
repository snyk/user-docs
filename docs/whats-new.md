---
cover: .gitbook/assets/docs-banner-nov25.webp
coverY: 0
---

# What's new?

The most recent updates include significant changes to the user docs, such as features added or removed, structural changes that affect how you find relevant information, and other improvements to enhance your interaction with the Snyk knowledge base.

## January 2026

### Snyk Code

* The [Snyk Code Security policies](manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/#manage-ignores-at-the-group-level-through-snyk-code-security-policies) documentation has been updated to clarify that Snyk Code Security policies are different to Snyk Security Policies.
* The [Enable Snyk Agent Fix](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#enable-snyk-agent-fix) section has been enhanced with more details and clear configuration steps.

### Snyk CLI&#x20;

* [SBOM test](developer-tools/snyk-cli/commands/sbom-test.md) command is now in Early Access.
* The latest [Snyk CLI version](developer-tools/snyk-cli/install-or-update-the-snyk-cli/) available is v1.1302.1.
* The [debugging mode of the Snyk CLI](developer-tools/snyk-cli/debugging-the-snyk-cli.md) has been enhanced with error catalog codes and exit codes.

### Snyk IDE

* You can now use the Issue View Options in [Eclipse](developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects.md#issue-view-options), [JetBrains](developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#general-settings), and [Visual Studio Code ](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-configuration-environment-variables-and-proxy.md#scan-configuration)to filter issues by their Code Consistent Ignores status.
* The Risk Score Threshold option has been added to the [Visual Studio Code ](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-configuration-environment-variables-and-proxy.md#scan-configuration)extension to allow you to filter Open Source issues by their risk score.

### Snyk Studio

* The [Agentic security with Snyk Studio](integrations/snyk-studio-agentic-integrations/) documentation has been restructured to accurately reflect the workflow you would go through when using Snyk Studio for the first time, and using the available [Quickstart guides](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/).
* The [Snyk Studio Adoption](integrations/snyk-studio-agentic-integrations/usage-analytics.md) report is now available under [Redesigned analytics](manage-risk/analytics/overview-tab.md#snyk-studio-adoption).
* The [supported tools list](integrations/snyk-studio-agentic-integrations/#mcp-server-supported-tools) was updated to clarify that Snyk supports `snyk_send_feedback`.

### Snyk supported languages

* [PHP](supported-languages/supported-languages-list/php.md) has been updated to include support for PHP version 8.5.
* [JavaScript](supported-languages/supported-languages-list/javascript/) was updated with support for Yarn 4.
* [Ruby](supported-languages/supported-languages-list/ruby.md) was updated with support for Ruby 4.
* [Python](supported-languages/supported-languages-list/python/) has been updated to remove the limitation note for Projects with downloaded dependencies.
* [Go](supported-languages/supported-languages-list/go.md) has been updated to include support for Go standard library, for Go with Open Source.
* Several supported language pages and their rules have been updated with Code analysis support in Early Access:  [Rust](supported-languages/supported-languages-list/rust.md), [Swift and Objective-C](supported-languages/supported-languages-list/swift-and-objective-c.md), [Dart and Flutter](supported-languages/supported-languages-list/dart-and-flutter.md), [Groovy](supported-languages/supported-languages-list/groovy.md), [Rust rules](scan-with-snyk/snyk-code/snyk-code-security-rules/rust-rules.md), [Objective-C rules](scan-with-snyk/snyk-code/snyk-code-security-rules/objective-c-rules.md), [Dart and Flutter rules](scan-with-snyk/snyk-code/snyk-code-security-rules/dart-and-flutter-rules.md), [Groovy rules](scan-with-snyk/snyk-code/snyk-code-security-rules/groovy-rules.md).

### Other updates

* [Container monitor](developer-tools/snyk-cli/commands/container-monitor.md) has been updated to include a new CLI option `--prune-repeated-subdependencies` to prune depgraphs for large container scans.&#x20;
* The [redesigned Snyk Analytics experience](manage-risk/analytics/overview-tab.md) is now in General Availability.
* The [Bitbucket Cloud documentation](developer-tools/scm-integrations/organization-level-integrations/bitbucket-cloud.md) has been updated to make it clear that scopeless API tokens are not supported for this integration.

## December 2025

### **Snyk API**

* The API docs navigation was enhanced with additional package-related reference pages (including `ContainerRegistryImagePolicy`).&#x20;

### **Snyk integrations**

* The [Partner integrations](integrations/partner-integrations.md) page has been updated, including additional coverage for Coding Assistants and how they can use Snyk Studio (MCP) in agentic workflows.
* [JavaScript](supported-languages/supported-languages-list/javascript/) navigation has been enhanced with better redirect and routing features.&#x20;

### **Snyk Studio**

* The [Snyk Studio - Agentic integrations](integrations/snyk-studio-agentic-integrations/) documentation has been updated to provide a clearer explanation of MCP usage and the available Snyk Studio tools.
* The [Quickstart guides for Snyk Studio](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/) were updated with new and refreshed setup guidance, including [Cursor](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/cursor-guide.md) and [Windsurf](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/windsurf-guide.md).

### **Other updates**

* GitHub Cloud App and GitHub Server App have been added to the list of [supported SCMs for Dockerfile analysis](scan-with-snyk/snyk-container/scan-your-dockerfile/#supported-scms-for-dockerfile-analysis).
* The `snyk-scm-contributors-count` docs were updated with prerequisites and setup notes. See [snyk-scm-contributors-count](developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-scm-contributors-count/).
* The Declining Balance documentation from the [Snyk Declining Balance of Hours Service Description](snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/declining-balance.md) page was updated with service description and expiration details.&#x20;
* The IaC issue-reporting walkthrough from the 'Getting started with Snyk IaC' page was updated to remove outdated screenshots and copy.&#x20;
* PR template variables were updated on the [Variables list and description](scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/customize-pr-templates/variables-list-and-description.md) page with a new container base image `short name` values for cleaner PR titles and messages.&#x20;

## November 2025

### **Snyk Container**

* The list of [operated distribution systems supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated with support for Chisel.

### **Snyk CLI**

* The latest [Snyk CLI version](developer-tools/snyk-cli/install-or-update-the-snyk-cli/) available is v1.1301.0.
* The [CLI help](manage-risk/prioritize-issues-for-fixing/reachability-analysis.md#using-reachability-analysis-with-snyk-cli) has been updated with commands for reachability analysis.

### Snyk IDE

* The Automated Org Selection feature uses repository context to choose an Organization. Manual configuration overrides this automated selection. If the selection fails, Snyk defaults to your preferred Organization setting. The feature is available for the [Eclipse plugin](developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/configuration-of-the-eclipse-plugin.md), the [JetBrains plugin](developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md), the [Visual Studio extension](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-configuration-environment-variables-and-proxy.md), and the [Visual Studio Code extension](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-configuration-environment-variables-and-proxy.md).

### **Snyk integrations**

* The Amazon Q guide for Snyk Studio now includes [updated instructions](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/amazon-q-guide.md#install-the-snyk-mcp-server-in-the-amazon-q-ide-extension) for configuring the Snyk MCP Server in VS Code and JetBrains.

### **Other updates**

* [Reachabilty analysis](manage-risk/prioritize-issues-for-fixing/reachability-analysis.md) has been updated with instructions on how it works and how to use it in both the Snyk Web UI and the Snyk CLI and clear support for specific languages and package managers.
* The [Pre-defined roles](snyk-platform-administration/user-roles/pre-defined-roles.md#role-types) documentation has been updated to communicate that the Organization Admin role and associated permissions supersede any Group Member role restrictions.
* The [severity condition](manage-risk/policies/security-policies/security-policies-conditions.md) is now available in Group-level policies. Use this feature to create more granular policies for Snyk Code and Snyk Open Source findings, for example, ignoring a finding or changing its severity.

## October 2025

### **Snyk API**

* A new [API migration guide](snyk-api/api-end-of-life-eol-process-and-migration-guides/guides-to-migration/v1-reporting-apis-to-export-api-migration-guide.md) is available to help you migrate from the v1 Reporting API to the REST Exporting API.
* The Export API has been improved with the option to [limit the link expiration](snyk-api/using-specific-snyk-apis/export-api-specifications-columns-and-filters.md#data-retention).

### **Snyk Broker**

* The [Universal Broker](snyk-api/reference/universal-broker.md) release status has transitioned to Generally Available.
* The page [Upgrade an Organization from Classic Broker to Universal Broker](implementation-and-setup/enterprise-setup/snyk-broker/universal-broker/upgrade-an-organization-integration-from-classic-broker-to-universal-broker.md#migrating-multiple-organizations) has been updated with steps to migrate multiple Organizations at a time.

### **Snyk CLI**

* Snyk CLI now supports uploading files and folders for Snyk Code scanning. The command [`code-test`](developer-tools/snyk-cli/commands/code-test.md) has been updated with options reflecting these capabilities.
* The latest [Snyk CLI version](developer-tools/snyk-cli/install-or-update-the-snyk-cli/) available is v1.1300.2.

### **Snyk integrations**

* The list of Snyk MCP quick guides now includes [Devin guide](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/devin-guide.md), [Factory guide](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/factory-guide.md), [Factory terminal guide](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/factory-terminal-ide-guide.md).
* The Snyk MCP Server has been rebranded as [Snyk Studio](integrations/snyk-studio-agentic-integrations/).
* [SCM integration support for Python](supported-languages/supported-languages-list/python/scm-integrations-and-python.md) has been updated with support for Python 3.14.

### **Other updates**

* The [Operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md#minimus) has been updated to include include support for Minimus, Ubuntu 25.10 - Questing Quokka, and Ubuntu 25.04 - Plucky Puffin.
* For [Ruby](supported-languages/supported-languages-list/ruby.md), versions 2.3.X are no longer supported. The Ruby-specific versions have been updated to include more version patches.
* PR Check report was added as Early Access to the available reports to identify Snyk PR check locations, increase adoption, and pinpoint common failure impacts on developer workflows.
* You can now label your assets with metadata on repository assets and build artifacts, helping tag, manage security, and group items by features. An asset label differs from an asset tag, which enables key-value tags for structured metadata, allowing for granular filtering, policy creation, and improved system alignment.
* [JavaScript for open source](supported-languages/supported-languages-list/javascript/#javascript-for-snyk-open-source) has been updated to include full support for pnpm Projects.

## September 2025

### **Snyk Container**

* The instructions for [installing the Snyk Controller on Amazon Elastic Kubernetes Service (Amazon AKS)](scan-with-snyk/snyk-container/kubernetes-integration/install-the-snyk-controller/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md#create-an-eks-node-role-for-your-node-group-and-add-the-trust-relationship-for-the-iam-role) have been updated with details for configuring trust relationships for the IAM role.
* The list of [operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated to include SUSE Linux Enterprise Server 15.7 and Rocky Linux 10.

### **Snyk integrations**

* The SCM integration for Bitbucket Data Center/Server now supports the Required Builds feature for granular control over pull requests. To learn more, visit [Required Builds](developer-tools/scm-integrations/organization-level-integrations/bitbucket-data-center-server.md#required-builds).
* [GitLab](scan-with-snyk/pull-requests/pull-request-checks/analyze-pr-checks-results.md#gitlab) is supported for PR check results. This feature blocks merge requests with security issues when the "Pipelines must succeed" setting is enabled.
* The Snyk MCP quick guides list has been enriched with the following guides: [Claude Code](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/claude-code-guide.md), [Continue](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/continue-guide.md), [JetBrains AI Assistant](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/jetbrains-ai-assistant.md), [JetBrains Junie](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/jetbrains-junie.md)

### **Other updates**

* For Java and Kotlin, the list of [supported Gradle versions](supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-package-managers-and-package-registries) now includes Gradle 9.
* For [Ruby](supported-languages/supported-languages-list/ruby.md), an end-of-support notice has been added to say that starting Oct 1, 2025, Fix PRs are no longer supported for Projects using Ruby versions 3.1.x and lower. The table of supported Ruby versions has also been updated.
* For Javascript, [support for pnpm Projects](supported-languages/supported-languages-list/javascript/#support-for-pnpm) has been added.
* `Raise Support Community Cases` and `View Support Community Cases` Tenant level permissions have been added. To learn more about which Tenant roles these permissions apply to, visit Pre-defined roles, [Tenant-level permissions](snyk-platform-administration/user-roles/pre-defined-roles.md#tenant-level-permissions).
* The [Analytics](manage-risk/analytics/) menu now updates its data daily instead of hourly.
* Learn how to resolve duplicated and unenriched assets discovered outside Group and Organization-level SCM integrations.
* You can now [exclude specific values](/broken/pages/ESO9SlllNm4KuC5roygB) when you filter your reports.

## August 2025

### **Snyk API**

* The [Export API](snyk-api/reference/export.md) has been enhanced with the project\_target\_file field.
* A new dataset for usage events has been added to the [Export API.](snyk-api/reference/export.md)

### **Snyk CLI**

* [Experimental builds ](developer-tools/snyk-cli/releases-and-channels-for-the-snyk-cli.md#experimental-builds)information is now available for the CLI releases and channels.
* The [AI-BOM](developer-tools/snyk-cli/commands/aibom.md) Snyk CLI command is now available with any stable CLI release.
* A new Snyk CLI analytics page is now available, providing information about [Essential Operational Analytics](developer-tools/snyk-cli/snyk-cli-analytics.md#essential-operational-analytics) and [Optional Usage Analytics](developer-tools/snyk-cli/snyk-cli-analytics.md#optional-usage-analytics).

### **Snyk integrations**

* You can now add the Snyk MCP server to [Goose CLI](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/gemini-cli-guide-1.md) to secure code generated with agentic workflows through an LLM.
* You can now integrate Akamai with the Snyk API & Web to discover and scan your API. See the [API Security](integrations/partner-integrations.md#api-security) section under Partner integrations page for more details.
* The [Jira Cloud documentation](integrations/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration.md) has been updated for parity with the current version.

### **Other updates**

* A new [Risk exposure report](/broken/pages/Xet9BNiX2RviFlrTiatG#risk-exposure-report) has been released, providing you with a single, consolidated view of your security risks.
* The rollout to General Availability has started for the [Pull Request Experience](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md).
* The [Operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) now includes Debian 14 - Forky.
* Snyk now supports [Ruby versions](supported-languages/supported-languages-list/ruby.md#technical-specifications) 3.3 \[3.3.9] and 3.4 \[3.4.5]. If the Ruby version is not specified in the gemfile, it will default to version 3.1.

## July 2025

### **Snyk API**

* The [Export API](snyk-api/reference/export.md) is now available as GA.
* The Assets API is now available as Early Access.

### **Snyk CLI**

* MCP updates:
  * [Updated the list of supported Snyk security tools into an AI system](integrations/snyk-studio-agentic-integrations/#snyk-studio-tools).
  * Updated release status from experimental to [Early access](discover-snyk/getting-started/snyk-release-process.md#early-access-features) and removed the experimental flag.
  * Added [Cursor](integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/cursor-guide.md) as a new supported agentic IDE for MCP.
* PAT updates:
  * Added PAT support for [Snyk CLI](developer-tools/snyk-cli/authenticate-to-use-the-cli.md).
  * Added PAT support for Snyk CI/CD integrations ([CircleCI](developer-tools/snyk-ci-cd-integrations/circleci-integration-using-a-snyk-orb.md), [Jenkins](developer-tools/snyk-ci-cd-integrations/jenkins-plugin-integration-with-snyk.md), [Maven](developer-tools/snyk-ci-cd-integrations/maven-plugin-integration-with-snyk.md)).

### **Snyk Code**

* Support for Python, JavaScript and Typescript now includes more frameworks.

### **Snyk Container**

[Operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated to include: SUSE Linux Enterprise (SLE) 15.3+, Red Hat Enterprise Linux 10, and Oracle Linux 10.

### **Snyk IDE**

* Added PAT support for all [Snyk IDE](developer-tools/snyk-ide-plugins-and-extensions/) plugins and extensions.
* Added an [IDE Plugin Compatibility Matrix](developer-tools/snyk-ide-plugins-and-extensions/compatibility-matrix.md) for all supported versions.

### **Snyk integrations**

* [Snyk Agent Fix in the PR](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md#snyk-agent-fix-in-the-pr) has added support for Bitbucket integrations, still in Early Access.
* The [minimum version](scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md) of Bitbucket Server and Bitbucket Data Center required to use the integrations with PR checks has been updated to 7.4 and 8 respectively.

### **Snyk Open Source**

[Scan open-source libraries and licenses](scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/), [Snyk License Compliance Management](scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md), and [Fix your vulnerabilities](scan-with-snyk/snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities.md) have been updated with the new **Issues** tab layout.

### **Other updates**

* A new architecture for user documentation on developer tools is now available. This update groups the main developer tools into a single section and distinctly separates them from the integrations documentation.
* [Analytics](manage-risk/analytics/overview-tab.md) has a fresh new look.
* Added [Snyk Assist](discover-snyk/snyk-learn/snyk-assist.md) documentation.
* The [Developer IDE and CLI usage](/broken/pages/Xet9BNiX2RviFlrTiatG#developer-ide-and-cli-usage) report has been improved with MCP-related data to provide better visibility into MCP usage.
* [Okta custom mapping documentation](implementation-and-setup/enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/custom-mapping/examples-setting-up-custom-mapping-for-idps/example-setting-up-custom-mapping-for-okta.md#construct-a-value-expression-that-creates-a-roles-array-to-be-sent-to-snyk) has been updated to clarify handling of the `Arrays.flatten(appuser.snyk_orgs)` value during setup.

## June 2025

### **Snyk Broker**

* Updated the Snyk Broker documentation to include distinct steps for setting up the [Container Registry Agent with Docker](implementation-and-setup/enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/#configuring-and-running-the-container-registry-agent), whether using the Classic or Universal Broker.
* Updated the [Using the API to set up Universal Broker](implementation-and-setup/enterprise-setup/snyk-broker/universal-broker/using-the-api-to-set-up-universal-broker/) documentation with a Prerequisites section and clarified that the Snyk Broker App ID differs for each [region](snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-client-urls).
* Snyk Learn courses have been integrated into the [Universal Broker](enterprise-setup/snyk-broker/universal-broker/) pages.

### **Other updates**

* [Usage settings](snyk-platform-administration/groups-and-organizations/usage-settings.md) has been updated with the new **Billing and Usage** dashboard, available with the new Snyk Platform Access plan.
* [Snyk Platform Access credits](snyk-data-and-governance/snyk-platform-access-credits.md) has been added with brief information on the new Snyk Platform Access plan.
* The troubleshooting sections for all [Snyk IDE plugins](developer-tools/snyk-ide-plugins-and-extensions/), have been updated to include clear steps for working with the Logs details, which are available across all plugins.
* A new feature, the [Snyk Agent Fix in the PR](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md#snyk-agent-fix-in-the-pr), has been released, enabling the user to interact with inline comments by requesting an initial fix or a different suggestion, or by applying a specific fix by using the `@snyk /apply #` command.
* [Consistent Ignores](manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/) for Snyk Code now fully supports CLI Upload.
* The page on Docker Desktop Extension integration has been removed, due to the end of support.

## May 2025

### **Snyk CLI**

* The `--platform` option was added to the [`container sbom`](developer-tools/snyk-cli/commands/container-sbom.md) command.
* The MCP information was expanded to [Developer guardrails for agentic workflows](integrations/snyk-studio-agentic-integrations/).

### **IDE plugins and extensions**

* Information was added to the [JetBrains plugin troubleshooting](developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/troubleshooting-for-the-jetbrains-plugin.md).
* Region information was updated on all [IDE pages](developer-tools/snyk-ide-plugins-and-extensions/).

### **Snyk Code**

* Legacy ignores can be converted using [bulk ignore conversion](manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/convert-project-scoped-ignores-to-asset-scoped-ignores.md#bulk-ignore-conversion).
* DeepCode AI Fix has a new name: [Snyk Agent Fix](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md).

### **Snyk Container**

[Configure the integration with Docker Hub](scan-with-snyk/snyk-container/container-registry-integrations/integrate-with-docker-hub/configure-the-integration-with-docker-hub.md) has been updated to state that Snyk does not support Organization Access Tokens (OAT).

### **Snyk Integrations**

The [Bitbucket Cloud App](developer-tools/scm-integrations/organization-level-integrations/bitbucket-cloud-app.md) and [Jira App](integrations/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration.md) integrations are now available in the `SNYK-US-02` environment.

### **Other updates**

* For [SCM integrations with Python](supported-languages/supported-languages-list/python/scm-integrations-and-python.md), the list of dependencies that are not supported has been updated to include `pip` for Python 2.7 and 3.7.
* [Python dependency filtering results](supported-languages/supported-languages-list/python/scm-integrations-and-python.md) have been updated to clarify the conditions in which certain packages and configurations are skipped by SCM scans.
* The list of supported package managers has been updated to include `conan`. See [C/C++](supported-languages/supported-languages-list/c-c++/), [SBOM test](developer-tools/snyk-cli/commands/sbom-test.md), [Test an SBOM document for vulnerabilities](snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).
* [Instructions for upgrading an Organization integration from Classic Broker to Universal Broker](implementation-and-setup/enterprise-setup/snyk-broker/universal-broker/upgrade-an-organization-integration-from-classic-broker-to-universal-broker.md) were clarified.

## April 2025

### **Snyk API**

* Several APIs have been updated; see the [Changelog](snyk-api/changelog.md).
* The navigation in the API section now reflects the use of Authentication and the Changelog for both the V1 and REST APIs.
* The [Authentication for API](snyk-api/authentication-for-api/) page has been updated with region information and clarity on using the bearer token.
* The [API endpoints index and tips](snyk-api/api-endpoints-index-and-tips/) page now has a note about how to find your `org_id`.

### **Snyk Essentials**

* [The Inventory Overview tab](manage-assets/assets-inventory-layouts.md) is now available to provide insights and prescriptive guidance to improve your application security.
* [The Visibility column](manage-assets/assets-inventory-components.md#visibility) has been added to show the visibility status of your repositories.

### **Snyk Broker**

Additional updates have been made to the [Universal Broker](snyk-api/reference/universal-broker.md) documentation to clarify the instructions and add details about the use of the APIs.

### **Snyk CLI**

Information has been added about Snyk support for the Model Context Protocol (MCP) through the [`snyk mcp` experimental CLI command](integrations/snyk-studio-agentic-integrations/usage-analytics.md).

### **Snyk Code**

* [Consistent Ignores ](manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/)is now available in Early Access. Your development teams can create ignores that are consistently respected regardless of how and where the test is run and what branch is being tested.
* Snyk Code supports gRPC libraries.

### **Snyk Container**

* [Using Custom Base Image Recommendation](scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/) has been updated with clarifications on how Snyk recommends images.
* The list of [Operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated to include Alpine Linux 3.21, Ubuntu 25.04 - Plucky Puffin, and Ubuntu 24.10 - Oracular Oriole.
* The section describing the automated integration process for Amazon Elastic Container Registry (ECR) has been removed, as Snyk no longer supports this method.

### **Snyk Integrations**

* For the [Jira integration](integrations/jira-and-slack-integrations/jira-integration.md#prerequisites-for-jira-integration-with-snyk), Snyk now supports Jira versions 5 to 10.
* For [SCM integrations with Gradle](supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md), Snyk now supports `allprojects` and `subprojects` blocks, as well as Spring Boot plugins BOMs.

### **Other updates**

* DAST scanning is now available with [Snyk API & Web](scan-with-snyk/overview.md#select-scanning-methods), enabling users to discover and test the security of their APIs and web apps, including AI-generated ones.
* PR Checks is now available with a General Availability release status.

## March 2025

### **Snyk Broker**

* The Snyk Broker section has been divided into [Universal Broker](enterprise-setup/snyk-broker/universal-broker/) and [Classic Broker](implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/) documentation and the [main page](implementation-and-setup/enterprise-setup/snyk-broker/) has been updated.
* The Classic Broker installation instructions now include the command to set the `BROKER_SERVER_URL` for [Docker](enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md) and the `brokerServerUrl` for [Helm](enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md).

### **Snyk API**

* The [V1 API overview](snyk-api/v1-api.md) and [reference](snyk-api/reference/) are now on the user docs site only. Additional details from Apiary have been added to the V1 reference on the user docs site. The API reference has been removed from the V1 API Apiary site.
* A section has been added for [pages that explain how to use specific APIs in depth](snyk-api/using-specific-snyk-apis/).

### **Snyk CLI, CI/CD, IDE**

* [Advanced use of Snyk Container CLI](developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/advanced-use-of-snyk-container-cli.md) now includes support for scanning Kaniko image archives.
* The [support policy for the CI/CD plugins](developer-tools/snyk-ci-cd-integrations/#support-policy) was updated to align with the CLI support policy.
* The Net new issues feature was added to the IDE documentation for [Eclipse](developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects.md#net-new-issues-versus-all-issues), [JetBrains](developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/run-an-analysis-with-the-jetbrains-plugin.md#net-new-issues-versus-all-issues), [Visual Studio](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/view-analysis-results-from-visual-studio-extension.md#net-new-issues-versus-all-issues), and [Visual Studio Code](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/view-analysis-results-from-visual-studio-code-extension/#net-new-issues-versus-all-issues), and [troubleshooting information](developer-tools/snyk-ide-plugins-and-extensions/troubleshooting-ides/net-new-issues-delta-scan-troubleshooting.md) was added.

### **Snyk Code**

* The Generated Pull Requests report is now available in Early Access. This report provides an overview of how Fix, Backlog, and Upgrade PRs are used and highlights the efficiency of PR merges.
* [The Pull Request Experience](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md) now supports GitLab and Azure Repos SCM integrations, with a few [limitations](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md#inline-comments).
* New Snyk Code filters and columns were added to [Snyk Reports](manage-risk/analytics/reports-tab/issue-columns-dictionary.md#issue-characteristics) and [Snowflake Data Share](manage-risk/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary.md): File Path, Code Region, and Asset Finding ID.
* Snyk Code now supports [Rust](supported-languages/supported-languages-list/rust.md) and [Groovy](supported-languages/supported-languages-list/groovy.md) available in Early Access and accessible from Snyk Preview.

### Snyk Essentials

* A new feature is now available in Snyk Essentials, introducing a new type of [asset tag](manage-risk/policies/assets-policies/#asset-tagging) known as GitHub custom properties.
* [Asset tags](manage-assets/assets-inventory-components.md#tags) have been redefined and are now clearly separated into system tags and user-defined tags.

### **Snyk Integrations**

* The [GitHub Server App](developer-tools/scm-integrations/organization-level-integrations/github-server-app.md) has moved into General Availability.
* The [Jira integration documentation](integrations/jira-and-slack-integrations/jira-integration.md#prerequisites-for-jira-integration-with-snyk) has been updated to state that Snyk supports version 5 to version 9.

### **Other updates**

* The PCI-DSS v4.0.1 report is now available in Early Access. This report leverages Snyk scan results to assess, prove, and improve readiness for PCI-DSS AppSec compliance regarding SCA and SAST vulnerabilities.
* The [Repositories Tested in CI/CD report](/broken/pages/Xet9BNiX2RviFlrTiatG#repositories-tested-in-ci-cd-report) is available in Early Access. This report tracks Snyk CI/CD testing to prevent vulnerable production deployments.
* [Severity levels](manage-risk/prioritize-issues-for-fixing/severity-levels.md#why-are-there-multiple-cvss-scores-for-the-same-vulnerability) now provide more details about the CVSS v4.0.

## February 2025

### Snyk Essentials

* The Integrations UI at the Group level has been enhanced to improve readability and actionability and provide inline instructions and inline profile helpers.
* Group-level [Integrations documentation](integrations/integrate-with-snyk.md#integrations-syncing-time) has been updated with new, more accurate sync times.
* The [asset filter](manage-risk/policies/assets-policies/create-policies.md) documentation has been consolidated into one section, and it now links to all relevant areas, such as Inventory and Asset Policy filters.
* [Issues Analytics](/broken/pages/1m1GoANo0CromJXOD7H1) is now available with a General Availability release status.

### Other updates

* A new [Automated Provisioning guide](implementation-and-setup/enterprise-setup/auto-provisioning-guide.md) has been created for **Pilot** and **Enterprise** **users**, detailing the steps of the auto-provisioning process for new and existing user accounts.
* [Snyk Code PR Checks](scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md#configure-for-code-analysis-click-to-expand) are in General Availability.

## December 2024 and January 2025

### **Snyk Container**

* Page "Integrate with Docker Desktop Extension" has been updated to include an end-of-support notice. Effective June 20, 2025, the integration with Docker Desktop will no longer receive updates or technical support.

### **Snyk CLI and IDEs**

* [Eclipse IDE](developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/) major update
* [Visual Studio IDE](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/) major update
* Region configuration update for [IDEs](developer-tools/snyk-ide-plugins-and-extensions/)
* [Snyk images EOL policy updated](developer-tools/snyk-ci-cd-integrations/snyk-images-and-eol-image-policy.md)
* [`snyk container test`](developer-tools/snyk-cli/commands/container-test.md) and [`snyk container monitor`](developer-tools/snyk-cli/commands/container-monitor.md) option `--exclude-node-modules` added

### **Other updates**

* [Snyk Admin](snyk-platform-administration/snyk-admin.md) pages have been updated to reflect the addition of [Tenants](snyk-platform-administration/groups-and-organizations/tenant/) in the Snyk hierarchy, including a new infographic to illustrate the Tenant position in the [hierarchy](snyk-platform-administration/groups-and-organizations/#the-snyk-hierarchy).

## November 2024

### **Snyk Essentials**

### **Snyk Container**

* The list of [operating system distributions supported by Snyk Container](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) has been updated to include Ubuntu 24.10 - Oracular Oriole and Ubuntu 24.04 - Noble Numbat 04.
* [How Snyk Container works](scan-with-snyk/snyk-container/how-snyk-container-works/) has been updated with details on the logic Snyk applies when providing public base image recommendations.

### **Other updates**

* The Pull Request Checks section has been updated to include the new [Pull Request Experience](scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md) for PR Checks.
* The [Supported languages](supported-languages/supported-languages-package-managers-and-frameworks.md) page has been reorganized to provide detailed information about language availability for each Snyk product. Additionally, it provides a list of package managers, frameworks, and features for each supported language.
* A service account using OAuth 2.0 can now be [created through the Snyk Web UI](implementation-and-setup/enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md#create-oauth-service-accounts-through-the-ui).
* The [API index ](snyk-api/api-endpoints-index-and-tips/)now includes entries for each endpoint mentioned in the Snyk user docs.
* The[ Developer IDE and CLI usage](/broken/pages/Xet9BNiX2RviFlrTiatG#developer-ide-and-cli-usage) report has been enhanced with additional functionalities: **Developer email address** and **PDF export**.
* The [Vulnerabilities Detail](/broken/pages/Xet9BNiX2RviFlrTiatG#vulnerabilities-detail-report) report has been enhanced with additional functionalities, such as **Target indication** and **Column picker**.

## October 2024

### **Snyk API**

* [Asset inventory components](manage-assets/assets-inventory-components.md#clusters) has been updated to include details on clusters.

### **Snyk CLI and IDEs**

* The [CLI authentication page](developer-tools/snyk-cli/authenticate-to-use-the-cli.md) has been updated for the OAuth 2.0 protocol.
* The page [Debugging the Snyk CLI](developer-tools/snyk-cli/debugging-the-snyk-cli.md) has been added.
* [CLI standalone executables](developer-tools/snyk-cli/install-or-update-the-snyk-cli/#install-with-standalone-executables) have been updated to include Alpine Arm64.
* IDE Eclipse[ plugin](developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/) and [JetBrains plugin ](developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/)documentation pages have been updated.
* [Authentication information](developer-tools/snyk-ide-plugins-and-extensions/) has been updated for all IDEs.

### **Snyk Integrations**

* [Snowflake Data Share](manage-risk/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share/) is now in [GA](discover-snyk/getting-started/snyk-release-process.md).
* [Snyk SCM integrations](developer-tools/scm-integrations/organization-level-integrations/) has been updated with additional notices relating to repository retrieval and permission or scope modifications after initial configuration.
* GitHub Cloud App has been added to feature support notices for Fix, Backlog, and Upgrade PRs.
* Snyk SCM integrations has been updated to include a table detailing the [permissions and scopes](developer-tools/scm-integrations/user-permissions-and-access-scopes.md#github-cloud-app-permission-requirements) required for the GitHub Cloud App.

### **Other updates**

* [Getting started](discover-snyk/getting-started/) has been updated to centralize content related to everything you need to know before using Snyk.
* Scanning methods have been added for the [Dart and Flutter](supported-languages/supported-languages-list/dart-and-flutter.md) languages.

## September 2024

### Snyk API

* A prerequisites section has been added to the Group level of [GitHub integration](developer-tools/scm-integrations/organization-level-integrations/github-enterprise.md#prerequisites), and more details about the [Pull personal repositories](developer-tools/scm-integrations/group-level-integrations/github-for-snyk-essentials.md) option have been added to the same documentation page.
* The [Set up Insights](manage-risk/prioritize-issues-for-fixing/set-up-insights/) section was updated to emphasize the risk factors availability for each integration option.
* The Snyk Runtime Sensor has been updated to reflect the importance of adopting it to achieve the most effective integration and to access its continuously expanded set of features.

### Snyk Broker

The Universal Broker feature is now available in Early Access. The Universal Broker separates deployment and container concerns from connection concerns. It allows for a smaller or a single deployment to support numerous connections of varied types.

### Snyk CLI

* The [CLI commands and options summary](developer-tools/snyk-cli/cli-commands-and-options-summary.md) was updated.
* [Authentication](developer-tools/snyk-cli/authenticate-to-use-the-cli.md) has been updated.
* Configuration has been updated: Environment variables for Snyk CLI, [`snyk config`](developer-tools/snyk-cli/commands/config.md) help, [`snyk config environment`](developer-tools/snyk-cli/commands/config-environment.md) help.

### Snyk Integrations

The Snowflake Data Share section has been updated to include a [Data Share Dictionary](manage-risk/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary.md), designed to help you navigate and build your dataset.

### Other updates

* The updated [Regional hosting and data residency](snyk-data-and-governance/regional-hosting-and-data-residency.md) page was published.
* [Glossary](discover-snyk/getting-started/glossary.md) terms were updated for SCA, SAST, DAST, and IAST as well as Software Composition Analysis.
* [Early Access](discover-snyk/getting-started/snyk-release-process.md#early-access) release status notices were updated.

## August 2024

### Snyk API

* Links in the API reference docs have been updated.
* The [API endpoints index and notes](snyk-api/api-endpoints-index-and-tips/) have been updated.

### Snyk Essentials

### Snyk CLI

* [`snyk auth`](developer-tools/snyk-cli/commands/auth.md) command help updated to reflect OAuth default.
* [CLI authentication](developer-tools/snyk-cli/authenticate-to-use-the-cli.md) instructions updated for OAuth default and improved flow.
* [`snyk config environment`](developer-tools/snyk-cli/commands/config-environment.md) command help has been added.
* CLI [support for pnpm added](supported-languages/supported-languages-list/javascript/#support-for-pnpm).

### Snyk IDE

* [CLI authentication](developer-tools/snyk-cli/authenticate-to-use-the-cli.md) instructions updated for IDE.
* IDE authentication instructions updated: [Eclipse](developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/authentication-for-the-eclipse-plugin.md), [Jetbrains](developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/authentication-for-the-jetbrains-plugins.md), [VS extension](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/authentication-for-visual-studio-extension.md), [VS Code extension](developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/authentication-for-visual-studio-code-extension.md)

### **Snyk Integrations**

* Git repository cloning has been renamed [Workspaces for SCM integrations](developer-tools/scm-integrations/workspaces.md) to better reflect its functionality. Additional detail on [enablement](developer-tools/scm-integrations/workspaces.md#manage-workspaces) has been added.
* The [relationship](developer-tools/scm-integrations/organization-level-integrations/github-cloud-app.md#how-to-set-up-the-github-cloud-app) between GitHub organizations and Snyk Organizations when integrating with the GitHub Cloud App has been clarified.

## July 2024

### **Snyk API**

* The API documentation now provides the API Reference and explanatory documentation in the [API section](snyk-api/snyk-api.md).
* The [API End of Life (EOL) process and migration guides](snyk-api/api-end-of-life-eol-process-and-migration-guides/) are now published and updated to support the process, which began in July.
* [Asset inventory filtering](manage-assets/assets-inventory-components.md#asset-tabs) describes the new, simplified view that provides an improved experience of filtering the assets.
* The [Asset inventory layouts](manage-assets/assets-inventory-layouts.md) have been renamed to better reflect their functionality.
* Four new SCM integrations are now available for Snyk:
  * [Atlassian Compass](developer-tools/scm-integrations/application-context-for-scm-integrations/#atlassian-compass)
  * [Harness](developer-tools/scm-integrations/application-context-for-scm-integrations/#harness)
  * [OpsLevel](developer-tools/scm-integrations/application-context-for-scm-integrations/#opslevel)
  * [Datadog Service Catalog](developer-tools/scm-integrations/application-context-for-scm-integrations/#datadog-service-catalog)

### Snyk Integrations

* A comparison of the GitHub and GitHub Enterprise integrations functions now resides on the [SCM, IDE, and CI/CD integrations](developer-tools/scm-integrations/#github-vs-github-enterprise) page.
* Steps for [migrating from the GitHub integration to the GitHub Enterprise integration](developer-tools/scm-integrations/organization-level-integrations/github.md#migrate-to-the-github-enterprise-integration) now reside on the GitHub integration page.
* The [Snyk SCM Integrations](developer-tools/scm-integrations/organization-level-integrations/) page now contains information critical to using these integrations successfully in your SDLC. This includes:
  * [Git repository cloning](developer-tools/scm-integrations/workspaces.md) details
  * [Deployment recommendations](developer-tools/scm-integrations/deployment-recommendations.md) for Enterprise customers
  * [User permissions and access scope requirements](developer-tools/scm-integrations/user-permissions-and-access-scopes.md) for each SCM integration
  * Instructions on how to generate [integrated SCM tokens for Snyk Broker](developer-tools/scm-integrations/scm-integrations-and-snyk-broker.md#integrated-scm-tokens-for-classic-broker)

### **Other updates**

* **Snyk Reports:** The [issue column dictionary](manage-risk/analytics/reports-tab/issue-columns-dictionary.md#issue-vulnerability-details) includes new filters and columns for Jira (JIRA ISSUES LIST, LATEST JIRA ISSUE) and EPSS (EPSS SCORE, EPSS PERCENTILE). This allows you to manage your work with Jira and to include EPSS in your prioritization steps.
* **Snyk Security:** Snyk has improved the prioritization workflow and risk assessment by adopting [CVSS V4.0](manage-risk/prioritize-issues-for-fixing/severity-levels.md#severity-levels-and-cvss) as the default evaluation for new vulnerabilities.
* **Fix code vulnerabilities automatically:** [DeepCode AI Fix](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#snyk-agent-fix-language-support) is now available in AWS Environments and JetBrains IDEs. If you use AWS multi-tenant environments, turn on the Snyk Preview [Snyk Code Fix Suggestions](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#enable-snyk-agent-fix) and retest with Snyk in your IDE.
