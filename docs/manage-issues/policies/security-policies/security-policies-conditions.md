# Security policy conditions

A condition is a variable on which to set a rule.

After you select a condition category, you are prompted to select **Includes** or **Does not include** and the desired condition, for example, Mature, CWE-1234.

You can stack multiple conditions in the same rule using the **AND** function.

<table data-header-hidden><thead><tr><th width="210"></th><th></th></tr></thead><tbody><tr><td><strong>Condition category</strong></td><td><strong>Condition definitions and variables</strong></td></tr><tr><td>Exploit Maturity</td><td>Matches all issues that contain a specified exploit maturity value. When multiple values are selected, the condition will apply to all issues containing any of the selected values. Values include <code>Mature</code>, <code>Proof of Concept</code>, <code>No known exploit</code>, and <code>No data available</code>.</td></tr><tr><td>CWE</td><td>Matches all issues that contain a specified CWE. Supports multiple CWEs.</td></tr><tr><td>CVE</td><td>Matches all issues that contain a specified CVE. Supports multiple CVEs.</td></tr><tr><td>Snyk ID</td><td>Matches all issues that contain a specified Snyk ID. Supports multiple Snyk IDs. Not every issue has a CVE, so this is a good way to specify those.</td></tr></tbody></table>
