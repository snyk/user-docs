# Ruby for open source

## Ruby for open source support

Refer to the [Ruby details](./) for supported package managers and features.

## Open source and licensing

{% hint style="info" %}
**Feature availability**\
Depending on your plan, some features may not be available. For more information, see  [plans and pricing](https://snyk.io/plans/).

Platform-specific packages are not supported. If these are present in your `Gemfile.lock`, this can cause an invalid Fix PR to be created. If possible, use the non-platform-specific variant of a package
{% endhint %}

## Bundler support

Snyk supports testing, monitoring, and fixing Ruby Projects in the CLI and Git integrations that have their dependencies managed by [Bundler](https://bundler.io/) and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups. It is not possible to exclude certain groups, such as test or development groups.

## Manifest files supported with Snyk for Ruby

The following manifest files are supported:

* `Gemfile`
* `Gemfile.lock`

Snyk requires both files to be present to correctly test, monitor, and fix Ruby Projects.

## **Private Gem sources**

If your Gemfile needs access to private Gem sources, see [Private gem sources for Ruby configuration](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/private-gem-sources-for-ruby-configuration.md).

Using private Gem sources should work normally when you are using the Snyk CLI.

When creating Fix PRs for Ruby Projects using private Gem sources, Snyk may need access to the service hosting the Gems to update the file correctly.

## Fixing vulnerabilities in your Ruby Projects

{% hint style="info" %}
For Ruby versions < 3.2, Snyk does not support pinning a specific version of Ruby in the Gemfile, for example, `ruby "2.7.7".` You must use a more permissive version range that encapsulates all point versions, such as`ruby "~> 2.7.x"`
{% endhint %}

Snyk can fix vulnerabilities by updating vulnerable gems using `bundle update`after modifying your Gemfile, adhering to the rules you have specified there as far as possible.

In some scenarios, Snyk cannot upgrade all dependencies to non-vulnerable versions. In this case, consider updating the rules in your Gemfile.
