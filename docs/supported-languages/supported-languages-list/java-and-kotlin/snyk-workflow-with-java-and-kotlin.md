# Snyk workflow with Java and Kotlin

The Snyk team has built plugins to integrate Snyk into your workflows:

* [Gradle Plugin](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) (Community project)
* [Maven Plugin](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)

## Validating, monitoring, alerting, and gating

The following capabilities are available for all Snyk users:

### **With Git integrations**

Snyk allows you to [run PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis, and show results.

These results are viewable on the Snyk projects screen, for:

* Your code with Snyk Code
* Open Source with Snyk Open Source
  * Check for known vulnerabilities (Snyk Open Source)
    * Create Fix Pull Requests to fix known vulnerabilities (Maven)
  * License compliance checks (Snyk Open Source)(Maven)
  * Dependency upgrade - positioning updates to address technical debt (Snyk Open Source) (Maven)

With the Git Integration, you can monitor the following on a daily basis:

* Infrastructure as code (IaC) with Snyk Infrastructure as Code

### **With CI/CD integrations**

Snyk can passively monitor and provide a QA gate by failing build checks during testing for policy violations.

Snyk provides flexible capabilities, including:

* [Gradle Plugins](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) (Community project)
* [Maven Plugins](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)
* Dedicated plugins for Jenkins, Circle CI, and others (see relevant marketplaces)
* Using [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/)
* The Snyk CLI can be used in most CI/CD systems (see [examples](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * Fail the build based on criteria using options or the [snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md) tool
  * There are [containerized](https://hub.docker.com/r/snyk/snyk) versions available
* With Partner Platforms: Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk.
  * Note for Java: using the Git integration with Bitbucket Cloud or using the CLI instead of the prepackaged Bitbucket Pipe is suggested.

## Production monitoring

* (Snyk Enterprise plan only) Snyk can monitor container images and their open source or Linux based packages being used in production using Kubernetes integration, to notify customers of known vulnerabilities for applications in production.
* (All plans) Where a production integration does not exist, use the [snyk monitor](../../../developer-tools/snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.
