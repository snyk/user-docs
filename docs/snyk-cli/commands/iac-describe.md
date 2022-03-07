# IAC Describe

## Usage

`snyk iac describe [<OPTIONS>]`

## Description

The `snyk iac describe` command detect, track and alert on infrastructure drift and unmanaged resources.

## Exit codes

Possible exit codes and their meaning:

**0**: success, no drift found\
**1**: drifts or unmanaged resources found\
**2**: failure\


## Configure the Snyk CLI

You can use environment variables and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI.](../../features/snyk-cli/configure-the-snyk-cli.md)

## Configure the terraform provider

You can use environment variables and also set variables to configure terraform provider used by `describe` command. See [Configure terraform providers](../../products/snyk-infrastructure-as-code/describe-your-current-infrastructure/configuring-providers/).

## Debug

Use the `-d` option to output the debug logs.

## Options

### `--from`

Currently, `snyk iac describe` supports reading IaC from a Terraform state.

Multiple states can be read by passing `--from` a comma separated list. You can also use glob patterns to match multiple state files at once.

Se [documentation](../../products/snyk-infrastructure-as-code/describe-your-current-infrastructure/iac-sources-usage.md) on ho to use the different iac sources.

Examples:

I want to read a local state and a state stored in an S3 bucket: `$ snyk iac describe --from="tfstate+s3://statebucket/terraform.tfstate,tfstate://terraform_toto.tfstate"`

You can also read all files under a given prefix for S3 `$ snyk iac describe --from=tfstate+s3://statebucket/states`

Or in a given directory `$ snyk iac describe --from=tfstate://directory/*`

#### Supported IaC sources

* Terraform state
* Local: `--from=tfstate://terraform.tfstate`
* S3: `--from=tfstate+s3://my-bucket/path/to/state.tfstate`
* GCS: `--from=tfstate+gs://my-bucket/path/to/state.tfstate`
* HTTPS: `--from=tfstate+https://my-url/state.tfstate`
* Terraform Cloud / Terraform Enterprise: `--from=tfstate+tfcloud://WORKSPACE_ID`
* Azure blob storage: `--from=tfstate+azurerm://container-name/path/to/state.tfstate`

You can use any unsupported backend by using `terraform` to pipe your state in a file and then use this file:

`$ terraform state pull > state.tfstate` `$ snyk iac describe --from=tfstate://state.tfstate`

### `--to`

`describe` commands supports multiple providers. By default, it will scan against AWS, but you can change this using `--to`.

#### Usage

`$ snyk iac describe --to=PROVIDER+TYPE`

Examples:

`$ snyk iac describe --to=aws+tf`

#### Supported Providers

* `github+tf`
* `aws+tf`
* `gcp+tf`
* `azure+tf`

### `--service`

The services(s) specified with this argument controls which resources are going to be included/ignored in drift detection.

When specifying multiple services, use a comma to separate them e.g. `snyk iac describe --service=aws_s3,aws_ec2`

This flag can't be used at the same time with a driftignore file, driftignore will be ignored.

Here are the available services: `aws_s31`, `aws_ec21`, `aws_lambda1`, `aws_rds1`, `aws_route531`, `aws_iam1` , `aws_vpc1`, `aws_api_gateway1`, `aws_apigatewayv21`, `aws_sqs1`, `aws_sns1`, `aws_ecr1`, `aws_cloudfront1`, `aws_kms1`, `aws_dynamodb1`, `azure_base1` , `azure_compute1`, `azure_storage1`, `azure_network1`, `azure_container1`, `azure_database1`, `azure_loadbalancer1`, `azure_private_dns1`, `google_cloud_platform1`, `google_cloud_storage1`, `google_compute_engine1`, `google_cloud_dns1` , `google_cloud_bigtable1`, `google_cloud_bigquery1`, `google_cloud_functions1`, `google_cloud_sql1`, `google_cloud_run`

### `--only-managed`

This flag will display a report that show only change for managed resources. It will filter out the unamanged resources.

### `--only-unmanaged`

This flag will display a report that show only unmanaged resources. It will filter out drift for managed resources.

### `--quiet`

This flag prevents stdout to be use for anything but the scan result. This can be useful to pipe the output into some other command.

### `--filter`

Filter rules allow you to build complex expression to include and exclude a set of resources in your workflow. Powered by expression language JMESPath you could build a complex include and exclude expression.

See [documentation for more details](../../products/snyk-infrastructure-as-code/describe-your-current-infrastructure/filtering-results.md).

### `--json`

Output report as json to stdout.

### `--json-output-file=`

Output report as json into a file.

### `--html`

`--html-output-file=`

Output report as html to stdout or into a file.

### `--headers`

Use a specific HTTP header(s) for the HTTP backend.

Example:

`$ GITLAB_TOKEN=<access_token> \ snyk iac describe \ --from tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state> \ --headers "Authorization=Bearer ${GITLAB_TOKEN}"`

