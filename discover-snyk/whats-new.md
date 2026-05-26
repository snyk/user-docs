---
cover: .gitbook/assets/docs-banner-nov25.webp
coverY: 0
---

# What's new?

The most recent updates include significant changes to the user docs, such as features added or removed, structural changes that affect how you find relevant information, and other improvements to enhance your interaction with the Snyk knowledge base.

## April 2026

### Snyk CLI

* The [Install the Snyk CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/install-the-snyk-cli) page and sub-pages have been updated to reflect configuration best practice, including easily copied code snippets.
* The [Authenticate to use the Snyk CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/authenticate-to-use-the-cli) page has been updated to reflect best practice, including easily copied code snippets.

### Snyk supported languages

* [CLI support for uv](supported-languages/supported-languages-list/python/cli-support-for-uv.md) is now in Early Access.
* Snyk now supports [interfile analysis for Ruby](supported-languages/supported-languages-list/ruby.md#ruby-for-snyk-code).

### Other updates

* The `snyk_package_health_check` directive is now available for the Full profile on the [Directives](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/directives#secure-at-inception-package-health-check-experimental) page.
* CISA KEV has been added to the list of filters available in [Issue vulnerability details](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/issue-columns-dictionary#issue-vulnerability-details).
* The [PR Checks Report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/prevention-reports#pull-request-checks-usage-and-performance-report) is now General Available, with updates to Prevention Reports, Export API, and Snowflake Data Share.
* The [Pull Request experience](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience#pull-request-experience-feature-requirements) documentation has been updated to reflect that if you are using inline comments or Agent Fix, you must now specify a dedicated GitHub account by providing a GitHub Personal Access Token (PAT) in your integration settings.
* The [Enterprise implementation guide](implementation-and-setup/enterprise-implementation-guide/) now has embedded video tutorials to guide you in your Enterprise setup as a new user of Snyk.
* The [High availability mode](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/high-availability-mode) from Snyk Broker is now enabled by default.
* The [Container registry sync](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/use-snyk-container/sync-your-container-registry) from Snyk Container is now Generally Available.
* The [Container registry import policy](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/containerregistryimportpolicy) API was enhanced by refactoring schema names, adding test components, and full CRUD operations.

## March 2026

### Evo by Snyk

* Added the [`aibom test`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/aibom-test) command under [Snyk CLI Help](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands#snyk-aibom-test).
* Updated [`redteam`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/redteam) with additional options, and also added [Snyk Agent Red Teaming](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-agent-red-teaming).

### Snyk Analytics

* The [Analytics Overview](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/overview-tab) tab now includes the **Projects Monitored** widget.
* The [Snyk Pull request checks usage & performance report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/prevention-reports#pull-request-checks-usage-and-performance-report) is now in **General Availability** for Enterprise plan users.
* Added the **Assessing active security incidents** option to the [Zero-Day report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/remediation-reports#zero-day-report).

### Snyk CLI

* The latest Snyk CLI version is [1.1303.2](https://github.com/snyk/cli/releases/tag/v1.1303.2).
* The [environment variables page](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli#configure-max-network-attempts) has been updated to include `SNYK_MAX_ATTEMPTS` .

### Other updates

* The [Open Source license compliance](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance#license-updates) page has been updated to clarify that the Snyk database of supported licenses is regularly updated to match new releases of the SPDX License List.
* The `package-health-check` directive [sample script links](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/directives#secure-at-inception-package-health-check-experimental) now point to the improved recipes repository in GitHub. Use the [skills](https://github.com/snyk/studio-recipes/tree/main/command_directives/synchronous_remediation/skills/secure-dependency-health-check) script or the [hook](https://github.com/snyk/studio-recipes/tree/main/guardrail_directives/package_enforcement/cursor/hooks) script.
* For Snyk Container, the [Configure repository monitoring](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/configure-repository-monitoring) feature is now in Early Access.
* [Snyk 2.0](getting-started/snyk-2.0-platform-improvements.md) introduces UI enhancements to the platform navigation and is available in Early Access. This is being rolled out gradually, so not all users see the new navigation at the same time. If you are an existing user, you can switch between the new and classic navigation at any time using the toggle in your user profile menu. What is different:
  * **Global scope selector:** the top bar serves as the primary tool to navigate between different levels of your account. Use the scope selector to switch between Tenant, Groups, and Organizations. When you select a scope, the side menu automatically displays the relevant tools and data for that area.
  * **Analytics** becomes the centralized location for all reporting, including overview reports, dependencies, and license information.
  * **Settings** becomes the unified area for managing members, billing, integrations, and account preferences. When you use the scope selector to switch between Groups and Organizations, all relevant settings for that area are displayed under **Settings**.
  * **The Organization Dashboard** has been replaced by the scope selector and the **Analytics** overview page (accessible only to Tenant users, at Tenant-level).
* The [License Policies](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/policies/license-policies) page has been updated to reflect that newly supported licenses now have a default **Severity** of **None** and only appear in results if you explicitly configure this behavior.
* The [Enterprise implementation guide](implementation-and-setup/enterprise-implementation-guide/) has been updated to reflect the actual journey you would take as a new user onboarding with Snyk on the Enterprise plan. This includes adding guidance on how to create your Organization Template, configure all available features, and includes key decision callouts to help guide you when making essential decisions in this process.

## February 2026

### Snyk CLI

* [Container SBOM](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/container-sbom) has been updated with additional options.
* The Snyk CLI latest release version is [v1.1303.0](https://github.com/snyk/cli/releases/tag/v1.1303.1).

### Snyk Open Source

* [Breakability risk levels](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/snyk-pull-or-merge-requests/breakability-risk-levels) is now in Early Access.
* Improved .NET scanning is now a General Availability feature. The [.NET (C# and VB.NET](supported-languages/supported-languages-list/.net/) section has been updated to reflect this change.

### Snyk supported languages

* [Python](supported-languages/supported-languages-list/python/) has been updated to include support for Python version 3.12 for Snyk Code.
* Snyk Code now supports C# 14 and .NET 10, Kotlin and Java (including Spring WebFlux and JAX-RS), JavaScript and TypeScript (including Sequelize), Go (including Fiber), and Swift (including grpc-swift).
* Ruby 4.0 is supported in Snyk Code, starting with core parser improvements and stronger handling of Ruby modules.

### Snyk Studio

* A [`snyk_package_health_check` directive](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/directives#secure-at-inception-package-health-check-experimental) has been added to assist you in evaluating open-source packages for security vulnerabilities, maintenance health, community engagement, and popularity. Snyk has provided a secure dependency health check skill with a sample script and an enforce security scan on new packages hook with a sample script to integrate the `snyk_package_health_check` into your workflow.
* Additional guidance on [available MCP profile types](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/getting-started-with-snyk-studio#configure-the-snyk-mcp-profile) and their applicable tools, with configuration instructions, has been added.

### Other updates

* Updated documentation references and rule mappings from OWASP Top 10 (2021) to the OWASP Top 10 (2025) revision. This keeps security-category labels and cross-references aligned with the OWASP taxonomy.
* [Broker Contexts](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker/broker-context) was released for Universal Broker, including how contexts help segment and route Broker connections across environments and Organizations. This improves guidance for running multiple Broker deployments with clearer isolation.

## January 2026

### Snyk Code

* The [Snyk Code Security policies](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code#manage-ignores-at-the-group-level-through-snyk-code-security-policies) documentation has been updated to clarify that Snyk Code Security policies are different to Snyk Security Policies.
* The [Enable Snyk Agent Fix](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically#enable-snyk-agent-fix) section has been enhanced with more details and clear configuration steps.

### Snyk CLI

* [SBOM test](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/sbom-test) command is now in Early Access.
* The latest [Snyk CLI version](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/install-the-snyk-cli) available is v1.1302.1.
* The [debugging mode of the Snyk CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-cli/debugging-the-snyk-cli) has been enhanced with error catalog codes and exit codes.

### Snyk IDE

* You can now use the Issue View Options in [Eclipse](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects#issue-view-options), [JetBrains](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy#general-settings), and [Visual Studio Code](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-configuration-environment-variables-and-proxy#scan-configuration) to filter issues by their Code Consistent Ignores status.
* The Risk Score Threshold option has been added to the [Visual Studio Code](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-configuration-environment-variables-and-proxy#scan-configuration) extension to allow you to filter Open Source issues by their risk score.

### Snyk Studio

* The [Agentic security with Snyk Studio](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio) documentation has been restructured to accurately reflect the workflow you would go through when using Snyk Studio for the first time, and using the available [Quickstart guides](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides).
* The [Snyk Studio Adoption](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/usage-analytics) report is now available under [Redesigned analytics](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/overview-tab#snyk-studio-adoption).
* The [supported tools list](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio#mcp-server-supported-tools) was updated to clarify that Snyk supports `snyk_send_feedback`.

### Snyk supported languages

* [PHP](supported-languages/supported-languages-list/php.md) has been updated to include support for PHP version 8.5.
* [JavaScript](supported-languages/supported-languages-list/javascript/) was updated with support for Yarn 4.
* [Ruby](supported-languages/supported-languages-list/ruby.md) was updated with support for Ruby 4.
* [Python](supported-languages/supported-languages-list/python/) has been updated to remove the limitation note for Projects with downloaded dependencies.
* [Go](supported-languages/supported-languages-list/go.md) has been updated to include support for the Go standard library, for Go with Open Source.
* Several supported language pages and their rules have been updated with Code analysis support in Early Access: [Rust](supported-languages/supported-languages-list/rust.md), [Swift and Objective-C](supported-languages/supported-languages-list/swift-and-objective-c.md), [Dart and Flutter](supported-languages/supported-languages-list/dart-and-flutter.md), [Groovy](supported-languages/supported-languages-list/groovy.md), [Rust rules](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/snyk-code-security-rules/rust-rules), [Objective-C rules](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/snyk-code-security-rules/objective-c-rules), [Dart and Flutter rules](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/snyk-code-security-rules/dart-and-flutter-rules), and [Groovy rules](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/snyk-code-security-rules/groovy-rules).

### Other updates

* [Container monitor](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/container-monitor) has been updated to include a new CLI option `--prune-repeated-subdependencies` to prune depgraphs for large container scans.
* The [redesigned Snyk Analytics experience](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/overview-tab) is now in General Availability.
* The [Bitbucket Cloud documentation](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations/bitbucket-cloud) has been updated to make it clear that scopeless API tokens are not supported for this integration.

## 2025

<details>

<summary>December - January 2025 Documentation updates</summary>

### December 2025

#### *Snyk API

* The API docs navigation was enhanced with additional package-related reference pages (including `ContainerRegistryImagePolicy`).

#### Snyk Integrations

* The [Partner integrations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/partner-integrations) page has been updated, including additional coverage for Coding Assistants and how they can use Snyk Studio (MCP) in agentic workflows.
* [JavaScript](supported-languages/supported-languages-list/javascript/) navigation has been enhanced with better redirect and routing features.

#### Snyk Studio

* The [Snyk Studio - Agentic integrations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio) documentation has been updated to provide a clearer explanation of MCP usage and the available Snyk Studio tools.
* The [Quickstart guides for Snyk Studio](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides) were updated with new and refreshed setup guidance, including [Cursor](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/cursor-guide) and [Windsurf](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/windsurf-guide).

#### Other updates

* GitHub Cloud App and GitHub Server App have been added to the list of [supported SCMs for Dockerfile analysis](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/scan-your-dockerfile#supported-scms-for-dockerfile-analysis).
* The `snyk-scm-contributors-count` docs were updated with prerequisites and setup notes. See [snyk-scm-contributors-count](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-scm-contributors-count).
* The Declining Balance documentation from the [Snyk Declining Balance of Hours Service Description](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/ELvljsaLKPkSpffOkmsQ/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/declining-balance) page was updated with service description and expiration details.
* The IaC issue-reporting walkthrough from the 'Getting started with Snyk IaC' page was updated to remove outdated screenshots and copy.
* PR template variables were updated on the [Variables list and description](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/snyk-pull-or-merge-requests/customize-pr-templates/variables-list-and-description) page with a new container base image `short name` values for cleaner PR titles and messages.

### November 2025

#### Snyk Container

* The list of [operated distribution systems supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container) has been updated with support for Chisel.

#### Snyk CLI

* The latest [Snyk CLI version](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/snyk-cli/snyk-cli/install-the-snyk-cli) available is v1.1301.0.
* The [CLI help](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/reachability-analysis#using-reachability-analysis-with-snyk-cli) has been updated with commands for reachability analysis.

#### Snyk IDE

* The Automated Org Selection feature uses repository context to choose an Organization. Manual configuration overrides this automated selection. If the selection fails, Snyk defaults to your preferred Organization setting. The feature is available for the [Eclipse plugin](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-ide-plugins-and-extensions/eclipse-plugin/configuration-of-the-eclipse-plugin), the [JetBrains plugin](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy), the [Visual Studio extension](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-configuration-environment-variables-and-proxy), and the [Visual Studio Code extension](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-configuration-environment-variables-and-proxy).

#### Snyk integrations

* The Amazon Q guide for Snyk Studio now includes [updated instructions](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/amazon-q-guide#install-the-snyk-mcp-server-in-the-amazon-q-ide-extension) for configuring the Snyk MCP Server in VS Code and JetBrains.

#### Other updates

* [Reachabilty analysis](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/reachability-analysis) has been updated with instructions on how it works and how to use it in both the Snyk Web UI and the Snyk CLI and clear support for specific languages and package managers.
* The [Pre-defined roles](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-platform-administration/user-roles/pre-defined-roles#role-types) documentation has been updated to communicate that the Organization Admin role and associated permissions supersede any Group Member role restrictions.
* The [severity condition](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/policies/security-policies/security-policies-conditions) is now available in Group-level policies. Use this feature to create more granular policies for Snyk Code and Snyk Open Source findings, for example, ignoring a finding or changing its severity.

### October 2025

#### Snyk API

* A new [API migration guide](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/api-end-of-life-eol-process-and-migration-guides/guides-to-migration/v1-reporting-apis-to-export-api-migration-guide) is available to help you migrate from the v1 Reporting API to the REST Exporting API.
* The Export API has been improved with the option to [limit the link expiration](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/using-specific-snyk-apis/export-api-specifications-columns-and-filters#data-retention).

#### Snyk Broker

* The [Universal Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/universal-broker) release status has transitioned to Generally Available.
* The page [Upgrade an Organization from Classic Broker to Universal Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker/upgrade-an-organization-integration-from-classic-broker-to-universal-broker#migrating-multiple-organizations) has been updated with steps to migrate multiple Organizations at a time.

#### Snyk CLI

* Snyk CLI now supports uploading files and folders for Snyk Code scanning. The command [`code-test`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/code-test) has been updated with options reflecting these capabilities.
* The latest [Snyk CLI version](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/snyk-cli/snyk-cli/install-the-snyk-cli) available is v1.1300.2.

#### Snyk integrations

* The list of Snyk MCP quick guides now includes [Devin guide](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/devin-guide), [Factory guide](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/factory-guide), [Factory terminal guide](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/factory-terminal-ide-guide).
* The Snyk MCP Server has been rebranded as [Snyk Studio](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio).
* [SCM integration support for Python](supported-languages/supported-languages-list/python/scm-integrations-and-python.md) has been updated with support for Python 3.14.

#### Other updates

* The [Operating system distributions supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container#minimus) have been updated to include include support for Minimus, Ubuntu 25.10 - Questing Quokka, and Ubuntu 25.04 - Plucky Puffin.
* For [Ruby](supported-languages/supported-languages-list/ruby.md), versions 2.3.X are no longer supported. The Ruby-specific versions have been updated to include more version patches.
* PR Check report was added as Early Access to the available reports to identify Snyk PR check locations, increase adoption, and pinpoint common failure impacts on developer workflows.
* You can now label your assets with metadata on repository assets and build artifacts, helping tag, manage security, and group items by features. An asset label differs from an asset tag, which enables key-value tags for structured metadata, allowing for granular filtering, policy creation, and improved system alignment.
* [JavaScript for open source](supported-languages/supported-languages-list/javascript/#javascript-for-snyk-open-source) has been updated to include full support for pnpm Projects.

### September 2025

#### Snyk Container

* The instructions for [installing the Snyk Controller on Amazon Elastic Kubernetes Service (Amazon AKS)](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/kubernetes-integration/install-the-snyk-controller/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks#create-an-eks-node-role-for-your-node-group-and-add-the-trust-relationship-for-the-iam-role) have been updated with details for configuring trust relationships for the IAM role.
* The list of [operating system distributions supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container) has been updated to include SUSE Linux Enterprise Server 15.7 and Rocky Linux 10.

#### Snyk integrations

* The SCM integration for Bitbucket Data Center/Server now supports the Required Builds feature for granular control over pull requests. To learn more, visit [Required Builds](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations/bitbucket-data-center-server#required-builds).
* [GitLab](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/analyze-pr-checks-results#gitlab) is supported for PR check results. This feature blocks merge requests with security issues when the "Pipelines must succeed" setting is enabled.
* The Snyk MCP quick guides list has been enriched with the following guides: [Claude Code](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/claude-code-guide), [Continue](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/continue-guide), [JetBrains AI Assistant](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/jetbrains-ai-assistant), and [JetBrains Junie](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/jetbrains-junie)

#### Other updates

* For Java and Kotlin, the list of [supported Gradle versions](supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-package-managers-and-package-registries) now includes Gradle 9.
* For [Ruby](supported-languages/supported-languages-list/ruby.md), an end-of-support notice has been added to say that starting Oct 1, 2025, Fix PRs are no longer supported for Projects using Ruby versions 3.1.x and lower. The table of supported Ruby versions has also been updated.
* For JavaScript, [support for pnpm Projects](supported-languages/supported-languages-list/javascript/#support-for-pnpm) has been added.
* `Raise Support Community Cases` and `View Support Community Cases` Tenant-level permissions have been added. To learn more about which Tenant roles these permissions apply to, visit Pre-defined roles, [Tenant-level permissions](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-platform-administration/user-roles/pre-defined-roles#tenant-level-permissions).
* The [Analytics](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics) menu now updates its data daily instead of hourly.
* Learn how to resolve duplicated and unenriched assets discovered outside Group and Organization-level SCM integrations.
* You can now [exclude specific values](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab#exclude-filters) when you filter your reports.

### August 2025

#### Snyk API

* The [Export API](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/export) has been enhanced with the project\_target\_file field.
* A new dataset for usage events has been added to the [Export API.](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/export)

#### Snyk CLI

* [Experimental builds](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/releases-and-channels-for-the-snyk-cli#experimental-builds) information is now available for the CLI releases and channels.
* The [AI-BOM](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/aibom) Snyk CLI command is now available with any stable CLI release.
* A new Snyk CLI analytics page is now available, providing information about [Essential Operational Analytics](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/snyk-cli-analytics#essential-operational-analytics) and [Optional Usage Analytics](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/snyk-cli-analytics#optional-usage-analytics).

#### Snyk integrations

* You can now add the Snyk MCP server to [Goose CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/goose-cli-guide) to secure code generated with agentic workflows through an LLM.
* You can now integrate Akamai with the Snyk API & Web to discover and scan your API. See the [API Security](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/partner-integrations#api-security) section under the Partner integrations page for more details.
* The [Jira Cloud documentation](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration) has been updated for parity with the current version.

#### Other updates

* A new [Risk exposure report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/exposure-and-coverage-reports#risk-exposure-report) has been released, providing you with a single, consolidated view of your security risks.
* The rollout to General Availability has started for the [Pull Request Experience](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience).
* The [Operating system distributions supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container) now include Debian 14 - Forky.
* Snyk now supports [Ruby versions](supported-languages/supported-languages-list/ruby.md#technical-specifications) 3.3 \[3.3.9] and 3.4 \[3.4.5]. If the Ruby version is not specified in the Gemfile, it defaults to version 3.1.

### July 2025

#### Snyk API

* The [Export API](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/export) is now available as GA.
* The Assets API is now available as Early Access.

#### Snyk CLI

* MCP updates:
  * [Updated the list of supported Snyk security tools into an AI system](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio#mcp-server-supported-tools).
  * Updated release status from experimental to [Early access](getting-started/snyk-release-process.md#early-access-features) and removed the experimental flag.
  * Added [Cursor](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/quickstart-guides/cursor-guide) as a new supported agentic IDE for MCP.
* PAT updates:
  * Added PAT support for [Snyk CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/authenticate-to-use-the-cli).
  * Added PAT support for Snyk CI/CD integrations ([CircleCI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ci-cd-integrations/circleci-integration-using-a-snyk-orb), [Jenkins](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ci-cd-integrations/jenkins-plugin-integration-with-snyk), [Maven](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ci-cd-integrations/maven-plugin-integration-with-snyk)).

#### Snyk Code

* Support for Python, JavaScript, and Typescript now includes more frameworks.

#### Snyk Container

[Operating system distributions supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container) have been updated to include: SUSE Linux Enterprise (SLE) 15.3+, Red Hat Enterprise Linux 10, and Oracle Linux 10.

#### Snyk IDE

* Added PAT support for all [Snyk IDE](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/) plugins and extensions.
* Added an [IDE Plugin Compatibility Matrix](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-ide-plugins-and-extensions/compatibility-matrix) for all supported versions.

#### Snyk integrations

* [Snyk Agent Fix in the PR](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience#snyk-agent-fix-in-the-pr) has added support for Bitbucket integrations, still in Early Access.
* The [minimum version](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/configure-pull-request-checks) of Bitbucket Server and Bitbucket Data Center required to use the integrations with PR checks has been updated to 7.4 and 8 respectively.

#### Snyk Open Source

[Scan open-source libraries and licenses](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/), [Snyk License Compliance Management](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management), and [Fix your vulnerabilities](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities) have been updated with the new **Issues** tab layout.

#### Other updates

* A new architecture for user documentation on developer tools is now available. This update groups the main developer tools into a single section and distinctly separates them from the integrations documentation.
* [Analytics](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/overview-tab) has a fresh new look.
* Added [Snyk Assist](snyk-learn/snyk-assist.md) documentation.
* The [Developer IDE and CLI usage report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/prevention-reports#developer-ide-and-cli-usage-report) has been improved with MCP-related data to provide better visibility into MCP usage.
* [Okta custom mapping documentation](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/implementation-and-setup/enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/custom-mapping/examples-setting-up-custom-mapping-for-idps/example-setting-up-custom-mapping-for-okta#construct-a-value-expression-that-creates-a-roles-array-to-be-sent-to-snyk) has been updated to clarify handling of the `Arrays.flatten(appuser.snyk_orgs)` value during setup.

### June 2025

#### Snyk Broker

* Updated the Snyk Broker documentation to include distinct steps for setting up the [Container Registry Agent with Docker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/snyk-broker-container-registry-agent#configuring-and-running-the-container-registry-agent), whether using the Classic or Universal Broker.
* Updated the [Using the API to set up Universal Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker/using-the-api-to-set-up-universal-broker/README.md) documentation with a Prerequisites section and clarified that the Snyk Broker App ID differs for each [region](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency#broker-client-urls).
* Snyk Learn courses have been integrated into the [Universal Broker](ehttps://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker/) pages.

#### Other updates

* [Usage settings](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/usage-settings) have been updated with the new **Billing and Usage** dashboard, available with the new Snyk Platform Access plan.
* [Snyk Platform Access credits](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/ELvljsaLKPkSpffOkmsQ/snyk-platform-access-credits) have been added with brief information on the new Snyk Platform Access plan.
* The troubleshooting sections for all [Snyk IDE plugins](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/) have been updated to include clear steps for working with the Logs details, which are available across all plugins.
* A new feature, the [Snyk Agent Fix in the PR](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience#snyk-agent-fix-in-the-pr), has been released, enabling the user to interact with inline comments by requesting an initial fix or a different suggestion, or by applying a specific fix by using the `@snyk /apply #` command.
* [Consistent Ignores](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/) for Snyk Code now fully supports CLI Upload.
* The page on Docker Desktop Extension integration has been removed due to the end of support.

### May 2025

#### Snyk CLI

* The `--platform` option was added to the [`container sbom`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/container-sbom) command.
* The MCP information was expanded to [Developer guardrails for agentic workflows](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio).

#### IDE plugins and extensions

* Information was added to the [JetBrains plugin troubleshooting](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/troubleshooting-for-the-jetbrains-plugin).
* Region information was updated on all [IDE pages](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/).

#### Snyk Code

* Legacy ignores can be converted using [bulk ignore conversion](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/convert-project-scoped-ignores-to-asset-scoped-ignores#bulk-ignore-conversion).
* DeepCode AI Fix has a new name: [Snyk Agent Fix](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically).

#### Snyk Container

[Configure the integration with Docker Hub](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/container-registry-integrations/integrate-with-docker-hub/configure-the-integration-with-docker-hub) has been updated to state that Snyk does not support Organization Access Tokens (OAT).

#### Snyk Integrations

The [Bitbucket Cloud App](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations/bitbucket-cloud-app) and [Jira App](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration) integrations are now available in the `SNYK-US-02` environment.

#### Other updates

* For [SCM integrations with Python](supported-languages/supported-languages-list/python/scm-integrations-and-python.md), the list of dependencies that are not supported has been updated to include `pip` for Python 2.7 and 3.7.
* [Python dependency filtering results](supported-languages/supported-languages-list/python/scm-integrations-and-python.md) have been updated to clarify the conditions in which certain packages and configurations are skipped by SCM scans.
* The list of supported package managers has been updated to include `conan`. See [C/C++](supported-languages/supported-languages-list/c-c++.md), [SBOM test](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/sbom-test), [Test an SBOM document for vulnerabilities](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities).
* [Instructions for upgrading an Organization integration from Classic Broker to Universal Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker/upgrade-an-organization-integration-from-classic-broker-to-universal-broker) were clarified.

### April 2025

#### Snyk API

* Several APIs have been updated; see the [Changelog](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/changelog).
* The navigation in the API section now reflects the use of Authentication and the Changelog for both the V1 and REST APIs.
* The [Authentication for API](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/authentication-for-api) page has been updated with region information and clarity on using the bearer token.
* The [API endpoints index and tips](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/api-endpoints-index-and-tips) page now has a note about how to find your `org_id`.

#### Snyk Essentials

* [The Inventory Overview tab](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/assets-inventory-layouts) is now available to provide insights and prescriptive guidance to improve your application security.
* [The Visibility column](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/assets-inventory-components#visibility) has been added to show the visibility status of your repositories.

#### Snyk Broker

Additional updates have been made to the [Universal Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/universal-broker) documentation to clarify the instructions and add details about the use of the APIs.

#### Snyk CLI

Information has been added about Snyk support for the Model Context Protocol (MCP) through the [`snyk mcp` experimental CLI command](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/usage-analytics).

#### Snyk Code

* [Consistent Ignores ](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code)is now available in Early Access. Your development teams can create ignores that are consistently respected regardless of how and where the test is run and what branch is being tested.
* Snyk Code supports gRPC libraries.

#### Snyk Container

* [Using Custom Base Image Recommendation](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations) has been updated with clarifications on how Snyk recommends images.
* The list of [Operating system distributions supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container) has been updated to include Alpine Linux 3.21, Ubuntu 25.04 - Plucky Puffin, and Ubuntu 24.10 - Oracular Oriole.
* The section describing the automated integration process for Amazon Elastic Container Registry (ECR) has been removed, as Snyk no longer supports this method.

#### Snyk Integrations

* For the [Jira integration](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/jira-and-slack-integrations/jira-integration#prerequisites-for-jira-integration-with-snyk), Snyk now supports Jira versions 5 to 10.
* For [SCM integrations with Gradle](supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md), Snyk now supports `allprojects` and `subprojects` blocks, as well as Spring Boot plugins BOMs.

#### Other updates

* DAST scanning is now available with [Snyk API & Web](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/overview#select-scanning-methods), enabling users to discover and test the security of their APIs and web apps, including AI-generated ones.
* PR Checks is now available with a General Availability release status.

### March 2025

#### Snyk Broker

* The Snyk Broker section has been divided into [Universal Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker) and [Classic Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/classic-broker) documentation and the [main page](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/snyk-broker) has been updated.
* The Classic Broker installation instructions now include the command to set the `BROKER_SERVER_URL` for [Docker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker) and the `brokerServerUrl` for [Helm](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm).

#### Snyk API

* The [V1 API overview](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/v1-api) and [reference](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference) are now on the user docs site only. Additional details from Apiary have been added to the V1 reference on the user docs site. The API reference has been removed from the V1 API Apiary site.
* A section has been added for [pages that explain how to use specific APIs in depth](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/using-specific-snyk-apis).

#### Snyk CLI, CI/CD, IDE

* [Advanced use of Snyk Container CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/advanced-use-of-snyk-container-cli) now includes support for scanning Kaniko image archives.
* The [support policy for the CI/CD plugins](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ci-cd-integrations#support-policy) was updated to align with the CLI support policy.
* The Net new issues feature was added to the IDE documentation for [Eclipse](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/dintegrations/snyk-ide-plugins-and-extensions/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects#net-new-issues-versus-all-issues), [JetBrains](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/run-an-analysis-with-the-jetbrains-plugin#net-new-issues-versus-all-issues), [Visual Studio](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/view-analysis-results-from-visual-studio-extension#net-new-issues-versus-all-issues), and [Visual Studio Code](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/view-analysis-results-from-visual-studio-code-extension#net-new-issues-versus-all-issues), and [troubleshooting information](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/troubleshooting-ides/net-new-issues-delta-scan-troubleshooting) was added.

#### Snyk Code

* The Generated Pull Requests report is now available in Early Access. This report provides an overview of how Fix, Backlog, and Upgrade PRs are used and highlights the efficiency of PR merges.
* [The Pull Request Experience](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience) now supports GitLab and Azure Repos SCM integrations, with a few [limitations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience#inline-comments).
* New Snyk Code filters and columns were added to [Snyk Reports](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/issue-columns-dictionary#issue-characteristics) and [Snowflake Data Share](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary): File Path, Code Region, and Asset Finding ID.
* Snyk Code now supports [Rust](supported-languages/supported-languages-list/rust.md) and [Groovy](supported-languages/supported-languages-list/groovy.md) available in Early Access and accessible from Snyk Preview.

#### Snyk Essentials

* A new feature is now available in Snyk Essentials, introducing a new type of [asset tag](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/policies/assets-policies#asset-tagging) known as GitHub custom properties.
* [Asset tags](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/assets-inventory-components#tags) have been redefined and are now clearly separated into system tags and user-defined tags.

#### Snyk Integrations

* The [GitHub Server App](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations/github-server-app) has moved into General Availability.
* The Jira integration documentation has been updated to state that Snyk supports version 5 to version 9.

#### Other updates

* The PCI-DSS v4.0.1 report is now available in Early Access. This report leverages Snyk scan results to assess, prove, and improve readiness for PCI-DSS AppSec compliance regarding SCA and SAST vulnerabilities.
* The [Repositories Tested in CI/CD report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/prevention-reports#repositories-tested-in-ci-cd-report) is available in Early Access. This report tracks Snyk CI/CD testing to prevent vulnerable production deployments.
* [Severity levels](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/severity-levels#why-are-there-multiple-cvss-scores-for-the-same-vulnerability) now provide more details about the CVSS v4.0.

### February 2025

#### Snyk Essentials

* The Integrations UI at the Group level has been enhanced to improve readability and actionability and provide inline instructions and inline profile helpers.
* Group-level [Integrations documentation](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/integrate-with-snyk#integrations-syncing-time) has been updated with new, more accurate sync times.
* The [asset filter](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/policies/assets-policies/create-policies) documentation has been consolidated into one section, and it now links to all relevant areas, such as Inventory and Asset Policy filters.

#### Other updates

* A new [Automated Provisioning guide](implementation-and-setup/enterprise-setup/auto-provisioning-guide.md) has been created for **Pilot** and **Enterprise** **users**, detailing the steps of the auto-provisioning process for new and existing user accounts.
* [Snyk Code PR Checks](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/configure-pull-request-checks#configure-for-code-analysis-click-to-expand) are in General Availability.

</details>

## 2024

<details>

<summary>December - January 2024 Documentation updates</summary>

#### December 2024 and January 2025

**Snyk Container**

* Page "Integrate with Docker Desktop Extension" has been updated to include an end-of-support notice. Effective June 20, 2025, the integration with Docker Desktop will no longer receive updates or technical support.

**Snyk CLI and IDEs**

* [Eclipse IDE](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/) major update
* [Visual Studio IDE](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/) major update
* Region configuration update for [IDEs](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/)
* [Snyk images EOL policy updated](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-ci-cd-integrations/snyk-images-and-eol-image-policy)
* [`snyk container test`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/container-test) and [`snyk container monitor`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/container-monitor) option `--exclude-node-modules` added

**Other updates**

* [Snyk Admin](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-platform-administration/snyk-admin) pages have been updated to reflect the addition of [Tenants](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/tenant) in the Snyk hierarchy, including a new infographic to illustrate the Tenant position in the [hierarchy](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/groups-and-organizations#the-snyk-hierarchy).

#### November 2024

**Snyk Container**

* The list of [operating system distributions supported by Snyk Container](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container) has been updated to include Ubuntu 24.10 - Oracular Oriole and Ubuntu 24.04 - Noble Numbat 04.
* [How Snyk Container works](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/how-snyk-container-works) has been updated with details on the logic Snyk applies when providing public base image recommendations.

**Other updates**

* The Pull Request Checks section has been updated to include the new [Pull Request Experience](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/pull-request-checks/pull-request-experience) for PR Checks.
* The [Supported languages](supported-languages/supported-languages-package-managers-and-frameworks.md) page has been reorganized to provide detailed information about language availability for each Snyk product. Additionally, it provides a list of package managers, frameworks, and features for each supported language.
* A service account using OAuth 2.0 can now be [created through the Snyk Web UI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/service-accounts/service-accounts-using-oauth-2.0#create-oauth-service-accounts-through-the-ui).
* The [API index](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/api-endpoints-index-and-tips) now includes entries for each endpoint mentioned in the Snyk user docs.
* The [Developer IDE and CLI usage report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/prevention-reports#developer-ide-and-cli-usage-report) has been enhanced with additional functionalities: **Developer email address** and **PDF export**.
* The [Vulnerabilities Detail report](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/remediation-reports#vulnerabilities-detail-report) has been enhanced with additional functionalities, such as **Target indication** and **Column picker**.

#### October 2024

**Snyk API**

* [Asset inventory components](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/assets-inventory-components#clusters) has been updated to include details on clusters.

**Snyk CLI and IDEs**

* The [CLI authentication page](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/authenticate-to-use-the-cli) has been updated for the OAuth 2.0 protocol.
* The page [Debugging the Snyk CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/debugging-the-snyk-cli) has been added.
* [CLI standalone executables](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/install-the-snyk-cli#direct-binary-download) have been updated to include Alpine Arm64.
* IDE Eclipse[ plugin](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/) and [JetBrains plugin](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/) documentation pages have been updated.
* [Authentication information](https://app.gitbook.com/o/-m4tdxg8qotlggznlpfr/s/ieejsxqqu36y0vmfv8zf/integrations/snyk-ide-plugins-and-extensions/) has been updated for all IDEs.

**Snyk Integrations**

* [Snowflake Data Share](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share) is now in [GA](getting-started/snyk-release-process.md).
* [Snyk SCM integrations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations) has been updated with additional notices relating to repository retrieval and permission or scope modifications after initial configuration.
* GitHub Cloud App has been added to feature support notices for Fix, Backlog, and Upgrade PRs.
* Snyk SCM integrations has been updated to include a table detailing the [permissions and scopes](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/user-permissions-and-access-scopes#github-cloud-app-permission-requirements) required for the GitHub Cloud App.

**Other updates**

* [Getting started](getting-started/README.md) has been updated to centralize content related to everything you need to know before using Snyk.
* Scanning methods have been added for the [Dart and Flutter](supported-languages/supported-languages-list/dart-and-flutter.md) languages.

#### September 2024

**Snyk API**

* A prerequisites section has been added to the Group level of [GitHub integration](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations/github-enterprise#prerequisites), and more details about the [Pull personal repositories](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/group-level-integrations/github-for-snyk-essentials) option have been added to the same documentation page.
* The [Set up Insights](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/set-up-insights) section was updated to emphasize the risk factors availability for each integration option.
* The Snyk Runtime Sensor has been updated to reflect the importance of adopting it to achieve the most effective integration and to access its continuously expanded set of features.

**Snyk Broker**

The Universal Broker feature is now available in Early Access. The Universal Broker separates deployment and container concerns from connection concerns. It allows for a smaller or a single deployment to support numerous connections of varied types.

**Snyk CLI**

* The [CLI commands and options summary](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/cli-commands-and-options-summary) was updated.
* [Authentication](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/authenticate-to-use-the-cli) has been updated.
* Configuration has been updated: Environment variables for Snyk CLI, [`snyk config`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/config) help, [`snyk config environment`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/config-environment) help.

**Snyk Integrations**

The Snowflake Data Share section has been updated to include a [Data Share Dictionary](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary), designed to help you navigate and build your dataset.

**Other updates**

* The updated [Regional hosting and data residency](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency) page was published.
* [Glossary](getting-started/glossary.md) terms were updated for SCA, SAST, DAST, and IAST as well as Software Composition Analysis.
* [Early Access](getting-started/snyk-release-process.md#early-access) release status notices were updated.

#### August 2024

**Snyk API**

* Links in the API reference docs have been updated.
* The [API endpoints index and notes](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/api-endpoints-index-and-tips) have been updated.

**Snyk CLI**

* [`snyk auth`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/auth) command help updated to reflect OAuth default.
* [CLI authentication](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/authenticate-to-use-the-cli) instructions updated for OAuth default and improved flow.
* [`snyk config environment`](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/config-environment) command help has been added.
* CLI [support for pnpm added](supported-languages/supported-languages-list/javascript/#support-for-pnpm).

**Snyk IDE**

* [CLI authentication](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/authenticate-to-use-the-cli) instructions updated for IDE.
* IDE authentication instructions updated: [Eclipse](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/authentication-for-the-eclipse-plugin), [Jetbrains](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/jetbrains-plugin/authentication-for-the-jetbrains-plugins), [VS extension](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/authentication-for-visual-studio-extension), [VS Code extension](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/authentication-for-visual-studio-code-extension)

**Snyk Integrations**

* Git repository cloning has been renamed [Workspaces for SCM integrations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/workspaces) to better reflect its functionality. Additional detail on [enablement](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/workspaces#manage-workspaces) has been added.
* The [relationship](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations/github-cloud-app#how-to-set-up-the-github-cloud-app) between GitHub organizations and Snyk Organizations when integrating with the GitHub Cloud App has been clarified.

#### July 2024

**Snyk API**

* The API documentation now provides the API Reference and explanatory documentation in the [API section](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/snyk-api).
* The [API End of Life (EOL) process and migration guides](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/api-end-of-life-eol-process-and-migration-guides) are now published and updated to support the process, which began in July.
* [Asset inventory filtering](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/assets-inventory-components#asset-tabs) describes the new, simplified view that provides an improved experience of filtering the assets.
* The [Asset inventory layouts](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/manage-assets/assets-inventory-layouts) have been renamed to better reflect their functionality.
* Four new SCM integrations are now available for Snyk:
  * [Atlassian Compass](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/application-context-for-scm-integrations#atlassian-compass)
  * [Harness](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/application-context-for-scm-integrations#harness)
  * [OpsLevel](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/application-context-for-scm-integrations#opslevel)
  * [Datadog Service Catalog](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/application-context-for-scm-integrations#datadog-service-catalog)

**Snyk Integrations**

* A comparison of the GitHub and GitHub Enterprise integrations functions now resides on the [SCM, IDE, and CI/CD integrations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations#github-vs-github-enterprise) page.
* Steps for [migrating from the GitHub integration to the GitHub Enterprise integration](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/scm-integrations/organization-level-integrations/github#migrate-to-the-github-enterprise-integration) now reside on the GitHub integration page.
* The [Snyk SCM Integrations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations) page now contains information critical to using these integrations successfully in your SDLC. This includes:
  * [Git repository cloning](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/workspaces) details
  * [Deployment recommendations](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/deployment-recommendations) for Enterprise customers
  * [User permissions and access scope requirements](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/user-permissions-and-access-scopes) for each SCM integration
  * Instructions on how to generate [integrated SCM tokens for Snyk Broker](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/scm-integrations-and-snyk-broker#integrated-scm-tokens-for-classic-broker)

**Other updates**

* **Snyk Reports:** The [issue column dictionary](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/issue-columns-dictionary#issue-vulnerability-details) includes new filters and columns for Jira (JIRA ISSUES LIST, LATEST JIRA ISSUE) and EPSS (EPSS SCORE, EPSS PERCENTILE). This allows you to manage your work with Jira and to include EPSS in your prioritization steps.
* **Snyk Security:** Snyk has improved the prioritization workflow and risk assessment by adopting [CVSS V4.0](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/fix/prioritize-issues-for-fixing/severity-levels#severity-levels-and-cvss) as the default evaluation for new vulnerabilities.
* **Fix code vulnerabilities automatically:** [DeepCode AI Fix](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically#snyk-agent-fix-language-support) is now available in AWS Environments and JetBrains IDEs. If you use AWS multi-tenant environments, turn on the Snyk Preview [Snyk Code Fix Suggestions](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically#enable-snyk-agent-fix) and retest with Snyk in your IDE.

</details>
