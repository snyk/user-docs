# Snyk for Bazel

Snyk supports testing projects that have their dependencies managed by Bazel.

Support is available through the [Snyk API](../../../snyk-api/).

## Features of Snyk for Bazel

{% hint style="info" %}
**Feature availability**\
The Snyk API is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
The Dep Graph API requires additional permissions. [Contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new) to request access.
{% endhint %}

The following describes how to use Snyk to test your Bazel projects.

## Bazel overview

Bazel is defined on [bazel.build](https://docs.bazel.build/versions/master/bazel-overview.html) as follows:

> _Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users._

Bazel does not have dependency manifest files or lock files that package managers such as npm have. Instead, build configuration is managed in [BUILD](https://docs.bazel.build/versions/master/build-ref.html#BUILD\_files) files, using [Starlark](https://docs.bazel.build/versions/master/skylark/language.html), a domain specific language based on Python3.

Bazel has limited native integration with package registries such as npmjs.org or Maven Central. There are some Bazel rules that can be added to help with installing dependencies from external registries, for example, [from Maven](https://docs.bazel.build/versions/master/external.html#maven-artifacts-and-repositories).

However, in many cases, users must manually specify their dependency information (package name, location. and version), including all transitives. These can then be fetched by Bazel during builds.

Because Bazel dependencies are specified as code in BUILD files using Starlark, Snyk cannot easily discover which dependencies a project has.

The recommended approach is to test your dependencies using the [Snyk Dep Graph](https://snyk.docs.apiary.io/#reference/test/dep-graph) API. For more information see [Snyk Dep Graph Test API](https://github.com/snyk/dep-graph).

## How Snyk for Bazel works

1. For each type of dependency (for example, Maven, Cocoapods), create a [Dep Graph JSON object](https://github.com/snyk/dep-graph) listing all the dependency packages and versions (see the following).
2.  As part of a Bazel test rule, send this object as a POST request to the [Dep Graph Test API](https://snyk.docs.apiary.io/#reference/test/dep-graph), (along with your [auth token](../../../snyk-api-info/authentication-for-api.md)). An example curl request follows:

    ```
    curl -X POST 'https://snyk.io/api/v1/test/dep-graph' \
      -H 'Authorization: token {{your token}}' \
      -H 'Content-Type: application/json; charset=utf-8' \
      -d @dep-graph.json
    ```
3. Check the API response for pass/fail status and any resulting vulnerabilities. See the next section of this doc.

### Snyk Dep Graph Test API

The Snyk Dep Graph Test API takes a generic dependency graph and returns a report containing any relevant vulnerabilities for those dependencies.

The supported package managers and repository ecosystems are listed in the [API documentation](https://snyk.docs.apiary.io/#reference/test/dep-graph) (at the time of writing these are `deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, and `cocoapods`).

Any of your Bazel dependencies that are available in these ecosystems can be tested using the Snyk API.

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

Further notes on specific components in the dep graph object follow:

* `schemaVersion` - version of the dep-graph schema. Set this to `1.2.0`.
* `pkgManager.name` - one of `deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, or `cocoapods`
* `pkgs` - array of objects containing `id`, `name`and`version` of all packages in the dep-graph. Note that the `id` _must_ be of the form `name@version`. List each of your dependencies in this array, including an item representing the Project itself.
* `graph.nodes` - array of objects describing the relationships between entries in `pkgs`. In Bazel. This is typically the project node with all other packages defined as a flat array of direct dependencies in `deps.`
* `graph.rootNodeId` - specifies the `id` of the entry in `graph.nodes` to use as the root node of the graph. Set this to the `nodeId` of the project node.

### Snyk Dep Graph Test API response

The Dep Graph Test API returns a JSON object describing any issues (vulnerabilities and licenses) found in the dep graph dependencies.

An example response with a single vulnerability. follows:

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

* `ok` - Boolean value summarizing whether Snyk found any vulnerabilities in the supplied dependencies. You can use this for a quick pass/fail test
* `issuesData` - a hash of each unique vulnerability found. Each vulnerability contains many useful properties, such as `title`, `description`, `identifiers`, `publicationTime`, `severity`, and so on.
* `issues` - an simple array of mappings from vulnerabilities in `issuesData` to package. As a vulnerability may be relevant to multiple packages, this mapping is used to keep the response length as short as possible

## Examples for Snyk for Bazel

{% hint style="info" %}
See [Manually creating a Dep Graph from Bazel Java project](https://github.com/snyk/bazel-simple-app) for a full example Bazel Java project and corresponding Snyk Dep Graph object.
{% endhint %}

For a simple Bazel project with a single dependency on a Maven package, you may specify the dependency like this:

```
maven_jar(
    name = "logback-core",
    artifact = "ch.qos.logback:logback-core:1.0.13",
    sha1 = "dc6e6ce937347bd4d990fc89f4ceb469db53e45e",
)
```

From this you could construct the following Dep Graph JSON object:

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

This particular package (`ch.qos.logback:logback-core@1.0.13`) contains a vulnerability, described in detail in the resulting JSON response object.
