# What counts as a test?

{% hint style="info" %}
The information on this page, including any reference to “test limits”, does not apply to capabilities for which customers are on a credit-based license. For the credit and usage policies applicable to these customers, see [Snyk Platform Access Plan](https://docs.snyk.io/snyk-data-and-governance/snyk-platform-access-credits) (for credit-based licenses purchased on or before 31 December 2025), or [Platform Credit Consumption Licensing ](https://snyk.io/policies/credit-based-billing/)(for credit-based licenses purchased on or after 1 January 2026).

Capitalized terms used but not defined herein shall have the meaning as set forth in the Customer’s purchase agreement or other applicable documentation found on [snyk.docs.io](http://snyk.docs.io/).
{% endhint %}

The Snyk Open Source, Snyk Code, Snyk IaC, and Snyk Container capabilities entitle customers to run security tests on their code-based assets as applicable based on the functionality of each capability.

This document outlines what constitutes a test, in order for the customer to understand usage against the subscription allocation.

There are two main types of tests:

* Recurring: Tests are triggered by the Snyk application, based on the customer’s configurations, and occur at a set cadence (e.g. daily or weekly). These tests are triggered by the Web UI, CLI, or API and implemented through a cron job, typically within the SCM.\
  \
  Note that if you have chosen the weekly retesting interval, and a manual test occurs before one week ends, the next test will occur a week after the manual test. Re-tests all begin at a set time each day. These re-tests may take sufficient time to process, such that a particular test takes place on the eighth day after the last test. If you see a test interval that is longer than you expect and you are concerned, contact Snyk support.
* Manual: Tests are triggered by the customer through a specific action. These tests can occur at any cadence within the available functionality of the capability. These tests can be triggered in a number of different ways, including:
  * API - triggered by API call
  * CLI - triggered by CLI commands
  * IDE - triggered by save or autosave (may vary by IDE)
  * Pull request test or check - triggered by generation of a new Pull Request
  * Push tests - triggered by the customer's SCM
  * Web UI Import or retest - triggered by a button in Web UI

Each customer’s specific usage and configurations are distinct; therefore Snyk uses the criteria described here to determine what constitutes a test.

The following explains how Snyk determines the number of tests for each capability:

* Snyk Open Source: the number of manifest files where vulnerabilities are identified by the application; note that one repository can have many manifest files. When a scan is conducted, each manifest file in scope will generate a distinct test.
* Snyk Code: the number of repositories scanned by the customer in the application.
* Snyk Container: the number of dockerfiles or other container files & container images scanned by the customer in the application.
* Snyk IaC: the number of terraform or other IaC files scanned by the customer in the application.

## Test limits and usage policy

Snyk maintains separate test counts and limits for each of the following products: Snyk Open Source, Snyk Code, Snyk Container, and Snyk IaC. For more details about the available Snyk plans and associated test limits, see [Plans and pricing](https://snyk.io/plans?_gl=1*1d4rb1a*_ga*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.*_ga_X9SH3KP7B4*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w).

The applicable test limits apply to private repositories only.&#x20;

Plan specific considerations:&#x20;

* Free Snyk Plan: test limits apply to all products including Snyk Container and Snyk IaC.&#x20;
* Team and Enterprise Snyk Plans: test limits only apply to Snyk Code and Snyk Open Source. Tests executed via Snyk’s IDE plugin do not count towards test limits.

If you reach your test limit or would like to increase your recurring test frequency, you can [upgrade your plan](https://snyk.io/plans?_gl=1*1d4rb1a*_ga*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.*_ga_X9SH3KP7B4*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w).
