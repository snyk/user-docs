# Getting started

### Configure the local environment

Most of the work we will do will involve using the [Azure Command-Line Interface \(CLI\)](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest). Detailed documentation on installing the Azure CLI for [Windows](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-windows?view=azure-cli-latest), [macOS](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-macos?view=azure-cli-latest), and [Linux](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-yum?view=azure-cli-latest) is available in [Azure documentation](https://docs.microsoft.com/en-us/azure/). These examples will be based on macOS. Additionally, sample code, templates, and other resources are provided in a [repository](https://github.com/snyk-partners/snyk-azure-resources) for this workshop. You are encouraged to [`clone`](https://github.com/snyk-partners/snyk-azure-resources.git) or [`fork`](https://github.com/snyk-partners/snyk-azure-resources/fork) this repository as we will reference that content throughout these exercises.

#### Install Homebrew

If you don't already have it, [install Homebrew](https://docs.brew.sh/Installation.html) then install the Azure CLI with the following command:

```bash
brew update && brew install azure-cli
```

#### Authenticate with the Azure CLI

Once installed, you will need to sign in to your Azure account from the CLI. Run the following command:

```bash
az login
```

The CLI will attempt to open your default browser and load the Azure login page. Provide your Azure account credentials in the browser and upon successful authentication you will see the following response in your browser window:

![](../../.gitbook/assets/azure_cli_login.png)

You should also see output similar to the following in your terminal:

```javascript
[
  {
    "cloudName": "AzureCloud",
    "homeTenantId": "<guid>",
    "id": "<guid>",
    "isDefault": true,
    "state": "Enabled",
    "tenantId": "<guid>",
    "user": {
      "name": "developer@microsoft.com",
      "type": "user"
    }
  }
]
```

If you encounter a problem, please review the [Install Azure CLI on macOS](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-macos?view=azure-cli-latest) documentation pages for additional guidance.

## 

