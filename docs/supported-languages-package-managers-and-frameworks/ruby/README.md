# Ruby

## Applicability

Snyk supports [Ruby for code analysis](ruby-for-code-analysis.md) and [Ruby for open source](ruby-for-open-source.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:gem`
* Test your app's packages using `pkg:gem`

## Package managers and supported file extensions

Snyk for Ruby supports Bundler as a package manager and all Gemfile and Gemfile.lock are compatible with the Snyk supported Ruby versions.

As a package registry, [rubygems.org](https://rubygems.org/) is supported.

**Supported Ruby versions**

| Ruby main version | Ruby specific version                                         |
| ----------------- | ------------------------------------------------------------- |
| **`2.3.X`**       | `2.3.1`, `2.3.6`                                              |
| **`2.4.X`**       | `2.4.0`, `2.4.1`, `2.4.2`, `2.4.5`, `2.4.6`, `2.4.9`          |
| **`2.5.X`**       | `2.5.0`, `2.5.1`, `2.5.3`                                     |
| **`2.6.X`**       | `2.6.1`, `2.6.3`, `2.6.5`, `2.6.6`                            |
| **`2.7.X`**       | `2.7.2`, `2.7.3`, `2.7.4`, `2.7.5`, `2.7.6`, `2.7.7`, `2.7.8` |
| **`3.0.X`**       | `3.0.0`                                                       |
| **`3.1.X`**       | `3.1.0`, `3.1.1`, `3.1.2`, `3.1.3`                            |
| **`3.2.X`**       | `3.2.0`, `3.2.1`                                              |

Snyk for Ruby supports the following file formats:

* Snyk Open Source: `gemfile`, `gemfile.lock`
* Snyk Code: `.erb`, `.haml`, `.rb`, `.rhtml`, `.slm`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Ruby:

* ActiveRecord - Partial
* Connection - Comprehensive
* grpc-ruby - Comprehensive
* LibXML - Comprehensive
* mysql2 - Comprehensive
* Nokogiri - Comprehensive
* OpenSSL - Comprehensive
* openai ruby client - Comprehensive
* ruby-openai - Comprehensive
* rexml - Comprehensive
* Ruby On Rails - Comprehensive
* sinatra - Comprehensive
* sqlite3-ruby - Comprehensive

## Features

The following features are supported in Snyk for Ruby:

| Snyk Open Source                                                   | Snyk Code                                      |
| ------------------------------------------------------------------ | ---------------------------------------------- |
| <ul><li>Fix PRs</li><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li></ul> |
