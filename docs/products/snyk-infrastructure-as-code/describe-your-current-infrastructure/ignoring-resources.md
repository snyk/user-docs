# Ignoring resources

The following information is from the [.driftignore documentation](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore).

**.driftignore** is a simple way to ignore resources, you put resources in a `.driftignore` file like a `.gitignore`.

If you need only to exclude a set of resources you should use .driftignore, if you need something more advanced, check filter rules.

Create the .driftignore file where you launch [snyk iac describe](../../../snyk-cli/commands/iac-describe.md) (usually the root of your IaC repo).

Each line must be of kind

* `resource_type.resource_id`, resource\_id could be a wildcard to exclude all resources of a given type.
* `resource_type.resource_id.path.to.field_name`, resource\_id can be wildcard to ignore a drift on given field for a given type, path could also contain wildcards.

The .driftignore file also supports negation of rules. This way you could ignore everything except certain types.

For example, if you want to ignore everything but IAM drifts:

```
*
!aws_iam_*
```

#### Examples[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#examples) <a href="#examples" id="examples"></a>

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

### Precedence over filter rules[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#precedence-over-filter-rules) <a href="#precedence-over-filter-rules" id="precedence-over-filter-rules"></a>

The above mechanism to ignore resources can be used in combination with filter rules. Bear in mind that if the same resource is included by a filter rule and excluded inside the .driftignore file, snyk iac describe will just ignore this resource.

### Automatically generate a driftignore file[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#automatically-generate-a-driftignore-file) <a href="#automatically-generate-a-driftignore-file" id="automatically-generate-a-driftignore-file"></a>

See [driftignore generator command](../../../snyk-cli/commands/iac-gen-driftignore.md).
