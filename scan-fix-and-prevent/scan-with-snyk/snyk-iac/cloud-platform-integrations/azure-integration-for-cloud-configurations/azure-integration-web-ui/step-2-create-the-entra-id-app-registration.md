# Step 2: Create the Entra ID app registration

{% hint style="info" %}
**Recap**\
You have downloaded the Terraform template declaring the [Azure Active Directory (AD) application registration](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration), [federated identity credential](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation), and [service principal](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object) for Snyk. Now you need to provision the infrastructure.
{% endhint %}

To scan an Azure subscription, Snyk takes the permissions of a service principal with a [Reader](https://learn.microsoft.com/en-us/azure/role-based-access-control/built-in-roles#reader) role that allows Snyk to scan the configuration of your subscription resources.

Additionally, Snyk has a security feature that locks the federated credential for a subscription and tenant to the Organization that onboards it. This ensures that nobody can guess the credential's name and onboard it into a separate Organization to see those resources.

## Create the infrastructure with Terraform or the Azure CLI

You can create the app registration, federated identity credential, and service principal using one of the following tools, according to the type of file you downloaded from Snyk:

* [Terraform](step-2-create-the-entra-id-app-registration.md#create-azure-app-registration-infrastructure-using-terraform): Terraform CLI
* [Bash](step-2-create-the-entra-id-app-registration.md#create-azure-app-registration-infrastructure-using-the-azure-cli): Azure CLI, installed locally or accessed via Cloud Shell

### Create Azure app registration infrastructure using Terraform

{% hint style="info" %}
Before you use the [Terraform CLI](https://www.terraform.io/downloads), ensure you [configure it to use your Azure credentials](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory). Your user must have either the Application Administrator or Global Administrator directory role. This is required to create the [federated identity credential](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application_federated_identity_credential#api-permissions) and [service principal](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/service_principal) via Terraform.
{% endhint %}

1. In your terminal, navigate to the directory containing the Terraform file you downloaded (named `snyk-permissions-azure.tf` if downloaded from the Snyk Web UI).
2. Using the Terraform CLI, initialize the Terraform Project:

```
terraform init
```

3. Review and apply the Terraform plan:

```
terraform apply
```

4. Enter `yes` when Terraform asks if you want to perform the actions.

Terraform then provisions the infrastructure. When it is finished, you will see the following output:

```
Apply complete! Resources: 4 added, 0 changed, 0 destroyed.

Outputs:

application_id = 12345678-9012-3456-7890-12345678abcd
```

Copy the application ID for use in the next step.

### Create Azure app registration infrastructure using the Azure CLI

{% hint style="info" %}
If you use the [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) locally, ensure you have logged in with [`az login`](https://learn.microsoft.com/en-us/cli/azure/authenticate-azure-cli) first. This step is not necessary when you are using the CLI in the [Azure Cloud Shell](https://portal.azure.com/#cloudshell/).
{% endhint %}

1. In your terminal or in the [Cloud Shell](https://portal.azure.com/#cloudshell/), navigate to the directory containing your `.sh` file.
2. Make the Bash script executable using the `chmod` command:

```bash
chmod 755 snyk-permissions-azure.sh
```

3. Run the script:

```bash
./snyk-permissions-azure.sh
```

The Azure CLI then creates the AD app registration, federated identity credential, and service principal. When it is finished, you will see JSON output with information about the created infrastructure.

**Copy the application ID,** the last item in the output:

```json
{
  "application_id": "12345678-9012-3456-7890-12345678abcd"
}
```

Copy this application ID for use in the next step.

## What's next?

The next step is to create and scan the Cloud Environment. See [Step 3: Create and scan a Cloud Environment for Azure (Web UI)](step-3-create-and-scan-a-cloud-environment-for-azure-web-ui.md) or [Step 3: Create and scan a Cloud Environment for Azure (API)](../azure-integration-api/step-3-create-and-scan-a-cloud-environment-for-azure-api.md).
