# Bazel

{% hint style="info" %}
Snyk for Bazel is supported only for Snyk Open Source. The instructions on this page apply to Bazel v 7 only.
{% endhint %}

Snyk supports testing Projects that have their dependencies managed by Bazel. The recommended approach is to test and monitor using the [Snyk API v1 Dep Graph endpoints](../../snyk-api/snyk-api-v1-dep-graph-endpoints.md). While you can use Bazel for many languages including C++, the Dep Graph endpoints do not support C++.

## Getting started with Snyk for Bazel&#x20;

Bazel is defined on [bazel.build](https://docs.bazel.build/versions/master/bazel-overview.html) as follows:

> _Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users._

Bazel does not have dependency manifest files or lock files that package managers such as npm have. Instead, build configuration is managed in [BUILD](https://docs.bazel.build/versions/master/build-ref.html#BUILD\_files) files, using [Starlark](https://docs.bazel.build/versions/master/skylark/language.html), a domain-specific language based on Python3.

Bazel has limited native integration with package registries such as npmjs.org or Maven Central. Some Bazel rules can be added to help with installing dependencies from external registries, for example, [from Maven](https://docs.bazel.build/versions/master/external.html#maven-artifacts-and-repositories).

You must often manually specify all dependencies (package name, location, and version), including transitive dependencies, which can then be fetched by Bazel during builds.

Because Bazel dependencies are specified as code in BUILD files using Starlark, Snyk cannot easily discover the dependencies from a Project. For details about testing Bazel Projects using Snyk, see [Snyk API v1 Dep Graph endpoints](../../snyk-api/snyk-api-v1-dep-graph-endpoints.md).

## Examples of Snyk for Bazel

{% hint style="info" %}
See [Manually creating a Dep Graph from Bazel Java project](https://github.com/snyk/bazel-simple-app) for a full example Bazel Java project and the corresponding Snyk Dep Graph object.
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

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
