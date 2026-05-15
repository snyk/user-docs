---
description: SCM-Contributors-Count Modes and Levels
---

# Usage

## General command

```
snyk-scm-contributors-count <command> <command-options>
```

**`<command>`** is one of:

* `azure-devops`
* `bitbucket-cloud`
* `bitbucket-server`
* `github`
* `github-enterprise`
* `gitlab`
* `consolidateResults`

**`<command-options>`**: see SCM-specific pages (example pages) in the [scripts](scripts-for-scm-contributors-count/) section.

## Modes

### Scoping usage prior to onboarding

This mode works only with Bitbucket and Azure.

Apply the `skipSnykMonitoredRepos` flag, for example:

```
snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --skipSnykMonitoredRepos
```

### Snyk license consumption

This mode works only with Bitbucket and Azure.

Make sure to export your `SNYK_TOKEN`, for example:

```
export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD
```

## Levels

### Top level

In this level of usage, the tool starts from the top of the SCM to get the Orgs/Groups, then goes down to the repo level to get all the repos, then counts the commits for the past 90 days.

To use this level, provide the credentials (and host/url where applicable), and the tool will get the contributors count for all your orgs/groups and all their repos, for example:

```
snyk-scm-contributors-count github --token TOKEN
```

### Mid level

In this level of usage, the tool starts from the Orgs/Groups that the user provides, then goes down to the repo level to get all the repos, then count the commits for the past 90 days.

To use this level, provide the credentials and a comma-separated list of groups or orgs for which you'd like to fetch the repos and their contributors count, for example:

```
snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
```

### Low level

In this level of usage, the tool focuses on only one repo for which to get the contributors count.

To use this level, provide the credentials (host/url where applicable), one org/group, and one repo, for example:

```
snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO
```

## Debug mode

Add `DEBUG=snyk*` to the beginning of the command, for example:

```
DEBUG=snyk* snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
```

## Additional optional flags

Additional flags can be set for the command:

* Create an **import file** with unmonitored repos data to use with the `snyk-api-import` tool and import the repos to `my Snyk account`. Works only with Bitbucket and Azure. Apply the `importConfDir` flag with a path to a valid and writable folder where you will save the import files. This flag correlates with the `importFileRepoType` flag.
* Choose which **types of repos to add to the import file**. Works only with Bitbucket and Azure. Apply the `importFileRepoType` flag with one of these options: `all,` `private`, or `public`**.**
* **Exclude committers from being counted**. Apply the `exclusionFilePath` flag to the command with a path to a text file that contains the emails of the committers that you would like to be excluded from the count.
* Output the summary and results in a json format. Apply the `json` flag to the command.

## The consolidateResults command

Used for consolidating results from several commands across different SCMs, into a single file with unique contributors count. Refer to the [command page](consolidate-results.md).
