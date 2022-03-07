# Azure

## Authentication

To use describe, we need credentials to make authenticated requests to your Azure account. We retrieve configuration from [environment variables](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization#use-environment-based-authentication).

You can check the [Terraform documentation](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure) for a guide to configure Azure authentication.

```
# Here we use a service principal account with a client secret
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_SECRET=password\
  snyk iac describe --to=azure+tf
```

You can also authenticate using **az CLI**. In that case you will only have `AZURE_SUBSCRIPTION_ID` to specify:

```
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  snyk iac describe --to=azure+tf
```

### Least privileged policy[​](https://docs.driftctl.com/0.22.0/providers/azure/authentication#least-privileged-policy) <a href="#least-privileged-policy" id="least-privileged-policy"></a>

Describe needs to have read only access to your account, if you want to scan your whole Azure account you can set up the **Reader** role on your subscription.

![auth](https://docs.driftctl.com/assets/images/auth-d38df6fe7a4318ec9ebf82d0e5f9edae.png)

You may want to scan only a resource group, you can assign **Reader** role only on some restricted resources groups too.
