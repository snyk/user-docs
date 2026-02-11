# Breakability risk levels

{% hint style="info" %}
**Release status**

Breakability analysis is in Early Access and available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://docs.snyk.io/snyk-admin/manage-settings/snyk-preview).
{% endhint %}

Snyk analyzes dependency upgrades to predict if a proposed change will break your build or application. Breakability analysis assigns a risk level to each upgrade to help you decide whether to auto-merge a fix or review it manually.

Snyk assesses the practical impact of a change rather than relying only on version numbers. For example, a major version update (v2.0) can appear risky based on the version number, even if the code change is trivial.

To determine risk levels, Snyk uses official release notes, change logs, and ecosystem data. This page outlines the logic and definitions Snyk uses to classify upgrades.

Snyk classifies risk based on two factors:

* Evidence: what the maintainers of the system say has changed in the system.
* Confidence: how certain Snyk is about that evidence.

{% hint style="info" %}
With breakability analysis, Snyk sends package information to a Large Language Model (LLM). This information includes the current and proposed upgrade versions. Snyk does not transmit application data. Future features involving application context require you to enable them separately.

Ensure to review AI-generated content for accuracy before use.
{% endhint %}

## Risk level definitions and actions

### Low risk

Based on standard usage patterns, Snyk believes this upgrade is unlikely to introduce breaking changes.

{% hint style="info" %}
You can most likely auto-merge or approve low-risk PRs.
{% endhint %}

Criteria for low risk assessment include:

* Explicit non-breaking evidence: Snyk retrieved release notes and verified they contain only bug fixes, security patches, or performance improvements.
* Minimal practical impact: the upgrade is a major version, but only drops support for an element that has been End-of-Life (EOL) for several years.
* Patch or minor upgrades: semantic versioning suggests compatibility, and the change log shows no anomalies.

### Medium risk (review recommended)

The upgrade is likely safe, but requires developer validation due to uncertainty or specific environmental constraints.

{% hint style="warning" %}
Do not auto-merge medium risk PRs. Review the release notes or run your test suite to confirm compatibility.
{% endhint %}

Criteria for medium-risk assessment include:

* Low confidence or missing data: official release notes or change logs are missing or are unparsable. Without evidence, Snyk defaults to caution.
* Environment flags: The upgrade drops support for an element recently EOL or still in wide use. You must verify deployment environment compatibility.
* Ambiguous changes: The change log mentions behavioral changes that are not strictly breaking, but have the potential to affect edge-case implementations.

### High risk (action required)

Snyk confirms or suspects that the upgrade introduces breaking changes that require code refactoring.

{% hint style="danger" %}
Do not merge a high risk PR without code changes and thorough testing. Snyk recommends treating it as a development task.
{% endhint %}

Criteria for high-risk assessment include:

* Confirmed API changes: release notes explicitly detail the removal or renaming of functions, classes, or methods used in standard implementations.
* Configuration changes: the upgrade requires changes to config files, middleware ordering, or build pipelines.
* Major runtime shifts: The package drops support for an active Long Term Support (LTS) runtime version.

## End-of-Life (EOL) threshold

Major version upgrades often occur because a maintainer stops supporting an old language version. Snyk evaluates this as:

* Low impact: if a package drops support for a runtime that has been EOL for more than two years. Snyk classifies this as low risk. Snyk assumes secure enterprise environments have migrated away from these legacy versions.
* High impact: if a package drops support for a runtime that is current, LTS, or recently EOL. Snyk classifies this as high risk because it can break CI/CD pipelines that have not been upgraded.

## Deprecations and removals

* Deprecations: in case of deprecations, a feature is marked as deprecated, but still functions in the code. Snyk considers this to be low risk.
* Removals: if a deprecated feature is removed from the codebase, Snyk considers this high risk.

## Examples

The following use cases show how the logic applies.

| Use case               | Version upgrade         | Snyk risk assessment | Rationale                                                                                                                            |
| ---------------------- | ----------------------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| Security patch         | From `1.4.4` to `1.4.5` | Low                  | The changelog confirms only security fixes and no behavioral changes.                                                                |
| Ancient runtime drop   | From `3.0` to `4.0`     | Low                  | The upgrade is a major version, but the only change is dropping Node.js v4 (EOL 2018). There is no practical impact.                 |
| Missing documentation  | From `2.1` to `2.2`     | Medium               | The upgrade is a minor version, but there is no changelog. Snyk cannot detemine if it is safe, so it flags it for review.            |
| API removal            | From `3.0` to `4.0`     | High                 | The changelog confirms that `_.pluck` was removed and must be replaced with `_.map`. Code changes are required.                      |
| Long term support drop | From `5.0` to `6.0`     | High                 | The upgrade drops support for Java 17, which is an active LTS. This upgrade breaks builds that are running on standard environments. |
|                        |                         |                      |                                                                                                                                      |
