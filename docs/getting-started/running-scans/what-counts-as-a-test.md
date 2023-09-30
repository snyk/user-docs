# What counts as a scan?

Snyk keeps separate scan (test) counts and sets limits for each Snyk product: Snyk Open Source, Snyk Code, Snyk Container, and Snyk IaC.

If you are on the free Snyk plan, you may run unlimited scans against public open-source Projects and limited scans on private Projects. See [Plans and pricing](https://snyk.io/plans?\_gl=1\*1d4rb1a\*\_ga\*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.\*\_ga\_X9SH3KP7B4\*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w) for more details. If you reach your limit, you can [upgrade your plan](https://snyk.io/plans?\_gl=1\*1d4rb1a\*\_ga\*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.\*\_ga\_X9SH3KP7B4\*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w).

## Counting Git repository integration scans

These Snyk features for [Git Repositories (SCM)](../../integrations/git-repository-scm-integrations/) integrations run scans automatically by default:

* Daily recurring scans
* An automatic scan, which runs if the dependencies change on your default branch.
* PR checks, which run when you create a pull request that changes those dependencies.

If you have a Dockerfile in your source code repository, the default settings will detect and scan it, but Dockerfiles count as a Snyk Container scan, not a Snyk Open Source scan.

Terraform and Kubernetes configuration files scanned from source code repositories are counted as Snyk IaC scans.

For container scans from a registry or your Kubernetes cluster, Snyk counts the initial scan and subsequent recurring scans. By default, these are run once per day.

## Counting recurring scans

Snyk periodically checks whether your code is affected by newly disclosed vulnerabilities.

The frequency is set to a default for each product. For information about changing the frequency, see [Usage settings](../../snyk-admin/manage-settings/usage-settings.md) and [View and edit Project settings](../../snyk-admin/introduction-to-snyk-projects/view-and-edit-project-settings.md).

## Counting CLI scans

A scan is counted each time you run one of the following commands:

* For Snyk Open Source: `snyk test` or `snyk monitor`.
* For Snyk Container: `snyk container test` or `snyk container monitor.`
* For Snyk Code: `snyk code test`.

For Snyk IaC, the command is `snyk iac test`. Since this can scan multiple Projects, a scan is counted for every Project being scanned. For example, If a `snyk iac test` command scans 11 Projects, the count is increased by 11.

## Counting app-based scans

A scan is run when you add a new Project or click the re-test button. This is in addition to any automated recurring scans which are run.

## Counting API scans

Scans are counted when calls are made to the `https://snyk.io/api/v1/test` endpoint.
