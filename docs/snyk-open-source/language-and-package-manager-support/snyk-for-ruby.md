# Snyk for Ruby

Snyk supports testing, monitoring and fixing Ruby projects in the CLI and Git integrations that have their dependencies managed by [Bundler](https://bundler.io/), and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups, and currently it is not possible to exclude certain groups \(such as test or development groups\).

If your Gemfile needs access to private Gem sources please [see below.](snyk-for-ruby.md)

The following manifest files are supported:

* `Gemfile`
* `Gemfile.lock`

{% hint style="info" %}
**Note**  
Snyk requires both files to be present in order to correctly test, monitor & fix Ruby projects
{% endhint %}

## Fixing vulnerabilities in your Ruby projects

Snyk can fix vulnerabilities by updating vulnerable gems, using bundle update, after modifying your Gemfile \(sticking to the rules you have specified there as far as possible\).

This means that in some scenarios we won’t be able to upgrade all dependencies to non-vulnerable versions. In this case, you should consider updating the rules in your Gemfile.

In future releases, we are planning to provide suggestions to make this easier.

## **Private Gem Sources**

If you use private Gem sources this should work as normal when testing via the Snyk CLI.

You must take additional steps to [configure private Gem sources for projects imported from Git](https://docs.snyk.io/integrations/private-registry-integrations/private-gem-sources-for-ruby).



