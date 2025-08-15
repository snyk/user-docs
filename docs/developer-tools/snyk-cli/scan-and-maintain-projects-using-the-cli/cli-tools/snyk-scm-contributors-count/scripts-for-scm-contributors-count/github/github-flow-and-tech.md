# GitHub - Flow and Tech

## Flow <a href="#flow" id="flow"></a>

1. Fetch `one`/`some`/`all` the orgs that the credentials have access to from SCM and create an orgs list.
2. Fetch `one`/`all` repos under the fetched/provided orgs.
3. Fetch the commits for the fetched/provided repo/s and create a Contributors list.
4. Count the commits for the repo/s by the contributors.
5. Remove the contributors that were specified in the exclusion file (if `the exclusionFilePath` flag was set and a valid path to a text file was provided).
6. Print the results.

## GitHub API endpoints used <a href="#azure-api-endpoints-used" id="azure-api-endpoints-used"></a>

* To get the orgs from GitHub: `/user/orgs`
* To get the list of the repo/s that correlate with the fetched/provided orgs list: `/orgs/{Org}/repos`
* To get the commits for the fetched/provided repo/s list: `repos/{Org}/{Repo}/commits?since={threeMonthsDate}`
