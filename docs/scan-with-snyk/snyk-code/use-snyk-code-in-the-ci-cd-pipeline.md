# Use Snyk Code in the CI/CD pipeline

Use a CI/CD integration to test your code for vulnerabilities and ensure your changes do not introduce new vulnerabilities, keeping your applications secure.

When you set up your CI/CD integration, note the following:

* Snyk Code is not yet supported in the Snyk CI plugins, for example, the Snyk Jenkins plugin, but you can use the [Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/) to integrate with your CI server.
* You can [filter the results](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#filter-results-by-severity-level) by severity, for example, fail jobs only when high-severity vulnerabilities are introduced.
* You can [export the CLI output](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results) to JSON or SARIF standard formats.
* You can generate more visual results using the [Snyk-to-HTML](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) tool.

To integrate Snyk Code into your CI/CD pipeline, see [Snyk CI/CD integrations](../../integrate-with-snyk/snyk-ci-cd-integrations/).
