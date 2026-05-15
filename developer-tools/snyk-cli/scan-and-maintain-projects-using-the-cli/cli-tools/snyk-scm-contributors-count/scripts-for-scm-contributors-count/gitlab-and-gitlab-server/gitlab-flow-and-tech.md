# GitLab - Flow and Tech

## Flow <a href="#flow" id="flow"></a>

1. Set GitLab or GitLab Server mode (if a host was provided or not through the `url` flag).
2. Fetch `one`/`some`/`all` the groups that the credentials have access to from the SCM and create a groups list.
3. Fetch `one`/`all` projects under the fetched/provided groups.
4. Create an import file for unmonitored repos to use for easily importing repos into your Snyk account (if the `importConfDir` flag was set)
5. Fetch the commits for the fetched/provided project/s and create a Contributors list.
6. Count the commits for the project/s by the contributors.
7. Remove the contributors that were specified in the exclusion file (if `the exclusionFilePath` flag was set and a valid path to a text file was provided).
8. Print the results.

## GitLab API endpoints used <a href="#bitbucket-cloud-api-endpoints-used" id="bitbucket-cloud-api-endpoints-used"></a>

* To get the groups paths from GitLab if a group/s names were provided: `api/v4/groups?all_available=true&search={groupName}`
* To get the projects from GitLab if a host url was **not** provided: `/api/v4/projects?membership=true`
* To get the projects from GitLab Server if a host url **was** provided: `/api/v4/projects`
* To get the commits for the fetched/provided project/s list: `/api/v4/projects/{ProjectPath}/repository/commits?since=${threeMonthsDate}`
