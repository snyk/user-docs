# Bazel

**Note:** The language pages are being restructured and updated on June 22 and June 23.

{% hint style="info" %}
Snyk for Bazel is supported only for Snyk Open Source. The instructions on this page apply to Bazel v 7 only.
{% endhint %}

Snyk supports testing Projects that have their dependencies managed by Bazel. The recommended approach is to test and monitor using the [Snyk API v1 Dep Graph endpoints](dep-graph-api.md). While you can use Bazel for many languages including C++, the Dep Graph endpoints do not support C++.

## Getting started with Snyk for Bazel&#x20;

Bazel is defined on [bazel.build](https://docs.bazel.build/versions/master/bazel-overview.html) as follows:

> _Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users._

Bazel does not have dependency manifest files or lock files that package managers such as npm have. Instead, build configuration is managed in [BUILD](https://docs.bazel.build/versions/master/build-ref.html#BUILD\_files) files, using [Starlark](https://docs.bazel.build/versions/master/skylark/language.html), a domain-specific language based on Python3.

Bazel has limited native integration with package registries such as npmjs.org or Maven Central. Some Bazel rules can be added to help with installing dependencies from external registries, for example, [from Maven](https://docs.bazel.build/versions/master/external.html#maven-artifacts-and-repositories).

You must often manually specify all dependencies (package name, location, and version), including transitive dependencies, which can then be fetched by Bazel during builds.

Because Bazel dependencies are specified as code in BUILD files using Starlark, Snyk cannot easily discover the dependencies from a Project. For details about testing Bazel Projects using Snyk, see [Snyk API v1 Dep Graph endpoints](dep-graph-api.md).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
