# Rust

## Applicability and integration

{% hint style="info" %}
Rust is supported for Snyk Code and for Snyk Open Source.
{% endhint %}

Available functions:

* Test your app's SBOM using `pkg:cargo` For more information, see [Test an SBOM document for vulnerabilities](../../snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).
* Test your app's packages using `pkg:cargo`
* Test your individual Rust packages from the Cargo package manager. For more information, see [List issues for a package](../../snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package.md).

## Technical specifications

### Supported frameworks and libraries

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

### Supported package managers and registries

* Supported package manager: Cargo
* Supported package registry: [crates.io](https://crates.io/)

## Rust for Snyk Code

For Rust with Snyk Code, the following features are supported:

* Support for Interfile analysis
* Reports

{% hint style="info" %}
Code analysis support for Rust is available with Snyk Preview.&#x20;
{% endhint %}

## Rust for Snyk Open Source

For Rust with Snyk Open Source, there are no features supported.

{% hint style="warning" %}
If you want to use Open Source PR checks in Rust Projects for which the **Fix PRs** feature is not supported, it is possible that vulnerable dependencies are introduced without being flagged.
{% endhint %}
