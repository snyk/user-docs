# Understanding the Snyk Open Source CLI results

After you run the `snyk test` command in the CLI, the results of the Snyk Open Source test are presented in the terminal in one block that includes three sections, summary of the test findings, list of vulnerability issues detected, and descriptive information about the Snyk Project tested.

<figure><img src="../../../.gitbook/assets/image (107) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1) (2).png" alt="Snyk Open Source CLI test results."><figcaption><p>Snyk Open Source CLI test results</p></figcaption></figure>

## Summary of test findings

The summary of the test findings at the top shows the following:

<figure><img src="../../../.gitbook/assets/image (21).png" alt="Snyk Open Source CLI dependencies tested, issues and vulnerabilities found."><figcaption><p>Snyk Open Source CLI dependencies tested, issues and vulnerabilities found</p></figcaption></figure>

* The number of direct and transitive dependencies scanned
* Total number of issues found across one or more paths
* The number of paths through which vulnerable dependencies are introduced

## **List of vulnerability issues detected by Snyk Open Source**

The list of issues discovered in the Snyk Open Source test is split into three sections.

### **Issues to fix by upgrading**

These are issues which can be fixed by upgrading a direct dependencies version. They contain the following information:

* Which dependency to upgrade in order to resolve the issue
* The issue type
* The [severity rating](../../../manage-issues/issue-management/severity-levels.md) for the issue
* A link to the related issue in the [Snyk Vulnerability Database](https://security.snyk.io/)
* Which dependency this vulnerability is introduced through and its path

### **Issues with no direct upgrade or patch**

These are issues that cannot be resolved by upgrading a direct dependency. They contain the same information as “Issues to fix by upgrading” with the addition of versions which the vulnerability could be fixed in if the dependency were upgradable.

### License Issues

These are license issues which are determined by your organization’s [license policy](../../../manage-issues/policies/license-policies/). They contain the following information:

* Type of license
* License severity as determined by your organization's license policy
* Link to the [Snyk Vulnerability Database](https://security.snyk.io/) which provides more information about the license
* Dependency path through which this license is introduced

## Descriptive information about the Snyk Project scanned

The general information section about the test results includes the following details:\\

* **Organization:** The Snyk ID or internal name of the organization under which the test ran. For more information, see the [CLI test command help](../../../snyk-cli/commands/test.md).
* **Package manager:** The package manager associated with this open source scan
* **Target File:** The target file which was scanned for open source vulnerabilities
* **Project Name:** The name of the directory in which this project is located
* **Open Source:** Information about whether or not this scan was performed on an open source project. For more information, see [Test public repositories before use](../../../snyk-cli/test-for-vulnerabilities/test-public-repositories-before-use.md).
* **Project Path:** The path through which the target file is introduced
* **Local Snyk Policy: I**nformation as to whether or not this scan was performed on an open source project. For more information, see [The .snyk file](../../../snyk-cli/test-for-vulnerabilities/the-.snyk-file.md).
* **Licenses:** Information about whether or not this project was scanned for license issues
