# Getting started with Snyk Cloud: Google

Snyk Cloud scans the infrastructure configuration in your [Google Cloud](https://cloud.google.com/) projects and detects misconfigurations that can lead to vulnerabilities.

You can onboard a Google Cloud account to Snyk using the following methods:

* [Snyk Web UI](snyk-cloud-for-google-web-ui/)
* [Snyk API](snyk-cloud-for-google-api/)

## Prerequisites

To start using Snyk Cloud, you need the following:

* A Snyk Business or Enterprise [plan](https://snyk.io/plans/)
* A new Snyk Organization, with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator [role](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions)
* An Organization-level [service account](https://docs.snyk.io/features/user-and-group-management/structure-account-for-high-application-performance/service-accounts#set-up-a-service-account) with an Org Admin role
* Access to a [Google Cloud](https://cloud.google.com/) project and associated credentials with permissions to create a read-only Google service account
* Access to the [Terraform CLI](https://www.terraform.io/downloads), [configured with your Google credentials](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting\_started), to create the Google service account for Snyk
* **API only:** An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
* **API only (optional)**: [jq](https://stedolan.github.io/jq/), to unescape JSON containing the service account Terraform template
