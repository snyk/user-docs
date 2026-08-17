# Monitor and Manage Scans

The scans and jobs view gives you access to every scan performed across Evo, in one place. Where a target shows you the scan history for that one application, this view shows all activity regardless of which target it belongs to.

### What the view shows

| Column        | What it tells you                               |
| ------------- | ----------------------------------------------- |
| **Status**    | Whether the job is running, queued, or canceled |
| **Type**      | What kind of job it is                          |
| **Initiator** | Who or what started it                          |
| **Target**    | Which target the scan is running against        |
| **Progress**  | How far the scan has advanced                   |

### When to use it

Use the target's own scan list when you are working on one application. Use scans and jobs when your question spans targets:

| Question                                            | Why this view answers it                                                         |
| --------------------------------------------------- | -------------------------------------------------------------------------------- |
| What is running across the whole program right now? | Shows every running scan at once, rather than checking targets one at a time     |
| Who started this scan?                              | The initiator tells you whether a person triggered it or it came from automation |
| Why has my scan not started yet?                    | A queued status means it is accepted and waiting, not failed                     |
| What has been canceled recently, and by whom?       | Cancellations across all targets, with their initiator                           |
| Is scan activity concentrated on one target?        | Shows the distribution of activity across the portfolio                          |

### Reading the initiator

The initiator identifies what started the job. This is the fastest way to answer questions about unexpected activity: a scan you did not expect is usually explained by either a colleague or an automated trigger, and the initiator tells you which without having to ask.

It is also what makes shared ownership workable. When several people manage the same set of targets, the initiator is the record of who ran what.

### Reading the status

| Status       | Meaning                                                                          |
| ------------ | -------------------------------------------------------------------------------- |
| **Running**  | Executing now. Open the scan for per-stage progress                              |
| **Queued**   | Accepted and waiting to start. No action needed                                  |
| **Canceled** | Stopped before completion. Confirmed findings are preserved, coverage is partial |

A queued scan has not failed. It has been accepted and will start on its own.

One common cause of a queued scan is the target already having a scan running, since a target can only run one scan at a time.

### Scans and other Evo jobs

This view covers jobs across Evo, not only Snyk Continuous Offensive Security scans. Snyk Continuous Offensive Security scans appear here alongside other Evo activity, which gives you a single place to see everything Evo is doing on your behalf. Use the target and type information to tell them apart.

## Monitor scan progress

Scans take time, so it is worth knowing how to read progress while one runs. Doing so lets you catch a misconfigured target early instead of discovering it hours later from a report with nothing in it.

### Where to look

| View                 | What it gives you                                                                            |
| -------------------- | -------------------------------------------------------------------------------------------- |
| **Scan detail**      | Per-stage progress for one scan: recon counts, attack counts, findings so far, report status |
| **Target scan list** | Every scan for one target, with status and outcome. See Interpret scan results               |
| **Targets list**     | Which of your targets have a scan running right now. See Manage targets                      |
| **Scans and jobs**   | Every scan across Evo, with status, initiator, target, and progress                          |

### Progress reported during a scan

The scan detail view reports progress per stage.

| Metric                        | Stage                 | How to read it                                                                                         |
| ----------------------------- | --------------------- | ------------------------------------------------------------------------------------------------------ |
| **Pages and endpoints found** | Recon                 | The size of the discovered attack surface. Compare against what you know your application to contain   |
| **Attacks performed**         | Attack and validation | How much of the discovered surface has been tested, including attacks in progress and attacks finished |
| **Findings so far**           | Attack and validation | Confirmed, exploited vulnerabilities. This count only includes attacks that succeeded                  |
| **Report generated**          | Reporting             | Whether the downloadable report is ready                                                               |

### What is visible while a scan runs

Progress reporting during a scan is deliberately high level. You can see how far the scan has advanced, how many findings have been confirmed, and at what severity. You cannot yet open a finding to read its steps, proof of work, or remediation guidance.

That means triage begins when the scan completes. What in-flight progress is for is a different job: confirming the scan is configured correctly and working, early enough to cancel and fix the target if it is not. Given a scan takes a few hours, that check is worth making rather than discovering a misconfiguration at the end.

### Reading the numbers on a first scan

On a first scan of a target you have no baseline, so use your own knowledge of the application.

Check the recon count first. This is the highest-value early signal, because attacks only ever cover what recon discovered.

| Recon result                                       | What it suggests                                        |
| -------------------------------------------------- | ------------------------------------------------------- |
| Roughly the surface you expect                     | Configuration is working. Let the scan proceed          |
| Far fewer pages and endpoints than expected        | Something is blocking discovery. Check the causes below |
| Only your login page, or a handful of public pages | Authentication is almost certainly failing              |

