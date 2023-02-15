# Using Snyk Code from CI/CD pipeline

Use CI/CD integration to test your code for vulnerabilities and make sure that your changes do not introduce new vulnerabilities and keep your applications secure.

Before you set up your CI/CD integration, please be aware of the following:

* Snyk Code is not yet supported in the Snyk CI Plugins (for example, the Snyk Jenkins plugin), but you can use the [Snyk Code CLI](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code) to integrate with your CI server.
* You can [filter the results](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/displaying-only-discovered-issues-above-a-specific-severity-level) by severity (for example, only fail jobs if high-severity vulnerabilities are introduced).
* You can [export the CLI output](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-a-json-or-sarif-format-in-the-terminal) to JSON or SARIF standard formats.
* You can generate more visual results using the [Snyk-to-HTML](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature) tool.

To integrate Snyk Code into your CI/CD pipeline, see [Snyk CI/CD integrations](../../integrations/ci-cd-integrations/).
