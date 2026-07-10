---
description: The states a Snyk API and Web finding can have
---

# Finding states

During a target scan, the scanner identifies vulnerabilities within the target's URLs. When the scanner discovers a vulnerability, Snyk creates a finding. The state of a finding can change either automatically (by the scanner, as a result of a target scan or re-test) or manually (by user actions).

## Available finding states

A finding can have the following states:

* **Not fixed** - The scanner found a vulnerability that is waiting to be fixed. The user does not control this state. As long as a target scan or re-test finds the vulnerability, the finding stays in the **Not fixed** state.
* **Invalid** - The user marked the vulnerability as invalid. This state results from a user action. Use it to report a false positive.
* **Accepted risk** - The user marked the vulnerability as accepted risk. This state results from a user action. Use it to identify vulnerabilities that the user does not consider in need of being fixed.
* **Fixed** - The scanner could not find a previously existing vulnerability while running a subsequent target scan using the same profile or a broader one, so it marked the vulnerability as fixed. The user does not control this state.
* **Re-testing** - The scanner is re-testing a previously existing vulnerability. This state results from a user action. It can lead to either a **Fixed** vulnerability (if the scanner cannot replicate the vulnerability during the re-test) or a **Not fixed** vulnerability (if the scanner finds it again during the re-test).

## User-controlled vs. scanner-controlled states

In summary, the user controls the **Invalid**, **Accepted risk**, and **Re-testing** states, while the scanner sets the **Fixed** and **Not fixed** states.
