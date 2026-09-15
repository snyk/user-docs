---
nav_context: agnostic
description: What happens during a Snyk API and Web scan
---

# What happens during a scan

During a scan, Snyk API & Web goes through the target's URLs and interacts with every element found, filling out forms and clicking buttons, among other actions, to perform extensive tests on your target and identify as many vulnerabilities as possible.

Because of these thorough interactions, expect many requests and an influx of information into the target.

## Scan components

A scan involves three major components, each with a specific job:

* The **fingerprinter** identifies the technologies used on the target.
* The **crawler** goes through the target's URLs and interacts with every element found, clicking buttons and filling in forms, among other things.
* The **scanner** finds vulnerabilities within the target's URLs. The process involves two stages. First, it tests the target and identifies candidate vulnerabilities. Then, it validates these candidates to confirm which ones are genuinely exploitable. Validation processes one candidate at a time and continues after the testing stage completes.

## Scan states

A scan has several possible states:

* As soon as you request a scan, it gets the **Queued** state.
* After a queued scan begins, its state changes to **Started**. During a scan, the scan results page displays a progress bar for each of the three components. You can see which component is active and its progress.
* After the fingerprinter, the crawler, and the scanner complete their jobs, the scan ends with the **Completed** state.

### Additional scan states

A scan can have these additional states:

* If a user stops an ongoing scan, the scan state changes to **Canceled**.
* If the target is unreachable or a connection times out, the scan ends with the **Failed** state and a message indicating the error. Snyk uses the same state if a scan fails during execution.
* If the Snyk team must manually confirm some vulnerabilities, the scan gets the **Under Review** state. After this manual review, the scan changes to **Completed**.

## After the scan completes

After a scan finishes successfully, you can generate its scan reports and coverage reports.
