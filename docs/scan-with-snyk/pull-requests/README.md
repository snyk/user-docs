# Pull Requests

## Snyk Fix PRs

Fix pull or merge requests are created automatically by Snyk when new issues are identified in Project tests or a retest is run on a Project that has identified vulnerabilities. This feature applies to Projects imported through an SCM integration such as GitHub Enterprise or Azure.

For more information on how integrations use fix and upgrade pull requests, see [Upgrade open source dependencies with automatic PRs](snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/upgrade-open-source-dependencies-with-automatic-prs.md).

For instructions on opening pull requests from a GitHub account, see [Opening fix and upgrade pull requests from a fixed GitHub account](snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md).

For a full description of Snyk Fix PRs, see [Snyk Fix Pull Requests](snyk-pull-or-merge-requests/).

## PR Checks

Pull request checks are tests that run on generated pull requests to identify new issues with Projects. This allows you to prevent issues from being introduced into your code before merging to the main branch.

PR Checks that are configured to “Only fail when the issues found have a fix available” rely on Snyk FixPR support and, therefore, will not alert for projects in languages that do not support FixPRs.&#x20;

For a full description of PR Checks, see [Pull Request Checks](pull-request-checks/).

## The differences between Fix PRs and PR Checks

* Snyk Fix pull or merge requests are initiated by Snyk when a new issue is identified in a Project test or retest. PR Checks or merge requests are initiated by Snyk when there are code changes in an Organization or Project.
* Snyk Fix PRs are used for the remediation of issues. PR checks are used for the prevention of issues.
* Snyk Fix PRs are automatic or manual requests made to resolve issues with fixes available, for example, an upgrade or a patch. PR checks are tests made on code changes based on conditions you [configured](pull-request-checks/configure-pull-request-checks.md) during setup, which identify vulnerabilities that could affect your code before the merge.
* Snyk Fix PRs trigger when an issue is detected from a daily scan or test. PR checks trigger when you make a code change in a PR.
