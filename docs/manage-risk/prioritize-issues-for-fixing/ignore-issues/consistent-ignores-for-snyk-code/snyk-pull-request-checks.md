# Consistent Ignores for Snyk Code Pull Request Checks

## Pull Request Check default ignore behavior

When viewing a pull request (PR) check from Snyk in your integrated SCM, ignored findings do not contribute to the PR check outcome. A PR check does not fail due to an ignored finding. Additionally, the Snyk PR experience includes the count of active (unignored) findings in the summary comment and displays each finding as an inline comment within the PR.

If a finding is ignored after a PR check has already been completed, the PR check must be retriggered by committing again to the PR. Upon retriggering, the PR check the following changes occur:

• The ignored finding is no longer counted in the summary table.

• The inline comment for the ignored finding is collapsed by default and marked as resolved.

Ignores are respected in[ Snyk Code Pull Request Checks](../../../../scan-with-snyk/pull-requests/pull-request-checks/) regardless of whether they are created through [policy](./#manage-ignores-at-the-group-level-through-snyk-code-security-policies) or for an [individual `snyk/assets/finding/v1` value](./#manage-ignores-in-snyk-projects).

## Example: Snyk Pull Request Check with ignored finding

This example demonstrates Snyk Code behavior in a GitHub Pull Request (PR) Check before and after you ignore an issue.

Snyk detects a high-severity issue and fails the PR Check. To view all findings, including ignored issues, select **View Details** on the SCM integration PR page.

* **Before you ignore the issue:** The finding appears as an inline comment.
* **After you ignore the issue:** When you re-trigger the PR check, Snyk collapses the inline comment and marks it as resolved.
