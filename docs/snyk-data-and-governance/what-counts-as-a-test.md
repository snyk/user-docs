# What counts as a test?

{% hint style="info" %}
The information on this page applies to Free and Team plans only.

Capitalized terms used but not defined herein shall have the meaning as set forth in the Customer’s purchase agreement or other applicable documentation found on [snyk.docs.io](http://snyk.docs.io/).
{% endhint %}

Snyk keeps separate test counts and sets limits for each Snyk product: Snyk Open Source, Snyk Code, Snyk Container, and Snyk IaC.

If you are on the Free Snyk plan, you may run unlimited tests for public repositories, and limited tests on private repositories. Recurring tests may only be run on a weekly basis. See [Plans and pricing](https://snyk.io/plans?_gl=1*1d4rb1a*_ga*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.*_ga_X9SH3KP7B4*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w) for more details about Snyk plans. For information about unlimited tests against public repositories, see [Running out of tests](../developer-tools/snyk-cli/getting-started-with-the-snyk-cli.md#running-out-of-tests). If you reach your limit or would like to increase your recurring test frequency, you can [upgrade your plan](https://snyk.io/plans?_gl=1*1d4rb1a*_ga*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.*_ga_X9SH3KP7B4*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w).

The Snyk Open Source, Snyk Code, Snyk IaC, and Snyk Container applications allow customers to scan and run tests on their code-based assets as applicable based on the functionality of the application. The Customer’s Order Form indicates a plan type that comes with a certain number of tests as part of the Customer’s Subscription Allocation (Tests).

This document outlines what Snyk counts as a test, in order for the customer to understand its usage against its subscription allocation. Currently, test limits are focused on Snyk Open Source and Snyk Code Applications only, as is the discussion of test limits in this document.

There are two main types of tests:

* Recurring: Tests are triggered by the Snyk application, based on the customer’s configurations, and occur at a set cadence (daily or weekly). These tests are triggered by the Web UI, CLI, or API and implemented through a cron job, typically within the SCM.\
  \
  Note that if you have chosen the weekly retesting interval, and a manual test occurs before one week ends, the next test will occur a week after the manual test. Re-tests all begin at a set time each day. These re-tests may take sufficient time to process, such that a particular test takes place on the eighth day after the last test. If you see a test interval that is longer than you expect and you are concerned, contact Snyk support.
* Manual: Tests are triggered by the Customer through a specific election within the application. These tests can occur at any cadence within the available functionality of the application. These tests can be triggered in a number of different ways, including:
  * API - triggered by API call
  * CLI - triggered by CLI commands
  * IDE - triggered by save or autosave (may vary by IDE)
  * Pull request test or check - triggered by generation of a new Pull Request
  * Push tests - triggered by the customer's SCM
  * Web UI Import or retest - triggered by a button in WebUI

Each customer’s specific usage and configurations are different from others; therefore Snyk uses the criteria described here to determine what constitutes a test.

The following explains how Snyk determines the number of tests for each application:

* Snyk Open Source: the number of manifest files where vulnerabilities are identified by the application; note that one repository can have many manifest files.
* Snyk Code: the number of repositories scanned by the customer in the application.

The following are examples of ways the customer may initiate a test in the Snyk applications:

| Term                 | Definition                                                                                                                                                                                                                                                                                                                                                                                                                         | Example                                                                                                                                                                                                 |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| API                  | Using the Application Programming Interface to integrate programmatically with Snyk                                                                                                                                                                                                                                                                                                                                                | Using API endpoints: `https://api.snyk.io/rest`, `https://api.snyk.io/v1/test` endpoint                                                                                                                 |
| PR                   | Submitting a Pull Request (PR) within Source Control Manager (SCM)                                                                                                                                                                                                                                                                                                                                                                 | Test triggered using PR within Github                                                                                                                                                                   |
| Push                 | Push test triggered by customer SCM                                                                                                                                                                                                                                                                                                                                                                                                | Customer using Github as SCM and Jenkins as CI/CD. Customer creates a cron job within Jenkins to run at specific intervals, and Jenkins pulls the latest changes from Github to run predefined scripts. |
| Recurring            | Test triggered by the Snyk application based on the dustomer’s configurations and occurring at a set cadence (daily or weekly). These tests are triggered by Web UI, CLI, or API and implemented through a cron job typically within the SCM.                                                                                                                                                                                      | Using Snyk Github integration, customers can set daily or weekly tests.                                                                                                                                 |
| Retest               | Using the retest button within the Snyk web app                                                                                                                                                                                                                                                                                                                                                                                    | User clicks retest within the Snyk web app.                                                                                                                                                             |
| Snyk monitor command | Using the CLI to create a Project in your Snyk account to be continuously monitored for open-source vulnerabilities and license issues, sending the results to snyk.io. This applies only to Snyk Open Source and Snyk Container.                                                                                                                                                                                                  | User runs `snyk monitor` in the CLI                                                                                                                                                                     |
| Snyk test commands   | <p>Using the CLI to check Projects for open-source vulnerabilities and license issues. The <code>test</code> command tries to auto-detect supported manifest files with dependencies and test those.</p><p><br>Note: There are specific <code>snyk test</code> commands for the Snyk Code, Container, and IaC scanning methods: <code>snyk code test</code>, <code>snyk container test</code>, and <code>snyk iac test</code>.</p> | User runs `snyk test` in the CLI                                                                                                                                                                        |
| User                 | Tests triggered by importing repositories                                                                                                                                                                                                                                                                                                                                                                                          | In the Snyk web app, when user clicks import button                                                                                                                                                     |
| IDE                  | Integrated development environment, VS Code, JetBrains, and so on                                                                                                                                                                                                                                                                                                                                                                  |                                                                                                                                                                                                         |

## Counting Git repository integration scans

These Snyk features for [Git Repositories (SCM)](../developer-tools/scm-integrations/organization-level-integrations/) integrations run scans automatically by default:

* Daily recurring tests
* An automatic scan, which runs if the dependencies change on your default branch
* PR checks, which run when you create a pull request that changes those dependencies

If you have a Dockerfile in your source code repository, the default settings will detect and scan it, but Dockerfiles count as a Snyk Container scan, not a Snyk Open Source scan.

Terraform and Kubernetes configuration files scanned from source code repositories are counted as Snyk IaC scans.

For container scans from a registry or your Kubernetes cluster, Snyk counts the initial scan and subsequent recurring scans. By default, recurring scans run once a day.

## Counting recurring scans

Snyk periodically checks whether your code is affected by newly disclosed vulnerabilities.

The test frequency is set to a default for each product. For information about changing the frequency, see [Usage settings](../snyk-platform-administration/groups-and-organizations/usage-settings.md), [View and edit Project settings](../snyk-platform-administration/snyk-projects/view-and-edit-project-settings.md), and [Test frequency settings](../snyk-platform-administration/snyk-projects/#test-frequency-settings) on the Snyk Projects page.

## Counting CLI tests

A test is counted each time you run one of the following commands:

* For Snyk Open Source: `snyk test` or `snyk monitor`.
* For Snyk Container: `snyk container test` or `snyk container monitor.`
* For Snyk Code: `snyk code test`.

For Snyk IaC, the command is `snyk iac test`. Since this can scan multiple Projects, a scan is counted for every Project being scanned. For example, If a `snyk iac test` command scans 11 Projects, the count is increased by 11.

## Counting app-based tests

A scan runs when you add a new Project or click the re-test button. This is in addition to any automated tests that run.

## Counting API tests

Tests are counted when calls are made to the `https://api.snyk.io/v1/test` endpoint.

## Test usage policy

Snyk may monitor customer test volumes on a daily basis and actively review customer usage on a rolling thirty (30) day period. If a customer’s test usage exceeds the limit granted by the plan purchased by twenty percent (20%) on a rolling ninety (90) day period or by one-hundred percent (100%) for a period of thirty (30) days, Snyk may notify the customer to discuss the overage and, if action is required, provide an expansion invoice to increase the test allocation on the subscription or ask the customer to reduce its use. Except in unusual circumstances, Snyk does not invoice retroactively for test overages.
