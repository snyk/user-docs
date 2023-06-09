# IaC describe command examples

For a full list of `snyk iac describe` options, see [`snyk iac describe`](../../../snyk-cli/commands/iac-describe.md) command help or display the help by running:

```
snyk iac describe --help
```

## Use `--from` to specify state files

Scan for AWS drift and aggregate both a local Terraform state and one stored in an S3 bucket:

```
snyk iac describe --all --from="tfstate+s3://statebucket/terraform.tfstate,tfstate://other_terraform.tfstate"
```

Scan for AWS drift and aggregate all Terraform states under a given prefix for S3:

```
snyk iac describe --all --from="tfstate+s3://statebucket/states"
```

Read and aggregate all Terraform states in a given directory:

```
snyk iac describe --all --from="tfstate://directory/*.tfstate"
```

Use any unsupported backend by using `terraform` to pipe your state into a file and then use the file:

```
terraform state pull > state.tfstate

snyk iac describe --all --from="tfstate://state.tfstate"
```

## Use `--to` to specify the cloud provider to scan

Explicitly scan AWS in a Terraform context:

```
snyk iac describe --to="aws+tf"
```

## Use `--tf-provider-version` to specify the Terraform provider version

Specify terraform provider 3.43.0 to use this provider to avoid scan errors:

```
snyk iac describe --only-unmanaged --tf-provider-version=3.43.0
```

Use the same parameter for every cloud provider:

```
snyk iac describe --all --to="github+tf" --tf-provider-version=4.10.1
```

## Use `--tf-lockfile` to specify the Terraform lock file

Specify a custom path for the Terraform lock file (`.terraform.lock.hcl`):

```
snyk iac describe --all --to="aws+tf" --tf-lockfile="/path/to/.terraform.lock.hcl"
```

## Use `--fetch-tfstate-headers` to specify HTTP headers when fetching Terraform state

Specify HTTPS authentication to use a Terraform state stored on GitLab:

```
GITLAB_TOKEN=<access_token> \
  snyk iac describe --all \
  --from="tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state>" \
 --fetch-tfstate-headers='Authorization="Bearer ${GITLAB_TOKEN}"'
```

## Use `--tfc-endpoint` to read state from a Terraform Enterprise workspace

You can obtain your workspace ID from the **General Settings** of the Terraform Enterprise workspace.

Remember to provide your Terraform Enterprise API token.

Example:

```
snyk iac describe --all --from="tfstate+tfcloud://$WORKSPACE_ID" --tfc-token="$TFC_TOKEN" --tfc-endpoint="https://tfe.example.com/api/v2"
```

## Use `--service` to specify multiple services to inspect

Include AWS S3 and AWS EC2 resources in the report:

```
snyk iac describe --all --service="aws_s3,aws_ec2"
```

## Use `--strict` to include service-linked resources in the report

**Note:** When using strict mode with an AWS account, you may experience unnecessary noise from resources that do not belong to you.

This can happen if you have an Organization account in which you, by default, have a service-linked role associated with your account; for example, `AWSServiceRoleForOrganizations`.

Example to enable strict mode:

```
snyk iac describe --all --strict
```

## Use `--json` to output the report as JSON

Save the report to a JSON file through redirection:

```
snyk iac describe --all --json > report.json
```
