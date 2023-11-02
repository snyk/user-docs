# Use Snyk Code in the CI/CD pipeline

Use a CI/CD integration to test your code for vulnerabilities and ensure your changes do not introduce new vulnerabilities, keeping your applications secure.

When you set up your CI/CD integration, note the following:

* Snyk Code is not yet supported in the Snyk CI plugins, for example, the Snyk Jenkins plugin, but you can use the [Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/using-snyk-code-from-the-cli/) to integrate with your CI server.
* You can [filter the results](broken-reference) by severity, for example, fail jobs only when high-severity vulnerabilities are introduced.
* You can [export the CLI output](broken-reference) to JSON or SARIF standard formats.
* You can generate more visual results using the [Snyk-to-HTML](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html/) tool.

To integrate Snyk Code into your CI/CD pipeline, see [Snyk CI/CD integrations](../../integrations/snyk-ci-cd-integrations/).
