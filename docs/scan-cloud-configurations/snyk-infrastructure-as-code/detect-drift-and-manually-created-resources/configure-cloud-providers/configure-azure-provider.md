# Configure Azure provider

## Authentication of Azure provider

To use `iac describe`, set up credentials to make authenticated requests to your Azure account. Snyk retrieves configuration information from [environment variables](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization#use-environment-based-authentication).

For a guide to configuring Azure authentication, see the [Terraform documentation](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure).

```
# Here we use a service principal account with a client secret
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_SECRET=password\
  snyk iac describe --to="azure+tf"
```

You can also authenticate using the **az** CLI. Then you must specify only the `AZURE_SUBSCRIPTION_ID`:

```
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  snyk iac describe --to=azure+tf
```

## Least privilege policy​ <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The `iac describe` command needs read-only access to your account. If you want to scan your whole Azure account, set up the **Reader** role on your subscription, as shown in the following screenshot.

<figure><img src="https://docs.driftctl.com/assets/images/auth-d38df6fe7a4318ec9ebf82d0e5f9edae.png" alt="Set up Reader role for the Azure provider"><figcaption><p>Set up Reader role for the Azure provider</p></figcaption></figure>

You may want to scan only a resource group; you can assign the **Reader** role only on some restricted resource groups.
