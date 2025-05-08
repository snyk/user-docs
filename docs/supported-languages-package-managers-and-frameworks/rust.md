# Rust

{% hint style="info" %}
Rust is supported for Snyk Code and for Snyk Open Source.
{% endhint %}

## Applicability

The following functions are available for Rust:

* Test your app's SBOM using `pkg:cargo` For more information, see [Test an SBOM document for vulnerabilities](../snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).
* Test your app's packages using `pkg:cargo`
* Test your individual Rust packages from the Cargo package manager. For more information, see [List issues for a package](../snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package.md).

## Package managers and supported file extensions

For Rust, Snyk supports Cargo as a package manager and [crates.io](https://crates.io/) as a package registry. It does not support any file formats.

## Frameworks and libraries

For Rust with Snyk Open Source, the following frameworks and libraries are supported:

* Warp - Comprehensive
* Rust standard library - Comprehensive
* Iron - Comprehensive
* tokio - Comprehensive
* Hyper - Comprehensive
* axum - Comprehensive
* tower - Partial
* age - Comprehensive
* ammonia - Comprehensive
* diesel - Comprehensive
* orion - Comprehensive
* postgres - Comprehensive
* ring - Comprehensive
* rustcrypto - Comprehensive
* sqlx - Comprehensive
* Reqwest - Comprehensive

## Features

For Rust with Snyk Code, the following features are supported:

* Support for Interfile analysis
* Reports

{% hint style="info" %}
Code analysis support for Rust is available with **Snyk Preview**.&#x20;
{% endhint %}

For Rust with Snyk Open Source, there are no features supported.

{% hint style="info" %}
If you want to use Open Source PR checks in Rust Projects for which the **Fix PRs** feature is not supported, it is possible that vulnerable dependencies are introduced without being flagged.
{% endhint %}
