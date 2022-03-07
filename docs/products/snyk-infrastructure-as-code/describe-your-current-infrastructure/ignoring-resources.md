# Ignoring resources

Using a `.driftignore` file is a simple way to ignore resources. Put resources in a `.driftignore` file like a `.gitignore`.

To exclude only a set of resources, use `.driftignore`. If you need something more complex, consider using filter rules. See [Filtering results](filtering-results.md) for more information.

Create the `.driftignore` file in the directory where you launch [`snyk iac describe`](../../../snyk-cli/commands/iac-describe.md) (usually the root of your IaC repo).

Each line must be one of the following types:

* `resource_type.resource_id`, where resource\_id is a wildcard to exclude all resources of a given type
* `resource_type.resource_id.path.to.field_name`, where resource\_id is a wildcard to ignore a drift on given field for a given type, and `path` could also contain wildcards

The `.driftignore` file also supports negation of rules, allowing you to ignore everything except certain types.

For example, if you want to ignore everything but IAM drifts, use the following code:

```
*
!aws_iam_*
```

## Examples of .driftignore[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#examples) <a href="#examples" id="examples"></a>

```
# Will ignore S3 bucket called my-bucket
aws_s3_bucket.my-bucket
# Will ignore every aws_instance resource
aws_instance.*
# Will ignore environment for all lambda functions
aws_lambda_function.*.environment
# Will ignore resources like aws_iam_role.AmazonSSMRoleForInstances and aws_iam_role.AWSServiceRoleForAmazonSSM
*role.*Amazon*
# Will ignore lastModified for my-lambda-name lambda function
aws_lambda_function.my-lambda-name.last_modified
```

## Precedence over filter rules[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#precedence-over-filter-rules) <a href="#precedence-over-filter-rules" id="precedence-over-filter-rules"></a>

The means to ignore resources described on this page can be used in combination with filter rules. Note that if the same resource is included by a filter rule and excluded inside the `.driftignore` file, `snyk iac describe` ignores the resource.

## Automatically generate a .driftignore file[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#automatically-generate-a-driftignore-file) <a href="#automatically-generate-a-driftignore-file" id="automatically-generate-a-driftignore-file"></a>

See the [`IAC Gen-driftignore`](../../../snyk-cli/commands/iac-gen-driftignore.md) help.
