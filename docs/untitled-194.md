# Scan and fix security issues in your Helm Charts

* [ Configure your integration to find security issues in your Kubernetes configuration files](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360006402818-Configure-your-integration-to-find-security-issues-in-your-Kubernetes-configuration-files/README.md)
* [ Scan and fix security issues in your Kubernetes configuration files](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360006368877-Scan-and-fix-security-issues-in-your-Kubernetes-configuration-files/README.md)
* [ Scan and fix security issues in your Helm Charts](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360007673117-Scan-and-fix-security-issues-in-your-Helm-Charts/README.md)
* [ Working with your Kubernetes configuration file test results](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360006369437-Working-with-your-Kubernetes-configuration-file-test-results/README.md)
* [ Detecting Kubernetes configuration files using a broker](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360010797537-Detecting-Kubernetes-configuration-files-using-a-broker/README.md)

## Scan and fix security issues in your Helm Charts

Snyk scans Helm Charts, in addition to Kubernetes configuration files, for misconfigurations and security issues. Once Helm Charts are scanned, Snyk creates projects for each template and dependency template, generates reports on any misconfigurations, and makes recommendations for fixing them.

### Prerequisites

* An administrator should connect your organization with your preferred Git repository and enable detection of configuration files as described [here](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/articles/360006402818/README.md#UUID-c1919782-6bfa-b84b-a638-3913cee39fc5).
* The repository should follow the [standard Chart directory structure](https://helm.sh/docs/topics/charts/#the-chart-file-structure). Specifically we look for:
  * `Chart.yaml` - YAML file containing information about the chart
  * `values.yaml` - The default configuration values for this chart
  * `templates/` - A directory of templates that, when combined with values will generate valid Kubernetes manifest files.
  * `Values` - An optional directory of values that used to configure different environments.
  * `requirements.yaml` - Optional file - Additional way to declare dependencies

### Scan and fix your Charts

1. 1. If you already imported your repositories for testing before cloud configuration file detection was enabled by your administrator, then you should re-import that repository again in order to import the Helm chart: 
2. Every time a repository is scanned:
   * Each template in your Helm Chart creates a Snyk a project, grouped together by repository, similar to this example:
   * If you re-imported the repository in order to import the cloud configuration files, then Snyk imports and tests the configuration files and also re-tests the already imported application manifest files - displaying the test time as "now".
3. Click the project link you're interested in, to view the scan results and to correct your configuration files accordingly:
   * Projects that were created from external dependencies will also be scanned and issues shown.

