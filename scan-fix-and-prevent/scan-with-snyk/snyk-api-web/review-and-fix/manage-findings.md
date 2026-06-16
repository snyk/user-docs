# Manage findings

During a scan, Snyk API & Web finds vulnerabilities within the target URLs. When a vulnerability is found, a finding is created. These findings are registered in Snyk API & Web, and you can perform the following actions:

* Change a finding's state.
* Change a finding's severity.
* Change a finding's assignee.
* Change a finding's label.
* Re-test a finding.
* Add a note to a finding.

## Change state

A finding's state can change automatically (by the scanner during a target scan or re-test) or manually (through user actions).

Using the Snyk API & Web interface, you can define a finding as **Accepted** if you acknowledge and accept its risk, or as **Invalid** if you consider it to be a false positive.

You can change the state from the following locations:

* From the side panel or small details screen, click the **three vertical dots** that appear on the bottom-right corner and choose the respective action.
* From the full details page, click the respective button on the bottom right of the screen.
* From the list of findings, click the **State** dropdown.

These actions are reflected in the **State** field shown below.

Visit [Finding states](finding-states.md) to learn more about the findings' states and how they can change.

## Change severity

Depending on the type of vulnerability found, its exploitability, impact, and scope, a CVSS score and severity classification are attributed to the finding, helping you prioritize vulnerability fixes.

While the CVSS score cannot be manually changed, you can still change the finding's severity. This can be done directly from the finding's details page:

* From the small details panel, click the **three vertical dots** to open the dropdown menu, and click the **Change severity** option.
* From the full details page, click the respective button on the bottom right of the screen.

Regardless of which method you choose, you can then define the intended value and save the change.

Once you save the change, Snyk API & Web will not change the severity back, so make sure you intend to make the change. Visit [Severity levels in findings](severity-levels-in-findings.md) for more information.

## Change assignee

After a target scan or a re-test, you may want to assign a vulnerability to be handled by a specific team member. This can be done either through the finding's details page by clicking the respective **pencil icon** next to the **Assignee** field, or through the target page or the scan results page by selecting one or more findings and clicking the respective dropdown.

Visit [Assign vulnerabilities to a team member](assign-vulnerabilities-to-team-member.md) to learn more about changing a finding's assignee.

## Change labels

To help you filter scan results, you can use finding labels. You can create and apply them to your findings:

* Through the finding's details page, click the **pencil icon** that appears next to the respective field.
* Through the findings list, click the **Set labels** dropdown.

## Re-test

After fixing vulnerabilities previously reported by Snyk API & Web, you can re-test them to confirm they can no longer be exploited and are resolved.

To start a re-test:

* Visit the finding's details and click the **Re-test** button.
* Access any list in which it is displayed (target page, scan results page, or findings list), select it, and click the **Re-test** button.

If during a re-test the scanner is not able to replicate the vulnerability, the finding is marked as **Fixed**. Otherwise, it remains listed as **Not Fixed** until it can no longer be replicated by the scanner.

## Add a note

When viewing a vulnerability's details, you can add comments or notes for your teammates. Scroll down to the bottom of the page, write the intended note, and click the **Add note** button. These notes are not a way to contact Snyk API & Web and should only be used to leave contextualized information available for your teammates.
