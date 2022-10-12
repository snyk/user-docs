# Scan and fix security issues in Helm Charts

In addition to scanning Kubernetes configuration files for misconfigurations and security issues, Snyk has support for templating Helm charts and scanning the resultant manifests. This templating functionality is only available when importing repositories via the Snyk UI. Please see the below sections for prerequisites and guidance on how to scan templated Helm charts using the Snyk CLI. Once Helm charts are scanned, Snyk creates projects for each template and dependency template, generates reports on any misconfigurations, and makes recommendations for fixing them.

## Prerequisites

* An administrator should [connect your organization](../scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-filess.md) with your preferred Git repository and enable detection of configuration files as described.
* The repository should follow the [standard Chart directory structure](https://helm.sh/docs/topics/charts/#the-chart-file-structure).
* We currently only support templating Helm charts using the default values file, `values.yaml`. If you want to scan particular configurations of Helm values, then the supported workflow is to template the chart outside of Snyk and scan the manifests as regular Kubernetes files.
  * Helm charts that cannot be templated from their default values file are currently unsupported.
* Any chart dependencies must either be publicly downloadable from the configured Helm repository, or be found in the same git repository as the chart under test.

## Scan and fix your Charts

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. If you already imported your repositories for testing before cloud configuration file detection was enabled by your administrator, then you should re-import that repository again in order to import the Helm chart:
3. Every time a repository is scanned:
   1. Snyk creates a project for each template in your Helm Chart, grouped together by repository.
   2. If you re-imported the repository in order to import the cloud configuration files, then Snyk imports and tests the configuration files and also re-tests the already imported application manifest files - displaying the test time as "now".
4. Click the project link you're interested in, to view the scan results and to correct your configuration files accordingly.
   1. Projects that were created from external dependencies will also be scanned and issues shown.

![](<../../../.gitbook/assets/image (316) (1).png>)

## Testing custom Helm values configurations

Sometimes, testing a chart using only the default values isn’t enough. Snyk does not currently support passing custom values into imports. This section is intended to offer guidance on how to template custom configurations outside of Snyk, and scan the resultant Kubernetes manifests.

You can use the Snyk CLI and Helm in conjunction:

```bash
cd /path/to/helm/chart
helm dependency update
helm template . --output-dir out
snyk iac test out/
```

You can pass standard Helm values flags (e.g. `--set` and/or -`-values`) to `helm template` in order to test a non-default configuration.

You can script this process and run it in a CLI pipeline, or alternatively helm-template files into a repository that can be imported into Snyk as projects.

To share CLI results with the Snyk Web UI, use the `--report` flag. For example:

```
snyk iac test out/ --report
```
