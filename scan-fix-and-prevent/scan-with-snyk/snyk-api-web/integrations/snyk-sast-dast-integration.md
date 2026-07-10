---
description: How to integrate Snyk SAST and Snyk API and Web DAST findings
---

# Snyk SAST/DAST integration

This guide explains how to set up and use the static application security testing (SAST) and dynamic application security testing (DAST) integration to correlate findings from Snyk API & Web (DAST) with your static analysis results in Snyk (SAST).

## Overview

By connecting your dynamic and static scan results, you can streamline triage and remediation. This integration links DAST findings directly to the vulnerable location in your source code, helping your developers fix vulnerabilities faster.

## Prerequisites

* You must have active accounts in both Snyk API & Web and the Snyk platform.
* You must have a target application that both Snyk for SAST (Snyk Code) and Snyk API & Web for DAST scan.

## Step 1: Connect your Snyk accounts

Establish a connection between your Snyk API & Web account and your main Snyk account.

1. In Snyk API & Web, navigate to **Settings > Integrations**.
2. Locate the **Snyk** module.
3. Click the link to **Snyk group**. This starts the authentication and authorization process to connect your two accounts.

## Step 2: Map a target to your Snyk projects

Next, tell Snyk API & Web which Snyk Code projects (code repository) correspond to your DAST target.

1. Navigate to the **Targets** page and identify the target you want to integrate.
2. Navigate to that target **Settings** and click the **Integrations** tab.
3. In the **Snyk** module, click **Select projects** to open a new modal.
4. Map the current Snyk API & Web target to the corresponding code analysis projects from Snyk and click **Save**.

## Step 3: Run a DAST scan

Run a new scan on the target you configured in Step 2. Snyk API & Web correlates the DAST findings from this scan with the SAST findings from your mapped Snyk projects.

## Step 4: Analyze correlated findings

After the scan is complete, you can view the correlated results. Any correlated finding has an associated SAST label.

1. From the list of findings for your target, click a finding to open its details page.
2. Select the **SAST Findings** tab.

## Verify the outcome

The **SAST Findings** tab provides the connection between your DAST and SAST results.

A DAST finding provides proof that a vulnerability is exploitable. This integration links that finding directly to the specific line in your source code, so you can fix it immediately. The correlation provides you with a link to the vulnerability within Snyk Code, a link to the repository, and a snippet of code that is triggering a given vulnerability.

## Manage the feature and provide feedback

You can provide feedback directly on the finding details page.

* In the **SAST Findings** tab, you can report whether you believe the correlation was a match or a mismatch. This helps fine-tune the correlation process and provide better results.
* You can also provide additional qualitative feedback to help improve the accuracy of the correlation engine.
