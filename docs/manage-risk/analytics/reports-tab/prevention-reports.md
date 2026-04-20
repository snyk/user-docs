# Prevention reports

The Prevention reports section includes the following reports:

* [Developer IDE and CLI usage](prevention-reports.md#developer-ide-and-cli-usage-report) report
* [Pull request checks usage & performance](prevention-reports.md#pull-request-checks-usage-and-performance-report) report
* [Repositories tested in CI/CD](prevention-reports.md#repositories-tested-in-ci-cd-report) report
* [Snyk Generated Pull Requests](prevention-reports.md#snyk-generated-pull-requests-report) report

## Developer IDE and CLI usage report

To use this report, you must ensure you have installed the following prerequisites:

* Snyk CLI
  * version 1.1292.1 or newer (for CLI and IDE plugins usage)
  * version 1.1297.0 or newer for general Agentic scans (Snyk Studio using MCP)
  * version 1.1298.1 or newer for granular Agentic scans (such as MCP host)
* VS Code 1.86.0 or newer and Snyk Security plugin 2.3.3 or newer
* IntelliJ IDEs 2023.3 or newer and Snyk Security plugin 2.7.3 or newer
* Visual Studio 2019, 2022 and Snyk Security Plugin 1.1.47 or newer
* Eclipse 2023.12 or newer and Snyk Security plugin 2.1.0 or newer

This report shows the adoption of Snyk testing in local development through the IDE plugins, using the CLI locally or incorporating Snyk Studio into agentic workflows. The report is available under the Change Report dropdown at the Group and Organization levels.

{% hint style="info" %}
This report focuses on the local developer experience and does not include the use of CI/CD. In addition, it does not show Organizations or developers that have never used the CLI, IDE, or Snyk Studio (through MCP).
{% endhint %}

Security teams can use this report to demonstrate strong shift-left behavior as a model behavior to bring to other teams. This report also shows where teams or individual developers are not adopting Snyk locally. Companies can use this report to encourage more shift-left behavior.

This report shows the test usage in the IDE, CLI, and Snyk Studio by developers. Teams can filter by date and Organization. The report includes visibility into metrics such as:

* Total number of developers running scans and the number of scans in IDE, CLI, and Agentic integrations (Snyk Studio)

<figure><img src="../../../.gitbook/assets/Screenshot 2025-07-22 at 10.16.07.png" alt="Total number of developers running scans and the number of scans in IDE, CLI, and Agentic integrations (Snyk Studio)"><figcaption></figcaption></figure>

* Charts and summary tables breaking down this data by the environment of the scan

<figure><img src="../../../.gitbook/assets/Screenshot 2025-07-22 at 10.18.03.png" alt="Developers running scans by type of environment and Total number of scans by developer graphs"><figcaption></figcaption></figure>

* Charts and summary tables breaking down this data by different dimensions, such as IDE plugins or Agentic integrations

<div><figure><img src="../../../.gitbook/assets/Screenshot 2025-07-22 at 10.19.11.png" alt="Developers running scans by IDE and Total number of IDE scans graphs"><figcaption></figcaption></figure> <figure><img src="../../../.gitbook/assets/Screenshot 2025-07-22 at 10.19.17.png" alt="Developers running scans by Agentic integrations and Total number of scans run by Agentic integrations graphs"><figcaption></figcaption></figure></div>

* Charts and summary tables breaking down this data by the Snyk scan type

<figure><img src="../../../.gitbook/assets/Screenshot 2025-07-22 at 10.20.04.png" alt="Charts and summary tables breaking down this data by the Snyk scan type"><figcaption></figcaption></figure>

* List of organizations and developers adopting Snyk locally

<figure><img src="../../../.gitbook/assets/ide usage by developer.png" alt="List of Organizations where Snyk was adopted locally "><figcaption></figcaption></figure>

## Pull request checks usage & performance report

{% hint style="info" %}
**Feature availability**

Snyk Pull request checks usage & performance report is in Generally Availability for Enterprise plan customers. For more information, visit [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

This report enables you to track and analyze the results and adoption of Snyk pull request checks.&#x20;

By monitoring trends in successful, failed, and errored PR checks, AppSec teams can detect and investigate spikes in failures and identify source and target configurations that need attention. Comparing current adoption against a previous period also surfaces whether coverage is improving or regressing over time while identifying the coverage gaps that need to be addressed.

### PR check performance & status

High level metrics provide an overview of how often PR checks are passing along with how often checks are being marked as successful to override a failed status. The PR checks by status over time trend gives you a weekly view of PR check outcomes while additional views break these outcomes down by specific scan types.

A high rate of failed PR checks may indicate emerging risk areas while a high rate of overridden checks may signal that developers are bypassing security gates and warrant further investigation. An increasing success rate over time can demonstrate that developers are producing more secure code earlier in your software development life cycle.

<figure><img src="../../../.gitbook/assets/unknown (24).png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/unknown (25).png" alt=""><figcaption></figcaption></figure>

#### Error PR checks by error message

This table surfaces recurring technical issues that may be preventing PR checks from completing, such as misconfigured integrations or unsupported project types. By surfacing these directly, the report helps teams identify recurring configuration issues, making it easier to prioritize fixes that most improve developer experience.

<figure><img src="../../../.gitbook/assets/image (281).png" alt=""><figcaption></figcaption></figure>

### Pull-request overview & adoption

This section focuses on pull request outcomes and how broadly PR scanning is enabled across your repositories and organizations. You can quickly see your PR check coverage across all repositories for Snyk Code and Snyk Open Source and how often Snyk PR checks are surfacing vulnerabilities in PRs before they reach a production branch. Different tables allow you to drill down into PR scanning adoption and PR scanning performance by Group, Organization, and repository.

<figure><img src="../../../.gitbook/assets/image (284).png" alt=""><figcaption></figcaption></figure>

#### PR scanning by

Viewing the **Last Commit** data when viewing by repository requires [SCM Group integration](../../../developer-tools/scm-integrations/group-level-integrations/).

PR check enablement for all targets and projects imported through an integration is determined by the settings for that integration unless those settings are overridden at the project-level. A repository is considered enabled for PR checks for a specific Snyk product if at least one organization importing that repository has PR Checks enabled for all projects associated with that product type for at least one integration.

**Integrations With PR Scan** under the Organization view captures the integrations with at least one PR check scan type enabled.

**Integrations With Scan Enabled** under the Group view counts enablement by integration settings for all integrations that exist under all the organizations under that group.

Repository status will be "N/A" under the repository view when a repository is new to Snyk or if the specific Snyk product does not apply (there are no projects for that product type imported for the target repository).

<figure><img src="../../../.gitbook/assets/image (252).png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (276).png" alt=""><figcaption></figcaption></figure>

#### PR scanning performance by

The **% PRs With Failed Checks** indicates how often Snyk detected a vulnerability that violated the policy set under your PR check settings for at least one of the checks for that PR. This can help you understand the value in running these scans before your code is already merged for specific organizations and repositories.&#x20;

<figure><img src="../../../.gitbook/assets/image (290).png" alt=""><figcaption></figcaption></figure>

## Repositories tested in CI/CD report

The Repositories tested in CI/CD report analyzes Snyk tests executed in your CI/CD pipelines using the Snyk CLI. Use this report to track testing adoption across your company, ensure Snyk tests your repositories, and prevent critical vulnerabilities and misconfigurations from reaching production.

This report provides visibility into Snyk tests (`snyk test`, `snyk code test`, `snyk container test`, `snyk iac test`) executed in your CI pipeline. Use these results to evaluate security checks and determine whether to pass or fail the build process.

{% hint style="info" %}
The report does not include `snyk monitor` commands. While `snyk monitor` identifies new vulnerabilities and maintains your security posture, this report tracks only the tests that actively gate your CI/CD pipeline.
{% endhint %}

### Prerequisites

Before you use this report, ensure you meet the following requirements:

* Snyk CLI version 1.1292.1 or later.
* SCM Group integration to view the last commit data. For more information, see SCM integrations.
* Include the `.git` context in the `snyk container test` command when you test containers.

### Report filters and metrics

By default, the report filters results to the last 30 days. You can adjust the date range filter to review specific periods, filter the data by specific products or organizations, or extend the viewed period.

The main view displays the number of repositories tested per Snyk product within the selected date range. You can also view the change in tested repositories compared to the previous sequential period to measure CI/CD testing adoption.

A green upward arrow indicates an increase in tested repositories compared to the previous period, and a red downward arrow indicates a decrease. The absolute change value appears next to the arrow, with the percentage of change below it.

{% hint style="info" %}
A sequential period is a date range covering the last seven days. The previous sequential period covers the seven days before that, ensuring both periods have the same duration.
{% endhint %}

#### Repository Test Adoption <a href="#repository-test-adoption" id="repository-test-adoption"></a>

The **Repository Test Adoption** trend shows testing adoption over time. A green line represents the weekly number of tested repositories. A purple line represents repositories with commits in the last 30 days.

Compare these lines to determine if CI/CD testing adoption is increasing and to identify repositories that received commits but lack CI/CD tests.

You can filter by specific products or by specific organizations or extend the viewed period using the date range filter.

{% hint style="info" %}
You must configure SCM Group integration to view the last commit data. For more information, see [SCM integrations](../../../developer-tools/scm-integrations/group-level-integrations/).
{% endhint %}

#### Test Success Rate Trend <a href="#test-success-rate-trend" id="test-success-rate-trend"></a>

The test success rate indicates how well your engineering department or specific Snyk Organizations adopt a shift-left approach to resolve issues before the build process. Snyk calculates this rate by dividing the number of passed tests by the total number of applicable tests. An applicable test is one that did not fail due to technical issues or an unsupported Project.

A low success rate can indicate the following:

* Snyk tests fail due to security issues you can prevent in local development or pull request checks. Snyk recommends you test with the Snyk IDE plugin, use Snyk PR Checks, and enroll in a Snyk Learn program.
* The test success criteria are too strict. Review the test definitions for organizations with the lowest success rates in the **Adoption by Organizations** widget. For more information about defining test success criteria, see [Failing of builds in Snyk CLI](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/failing-of-builds-in-snyk-cli.md).

#### Adoption by Organizations <a href="#adoption-by-organizations" id="adoption-by-organizations"></a>

Launching an application security program to boost CI/CD testing adoption requires collaboration between AppSec and R\&D teams, gradual implementation, and regular progress monitoring.

Use the **Adoption by Organization** table to track and compare adoption rates across Snyk Organizations and identify those lagging behind. Review the **Success Rate** column to find organizations with lower success rates.

The table includes the following columns:

* **Tested Repositories**: The number of repositories tested in the selected time range, including the percentage of change compared to the previous sequential period.
* **Committed Repositories**: The number of repositories with commits in the last 30 days during the selected time range, including the percentage of change compared to the previous sequential period.
* **Success Rate**: The percentage of successful CI/CD tests compared to all executed tests.

#### Repository Test Summary <a href="#repository-test-summary" id="repository-test-summary"></a>

The **Repository test summary** table displays the tests performed during the selected date range.

By default, the table sorts repositories by their last commit. This helps you identify repositories expected to undergo CI/CD testing and verify Snyk tested them. Click a column name to sort the table by that column. You can sort by multiple columns simultaneously.

{% hint style="info" %}
You must configure SCM Group integration to view the last commit data. For more information, see [Group-level integrations](../../../developer-tools/scm-integrations/group-level-integrations/).
{% endhint %}

You can execute tests on a specific repository branch from the table. The **TESTED** indicator shows that Snyk tested at least one branch of this repository during the selected date range. Hover over the **TESTED** tag to view the last test performed during the selected date range.

## Snyk Generated Pull Requests report

{% hint style="info" %}
**Feature availability**

Snyk Generated Pull Requests report is available only for Enterprise plan customers, for all SCM integrations. For more information, visit [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

### Access the report

The Generated Pull Requests report can be accessed at both Group and Organization level from the **Change Report** drop down in the Reports menu.

This report type provides an overview of how [Fix](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/enable-automatic-fix-prs.md), [Backlog](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/enable-automatic-backlog-prs-for-previously-known-vulnerabilities.md), and [Upgrade PRs](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/enable-automatic-upgrade-prs-for-new-dependency-upgrades.md) are used and highlights the efficiency of PR merges.

{% hint style="warning" %}
This report type does not include PR checks.
{% endhint %}

The analytics report covers the following:

* Overview of PRs status by type and the PR merge ratio.
* Visibility of issues.
* Breakdown by repository for PR status.

The report summary enables you to check the total number of Snyk PRs created, the total pull requests merged, and the mean time to merge for those pull requests.

### Report features

Use the date filter in the upper-right corner of the report to display data for a specific interval.

Add various filters to narrow down results to specific configurations. The filter options are Organization, SCM, Project, and Repository.

#### Snyk Generated Pull Requests usage

Pull Request usage is visualized in a **Pull requests by type** graph and a **Pull requests by status** table, displaying the same data in different formats. These distinguish the number of PRs into Fix, Backlog, and Dependency upgrade categories, segmented by Open, Merged, and Closed status types. Merge rate is presented as a percentage for each row.

#### Open vs Fixed issues

The Open vs Fixed issues in Snyk PRs graph and table displays the number of open and fixed issues based on severity.

#### Snyk Generated Pull Requests by repository

The **Projects/Org/Repository** table displays the number of Total, Open, Merged, and Closed PRs for each Organization and repository relationship. Merge rate is presented as a percentage for each row.

Select a repository name to open a modal containing additional metrics for that specific repository.

The repository breakdown details the number of PRs segmented by PR type and PR status. Merge rate is presented as a percentage for each row. It also lists the Projects within that repository, with the number of issues categorised by severity.
