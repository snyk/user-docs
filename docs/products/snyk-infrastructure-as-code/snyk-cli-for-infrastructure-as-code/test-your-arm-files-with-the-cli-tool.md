# Test your ARM files with the CLI tool

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI.

Snyk Infrastructure as Code for Azure Resource Manager (ARM) supports scanning json format.

{% hint style="info" %}
You can also scan bicep format by converting the configuration files to json by using the Bicep CLI.
{% endhint %}

You can use the CLI as follows:

## To test for an issue on specified json files:

From the CLI enter the following:

```
snyk iac test deploy.json
```

You can also specify multiple files by appending the file names after each other, such as:

```
snyk iac test file-1.json file-2.json
```

## To test for an issue on specified bicep files:

First of all, make sure you have the [Bicep CLI installed](https://docs.microsoft.com/en-us/azure/azure-resource-manager/bicep/install).

After installing the Bicep CLI navigate to the directory that contains your bicep files and convert the relevant bicep file to json by entering the following:

```
az bicep build -f deploy.bicep
```

From here you can scan the newly created json file just like any other file, from the CLI enter the following:

```
snyk iac test deploy.json
```
