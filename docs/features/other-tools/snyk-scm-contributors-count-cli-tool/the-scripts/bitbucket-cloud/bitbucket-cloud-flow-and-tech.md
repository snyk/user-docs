# Bitbucket Cloud - Flow and Tech

## Flow

1. Fetch the monitored projects from Snyk (if the `skipSnykMonitoredRepos` flag was **not set** and the `SNYK_TOKEN` was exported).
2. Fetch `one`/`some`/`all` the workspaces that the credentials have access to from SCM and create a workspaces list.
3. Fetch `one`/`all` repos under the fetched/provided workspaces.
4. Remove the repos that are not monitored by Snyk (if the `skipSnykMonitoredRepos` flag was **not set** and the `SNYK_TOKEN` was exported) and create a Repo list.
5. Create an import file for unmonitored repos to use for easily importing repos into your Snyk account (if the `importConfDir` flag was set)
6. Fetch the commits for the fetched/provided repo/s and create a Contributors list.
7. Count the commits for the repo/s by the contributors.
8. Remove the contributors that were specified in the exclusion file (if `the exclusionFilePath` flag was set and a valid path to a text file was provided).
9. Print the results.

## Bitbucket Cloud API endpoints used

* To get the repositories from BB Cloud, if a workspace was **not** provided: `/api/2.0/repositories`
* To get the repositories from BB Cloud, if a workspace/s **was** provided: `/api/2.0/repositories/{Workspace}`
* To get the commits for the fetched/provided repo/s list:`/api/2.0/repositories/{Workspace}/{Repo}/commits`
