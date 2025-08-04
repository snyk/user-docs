# Snyk IaC CLI test results (v. 1.939.0 and later)

{% hint style="info" %}
The instructions in this section apply to any file format supported by Snyk Infrastructure as Code, including Terraform, Kubernetes, CloudFormation, and ARM.
{% endhint %}

Snyk CLI analyzes your configuration file for issues and provides information and advice on how to resolve the issues discovered.

For example, when scanning a Terraform file,  run the following command:

```
snyk iac test aws_api_gateway_stage_logging.tf
```

The results from running this command follow:

```
Snyk Infrastructure as Code

✔ Test completed.

Issues

Low Severity Issues: 1

  [Low] API Gateway access logging disabled
  Info:    Amazon Api Gateway access logging is not enabled. Audit records may not be available during investigation
  Rule:    https://security.snyk.io/rules/cloud/SNYK-CC-TF-138
  Path:    resource > aws_api_gateway_stage[denied] > access_log_settings
  File:    aws_api_gateway_stage_logging.tf
  Resolve: Set `access_log_settings` attribute

-------------------------------------------------------

Test Summary

  Organization: demo-org

✔ Files without issues: 0
✗ Files with issues: 1
  Invalid files: 0
  Ignored issues: 0
  Total issues: 1 [ 0 critical, 0 high, 0 medium, 1 low ]
```

These results include a list of issues sorted by severity, where each issue reported includes the following details:

* **Heading** - the issue that was detected and the severity level of that issue.
* **Info** - a short description of the detected issue.
* **Rule** - a link to the rule documentation.
* **Path** - the property path within the configuration file where the issue was identified. See the example following this list.
* **File** - the file where the issue is located.
* **Resolve** - a short explanation of how to resolve the issue.

An example of the property path follows.

```
resource > aws_api_gateway_stage[denied] > access_log_settingsresource > aws_api_gateway_stage[denied] > access_log_settings
```

The following example represents the content of the `aws_api_gateway_stage` block, called "**denied**", which lacks the `access_log_settings` field:

{% code title="aws_api_gateway_stage_logging.tf" %}
```
resource "aws_api_gateway_stage" "denied" {
  xray_tracing_enabled = true
}
```
{% endcode %}
