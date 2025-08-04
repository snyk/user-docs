# Snyk Code in the CI/CD pipeline

You can use CI/CD integration to test your code for vulnerabilities and ensure your changes do not introduce new vulnerabilities, keeping your applications secure.

* Snyk Code may not be supported by a Snyk CI/CD plugin, such as the Snyk Jenkins plugin. If this is the case, you can integrate [Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/) with your CI server.&#x20;
* You can filter the results by severity, for example, fail jobs only when high-severity vulnerabilities are introduced. See [Filter results by Severity](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#filter-results-by-severity-level).
* You can export the CLI output to JSON or SARIF standard formats. See [Output test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results).
* You can generate more visual results using the Snyk-to-HTML tool. See [`snyk-to-html`](../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
