---
description: The list of options and some examples
---

# Bitbucket Cloud - Examples

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --user                    Bitbucket cloud username                   [required]
  --password                Bitbucket cloud password                   [required]
  --workspaces              [Optional] Bitbucket cloud workspace name/uuid to count contributors for
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
2.  Get your Bitbucket-Cloud username (**not email**) and password.

    **Note**: Make sure your credentials have read access to the repos.

### Running the command

Consider the following levels of usage and options:

#### Usage levels:

*   To get commits for all workspaces and their repos in Bitbucket-Cloud: provide the Bitbucket-Cloud user and password:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD
    ```
*   To get commits for some workspaces and their repos in Bitbucket-Cloud: Provide the Bitbucket-Cloud user, Bitbucket-Cloud password and a comma-separated list of workspaces:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --workspaces Workspace1,Workspace2...
    ```
*   To get commits for a specific repo in Bitbucket-Cloud: Provide your Bitbucket-Cloud user, Bitbucket-Cloud password, a workspace and a repo name:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --workspaces Workspace1 --repo Repo1
    ```

#### Options:

*   To get all the commits from Bitbucket-Cloud regardless of the repos that are already monitored by Snyk (You might have repos in Bitbucket-Cloud that are not monitored in Snyk, using this flag will skip checking for Snyk monitored repos and will go directly to Bitbucket-Cloud to fetch tha commits): add the `--skipSnykMonitoredRepos` flag:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore (separated by commas) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --workspaces Workspace1,Workspace2 --exclusionFilePath PATH_TO_FILE
    ```
*   To se the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --workspaces Workspace1 --repo Repo1 --json
    ```
*   To create an import file for me with my unmonitored repos: add the `--importConfDir` flag  with a valid (writable) path to a folder in which the import files will be stored, and add the `--importFileRepoType` flag (optional) with the repo types to add to the file (all/private/public, defaults to all). (**Note that these flags can not be set with the `--repo` flag**):

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --importConfDir ValidPathToFolder --importFileRepoType private/public/all
    ```

    For more details about these flag, refer to this [page](../../creating-and-using-the-import-files.md)
*   To run in debug mode for verbose output, prefix with `DEBUG=snyk*`:

    ```
    DEBUG=snyk* snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --workspaces Workspace1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    ```
