# Review pending findings

The self-review feature in Snyk API & Web gives you visibility and control over findings that are in a Pending Review state. This feature allows you to view and act on pending findings immediately, helping you speed up your security reviews and development workflows.

Previously, scans containing low-confidence findings were marked as Under Review, and those pending findings remained hidden while awaiting manual verification by the Snyk team.

## Prerequisites

To enable or disable this feature for your Organization, you must have Account Owner permissions.

## Enable the self-review feature

As an Account Owner, enable this feature in your account settings.

1. Navigate to **Settings > Scan Settings**.
2. Locate the **REVIEW PENDING FINDINGS** section.
3. Use the toggle to enable the self-review feature.

By enabling this feature, you acknowledge the following:

* **Concurrent reviews**: Your team and the Snyk team may review findings at the same time. Your team's decision always takes precedence.
* **Justification visibility**: The justification you provide for a review decision is visible to Snyk to help improve the products.

## Review pending findings

Once the feature is enabled, you can review findings individually or in bulk directly from the findings list. First, use the **State** or **Review status** filter to locate findings in a Pending Review state.

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

1. Select the checkboxes next to the findings you want to review.
2. Click **Review** that appears above the table.
3. Choose to **Approve** or **Reject** the selected findings.
4. Enter a single justification applied to all findings in the batch.
5. Click **Save**.

## Verify the outcome

After you review a pending finding, its status updates accordingly:

* Approved findings move to the Not fixed state and are treated as active vulnerabilities.
* Rejected findings are updated to the N/A state and are considered inactive.
* Your decision and justification are recorded in each finding's history log for auditing purposes.

## Additional information

* **Concurrent reviews**: If a Snyk team member reviews a finding, their decision is logged in the finding's history. However, you always have the final say and can override their decision at any time. A Snyk team member cannot override your decision.
* **SDK and CLI access**: You can also view and review pending findings using the Snyk API & Web SDK and CLI.
