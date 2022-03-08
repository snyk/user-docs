# IAC describe

## Usage

`snyk iac describe [<OPTIONS>]`

## Description

The `snyk iac describe` command detects, tracks, and alerts on infrastructure drift and unmanaged resources.

## Exit codes

Possible exit codes and their meaning:

**0**: success, no drift found\
**1**: drifts or unmanaged resources found\
**2**: failure\


## Configure the Snyk CLI

You can use environment variables and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI.](../../features/snyk-cli/configure-the-snyk-cli.md)

## Configure the terraform provider

You can use environment variables and also set variables to configure the terraform provider used by the `describe` command. For more information see [Configuring providers](https://docs.snyk.io/products/snyk-infrastructure-as-code/describe-your-current-infrastructure/configuring-providers).

## Debug

Use the `-d` option to output the debug logs.

## Options

### `--from=<STATE>`\[,\<STATE]...>

Specify multiple states to be read.

At this time the`snyk iac describe` command supports reading IaC from a Terraform state.

To read multiple states, pass `--from=` followed by a comma-separated list. You can also use glob patterns to match multiple state files at once.

For more information including **a list of supported IaC sources** and how to use them see [IAC Sources usage](https://docs.snyk.io/products/snyk-infrastructure-as-code/describe-your-current-infrastructure/iac-sources-usage).

Examples:

Read a local state and a state stored in an S3 bucket:

```
$ snyk iac describe --from="tfstate+s3://statebucket/terraform.tfstate,tfstate://terraform_toto.tfstate"
```

Read all files under a given prefix for S3:

&#x20;`$ snyk iac describe --from=tfstate+s3://statebucket/states`

Read all files in a given directory:

&#x20;`$ snyk iac describe --from=tfstate://directory/*`

Use any unsupported backend by using `terraform` to pipe your state in a file and then use this file:

`$ terraform state pull > state.tfstate` `$ snyk iac describe --from=tfstate://state.tfstate`

### `--to=<PROVIDER+TYPE>`

Specify the provider to scan.

The `iac describe` command supports multiple providers. By default the describe command scans against AWS, but you can change this using  the --to option.

Usage:

`$ snyk iac describe --to=PROVIDER+TYPE`

Example:

`$ snyk iac describe --to=aws+tf`

Supported providers

* `github+tf`
* `aws+tf`
* `gcp+tf`
* `azure+tf`

### `--service=<SERVICE>`\[,\<SERVICE]...>

Specify the services that control which resources are included, ignored, or both in drift detection.

Specify multiple services as a comma-separated list, for example:

&#x20;`snyk iac describe --service=aws_s3,aws_ec2`

This option cannot be used with a `.driftignore` file, because the `.driftignore` file will be ignored.

The supported services are: `aws_s31`, `aws_ec21`, `aws_lambda1`, `aws_rds1`, `aws_route531`, `aws_iam1` , `aws_vpc1`, `aws_api_gateway1`, `aws_apigatewayv21`, `aws_sqs1`, `aws_sns1`, `aws_ecr1`, `aws_cloudfront1`, `aws_kms1`, `aws_dynamodb1`, `azure_base1` , `azure_compute1`, `azure_storage1`, `azure_network1`, `azure_container1`, `azure_database1`, `azure_loadbalancer1`, `azure_private_dns1`, `google_cloud_platform1`, `google_cloud_storage1`, `google_compute_engine1`, `google_cloud_dns1` , `google_cloud_bigtable1`, `google_cloud_bigquery1`, `google_cloud_functions1`, `google_cloud_sql1`, `google_cloud_run`

### `--only-managed`

Display a report that shows change only for managed resources; filter out drift for unamanged resources.

### `--only-unmanaged`

Display a report that shows change only for unmanaged resources; filter out drift for managed resources.

### `--quiet`

Prevent stdout from being used for anything but the scan result. This can be useful to pipe the output into some other command.

### `--filter`

Use filter rules.

Filter rules allow you to build complex expression to include and exclude a set of resources in your workflow. You can build a complex include and exclude expression powered by the expression language JMESPath.

For more information see [Filtering results](https://docs.snyk.io/products/snyk-infrastructure-as-code/describe-your-current-infrastructure/filtering-results).

### `--json`

Output the report as JSON to stdout.

### `--json-output-file=<OUTPUT_FILE_PATH>`

Save test output in JSON format directly to the specified file, regardless of whether or not you use the `--json` option.

This is especially useful if you want to display the human-readable test output using stdout and at the same time save the JSON format output to a file.

### `--html`

`--html-output-file=<OUTPUT_FILE_PATH>`

Output the report as html to stdout or into a file.

### `--headers`

Use a specific HTTP header or headers for the HTTP backend.

Example:

```
$ GITLAB_TOKEN=<access_token> \ snyk iac describe \ --from tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state> \ --headers "Authorization=Bearer ${GITLAB_TOKEN}"
```

### `--tfc-token`

Specify an API token to authenticate to the Terraform Cloud or Enterprise API.

### `--tfc-endpoint`

Read the current state for a given workspace from Terraform Enterprise by passing the tfc-endpoint value that is specific to your Org's Terraform Enterprise installation.

You can obtain your workspace id from the **General Settings** of the workspace.

Remember to provide your Terraform Enterprise API token.

Example:

```
$ snyk iac describe --from tfstate+tfcloud://$WORKSPACE_ID --tfc-token $TFC_TOKEN --tfc-endpoint 'https://tfe.example.com/api/v2'
```

### `--tf-provider-version`

Specify a terraform provider version to use. If none is specified, default versions are used as follows:

* aws@3.19.0
* github@4.4.0

Usage:

Specify terraform provider 3.43.0 to use this provider to avoid scan errors:

`$ DCTL_TF_PROVIDER_VERSION=3.43.0 snyk iac describe`

Use the same parameter for every cloud provider:

`$ DCTL_TF_PROVIDER_VERSION=4.10.1 snyk iac describe --to github+tf`

### `--strict`

Enable strict mode.

When running this command against an AWS account, you may experience unnecessary noise from resources that do not belong to you. This can happen if you have an organization account in which you by default have a service-linked role associated to your the account, for example, **AWSServiceRoleForOrganizations**. The `iac describe` command ignores those service-linked resources by default. To include those resources in the report you can enable **strict mode**.

Resources include service-linked AWS IAM roles, including their policies and policy attachments

Usage:

`$ snyk iac describe --strict`

### `--deep`

Enable deep mode.

**WARNING:**

This option is **EXPERIMENTAL**. Enabling deep mode while using a Terraform state as the IaC source can lead to unexpected behaviors: false positive drifts and undetected drifts.

Deep mode enables retrieval of details for resources.

* In **deep** mode snyk compares details for resources to expected details, for example, a terraform plan.
* In **non-deep** mode (the default) snyk enumerates only resources and displays which ones are out of IaC scope.

Since the original behavior overlaps the new `terraform plan` behavior (as of Terraform 0.15 it shows diffs between your state and the remote) Snyk moved the original behavior under the `--deep` **experimental** flag.

Usage:

`$ snyk iac describe --deep`

### `--driftignore`

Specify a custom filename for a driftignore file.

The default name for a driftignore file is `.driftignore`. If for some reason you want to use a custom filename, you can do so using the `--driftignore` flag. This is especially useful when you have multiple driftignore files, where each of them represents a particular use case.

Note: You can use only one driftignore file at once.

See [Ignoring resources](https://docs.snyk.io/products/snyk-infrastructure-as-code/describe-your-current-infrastructure/ignoring-resources).

Example:

Apply ignore directives from the /path/to/driftignore file:

`$ snyk iac describe --driftignore /path/to/driftignore`

### `--tf-lockfile`

By default, the command tries to read a Terraform lock file (.terraform.lock.hcl) in the current directory, so it can automatically detect which provider to use, according to the --to flag. You can specify a custom path for that file using the --tf-lockfile flag. If parsing the lockfile fails for some reason, errors will be logged and scan will continue.

**Note**: When using both the --tf-lockfile and --tf-provider-version opton s together, --tf-provider-version takes precedence overall.

Example:

`$ snyk iac describe --to aws+tf --tf-lockfile path/to/.terraform.lock.hcl`

### `--org=<ORG_ID>`

Specify the `<ORG_ID>` to run Snyk commands tied to a specific organization. The `<ORG_ID>` influences some features availability and private test limits.

If you have multiple organizations, you can set a default from the CLI using:

`$ snyk config set org=<ORG_ID>`

Set a default to ensure all newly tested projects are tested under your default organization. If you need to override the default, use the `--org=<ORG_ID>` option.

Default: `<ORG_ID>` that is the current preferred organization in your [Account settings](https://app.snyk.io/account).

For more information see the article [How to select the organization to use in the CLI](https://support.snyk.io/hc/en-us/articles/360000920738-How-to-select-the-organization-to-use-in-the-CLI).

### `--config-dir`

You can change the directory path used for configuration. By default, it is the `$HOME` directory.

This can be useful, for example, if you want to invoke this command in an AWS Lambda function where you can only use the `/tmp` folder.

Usage:

```
$ snyk iac describe --config-dir path_to_config_dir $ DCTL_CONFIG_DIR=path_to_config_dir snyk iac describe
```

## Examples for the `iac describe` command

### With a local state

```
$ snyk iac describe
```

### Same as

```
$ snyk iac describe --from=tfstate://terraform.tfstate
```

### Specify AWS credentials

```
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=XXX snyk iac describe
```

### Use a named profile

```
$ AWS_PROFILE=profile_name snyk iac describe
```

### With state stored on an s3 backend

```
$ snyk iac describe --from=tfstate+s3://my-bucket/path/to/state.tfstate
```

### With multiple states

```
$ snyk iac describe --from=tfstate://terraform_S3.tfstate --from tfstate://terraform_VPC.tfstate
```

### Using glob pattern

```
$ snyk iac describe --from=tfstate://path/to/**/*.tfstate
$ snyk iac describe --from=tfstate+s3://path/to/**/*.tfstate
```

For more information see [Snyk CLI for Infrastructure as Code](../../products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/) .
