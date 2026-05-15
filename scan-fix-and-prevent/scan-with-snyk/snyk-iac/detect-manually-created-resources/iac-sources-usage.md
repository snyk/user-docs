# IAC sources usage

## **Supported IaC sources**

At this time, the `snyk iac describe` command supports reading Terraform states as follows:

* Local: `--from="tfstate://terraform.tfstate"`
* S3: `--from="tfstate+s3://my-bucket/path/to/state.tfstate"`
* GCS: `--from="tfstate+gs://my-bucket/path/to/state.tfstate"`
* HTTPS: `--from="tfstate+https://my-url/state.tfstate"`
* Terraform Cloud / Terraform Enterprise: `--from="tfstate+tfcloud://WORKSPACE_ID"`
* Azure blob storage: `--from="tfstate+azurerm://container-name/path/to/state.tfstate"`

You can use any unsupported backend by using `terraform` to pipe your state in a file and then use the file with `snyk iac describe`:

```
$ terraform state pull > state.tfstate
$ snyk iac describe --from="tfstate://state.tfstate"
```

## **S3 read-only access IAM policy**

The `snyk iac describe` command needs read-only access. The following policy ensures minimal access to your state file.

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

## **HTTP + GitLab**

The HTTP backend supports the GitLab-managed Terraform state using the GitLab API.

You need a GitLab repository that contains a Terraform state and an access token with the `read_api` scope.

Use the following command:

```
$ GITLAB_TOKEN=<access_token> \
snyk iac describe \
--from="tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state>" \
--headers "Authorization=Bearer ${GITLAB_TOKEN}"
```

For more information about the GitLab-managed Terraform State, see [GitLab-managed Terraform state](https://docs.gitlab.com/ee/user/infrastructure/terraform\_state.html) on the GitLab documentation website.

## **Azure Blob Storage**

To access state from Azure Blob Storage, define the following environment variables:

```
$ export AZURE_STORAGE_ACCOUNT=...
$ export AZURE_STORAGE_KEY=...
$ snyk iac describe --from="tfstate+azurerm://my-container/terraform.tfstate"
```

You can find these values in your Azure console, as shown in the following screenshot:

<figure><img src="https://docs.driftctl.com/assets/images/azure_storage_account_keys-ccb38d8792616d4376050fc6b715a6ef.png" alt="Azure storage environment variables"><figcaption><p>Azure storage environment variables</p></figcaption></figure>
