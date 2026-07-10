---
description: Snyk support for Bazel with Snyk Open Source, including testing Bazel v7 Projects through the Dep Graph API
---

# Bazel

## Applicability

{% hint style="info" %}
Snyk supports Bazel only for Snyk Open Source.

Snyk for Bazel provides support for using the [Bazel build and test tool](https://docs.bazel.build/versions/master/bazel-overview.html) with Snyk Open Source. The instructions in this documentation apply to Bazel v 7 only.
{% endhint %}

Snyk supports testing Projects whose dependencies are managed by Bazel. Snyk recommends testing and monitoring using the Dep Graph API.

Unlike npm, Bazel does not rely on dependency manifest files or lock files. Instead, you manage build configurations in [BUILD](https://docs.bazel.build/versions/master/build-ref.html#BUILD_files) files using [Starlark](https://docs.bazel.build/versions/master/skylark/language.html), a domain-specific language based on Python 3.

You manually specify all dependencies (package name, location, and version), including transitive dependencies. Bazel fetches these dependencies during builds.

Bazel has limited native integration with package registries, such as npmjs.org or Maven Central. You can add Bazel rules to help install dependencies from external registries.

Because Bazel dependencies are specified as code in BUILD files using Starlark, Snyk cannot easily discover the dependencies from a Project.

## Dep Graph API

To secure Bazel Projects, you must use the Snyk Dep Graph API. This API accepts a generic dependency graph and returns a report containing any relevant vulnerabilities for those dependencies.

### Requirements and considerations

The Dep Graph API requires specific permissions. If you do not have access, contact Snyk Support.

You can test Bazel dependencies across any supported ecosystem, except C++, which is not supported by these endpoints.

Use the Snyk Dep Graph API endpoints [Test Dep Graph](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/test-v1) and [Monitor Dep Graph](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/monitor-v1) to test and monitor dependencies managed by Bazel. The monitor capability allows you to submit a tree for Snyk to monitor for vulnerabilities.

### Test and monitor dependencies

To integrate Snyk into your Bazel workflow, follow these steps to manually generate and submit a dependency graph to the Snyk API:

1. Create a [Dep Graph JSON object](https://github.com/snyk/dep-graph) listing all the dependency packages and versions for each type of dependency (for example, Maven or CocoaPods).
2. Send the Dep Graph JSON object as a POST request to the Test Dep Graph endpoint, along with your [auth token](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-api/authentication-for-api), as part of a Bazel test rule.
3. Check the API response for pass or fail status and any resulting vulnerabilities.

For example:

```bash
curl -X POST 'https://api.snyk.io/v1/test/dep-graph' \
  -H 'Authorization: token {{your token}}' \
  -H 'Content-Type: application/json; charset=utf-8' \
  -d @dep-graph.json
```

### Dep Graph JSON syntax

The Test Dep Graph API accepts a Snyk Dep Graph JSON object. This object describes the root application and the graph of direct and transitive dependencies.

The [schema](https://github.com/snyk/dep-graph#depgraphdata) for this format is:

{% code overflow="wrap" fullWidth="false" %}
```json
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

Specific components in the Dep Graph object include:

* `schemaVersion` - the version of the Dep Graph schema. Set this to `1.2.0`.
* `pkgManager.name` - can be one of `deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, or `cocoapods`.
* `pkgs` - an array of objects containing `id`, `name` and `version` of all packages in the Dep Graph. The `id` must be in the form `name@version`. List each of your dependencies in this array, including an item representing the Project itself.
* `graph.nodes` - an array of objects describing the relationships between entries in `pkgs`. This is typically the Project node with all other packages defined as a flat array of direct dependencies in `deps.`
* `graph.rootNodeId` - specifies the `id` of the entry in `graph.nodes` to use as the root node of the graph. Set this to the `nodeId` of the Project node.

### Dep Graph Test API response

The Test Dep Graph API returns a JSON object describing any issues (vulnerabilities and licenses) found in the Dep Graph dependencies.

An example response with a single vulnerability:

{% code overflow="wrap" %}
```json
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

Specific components in the response object include:

* `ok` - Boolean value summarizing whether Snyk found any vulnerabilities in the supplied dependencies. You can use this for a quick pass or fail test.
* `issuesData` - a hash of each unique vulnerability found. Each vulnerability contains useful properties, such as `title`, `description`, `identifiers`, `publicationTime`, `severity`, and so on.
* `issues` - an array of mappings from vulnerabilities in `issuesData` to package. This mapping shortens the response length because a vulnerability can apply to multiple packages.

### Example of dependency mapping for a Bazel Project

For a Bazel Project with a single dependency on a Maven package, you can specify the dependency as follows:

```python
maven_jar(
    name = "logback-core",
    artifact = "ch.qos.logback:logback-core:1.0.13",
    sha1 = "dc6e6ce937347bd4d990fc89f4ceb469db53e45e",
)
```

Use the provided template to construct the following Dep Graph JSON object:

```json
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

This package (`ch.qos.logback:logback-core@1.0.13`) contains a vulnerability described in detail in the resulting JSON response object.
