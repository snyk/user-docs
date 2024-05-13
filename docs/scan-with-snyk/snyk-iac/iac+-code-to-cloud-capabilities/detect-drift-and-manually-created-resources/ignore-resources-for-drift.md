# Ignore resources for drift

The `.snyk` policy file can be used to exclude resources from being considered IaC drift by `snyk iac describe`. See [the `.snyk` policy file doc](../../../../manage-risk/policies/the-.snyk-file.md) for general information.

If you need to exclude only a set of resources, use `.snyk`. If you have more complex requirements, consider using filter rules. For more information see [Filter rules](filter-rules.md).

Create the `.snyk` file in the directory where you launch the `snyk iac describe` command, typically the root of your IaC repo.

Each line must be structured as follows:

* `resource_type.resource_id`, where `resource_id` is a wildcard to exclude all resources of a given type
* `resource_type.resource_id.path.to.field_name`, where `resource_id` is a wildcard to ignore a drift on a given field for a given type and `path,` can also contain wildcards.

## Examples of IaC ignores

Ignore a single IAM user (_aws\_iam\_user_) named "_tfc-demo_".

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
exclude:
  iac-drift:
    - aws_iam_user.tfc-demo
```
{% endcode %}

Ignore all S3 buckets drifts.

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
exclude:
  iac-drift:
    - aws_s3_bucket.*
```
{% endcode %}

The `.snyk` policy file also supports the negation of rules. This allows you to ignore everything except certain types. In this example, only S3 buckets will **not** be ignored:

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
exclude:
  iac-drift:
    - '*'
    - '!aws_s3_bucket'
```
{% endcode %}

Ignore a specific IAM Policy Attachment (AWSServiceRoleForRDS) using its ARN (arn:aws:iam::aws:policy/aws-service-role/AmazonRDSServiceRolePolicy).

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
exclude:
  iac-drift:
    - aws_iam_policy_attachment.AWSServiceRoleForRDS-arn:aws:iam::aws:policy/aws-service-role/AmazonRDSServiceRolePolicy
```
{% endcode %}

Ignore the S3 bucket called my-bucket and so on, as shown.

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
exclude:
  iac-drift:
      # Will ignore S3 bucket called my-bucket
    - aws_s3_bucket.my-bucket
      # Will ignore every aws_instance resource
    - aws_instance.*
      # Will ignore environment for all lambda functions
    - aws_lambda_function.*.environment
      # Will ignore resources like aws_iam_role.AmazonSSMRoleForInstances and aws_iam_role.AWSServiceRoleForAmazonSSM
    - *role.*Amazon*
      # Will ignore lastModified for my-lambda-name lambda function
    - aws_lambda_function.my-lambda-name.last_modified
```

## Precedence over filter rules[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#precedence-over-filter-rules) <a href="#precedence-over-filter-rules" id="precedence-over-filter-rules"></a>

You can use the means to ignore resources explained on this page in combination with filter rules.

**Note:** If the same resource is included by a filter rule and excluded inside the `.snyk` file, `snyk iac describe` ignores this resource.

## Automatically generate drift exclusion rules[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#automatically-generate-a-driftignore-file) <a href="#automatically-generate-a-driftignore-file" id="automatically-generate-a-driftignore-file"></a>

For details, run`snyk iac update-exclude-policy --help.`

This command helps to generate a `.snyk` policy file, adding all the detected drifts to it in order to ignore them all.

For example, to ignore all the unmanaged resources at once, run the following command:

```
$ snyk iac describe --json | snyk iac update-exclude-policy
```
