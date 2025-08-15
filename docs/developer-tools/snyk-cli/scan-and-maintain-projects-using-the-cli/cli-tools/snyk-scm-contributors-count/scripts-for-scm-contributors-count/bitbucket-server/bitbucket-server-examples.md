---
description: The list of options and some examples for Bitbucket Server
---

# Bitbucket Server - Examples

The following options are available for the `snyk-scm-contributors-count bitbucket-server` command:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Bitbucket server token                     [required]
  --url                     Bitbucket server base url e.g. (https://bitbucket.mycompany.com)         [required]
  --projectKeys             [Optional] Bitbucket server project key to count contributors for
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, requiered when using the "consolidateResults" command
  --skipSnykMonitoredRepos  [Optional] Skip Snyk monitored repos and count contributors for all repos
  --importConfDir           [Optional] Generate an import file with the unmonitored repos: A path to a valid folder for the generated import files
  --importFileRepoType      [Optional] To be used with the importConfDir flag: Specify the type of repos to be added to the import file. Options: all/private/public. Default: all
```

## Before running the command

1. Export SNYK\_TOKEN (if you want to get the contributors ONLY for repos that are already monitored by Snyk):
   * Make sure that your token has Group level access or use a service account's token that has Group level access. To learn more on how to create a service account, refer to [Service accounts](../../../../../../../implementation-and-setup/enterprise-setup/service-accounts/).
   * Copy the token value.
   *   Export the token in your environment:

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2. Get your Bitbucket Server token and URL:
   *   Create a Token if one does not exist, using this [guide](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver).

       **Note**: Make sure your token has read access to the repos.
   * The URL is the actual URL of your Bitbucket Server instance, for example, http://bitbucket-server.mycompany.com.

## Running the command

Consider the following levels of usage and options:

### Usage levels

*   To get commits for all projects and their repos in Bitbucket Server, provide the Bitbucket Server token and url:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL
    ```
*   To get commits for some projects and their repos in Bitbucket Server, provide the Bitbucket Server token, Bitbucket Server url ,and the projects, separated by a comma:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1,Key2...
    ```
*   To get commits for a specific repo in Bitbucket Serve, provide your Bitbucket Server token, Bitbucket Server url, a project, and a repo name:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1
    ```

### Options

*   To get all the commits from Bitbucket Server regardless of the repos that are already monitored by Snyk, add the `--skipSnykMonitoredRepos` flag.\
    You might have repos in Bitbucket Server that are not monitored in Snyk,. Use this flag to skip checking for Snyk monitored repos and go directly to Bitbucket Server to fetch the commits.

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits, add an exclusion file with the emails to ignore(separated by a new line) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1,Key2 --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --json
    ```
*   To create an import file for me with my unmonitored repos, add the `--importConfDir` flag with a valid (writable) path to a folder in which the import files will be stored and add the `--importFileRepoType` flag (optional) with the repo types to add to the file (`all`/`private`/`public`, defaults to `all`). Note that these flags **can not** be set with the `--repo` flag.

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --importConfDir ValidPathToFolder --importFileRepoType private/public/all
    ```

    For more information about these flag, refer to [Creating and using the import file](../../creating-and-using-the-import-file.md).
*   To run in debug mode for verbose output, add `DEBUG=snyk*` to the beginning of the command:

    ```
    DEBUG=snyk* snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    ```
