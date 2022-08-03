# Understanding the Snyk Open Source CLI results



After you run the `snyk test` command in the CLI, the results of the Snyk Open Source test are presented in the terminal in one block that includes three sections, summary of the test findings, list of vulnerability issues detected, and descriptive information about the Snyk Project tested.

![Snyk Open Source CLI test results](../../.gitbook/assets/image.png)

## Summary of test findings

The summary of the test findings at the top shows the following:

![Snyk Open Source CLI dependencies tested, issues and vulnerabilities found](<../../.gitbook/assets/image (3).png>)

* The number of direct and transitive dependencies scanned
* Total number of issues found across one or more paths
* The number of paths through which vulnerable dependencies are introduced

## **List of vulnerability issues detected by Snyk Open Source**

The list of issues discovered in the Snyk Open Source test is split into three sections.

### **Issues to fix by upgrading**

These are issues which can be fixed by upgrading a direct dependencies version. They contain the following information:

![Snyk Open Source CLI  test results, issues to fix by upgrading](https://lh5.googleusercontent.com/9jt74X2HNuIYhK17OJrsPE1kktMvn0D9tJ8EATxwvtPaZMJud9imPh1rqVcB7Ya-G4vMFgStgYVFqieQpttZiR\_BFUzvfIqB8Evt3-DDqXB5wJNtjeQWrhEQpp6tRy30psgxkE3GjhdjatRFmd0g17c)

* Which dependency to upgrade in order to resolve the issue&#x20;
* The issue type&#x20;
* The [severity rating](https://docs.snyk.io/introducing-snyk/snyks-core-concepts/severity-levels) for the issue&#x20;
* A link to the related issue in the [Snyk Vulnerability Database](https://security.snyk.io/)&#x20;
* Which dependency this vulnerability is introduced through and its path

### **Issues with no direct upgrade or patch**

These are issues that cannot be resolved by upgrading a direct dependency. They contain the same information as “Issues to fix by upgrading” with the addition of versions which the vulnerability could be fixed in if the dependency were upgradable.\


![Snyk Open Source CLI test results, issues with no direct upgrade or patch](https://lh4.googleusercontent.com/-xyYz9oJXYQs-Q5mmQlMzOlUwNVFsc7ftAlYmGNIC51h1g6SgWfmIg7KntXy8V-IiZbzfaIqSrHRyd6p4S2tp82GCOrWOcJDnNG-kUA\_zFnX3IwzgmGnGnu3PKlQ0og713peu7EGZT5GsWQkLbJ5mks)

### License Issues

These are license issues which are determined by your organization’s [license policy](https://docs.snyk.io/products/snyk-open-source/license-policies). They contain the following information:\


![Snyk Open Source CLI test results, license issues](https://lh5.googleusercontent.com/4nxJXQtFicCcC48k6KAxYBVX\_FTlrbB6Ir1ZUi-IJG-5a0c84jm5cmY\_s0yq-TD6X8dOtZzIP2\_boqtX4YbPJa3rjGYOA5ne1EtfId6WNZJ34rThP2Byj2uAiMzy935ItuQWd9zWqh7ilK3uBsgqJbo)

* Type of license&#x20;
* License severity as determined by your organization's license policy&#x20;
* Link to the [Snyk Vulnerability Database](https://security.snyk.io/) which provides more information about the license
* Dependency path through which this license is introduced

## Descriptive information about the Snyk Project scanned

The general information section about the test results includes the following details:\


![Snyk Open Source CLI test results description of project scanned](https://lh4.googleusercontent.com/qqa6EAFsA7TNLYnfatSTRDrO63EVAruymozZnyrsNmZStgGAVNMeEIenl6XA0KA9jaSN4BDJD90Y61IrHxTkCf4UR5wk7-NuIADvNG7AomLTwZsTNa5Kep7WAGUa-Iq8qeRHBBJQMgfMCFuiTNqie5A)

* **Organization:** The Snyk ID or internal name of the organization under which the test ran. For more information, see the [CLI test command help](https://docs.snyk.io/snyk-cli/commands/test#org-less-than-org\_id-greater-than).
* **Package manager:** The package manager associated with this open source scan
* **Target File:** The target file which was scanned for open source vulnerabilities
* **Project Name:** The name of the directory in which this project is located
* **Open Source:** Information about whether or not this scan was performed on an open source project. For more information, see [Test public repositories before use](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/test-public-repositories-before-use).
* **Project Path:** The path through which the target file is introduced
* **Local Snyk Policy: I**nformation as to whether or not this scan was performed on an open source project. For more information, see [The .snyk file](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/the-.snyk-file).
* **Licenses:** Information about whether or not this project was scanned for license issues

