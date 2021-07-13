# Snyk checks on pull requests

* [ GitHub integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032117-GitHub-integration/README.md)
* [ GitHub Enterprise Integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360015951318-GitHub-Enterprise-Integration/README.md)
* [ Bitbucket Cloud integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032097-Bitbucket-Cloud-integration/README.md)
* [ Bitbucket Data Center / Server integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002218-Bitbucket-Data-Center-Server-integration/README.md)
* [ GitLab integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002238-GitLab-integration/README.md)
* [ Azure Repos integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002198-Azure-Repos-integration/README.md)
* [ GitHub Read-Only Projects](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360010766557-GitHub-Read-Only-Projects/README.md)
* [ Opening fix and upgrade pull requests from a fixed GitHub account](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360010843797-Opening-fix-and-upgrade-pull-requests-from-a-fixed-GitHub-account-/README.md)
* [ Test your PRs for vulnerabilities before merging](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360006528057-Test-your-PRs-for-vulnerabilities-before-merging/README.md)
* [ Snyk checks on pull requests](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360006581938-Snyk-checks-on-pull-requests/README.md)

  [See more](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/sections/360001138098-Git-repository-SCM-integrations/README.md)

## Snyk checks on pull requests

By default, Snyk scans every pull request submitted on your monitored repositories, displaying the results and recommendations grouped together in a single security check and a single license check, regardless of the number of manifest files in the repository.

Administrators and account owners manage settings for Snyk PR tests from our app on both the organization and the project levels, configuring whether the feature is on \(enabled by default\) and under what conditions Snyk should fail your PR checks.

If a test fails for any of the lines in your pull request, the check itself appears as failed from the pull request; if all of the tests pass, the check itself appears as successful from the pull request:

To view the check results for all of the manifest files, click the Details link for the full list of tests and the results per file, directly from our interface.

From this view, click the links for additional information as follows:

* Click the repository link \(1\) to go back to the Git repository
* Click the Organization link \(2\) to view all projects in this Snyk organization
* Click the manifest file link \(3\) to view the Project details page with full details for all vulnerabilities affecting this project
* Click the View test page link \(4\) to view full details regarding this pull request and the issues preventing the check from passing

When Snyk tests your pull requests, the following are the possible statuses that can be displayed from this page, in the Results field:

* Success - no issues are identified and all checks pass
* Processing - this status appears until the Snyk test ends
* Failure - when issues are identified that must be fixed in order for the check to pass
* Error - an error occurs when your manifest file is out of sync, Snyk couldn't read the manifest file, or Snyk couldn't find the manifest file
* Canceled - Snyk test can't run because you've reached your monthly test limit

