# Critical severity migration

##  Critical severity migration

A new **Critical** severity level is assigned to vulnerabilities with a [CVSS](https://www.first.org/cvss/) score between 9.0-10.0.

Use this documentation to understand the change, and understand the actions your organization needs to take, to support this change.

#### Background

At Snyk, we use [CVSS framework version 3.1](https://www.first.org/cvss/v3-1/) to communicate the characteristics and severity of vulnerabilities.

Previously, we used three severity levels: **High**, **Medium** and **Low**. We have now added a **Critical** level to align with CVSS severity levels and provide better visibility on the highest severity issues.

This change applies to Snyk Open Source and Snyk Container only. It does not apply to licenses, Snyk Code or Snyk IaC functionality.

#### CVSS levels

| Severity level | CVSS 2 | CVSS 3 |
| :--- | :--- | :--- |
| Low | 0.0 - 3.9 | 0.0 - 3.9 |
| Medium | 4.0 - 6.9 | 4.0 - 6.9 |
| High | 7.0 - 10.0 | 7.0 - 8.9 |
| Critical | N/A | 9.0 - 10.10 |

#### Change timescale

* **24th May 2021**, this functionality is available at an "opt-in" level. We recommend you use this period to make any adjustments required to your own systems and processes \(for example, updating SLAs\), then opt-in to the change using the **Snyk Preview** feature.
* **28th June 2021**, this functionality is enabled for all customers.

#### Impacts of change

After the new functionality is enabled, all vulnerabilities with CVSS above 9.0 become **Critical**, not **High**. This will likely reduce the number of High severity vulnerabilities for your project, as some High severity vulnerabilities will become Critical.

#### Processing the change

When implemented, you may notice a reduction in the number of High severity issues across your projects, as some of these issues now have a Critical severity level.

Also, this change affects automated CLI/API-based pipelines, so consider if you have any internal pipelines or workflows that will be affected by this change, read below to learn about the changes in detail, then update your internal processes accordingly.

### Enabling critical severity

Before enabling the feature, ensure you are using the latest version of the Snyk CLI, CI/CD plugins, and IDE Extensions. If you use the API, or parse the JSON output from the CLI, also update your workflow to accept the new values.

To enable the feature during the "opt-in" period, click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Snyk Preview**, then enable the critical severity feature. 

Snyk Preview is available at the Group and Organization levels. If you enable the feature in the Group settings, the feature will be enabled for all Organizations within the same Group.

You must retest each project for the changes to take effect \(either retest manually or **wait for the next scheduled testing**\).

#### CLI Prerequisites

The minimum required version of the [Snyk CLI](https://support.snyk.io/hc/en-us/categories/360000456217-Snyk-CLI) to see critical issues is `1.514.0` . Use the command `snyk -v` to verify your version. To upgrade to the latest version, run `npm i -g snyk`.

#### CI/CD Prerequisites

The minimum required versions for Snyk's [CI/CD Integration plugins,](https://support.snyk.io/hc/en-us/sections/360001152577-CI-CD-integrations) to support critical severity functionality, are as follows.

| **Integration** | **Minimum supported version** | **Update necessity** |
| :--- | :--- | :--- |
| [Nexus Gatekeeper](https://support.snyk.io/hc/en-us/articles/360004127697-Nexus-Repository-Manager-Gatekeeper-plugin) | [0.3.0](https://github.com/snyk/nexus-snyk-security-plugin/releases/tag/v0.3.0) | **Required** |
| [Artifactory Gatekeeper](https://support.snyk.io/hc/en-us/articles/360004032417-Artifactory-Gatekeeper-plugin-overview) | [1.1.0](https://github.com/snyk/artifactory-snyk-security-plugin/releases/tag/1.1.0) | **Required** |
| [TeamCity](https://support.snyk.io/hc/en-us/articles/360004032297-TeamCity-integration-overview) | [2018.2-13](https://plugins.jetbrains.com/plugin/12227-snyk-security) | **Required** |
| [Maven](https://support.snyk.io/hc/en-us/articles/360004570477-Maven-plugin-integration) | [2.0.0](https://github.com/snyk/snyk-maven-plugin/releases/tag/v2.0.0) | **Required** |
| [Azure Pipelines](https://support.snyk.io/hc/en-us/articles/360004127677-Azure-Pipelines-integration) | [0.4.0](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan) | Optional |
| [Bitbucket Pipelines](https://support.snyk.io/hc/en-us/articles/360004032177-Bitbucket-Pipelines-integration-overview) | [0.5.0](https://bitbucket.org/product/features/pipelines/integrations?p=snyk/snyk-scan) | Optional |
| [CircleCI](https://support.snyk.io/hc/en-us/articles/360004032197-Getting-Snyk-Orb-details-from-the-CircleCI-registry) | [0.1.0](https://circleci.com/developer/orbs/orb/snyk/snyk) | Optional |
| [Jenkins](https://support.snyk.io/hc/en-us/articles/360004032217-Jenkins-integration-overview) | [2.14.0](https://plugins.jenkins.io/snyk-security-scanner/) | Optional |

\* Optional = Needs to be updated in order to support Critical as a Severity Threshold

#### IDE Extension prerequisites

| **Integration** | **Minimum supported version** | **Update necessity** |
| :--- | :--- | :--- |
| [JetBrains](https://support.snyk.io/hc/en-us/articles/360004032317-JetBrains-IDE-Plugins) | [2.1.4](https://plugins.jetbrains.com/plugin/10972-snyk-vulnerability-scanner) | Recommended |
| [Visual Studio](https://support.snyk.io/hc/en-us/articles/360020073958-Visual-Studio-extension) | [1.0.5](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs) | Recommended |
| [Eclipse](https://support.snyk.io/hc/en-us/articles/360004032337-Eclipse-Snyk-plugin-overview) | [1.3.1](https://marketplace.eclipse.org/content/snyk-security-scanner) | Recommended |

\* Recommended = Update is recommended in order to see issues as Critical

### Snyk web app changes

Filtering by Critical severity is now available in the **Project** and **Reports** pages:

#### PR Checks

If you enabled **Fail for high severity issues** under Pull-request test in the **Settings** page for your integration, you see text to confirm that critical and high severity issues fail for the PR when there is a fix available:

After implementing the change:

* `[Critical Severity]` appears for all issues with CVSS above 9.0.
* `"severity": "critical"` appears when using `--json` or when exporting data to external json file \(using `--json-file-output` flag\)
* Set the flag `snyk test --severity-threshold=critical` to see reports for critical vulnerabilities only. This flag is also available when testing docker images: `snyk test --docker debian:8 --severity-threshold=critical`

On every endpoint with the attribute `severity` or `issuesCountsBySeverity` you can now also see **Critical**.

This includes the following endpoints:

* [Projects Endpoints](https://snyk.docs.apiary.io/#reference/projects)
  * [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)
  * [Retrieve a single project](https://snyk.docs.apiary.io/#reference/projects/individual-project/retrieve-a-single-project)
  * [Update a project](https://snyk.docs.apiary.io/#reference/projects/all-projects/update-a-project)
  * [List all aggregated issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues)
* [Dependencies](https://snyk.docs.apiary.io/#reference/dependencies/)
  * [List all dependencies](https://snyk.docs.apiary.io/#reference/dependencies/dependencies-by-organization/list-all-dependencies)
* [All Test endpoints](https://snyk.docs.apiary.io/#reference/test)
* [Reporting API](https://snyk.docs.apiary.io/#reference/reporting-api)
  * [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues)
  * [Get list of issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues)
  * [Get latest issues count](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issue-counts/get-latest-issue-counts)
  * [Get issues counts over time](https://snyk.docs.apiary.io/#reference/reporting-api/issue-counts-over-time/get-issue-counts)

#### Response examples

```text
"issueCountsBySeverity": {
     "critical": 5
     "high": 8,       
     "medium": 13,        
     "low": 15      
},
```

```text
"issueData": {
  "id": "SNYK-JS-LODASH-590103",
  "title": "Prototype Pollution",
  "severity": "critical",
  "url": "<https://snyk.io/vuln/SNYK-JS-LODASH-590103>",
  "identifiers": {
    "CVE": [],
    "CWE": [
      "CWE-400"
    ]
  },
```

For every endpoint that you filter by `severities`, you should add `critical` value.

For example, for List of all Aggregated endpoint, you should add to the POST message the Critical filter

```text
{  
  "filters": {    
	"severities": [  
		"critical",   
		"high",    
		"medium",     
		"low"
		],
	}
}
```

If you have [security policies](https://support.snyk.io/hc/en-us/sections/360004225818-Security-Policies) with rules that change a severity to High, consider changing it to Critical instead.

