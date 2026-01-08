# Consistent Ignores for Snyk Code Pull Request Checks

## Pull Request Check default ignore behavior

When viewing a pull request (PR) check from Snyk in your integrated SCM, ignored findings do not contribute to the PR check outcome. A PR check does not fail due to an ignored finding. Additionally, the Snyk PR experience includes the count of active (unignored) findings in the summary comment and displays each finding as an inline comment within the PR.

If a finding is ignored after a PR check has already been completed, the PR check must be retriggered by committing again to the PR. Upon retriggering, the PR check the following changes occur:

• The ignored finding is no longer counted in the summary table.

• The inline comment for the ignored finding is collapsed by default and marked as resolved.

Ignores are respected in[ Snyk Code Pull Request Checks](../../../../scan-with-snyk/pull-requests/pull-request-checks/) regardless of whether they are created through [policy](./#manage-ignores-at-the-group-level-through-snyk-code-security-policies) or for an [individual `snyk/assets/finding/v1` value](./#manage-ignores-in-snyk-projects).

## Example: Snyk Pull Request Check with ignored finding

The following example shows the behavior of a Snyk Code issue in a Pull Request (PR) Check in GitHub before and after it is ignored.

In the following screenshots, a high-severity issue was detected, causing the Snyk PR Check to fail.

<figure><img src="../../../../.gitbook/assets/view-pr-check-details.png" alt=""><figcaption><p>PR check ignored finding overview in GitHub</p></figcaption></figure>

You can see all findings, including ignored ones, by selecting **View Details** from the PR page of your SCM integration.

<figure><img src="../../../../.gitbook/assets/finding-after-ignore-snyk-web-ui.png" alt=""><figcaption><p>Open and ignored finding in Snyk Web PR check UI</p></figcaption></figure>

Before being ignored, the finding appears as an inline comment in your SCM integration.

<figure><img src="../../../../.gitbook/assets/finding-before-ignore.png" alt=""><figcaption><p>PR check ignored finding details in GitHub</p></figcaption></figure>

After an ignore is applied and a PR check is re-triggered, the inline comment for the ignored finding is collapsed by default and marked as resolved.

<figure><img src="../../../../.gitbook/assets/finding-after-ignored-collapsed-view.png" alt=""><figcaption><p>Collapsed view of a finding in Github</p></figcaption></figure>
