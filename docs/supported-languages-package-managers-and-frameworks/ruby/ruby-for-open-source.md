# Ruby for open source

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## [Ruby](./) for open source

**Package manager**: Bundler

**Package registry**: [rubygems.org](https://rubygems.org/)

**Import your app through SCM**: Available&#x20;

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:gem`

**Test your app's packages**: Available, `pkg:gem`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

**Package manager versions**: All Gemfile and Gemfile.lock are compatible with the [Snyk supported Ruby versions](./#supported-ruby-versions)

## Open source and licensing

{% hint style="info" %}
**Feature availability**\
Features may not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk supports testing, monitoring, and fixing Ruby Projects in the [CLI ](ruby-for-open-source.md#snyk-cli)and Git [integrations](ruby-for-open-source.md#snyk-integrations) that have their dependencies managed by [Bundler](https://bundler.io/) and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups. Currently, it is not possible to exclude certain groups, such as test or development groups.

If your Gemfile needs access to private Gem sources, see [Private Gem sources](ruby-for-open-source.md#private-gem-sources).

{% hint style="warning" %}
Platform-specific packages are currently not supported. If these are present in your `Gemfile.lock`, this can cause an invalid **Fix PR** to be created. If possible, use the non-platform-specific variant of a package.
{% endhint %}

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).
