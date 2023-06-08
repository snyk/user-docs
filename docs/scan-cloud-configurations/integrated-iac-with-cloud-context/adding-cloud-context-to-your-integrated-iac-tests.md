# Adding cloud context to your Integrated IaC tests

{% hint style="info" %}
The cloud context feature is available for [Integrated IaC](broken-reference) only, and supports AWS.
{% endhint %}

## What is cloud context?

Snyk IaC's cloud context feature uses information from deployed cloud infrastructure, either via [Integrated IaC with cloud context ](./)or through local enumeration, to suppress certain issues from your IaC tests.

For example, suppose your Terraform configuration declares an Amazon Web Services (AWS) S3 bucket that does not have a public access block, but you have an account-level public access block. Snyk applies the cloud context from your AWS account to suppress false positive issues stating your bucket is not secured by a public access block.

Example results without cloud context:

```
Test Summary

  Organization: demo-production
  Project name: terraform

✔ Files without issues: 0
✗ Files with issues: 1
  Ignored issues: 0
  Total issues: 15 [ 0 critical, 7 high, 3 medium, 5 low ]
```

Example results with cloud context:

```
Test Summary

  Organization: demo-production
  Project name: terraform

✔ Files without issues: 0
✗ Files with issues: 1
  Ignored issues: 0
  Cloud context - suppressed issues: 5
  Total issues: 10 [ 0 critical, 2 high, 3 medium, 5 low ]
```

The output summary lists the number of suppressed issues, for example, `Cloud context - suppressed issues: 5`. These suppressed issues are not included in the total issue count, for example, `Total issues: 10 [ 0 critical, 2 high, 3 medium, 5 low ]`.

Currently, Terraform for Amazon Web Services (AWS) is supported.

There are two ways that Snyk IaC can apply cloud context and suppress issues in your IaC test results:

* [Bringing context from Snyk](adding-cloud-context-to-your-integrated-iac-tests.md#bringing-context-from-a-snyk-cloud-scan)
* [Bringing context from local enumeration](adding-cloud-context-to-your-integrated-iac-tests.md#bringing-context-from-a-live-scan)

## Bringing context from Snyk <a href="#bringing-context-from-a-snyk-cloud-scan" id="bringing-context-from-a-snyk-cloud-scan"></a>

If you have a [Snyk cloud environment](key-concepts.md#environments), you can leverage what Snyk already knows about your cloud provider account to apply cloud context and reduce false positives in your IaC tests.

Use the `--snyk-cloud-environment=<ENVIRONMENT_ID>` option with [`snyk iac test`](../../snyk-cli/commands/iac-test.md) to tell Snyk which cloud environment to use as context for your IaC test.

For example, the following command tests the IaC in the present working directory and applies cloud context from the results of the latest scan for the Snyk cloud environment `93786877-c9f8-0000-1234-abcd1234efgh`:

```
snyk iac test --snyk-cloud-environment=93786877-c9f8-0000-1234-abcd1234efgh
```

To find your environment ID, see [Find an environment ID](snyk-environments/find-an-environment-id.md).

See the [Integrated IaC with cloud context](./) documentation for information about creating a Snyk cloud environment.

## Bringing context from local enumeration <a href="#bringing-context-from-a-live-scan" id="bringing-context-from-a-live-scan"></a>

You can use the `--cloud-context=<PROVIDER>` option with the [`snyk iac test`](../../snyk-cli/commands/iac-test.md) command to apply cloud context to an IaC test without referencing a Snyk cloud environment. Snyk does this by authenticating with your cloud provider locally and scanning the associated cloud provider account.

The process for authenticating with a cloud provider account is the same as for the [`snyk iac describe`](../../snyk-cli/commands/iac-describe.md) command. See **Step 1** of [Get started with Snyk IaC Describe](../snyk-infrastructure-as-code/detect-drift-and-manually-created-resources/get-started-with-snyk-iac-describe-on-aws.md) on AWS for details.

For example, the following command tests the IaC in the present working directory and executes a scan of the authenticated AWS account, then applies the cloud context to suppress issues:

```
snyk iac test --cloud-context=aws
```
