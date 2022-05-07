# Bitbucket Server - Flow and Tech

## Flow <a href="#flow" id="flow"></a>

1. Fetch the monitored projects from Snyk (if the `skipSnykMonitoredRepos` flag was **not set** and the `SNYK_TOKEN` was exported).
2. Fetch `one`/`some`/`all` the projects that the credentials have access to from SCM and create a projects list.
3. Fetch `one`/`all` repos under the fetched/provided projects.
4. Remove the repos that are not monitored by Snyk (if the `skipSnykMonitoredRepos` flag was **not set** and the `SNYK_TOKEN` was exported) and create a Repo list.
5. Create an import file for unmonitored repos to use for easily importing repos into Snyk account (if the `importConfDir` flag was set).
6. Fetch the commits for the fetched/provided repo/s and create a Contributors list.
7. Count the commits for the repo/s by the contributors.
8. Remove the contributors that were specified in the exclusion file (if `the exclusionFilePath` flag was set and a valid path to a text file was provided).
9. Print the results.

## Bitbucket Server API endpoints used <a href="#bitbucket-cloud-api-endpoints-used" id="bitbucket-cloud-api-endpoints-used"></a>

* To get the repositories from BB Cloud, if a workspace was **not** provided: `/rest/api/1.0/repos`
* To get the repositories from BB Cloud, if a workspace/s **was** provided: `/rest/api/1.0/projects/{Project}/repos`
* To get the commits for the fetched/provided repo/s list: `/rest/api/1.0/projects/{Project}/repos/{Repo}/commits`
