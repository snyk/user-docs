---
hidden: true
---

# V1 Issues APIs to REST Issues API migration guide

{% hint style="info" %}
This page applies to the following V1  API endpoints:

* [Get list of issues](../../reference/reporting-api-v1.md#post-reporting-issues) (`POST /reporting/issues`) - not available for customers or partners in the `Snyk-US-02`, `Snyk-EU-01`, and `Snyk-AU-01` regions.
* [Get list of latest issues](../../reference/reporting-api-v1.md#post-reporting-issues-latest) (`POST /reporting/issues/latest`) - not available for customers or partners in the `Snyk-US-02`, `Snyk-EU-01`, and `Snyk-AU-01` regions.
* [List all aggregated issues](../../reference/projects-v1.md#post-org-orgid-project-projectid-aggregated-issues) (`POST /org/{orgId}/project/{projectId}/aggregated-issues`)

These APIs are subject to deprecation, and migrating to the REST Issues API immediately is highly recommended. The REST Issues API accommodates multiple regions.
{% endhint %}

## Key benefits of the REST Issues API

{% hint style="info" %}
Use the following REST API endpoints instead to get lists of issues:

* [Get issues by Group ID](../../reference/issues.md#get-groups-group_id-issues)
* [Get issues by Org ID](../../reference/issues.md#get-orgs-org_id-issues)
{% endhint %}

The REST Issues API provides:

* Consistency through improved performance and reliability of the REST Issues API.
* In-depth coverage of Open Source packages, SAST, Container, and IaC issues.
* Flexibility through new filters for tailored API responses.
* Enhanced usability through better pagination and response handling, streamlining API interactions.

This migration guide is intended to facilitate a seamless transition by providing step-by-step instructions, code examples, and best practices to help you smoothly integrate with the new API.

If you are using the V1 Issues endpoints, Snyk recommends that you migrate all your automation, API-based reporting, and integrations accordingly, by using the guidance provided here.

## Comparison of V1 Issues APIs vs REST Issues API

{% hint style="info" %}
In the V1 Issues APIs, the `id` or `issue.id` is relative to the Project to define uniqueness. &#x20;

In the REST Issues API, a UUID is generated for each issue to define uniqueness.&#x20;
{% endhint %}

| V1 Project issues                     | V1 Reporting Issues       | REST unified issues                                                             | Notes                                                                                                               |
| ------------------------------------- | ------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `id`                                  | `issue.id`                | `problems.id`\*                                                                 | \*where the source is "Snyk"                                                                                        |
| N/A                                   | N/A                       | `id`                                                                            | The unique value of the issue in the REST API                                                                       |
| `issue.url`                           | `issue.url`               | For SCA, built from “http://security.snyk.io/vuln/” + key                       |                                                                                                                     |
| `issueType`                           | `issue.type`              | `problems.type`                                                                 |                                                                                                                     |
| `pkgName`                             | `issue.package`           | `coordinates.representations.dependency.package_name`                           |                                                                                                                     |
| `pkgVersions`                         | `issue.version`           | `coordinates.representations.dependency.package_version`                        | Each package name representation has its own version coordinate.                                                    |
| `issueData.id`                        | `issue.id`                | `problems.id`\*                                                                 | \*where the source is "Snyk"                                                                                        |
| `issueData.title`                     | `issue.title`             | `title`                                                                         |                                                                                                                     |
| `issueData.severity`                  | `issue.severity`          | `effective_severity`                                                            |                                                                                                                     |
| `issueData.originalSeverity`          | `issue.originalSeverity`  | `original_severity`                                                             | Not returned if no changes to severity have been made.                                                              |
| `issueData.url`                       | `issue.url`               | `problem.url`                                                                   |                                                                                                                     |
| `issueData.description`               | N/A                       | N/A                                                                             | Availability of this field may be limited to projects that have been updated (new issues found).                    |
| `issueData.identifiers.CVE`           | `issue.identifiers`       | `problems.id`\*                                                                 | \*where the source is "NVD"                                                                                         |
| `issueData.identifiers.CWE`           | `issue.identifiers`       | `classes.id`                                                                    |                                                                                                                     |
| `issueData.identifiers.OSVDB`         | `issue.identifiers`       | N/A                                                                             |                                                                                                                     |
| `issueData.credit`                    | `issue.credit`            | N/A                                                                             |                                                                                                                     |
| `issueData.exploitMaturity`           | `issue.exploitMaturity`   | `exploit_details.maturity_levels`                                               | For REST, the maturity levels have a different format (standard CVSS4 naming convention).                           |
| `issueData.semver.vulnerable`         | `issue.semver.vulnerable` | N/A                                                                             | Fix information is not available yet in the REST API.                                                               |
| `issueData.semver.unaffected`         | `issue.semver.vulnerable` | N/A                                                                             |                                                                                                                     |
| `issueData.publicationTime`           | `issue.publicationTime`   | N/A                                                                             |                                                                                                                     |
| `issueData.disclosureTime`            | `issue.disclosureTime`    | N/A                                                                             |                                                                                                                     |
| `issueData.CVSSv3`                    | `issue.CVSSv3`            | `severities.vector`                                                             | This is an array in the REST response by source (NIST, Snyk, RedHat, and so on).                                    |
| `issueData.cvssScore`                 | `issue.cvssScore`         | `severities.score`                                                              | This is an array in the REST response by source (NIST, Snyk, RedHat, and so on).                                    |
| `issueData.language`                  | `issue.language`          | Data exists at the Project level, inferred from the project type in field: type |                                                                                                                     |
| `issueData.patches.id`                | `issue.patches`           | N/A                                                                             |                                                                                                                     |
| `issueData.patches.urls`              | N/A                       | N/A                                                                             |                                                                                                                     |
| `issueData.patches.version`           | N/A                       | N/A                                                                             |                                                                                                                     |
| `issueData.patches.comments`          | N/A                       | N/A                                                                             |                                                                                                                     |
| `issueData.patches.modificationTime`  | N/A                       | N/A                                                                             |                                                                                                                     |
| `issueData.nearestFixedInVersion`     | N/A                       | N/A                                                                             |                                                                                                                     |
| `issueData.path`                      | N/A                       | N/A                                                                             | Represented across multiple attributes only for Cloud issues.                                                       |
| `issueData.violatedPolicyPublicId`    | N/A                       | problem.id or class.id                                                          |                                                                                                                     |
| `issueData.isMaliciousPackage`        | N/A                       | N/A                                                                             | No data.                                                                                                            |
| `introducedThrough.kind`              | N/A                       | N/A                                                                             | Specific to dependency chains which may not be available.                                                           |
| `introducedThrough.data`              | N/A                       | N/A                                                                             | Specific to dependency chains which may not be available.                                                           |
| `isPatched`                           | `issue.isPatched`         | N/A                                                                             |                                                                                                                     |
| `isIgnored`                           | `issue.isIgnored`         | `ignored`                                                                       |                                                                                                                     |
| `ignoreReasons.reason`                | `issue.ignored.reason`    | N/A                                                                             | Can be retrieved with [V1 List Ignores API](../../reference/ignores-v1.md#get-org-orgid-project-projectid-ignores). |
| `ignoreReasons.expires`               | `issue.ignored.expires`   | N/A                                                                             | Can be retrieved with [V1 List Ignores API](../../reference/ignores-v1.md#get-org-orgid-project-projectid-ignores). |
| `ignoreReasons.source`                | `issue.ignored.source`    | N/A                                                                             | Can be retrieved with [V1 List Ignores API](../../reference/ignores-v1.md#get-org-orgid-project-projectid-ignores). |
| `fixInfo.isUpgradable`                | `issue.isUpgradeable`     | `coordinates.is_upgradable`                                                     |                                                                                                                     |
| `fixInfo.isPinnable`                  | `issue.isPinnable`        | `coordinates.is_pinnable`                                                       |                                                                                                                     |
| `fixInfo.isPatchable`                 | `issue.isPatchable`       | `coordinates.is_patchable`                                                      |                                                                                                                     |
| `fixInfo.isFixable`                   | N/A                       | `coordinates.is_fixable`                                                        |                                                                                                                     |
| `fixInfo.isPartiallyFixable`          | N/A                       | N/A                                                                             | Inferred by looking at all of the coordinates that belong to an issue.                                              |
| `fixInfo.nearestFixedInVersion`       | part of semver            | N/A                                                                             |                                                                                                                     |
| `priority.score`                      | `issue.priorityScore`     | `risk.score`                                                                    | Priority score value is replaced by risk score. Previous score migrated to this field in REST.                      |
| `priority.factors`                    | N/A                       | `risk.factors`                                                                  | Priority factors are replaced by risk score.                                                                        |
| `priority.factors.name = “Reachable”` | `issue.reachability`      | `coordinates.reachability`                                                      |                                                                                                                     |
| `links`                               | N/A                       | N/A                                                                             | New links object contains pagination and relationship links.                                                        |
