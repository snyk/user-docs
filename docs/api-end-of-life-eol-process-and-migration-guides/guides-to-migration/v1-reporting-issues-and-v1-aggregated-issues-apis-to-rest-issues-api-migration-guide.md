# V1 Reporting Issues and V1 Aggregated Issues APIs to REST Issues API migration guide



{% hint style="info" %}
This page applies to the [V1 List All Aggregated Issues API](../../snyk-api/reference/projects-v1.md#post-org-orgid-project-projectid-aggregated-issues) (`post-org-orgid-project-projectid-aggregated-issues`) and to the V1 Reporting Issues API endpoints: [Get list of issues](../../snyk-api/reference/reporting-api-v1.md#post-reporting-issues) (`post-reporting-issues`) and [Get list of latest issues](../../snyk-api/reference/reporting-api-v1.md#post-reporting-issues-latest) (`post-reporting-issues-latest`).

While the deprecation schedule for the V1 Issues API endpoints has not been established, these APIs will be deprecated soon, and migration to the REST Issues API is highly recommended as soon as possible. Additionally, the [V1 Reporting Issues API](../../snyk-api/reference/reporting-api-v1.md) is not available for customers or partners in the `Snyk-US-02`, `Snyk-EU-01`, and `Snyk-AU-01` regions.&#x20;
{% endhint %}

## Highlights of the GA REST Issues API

{% hint style="info" %}
GA REST Issues API documentation [`/groups/{group_id}/issues`](https://apidocs.snyk.io/#get-/groups/-group_id-/issues) and [`/orgs/{org_id}/issues`](https://apidocs.snyk.io/#get-/orgs/-org_id-/issues). Updated link: [Get issues by Org ID](../../snyk-api/reference/issues.md#get-orgs-org_id-issues)
{% endhint %}

The REST version of the API delivers:

* **Consistency:** Improved performance and reliability of the REST Issues API
* **Depth:** detailed representations for Open Source packages, SAST, Container, and IaC issues
* **Flexibility:** new filters for tailored API responses
* **Usability:** improved pagination and response management, simplifying the API interaction

Snyk understands that migrating to a new API can be a significant undertaking, and we want to support you throughout the process. This comprehensive migration guide is intended to facilitate a seamless transition by providing step-by-step instructions, code examples, and best practices to help you smoothly integrate with the new API.

If you are using the V1 Issues endpoints, Snyk recommends reviewing this migration guide and migrating all your automations, API based reporting, and integrations accordingly.

## Comparison of V1 Issues APIs vs REST GA API

{% hint style="info" %}
**Mapping V1 Issues API to REST GA API issues**

In the V1 Issues APIs, the id or issue.id is relative to the project to define uniqueness. To define the uniqueness in the REST API, a UUID is generated for each issue.&#x20;
{% endhint %}

| **v1 Aggregated Project Issues**    | **V1 Reporting Issues** | **REST unified issues schema**                                                  | **Notes**                                                                                                                                                |
| ----------------------------------- | ----------------------- | ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id                                  | issue.id                | problems.id\*                                                                   | \*where the source is "Snyk"                                                                                                                             |
|                                     |                         | id                                                                              | The unique value of the issue in the REST API                                                                                                            |
| issue.url                           | issue.url               | For SCA, built from “http://security.snyk.io/vuln/” + key                       |                                                                                                                                                          |
| issueType                           | issue.type              | problems.type                                                                   |                                                                                                                                                          |
| pkgName                             | issue.package           | coordinates.representations.dependency.package\_name                            |                                                                                                                                                          |
| pkgVersions                         | issue.version           | coordinates.representations.dependency.package\_version                         | Each package name representation has it's own version coordinate.                                                                                        |
| issueData.id                        | issue.id                | problems.id\*                                                                   | \*where the source is "Snyk"                                                                                                                             |
| issueData.title                     | issue.title             | title                                                                           |                                                                                                                                                          |
| issueData.severity                  | issue.severity          | effective\_severity                                                             |                                                                                                                                                          |
| issueData.originalSeverity          | issue.originalSeverity  | original\_severity                                                              | Not returned if no changes to severity have been made.                                                                                                   |
| issueData.url                       | issue.url               | problem.url                                                                     |                                                                                                                                                          |
| issueData.description               | N/A                     | Available for issue type = code                                                 | The description attribute for CIM\* issues is currently populated only for Code issues. For all the other issue types, this information is not available |
| issueData.identifiers.CVE           | issue.identifiers       | problems.id\*                                                                   | \*where the source is "NVD"                                                                                                                              |
| issueData.identifiers.CWE           | issue.identifiers       | classes.id                                                                      |                                                                                                                                                          |
| issueData.identifiers.OSVDB         | issue.identifiers       | N/A                                                                             |                                                                                                                                                          |
| issueData.credit                    | issue.credit            | N/A                                                                             |                                                                                                                                                          |
| issueData.exploitMaturity           | issue.exploitMaturity   | exploit\_details.maturity\_levels                                               | For REST, the maturity levels have a different format (standard CVSS4 naming convention).                                                                |
| issueData.semver.vulnerable         | issue.semver.vulnerable | N/A                                                                             | Fix information is not available yet in the REST API.                                                                                                    |
| issueData.semver.unaffected         | issue.semver.vulnerable | N/A                                                                             |                                                                                                                                                          |
| issueData.publicationTime           | issue.publicationTime   | N/A                                                                             |                                                                                                                                                          |
| issueData.disclosureTime            | issue.disclosureTime    | N/A                                                                             |                                                                                                                                                          |
| issueData.CVSSv3                    | issue.CVSSv3            | severities.vector                                                               | This is an array in the REST response by source (NIST, Snyk, RedHat, and so on).                                                                         |
| issueData.cvssScore                 | issue.cvssScore         | severities.score                                                                | This is an array in the REST response by source (NIST, Snyk, RedHat, and so on).                                                                         |
| issueData.language                  | issue.language          | Data exists at the project level, inferred from the project type in field: type |                                                                                                                                                          |
| issueData.patches.id                | issue.patches           | N/A                                                                             | Likely added in a future update.                                                                                                                         |
| issueData.patches.urls              |                         | N/A                                                                             | Likely added in a future update.                                                                                                                         |
| issueData.patches.version           |                         | N/A                                                                             | Likely added in a future update.                                                                                                                         |
| issueData.patches.comments          |                         | N/A                                                                             | Likely added in a future update.                                                                                                                         |
| issueData.patches.modificationTime  |                         | N/A                                                                             | Likely added in a future update.                                                                                                                         |
| issueData.nearestFixedInVersion     |                         | N/A                                                                             | Likely added in a future update.                                                                                                                         |
| issueData.path                      |                         |                                                                                 | Represented across multiple attributes only for Cloud issues.                                                                                            |
| issueData.violatedPolicyPublicId    |                         | problem.id or class.id                                                          |                                                                                                                                                          |
| issueData.isMaliciousPackage        |                         | N/A                                                                             | No data.                                                                                                                                                 |
| introducedThrough.kind              |                         | N/A                                                                             | Specific to dependency chains which are not supported by the CIM\*.                                                                                      |
| introducedThrough.data              |                         | N/A                                                                             | Specific to dependency chains which are not supported by the CIM\*.                                                                                      |
| isPatched                           | issue.isPatched         | N/A                                                                             |                                                                                                                                                          |
| isIgnored                           | issue.isIgnored         | ignored                                                                         |                                                                                                                                                          |
| ignoreReasons.reason                | issue.ignored.reason    | N/A                                                                             | Can be retrieved with V1 List Ignores API.                                                                                                               |
| ignoreReasons.expires               | issue.ignored.expires   | N/A                                                                             | Can be retrieved with V1 List Ignores API.                                                                                                               |
| ignoreReasons.source                | issue.ignored.source    | N/A                                                                             | Can be retrieved with V1 List Ignores API.                                                                                                               |
| fixInfo.isUpgradable                | issue.isUpgradeable     | coordinates.is\_upgradable                                                      |                                                                                                                                                          |
| fixInfo.isPinnable                  | issue.isPinnable        | coordinates.is\_pinnable                                                        |                                                                                                                                                          |
| fixInfo.isPatchable                 | issue.isPatchable       | coordinates.is\_patchable                                                       |                                                                                                                                                          |
| fixInfo.isFixable                   | N/A                     | coordinates.is\_fixable                                                         |                                                                                                                                                          |
| fixInfo.isPartiallyFixable          | N/A                     | N/A                                                                             | Infered by looking at all of the coordinates that belong to an issue.                                                                                    |
| fixInfo.nearestFixedInVersion       | part of semver          | N/A                                                                             |                                                                                                                                                          |
| priority.score                      | issue.priorityScore     | risk.score                                                                      | Priority score value is replaced by risk score. Previous score migrated to this field in REST.                                                           |
| priority.factors                    | N/A                     | risk.factors                                                                    | Priority factors are replaced by risk score.                                                                                                             |
| priority.factors.name = “Reachable” | issue.reachability      | coordinates.reachability                                                        |                                                                                                                                                          |
| links                               | N/A                     | N/A                                                                             | New links object contains pagination and relationship links.                                                                                             |

\* CIM refers to Snyk's "Common Issues Model."

## Notes on migration

* Filtering options vary between API versions. Refer to the [API Reference](../../snyk-api/reference/) for each API to validate your filtering choices.
* As noted above, the REST API  paginates results. Code accordingly.