When recon looks too small, check in this order:

1. **Authentication.** The most common cause, and it fails quietly: agents never get past the login page and no error tells you so. Verify the credentials, and for two-factor accounts verify the secret or seed rather than a code. See Configure authentication.
2. **Identity provider scope.** If login redirects to a third-party provider that is not in scope, the redirect is rejected at the proxy and authentication cannot complete.
3. **Application scope.** An API or service on a hostname you did not add to the scope is rejected at the proxy, so nothing behind it is discovered. This is the most common reason a front end scans fine but produces no serious findings. See Define the target scope.
4. **Firewall or WAF blocking.** If a firewall starts rate limiting or blocking scan traffic, discovery stalls. Configure a header the firewall allows, or allow the scanner IP addresses. See Add a target and Scanner IP addresses.

If any of these apply, cancel the scan, fix the target configuration, and start again. There is no value in letting a scan run for hours against a surface it cannot reach.

### Reading the numbers on a repeat scan

Once a target has scan history, compare against the previous scan rather than against intuition.

| Observation                      | What it suggests                                                                              |
| -------------------------------- | --------------------------------------------------------------------------------------------- |
| Recon count similar to last scan | Configuration is still valid                                                                  |
| Recon count dropped sharply      | Credentials expired, a host moved, or the application changed. Check authentication and scope |
| Recon count rose sharply         | New surface was deployed. Worth confirming the new hosts are all in scope                     |
| Fewer findings than last scan    | Fixes are landing, or coverage dropped. Check the recon count to tell which                   |
| More findings than last scan     | New surface, new vulnerabilities, or previously blocked surface now reachable                 |

A drop in findings is only good news if the recon count held steady. If both dropped, coverage fell rather than security improving.

### Duration

Every completed scan records its duration, visible in the target's scan list. After a few scans of the same target you will have a dependable expectation for it. A scan running well beyond its usual duration is worth checking, most often for a firewall throttling traffic.

## Actions on scans

Two actions are available on a scan: cancel it, and download its report. Both are available from the target's scan list and from the scan detail view.

Canceling stops a running scan.

Findings already confirmed are preserved. A canceled scan is not a discarded scan. Every finding agents had validated before cancellation remains available on the target, with its full detail and evidence intact, because each finding was already proven at the moment it was reported.

What you lose is coverage. The remaining attack surface is not tested, and if the scan had not reached the reporting stage, there is no downloadable report for it.

Situations to consider when to stop a scan:

| Situation                                         | Why cancel                                                                                                     |
| ------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Recon found far less than expected                | The target configuration is wrong. Fix it and rescan rather than letting hours run against unreachable surface |
| The scan is affecting the environment             | Someone needs the environment stable, or scan traffic is causing problems                                      |
| You started it against the wrong target           | Stop it before it sends attack traffic where it should not go                                                  |
| A critical finding needs immediate attention      | Cancel, act on what you have, and rescan after the fix                                                         |
| The scan is running far beyond its usual duration | Investigate before letting it continue                                                                         |

You can cancel a running scan directly from the targets list without opening the target, which is convenient when you are working across several targets. See[ Manage targets](../targets/manage-targets.md).

Canceled scans remain in the target's scan history with a canceled status. This ensures the history accurately records all executed scans. Treat findings from canceled scans as valid, but consider their coverage partial. A canceled scan does not indicate that untested areas are secure.

### Download the report

The report is generated during the reporting stage, at the end of a scan. Once generated, download it from the scan as a PDF.

The report contains more detail than the UI list view. It is the artifact to hand to an auditor, attach to a compliance review, or send to a team that does not use Evo.

Historical reports remain available. Reports stay downloadable on past scans, so you can retrieve the report for an earlier assessment without rescanning.

The scan detail view shows whether the report has been generated. If a scan has findings but no report, either reporting has not finished or the scan was canceled before reaching that stage.

See Reports for what the report contains.

{% hint style="info" %}
Download the reports you need before deleting a target. Deleting a target removes its scan history, including the ability to regenerate historical reports. See Manage targets.
{% endhint %}

### Scan statuses

| Status        | Meaning                                                                                             |
| ------------- | --------------------------------------------------------------------------------------------------- |
| **Queued**    | The scan is accepted and waiting to start                                                           |
| **Running**   | The scan is executing. Progress is visible per stage                                                |
| **Completed** | The scan finished all three stages. The report is available                                         |
| **Canceled**  | The scan was stopped before completion. Confirmed findings are preserved; coverage is partial       |
| **Failed**    | The scan could not complete. Check the target configuration, starting with authentication and scope |

