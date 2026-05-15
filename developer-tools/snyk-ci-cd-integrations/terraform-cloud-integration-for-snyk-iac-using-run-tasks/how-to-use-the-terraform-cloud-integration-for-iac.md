# How to use the Terraform Cloud integration for IaC

After your integration is set up, Snyk scans Terraform plans for each run triggered in your workspace.

## View Terraform plan scanning results

1. For each run triggered in the Terraform Cloud workspace, the result of Snyk Terraform plan scanning appears under the `run tasks` step, which starts after the `Plan` stage finishes.
2. The scan results in either a `passed` or a `failed` status. If Snyk finds issues in your Terraform plan file, the scan results in a failure.
3. Click on the **Details** link of the run task results in Terraform Cloud to view further details in Snyk.
   1. You can also find the results under the **Projects** tab in Snyk by searching for terraform-plan.json, which is under a **Target** named by `{YOUR_TFC_ORG_NAME}/{YOUR_TFC_WORKSPACE_NAME}`
   2. You can also use the filter in the left pane to show only Terraform Cloud Projects.
4. A single Project in Snyk (`terraform-plan.json`) is created for each workspace that uses the Snyk integration. Every Project page shows the latest scanning results.
5. To see historical scan results, navigate to the **History** tab under the relevant Project and choose the historic snapshot you wish to view.

## Customize Terraform plan scanning

The Snyk Terraform Cloud integration provides the following levels of customization:

* **Severity Threshold**: Set the minimum level of severity for failure. This can be set on the integration page in Snyk.
* **Custom Severities**: Set custom severities for issues that overwrite the defaults (for example, [SNYK-CC-00172](https://security.snyk.io/rules/cloud/SNYK-CC-00172)).
* **Enforcement Level**: Determine whether a failure blocks the `apply` or not. This setting is controlled through Terraform Cloud. For example, the `Advisory` level does not block the `apply` even if Snyk finds issues within the minimum severity threshold.

## Notes and limitations

* Snyk receives an event from Terraform Cloud for each `plan` stage finished within the latest run in Terraform Cloud.
* The only way to trigger a scan is through Terraform Cloud by triggering a new run.
* You cannot trigger a re-scan of the Terraform plan file through the Snyk Web UI.
* If you customize the Snyk integration, for example, change severity threshold or customize policy severities, you must trigger a new run in Terraform Cloud for the changes to take effect in Snyk.
