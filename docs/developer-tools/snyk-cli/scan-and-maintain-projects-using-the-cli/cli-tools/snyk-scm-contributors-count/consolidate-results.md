---
description: How to use the consolidateResults command
---

# Consolidate results

## consolidateResults command

When you work with the SCM-Contributors-Count tool, you may work with more than one source control manager (SCM). You will run a separate command for each SCM to get the contributors count for your repos there.

For example: If you have a contributor that commits to both GitHub repos and Bitbucket projects, you'll see that contributor's details in the output for both SCMs.

If you want to get an overall picture of all your contributors across all the SCMs, use the `consolidateResults` command.

This command allows you to take multiple (`json`) outputs of `snyk-scm-contributors-count` commands for different SCMs and consolidate them into one file, with a unique contributors count and a total of the repos from all SCMs.

The following options are available for the `consolidateResults` command:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --folderPath              Path to a folder containing the json outputs        [required]
```

## Running the commands

* Run the `snyk-scm-contributors-count` command for each repo with the `--json` flag and send the output to a designated folder, for example:

```
snyk-scm-contributors-count github --token TOKEN --json > PathToFolder/FileName
snyk-scm-contributors-count github-enterprise --token TOKEN --json > PathToFolder/OtherFileName
```

* Run the `consolidateResults` command and apply the `--folderPath` flag with the path to the designated, read/write accessible folder that contains the different output json files with the individual SCM results.

```
snyk-scm-contributors-count consolidateResults --folderPath PathToFolder
```

* The tool will then look for valid files in the folder, read the content of the files, create a new file with consolidated, unique results from all the files that have been read, and name the new file`consolidated-results.json`.
