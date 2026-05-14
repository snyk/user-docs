# Prerequisites for Bitbucket Pipelines integration

The following are prerequisites for Bitbucket Pipelines integration:

* For your Bitbucket Pipelines, ensure you have build minutes in your account, which are necessary to enable ongoing CI/CD workflows.
* Create a Snyk account and retrieve the Snyk API token or Snyk PAT from your **Account settings**.
* Create a Repository variable from Bitbucket for your token. Call the variable SNYK\_TOKEN.
* If you supply a custom image, ensure it meets the [requirements](../user-defined-custom-images-for-cli.md#requirements-for-user-defined-custom-images-for-cli) for custom images
