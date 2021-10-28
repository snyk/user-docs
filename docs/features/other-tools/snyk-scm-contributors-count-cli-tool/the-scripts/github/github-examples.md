---
description: The list of options and some examples
---

# Github - Examples

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Github token                               [required]
  --orgs                    [Optional] A list of Github organizations, separeted by comma, 
                            to fetch and count contributors for their repositories              
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output
  --skipSnykMonitoredRepos  [Optional] Skip Snyk monitored repos and count contributors for all repos
  --importConfDir           [Optional] Generate an import file with the unmonitored repos: A path to a valid folder for the generated import files
  --importFileRepoType      [Optional] To be used with the importConfDir flag: Specify the type of repos to be added to the import file. Options: all/private/public. Default: all
```

### Before running the command:&#x20;

1. Export SNYK\_TOKEN (if you want to get the contributors only for repos that are already monitored by Snyk):
   * Go to [Snyk-account](https://app.snyk.io/account) and create a token if not already exists.
   * Copy the token value
   *   Export the token in your environment:&#x20;

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2.  Get your Github token or create a new one with this [guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

    **Note:** make sure your token has read access to the repos.

### Running the command

Consider the following levels of usage and options:

#### Usage levels:

*   To get commits for all repos in all my orgs in Github: provide the Github token:

    ```
    snyk-scm-contributors-count github --token TOKEN
    ```
*   To get commits for some Orgs and their repos in Github: Provide the Github token and the org names, separated by a comma:

    ```
    snyk-scm-contributors-count github --token TOKEN --orgs ORG_ONE,ORG_TWO,ORG_THREE
    ```
*   To get commits for only one repo in Github: Provide the Github token, one org name and one repo name:

    ```
    snyk-scm-contributors-count github --token TOKEN --orgs ORG --repo REPO
    ```

### Options:

*   To get all the commits from Github regardless of the repos that are already monitored by Snyk (You might have repos in Github that are not monitored in Snyk, using this flag will skip checking for Snyk monitored repos and will go directly to Github to fetch the commits) => add the `--skipSnykMonitoredRepos` flag:

    ```
    snyk-scm-contributors-count github --token TOKEN --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore(separated by commas) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count github --token TOKEN --orgs ORG_ONE,ORG_TWO --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count github --token TOKEN --json
    ```
*   To create an import file for me with my unmonitored repos: add the `--importConfDir` flag  with a valid (writable) path to a folder in which the import files will be stored and add the `--importFileRepoType` flag (optional) with the repo types to add to the file (all/private/public, defaults to all). (**Note that these flags can not be set with the `--repo` flag**):

    ```
    snyk-scm-contributors-count github --token TOKEN --importConfDir ValidPathToWritableFolder --importFileRepoType private/public/all
    ```

    For more details about these flag, refer to this [page](../../creating-and-using-the-import-files.md)
*   To run in debug mode for verbose output: prefix with `DEBUG=snyk*`:

    ```
    DEBUG=snyk* snyk-scm-contributors-count github --token TOKEN --orgs ORG --repo REPO --exclusionFilePath PATH_TO_FILE --json
    ```
