# Understand finding details

Opening a finding gives you everything needed to triage it, reproduce it, and hand it to a developer, without switching tools.

## Finding metadata

At the top of the finding:

| Field           | What it tells you                                                                                                        |
| --------------- | ------------------------------------------------------------------------------------------------------------------------ |
| **First found** | When any scan first confirmed this finding. Shows how long the vulnerability has been present                            |
| **Last found**  | When it was most recently confirmed. If this stops advancing while scans continue, the finding is no longer reproducible |
| **CVSS score**  | Numeric severity. Visit Severity and scoring                                                                             |
| **CVSS vector** | The string showing how the score was derived, so you can re-evaluate it against your own environment                     |
| **Location**    | Where the finding was exploited: the URL or endpoint                                                                     |

The gap between first found and last found is worth reading. A finding first seen months ago and confirmed again in the latest scan has survived multiple releases, which usually says something about process rather than difficulty.

## The four tabs

Each finding has four tabs. They are ordered to match how you work through a finding: understand it, see how it was found, see what it yielded, then fix it.

You can expand the finding into a larger form factor when you need more room, which helps when reading through a long sequence of steps or a substantial exploit result.

### Overview

A summary of the vulnerability and an explanation of its impact.

Read this first. It answers what the vulnerability is and what it would mean if an attacker exploited it in your application. The impact explanation is what you use to argue priority with an engineering team, since it describes consequence rather than vulnerability class.

### Steps

The sequence of actions the agents took in order to find this vulnerability.

This is the reasoning trace: what was attempted at each stage and what was concluded from the result. It serves three purposes:

| Purpose                         | How to use it                                                                                                                 |
| ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| **Reproduction path**           | Follow the steps to reproduce the finding yourself, or hand them to a developer as reproduction instructions                  |
| **Understanding the real flaw** | The path agents took often reveals that the underlying problem is broader than the endpoint where it surfaced                 |
| **Scoping the fix**             | If agents reached the vulnerability by chaining several weaknesses, fixing only the final step leaves the chain mostly intact |

Reading the steps is what stops a fix from being too narrow. An endpoint that leaked data because an authorization check was missing at a shared layer will be exploitable through other endpoints too, and the steps are where that becomes visible.

### Proof of work

The result of running the exploit against your application.

This is the evidence that the finding is real. Agents did not infer the vulnerability from a response pattern; they exploited it, and this is what came back.

Proof of work is also the best available measure of severity in practice. The vulnerability class tells you what category of problem you have; the proof of work tells you what an attacker actually obtained. Two findings of the same class can differ enormously here, and that difference should drive which you fix first.

{% hint style="info" %}
When you need to convince someone that a finding warrants immediate attention, the proof of work is the artifact to show them. It is concrete in a way that a CVSS score is not.
{% endhint %}

### Remediation guidance

How to fix the underlying issue.

The specificity depends on whether the target has a source repository linked:

| Target configuration                | What guidance you get                                                              |
| ----------------------------------- | ---------------------------------------------------------------------------------- |
| **No repository linked** (blackbox) | An explanation of the vulnerability class and how to remediate it, with references |
| **Repository linked** (greybox)     | The specific file, the line number, and a suggested code change                    |

If you are receiving generic guidance and want file-level guidance, link a repository to the target. Subsequent scans then run as greybox automatically. See What is a Target.

## Working through a finding

1. **Overview** to understand what it is and what it would cost you.
2. **Proof of work** to see what the attack actually yielded, which is usually the real basis for prioritization.
3. **Steps** to understand how it was reached, and whether the underlying flaw is broader than where it surfaced.
4. **Remediation guidance** to make the change.
5. **Rescan** to confirm the fix holds.

Step 5 matters. A fix that closes the specific path the exploit used, without addressing the flaw beneath it, will be found again through a different path. Rescanning is the only reliable confirmation.

