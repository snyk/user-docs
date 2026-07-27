---
description: How to manage Snyk API and Web findings
nav_context: classic
---

# Manage findings

During a scan, Snyk API & Web finds vulnerabilities within the target URLs. When Snyk finds a vulnerability, it creates a finding. Snyk registers these findings, and you can perform the following actions:

* Change a finding's state.
* Change a finding's severity.
* Change a finding's assignee.
* Change a finding's label.
* Re-test a finding.
* Add a note to a finding.

## Change state

The state of a finding can change automatically (by the scanner during a target scan or re-test) or manually (through user actions).

Using the Snyk interface, you can define a finding as **Accepted** if you acknowledge and accept its risk, or as **Invalid** if you consider it to be a false positive.

You can change the state from the following locations:

* From the side panel or small details screen, click the **three vertical dots** that appear in the bottom-right corner and select the respective action.
* From the full details page, click the respective button on the bottom right of the screen.
* From the list of findings, click the **State** dropdown.

These actions are reflected in the **State** field.

Visit [Finding states](finding-states.md) to learn more about finding states and how they can change.

## Change severity

Depending on the type of vulnerability found, its exploitability, impact, and scope, Snyk attributes a CVSS score and severity classification to the finding, helping you prioritize vulnerability fixes.

You cannot manually change the CVSS score, but you can change the severity of the finding. Change it directly from the details page of the finding:

* From the small details panel, click the **three vertical dots** to open the dropdown menu, and select **Change severity**.
* From the full details page, click the respective button on the bottom right of the screen.

Regardless of which method you choose, you can then define the intended value and save the change.

After you save the change, Snyk does not change the severity back, so ensure you intend to make the change. Visit [Severity levels in findings](severity-levels-in-findings.md) for more information.

## Change assignee

After a target scan or a re-test, you can assign a vulnerability to be handled by a specific team member. Assign it from the details page of the finding by clicking the respective **pencil icon** next to the **Assignee** field, or from the target page or the scan results page by selecting one or more findings and clicking the respective dropdown.

Visit [Assign vulnerabilities to a team member](assign-vulnerabilities-to-team-member.md) to learn more about changing a finding's assignee.

## Change labels

To help you filter scan results, you can use finding labels. You can create and apply them to your findings:

* From the details page of the finding, click the **pencil icon** that appears next to the respective field.
* From the findings list, click the **Set labels** dropdown.

## Re-test

After fixing vulnerabilities previously reported by Snyk, you can re-test them to confirm they can no longer be exploited and are resolved.

To start a re-test:

* Visit the details of the finding and click **Re-test**.
* Access any list in which it is displayed (target page, scan results page, or findings list), select it, and click **Re-test**.

If the scanner cannot replicate the vulnerability during a re-test, Snyk marks the finding as **Fixed**. Otherwise, it remains listed as **Not fixed** until the scanner can no longer replicate it.

## Add a note

When viewing the details of a vulnerability, you can add comments or notes for your teammates. Scroll down to the bottom of the page, write the intended note, and click **Add note**. These notes are not a way to contact Snyk. Use them only to leave contextualized information available for your teammates.
