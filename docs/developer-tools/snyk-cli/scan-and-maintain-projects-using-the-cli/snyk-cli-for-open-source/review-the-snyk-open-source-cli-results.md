# Review the Snyk Open Source CLI results

After you run the `snyk test` command in the CLI, the Snyk Open Source test results are displayed. The report of results includes a summary of the test findings, a list of vulnerability issues detected, and descriptive information about the Snyk Project tested.

<figure><img src="../../../../.gitbook/assets/image (107) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1) (2).png" alt="Snyk Open Source CLI test results."><figcaption><p>Snyk Open Source CLI test results</p></figcaption></figure>

## Summary of test findings

The summary of the test findings at the beginning of the report of results shows the following:

<figure><img src="../../../../.gitbook/assets/image (28).png" alt="Snyk Open Source CLI dependencies tested, issues and vulnerabilities found."><figcaption><p>Snyk Open Source CLI dependencies tested, issues and vulnerabilities found</p></figcaption></figure>

* The number of direct and transitive dependencies scanned
* Total number of issues found across one or more paths
* The number of paths through which vulnerable dependencies are introduced

## **List of vulnerability issues detected by Snyk Open Source**

The list of issues discovered in the Snyk Open Source test includes the following sections.

### **Issues to fix by upgrading**

These are issues that can be fixed by upgrading a direct dependencies version. They contain the following information:

* Which dependency to upgrade to resolve the issue
* The issue type
* The [severity rating](../../../../manage-risk/prioritize-issues-for-fixing/severity-levels.md) for the issue
* A link to the related issue in the [Snyk Vulnerability Database](https://security.snyk.io/)
* The dependency through which this vulnerability is introduced and its path

### **Issues with no direct upgrade or patch**

These are issues that cannot be resolved by upgrading a direct dependency. They contain the same information as the issues to fix by upgrading, with versions in which the vulnerability could be fixed if the dependency were upgradeable.

### License Issues

License issues are determined by the [license policy](../../../../manage-risk/policies/license-policies/) of your Snyk Organization. License issues  contain the following information:

* Type of license
* License severity as determined by the license policy of your Snyk Organization
* Link to the [Snyk Vulnerability Database](https://security.snyk.io/), which provides more information about the license
* Dependency path through which this license is introduced

## Descriptive information about the Snyk Project scanned

The descriptive information about the test results includes the following details:

* **Organization:** The Snyk ID or internal name of the Organization under which the test ran. For more information, see [`--org=ORG_ID>` option](../../commands/test.md#org-less-than-org_id-greater-than) in the `snyk test` help.
* **Package manager:** The package manager associated with this Open Source scan
* **Target File:** The target file which was scanned for Open Source vulnerabilities
* **Project Name:** The name of the directory in which this Project is located
* **Open Source:** Information about whether or not this scan was performed on an Open Source Project.
* **Project Path:** The path through which the target file is introduced
* **Local Snyk Policy:** Information about whether this scan was performed on an Open Source Project. For more information, see [The .snyk file](../../../../manage-risk/policies/the-.snyk-file.md).
* **Licenses:** Information about whether this Project was scanned for license issues
