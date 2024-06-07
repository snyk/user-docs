# Pull Requests

## Snyk Fix PRs

Fix pull or merge requests are created automatically by Snyk when new issues are identified in Project tests or a retest is run on a Project that has identified vulnerabilities. This feature applies to Projects imported through an SCM integration such as GitHub Enterprise or Azure.

For more information on how integrations use fix and upgrade pull requests, see [View and understand Snyk upgrade pull requests for integrations](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/introduction-to-git-repository-integrations/view-and-understand-snyk-upgrade-pull-requests.md).

For instructions on opening pull requests from a GitHub account, see [Opening fix and upgrade pull requests from a fixed GitHub account](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/introduction-to-git-repository-integrations/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md).

For a full description of Snyk Fix PRs, see [Snyk Fix Pull Requests](snyk-fix-pull-or-merge-requests/).

## PR Checks

Pull request checks are tests that run on generated pull requests to identify new issues with Projects. This allows you to prevent issues from being introduced into your code before merging to the main branch.

For a full description of PR Checks, see [Pull Request Checks](pull-request-checks/).

## The differences between Fix PRs and PR Checks

* Snyk Fix pull or merge requests are initiated by Snyk when a new issue is identified in a Project test or retest. PR Checks or merge requests are initiated by Snyk when there are code changes in an Organization or Project.
* Snyk Fix PRs are used for the **remediation** of issues. PR checks are used for the **prevention** of issues.
* Snyk Fix PRs are automatic or manual requests made to resolve issues with fixes available, for example, an upgrade or a patch. PR checks are tests made on code changes based on conditions you [configured](pull-request-checks/configure-pull-request-checks.md) during setup, which identify vulnerabilities that could affect your code before merge.
* Snyk Fix PRs trigger when an issue is detected from a daily scan or test. PR checks trigger when you make a code change in a PR.

