# Issue columns dictionary

Snyk reporting includes tens of filters and columns, allowing users to develop refined views of the data and obtain the required insights with ease. Thus reaching accurate conclusions requires one to understand the columns used and the meaning of the filters. This dictionary describes the meaning behind the issue columns in Snyk Issues Detail report.

## Issue characteristics <a href="#issue-characteristics" id="issue-characteristics"></a>

Describes the main attributes of the issue.

* **AUTO FIXABLE** - Indicates whether the issue can be automatically fixed by Snyk.
* **COMPUTED FIXABILITY** - Indicates whether the issue can be fixed by Snyk. For details, see [Computed Fixability filters](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/vulnerability-fix-types.md#computed-fixability-filters).
  * **Fixable:** There is a fix for all the identified issues, meaning that all detailed paths have remediation.
  * **Partially fixable:** The issue has upgradable paths, but not all detailed paths have remediation.
  * **No supported fix**: The issue has no upgradable paths.
* **INTRODUCTION CATEGORY** - Categorizes how an issue was introduced:
  * **Baseline Issue** - issues identified right after the Project began to be monitored.
  * **Preventable Issue** - issues for which Snyk published the related problem at least seven days before detection.
  * **Non Preventable Issue** - issues that were created due to external factors such as a new vulnerability being published.
  * **Other New Issue** - issues for which Snyk cannot classify their preventability. For details, see [Delineation of how risk is introduced](../analytics/enterprise-analytics.md#delineation-of-how-risk-is-introduced).
* **ISSUE** - A combination of:
  * **Problem Title**: Snyk vulnerability name.
  * **Issue Type:** indicates whether the issue is related to a vulnerability, license, or configuration.
* **ISSUE STATUS** - Indicates whether the issue is open, resolved, or ignored.
* **PRODUCT NAME** - Snyk product name.
* **SEVERITY** - Indicates the issue severity according to the analysis by a specific Snyk product.
* **SCORE** - A score based on an analysis model. Priority score is released in General Availability, while Risk Score is in Early Access. For details, see [Priority Score vs Risk Score](../prioritize-issues-for-fixing/priority-score-vs-risk-score.md).
* **REACHABILITY** - The reachability of the issue indicates whether the issue is related to functions that are being called by the application and thus has a greater risk of exploitability.\
  Allowed values:
  * **Reachable** - A direct or indirect path was found from your application to the vulnerable code.
  * **No path found** - No path found from your application to the vulnerable code
  * **Not applicable** - Reachability is not a relevant attribute for the specific issue

## Issue vulnerability details <a href="#issue-vulnerability-details" id="issue-vulnerability-details"></a>

The vulnerability details refer to various issue attributes that are being defined by Snyk, Mitre, NVD, or any other trusted security organization.

* **CVE** - Mitre CVE ID
* **CWE** - Mitre CWE ID
* **EXPLOIT MATURITY** - Represents the existence and maturity of public exploits validated by Snyk. The allowed values follow. For details, see [View exploits in Projects](../prioritize-issues-for-fixing/view-exploits.md#view-exploits-in-projects).
  * **Mature:** Snyk has a published code exploit for this vulnerability.
  * **Proof of concept:** Snyk has a proof-of-concept or detailed explanation of how to exploit this vulnerability. Proof of concept vulnerability patches cannot be disabled and will appear in fix PRs where they are found
  * **No known exploit:** Snyk did not find a proof-of-concept or a published exploit for this vulnerability.
  * **No data**: The issue is not a vulnerability, but a license issue or a vulnerability advisory.
* **FIXED IN AVAILABLE** - Indicates if a new package version that includes a fix for the vulnerability exists
* **FIXED IN VERSION** - Indicates the package version with a fix for the vulnerability.
* **NVD SCORE** - The vuln score as calculated by NVD.
* **NVD SEVERITY** - The vuln severity as rated by NVD.
* **PACKAGE NAME AND VERSION** - The vuln associated package name and version.
* **PROBLEM ID** - Snyk Vuln DB ID that uniquely identifies the vulnerability.
* **PROBLEM TITLE** - The vulnerability name as described by Snyk.
* **SEMVER VULNERABLE RANGE** - The vulnerable range of package versions (based on semantic versioning).
* **VULNERABILITY PUBLICATION DATE** - Timestamp indicating when the vulnerability was published by Snyk.

## Issue context columns <a href="#issue-context-columns" id="issue-context-columns"></a>

The context columns help to understand the wider meaning of the issue’s impact and risk over its associated project, target, organization, etc.

* **ORG NAME** - The Snyk Organization name.
* **PROJECT COLLECTION** - Project collections are static collections of Projects.
* **PROJECT CRITICALITY** - Business criticality of the Project. For details, see  [Project attributes](../../snyk-admin/snyk-projects/project-attributes.md).
* **PROJECT ENVIRONMENT** - The environment of the Project. For details, see  [Project attributes](../../snyk-admin/snyk-projects/project-attributes.md).
* **PROJECT LIFECYCLE** - Project’s lifecycle. For details, see  [Project attributes](../../snyk-admin/snyk-projects/project-attributes.md).
* **PROJECT NAME** - The Project name
* **PROJECT ORIGIN** - The Project’s source of integration; can be the name of the originating SCM, Container registry, and so on.
* **PROJECT OWNER** - A user that is defined as the Project’s owner.
* **PROJECT TAGS** - Tags that are associated with the Project. For details, see [Project tags](../../snyk-admin/introduction-to-snyk-projects/project-tags.md).
* **PROJECT TARGET** - The Target name.
* **PROJECT TARGET REFERENCE** - Specify a reference that differentiates this Project, for example, a branch name or version. For details, see [Group Projects by branch or version for monitoring](../../snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring.md).
* **PROJECT TYPE** - The package manager of the Project

## Issue date columns <a href="#issue-date-columns" id="issue-date-columns"></a>

The date columns indicate important events that are used for various calculations and trend analysis

* **INTRODUCED** - Timestamp of the first scan that identified the issue.
* **LAST IGNORED** - Timestamp indicating when the issue was last marked as ignored,
* **LAST INTRODUCED** - Timestamp of the last scan that identified this issue,
* **LAST RESOLVED** - The last timestamp at which the issue was not found in a test.
