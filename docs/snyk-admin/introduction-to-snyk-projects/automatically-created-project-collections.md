# Automatically created Project collections

{% hint style="info" %}
**Release status**&#x20;

Project collections in [Closed Beta](../../getting-started/snyk-release-process.md) and available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

Scanning the same Project outside of Snyk through an SCM integration and the Snyk CLI creates duplicate Targets within Snyk, with duplicate Projects and Issues. These may not be exact duplicates.&#x20;

The Projects in these duplicate Targets will be detected and grouped automatically into a new Project Collection. You can filter Project Collections of this kind by the method used to scan the code.

You can create reports from these automatically created Project Collections so you see reports only for issues found by the method of scanning you specify.
