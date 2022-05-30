# Snyk IaC CLI test results (v. 1.939.0 and later)

**Note:** The instructions in this section apply to any supported file format of Snyk Infrastructure as Code, including Terraform, Kubernetes, CloudFormation, and ARM.

Snyk CLI analyzes your configuration file for issues, and provides information and advices on how to resolve the discovered issues.

For example, when scanning a Terraform file, we enter the following command:

```
snyk iac test aws_api_gateway_stage_logging.tf
```

Once the command run, the following results are received:

```
Snyk Infrastructure as Code

✔ Test completed.

Issues

Low Severity Issues: 1

  [Low] API Gateway access logging disabled
  Info:    Amazon Api Gateway access logging is not enabled. Audit records may not be available during investigation
  Rule:    https://snyk.io/security-rules/SNYK-CC-TF-138
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

The results include a list of issues sorted by severity, where each issue consists of the following details:

* **Heading** - the issue that was detected, and the severity level of that issue.
* **Info** - a short description of the detected issue.
* **Rule** - a link to the rule documentation.&#x20;
* **Path** - the property path within the configuration file, where the issue was identified. See the example below for more details.

For example**:**\
****The path of the issue is specified as:

```
resource > aws_api_gateway_stage[denied] > access_log_settingsresource > aws_api_gateway_stage[denied] > access_log_settings
```

In the following code, line 1 represents the content of the `aws_api_gateway_stage` block, called "**denied**", which is missing the `access_log_settings` field:

{% code title="aws_api_gateway_stage_logging.tf" %}
```
resource "aws_api_gateway_stage" "denied" {
  xray_tracing_enabled = true
}
```
{% endcode %}

* **File** - the file where the issue is located.
* **Resolve** - a short explanation on how to resolve the issue.
