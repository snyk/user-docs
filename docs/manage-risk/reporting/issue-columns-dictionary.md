# Issue columns dictionary

Snyk reporting includes many filters and columns, allowing users to develop refined views of the data and obtain the required insights with ease. Reaching accurate conclusions requires understanding the columns used and the meaning of the filters. This dictionary explains the meaning of the issue columns in Snyk Issues Detail report.

{% hint style="info" %}
Each inner list is sorted alphabetically
{% endhint %}

## Issue characteristics <a href="#issue-characteristics" id="issue-characteristics"></a>

Describes the main attributes of the issue.

* **ASSET FINDING ID** - A unique issue ID at the repository level that is only applicable for Snyk Code issues.
* **COMMIT ID** - The unique ID that the SCM integration assigns to commits so that they can be uniquely identified. Snyk provides a Commit ID only for Snyk Code issues.
* **CODE REGION** - The line numbers and columns range where the issues were found within a file.&#x20;
* **COMPUTED FIXABILITY** - Indicates whether the issue can be fixed based on the vulnerability remediation paths. Relevant for SCA vulnerabilities. For details, see [Computed Fixability filters](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/vulnerability-fix-types.md#computed-fixability-filters).
  * **Fixable:** There is a fix for all the identified issues, meaning that all detailed paths have remediation.
  * **Partially fixable:** The issue has upgradable paths, but not all detailed paths have remediation.
  * **No supported fix**: The issue has no upgradable paths.
* **FILE PATH** - The path to the file where Snyk Code identified the specific issue.
* **HAS JIRA ISSUE(S) ASSIGNED** - Displays `true` when at least one Jira issue is assigned, otherwise displays `false`.
* **INTRODUCTION CATEGORY** - Categorizes how an issue was introduced:
  * **Baseline Issue** - Issues identified right after the Project began to be monitored.
  * **Preventable Issue** - Issues for which Snyk published the related problem at least seven days before detection.
  * **Non Preventable Issue** - Issues that were created due to external factors such as a new vulnerability being published.
  * **Other New Issue** - Issues for which Snyk cannot classify their preventability. For details, see [Delineation of how risk is introduced](../analytics/issues-analytics.md#delineation-of-how-risk-is-introduced).
* **ISSUE** - A combination of:
  * **Problem Title**: Snyk vulnerability name.
  * **Issue Type:** indicates whether the issue is related to a vulnerability, license, or configuration.
* **ISSUE STATUS** - Indicates whether the issue is open, resolved, or ignored.
* **JIRA ISSUES LIST** - A list of all the attached Jira issue keys.
* **LATEST JIRA ISSUE** - The latest attached Jira issue key with a link to the issue card on the Project page.
* **PRODUCT NAME** - Snyk product name.
* **SEVERITY** - Indicates the issue severity according to the analysis by a specific Snyk product.
* **SCORE** - A score based on an analysis model. Priority score is released in General Availability, while Risk Score is in Early Access. For details, see [Priority Score vs Risk Score](../prioritize-issues-for-fixing/priority-score-vs-risk-score.md).
* **REACHABILITY** - The reachability of the issue indicates whether the issue is related to a code element (function, module, class, etc) that are being called by the application and thus has a greater risk of exploitability. Allowed values:
  * **Reachable** - A direct or indirect path was found from your application to the vulnerable code.
  * **No path found** - No path found from your application to the vulnerable code. Note that this value might change over time.
  * **Not applicable** - Reachability is not relevant to the specific issue. This value will appear for languages that are not yet supported and for issues that are not Snyk Open Source vulnerabilities.

## Issue vulnerability details <a href="#issue-vulnerability-details" id="issue-vulnerability-details"></a>

The vulnerability details refer to various issue attributes that are being defined by Snyk, Mitre, NVD, or any other trusted security organization.

* **ATTACK VECTOR -** Represents the context by which vulnerability exploitation is possible. For more details about the attack vector and its values (Network, Adjacent, Local, Physical), see the [specification document](https://www.first.org/cvss/specification-document).&#x20;
* **CVE** - Mitre CVE ID
* **CWE** - Mitre CWE ID
* **EPSS SCORE** - The probability of exploitation in the wild in the next 30 days.
* **EPSS PERCENTILE** -  The proportion of all vulnerabilities with the same or lower EPSS score.
* **EXISTS IN DIRECT DEPENDENCY -** Indicates if the vulnerability exists in a direct dependency. If false, the vulnerability only exists in transitive dependencies.
* **EXPLOIT MATURITY** - Represents the existence and maturity of public exploits validated by Snyk. For details, see [View exploits in Projects](../prioritize-issues-for-fixing/view-exploits.md#view-exploits-in-projects). The allowed values include:
  * **Mature:** Snyk has a published code exploit for this vulnerability.
  * **Proof of concept:** Snyk has a proof-of-concept or detailed explanation of how to exploit this vulnerability. Proof of concept vulnerability patches cannot be disabled and will appear in fix PRs where they are found.
  * **No known exploit:** Snyk did not find a proof-of-concept or a published exploit for this vulnerability.
  * **No data**: The issue is not a vulnerability, but a license issue or a vulnerability advisory.
* **FIXED IN AVAILABLE** - Indicates if a new package version that includes a fix for the vulnerability exists.
* **FIXED IN VERSION** - Indicates the package version with a fix for the vulnerability.
* **NVD SCORE** - The vuln score as calculated by NVD.
* **NVD SEVERITY** - The vulnerability severity as rated by NVD.
* **PACKAGE NAME AND VERSION** - The vulnerability-associated package name and version.
* **PROBLEM ID** - Snyk Vuln DB ID that uniquely identifies the vulnerability.
* **PROBLEM TITLE** - The vulnerability name as described by Snyk.
* **SEMVER VULNERABLE RANGE** - The vulnerable range of package versions (based on semantic versioning).
* **SNYK CVSS SCORE -** Snyk's CVSS (Common Vulnerability Scoring System) score.
* **VULNERABILITY PUBLICATION DATE** - Timestamp indicating when the vulnerability was published by Snyk.

## Issue context columns <a href="#issue-context-columns" id="issue-context-columns"></a>

The context columns help you understand the impact and risk for an issue based on various contexts, including hierarchy, project and target, asset, or application context.

### Snyk hierarchy context

* **ORG NAME** - The Snyk Organization name.

### Project and target context

* **PROJECT COLLECTION** - Project collections are static collections of Projects.
* **PROJECT CRITICALITY** - The business criticality of the Project. For details, see [Project attributes](../../snyk-admin/snyk-projects/project-attributes.md).
* **PROJECT ENVIRONMENT** - The environment of the Project. For details, see [Project attributes](../../snyk-admin/snyk-projects/project-attributes.md).
* **PROJECT LIFECYCLE** - The lifecycle of a Project. For details, see [Project attributes](../../snyk-admin/snyk-projects/project-attributes.md).
* **PROJECT NAME** - The Project name.
* **PROJECT ORIGIN** - The source of integration for the Project; can be the name of the originating SCM, Container registry, and so on.
* **PROJECT OWNER** - A user who is defined as the owner of the Project.
* **PROJECT TAGS** - Tags that are associated with the Project. For details, see [Project tags](../../snyk-admin/introduction-to-snyk-projects/project-tags.md).
* **PROJECT TARGET** - The Target name.
* **PROJECT TARGET REFERENCE** - Specify a reference that differentiates this Project, for example, a branch name or version. For details, see [Group Projects by branch or version for monitoring](../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring.md).
* **PROJECT TYPE** - The package manager of the Project.

### Asset context&#x20;

{% hint style="info" %}
When filtering issues by assets context, issues of archived assets will be excluded from the results
{% endhint %}

* **ASSET CLASS**- specifies the business criticality of the asset (A, most critical - D, least critical).
* **ASSET ID** - the unique identifier of the asset within a Snyk Group.
* **ASSET NAME** - the name of the asset.
* **ASSET TAGS** - the tags that were assigned to the asset based on imported data or user input.
* **ASSET TYPE** - specifies the type of the asset: repository, container image, or package.
* **PARENT ASSET ID** - the unique identifier of the parent asset.
* **PARENT ASSET NAME** - the name of the parent asset.
* **REPOSITORY FRESHNESS** - the repository activity status based on the last commit date.
  * **Active**: Had commits in the last three months.
  * **Inactive**: The last commits were made in the last three to six months.
  * **Dormant**: No commits in the last six months.
  * **N/A**: There are no commits detected by Snyk Essentials.

### Application context

Navigate to the [Application context for SCM integrations](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/) page for more details about enriching the application context data.

{% hint style="info" %}
When filtering issues by application context, issues of archived assets will be excluded from the results
{% endhint %}

* **APPLICATION** - represents the application or service that the asset is associated with.
* **ASSET LIFECYCLE** - represents the lifecycle state of the asset, for example `production`, `experimental`, `deprecated`.
* **ASSET OWNER** - represents the code owner of the asset, which should usually be a dev team.
* **CATALOG NAME** - the name of your application context catalog.
* **CATEGORY** - the category of a repository asset, for example, service or library.

## Issue date columns <a href="#issue-date-columns" id="issue-date-columns"></a>

The date columns indicate important events that are used for various calculations and trend analysis.

* **INTRODUCED** - Timestamp of the first scan that identified the issue.
* **LAST IGNORED** - Timestamp indicating when the issue was last marked as ignored.
* **LAST INTRODUCED** - Timestamp of the last scan that identified this issue.
* **LAST RESOLVED** - The last timestamp at which the issue was not found in a test.
