# What counts as a test?

If you are on the free Snyk plan, you may run unlimited tests against open-source projects and limited tests on private Projects. See [Plans and pricing](https://snyk.io/plans?\_gl=1\*1d4rb1a\*\_ga\*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.\*\_ga\_X9SH3KP7B4\*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w) for the limits on both free and paid Snyk accounts.

If you reach your limit, you may want to [choose a plan](https://snyk.io/plans?\_gl=1\*1d4rb1a\*\_ga\*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.\*\_ga\_X9SH3KP7B4\*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w) to continue being able to run tests.

Snyk keeps separate test counts and sets limits for each product: Snyk Open Source, Snyk Code, Snyk Container, and Snyk IaC.

Tests on private Projects count towards your test limit; tests on public Projects do not count.

## Counting SCM integration tests

Several features enabled by default cause tests to be executed. These are:

* Daily recurring tests; see the following section.
* An automatic test, which executes when the dependencies change on your default branch.
* Pull request checks, which execute when you create a pull request that changes those dependencies.

If you have a Dockerfile in your source code repository, the default settings will detect and scan it, but Dockerfiles count as a Snyk Container test, not a Snyk Open Source test.

Terraform and Kubernetes configuration files scanned from source code repositories are counted as Snyk IaC tests.

For container scans from a registry or your Kubernetes cluster, Snyk counts the initial test and subsequent recurring tests. By default, these are run once per day.

## Counting recurring tests

Snyk periodically checks whether your code is affected by newly disclosed vulnerabilities.

The frequency is set to a default for each product. For information about changing the frequency, see [Usage settings](../snyk-admin/manage-settings/usage-settings.md) and [View and edit Project settings](../manage-issues/snyk-projects/view-and-edit-project-settings.md).

## Counting CLI tests

A test is counted each time you run one of the following commands:

* For Snyk Open Source: `snyk test` or `snyk monitor`.
* For Snyk Container: `snyk container test` or `snyk container monitor.`
* For Snyk Code: `snyk code test`.

For Snyk IaC, the command is `snyk iac test`. Since this can test multiple Projects, a test is counted for every Project being scanned. For example, If a `snyk iac test` command tests 11 Projects, the test count is increased by 11.

## Counting app-based tests

A test is run when you add a new Project or click the re-test button. This is in addition to any automated recurring tests which are run.

## Counting API tests

Tests are counted when calls are made to the `https://snyk.io/api/v1/test` endpoint.
