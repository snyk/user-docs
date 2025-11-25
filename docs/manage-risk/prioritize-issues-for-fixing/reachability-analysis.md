# Reachability analysis

{% hint style="info" %}
**Release status**

Reachability analysis is in General Availability (GA) for specific [integrations and programming languages](reachability-analysis.md#supported-languages-and-integrations).
{% endhint %}

Snyk reachability analysis allows you to analyze risk by identifying whether your application is calling a code element (for example, functions, classes, modules, and more) related to the vulnerability, thus raising the chances of that vulnerability being exploited in the context of your application.

Reachability analysis can be used as an indicator to make decisions, or as part of a broader risk-based prioritization approach using the Snyk Risk Score.

Snyk uses a combination of static program analysis and various AI techniques to determine the reachability of a given vulnerability, with validation conducted by security research experts. These capabilities enable Snyk to quickly analyze the code without requiring the application to be built prior to the scan.

To use this feature, Snyk must analyze your source code. To learn more, visit [How Snyk handles your data](../../snyk-data-and-governance/how-snyk-handles-your-data.md).

## How reachable vulnerability analysis works

Snyk uses a combination of security expert analysis, program analysis, and various AI techniques to determine the reachability of a vulnerability. Analysis steps include:

* Enriching vulnerabilities with the patches applied to fix them - as part of the Snyk vulnerability curation process, Snyk references the fix commit that the maintainer applied.
* Related elements analysis - Based on the commit fix, Snyk uses DeepCode AI program analysis to analyze the code elements and other parameters related to the vulnerability.
* Root Cause analysis - Snyk uses DeepCode AI and Natural Language Processing (NLP) techniques to automatically rank the related code elements by their chances of being the root cause of the vulnerability.
* Reachability analysis - As issues are found in your application by a Snyk scan, the DeepCode program analysis engine is used to analyze the call graph of your application in relation to the call graph between the open-source dependencies used. A path between your application and a code element ranked as a root cause will yield a “Reachable” vulnerability.
* Security experts supervision - Snyk security experts will manually verify and mark elements as root causes in order to make the entire analysis more accurate over time

The following considerations related to false positives and false negatives apply to Reachable vulnerability analysis.

Program analysis requires a trade-off between accurate results, minimizing false positives, and recall rates, by avoiding potentially exploitable vulnerabilities.

To facilitate this trade-off, Snyk DeepCode analysis applies real-time decision-making to determine whether to under-approximate the set of reachable elements based on analysis of the likelihood that a reachable path will be found in a specific environment.

For example, it is not always possible to give a precise answer when reflection programming is used. In such a case, neither returning a large set of false positives nor returning “Not reachable” will suffice. The agentic capability of Snyk analysis optimizes in order to retrieve the most accurate and complete result possible for a given code structure.

{% hint style="info" %}
Changes in the first-party code, in the vulnerability analysis, and enhancements in the SAST engine can impact the results. Therefore, vulnerabilities labeled as **Not path found** can change to **Reachable** over time.
{% endhint %}

## Limits of reachability in static analysis

Snyk aims to identify and demonstrate when a code element is reachable, while also addressing the challenges of indicating that a code element cannot be reached.

Static analysis techniques can show that a vulnerability or code element can be reached through at least one execution path. However, just because there is no evidence of this does not mean that the element cannot be reached.

A code element that is not marked as reachable may still be accessible under conditions that were not considered during analysis. This can occur due to incomplete information, control flow issues, potential dynamic behaviors, or overlooked edge cases.

Reachability is an important factor in assessing security risks. When evaluating these risks, consider both reachable issues and those that are not reachable. This helps you evaluate security risks as a whole, making sure you do not overlook any potential threats.

## Supported languages and integrations

Reachability analysis is supported for the following languages and package managers:

| Language                                                                                                                                                     | Package manager     | Release status       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------- | -------------------- |
| [Java](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/)                                                                           | Maven, Gradle       | General Availability |
| [JavaScript](../../supported-languages/supported-languages-list/javascript/), [TypeScript](../../supported-languages/supported-languages-list/typescript.md) | npm, Yarn, pnpm     | General Availability |
| [Python](../../supported-languages/supported-languages-list/python/)                                                                                         | pip, poetry, pipenv | Early Access         |
| [C#](../../supported-languages/supported-languages-list/.net/)                                                                                               | NuGet, paket        | Early Access         |

Reachability analysis is supported in the following integrations:

| Integration                                                                                                                                              | Release status       |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------- |
| [GitHub](../../developer-tools/scm-integrations/organization-level-integrations/github.md)                                                               | General Availability |
| [GitHub Enterprise](../../developer-tools/scm-integrations/organization-level-integrations/github-enterprise.md)                                         | General Availability |
| [GitHub Cloud App](../../developer-tools/scm-integrations/organization-level-integrations/github-cloud-app.md)                                           | General Availability |
| [Bitbucket Cloud](../../developer-tools/scm-integrations/organization-level-integrations/bitbucket-cloud-app.md)                                         | General Availability |
| [Bitbucket Server](../../developer-tools/scm-integrations/organization-level-integrations/bitbucket-data-center-server.md)                               | General Availability |
| [GitLab](../../developer-tools/scm-integrations/organization-level-integrations/gitlab.md)                                                               | General Availability |
| [Azure Repos](../../developer-tools/scm-integrations/organization-level-integrations/azure-repositories-tfs.md)                                          | General Availability |
| [Brokered connections](../../implementation-and-setup/enterprise-setup/snyk-broker/broker-inbound-and-outbound-connections-and-allowed-requests.md)      | General Availability |
| [Snyk CLI](../../developer-tools/snyk-cli/)                                                                                                              | Early Access         |
| [AWS CodePipeline integration with CodeBuild](../../developer-tools/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage.md) | Early Access         |
| [Azure Pipelines](../../developer-tools/snyk-ci-cd-integrations/azure-pipelines-integration/)                                                            | Early Access         |
| [Bitbucket Pipelines](../../developer-tools/snyk-ci-cd-integrations/bitbucket-pipelines-integration-using-a-snyk-pipe/)                                  | Early Access         |
| [CircleCI](../../developer-tools/snyk-ci-cd-integrations/circleci-integration-using-a-snyk-orb.md)                                                       | Early Access         |
| [GitHub Actions](../../developer-tools/snyk-ci-cd-integrations/github-actions-for-snyk-setup-and-checking-for-vulnerabilities/)                          | Early Access         |
| [Jenkins](../../developer-tools/snyk-ci-cd-integrations/jenkins-plugin-integration-with-snyk.md)                                                         | Early Access         |
| [Maven](../../developer-tools/snyk-ci-cd-integrations/maven-plugin-integration-with-snyk.md)                                                             | Early Access         |

## Set up reachability analysis

### Supported configuration options

Snyk supports multiple reachability analysis configuration options:

* **At the Group level**: You can set the default Reachability analysis setting for all Organizations using the parent Group's Setting&#x73;**.** The default setting applies to existing Organizations with no reachability analysis configuration and all new Organizations.
* **At the Organization level**: You can enable reachability for a single Organization in the Organization Settings. The reachability analysis Organization setting overrides the Group level default.
* **Enable for all Organizations in a Group**: You can enable and disable reachability analysis for all Organizations in a Group using the parent Group's Settings. Bulk settings changes override any existing reachability analysis Organization setting.

### Enable reachability at the Group level

You can set the reachability analysis as the default setting for multiple Organizations. To do this:

1. Navigate to Group **Settings** > **Snyk Open Source**.
2. Under **Reachability analysis,** click **Enable Reachability** to apply it to new Organizations.
3. Confirm the selection to save changes.

<figure><img src="../../.gitbook/assets/image (10).png" alt="Enable reachability at the Group level"><figcaption><p>Enable reachability at the Group level</p></figcaption></figure>

After reachability analysis is enabled, Snyk performs the analysis as part of scanning Projects.

{% hint style="info" %}
You can apply the reachability analysis to existing Projects by triggering a [manual test](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/#manual-snyk-prs).
{% endhint %}

### Enable reachability at the Organization level

To enable reachability analysis for your Organization and begin analyzing Projects for reachable vulnerabilities:

1. Navigate to Organization **Settings** > **Snyk Open Source**.
2. Under **General**, select **Enable reachability analysis**.
3. Confirm the selection to save changes.

<figure><img src="../../.gitbook/assets/image (198).png" alt="Enable reachability at the Organization level"><figcaption><p>Enable reachability at the Organization level</p></figcaption></figure>

After reachability analysis is enabled, the analysis is done as part of scanning Projects.

{% hint style="info" %}
You can apply the reachability analysis to existing Projects by triggering a [manual test](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/#manual-snyk-prs).
{% endhint %}

### Enable reachability for all Organizations in a Group

You can enable and disable reachability for all Organizations in a Group, overriding all existing Organization-level settings.

1. In the Group **Settings**, navigate to **Snyk Open Source**.
2. Under **General**, select **Enable/Disable reachability analysis** to activate it for all Organizations.
3. Confirm the selection to save changes.

<figure><img src="../../.gitbook/assets/image (1) (5).png" alt="Enable reachability analysis for all Organizations in a Group"><figcaption><p>Enable reachability for all Organizations in a Group</p></figcaption></figure>

### Enable reachability for Snyk CLI and CI/CD integrations

{% hint style="info" %}
Reachability analysis for Snyk CLI and CI/CD integrations is in Early Access.
{% endhint %}

Snyk supports performing reachability analysis through the CLI and CI/CD integrations. When **Reachability analysis for Snyk CLI and CI/CD integrations** is enabled, the Snyk Preview toggle is on by default, for all Groups and Organizations.

To enable or disable **Reachability in the Snyk CLI and CI/CD integrations** for a Group or Org:

1. Navigate to Group or Org **Settings** > **Snyk Preview**.
2. Under **Reachability for Snyk CLI and CI/CD integrations**, toggle **Enable Reachability for Snyk CLI and CI/CD integrations** on or off.
3. Confirm the selection to save changes.

<figure><img src="../../.gitbook/assets/image (45) (1).png" alt="Enable Reachability for Snyk CLI and CI/CD integrations in Snyk Preview"><figcaption><p>Enable Reachability for Snyk CLI and CI/CD integrations in Snyk Preview</p></figcaption></figure>

### **Enable reachability** analysis **for brokered connections**

If you use a brokered connection to your SCM, configure [Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/) to provide access to your source files.

## Use reachability analysis in the Snyk Web UI

After it is identified, a vulnerability has one of the following reachability statuses:

* `REACHABLE` - A direct or indirect path was found from your application to the vulnerable code.
* `NO PATH FOUND` - No path found from your application to the vulnerable code.
* `NOT APPLICABLE` - Reachability is not supported for this vulnerability.

{% hint style="info" %}
A vulnerability with the status `NO PATH FOUND` it does not mean that the vulnerability is completely unreachable or unexploitable.
{% endhint %}

Reachability analysis status is available on the Project page, as part of the Risk Score, in the [Issues Detail report](../reporting/available-snyk-reports.md#issues-detail-report), and through the API endpoint [Get issues by Group ID](../../snyk-api/reference/issues.md#groups-group_id-issues).

### Reachability analysis as shown on the Project page

After you import or test a Project using the Snyk Web UI, Snyk monitors the Project, and the results of the reachable vulnerabilities analysis appear on the Project page as follows:

* As a filter - The filter **Reachable** allows you to filter results based on reachability.
* Reachability badge - Allows you to quickly see on the issue card when a vulnerability is reachable.
* Call path - Allows you to see the path from your code to the vulnerable code element to validate the result.

<figure><img src="broken-reference" alt="Reachability filters, badge and call path on the Projects UI"><figcaption><p>Reachability filters, badge and call path on the Projects UI</p></figcaption></figure>

### Reachability analysis as part of the Risk Score

[Risk Score](risk-score.md) helps you apply holistic risk-based prioritization that combines multiple factors, associated with either the vulnerability or the context of your application. Reachability analysis is such a contextual factor that will significantly increase the overall score.

Risk Score is available on the Projects page and through the API and Reports.

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (125).png" alt="Reachability as part of the Risk Score"><figcaption><p>Reachability as part of the Risk Score</p></figcaption></figure></div>

{% hint style="info" %}
[Priority score](priority-score.md), the legacy model preceding the Risk Score, also takes reachable vulnerabilities into account.
{% endhint %}

## Using reachability analysis with Snyk CLI

Snyk supports Reachability CLI version `1.1301.0` and greater. To perform Reachability analysis, use the `--reachability=true` option.

{% hint style="info" %}
When using `--reachability`, Snyk returns a new findings schema. Some legacy fields may not be available in this schema. Before using, please first check your usage to ensure automations do not fail.
{% endhint %}

By default, Snyk uploads source code from the current working directory for Reachability analysis. To specify source files for analysis in a separate directory, use the `--source-dir` option.

To filter findings on Reachability, use the `--reachability-filter=reachable` option.

For a full list of supported options, see the CLI help docs for `test`, `monitor`, and `sbom test` CLI commands.

### Human-readable output

By default, findings will be returned in human-readable output with **Reachability** as a property of each finding.

<figure><img src="../../.gitbook/assets/image (46).png" alt=""><figcaption></figcaption></figure>

### JSON output

To print the findings in JSON, use the `--json` option. `reachability` is a property of `vulnerabilities`.

```json
{
  "dependencyCount": 12,
  "displayTargetFile": "log4shell-server/pom.xml",
  "filesystemPolicy": false,
  "filtered": {
    "ignore": [],
    "patch": []
  },
  "hasUnknownVersions": false,
  "ignoreSettings": {
    "adminOnly": false,
    "autoApproveIgnores": false,
    "disregardFilesystemIgnores": false,
    "reasonRequired": false
  },
  "isPrivate": true,
  "licensesPolicy": {
    "orgLicenseRules": null,
    "severities": null
  },
  "ok": false,
  "org": "mysnykorg",
  "packageManager": "maven",
  "path": "/Users/jck/workspace/java-goof/log4shell-goof",
  "policy": "",
  "projectName": "io.snyk:log4shell-server",
  "remediation": {
    "ignore": null,
    "pin": {},
    "unresolved": [],
    "upgrade": {}
  },
  "summary": "10 vulnerable dependency paths",
  "targetFile": "",
  "uniqueCount": 10,
  "vulnerabilities": [
    {
      "CVSSv3": "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L/E:P/RL:O/RC:R",
      "creationTime": "2019-10-10T18:31:03.943542Z",
      "credit": [
        "Unknown"
      ],
      "cvssDetails": [
        {
          "assigner": "NVD",
          "cvssV3BaseScore": 7.3,
          "cvssV3Vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L",
          "modificationTime": "2024-03-11T07:51:53.729801Z",
          "severity": "high"
        }
      ],
      "cvssScore": 5.6,
      "cvssSources": [
        {
          "assigner": "Snyk",
          "baseScore": 5.6,
          "cvssVersion": "3.1",
          "modificationTime": "2024-03-06T13:45:01.631734Z",
          "severity": "medium",
          "type": "primary",
          "vector": "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L/E:P/RL:O/RC:R"
        },
        {
          "assigner": "NVD",
          "baseScore": 7.3,
          "cvssVersion": "3.1",
          "modificationTime": "2024-03-11T07:51:53.729801Z",
          "severity": "high",
          "type": "secondary",
          "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L"
        }
      ],
      "description": "## Overview\n\n[commons-collections:commons-collections](https://mvnrepository.com/artifact/commons-collections/commons-collections) is a library which contains types that extend and augment the Java Collections...",
      "disclosureTime": "2019-10-10T00:00:00.000000Z",
      "epssDetails": {
        "modelVersion": "v2025.03.14",
        "percentile": "0.93995",
        "probability": "0.13769"
      },
      "exploit": "Proof of Concept",
      "exploitDetails": {
        "maturityLevels": [
          {
            "format": "CVSSv3",
            "level": "Proof of Concept",
            "type": "secondary"
          },
          {
            "format": "CVSSv4",
            "level": "Proof of Concept",
            "type": "primary"
          }
        ],
        "sources": [
          "Snyk"
        ]
      },
      "fixedIn": [
        "3.2.2"
      ],
      "from": [
        "io.snyk:log4shell-server@0.0.1-SNAPSHOT",
        "commons-collections:commons-collections@3.1"
      ],
      "functions": [],
      "id": "SNYK-JAVA-COMMONSCOLLECTIONS-472711",
      "identifiers": {
        "CVE": [
          "CVE-2015-6420"
        ],
        "CWE": [
          "CWE-502"
        ],
        "GHSA": [
          "GHSA-6hgm-866r-3cjv"
        ]
      },
      "insights": {
        "triageAdvice": null
      },
      "isPatchable": false,
      "isUpgradable": true,
      "language": "java",
      "malicious": false,
      "modificationTime": "2024-03-11T07:51:53.729801Z",
      "name": "commons-collections:commons-collections",
      "packageManager": "maven",
      "packageName": "commons-collections:commons-collections",
      "patches": [],
      "publicationTime": "2020-02-24T00:00:00.000000Z",
      "reachability": "reachable",
      "references": [
        {
          "title": "GitHub Commit",
          "url": "https://github.com/apache/commons-collections/commit/5ec476b0b756852db865b2e442180f091f8209ee"
        },
        {
          "title": "GitHub PR",
          "url": "https://github.com/apache/commons-collections/pull/18"
        },
        {
          "title": "Jira Ticket",
          "url": "https://issues.apache.org/jira/browse/COLLECTIONS-580"
        }
      ],
      "semver": {
        "vulnerable": [
          "[,3.2.2)"
        ]
      },
      "severity": "medium",
      "severityWithCritical": "medium",
      "socialTrendAlert": false,
      "title": "Deserialization of Untrusted Data",
      "upgradePath": [
        false,
        "commons-collections:commons-collections@3.2.2"
      ],
      "version": "3.1"
    }
  ]
}Score": 9.8,
      "cvssSources": [
        {
          "assigner": "Snyk",
          "baseScore": 9.8,
          "cvssVersion": "3.1",
          "modificationTime": "2025-02-10T14:22:01.344278Z",
          "severity": "critical",
          "type": "primary",
          "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H/E:H/RL:O"
        },
        {
          "assigner": "NVD",
          "baseScore": 9.8,
          "cvssVersion": "3.0",
          "modificationTime": "2024-03-11T09:46:27.924934Z",
          "severity": "critical",
          "type": "secondary",
          "vector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H"
        },
        {
          "assigner": "Red Hat",
          "baseScore": 7.3,
          "cvssVersion": "3.1",
          "modificationTime": "2024-03-11T09:52:38.421377Z",
          "severity": "high",
          "type": "secondary",
          "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L"
        }
      ],
      "description": "## Overview\n[commons-collections:commons-collections](https://mvnrepository.com/artifact/commons-collections/commons-collections) is a library which contains types that extend and augment the Java Collections Framework.\n\nAffected versions of this package are ...\n",
      "disclosureTime": "2015-11-06T16:51:56.000000Z",
      "epssDetails": null,
      "exploitDetails": {
        "maturityLevels": [
          {
            "format": "",
            "level": "",
            "type": "secondary"
          },
          {
            "format": "",
            "level": "",
            "type": "primary"
          }
        ],
        "sources": []
      },
      "from": [],
      "isPatchable": false,
      "isUpgradable": false,
      "name": "commons-collections:commons-collections",
      "packageManager": "maven",
      "packageName": "commons-collections:commons-collections",
      "severity": "critical",
      "title": "Deserialization of Untrusted Data",
      "upgradePath": null,
      "version": "3.1"
    }
  ]
}
```

### Technical specifications and guidance

Using Reachability will perform an upload of source code files to Snyk for analysis. The following technical limitations exist:

* **Total number of files:** 300,000 files
* **Total size of files:** 3GB
* **Total size per file:** 1MB. See the [Technical specifications and guidance](https://docs.snyk.io/supported-languages/technical-specifications-and-guidance#file-size-limit-for-snyk-code-analysis) for more details.
* **File name length:** 256 characters
