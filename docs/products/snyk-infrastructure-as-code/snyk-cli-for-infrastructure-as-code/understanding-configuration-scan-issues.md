# Understanding the CLI Output

Snyk analyzes your provided configuration file for issues and provides advice on how to resolve the issue directly from the CLI.

For example - scanning a Terraform file:

```
snyk iac test aws_api_gateway_stage_logging.tf
```

could give an output as follows

![](../../../.gitbook/assets/screenshot-2021-09-28-at-19.58.22.png)

This example is of output from a Terraform file, but this guide applies to any file format including Kubernetes or CloudFormation.

## List of vulnerabilities—sorted by severity, where each is detailed as follows:

**A clear heading line** - specifying the issue that has been detected, the severity of that issue and the Snyk Policy Id for that particular issue.

**Location** - the property path within the configuration file at which the issue has been identified. See the example below for more details.

## **As an example:**

![](../../../.gitbook/assets/screenshot-2021-09-28-at-20.00.36.png)

The path of this issue is specified as:

```
resource > aws_api_gateway_stage[denied] > access_log_settings
```

In the following code you can see that line 1 represents the contents of the `aws_api_gateway_stage` block named "denied" which is missing the `access_log_settings` field.

{% code title="aws_api_gateway_stage_logging.tf" %}
```
resource "aws_api_gateway_stage" "denied" {
  xray_tracing_enabled = true
}
```
{% endcode %}
