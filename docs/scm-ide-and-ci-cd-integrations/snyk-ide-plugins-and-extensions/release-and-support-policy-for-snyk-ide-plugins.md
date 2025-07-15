# Release and support policy for Snyk IDE plugins

This page details the release policy for Snyk IDE plugins.

Builds are deployed at distinct points in time after comprehensive testing and are considered stable.

**Cadence**: Six versions per year, approximately every **eight weeks**. Hotfix releases are possible for newly discovered security vulnerabilities.

{% hint style="info" %}
Snyk recommends always using the latest version of an IDE plugin. You can install through the marketplace or the GitHub release page at any time. Each installation uninstalls the prior one. See the list of links for the [GitHub release pages and changelogs for the IDE plugins](release-and-support-policy-for-snyk-ide-plugins.md#github-release-pages-and-changelogs-for-the-ide-plugins) at the end of this page.
{% endhint %}

## Versioning for IDE plugins

All Snyk IDE plugins follow industry standard [Semantic Versioning](https://semver.org/), which has a three-part notation. Given a version number, `MAJOR.MINOR.PATCH`, increment the:

1. `MAJOR` version when you make **breaking** changes; examples follow.
2. `MINOR` version when you add functionality in a backward-compatible manner.
3. `PATCH` version when you make backward-compatible bug fixes.

Examples of **breaking** changes include the following:

* Deprecating features that produced findings before, which means users will get fewer findings after the release.
* Introducing new features that might introduce new findings.
* Introducing mandatory configuration changes that might affect the number of findings.

## Support policy

Snyk supports the latest 12 months of plugin versions, ensuring functionality and performance. Older versions are considered End-of-Support (EOS) and will not receive bug fixes or troubleshooting.

Snyk provides fixes only in new versions and cannot fix older versions. Customers must upgrade to benefit from improvements.

This policy fosters innovation while optimizing resources.

If you need help, submit a [request](https://support.snyk.io) to Snyk Support.

## Preview builds

Snyk offers preview versions for those who want to test in-progress features.

{% hint style="info" %}
These versions may contain bugs and are not officially supported.
{% endhint %}

Preview builds are deployed regularly, up to multiple times per day, and contain the latest changes. They are given a version number following the pattern { YEAR }. { MONTH }.{ DAY }{ HOUR } .

Preview versions of IDE plugins are deployed as standalone plugins named in the form `(Preview) Snyk Security`. Snyk recommends that you not install them together with the stable version.

## GitHub release pages and changelogs for the IDE plugins

### Releases

Release notes identify all changes, focusing on new features or significant changes.

* [Eclipse plugin releases](https://github.com/snyk/snyk-eclipse-plugin/releases)
* [JetBrains plugin releases](https://github.com/snyk/snyk-intellij-plugin/releases)
* [Visual Studio extension releases](https://github.com/snyk/snyk-visual-studio-plugin/releases)
* [Visual Studio Code extension releases](https://github.com/snyk/vscode-extension/releases)

### Changelogs

* [Eclipse plugin changelog](https://github.com/snyk/snyk-eclipse-plugin/blob/main/CHANGELOG.md)
* [JetBrains plugin changelog](https://github.com/snyk/snyk-intellij-plugin/blob/master/CHANGELOG.md)
* [Visual Studio extension changelog](https://github.com/snyk/snyk-visual-studio-plugin/blob/main/CHANGELOG.md)
* [Visual Studio extension changelog](https://github.com/snyk/snyk-visual-studio-plugin/blob/main/CHANGELOG.md)

