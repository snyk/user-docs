# Finding states

During a target scan, the scanner identifies vulnerabilities within the target's URLs. When a vulnerability is discovered, a finding is created. The finding's state can change either automatically (by the scanner, as a result of a target scan or re-test) or manually (by user actions).

## Available finding states

A finding can have the following states:

* **Not fixed** - A vulnerability was found and is waiting to be fixed. This state is not controlled by the user. As long as a target scan or re-test finds the vulnerability, the finding will be set as Not fixed.
* **Invalid** - The vulnerability was marked as invalid. This state is a result of a user action. You can use it to report a False Positive.
* **Accepted risk** - The vulnerability was marked as accepted risk. This state is a result of a user action. This can be used to identify vulnerabilities that the user doesn't consider in need of being fixed.
* **Fixed** - A previously existing vulnerability couldn't be found while running a subsequent target scan using the same profile (or a broader one), thus it has been marked as fixed. This state is not user controlled.
* **Re-testing** - A previously existing vulnerability is being re-tested. This state is a result of a user action and can lead to either a Fixed vulnerability (if the scanner isn't able to replicate the vulnerability during the re-test) or Not fixed vulnerability (if the scanner is able to find it again during the re-test).

## User-controlled vs. scanner-controlled states

In summary, Invalid, Accepted risk, and Re-testing are states controlled by the user, while Fixed and Not fixed are states set by the scanner.
