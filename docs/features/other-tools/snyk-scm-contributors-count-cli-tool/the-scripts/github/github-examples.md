---
description: The list of options and some examples for GitHub
---

# GitHub - Examples

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   GitHub token                               [required]
  --orgs                    [Optional] A list of GitHub organizations, separeted by comma, 
                            to fetch and count contributors for their repositories              
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, required when using the "consolidateResults" command
```

## Before running the command

Get your GitHub token or create a new one with this [guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

**Note:** Make sure your token has read access to the repos.

## Running the command

Consider the following levels of usage and options:

### Usage levels

*   To get commits for all repos in all my orgs in GitHub, provide the GitHub token:

    ```
    snyk-scm-contributors-count github --token TOKEN
    ```
*   To get commits for some Orgs and their repos in GitHub:, provide the GitHub token and the org names, separated by a comma:

    ```
    snyk-scm-contributors-count github --token TOKEN --orgs ORG_ONE,ORG_TWO,ORG_THREE
    ```
*   To get commits for only one repo in GitHub, provide the GitHub token, one org name ,and one repo name:

    ```
    snyk-scm-contributors-count github --token TOKEN --orgs ORG --repo REPO
    ```

### Options

* To exclude some contributors from being counted in the commits , add an exclusion file with the emails to ignore(separated by a new line) and apply the `--exclusionFilePath` with the path to that file:

```
snyk-scm-contributors-count github --token TOKEN --orgs ORG_ONE,ORG_TWO --exclusionFilePath PATH_TO_FILE
```

*   To set the output to json format, dd the `--json` flag:

    ```
    snyk-scm-contributors-count github --token TOKEN --json
    ```
* To run in debug mode for verbose output, prefix with `DEBUG=snyk*`:

```
DEBUG=snyk* snyk-scm-contributors-count github --token TOKEN --orgs ORG --repo REPO --exclusionFilePath PATH_TO_FILE --json
```
