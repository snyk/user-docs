# Rust

{% hint style="info" %}
Rust is supported for Snyk Code (full support) and for Snyk Open Source (limited support).
{% endhint %}

## Rust for Snyk Code

{% hint style="info" %}
Code analysis support for Rust is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](../../snyk-platform-administration/snyk-preview.md).
{% endhint %}

### Supported frameworks and libraries

For Rust with Snyk Code, the following frameworks and libraries are supported:

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

For an overview of the supported security rules, visit [Rust rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/rust-rules.md).

### Supported file formats

For Rust with Snyk Code, the following file formats are supported: `.rs`.

### Available features

* Support for Interfile analysis
* Reports

## Rust for Snyk Open Source

### Supported package managers and registries

For Rust with Snyk Open Source, the following are supported:

* Supported package registry: [crates.io](https://crates.io/)
* Supported files: CycloneDX and SPDX SBOMs

### Available features

* Test your SBOM containing `cargo` PURLs using the [SBOM test](https://docs.snyk.io/developer-tools/snyk-cli/commands/sbom-test) CLI command or API.
* Test your individual Rust packages using the [List issues for a package](../../snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package.md) API.

SCM import and the standard CLI commands `snyk test`, `snyk monitor` are not available.

Open Source scanning of Rust manifests and dependencies is limited to testing using the CLI command `snyk sbom test`, or through the Snyk API, using either SBOM testing or individual package testing. API access is available only with Ignite or Enterprise plans.

### Scan Rust dependencies in bulk using the CLI

To do this:&#x20;

* Use a third party tool to create a SBOM document from the `Cargo.toml` and `cargo.lock` file, in one of the supported SBOM formats.
* Use `snyk sbom test --file=<path to your SBOM>` to scan the SBOM document.&#x20;

Alternatively, you can use the REST API, as follows:

* Use the REST API to `POST` the SBOM document to the `sbom_tests` endpoint.&#x20;
* Retrieve the results by calling the `sbom_tests/{job_id}` endpoint. For more information, visit [SBOM](../../snyk-api/reference/sbom.md) and [Test an SBOM document for vulnerabilities](../../snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).

### Scan Rust dependencies individually using the API

To test your individual Rust packages from the Cargo package manager, you can use the [List issues for a package](../../snyk-api/reference/issues.md) API. You can obtain the PURL from the metadata section of the package on [crates.io](http://crates.io) and it must adhere to the [purl specification](https://github.com/package-url/purl-spec).

Before using it in the API, ensure you URL encode it. For example, `pkg:cargo/sd@0.1.0` becomes `pkg%3Acargo%2Fsd%400.1.0`

This reports only the direct vulnerabilities for that package. For more information, visit [List issues for a package](../../snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package.md).

### Gating commits and PRs

To avoid introducing vulnerabilities on commits and PRs, Snyk recommends incorporating the above testing into your CI/CD pipeline.
