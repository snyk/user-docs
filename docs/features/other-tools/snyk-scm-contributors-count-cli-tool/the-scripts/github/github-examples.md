## Github
Available options:
```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --token                   Github token                               [required]
  --orgs                    [Optional] A list of Github organizations, separated by comma, 
                            to fetch and count contributors for their repositories              
  --repo                    [Optional] Specific repo to count only for
  --exclusionFilePath       [Optional] Exclusion list filepath
  --json                    [Optional] JSON output, required when using the "consolidateResults" command
```

### Before running the command:
Get your Github token or create a new one with this [guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) =>
    - Make sure that your token has access to the repos that you want to get the contributors count for

### Running the command

Consider the following levels of usage and options:

#### Usage levels:
- I want to get commits for all repos in all my orgs in Github => Only provide the Github token: 
```
snyk-scm-contributors-count github --token TOKEN
```

- I want to get commits for some Orgs and their repos in Github => Provide the Github token
  and the org names, seperated by a comma:
```
snyk-scm-contributors-count github --token TOKEN --orgs ORG_ONE,ORG_TWO,ORG_THREE
```

- I want to get commits for only one repo in Github => Provide the Github token,
  one org name and one repo name:
```
snyk-scm-contributors-count github --token TOKEN --orgs ORG --repo REPO
```

#### Options:
- I want to exclude some contributors from being counted in the commits => add an exclusion file with the emails to ignore(seperated by commas) and apply the `--exclusionFilePath` with the path to that file to the command:
```
snyk-scm-contributors-count github --token TOKEN --orgs ORG_ONE,ORG_TWO --exclusionFilePath PATH_TO_FILE
```

- I want the output to be in json format => add the `--json` flag to the command:
```
snyk-scm-contributors-count github --token TOKEN --json
```

- I want to run in debug mode for verbose output => add `DEBUG=snyk*` to the beginning of the command:
```
DEBUG=snyk* snyk-scm-contributors-count github --token TOKEN --orgs ORG --repo REPO --exclusionFilePath PATH_TO_FILE --json
```