### `--tfc-token`

Specify an API token to authenticate to the Terraform Cloud or Enterprise API.

### `--tfc-endpoint`

You can also read the current state for a given workspace from Terraform Enterprise by passing the tfc-endpoint value that's specific to your Org's Terraform Enterprise installation.

You can obtain your workspace id from the General Settings of the workspace.

Don't forget to provide your Terraform Enterprise API token.

Example:

`$ snyk iac describe --from tfstate+tfcloud://$WORKSPACE_ID --tfc-token $TFC_TOKEN --tfc-endpoint 'https://tfe.example.com/api/v2'`

### `--tf-provider-version`

You can specify a terraform provider version to use. If none, defaults version will be used like below:

* aws@3.19.0
* github@4.4.0

#### Usage

I use terraform provider 3.43.0 so I can use this provider to avoid scan errors. `$ DCTL_TF_PROVIDER_VERSION=3.43.0 snyk iac describe`

Same parameter is used for every cloud provider. `$ DCTL_TF_PROVIDER_VERSION=4.10.1 snyk iac describe --to github+tf`

### `--strict`

When running this command against an AWS account, you may experience unnecessary noises with resources that don't belong to you. It can be the case if you have an organization account in which you will by default have a service-linked role associated to your account (e.g. AWSServiceRoleForOrganizations). For now, the command ignores those service-linked resources by default.

If you still want to include those resources in the report anyway, you can enable the strict mode.

For now, resources include:

* Service-linked AWS IAM roles, including their policies and policy attachments

#### Usage

`$ snyk iac describe --strict`

### `--deep`

#### Warning

This flag is **EXPERIMENTAL**. Enabling deep mode while using a Terraform state as IaC source can lead to unexpected behaviors: false positive drifts, undetected drifts.

Deep mode enables resources details retrieval.

* In **deep** mode we compare resources details to expected ones (like a terraform plan).
* In **non-deep** mode (the default one) we only enumerate resources and display which ones are out of IaC scope.

Since it overlaps the new `terraform plan` behavior (as of Terraform 0.15 it shows diffs between your state and the remote) we moved the original behavior under the `--deep` **experimental** flag.

#### Usage

`$ snyk iac describe --deep`

### `--driftignore`

The default name for a driftignore file is `.driftignore`. If for some reason you want to use a custom filename, you can do so using the `--driftignore` flag. This is especially useful when you have multiple driftignore files, where each of them represents a particular use case.

NOTE: You can use only one driftignore file at once.

See [ignoring ressources](../../products/snyk-infrastructure-as-code/describe-your-current-infrastructure/ignoring-resources.md).

#### Usage

Apply ignore directives from the /path/to/driftignore file

`$ snyk iac describe --driftignore /path/to/driftignore`

### `--tf-lockfile`

By default, the command tries to read a Terraform lock file (.terraform.lock.hcl) in the current directory, so it can automatically detect which provider to use, according to the --to flag. You can specify a custom path for that file using the --tf-lockfile flag. If parsing the lockfile fails for some reason, errors will be logged and scan will continue.

### `--org=<ORG_ID>`

Specify the `<ORG_ID>` to run Snyk commands tied to a specific organization. The `<ORG_ID>` influences some features availability and private test limits.

If you have multiple organizations, you can set a default from the CLI using:

`$ snyk config set org=<ORG_ID>`

Set a default to ensure all newly tested projects are tested under your default organization. If you need to override the default, use the `--org=<ORG_ID>` option.

Default: `<ORG_ID>` that is the current preferred organization in your [Account settings](https://app.snyk.io/account).

Example: `$ snyk test --org=my-team`

For more information see the article [How to select the organization to use in the CLI](https://support.snyk.io/hc/en-us/articles/360000920738-How-to-select-the-organization-to-use-in-the-CLI).

#### Note

When using both --tf-lockfile and --tf-provider-version flags together, --tf-provider-version will simply take precedence overall.

#### Example

`$ snyk iac describe --to aws+tf --tf-lockfile path/to/.terraform.lock.hcl`

### `--config-dir`

You can change the directory path used for configuration. By default, it is the `$HOME` directory.

This can be useful, for example, if you want to invoke this command in an AWS Lambda function where you can only use the `/tmp` folder.

#### Usage

`$ snyk iac describe --config-dir path_to_config_dir` `$ DCTL_CONFIG_DIR=path_to_config_dir snyk iac describe`

## Examples for the `iac describe` command

### With a local state

```
$ snyk iac describe
```

### Same as

```
$ snyk iac describe --from=tfstate://terraform.tfstate
```

### To specify AWS credentials

```
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=XXX snyk iac describe
```

### or using a named profile

```
$ AWS_PROFILE=profile_name snyk iac describe
```

### With state stored on a s3 backend

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
