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
                            *note* for sub-level groups, provide the lowest level group name                                             
  --project                 [Optional] Your Gitlab project path with namespaces to count contributors for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, requiered when using the "consolidateResults" command
```

**Before running the command:**

Get your Gitlab token or create a new one with this [guide](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html).

**Note: make sure your token has read access to the repos.**

## Running the command

Consider the following levels of usage and options:

### Usage levels

*   To get commits for all groups and their projects in Gitlab: provide the Gitlab token (and server url for Gitlab Enterprise):

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --url URL
    ```
*   To get commits for some groups and their projects in Gitlab: provide the Gitlab token and the group names, separated by a comma:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
    ```

{% hint style="info" %}
Please mind that for nested groups, you need to provide the lowest level group name, for example: for `TopLevelGroup/MidLevelGroup/LowLevelGroup` only provide "LowLevelGroup" with the `--groups` flag
{% endhint %}

*   To get commits for a specific project in Gitlab: provide the Gitlab token and **one** group name and **one** project name:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP --project PROJECT
    ```

### Options

* To exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore(separated by commas) and apply the `--exclusionFilePath` with the path to that file:

```
snyk-scm-contributors-count gitlab --token TOKEN --projectKeys ID1,ID2,Path1/Namespace1 --exclusionFilePath PATH_TO_FILE
```

*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --json
    ```
* To run in debug mode for verbose output: prefix with`DEBUG=snyk*`:

```
DEBUG=snyk* snyk-scm-contributors-count gitlab --token TOKEN --url URL --exclusionFilePath PATH_TO_FILE --json
```
