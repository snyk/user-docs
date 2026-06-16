# Use the Snyk SAST/DAST integration

This guide explains how to set up and use the SAST/DAST integration to correlate findings from Snyk API & Web (DAST) with your static analysis results in Snyk (SAST).

## Overview

By connecting your dynamic and static scan results, you can streamline triage and remediation. This integration links DAST findings directly to the vulnerable location in your source code, helping your developers fix issues faster.

## Prerequisites

* You must have active accounts in both Snyk API & Web and the Snyk platform.
* You need to have a target application that is being scanned by both Snyk for SAST (Snyk Code) and Snyk API & Web for DAST.

## Step 1: Connect your Snyk accounts

First, you need to establish a connection between your Snyk API & Web account and your main Snyk account.

1. In Snyk API & Web, navigate to **Settings > Integrations**.
1. Locate the **Snyk** module.
1. Follow the link to **Snyk group**. This starts the authentication and authorization process to connect your two accounts.

<figure><img src="../../../.gitbook/assets/how-to-use-the-snyk-sastdast-integration_image1.png" alt="Snyk integration module in Settings"></figure>

## Step 2: Map a target to your Snyk projects

Next, you need to tell Snyk API & Web which Snyk Code projects (code repository) correspond to your DAST target.

1. Navigate to the **Targets** page and identify the target you want to integrate.
1. Go to that target **Settings** and click the **Integrations** tab.
1. In the **Snyk** module, click **Select projects** to open a new modal.
1. Map the current Snyk API & Web target to the corresponding code analysis project(s) from Snyk and click **Save**.

<figure><img src="../../../.gitbook/assets/how-to-use-the-snyk-sastdast-integration_image2.png" alt="Mapping Snyk API & Web target to Snyk Code projects"></figure>

## Step 3: Run a DAST scan

Run a new scan on the target you configured in Step 2. Snyk API & Web now correlates the DAST findings from this scan with the SAST findings from your mapped Snyk projects.

## Step 4: Analyze correlated findings

After the scan is complete, you can view the correlated results. Any correlated finding will have a SAST label associated with it.

<figure><img src="../../../.gitbook/assets/how-to-use-the-snyk-sastdast-integration_image3.png" alt="Findings list showing SAST label"></figure>

1. From the list of findings for your target, click a finding to open its details page.
1. Select the **SAST Findings** tab.

## Verify the outcome

The **SAST Findings** tab provides the connection between your DAST and SAST results.

A DAST finding provides proof that a vulnerability is exploitable. This integration links that finding directly to the specific line in your source code, so you can fix it immediately. The correlation provides you with a link to the vulnerability within Snyk Code, a link to the repository, and a snippet of code that is triggering a given vulnerability.

<figure><img src="../../../.gitbook/assets/how-to-use-the-snyk-sastdast-integration_image4.png" alt="SAST Findings tab showing correlated results"></figure>

## Manage the feature and provide feedback

You can provide feedback directly on the finding details page.

* In the **SAST Findings** tab, you can report whether you believe the correlation was a match or a mismatch. This helps fine-tune the correlation process and provide better results.
* You can also provide additional qualitative feedback to help improve the accuracy of the correlation engine.
