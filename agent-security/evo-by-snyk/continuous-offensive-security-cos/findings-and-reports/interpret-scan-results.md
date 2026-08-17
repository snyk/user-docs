# Interpret Scan Results

When you open a target, you get every finding on it, plus a summary of the target's current state. You can also open an individual scan to see only what that scan found.

## The target summary

At the top of the target sits its current state:

| Element                  | What it tells you                                                                                                                         |
| ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------- |
| **Linked repositories**  | Whether a source repository is connected. When one is, findings carry file and line level remediation guidance rather than generic advice |
| **Last completed scan**  | When the target was last assessed. Findings reflect the application as it was at that point, not necessarily as it is now                 |
| **Scan duration**        | How long the last scan took. Useful as a baseline for spotting a scan that is running unusually long                                      |
| **Findings by severity** | Total open vulnerabilities, counted by severity                                                                                           |

{% hint style="info" %}
The severity counts here are open findings, not everything ever found. A finding that a later scan could no longer exploit is no longer counted, which is why these numbers fall as you fix things.
{% endhint %}

## The findings list

Every finding on the target appears in the list:

| Column          | What it tells you                                          |
| --------------- | ---------------------------------------------------------- |
| **Finding**     | What the vulnerability is                                  |
| **Location**    | Where it was found: the URL or endpoint that was exploited |
| **Severity**    | How serious it is. Visit Severity and scoring              |
| **First found** | When any scan first confirmed it                           |

Open any finding for its full detail, including the steps agents took and the proof of work. See [Understand finding details](understand-finding-details.md).

## Filter and sort

The list supports filtering and sorting so you can work through it in a deliberate order rather than top to bottom.

| Approach                              | Use when                                                                                                                                                  |
| ------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Sort by severity**                  | Standard triage. Work critical and high first                                                                                                             |
| **Filter by severity**                | You want to see only what needs attention this week                                                                                                       |
| **Sort by first found, oldest first** | Finding long-standing issues that have been repeatedly deferred. A critical finding first seen months ago is a process problem as well as a technical one |
| **Sort by first found, newest first** | Reviewing what the most recent scan introduced, typically after a release                                                                                 |
| **Filter by location**                | Concentrating on one service or one part of the application, or routing work to the team that owns it                                                     |

A practical pattern for a large first result set: filter to critical and high, work that list to completion, then widen. Everything in the list was exploited, so there is no low-value verification work to clear first.

## Viewing results for one scan

The target's findings list aggregates the current state across all scans. To see one scan in isolation, open that scan. It shows the findings identified during that assessment, with location, severity, first found, and last found.

Use the per-scan view when your question is about a point in time:

* What did the assessment we ran before the last release find?
* What did this particular scan cover, given it was canceled partway through?
* Which findings correspond to the report we sent to the auditor?

Use the target's list when your question is about the present state of the application.

## Scan history

The target shows every scan it has run:

| Column               | What it tells you                               |
| -------------------- | ----------------------------------------------- |
| **Date**             | When the scan started                           |
| **Ended**            | When it finished                                |
| **Findings**         | How many findings that scan produced            |
| **Duration**         | How long it took                                |
| **Endpoints tested** | Coverage achieved                               |
| **Status**           | Running, queued, completed, canceled, or failed |
| **Actions**          | Cancel a running scan, or download the report   |

Reading the history as a series is more informative than reading any single row. Compare endpoints tested across scans: a scan with markedly lower coverage than its predecessors had a configuration problem, and its lower finding count reflects reduced coverage rather than improved security. See [Monitor and manage scans](../scans/monitor-and-manage-scans.md).

## Triage begins when the scan completes

While a scan runs, only high level progress is available: how far it has advanced, how many findings are confirmed, and their severities. Finding detail becomes available once the scan finishes, so triage is work you do on a completed scan rather than alongside a running one. See Monitor and manage scans.

