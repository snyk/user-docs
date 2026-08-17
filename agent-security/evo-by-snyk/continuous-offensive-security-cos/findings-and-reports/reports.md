# Reports

Every completed scan produces a downloadable report. The report contains more detail than the UI list view, which makes it the artifact to use whenever the audience is outside Evo.

## Generating a report

Report generation is the third and final stage of a scan. You do not trigger it separately: when a scan completes its reporting stage, the report is ready.

The scan detail view shows whether the report has been generated. Download it from the scan, or from the actions on the target's scan list.

Reports download as PDF.

If a scan has findings but no report, either the reporting stage has not finished, or the scan was canceled before reaching it. Findings from a canceled scan are still available in the UI even where no report was produced. See [Monitor and manage scans](../scans/monitor-and-manage-scans.md).

## What the report contains

The report is paginated and print-ready, and contains the full record of the assessment:

* What was tested, and the coverage achieved.
* Every finding, with its severity, CVSS score and vector, and location.
* The steps agents took to reach each finding.
* The proof of work: the result of running each exploit.
* Remediation guidance for each finding.

Because validation happened during the scan, the report is self-supporting. A reader who was not involved in the assessment can see not just what was claimed but the evidence for it.

## Historical reports stay available

Reports remain downloadable on past scans. You can retrieve the report for an earlier assessment at any time, without rescanning.

This is what makes the scan history usable as an audit trail. When an auditor asks what you knew in a given quarter, the report from that quarter's scan is the answer, unchanged.

{% hint style="warning" %}
Deleting a target removes its scan history, and with it the ability to retrieve historical reports. Download anything you need to retain before deleting a target. See Manage targets.
{% endhint %}

## Using reports for audit and compliance

The report is designed to be handed over rather than edited. Three properties make it useful in a compliance context:

| Property                     | Why it matters to an auditor                                                                                              |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| **Evidence per finding**     | Each finding carries the result of the exploit, so the report demonstrates testing was performed rather than asserting it |
| **Coverage recorded**        | The report states what was tested, which addresses the scope question auditors ask first                                  |
| **Point in time, immutable** | The report for a past scan reflects that scan. It does not change as the application changes                              |

A regular scanning cadence, with the report retained from each scan, produces a continuous record of security testing. That is usually more valuable to an auditor than a single thorough assessment, because it demonstrates an ongoing process.

## Using reports with developers and stakeholders

| Audience                     | How to use the report                                                                                                                               |
| ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Developers**               | The steps and remediation guidance give a reproduction path and a fix. Where a repository is linked to the target, guidance names the file and line |
| **Engineering leadership**   | Severity distribution and the proof of work make the case for prioritization concretely                                                             |
| **Teams without Evo access** | The report is self-contained, so it works as a handover artifact                                                                                    |
| **Auditors and assessors**   | Evidence, coverage, and date are all present                                                                                                        |

## Reports and the UI

Use the UI for working through findings: it filters, sorts, and lets you triage while a scan is still running. Use the report when you need a fixed, complete, shareable record.

They are not alternatives. The normal pattern is to triage in the UI during and immediately after a scan, then download the report to retain as evidence of that assessment.

