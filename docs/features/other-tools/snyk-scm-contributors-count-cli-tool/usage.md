---
description: Usage Modes and Levels
---

# Usage

### General command

```
snyk-scm-contributors-count <command> <command-options>
```

**`<command>`**is one of:

* `azure-devops`
* `bitbucket-cloud`
* `bitbucket-server`
* `github`
* `github-enterprise`
* `gitlab`

**`<command-options>`**: see SCM-specific pages (example pages) in the [scripts](the-scripts/) section.

### Modes

#### Scoping usage prior to onboarding

&#x20;Apply the `skipSnykMonitoredRepos` flag. For example:

```
snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --skipSnykMonitoredRepos
```

#### Snyk license consumption

Make sure to export your SNYK\_TOKEN.  For example:

```
export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD
```

### Levels

#### Top level

In this level of usage, the tool starts from the top of the SCM to get the Orgs/Groups, then down to the repo level to get all the repos, then counts the commits for the past 90 days.

To use this level, provide the credentials (and host/url where applicable) and the tool will get the contributors count for all your orgs/groups and all their repos. For example:

```
snyk-scm-contributors-count github --token TOKEN
```

#### Mid level

In this level of usage, the tool will start from the Orgs/Groups that the user provides, then down to the repo level to get all the repos and then count the commits for the past 90 days.

To use this level, provide the credentials, and a comma-separated list of groups or orgs for which you'd like to fetch the repos and their contributors count. For example:

```
snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
```

#### Low level

In this level of usage, the tool will focus on only one repo to get the contributors count for.

To use this level, provide the credentials (host/url where applicable), one org/group and one repo. For example:

```
snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO
```

### Debug mode

Add `DEBUG=snyk*` to the beginning of the command. For example:

```
DEBUG=snyk* snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
```

## Additional optional flags

Additional flags can be set to the command:

* **Create an import file with unmonitored repos data to use with the snyk-api-import tool and import the repos to my Snyk account - **Apply the** `importConfDir`** flag with a path to a valid and writable folder in which to save the import files. This flag correlates with the **`importFileRepoType`** flag.
* **Choose which types of repos to add to the import file** - Apply the **`importFileRepoType `**flag** **with one of these options :** all, private or public**
* **Exclude committers from being counted** - Apply the **`exclusionFilePath`** flag to the command with a path to text file that contains the emails of the committers that you would like to be excluded from the count.
* **Output the summary and results in a json format** - Apply the **`json`** flag to the command
