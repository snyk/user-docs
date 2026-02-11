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

Snyk Pull request checks usage & performance report is in Early Access and available only for Enterprise plan customers. For more information, visit [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

Monitor the use and results of Snyk Pull Request Checks to verify that critical vulnerabilities are identified before merging.

## Repositories tested in CI/CD report

To use this report, consider the following prerequisites:

* Snyk CLI version 1.1292.1 or newer.
* Viewing the last commit data requires SCM Group integration. For more details, navigate to [SCM integrations](../../../developer-tools/scm-integrations/organization-level-integrations/).
* When testing containers, include the `.git` context as part of the `snyk container test` command.

This report analyzes Snyk tests performed as part of CI/CD pipelines executed using Snyk CLI. It will inform you about the usage of your company and adoption of testing in CI/CD, ensuring repositories are tested as expected and preventing critical vulnerabilities and misconfigurations from being deployed and reaching the production environment.

The report results are scoped by a date range filter that you can use to review specific periods. The filter is defaulted to the last 30 days.

This report provides visibility into Snyk tests (`snyk test`, `snyk code test`, `snyk container test`, `snyk iac test`) executed within your CI pipeline (using CLI). Its primary goal is to help you evaluate test results and determine whether to pass or fail the build process based on these security checks.

{% hint style="info" %}
&#x20;`snyk monitor` commands are **not** included in this report. While `snyk monitor` is crucial for ongoing security posture and identifying new vulnerabilities, this report specifically tracks tests that actively gate your CI/CD pipeline.
{% endhint %}

The numbers displayed on the main view of the report represent the number of repositories tested in the selected date range per Snyk product.

In addition, you can learn about the change in the number of tested repositories compared to the previous sequential period, so you can conclude whether the adoption of CI/CD tests across repositories improved.

A green upward arrow indicates that more repositories were tested compared to the previous sequential period, while a red downward arrow indicates the opposite. The absolute change value appears next to the arrow, and the perception of change appears right underneath to measure the degree of change.

{% hint style="info" %}
A sequential period refers to a date range covering the last seven days. In this case, the period starts seven days ago and ends today. The previous sequential period spans from 14 days ago to seven days ago. As a result, both sequential periods are of the same duration.
{% endhint %}

#### Repository Test Adoption <a href="#repository-test-adoption" id="repository-test-adoption"></a>

Review the Repository Test Adoption trend to learn more about ‌adoption over time.\
Represented by the green line, you can see the weekly number of repositories that have been tested compared to the repositories that had commits in the last 30 days, represented by the purple line.

This comparison helps determine whether Snyk tests in CI/CD are being increasingly adopted over time and highlights the number of repositories that have received commits but have not been tested in CI/CD.

{% hint style="info" %}
Viewing the last commit data requires SCM Group integration. For more details, visit the [SCM integrations](../../../developer-tools/scm-integrations/organization-level-integrations/) page.
{% endhint %}

You can filter by specific products or by specific organizations or extend the viewed period using the date range filter.

#### Test Success Rate Trend <a href="#test-success-rate-trend" id="test-success-rate-trend"></a>

The test success rate serves as an indicator of how well the engineering department or specific Snyk Organizations can adopt a "shift left" approach, which aims to identify and resolve issues before the code reaches the build process. This success rate is calculated by dividing the number of tests that passed by the total number of relevant tests conducted.

{% hint style="info" %}
An applicable test is a test that did not fail due to technical issues or a non-supported Project.
{% endhint %}

Having a low success rate can indicate that:

* Snyk tests are failing due to security issues that can be prevented in local development or in the PR Check stages. Snyk recommends testing with the [Snyk IDE](../../../developer-tools/snyk-ide-plugins-and-extensions/) plugin, using [Snyk PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/) and enroll in a [Snyk Learn](../../../discover-snyk/snyk-learn/) program.
* The test success criteria are too strict. To explore this option further, Snyk recommends reviewing the test definitions of the organizations with the lowest success rate, as shown by the Adoption by Organizations widget. For more details about defining test success criteria, navigate to the [Failing of builds in Snyk CLI](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/failing-of-builds-in-snyk-cli.md) page.

#### Adoption by Organizations <a href="#adoption-by-organizations" id="adoption-by-organizations"></a>

Launching an Application Security program to boost testing adoption in CI/CD pipelines can be challenging. This initiative requires collaboration between the AppSec and R\&D teams and will be implemented gradually, with regular progress monitoring.

The Adoption by Organization table enables you to track and compare the Snyk Organizations adoption rates, helping you identify those that are struggling or lagging behind.

In addition, you can examine the success rate column to surface organizations that have lower success rates.

**Columns descriptions:**

* **Tested Repositories:** the number of repositories that were tested in the selected time range, with an indication of the percentage of change compared to the previous sequential period.
* **Committed Repositories:** the number of repositories that had any commits in the last 30 days at any given time in the selected time range, with an indication of the percentage of change compared to the previous sequential period.
* **Success Rate:** the portion of successful tests in CI/CD against all other tests that were executed.

#### Repository Test Summary <a href="#repository-test-summary" id="repository-test-summary"></a>

The repository test summary table shows the performed tests during the selected date range.

The default sorting in the table surfaces repositories according to their last commit, allowing you to identify repositories that were expected to be tested in CI/CD pipelines and verify they were tested. Clicking a column name sorts the table to that column. You can sort the table by multiple columns at a time.

{% hint style="info" %}
Viewing the last commit data requires SCM Group integration. For more details, visit the [Group-level integrations](../../../developer-tools/scm-integrations/group-level-integrations/) page.
{% endhint %}

You can execute the test on a specific repository branch in the table. The `tested` indicator means that any branch of this repository was tested during the selected date range.

{% hint style="info" %}
Hovering over the TESTED tag reveals the last test performed during the selected date range
{% endhint %}

## Snyk Generated Pull Requests report

{% hint style="info" %}
**Feature availability**

Snyk Generated Pull Requests report is available only for Enterprise plan customers, for all SCM integrations. For more information, visit [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

### Access the report

The Generated Pull Requests report can be accessed at both Group and Organization level from the **Change Report** drop down in the Reports menu.

This report type provides an overview of how [Fix](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md), [Backlog](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md), and [Upgrade PRs](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/) are used and highlights the efficiency of PR merges.

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
