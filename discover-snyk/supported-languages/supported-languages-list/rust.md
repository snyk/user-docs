---
nav_context: agnostic
description: >-
  Snyk support for Rust, with full Snyk Code coverage in Early Access on
  Enterprise plans and limited Snyk Open Source support
---

# Rust

{% hint style="info" %}
Snyk supports Rust for Snyk Code (full support) and for Snyk Open Source (limited support).
{% endhint %}

## Rust for Snyk Code

{% hint style="info" %}
Code analysis support for Rust is in Early Access and is available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://docs.snyk.io/platform-administration/snyk-hierarchy/snyk-preview).
{% endhint %}

For an overview of the supported security rules, visit [Rust rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/rust-rules).

### Supported frameworks and libraries

For Rust with Snyk Code, Snyk supports the following frameworks and libraries:

{% columns %}
{% column %}
* actix\_files
* actix\_identity
* actix\_multipart
* actix\_session
* actix\_web
* age
* ammonia
* async\_graphql
* axum
* diesel
* handlebars
* hyper
* iron
{% endcolumn %}

{% column %}
* orion
* postgres
* reqwest
* ring
* rustcrypto
* sqlx
* tera
* tokio
* tokio\_dbs
* tonic
* uuid
* warp
{% endcolumn %}
{% endcolumns %}

### Supported file formats

For Rust with Snyk Code, Snyk supports the following file formats: `.rs`

### Available features

* Reports
* Interfile analysis

## Rust for Snyk Open Source

### Supported package managers and registries

For Rust with Snyk Open Source, Snyk supports the following:

* Supported package registry: [crates.io](https://crates.io/)
* Supported files: CycloneDX and SPDX SBOMs

### Available features

* Test your SBOM containing `cargo` PURLs through the [SBOM test](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/sbom-test) CLI command or API
* Test your individual Rust packages using the [List issues for a package](https://docs.snyk.io/developer-tools/snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package) API

{% hint style="info" %}
SCM import and the standard CLI commands `snyk test` and `snyk monitor` are not available.
{% endhint %}

Open Source scanning of Rust manifests and dependencies is limited to testing using the CLI command `snyk sbom test`, or through the Snyk API, using either SBOM testing or individual package testing. API access is available only with Enterprise plans.

{% hint style="info" %}
License scanning is not available for Cargo.
{% endhint %}

### Scan Rust dependencies in bulk using the CLI

To scan Rust dependencies in bulk using the CLI:

1. Use a third-party tool to create an SBOM document from the `Cargo.toml` and `Cargo.lock` files, in one of the supported SBOM formats.
2. Use `snyk sbom test --file=<path to your SBOM>` to scan the SBOM document.

Alternatively, you can use the REST API, as follows:

1. Use the REST API to `POST` the SBOM document to the `sbom_tests` endpoint.
2. Retrieve the results by calling the `sbom_tests/{job_id}` endpoint. For more information, visit [SBOM](https://docs.snyk.io/developer-tools/snyk-api/reference/sbom) and [Test an SBOM document for vulnerabilities](https://docs.snyk.io/developer-tools/snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities).

### Scan Rust dependencies individually using the API

To test your individual Rust packages from the Cargo package manager, you can use the [List issues for a package](https://docs.snyk.io/developer-tools/snyk-api/reference/issues) API. You can obtain the PURL from the metadata section of the package on [crates.io](http://crates.io). The PURL must adhere to the [purl specification](https://github.com/package-url/purl-spec).

Before you use the PURL in the API, URL-encode it. For example, `pkg:cargo/sd@0.1.0` becomes `pkg%3Acargo%2Fsd%400.1.0`.

This reports only the direct vulnerabilities for that package. For more information, visit [List issues for a package](https://docs.snyk.io/developer-tools/snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package).

### Gate commits and PRs

To avoid introducing vulnerabilities on commits and PRs, Snyk recommends incorporating the preceding testing into your CI/CD pipeline.
