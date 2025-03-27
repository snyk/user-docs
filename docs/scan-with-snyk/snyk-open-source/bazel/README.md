# Bazel - a build and test tool

{% hint style="info" %}
**Feature availability**\
Snyk for Bazel is a build and test tool supported by Snyk Open Source. The instructions in this documentation apply to Bazel v 7 only.
{% endhint %}

## Applicability

Snyk for Bazel is supported only for Snyk Open Source.

Snyk supports testing Projects that have their dependencies managed by Bazel. The recommended approach is to test and monitor using the [Dep Graph API](dep-graph-api.md). While you can use Bazel for many languages including C++, the Dep Graph endpoints do not support C++.

## Package managers

Bazel does not have dependency manifest files or lock files that package managers such as npm have. Instead, build configuration is managed in [BUILD](https://docs.bazel.build/versions/master/build-ref.html#BUILD_files) files, using [Starlark](https://docs.bazel.build/versions/master/skylark/language.html), a domain-specific language based on Python3.

You must often manually specify all dependencies (package name, location, and version), including transitive dependencies, which can then be fetched by Bazel during builds.

Bazel has limited native integration with package registries such as npmjs.org or Maven Central. Some Bazel rules can be added to help with installing dependencies from external registries, for example, [from Maven](https://docs.bazel.build/versions/master/external.html#maven-artifacts-and-repositories).

Because Bazel dependencies are specified as code in BUILD files using Starlark, Snyk cannot easily discover the dependencies from a Project. For information about testing and monitoring Bazel Projects using Snyk, see the page [Dep Graph API](dep-graph-api.md).

## Frameworks and libraries

The are no available frameworks and libraries listed for Snyk for Bazel.

## Features

There are no available features listed for Snyk for Bazel.

## Dep Graph API

{% hint style="info" %}
**Feature availability**\
The Snyk API is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

The Dep Graph API requires additional permissions. [Contact Snyk Support](https://support.snyk.io) to request access.

To test and monitor dependencies managed by [Bazel](./), it is recommended that you use the Snyk Dep Graph API endpoints [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph) and [Monitor Dep Graph](../../../snyk-api/reference/monitor-v1.md). The monitor capability allows customers to submit a tree for Snyk to monitor for vulnerabilities. While you can use Bazel for many languages including C++, the Dep Graph endpoints do not support C++.

Follow these basic steps:

1. For each type of dependency, for example, Maven, Cocoapods, create a [Dep Graph JSON object](https://github.com/snyk/dep-graph) listing all the dependency packages and versions. See [Example of Snyk for Baszel](example-of-snyk-for-bazel.md).
2.  As part of a Bazel test rule, send the Dep Graph JSON object as a POST request to the endpoint [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph), along with your [auth token](../../../snyk-api/rest-api/authentication-for-api/). An example curl request follows:

    ```
    curl -X POST 'https://api.snyk.io/v1/test/dep-graph' \
      -H 'Authorization: token {{your token}}' \
      -H 'Content-Type: application/json; charset=utf-8' \
      -d @dep-graph.json
    ```
3. Check the API response for pass/fail status and any resulting vulnerabilities.

### How the Test Dep Graph API works

The Test Dep Graph API takes a generic dependency graph and returns a report containing any relevant vulnerabilities for those dependencies.

The supported package managers and repository ecosystems are listed in the [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph) and [Monitor Dep Graph](../../../snyk-api/reference/monitor-v1.md) documentation.

Any of your Bazel dependencies that are available in the supported ecosystems can be tested using the Snyk API.

### Snyk Dep Graph JSON syntax

The Test Dep Graph API takes a [Snyk Dep Graph](https://github.com/snyk/dep-graph) JSON object describing the root application and the graph of direct and transitive dependencies.

The [schema](https://github.com/snyk/dep-graph#depgraphdata) for this format is as follows:

{% code overflow="wrap" fullWidth="false" %}
```java
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
{% endcode %}

Further notes on specific components in the Dep Graph object follow:

* `schemaVersion` - the version of the Dep Graph schema. Set this to `1.2.0`.
* `pkgManager.name` - can be one of `deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, or `cocoapods`.
* `pkgs` - an array of objects containing `id`, `name`and`version` of all packages in the Dep Graph. Note that the `id` must be in the form `name@version`. List each of your dependencies in this array, including an item representing the Project itself.
* `graph.nodes` - an array of objects describing the relationships between entries in `pkgs`. This is typically the Project node with all other packages defined as a flat array of direct dependencies in `deps.`
* `graph.rootNodeId` - specifies the `id` of the entry in `graph.nodes` to use as the root node of the graph. Set this to the `nodeId` of the Project node.

### Snyk Dep Graph Test API response

The TEst Dep Graph API returns a JSON object describing any issues (vulnerabilities and licenses) found in the Dep Graph dependencies.

An example response with a single vulnerability follows:

{% code overflow="wrap" %}
```java
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
{% endcode %}

Further notes on specific components in the response object follow:

* `ok` - Boolean value summarizing whether Snyk found any vulnerabilities in the supplied dependencies. You can use this for a quick pass or fail test.
* `issuesData` - a hash of each unique vulnerability found. Each vulnerability contains many useful properties, such as `title`, `description`, `identifiers`, `publicationTime`, `severity`, and so on.
* `issues` - a simple array of mappings from vulnerabilities in `issuesData` to package. As a vulnerability may be relevant to multiple packages, this mapping is used to keep the response length as short as possible.

## Example of Snyk for Bazel

{% hint style="info" %}
See [Manually creating a Dep Graph from Bazel Java project](https://github.com/snyk/bazel-simple-app) for a full example of a Bazel Java project and the corresponding Snyk Dep Graph object.
{% endhint %}

For a simple Bazel Project with a single dependency on a Maven package, you may specify the dependency like this:

```
maven_jar(
    name = "logback-core",
    artifact = "ch.qos.logback:logback-core:1.0.13",
    sha1 = "dc6e6ce937347bd4d990fc89f4ceb469db53e45e",
)
```

From this, you could construct the following Dep Graph JSON object:

```
{
  "depGraph": {
    "schemaVersion": "1.2.0",
    "pkgManager": {
      "name": "maven"
    },
    "pkgs": [
      {
        "id": "app@1.0.0",
        "info": {
          "name": "app",
          "version": "1.0.0"
        }
      },
      {
        "id": "ch.qos.logback:logback-core@1.0.13",
        "info": {
          "name": "ch.qos.logback:logback-core",
          "version": "1.0.13"
        }
      }
    ],
    "graph": {
      "rootNodeId": "root-node",
      "nodes": [
        {
          "nodeId": "root-node",
          "pkgId": "app@1.0.0",
          "deps": [
            {
              "nodeId": "ch.qos.logback:logback-core@1.0.13"
            }
          ]
        },
        {
          "nodeId": "ch.qos.logback:logback-core@1.0.13",
          "pkgId": "ch.qos.logback:logback-core@1.0.13",
          "deps": []
        }
      ]
    }
  }
}
```

This particular package (`ch.qos.logback:logback-core@1.0.13`) contains a vulnerability described in detail in the resulting JSON response object.
