---
description: The list of options and some examples for Bitbucket Cloud
---

# Bitbucket Cloud - Examples

The following options are available for the `snyk-scm-contributors-count bitbucket-cloud` command:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --user                    Bitbucket cloud username                   [required]
  --password                Bitbucket cloud app password               [required]
  --workspaces              [Optional] Bitbucket cloud workspace name/uuid to count contributors for
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, requiered when using the "consolidateResults" command
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
2.  Get your Bitbucket Cloud username (**not email**) and [app password](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#authentication)

    **Note**: Make sure your credentials have read access to the repos.

## Running the command

Consider the following levels of usage and options:

### Usage levels

*   To get commits for all workspaces and their repos in Bitbucket Cloud, provide the Bitbucket Cloud user and app password:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD
    ```
*   To get commits for some workspaces and their repos in Bitbucket Cloud, provide the Bitbucket Cloud user, Bitbucket Cloud app password, and a comma-separated list of workspaces:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1,Workspace2...
    ```
*   To get commits for a specific repo in Bitbucket Cloud, provide the Bitbucket Cloud user, Bitbucket Cloud app password, a workspace, and a repo name:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1 --repo Repo1
    ```

### Options

*   To get all the commits from Bitbucket Cloud regardless of the repos that are already monitored by Snyk, add the `--skipSnykMonitoredRepos` flag.\
    You might have repos in Bitbucket Cloud that are not monitored in Snyk; use this flag to skip checking for Snyk monitored repos and go directly to Bitbucket Cloud to fetch the commits.

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits , add an exclusion file with the emails to ignore (separated by a new line),and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1,Workspace2 --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1 --repo Repo1 --json
    ```
*   To create an import file for your unmonitored repos, add the `--importConfDir` flag with a valid (writable) path to a folder in which the import files will be stored, and add the `--importFileRepoType` flag (optional) with the repo types to add to the file (`all`/`private`/`public`, defaults to `all`). Note that these flags **can not** be set with the `--repo` flag.

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --importConfDir ValidPathToFolder --importFileRepoType private/public/all
    ```

    For more information about these flags, refer to this [Creating and using the import page](../../creating-and-using-the-import-file.md).
*   To run in debug mode for verbose output, prefix with `DEBUG=snyk*`:

    ```
    DEBUG=snyk* snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    .
    ```
