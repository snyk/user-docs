# Manage Targets

The targets page is where you manage every target you have configured. It lists all your managed targets and lets you act on them directly.

## What the targets list shows

Each row in the list carries the current state of one target:

| Column                        | What it tells you                                                                                                                                                |
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Target**                    | The name and main URL                                                                                                                                            |
| **Open findings by severity** | Active findings on that target, counted by severity, for example 5 critical, 9 high, 17 medium. These are findings that are still open, not the total ever found |
| **Created**                   | When the target was configured                                                                                                                                   |
| **Last completed scan**       | When the most recent scan finished. A date far in the past means the target's findings reflect an older version of the application                               |
| **Scan status and actions**   | Whether a scan is currently running, and the actions available                                                                                                   |

Above the list sits an aggregate view of your whole portfolio: fix rate, targets at risk, running scans, and open findings by severity.

## Actions on a target

### View the target

Opens the target, where you can see all its findings and its full scan history. Visit Interpret scan results.

### Edit the target

Change any part of the configuration: name, main URL, scope, users and credentials, or extra instructions.

Edits apply to future scans. A scan already running continues with the configuration it started with, and completed scans keep the configuration they ran under, so historical results stay reproducible.

Common reasons to edit a target:

| Change                          | Why                                                                                                         |
| ------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| Add a host to the scope         | A new service was deployed on a new hostname, or the first scan revealed a host being rejected at the proxy |
| Add a user                      | You configured one user initially and now want authorization testing, which needs at least two              |
| Rotate credentials              | The test account password changed, or the two-factor secret was re-enrolled                                 |
| Extend the reject list          | A scan reached something you would rather it did not touch                                                  |
| Sharpen the application context | Results were technically valid but not focused on what you care about                                       |
| Link a source repository        | Upgrades subsequent scans from blackbox to greybox                                                          |

### Cancel a running scan

Stops the scan. Findings already confirmed are preserved, so a canceled scan still gives you whatever agents had validated up to that point. See Monitor and manage scans.

### Delete the target

Removes the target and its scan history.

{% hint style="warning" %}
Deleting a target removes its assessment history along with it, including past findings and the ability to regenerate historical reports. If you need those results for an audit trail, download the reports before deleting. See Reports.
{% endhint %}

If a scan is running when you delete a target, you are asked to confirm.

## Keeping the target list healthy

Two patterns are worth watching for as your portfolio grows:

Stale targets. A target with a last completed scan several months old is reporting on an application that no longer exists in that form. Its open finding counts, and its contribution to your fix rate, are both misleading. Rescan it or delete it.

Targets that were never scanned successfully. A target that has been configured but has no completed scan usually has a configuration problem rather than a clean bill of health. Check whether authentication is succeeding and whether the application is in scope. See Configure authentication and Define the target scope.

