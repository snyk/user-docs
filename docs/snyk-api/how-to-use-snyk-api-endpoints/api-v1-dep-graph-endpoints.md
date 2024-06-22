# API v1 Dep Graph endpoints

{% hint style="info" %}
**Feature availability**\
The Snyk API is available with Enterprise plans. See [plans and pricing](https://snyk.io/plans/) for details.

The Dep Graph API requires additional permissions. [Contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new) to request access.
{% endhint %}

The recommended approach to testing and monitoring your dependencies managed by [Bazel](../../supported-languages-and-package-managers/bazel.md) is to use the [Snyk Dep Graph](https://snyk.docs.apiary.io/#reference/test/dep-graph) v1 APIs. The monitor capability allows customers to submit a tree for Snyk to monitor for vulnerabilities. While you can use Bazel for many languages including C++, the Dep Graph endpoints do not support C++. For more information, including a list of the supported package managers, see [Snyk Dep Graph Test v1 API](https://github.com/snyk/dep-graph).

Follow these basic steps:

1. For each type of dependency (for example, Maven, Cocoapods), create a [Dep Graph JSON object](https://github.com/snyk/dep-graph) listing all the dependency packages and versions. For an example, see the Snyk user docs [Bazel](../../supported-languages-and-package-managers/bazel.md) page.
2.  As part of a Bazel test rule, send the Dep Braph JSON object as a POST request to the [Dep Graph Test API](https://snyk.docs.apiary.io/#reference/test/dep-graph), along with your [auth token](../rest-api/authentication-for-api/). An example curl request follows:

    ```
    curl -X POST 'https://snyk.io/api/v1/test/dep-graph' \
      -H 'Authorization: token {{your token}}' \
      -H 'Content-Type: application/json; charset=utf-8' \
      -d @dep-graph.json
    ```
3. Check the API response for pass/fail status and any resulting vulnerabilities. For more information, continue with [Snyk Dpe Graph Test API](api-v1-dep-graph-endpoints.md#snyk-dep-graph-test-api).

### Snyk Dep Graph Test API

The Snyk Dep Graph Test API takes a generic dependency graph and returns a report containing any relevant vulnerabilities for those dependencies.

The supported package managers and repository ecosystems are listed in the [API documentation](https://snyk.docs.apiary.io/#reference/test/dep-graph) (`deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, and `cocoapods`).

Any of your Bazel dependencies that are available in these ecosystems can be tested using the Snyk API v1.

### Snyk Dep Graph JSON syntax

The Dep Graph Test API takes a [Snyk Dep Graph](https://github.com/snyk/dep-graph) JSON object describing the root application and the graph of direct and transitive dependencies.

The [schema](https://github.com/snyk/dep-graph#depgraphdata) for this format is as follows:

```
export interface DepGraphData {
  schemaVersion: string;
  pkgManager: {
    name: string;
    version?: string;
    repositories?: Array<{
      alias: string;
    }>;
  };
  pkgs: Array<{
    id: string;
    info: {
      name: string;
      version?: string;
    };
  }>;
  graph: {
    rootNodeId: string;
    nodes: Array<{
      nodeId: string;
      pkgId: string;
      info?: {
        versionProvenance?: {
          type: string;
          location: string;
          property?: {
            name: string;
          };
        },
        labels?: {
          [key: string]: string | undefined;
        };
      };
      deps: Array<{
        nodeId: string;
      }>;
    }>;
  };
}
```

Further notes on specific components in the Dep Graph object follow:

* `schemaVersion` - the version of the Dep Graph schema. Set this to `1.2.0`.
* `pkgManager.name` - can be one of `deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, or `cocoapods`.
* `pkgs` - an array of objects containing `id`, `name`and`version` of all packages in the Dep Graph. Note that the `id` must be in the form `name@version`. List each of your dependencies in this array, including an item representing the Project itself.
* `graph.nodes` - an array of objects describing the relationships between entries in `pkgs`. This is typically the Project node with all other packages defined as a flat array of direct dependencies in `deps.`
* `graph.rootNodeId` - specifies the `id` of the entry in `graph.nodes` to use as the root node of the graph. Set this to the `nodeId` of the Project node.

### Snyk Dep Graph Test API response

The Dep Graph Test API returns a JSON object describing any issues (vulnerabilities and licenses) found in the Dep Graph dependencies.

An example response with a single vulnerability:

```
{
    "ok": false,
    "packageManager": "maven",
    "issuesData": {
        "SNYK-JAVA-CHQOSLOGBACK-30208": {
            "CVSSv3": "CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
            "alternativeIds": [],
            "creationTime": "2017-03-19T14:58:38Z",
            "credit": [
                "Unknown"
            ],
            "cvssScore": 9.8,
            "description": "## Overview\n[ch.qos.logback:logback-core](https://mvnrepository.com/artifact/ch.qos.logback/logback-core) is a logback-core module.\n\nAffected versions of this package are vulnerable to Arbitrary Code Execution. A configuration can be ...",
            "disclosureTime": "2017-03-13T06:59:00Z",
            "exploit": "Not Defined",
            "fixedIn": [
                "1.1.11"
            ],
            "functions": [],
            "id": "SNYK-JAVA-CHQOSLOGBACK-30208",
            "identifiers": {
                "CVE": [
                    "CVE-2017-5929"
                ],
                "CWE": [
                    "CWE-502"
                ]
            },
            "language": "java",
            "mavenModuleName": {
                "artifactId": "logback-core",
                "groupId": "ch.qos.logback"
            },
            "modificationTime": "2020-06-12T14:36:56.271247Z",
            "moduleName": "ch.qos.logback:logback-core",
            "packageManager": "maven",
            "packageName": "ch.qos.logback:logback-core",
            "patches": [],
            "proprietary": false,
            "publicationTime": "2017-03-21T15:30:44Z",
            "references": [
                {
                    "title": "GitHub Commit #1",
                    "url": "https://github.com/qos-ch/logback/commit/f46044b805bca91efe5fd6afe52257cd02f775f8"
                },
                {
                    "title": "GitHub Commit #2",
                    "url": "https://github.com/qos-ch/logback/commit/979b042cb1f0b4c1e5869ccc8912e68c39f769f9"
                },
                {
                    "title": "Logback News",
                    "url": "https://logback.qos.ch/news.html"
                },
                {
                    "title": "NVD",
                    "url": "https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2017-5929"
                },
                {
                    "title": "NVD",
                    "url": "https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2017-5929/"
                }
            ],
            "semver": {
                "vulnerable": [
                    "[, 1.1.11)"
                ]
            },
            "severity": "high",
            "title": "Arbitrary Code Execution"
        }
    },
    "issues": [
        {
            "pkgName": "ch.qos.logback:logback-core",
            "pkgVersion": "1.0.13",
            "issueId": "SNYK-JAVA-CHQOSLOGBACK-30208",
            "fixInfo": {}
        }
    ],
    "org": {
        "id": "3e5fe3fe-9181-4f0f-a231-39764485e73f",
        "name": "stephen.elson-xnf"
    }
}
```

Further notes on specific components in the response object follow:

* `ok` - Boolean value summarizing whether Snyk found any vulnerabilities in the supplied dependencies. You can use this for a quick pass or fail test.
* `issuesData` - a hash of each unique vulnerability found. Each vulnerability contains many useful properties, such as `title`, `description`, `identifiers`, `publicationTime`, `severity`, and so on.
* `issues` - a simple array of mappings from vulnerabilities in `issuesData` to package. As a vulnerability may be relevant to multiple packages, this mapping is used to keep the response length as short as possible.
