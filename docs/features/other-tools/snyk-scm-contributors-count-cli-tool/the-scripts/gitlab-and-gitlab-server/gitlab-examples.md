---
description: The list of options and some examples
---

# Gitlab - Examples

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Gitlab token                               [required]
  --url                     [Optional] Your Gitlab host custom URL. If no host was provided
                            it will default to https://gitlab.com/
  --groups                  [Optional] Your Gitlab groups names to count contributors for 
                            *note* for sub-level groups, please provide the lowest level group name                                             
  --project                 [Optional] Your Gitlab project path with namespaces to count contributors for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output
  --skipSnykMonitoredRepos  [Optional] Skip Snyk monitored projects and count contributors for all projects
  --importConfDir           [Optional] Generate an import file with the unmonitored projects: A path to a valid folder for the generated import files
  --importFileRepoType      [Optional] To be used with the importConfDir flag: Specify the type of repos to be added to the import file. Options: all/private/public. Default: all
```

### Before running the command:&#x20;

1. Export SNYK\_TOKEN (if you want to get the contributors ONLY for projects that are already monitored by Snyk) =>
   * Go to [Snyk-account](https://app.snyk.io/account) and create a token if not already exists.
   * Copy the token value
   *   Export the token in your environment:&#x20;

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2.  Get your Gitlab token or create a new one with this [guide](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html)**:**

    **Note: **make sure your token has read access to the repos.

### Running the command

Consider the following levels of usage and options:

#### Usage levels:

*   To get commits for all groups and their projects in Gitlab: provide the Gitlab token (and server url for Gitlab Enterprise):

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --url URL
    ```
*   To get commits for some groups and their projects in Gitlab: provide the Gitlab token and the group names, separated by a comma:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
    ```

{% hint style="info" %}
Please mind that for nested groups, you need to provide the lowest level group name, for example: for `TopLevelGroup/MidLevelGroup/LowLevelGroup `only provide "LowLevelGroup" with the `--groups` flag
{% endhint %}

*   To get commits for a specific project in Gitlab: provide the Gitlab token and **one** group name and **one** project name:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP --project PROJECT
    ```

### Options:

*   I want to get all the commits from Gitlab regardless of the projects that are already monitored by Snyk (You might have projects in Gitlab that are not monitored in Snyk, using this flag will skip checking for Snyk monitored projects and will go directly to Gitlab to fetch the commits) => add the `--skipSnykMonitoredRepos` flag to the command:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --skipSnykMonitoredRepos
    ```
*   To exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore(separated by commas) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --projectKeys ID1,ID2,Path1/Namespace1 --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --json
    ```
*   To create an import file for me with my unmonitored projects: add the `--importConfDir` flag with a valid (writable) path to a folder in which the import files will be stored and add the `--importFileRepoType` flag (optional) with the projects types to add to the file (all/private/public, defaults to all). (**Note that these flags can not be set with the `--project` flag**):

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --importConfDir ValidPathToWritableFolder --importFileRepoType private/public/all
    ```

    For more details about these flag, refer to this [page](../../creating-and-using-the-import-files.md)
*   To run in debug mode for verbose output: prefix with`DEBUG=snyk*`:

    ```
    DEBUG=snyk* snyk-scm-contributors-count gitlab --token TOKEN --url URL --exclusionFilePath PATH_TO_FILE --json
    ```
