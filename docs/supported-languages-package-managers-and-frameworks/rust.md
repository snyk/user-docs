# Rust

## Applicability

Snyk supports [Rust for code analysis](rust/rust-for-code-analysis.md) and Rust for open source.

Snyk supports the testing of Rust applications through the API. For details, see [Test an SBOM document for vulnerabilities](../snyk-api/how-to-use-snyk-sbom-and-list-issues-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).

Snyk also supports the testing of individual Rust packages from the Cargo package manager. For details, see [List issues for a package](../snyk-api/how-to-use-snyk-sbom-and-list-issues-apis/list-issues-for-a-package.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* Test your app's SBOM using`pkg:cargo`
* Test your app's packages using`pkg:cargo`

## Package managers and supported file extensions

Snyk for Rust supports Cargo as a package manager and [crates.io](https://crates.io/) as a package registry and does not support any file formats.

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Rust:

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

If you want to use **Open Source** PR checks in Rust projects, for which the Fix PRs feature is not supported, then vulnerable dependencies may be introduced without being flagged.

## Features

The following features are supported in Snyk for Rust:

| Snyk Open Source | Snyk Code                                                        |
| ---------------- | ---------------------------------------------------------------- |
| None             | <ul><li>Support for Interfile analysis</li><li>Reports</li></ul> |

If you need help, [contact Snyk Support](https://support.snyk.io).
