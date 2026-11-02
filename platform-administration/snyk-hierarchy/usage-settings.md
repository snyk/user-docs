---
nav_context: classic
description: >-
  How to view Snyk usage details for your Group or Organization, including tests
  and contributing developers
---

# SKU-based licensing usage settings

{% hint style="info" %}
The following usage settings are relevant only for subscriptions to the legacy, SKU-based licensing offerings for Snyk Open Source, Snyk Code, Snyk Container, and Snyk IaC. These settings do not apply to any capabilities entitled through the [Snyk Platform Subscription](../snyk-platform-subscription/snyk-platform-subscription-mechanics.md).
{% endhint %}

In your Group or Organization, select **Settings** > **Usage** to view Snyk usage details for your Group or Organization:

* **Test usage**: the number of tests used
* **Contributing developers**: the number of developers contributing to Projects

{% hint style="info" %}
For more information about test frequency settings, see [Project actions](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects#project-actions-on-the-project-listing-page) on the Project Listings page, where bulk actions are also explained.
{% endhint %}

## Test usage

The **Test Usage** section shows how many tests you are using over the current billing period:

<figure><img src="../.gitbook/assets/usage_test_usage.png" alt="Test usage data"><figcaption><p>Test usage data</p></figcaption></figure>

{% hint style="info" %}
Test limits vary for Snyk products and plans.

For the Free plan, the limits are: Open Source, 400 tests; Code, 100 tests; IaC, 300 tests; Container, 100 tests.

To see the limits that apply to you and how much of them you have used, open **Usage** in your Organization settings. On an Enterprise plan, your limits are set by your contract rather than published here; your Snyk account team can confirm them. See the [Plans and pricing page](https://snyk.io/plans/) for more details. See [What counts as a test?](https://docs.snyk.io/snyk-data-and-governance/what-counts-as-a-test) for details of how Snyk counts tests.
{% endhint %}

## Contributing developers

Snyk defines contributing developers as developers having made a commit to a private repo monitored by Snyk in the last 90 days.

The **Contributing developers for SCM and CLI integrations** section shows contributing developer counts, both at the Organization level and the Group level.

The developer count is the number of developers who have contributed to the default branch of any private repository connected to an integration.

Snyk does not count contributions to public (open-source) repos because the pricing model is based on the number of contributing developers to private repositories.

The contributing developers count shows:

* **Total unique contributors across all integrations:** the number of contributors across all the integrations in your Snyk account. Contributing developers are only counted once, even if they have contributed to multiple integrations or multiple repositories.
* **Breakdown by integration**: the number of contributors, Organizations, and repos in that integration.

A contributor is recognized through the author email field, set in the local Git settings.
