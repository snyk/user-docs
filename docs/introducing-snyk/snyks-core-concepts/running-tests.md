# Running tests

You can use Snyk to test your code in multiple ways.

* Manually; using the [Snyk CLI](), the [Snyk app](), and the [Snyk API]().
* Snyk can also [run tests automatically](); on a recurring basis, or when a relevant repo change is made.

#### Run tests with the CLI

With our [**CLI**](https://snyk.io/docs/using-snyk): you can use the following commands:

* Scan open-source code with **snyk test**.
* Scan container images with **snyk container test**.
* Scan Infrastructure as Code \(IaC\) files with **snyk iac test**.

See [Getting started with the CLI](https://docs.snyk.io/snyk-cli/guides-for-our-cli/getting-started-with-the-cli) for details.

#### Run tests with the Snyk app

A test is run when you add a new project or click the re-test button. 

See [Getting started with Snyk products](https://support.snyk.io/hc/en-us/sections/360004349758-Getting-started-with-Snyk-products) for details.

#### Run tests with the API

Tests are counted when calls are made to the [**https://snyk.io/api/v1/test**](https://snyk.io/api/v1/test) endpoint.

See [API documentation](https://snyk.docs.apiary.io/#) for details.

### Run tests automatically

Snyk provides automatic scanning functions with recurring tests, allowing you to catch new vulnerabilities automatically. After you import a project, Snyk automatically runs periodic checks to see if your code is affected by newly disclosed vulnerabilities.

Test frequency is set to daily by default. To change frequency, go to either the **Usage** page \(see [Usage page details](https://docs.snyk.io/user-and-group-management/managing-settings/usage-page-details)\) or the project **Settings** page \(see [View project settings](https://docs.snyk.io/getting-started/introduction-to-snyk-projects/view-project-settings)\).

#### Snyk monitor

Use the `snyk monitor` CLI command to create a snapshot of a project on the Snyk website that will be continuously monitored for new vulnerabilities. 

See [Monitor your projects at regular intervals](https://docs.snyk.io/snyk-cli/secure-your-projects-in-the-long-term/monitor-your-projects-at-regular-intervals) for details.

#### Automatically testing PRs / MRs

By default, Snyk scans every pull request submitted on your monitored repositories, showing the results and recommendations grouped together in a single security check and a single license check.

See [Enable Snyk test on PRs](https://support.snyk.io/hc/en-us/articles/360018010597-Snyk-SCM-integration-good-practices#Stage3) for more details.

