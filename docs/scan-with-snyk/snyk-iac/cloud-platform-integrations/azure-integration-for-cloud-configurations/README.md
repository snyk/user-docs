# Azure integration for cloud configurations

Snyk integrates with your [Microsoft Azure](https://azure.microsoft.com/en-us/) subscription to find issues in your cloud configurations and generate cloud context to help you prioritize your vulnerabilities.

You can onboard an Azure subscription to Snyk using the following methods:

* [Snyk Web UI](azure-integration-web-ui/)
* [Snyk API](azure-integration-api/)

## Prerequisites for Azure integration for cloud configurations

To add an Azure integration for cloud configurations, you need the following:

* A Snyk Business or Enterprise [plan](https://snyk.io/plans/)
* A new Snyk Organization with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator [role](../../../../snyk-platform-administration/user-roles/pre-defined-roles.md)
* Access to a [Microsoft Azure](https://azure.microsoft.com/en-us/) subscription and associated credentials with permissions to create the following resources:
  * [An Active Directory (AD) application registration](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration)
  * [A federated identity credential](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation)\
    If you are using Terraform to create this [resource](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application_federated_identity_credential#api-permissions), your user must have either the Application Administrator or Global Administrator directory role.
  * [A service principal](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object) with read-only permissions\
    If you are .using Terraform to create this [resource](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/service_principal), your user must have either the Application Administrator or Global Administrator directory role.
* Access to the [Terraform CLI](https://www.terraform.io/downloads) or Azure CLI ([locally](https://learn.microsoft.com/en-us/cli/azure/) or via the [Cloud Shell](https://portal.azure.com/#home)) to create the above resources via Terraform or Bash script\
  If you are using Terraform or the Azure CLI locally, ensure you configure it with your Azure credentials. See the instructions for [Terraform](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory) or the [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/authenticate-azure-cli).
* API only: An Organization-level [service account](../../../../implementation-and-setup/enterprise-setup/service-accounts/) with an Org Admin role to use the Snyk API
* API only: An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
* API only, optional: [jq](https://stedolan.github.io/jq/), to unescape JSON containing the Terraform template or Bash script
