---
nav_context: agnostic
description: >-
  Snyk support for Ruby with Snyk Code and Snyk Open Source, including SCM
  import, CLI and IDE testing, and supported versions
last_reviewed: 2026-09-14
last_reviewed_source: focused-edit
---

# Ruby

## Applicability and integration

{% hint style="info" %}
Snyk supports Ruby for Snyk Code and Snyk Open Source.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

## Technical specifications

### Supported versions

Snyk supports the following Ruby versions:

| Ruby main version | Ruby specific version |
| ----------------- | --------------------- |
| `2.4.X`           | `2.4.0` to `2.4.10`   |
| `2.5.X`           | `2.5.0` to `2.5.9`    |
| `2.6.X`           | `2.6.0` to `2.6.10`   |
| `2.7.X`           | `2.7.0` to `2.7.8`    |
| `3.0.X`           | `3.0.0` to `3.0.7`    |
| `3.1.X`           | `3.1.0` to `3.1.7`    |
| `3.2.X`           | `3.2.0` to `3.2.9`    |
| `3.3.X`           | `3.3.0` to `3.3.9`    |
| `3.4.X`           | `3.4.0` to `3.4.5`    |

### Supported frameworks and libraries

For Ruby, Snyk supports the following frameworks and libraries:

* ActiveRecord
* Connection
* grpc-ruby
* LibXML
* mysql2
* Nokogiri
* OpenSSL
* openai ruby client
* ruby-openai
* rexml
* Ruby on Rails
* sinatra
* sqlite3-ruby

### Supported package managers and registries

For Ruby, Snyk supports [Bundler](https://bundler.io/) as a package manager. All Gemfile and Gemfile.lock files are compatible with the Ruby versions that Snyk supports.

Snyk supports [rubygems.org](https://rubygems.org/) as a package registry.

## Ruby for Snyk Code

For an overview of the supported security rules, visit [Ruby rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/ruby-rules).

For Ruby with Snyk Code, Snyk supports the following file formats: `.erb`, `.haml`, `.rb`, `.rhtml`, `.slm`

Available features:

* Reports
* Interfile analysis

## Ruby for Snyk Open Source

{% hint style="info" %}
Depending on your plan, some features may not be available. For more information, visit [plans and pricing](https://snyk.io/plans/).

Snyk does not support platform-specific packages. If these are present in your `Gemfile.lock`, Snyk can create an invalid Fix PR. If possible, use the non-platform-specific variant of a package.
{% endhint %}

For Ruby with Snyk Open Source, Snyk supports the following file formats: `Gemfile`, `Gemfile.lock`

Available features:

* Fix PRs
* License scanning
* Test your app's SBOM and packages using `pkg:gem` PURLs through the [SBOM test](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/sbom-test) CLI command

### Bundler support

For Ruby Projects that use [Bundler](https://bundler.io/) to manage dependencies, Snyk supports testing, monitoring, and fixing in the CLI and SCM integrations. Snyk compares the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups. You cannot exclude specific groups, such as test or development groups.

### Manifest files supported for Ruby

For Ruby, Snyk supports the following manifest files:

* `Gemfile`
* `Gemfile.lock`

Snyk requires both files to be present to correctly test, monitor, and fix Ruby Projects.

### Private gem sources

If your Gemfile needs access to private gem sources, visit [Private gem sources for Ruby configuration](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations/private-gem-sources-for-ruby-configuration).

Using private gem sources should work normally when you are using the Snyk CLI.

When creating Fix PRs for Ruby Projects using private gem sources, Snyk may need access to the service hosting the gems to update the file correctly.

### Fixing vulnerabilities in your Ruby Projects

{% hint style="info" %}
For Ruby versions earlier than 3.2, Snyk does not support pinning a specific version of Ruby in the Gemfile, for example, `ruby "2.7.7"`. You must use a more permissive version range that encapsulates all point versions, such as `ruby "~> 2.7.x"`.
{% endhint %}

Snyk can fix vulnerabilities by updating vulnerable gems using `bundle update` after modifying your Gemfile, adhering to the rules you have specified there as far as possible.

In some scenarios, Snyk cannot upgrade all dependencies to non-vulnerable versions. In this case, update the rules in your Gemfile.
