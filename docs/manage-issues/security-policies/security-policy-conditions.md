# Security policy conditions

A condition is a variable to set a rule on.&#x20;

After you select a condition category, you are prompted to select **Includes** or **Does not include** and the desired condition (for example Mature, CWE-1234).

You can stack multiple conditions in the same rule using the **AND** function.

| **Condition Category** | **Condition Variables/Definitions**                                                                                                                                                                                                                  |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Exploit Maturity       | Matches all issues that contain a specified Exploit maturity value. When multiple values are selected, it will apply to all issues containing any of these values. Values include Mature, Proof of Concept, No known exploit, and No data available. |
| CWE                    | Matches all issues that contain a specified CWE. Supports multiple CWEs.                                                                                                                                                                             |
| CVE                    | Matches all issues that contain a specified CVE. Supports multiple CVEs.                                                                                                                                                                             |
| Snyk ID                | Matches all issues that contain a specified Snyk ID. Supports multiple Snyk IDs. Not every issue has a CVE, so this is a good way of being able to specify those.                                                                                    |
