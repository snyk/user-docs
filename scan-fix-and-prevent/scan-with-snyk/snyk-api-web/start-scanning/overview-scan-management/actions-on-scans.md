---
description: The actions you can take on Snyk API and Web scans
nav_context: agnostic
---

# Actions on scans

The **Targets** section provides a set of actions to manage scans in Snyk API & Web.

On each line in the **Targets** list, choose the scan action to run for that target. You can also manage scans directly from the target details by clicking the target name.

Snyk provides the following actions to manage scans, regardless of where you are in the product:

| Action | What it does |
|--------|-------------|
| **Scan Now** | Starts a scan. |
| **Scan Later** | Schedules a scan. |
| **Stop** | Stops a running scan. |
| **Pause** | Pauses a running scan. If a scan is paused for seven days without being resumed, Snyk automatically stops it. |
| **Resume** | Resumes a paused scan. |

Depending on the state of the scan, you can perform the following actions:

* You can only start a scan if no scan is running on that target.
* You can only pause a scan if it is running.
* You can only resume a scan if it is paused.
* You can only stop a scan if it is running or paused.
* You can always schedule a scan.

## Summary of available actions by scan state

| Target scan state | Available actions |
|-------------------|-------------------|
| No scan is running. | Scan Now<br>Scan Later |
| The scan is running. | Stop<br>Pause<br>Scan Later |
| The scan is paused. | Resume<br>Stop<br>Scan Later |

Snyk starts or stops a scan as soon as you select those options. Pausing or resuming a scan can take a while, because the scanner must perform the necessary actions. Pausing and resuming scans is available only for some account plans.
