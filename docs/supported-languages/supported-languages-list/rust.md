# Rust

{% hint style="info" %}
Rust is supported for Snyk Code and for Snyk Open Source.
{% endhint %}

{% hint style="info" %}
Code analysis support for Rust is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](../../snyk-platform-administration/snyk-preview.md).
{% endhint %}

## Rust for Snyk Code

For an overview of the supported security rules, visit [Rust rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/rust-rules.md).

### Supported frameworks and libraries

The following frameworks and libraries are supported:

{% columns %}
{% column %}
* actix_files
* actix_identity
* actix_multipart
* actix_session
* actix_web
* age
* ammonia
* async_graphql
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
* tokio_dbs
* tonic
* uuid
* warp
{% endcolumn %}
{% endcolumns %}

### Supported file formats

The following file formats are supported: `.rs`.

### Available features

* Support for Interfile analysis
* Reports

## Rust for Snyk Open Source

For Rust with Snyk Open Source, there are no features supported.

{% hint style="warning" %}
If you want to use Open Source PR checks in Rust Projects for which the **Fix PRs** feature is not supported, it is possible that vulnerable dependencies are introduced without being flagged.
{% endhint %}

## Applicability and integration

Available functions:

* Test your app's SBOM using `pkg:cargo` For more information, see [Test an SBOM document for vulnerabilities](../../snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).
* Test your app's packages using `pkg:cargo`
* Test your individual Rust packages from the Cargo package manager. For more information, see [List issues for a package](../../snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package.md).

## Technical specifications

### Supported package managers and registries

* Supported package manager: Cargo
* Supported package registry: [crates.io](https://crates.io/)