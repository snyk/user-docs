---
description: How to use this command
---

# Consolidate Results

### ConsolidateResults

Available options:

```
  --version                 Show version number                        [boolean]
  --help                    Show help                                  [boolean]
  --folderPath              Path to a folder containing the json outputs        [requiered]
```

#### Overview:

You might work with not just one source control manager (SCM) but a few, when working with this tool, you'll need to run a separate command for each SCM to get the contributors count for your repos there.

For example: If you have a contributor that commits to both Github repos and Bitbucket projects, you'll see his/hers details on the outputs of both SCMs.

What if you want to get an overall picture of all your contributors across all the SCMs that you work with?

This command allows you to take multiple (json) outputs of different commands for different SCMs and consolidate them into one file, with a unique contributors count and a total of the repos from all SCMs.

#### Running the command:

* Run the different commands with the `--json` flag and send the output to a designated folder, for example:

```
snyk-scm-contributors-count github --token TOKEN --json > PathToFolder/FileName
snyk-scm-contributors-count github-enterprise --token TOKEN --json > PathToFolder/OtherFileName
```

* Run the command and apply the --folderPath flag with the path to the designated, read/write accessible, folder that contains the different output json files with the individual SCM results.

```
snyk-scm-contributors-count consolidateResults --folderPath PathToFolder
```

* The tool will then look for valid files in the applied folder, read the content of the files and create a new file with consolidated, unique results from all the read files and name it "consolidated-results.json"
