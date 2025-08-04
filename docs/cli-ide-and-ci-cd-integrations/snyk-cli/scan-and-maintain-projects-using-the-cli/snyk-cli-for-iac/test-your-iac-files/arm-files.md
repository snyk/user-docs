# ARM files

With Snyk Infrastructure as Code, you can test your configuration files using the CLI.

Snyk Infrastructure as Code for Azure Resource Manager (ARM) supports scanning JSON format files. You can also scan Bicep format files by converting the configuration files to JSON using the Bicep CLI. Snyk supports the ARM `languageversion` 1.0.

## Test for an issue on specified JSON files

Enter the following Snyk CLI command:

```
snyk iac test deploy.json
```

You can also specify multiple files by appending the file names after each other, for example:

```
snyk iac test file-1.json file-2.json
```

## Test for an issue on specified Bicep files

Be sure you have the [Bicep CLI installed](https://docs.microsoft.com/en-us/azure/azure-resource-manager/bicep/install).

After installing the Bicep CLI, **navigate** to the directory that contains your Bicep files and **convert** the relevant Bicep file to JSON by entering the following:

```
az bicep build -f deploy.bicep
```

You can then scan the newly created JSON file in the same way as any other file. Use the following Snyk CLI command:

```
snyk iac test deploy.json
```
