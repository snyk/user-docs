# Get started with Snyk IaC Describe on AWS

## **Step 1: Configure AWS authentication for your environment**

The [`snyk iac describe`](../../../../snyk-cli/commands/iac-describe.md) command requires authentication to your cloud provider in order to run properly. It requires only the lowest read-only access rights possible. You can use use the built-in AWS `ReadOnlyAccess` IAM policy for an IAM user as a the default to get started.

`snyk iac describe` can reuse standard authentication methods for AWS, such as the `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `AWS_REGION` [environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-list). When those are set, the Snyk CLI will automatically pick them up to authenticate on AWS.

Alternatively, you can configure the [AWS profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) in `~/.aws/credentials` and use the `AWS_PROFILE` environment variable.

## **Step 2: Use the `describe` command to report drift**

### **Unmanaged resources**

Snyk IaC can report drift of unmanaged resources, that is, resources present on your cloud provider but not on your Terraform state. You probably will want to import those resources into Terraform or delete them from your IaaS account.

### Select Terraform state files

To understand the drift that happens in your cloud environment, compare the state of your environment to one or multiple Terraform state file(s) (`.tfstate`).

The state file can be located locally or in an S3 bucket. Terraform Cloud is also available, but out of scope for this getting started document.

The `--from` option helps Snyk determine the path of the `.tfstate` file.

For a single local Terraform state file, use the command:

`$ snyk iac describe --from="tfstate://path/to/terraform.tfstate" --only-unmanaged`

To load all the Terraform states found in a given directory automatically, you can use glob patterns like this:

`$ snyk iac describe --from="tfstate://path/to/**/*.tfstate" --only-unmanaged`

For a single Terraform state stored on an S3 backend:

`$ snyk iac describe --from="tfstate+s3://my-bucket/path/to/state.tfstate" --only-unmanaged`

You can also aggregate multiple Terraform state files by listing them in the `--from` option. You can scan your local directory for different files, or use different paths from different sources. To choose two specific Terraform states, execute the following:

`$ snyk iac describe --from="tfstate://path/to/terraform_S3.tfstate,tfstate://path/to/terraform_VPC.tfstate" --only-unmanaged`

### Drift results and next steps

#### Create baseline

You ran `snyk iac describe` once and got a report of the current IaC coverage of your cloud infrastructure. When you are done fixing all the issues you identified and are left with known differences you are not planning to change, you can create a baseline so those differences will not be displayed in your next scan.

There are two options: ignore a specific resource or ignore multiple resources.

**Ignore multiple resources**

Use the output of the `describe` command and extract its results to [update the .snyk excluded policies](../../../../snyk-cli/commands/iac-update-exclude-policy.md):

`$ snyk iac describe --json --all | snyk iac update-exclude-policy`

**Ignore a specific resource**

To ignore a specific resource, you must exclude it manually by editing the `.snyk` file and adding the resource details to the `exclude` list. For more information, see [Ignore resources](ignore-resources-for-drift.md).

You are now ready to add `snyk iac describe` as a recurring cronjob to get alerts when a new resource is created outside of your IaC deployment.
