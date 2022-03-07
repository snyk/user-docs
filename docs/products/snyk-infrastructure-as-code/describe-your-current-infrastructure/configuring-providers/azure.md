# Azure

## Authentication

To use `iac describe`, set up credentials to make authenticated requests to your Azure account. Snyk retrieves configuration information from [environment variables](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization#use-environment-based-authentication).

See the [Terraform documentation](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure) for a guide to configure Azure authentication.

```
# Here we use a service principal account with a client secret
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_SECRET=password\
  snyk iac describe --to=azure+tf
```

You can also authenticate using **az CLI**. In that case you will need to specify only  `AZURE_SUBSCRIPTION_ID`:

```
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  snyk iac describe --to=azure+tf
```

## Least privileged policy[​](https://docs.driftctl.com/0.22.0/providers/azure/authentication#least-privileged-policy) <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The `iac describe` command needs read only access to your account, if you want to scan your whole Azure account you can set up the **Reader** role on your subscription.

![Setting up the Reader role](https://docs.driftctl.com/assets/images/auth-d38df6fe7a4318ec9ebf82d0e5f9edae.png)

You may want to scan only a resource group. You can also assign the **Reader** role only on some restricted resources groups.
