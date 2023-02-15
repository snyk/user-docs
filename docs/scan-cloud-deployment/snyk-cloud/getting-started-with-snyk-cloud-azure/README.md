# Getting started with Snyk Cloud: Azure

{% hint style="info" %}
Azure support is currently in Closed Beta. If you are interested in using Snyk Cloud to scan your Azure subscription, please contact your Snyk account team. See [Snyk feature release process](../../../snyk-processes/snyk-feature-release-process.md) for more details.
{% endhint %}

Snyk Cloud scans the infrastructure configuration in your [Microsoft Azure](https://azure.microsoft.com/en-us/) subscription and detects misconfigurations that can lead to vulnerabilities.

You can onboard an Azure subscription to Snyk using the following methods:

* [Snyk Web UI](snyk-cloud-for-azure-web-ui/)
* [Snyk API](snyk-cloud-for-azure-api/)

## Prerequisites

To start using Snyk Cloud, you need the following:

* A Snyk Business or Enterprise [plan](https://snyk.io/plans/)
* A new Snyk Organization, with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator [role](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions)
* An Organization-level [service account](https://docs.snyk.io/features/user-and-group-management/structure-account-for-high-application-performance/service-accounts#set-up-a-service-account) with an Org Admin role
* Access to a [Microsoft Azure](https://azure.microsoft.com/en-us/) subscription and associated credentials with permissions to create the following resources:
  * [An Active Directory (AD) application registration](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration)
  * [A federated identity credential](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation)
    * If using Terraform to create this [resource](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application\_federated\_identity\_credential#api-permissions), your user must have either the Application Administrator or Global Administrator directory role
  * [A service principal](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object) with read-only permissions
    * If using Terraform to create this [resource](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/service\_principal), your user must have either the Application Administrator or Global Administrator directory role
* Access to the [Terraform CLI](https://www.terraform.io/downloads) or Azure CLI ([locally](https://learn.microsoft.com/en-us/cli/azure/) or via the [Cloud Shell](https://portal.azure.com/#home)) to create the above resources via Terraform or Bash script
  * If using Terraform or the Azure CLI locally, ensure you configure it with your Azure credentials. See instructions for [Terraform](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory) or the [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/authenticate-azure-cli)
* **API only:** An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
* **API only, optional**: [jq](https://stedolan.github.io/jq/), to unescape JSON containing the Terraform template or Bash script
