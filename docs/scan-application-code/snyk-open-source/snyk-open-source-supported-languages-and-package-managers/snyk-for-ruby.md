# Snyk for Ruby

Snyk supports testing, monitoring, and fixing Ruby projects in the [CLI ](../../../snyk-cli/)and Git [integrations](../../../integrations/) that have their dependencies managed by [Bundler](https://bundler.io/), and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups. Currently it is not possible to exclude certain groups, such as test or development groups.

If your Gemfile needs access to private Gem sources, see [Private Gem sources](snyk-for-ruby.md#private-gem-sources).

{% hint style="warning" %}
Platform-specific packages are currently not supported. If these are present in your Gemfile.lock then this can cause an invalid **Fix PR** to be created. If possible, use the non-platform-specific variant of a package.
{% endhint %}

## Manifest files supported with Snyk for Ruby

The following manifest files are supported:

* `Gemfile`
* `Gemfile.lock`

{% hint style="info" %}
Snyk requires both files to be present in order to correctly test, monitor , and fix Ruby projects
{% endhint %}

## Fixing vulnerabilities in your Ruby projects

Snyk can fix vulnerabilities by updating vulnerable gems using bundle update, after modifying your Gemfile, adhering to the rules you have specified there as far as possible.

This means that in some scenarios Snyk is not able to upgrade all dependencies to non-vulnerable versions. In this case, consider updating the rules in your Gemfile.

In future releases, Snyk plans to provide suggestions to make this easier.

## **Private Gem sources**

If you use private Gem sources this should work as normal when you are testing using the Snyk CLI.

You must take additional steps to configure private Gem sources for projects imported from Git.

## Supported Ruby version

#### `2.3.X`

* '2.3.1', '2.3.6'

#### `2.4.X`

* '2.4.0', '2.4.1', '2.4.2', '2.4.5', '2.4.6', '2.4.9'

#### `2.5.X`

* '2.5.0', '2.5.1', '2.5.3'

#### `2.6.X`

* '2.6.1', '2.6.3', '2.6.5', '2.6.6'

#### `2.7.X`

* '2.7.2', '2.7.3', '2.7.4', '2.7.5', '2.7.6'

#### `3.0.X`

* '3.0.0'

#### `3.1.X`

* '3.1.0'

#### `3.2.X`

* '3.2.0'
