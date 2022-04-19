## Gitlab
Available options:
```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Gitlab token                               [required]
  --url                     [Optional] Your Gitlab host custom URL. If no host was provided
                            it will default to https://gitlab.com/
  --groups                  [Optional] Your Gitlab groups names to count contributors for 
                            *note* for sub-level groups, please provide the lowest level group name                
  --project                 [Optional] Your Gitlab project path with namespace to count contributors for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, requiered when using the "consolidateResults" command
```

### Before running the command:
Get your Gitlab token or create a new one with this [guide](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) =>
    - Make sure that your token has access to the projects that you want to get the contributors count for

### Running the command

Consider the following levels of usage and options:

#### Usage levels:
- I want to get commits for all groups and their projects in Gitlab => Only provide the Gitlab token (and server url for Gitlab Enterprise): 
```
snyk-scm-contributors-count gitlab --token TOKEN --url URL
```

- I want to get commits for some groups and their projects in Gitlab => Provide the Gitlab token
  and the groups names, seperated by a comma:
```
snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
```

- I want to get commits for only one project in Gitlab => Provide the Gitlab token,
  one group name and one project name:
```
snyk-scm-contributors-count gitlab --token TOKEN --groups Group --project PROJECT
```

#### Options:
- I want to exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore(seperated by commas) and apply the `--exclusionFilePath` with the path to that file to the command:
```
snyk-scm-contributors-count gitlab --token TOKEN --projectKeys Path1/Namespace1 --exclusionFilePath PATH_TO_FILE
```

- I want the output to be in json format => add the `--json` flag to the command:
```
snyk-scm-contributors-count gitlab --token TOKEN --json
```

- I want to run in debug mode for verbose output => add `DEBUG=snyk*` to the beginning of the command:
```
DEBUG=snyk* snyk-scm-contributors-count gitlab --token TOKEN --url URL --exclusionFilePath PATH_TO_FILE --json
```