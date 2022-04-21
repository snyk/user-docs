---
description: The list of options and some examples
---

# Github Enterprise - Examples

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Github Enterprise token                    [required]
  --url                     Your Github host custom URL.
                            e.g https://ghe.prod.company.org/          [required]
  --orgs                    [Optional] A list of Github Enterprise organizations, separeted by comma, 
                            to fetch and count contributors for their repositories              
  --repo                    [Optional] Specific repo to count only for
  --fetchAllOrgs            [Optional] When enabled, will fetch all orgs that the token has access to
                            rather than fetching only the orgs your authorized to operate in.
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, requiered when using the "consolidateResults" command
```

**Before running the command:**

Get your Github Enterprise token or create a new one with this [guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

**Note: make sure your token has read access to the repos.**

## Running the command

Consider the following levels of usage and options:

### Usage levels

*   To get commits for all repos in all my orgs in Github Enterprise: provide the Github Enterprise token:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL
    ```
*   To get commits for some orgs and their repos in Github Enterprise: provide the Github Enterprise token and the org names, separated by a comma:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG_ONE,ORG_TWO,ORG_THREE
    ```
*   To get commits for only one repo in Github Enterprise: provide the Github Enterprise token, one org name and one repo name:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO
    ```

### Options

* To map all the orgs in Github Enterprise and not only the ones I have operate rights to => add the `--fetchAllOrgs` flag:

```
snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --fetchAllOrgs
```

*   To exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore(separated by commas) and apply the `--exclusionFilePath` with the path to that file:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG_ONE,ORG_TWO --exclusionFilePath PATH_TO_FILE
    ```
*   To set the output to json format: add the `--json` flag:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --json
    ```
* To run in debug mode for verbose output: prefix with `DEBUG=snyk*` :

```
DEBUG=snyk* snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO --exclusionFilePath PATH_TO_FILE --json
```
