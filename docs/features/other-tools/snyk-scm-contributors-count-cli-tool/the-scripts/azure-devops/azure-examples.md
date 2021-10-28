---
description: The list of options and some examples
---

# Azure - Examples

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Azure Devops token                         [required]
  --org                     Your Org name in Azure Devops e.g. https://dev.azure.com/{OrgName}           [required]
  --projectKeys             [Optional] Azure Devops project key/name to count
                            contributors for
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output
  --skipSnykMonitoredRepos  [Optional] Skip Snyk monitored repos and count contributors for all repos
  --importConfDir           [Optional] Generate an import file with the unmonitored repos: A path to a valid folder for the generated import files
  --importFileRepoType      [Optional] To be used with the importConfDir flag: Specify the type of repos to be added to the import file. Options: all/private/public. Default: all
```

### Before running the command:

1. Export SNYK\_TOKEN (if you want to get the contributors only for repos that are already monitored by Snyk):
   * Go to [Snyk-account](https://app.snyk.io/account) and create a token if not already exists.
   * Copy the token value
   *   Export the token in your environment:&#x20;

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2. Get your Azure-Devops Token and Org:
   *   Create a Token if one does not exist, using this [guide](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page)&#x20;

       **Note**: make sure your token has read access to the repos.
   * Your Org name in Azure is listed on the left pane in [Azure-Devops-site](https://dev.azure.com)

### Running the command

Consider the following levels of usage and options:

#### Usage levels:

*   I want to get commits for all projects and their repos under my Org in Azure => Only provide the Azure token and Azure Org:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG
    ```
*   I want to get commits for some projects and their repos under my Org in Azure => Provide the Azure token, Azure Org and the project key/s separated by a comma:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1,Key2...
    ```
*   I want to get commits for a specific repo under my Org in Azure => Provide your Azure token, Azure Org, a project key and a repo name:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1
    ```

#### Options:

*   To get all the commits from Azure regardless of the repos that are already monitored by Snyk: add the `--skipSnykMonitoredRepos` flag.

    (You might have repos in Azure that are not monitored in Snyk, using this flag will skip checking for Snyk monitored repos and will go directly to Azure to fetch the commits)

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits:  add an exclusion file with the emails to ignore (separated by commas) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1,Key2 --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1 --json
    ```
*   To create an import file for me with my unmonitored repos: add the `--importConfDir` flag  with a valid (writable) path to a folder in which the import files will be stored and add the `--importFileRepoType` flag (optional) with the repo types to add to the file (all/private/public, defaults to all). (**Note that these flags can not be set with the `--repo` flag**):

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --importConfDir ValidPathToWritableFolder --importFileRepoType private/public/all
    ```

    For more details about these flag, refer to this [page](../../creating-and-using-the-import-files.md)
*   To run in debug mode for verbose output: prefix the command with`DEBUG=snyk*`:

    ```
    DEBUG=snyk* snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    ```
