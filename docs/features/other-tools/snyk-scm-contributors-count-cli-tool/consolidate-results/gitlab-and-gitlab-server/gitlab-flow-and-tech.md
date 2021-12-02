# Gitlab - Flow and Tech

### Flow <a href="flow" id="flow"></a>

1. Fetch the monitored projects from Snyk (if the `skipSnykMonitoredRepos` flag was **not set** and the `SNYK_TOKEN` was exported).
2. Set Gitlab or Gitlab Server mode (if a host was provided or not through the `url` flag).
3. Fetch one/some/all the groups that the credentials have access to from the SCM and create a groups list.
4. Fetch one/all projects under the fetched/provided groups.
5. Remove the projects that are not monitored by Snyk (if the `skipSnykMonitoredRepos` flag was **not set** and the `SNYK_TOKEN` was exported) and create a Projects list.
6. Create an import file for unmonitored repos to use for easily importing repos into Snyk account (if the `importConfDir` flag was set)
7. Fetch the commits for the fetched/provided project/s and create a Contributors list.
8. Count the commits for the project/s by the contributors.
9. Remove the contributors that were specified in the exclusion file (if `the exclusionFilePath` flag was set and a valid path to a text file was provided).
10. Print the results.

### Gitlab API endpoints used <a href="bitbucket-cloud-api-endpoints-used" id="bitbucket-cloud-api-endpoints-used"></a>

* To get the groups paths from Gitlab if a group/s names were provided: `api/v4/groups?all_available=true&search={groupName}`
* To get the projects from Gitlab if a host url was **not** provided : `/api/v4/projects?membership=true`
* To get the projects from Gitlab Server if a host url **was** provided: `/api/v4/projects`
* To get the commits for the fetched/provided project/s list: `/api/v4/projects/{ProjectPath}/repository/commits?since=${threeMonthsDate}`
