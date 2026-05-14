# Scan and fix security issues in Helm Charts

In addition to scanning Kubernetes configuration files for misconfigurations and security issues, Snyk has support for templating Helm charts and scanning the resultant manifests. This templating functionality is only available when you import repositories using the Snyk UI. See the following sections for prerequisites and guidance on how to scan templated Helm charts using the Snyk CLI. After Helm charts are scanned, Snyk creates Projects for each template and dependency template, generates reports on any misconfigurations, and makes recommendations for fixing them.

## Prerequisites for scanning and fixing issues in Helm Charts

* An administrator must connect your Organization with your preferred Git repository and enable the detection of configuration files as [explained for Terraform files](../scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-files-current-iac.md#configure-snyk-to-scan-your-terraform-configuration-files).
* The repository should follow the [standard Chart directory structure](https://helm.sh/docs/topics/charts/#the-chart-file-structure). However the `charts/` sub-directory is not supported.
* Snyk supports only templating Helm charts using the default values file, `values.yaml`.
  * If you want to scan particular configurations of Helm values, the supported workflow is to template the chart outside of Snyk and scan the manifests as regular Kubernetes files.
  * Helm charts that cannot be templated from their default values file are not supported.
* Any chart dependencies must be publicly downloadable from the configured Helm repository. Subcharts or non-publicly downloadable dependencies are not supported. The supported workflow in these cases is to template the chart outside of Snyk and scan the manifests as regular Kubernetes files.

## Scan and fix your Helm Charts

* Log in to your account and navigate to the relevant Group and Organization that you want to manage.
* If you imported your repositories for testing before configuration file detection was enabled by your administrator, you must re-import that repository in order to import the Helm chart:

Every time a repository is scanned, Snyk creates a Project for each template in your Helm Chart, grouped together by repository.

If you re-imported the repository in order to import the configuration files, then Snyk imports and tests the configuration files and also re-tests the already imported application manifest files, displaying the test time as "now".

* Click the Project link you are interested in to view the scan results and to correct your configuration files accordingly.\
  Projects that were created from external dependencies will also be scanned, and issues will be shown.

<figure><img src="../../../../.gitbook/assets/image (114).png" alt="Helm Charts Project detail"><figcaption><p>Helm Charts Project detail</p></figcaption></figure>

## Templating charts and scanning Kubernetes manifests

Sometimes testing a chart using only the default values is not enough, or a dependency is either a subchart or not publicly downloadable. Snyk does not support these scenarios during imports. This section is intended to offer guidance on how to template custom configurations outside of Snyk and scan the resultant Kubernetes manifests.

You can use the Snyk CLI and Helm in conjunction:

```bash
cd /path/to/helm/chart
helm dependency update
helm template . --output-dir out
snyk iac test out/
```

You can pass standard Helm values flags (for example, `--set`, `--value`, or both) to `helm template` in order to test a non-default configuration.

You can script this process and run it in a CLI pipeline or helm-template files into a repository that can be imported into Snyk as Projects.

To share CLI results with the Snyk Web UI, use the `--report` CLI option. For example:

```
snyk iac test out/ --report
```
