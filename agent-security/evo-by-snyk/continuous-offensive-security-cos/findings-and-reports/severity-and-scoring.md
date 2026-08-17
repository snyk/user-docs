# Severity and Scoring

Every finding carries a severity, a CVSS score, and a CVSS vector. Together these let you order work across findings and across targets.

## Severity levels

Findings are classified by severity, and severity is what the counts throughout the product are grouped by: the open findings on each target, the summary at the top of a target, and the aggregate posture across your portfolio.

| Severity     | How to treat it                                          |
| ------------ | -------------------------------------------------------- |
| **Critical** | Fix now. Contributes to targets at risk                  |
| **High**     | Fix in the current cycle. Contributes to targets at risk |
| **Medium**   | Schedule into planned work                               |
| **Low**      | Address opportunistically                                |

A target with at least one critical or high finding open is counted as at risk on your target risk overview view.

## CVSS score and vector

| Element         | What it is                                                                                                                                                                           |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **CVSS score**  | The numeric severity of the finding under the Common Vulnerability Scoring System                                                                                                    |
| **CVSS vector** | The string showing how the score was composed, from factors such as attack vector, required privileges, user interaction, and impact on confidentiality, integrity, and availability |

The score gives you a single comparable number. The vector tells you why, which is what lets you disagree with it intelligently.

## Severity is not confidence

This is the most important difference from a conventional scanner, and it is worth being explicit about.

In a scanner, severity is entangled with confidence. A high severity finding gets attention partly because it is more likely to be genuine, and triage begins by reproducing it. Most of the queue is noise.

Here, every finding was exploited before it was reported. Severity therefore carries no information about whether a finding is real:

* A medium severity finding is exactly as real as a critical one.
* You do not need to reproduce a finding to establish that it is genuine. The proof of work is the record of it being exploited.
* Severity orders work. It does not filter noise, because there is none to filter.

The practical consequence is that a low severity finding cannot be dismissed on the grounds that it is probably a false positive. It is a genuine vulnerability that is less consequential.

## Prioritizing well

Severity orders findings by their general seriousness. Your own context changes that order, and three inputs are worth weighing alongside the score.

### Read the proof of work

The proof of work shows what the attack actually obtained. Two findings of the same class and similar score can differ substantially in what they yielded, and that difference should usually win. A finding whose proof of work shows another tenant's data is more urgent than a same-score finding whose proof of work shows an internal version string. See Understand finding details.

### Read the CVSS vector, not just the score

The vector exposes the assumptions behind the score. A high score that assumes network reachability may matter less on a service only reachable internally. A moderate score that requires no privileges and no user interaction may matter more than its number suggests, because it is trivially automatable.

### Weigh the asset

The context you gave the target describes which parts of the application handle sensitive data. A medium finding on a payment path can outrank a high finding on a marketing page. See Add a target.

{% hint style="info" %}
A reasonable default: order by severity, then reorder the top of that list by proof of work and asset sensitivity. Do not reorder the whole list; the returns are not worth the effort below high.
{% endhint %}

## Age as a prioritization input

Each finding carries a first found date. A critical finding first seen months ago and still confirmed in the latest scan has survived several releases. That is a signal about your remediation process, and it is often a stronger argument for action than the score alone.

Sorting the findings list by first found, oldest first, surfaces these. See Interpret scan results.

