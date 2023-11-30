# Prerequisites for Bitbucket Pipelines integration

The following are prerequisites for Bitbucket Pipelines integration:

* For your Bitbucket Pipelines, ensure you have build minutes in your account, which are necessary to enable ongoing CI/CD workflows.
* Create a Snyk account and retrieve the Snyk API token from your **Account settings**.
* Create a Repository variable from Bitbucket for your Snyk API token. Call the variable SNYK\_TOKEN.
* If you supply a custom image, ensure it meets the [requirements](user-defined-custom-images.md#requirements) for custom images
