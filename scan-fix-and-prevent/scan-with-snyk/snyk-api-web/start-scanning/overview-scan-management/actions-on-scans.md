# Actions on scans

A set of actions is available in the Targets section to manage scans in Snyk API & Web.

On each line in the Targets list, choose the scan action to execute for that target. You can also manage scans directly from the target details by clicking the target name.

Regardless of where you are in Snyk API & Web, you have the following actions to manage scans:

| Action | What it does |
|--------|-------------|
| **Scan Now** | Starts a scan. |
| **Scan Later** | Schedules a scan. |
| **Stop** | Stops a running scan. |
| **Pause** | Pauses a running scan. If a scan is paused for seven days without being resumed, Snyk API & Web automatically stops it. |
| **Resume** | Resumes a paused scan. |

Depending on the state of the scan, you may be able to perform the following actions:

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

While starting or stopping a scan is generally executed as soon as you select those options, pausing or resuming a scan might take a while for the scanner to perform the necessary actions. Pausing and resuming scans is only available for some account plans.
