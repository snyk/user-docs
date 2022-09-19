# Share CLI results with the Snyk Web UI

### Introduction

You can use the [CLI](../../snyk-cli/) `snyk iac test` command to address known configuration issues.

To see these issues displayed in the Snyk Web UI, run the following CLI command:

`snyk iac test myproject --report`

### Example

```
> snyk iac test myproject --report

Testing arm-file.tf...


Infrastructure as code issues:
  ✗ VM Agent is not provisioned automatically for Windows [Low Severity] [SNYK-CC-AZURE-667] in Compute
    introduced by resource > azurerm_virtual_machine[my_terraformvm] > os_profile_windows_config > provision_vm_agent


Organization:      my.org
Type:              Terraform
Target file:       arm-file.tf
Project name:      myproject
Open source:       no
Project path:      myproject

Tested arm-file.tf for known issues, found 1 issues

Your test results are available at: https://app.snyk.io/org/my.org/projects under the name myproject
```

This sends a snapshot of your current configuration issues to your Snyk dashboard, to view in the Snyk Web UI.

### Viewing snapshots in the Snyk Web UI

Log in to the [Snyk Web UI](../../snyk-web-ui/) and go to your organization's project page, to see the specific snapshot of your scanned project.

![Newly scanned project in the projects page](<../../.gitbook/assets/image (349) (1).png>)

You can also open the project itself and see the usual individual project view:

![individual project view](<../../.gitbook/assets/image (395).png>)

### **Ignores**

You can ignore issues using the Snyk Web UI, or by creating a `.snyk` policy file along with your project when scanning. For more information, see [Iac Ignores using the .snyk policy file](snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file.md).

{% hint style="info" %}
Issues that are ignored by using the `.snyk` policy file can not be unignored in the Snyk web UI.
{% endhint %}

### Project tags

You can attach tags to the scanned projects using the `--project-tags` flag. The flag accepts a comma-separated list of tags, where each tag is a key-value pair. Keys and values are separated by an `=` sign. The `--project-tags` flag is valid only when used with `--report`.

The following example attaches the tags `department` and `team` to the scanned projects, with values `platform` and `persistence`, respectively.

```
> snyk iac test myproject --report \
    --project-tags=department=platform,team=persistence
```

### Project attributes

You can set attributes for the scanned projects using the `--project-business-criticality`, `--project-environment`, and `--project-lifecycle` tags. These flags are valid only when used with `--report`.

* `--project-business-criticality` accepts a comma-separated list of the following values: `critical`, `high`, `medium`, `low`.
* `--project-environment` accepts a comma-separated list of the following values: `frontend`, `backend`, `internal`, `external`, `mobile`, `saas`, `onprem`, `hosted`, `distributed`.
* `--project-lifecycle` accepts a comma-separated list of the following values: `production`, `development`, `sandbox`.

The following example sets the business criticality to `high`, the environment to the values `frontend` and `internal`, and the lifecycle to `development` for each scanned project.

```
> snyk iac test myproject --report \
    --project-business-criticality=high \
    --project-environment=frontend,internal \
    --project-lifecycle=development
```

### Target reference

You can set the target reference for the scanned projects using the `--target-reference` flag. This flag is valid only only when used with `--report`.

The following example sets the target reference for the scanned projects to the name of the current Git branch.

```
snyk iac test myproject --report \
    --target-reference="$(git branch --show-current)"
```

### **Notes**

Using [Custom rules](custom-rules/) and the Share Results functionality together is not currently supported.

Snyk does not share any file contents over the network, only the required metadata for the configuration issues that were just scanned.
