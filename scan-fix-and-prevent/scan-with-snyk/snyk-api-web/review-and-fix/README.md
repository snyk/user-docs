---
description: How to review and fix findings from Snyk API and Web scans
---

# Review and fix

After scanning your targets, use Snyk API & Web to review findings, assess their severity, and assign them to team members for remediation. The review and fix workflow helps you prioritize and manage vulnerabilities efficiently.

## Key features

Snyk provides several capabilities to help you manage findings:

* **Severity classification**: Understand the risk level of each finding based on exploitability and impact.
* **Self-review**: Review and approve or reject low-confidence findings without waiting for manual verification.
* **Findings assignment**: Delegate vulnerabilities to specific team members for resolution.
* **SAST/DAST correlation**: Link dynamic scan findings directly to vulnerable code locations in your source code.
* **Login verification**: Confirm successful authentication during scans to ensure comprehensive coverage.

## Review workflows

### Review scan results

After running a scan, you can view detailed results showing scan progress, discovered vulnerabilities, and generated reports. The scan results page provides insights into the scanner components (fingerprinter, crawler, and scanner), HTTP response codes, and the overall risk assessment of your target.

Visit [Interpret target scan results](interpret-target-scan-results.md) to learn how to understand and analyze scan results. For information about HTTP response codes, visit [HTTP status codes in target scans](http-status-codes-in-target-scans.md).

### Review findings severity

Each finding includes a severity level (Critical, High, Medium, Low) to help you prioritize remediation efforts. The severity considers the likelihood of exploitation, required skills to exploit, and potential impact.

Visit [Severity levels in findings](severity-levels-in-findings.md) for detailed information about how Snyk assigns severity levels.

### Review pending findings

Snyk lets you review findings in a **Pending Review** state without waiting for manual verification by the Snyk team. This feature gives you immediate access to low-confidence findings so you can speed up your security reviews.

Visit [Review pending findings](https://github.com/snyk/user-docs/blob/main/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/review-and-fix/review-pending-findings.md) for step-by-step instructions on enabling and using the self-review feature.

### Assign vulnerabilities to team members

You can assign findings to team members for tracking and remediation. Assign findings individually or in bulk from the **Target** page, **Scan Results** page, or **Finding Details** page.

Visit [Assign vulnerabilities to a team member](assign-vulnerabilities-to-team-member.md) for instructions on assigning findings.

### Review scan login attempts

When authentication is configured on a target, you can verify that the scanner successfully logged in by watching a video recording of the login attempt. This helps you confirm authentication worked as expected and alerts you to update target settings if needed.

Visit [Review scan login attempts](review-scan-login-attempts.md) for details on accessing and using login attempt recordings.

## SAST/DAST integration

The SAST/DAST integration connects dynamic scan findings from Snyk with static analysis results from Snyk Code. This correlation links DAST findings directly to the vulnerable location in your source code, helping developers fix vulnerabilities faster.

When a DAST finding correlates with SAST results, you see:

* A SAST label on the finding.
* Direct links to the vulnerability in Snyk Code.
* Links to the affected repository.
* Code snippets showing the vulnerable code.

### Prerequisites

* Active accounts in both Snyk API & Web and the Snyk platform.
* A target application scanned by both Snyk Code (SAST) and Snyk API & Web (DAST).

### Set up the integration

To connect your dynamic and static scan results, do the following.

#### Connect your Snyk accounts

1. In Snyk, navigate to **Settings** > **Integrations**.
2. Locate the **Snyk** module.
3. Follow the link to **Snyk group** to start the authentication and authorization process.

#### Map a target to your Snyk projects

1. Navigate to the **Targets** page and identify the target you want to integrate.
2. Navigate to that target **Settings** and click the **Integrations** tab.
3. In the **Snyk** module, click **Select projects** to open a new modal.
4. Map the current target to the corresponding code analysis projects from Snyk and click **Save**.

#### Run a DAST scan

Run a new scan on the configured target. Snyk now correlates the DAST findings from this scan with the SAST findings from your mapped Snyk projects.

#### Analyze correlated findings

After the scan completes, correlated findings display with a SAST label. To view details, do the following.

1. From the findings list for your target, click a finding to open its details page.
2. Select the **SAST Findings** tab to see the connection between your DAST and SAST results.

A DAST finding provides proof that a vulnerability is exploitable. The integration links that finding directly to the specific line in your source code, showing a link to the vulnerability in Snyk Code, a link to the repository, and a code snippet triggering the vulnerability.

#### Provide feedback

You can provide feedback on the correlation accuracy to help improve the correlation engine.

* In the **SAST Findings** tab, report whether the correlation was a match or a mismatch.
* Provide qualitative feedback to help fine-tune the correlation process.
