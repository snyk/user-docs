# IAC Sources usage

## **Supported IaC sources**[**​**](https://docs.driftctl.com/0.22.0/usage/cmd/scan-usage#supported-iac-sources)

* Terraform state
* Local: `--from=tfstate://terraform.tfstate`
* S3: `--from=tfstate+s3://my-bucket/path/to/state.tfstate`
* GCS: `--from=tfstate+gs://my-bucket/path/to/state.tfstate`
* HTTPS: `--from=tfstate+https://my-url/state.tfstate`
* Terraform Cloud / Terraform Enterprise: `--from=tfstate+tfcloud://WORKSPACE_ID`
* Azure blob storage: `--from=tfstate+azurerm://container-name/path/to/state.tfstate`

You can use any unsupported backend by using `terraform` to pipe your state in a file and then use this file with driftctl:

```
$ terraform state pull > state.tfstate
$ snyk iac describe --from=tfstate://state.tfstate
```

## **S3**[**​**](https://docs.driftctl.com/0.22.0/usage/cmd/scan-usage#s3)

The `snyk iac describe` command needs read only access, so you can use the following policy to ensure minimal access to your state file.

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::mybucket"
    },
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::mybucket/path/to/my/key"
    }
  ]
}
```

## **HTTP + GitLab**[**​**](https://docs.driftctl.com/0.22.0/usage/cmd/scan-usage#http--gitlab)

The HTTP backend supports the GitLab-managed Terraform state using their API.

Ensure you have a GitLab repository that contains a Terraform state and an access token with the `read_api` scope.

Use the following command:

```
$ GITLAB_TOKEN=<access_token> \
snyk iac describe \
--from=tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state> \
--headers "Authorization=Bearer ${GITLAB_TOKEN}"
```

You can find more information about the GitLab-managed Terraform state on the [GitLab docs site](https://docs.gitlab.com/ee/user/infrastructure/terraform\_state.html).

## **Azure blob storage**[**​**](https://docs.driftctl.com/0.22.0/usage/cmd/scan-usage#azure-blob-storage)

To access state from azure blob storage,  define the following environment variables:

```
$ export AZURE_STORAGE_ACCOUNT=...
$ export AZURE_STORAGE_KEY=...
$ snyk iac describe --from tfstate+azurerm://my-container/terraform.tfstate
```

You can find theses values in your Azure console as shown in the following screenshot:

![Azure account access keys](https://docs.driftctl.com/assets/images/azure\_storage\_account\_keys-ccb38d8792616d4376050fc6b715a6ef.png)
