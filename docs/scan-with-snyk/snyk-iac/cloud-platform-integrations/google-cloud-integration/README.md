# Google Cloud integration

Snyk integrations with your [Google Cloud](https://cloud.google.com/) Projects to find issues in your cloud configurations and generate cloud context to help you prioritize your vulnerabilities.

You can onboard a Google Cloud account to Snyk using the following methods:

* [Snyk Web UI](google-cloud-integration-web-ui/)
* [Snyk API](google-cloud-integration-api/)

## Prerequisites for Google Cloud integration

To add a Google Cloud integration, you need the following:

* A Snyk Enterprise [plan](https://snyk.io/plans/)
* A new Snyk Organization, with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator role
* Access to a [Google Cloud](https://cloud.google.com/) project and associated credentials with permissions to create a read-only Google service account
* Access to the [Terraform CLI](https://www.terraform.io/downloads), [configured with your Google credentials](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting_started), to create the Google service account for Snyk
* API only: An Organization-level [service account](../../../../implementation-and-setup/enterprise-setup/service-accounts/) with an Org Admin role to use the Snyk API
* API only: An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
* API only (optional): [jq](https://stedolan.github.io/jq/), to unescape JSON containing the service account Terraform template
