---
hidden: true
noIndex: true
---

# Review pending findings

Self-review in Snyk API & Web gives you visibility and control over findings in a **Pending Review** state. View and act on pending findings immediately to speed up security reviews and development workflows.

Previously, Snyk marked scans with low-confidence findings as **Under Review** and hid those pending findings until the Snyk team verified them manually.

## Prerequisites

To enable or disable this feature for your Organization, you must have Account Owner permissions.

## Enable the self-review feature

As an Account Owner, enable this feature in your account settings.

1. Navigate to **Settings > Scan Settings**.
2. Locate the **REVIEW PENDING FINDINGS** section.
3. Use the toggle to enable the self-review feature.

By enabling this feature, you acknowledge the following:

* Concurrent reviews: Your team and the Snyk team can review findings at the same time. The decision of your team always takes precedence.
* Justification visibility: The justification you provide for a review decision is visible to Snyk to help improve the products.

## Review pending findings

After you enable the feature, you can review findings individually or in bulk directly from the findings list. First, use the **State** or **Review status** filter to locate findings in a **Pending Review** state.

### Review a single finding

You can review a single finding from its details view.

1. From the findings list, open the finding details using one of two methods:
   * Click anywhere on the finding row to open the side drawer.
   * Click the finding name to open the full details page.
2. In the detail view, click **Review**.
3. Choose to **Approve** or **Reject** the finding.
4. Enter a justification for your decision and click **Save**.

### Review multiple findings in bulk

You can approve or reject multiple findings at once from the main findings list.

1. Select the check boxes next to the findings you want to review.
2. Click **Review** that appears above the table.
3. Choose to **Approve** or **Reject** the selected findings.
4. Enter a single justification applied to all findings in the batch.
5. Click **Save**.

## Verify the outcome

After you review a pending finding, its status updates accordingly:

* Approved findings move to the **Not fixed** state and are treated as active vulnerabilities.
* Snyk updates rejected findings to the **N/A** state and considers them inactive.
* Snyk records your decision and justification in the history log of each finding for auditing purposes.

## Additional information

* Concurrent reviews: If a Snyk team member reviews a finding, Snyk logs their decision in the history of the finding. However, you always have the final say and can override their decision at any time. A Snyk team member cannot override your decision.
* SDK and CLI access: You can also view and review pending findings using the Snyk API & Web SDK and CLI.
