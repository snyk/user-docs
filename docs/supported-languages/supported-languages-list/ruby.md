# Ruby

{% hint style="info" %}
Snyk does not support Fix PRs for Projects using Ruby versions 3.1.x and lower. To avoid disruption, Snyk recommends upgrading to a supported Ruby version.
{% endhint %}

## Applicability and integration

{% hint style="info" %}
Ruby is supported for Snyk Code and Snyk Open Source.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

Available functions:

* Test your app's SBOM using `pkg:gem`
* Test your app's packages using `pkg:gem`

## Technical specifications

### Supported versions

Snyk supports the following Ruby versions:

| Ruby main version | Ruby specific version |
| ----------------- | --------------------- |
| `2.4.X`           | `2.4.0` to `2.4.10`   |
| `2.5.X`           | `2.5.0` to `2.5.9`    |
| `2.6.X`           | `2.6.1` to `2.6.10`   |
| `2.7.X`           | `2.7.0` to `2.7.8`    |
| `3.0.X`           | `3.0.0` to `3.0.7`    |
| `3.1.X`           | `3.1.0` to `3.1.7`    |
| `3.2.X`           | `3.2.0` to `3.2.9`    |
| `3.3.X`           | `3.3.0` to `3.3.9`    |
| `3.4.X`           | `3.4.0` to `3.4.5`    |

### Supported frameworks and libraries

For Ruby, the following frameworks and libraries are supported:

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

### Supported package managers and registries

For Ruby, Snyk supports [Bundler](https://bundler.io/) as a package manager. All Gemfile and Gemfile.lock are compatible with the Ruby versions that Snyk supports.

As a package registry, [rubygems.org](https://rubygems.org/) is supported.

## Ruby for Snyk Code

For Ruby with Snyk Code, the following file formats are supported:  `.erb`, `.haml`, `.rb`, `.rhtml`, `.slm`

Available features:

* Reports


## Ruby for Snyk Open Source

{% hint style="info" %}
Depending on your plan, some features may not be available. For more information, see [plans and pricing](https://snyk.io/plans/).

Platform-specific packages are not supported. If these are present in your `Gemfile.lock`, this can cause an invalid Fix PR to be created. If possible, use the non-platform-specific variant of a package.
{% endhint %}

For Ruby with Snyk Open Source, the following file formats are supported: `gemfile`, `gemfile.lock`

Available features:

* Fix PRs
* License scanning
* Reports

### Bundler support

Snyk supports testing, monitoring, and fixing Ruby Projects in the CLI and Git integrations that have their dependencies managed by [Bundler](https://bundler.io/) and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups. It is not possible to exclude certain groups, such as test or development groups.

### Manifest files supported for Ruby

For Ruby, Snyk supports the following manifest files:

* `Gemfile`
* `Gemfile.lock`

Snyk requires both files to be present to correctly test, monitor, and fix Ruby Projects.

### **Private Gem sources**

If your Gemfile needs access to private Gem sources, see [Private gem sources for Ruby configuration](../../scan-with-snyk/snyk-open-source/package-repository-integrations/private-gem-sources-for-ruby-configuration.md).

Using private Gem sources should work normally when you are using the Snyk CLI.

When creating Fix PRs for Ruby Projects using private Gem sources, Snyk may need access to the service hosting the Gems to update the file correctly.

### Fixing vulnerabilities in your Ruby Projects

{% hint style="info" %}
For Ruby versions < 3.2, Snyk does not support pinning a specific version of Ruby in the Gemfile, for example, `ruby "2.7.7".` You must use a more permissive version range that encapsulates all point versions, such as`ruby "~> 2.7.x"`
{% endhint %}

Snyk can fix vulnerabilities by updating vulnerable gems using `bundle update` after modifying your Gemfile, adhering to the rules you have specified there as far as possible.

In some scenarios, Snyk cannot upgrade all dependencies to non-vulnerable versions. In this case, consider updating the rules in your Gemfile.
