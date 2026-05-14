---
description: The list of options and some examples for Azure
---

# Azure - Examples

The following options are available for the `snyk-scm-contributors-count azure devops` command:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Azure DevOps token                         [required]
  --org                     Your Org name in Azure DevOps, for example, https://dev.azure.com/{OrgName}           [required]
  --projectKeys             [Optional] Azure DevOps project key/name to count
                            contributors for
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, required when using the "consolidateResults" command
  --skipSnykMonitoredRepos  [Optional] Skip Snyk monitored repos and count contributors for all repos
  --importConfDir           [Optional] Generate an import file with the unmonitored repos: A path to a valid folder for the generated import files
  --importFileRepoType      [Optional] To be used with the importConfDir flag: Specify the type of repos to be added to the import file. Options: all/private/public. Default: all
```

## Before running the command

1. Export SNYK\_TOKEN (if you want to get the contributors only for repos that are already monitored by Snyk):
   * Make sure that your token has Group level access or use a service account's token that has Group level access. To learn more about how to create a service account, refer to [Service accounts](../../../../../../../implementation-and-setup/enterprise-setup/service-accounts/).
   * Copy the token value.
   *   Export the token in your environment:

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2. Get your Azure Devops Token and Org:
   *   Create a Token if one does not exist, using this [guide](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page).

       **Note**: Make sure your token has read access to the repos.
   * Find your Org name in Azure listed on the left pane on the [Azure DevOps site](https://dev.azure.com).

## Running the command

Consider the following levels of usage and options:

### Usage levels

*   To get commits for all projects and their repos under my Org in Azure , only provide the Azure token and Azure Org:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG
    ```
*   To get commits for some projects and their repos under my Org in Azure , provide the Azure token, Azure Org and the project key/s separated by a comma:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1,Key2...
    ```
*   To get commits for a specific repo under my Org in Azure , provide the Azure token, Azure Org, a project key and a repo name:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1
    ```

### Options

*   To get all the commits from Azure regardless of the repos that are already monitored by Snyk, add the `--skipSnykMonitoredRepos` flag.

    You might have repos in Azure that are not monitored in Snyk; use this flag to skip checking for Snyk monitored repos and go directly to Azure to fetch the commits.

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits, add an exclusion file with the emails to ignore (separated by a new line) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1,Key2 --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1 --json
    ```
*   To create an import file for your unmonitored repos: add the `--importConfDir` flag with a valid (writable) path to a folder in which the import files will be stored and add the `--importFileRepoType` flag (optional) with the repo types to add to the file (`all`/`private`/`public`, defaults to `all`). Note that these flags **can not** be set with the `--repo flag.`

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --importConfDir ValidPathToWritableFolder --importFileRepoType private/public/all
    ```

    For more details about these flags, refer to the [Creating and using the import page](../../creating-and-using-the-import-file.md).
*   To run in debug mode for verbose output, prefix the command with`DEBUG=snyk*`:

    ```
    DEBUG=snyk* snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    ```
