---
nav_context: agnostic
description: >-
  The release and lifecycle statuses Snyk uses for features, what each status
  means for access, support, and documentation, how much notice Snyk gives
  before a lifecycle change, and how legacy statuses map to the current ones.
last_reviewed: 2026-11-20
last_reviewed_source: reviewed
---

# Snyk release process

{% hint style="info" %}
New major capabilities go through every release status. Smaller features can ship without going through the release statuses. Timelines vary by feature.
{% endhint %}

Since September 1, 2026, Snyk uses the following statuses for every feature that was not yet Generally Available on that date. A status that Snyk announced before September 1, 2026 still applies, so some features show a legacy status. For those, visit [Legacy release statuses](snyk-release-process.md#legacy-release-statuses).

## Release statuses

Each release status lets more customers use the feature, and each comes with a different level of support.

<table><thead><tr><th width="150">Status</th><th width="230">Description</th><th>Available to</th><th>Access</th><th>Support</th><th>Docs</th></tr></thead><tbody><tr><td>Prototype</td><td>Snyk builds and tests the feature internally.</td><td>Snyk internal users only</td><td>Not available outside Snyk</td><td>No support commitment</td><td>No documentation</td></tr><tr><td>Experimental Preview</td><td>Snyk tests the feature with a small number of design partners. This is the first customer-facing release. Snyk makes no commitments about the feature.</td><td>A small number of existing Snyk customers that the Snyk team building the feature selects</td><td>Invitation only. The Snyk team building the feature turns it on.</td><td>No support commitment. The Snyk team building the feature handles questions directly.</td><td>Provided to design partners for onboarding, not public</td></tr><tr><td>Open Preview</td><td>Snyk tests whether the feature delivers value to customers beyond the design partners. The feature has stated limits. Snyk can move the feature back to Experimental Preview or discontinue it.</td><td>Customers who opt in and meet the criteria Snyk sets for the feature, primarily existing Snyk customers</td><td>Opt-in: on request through your Snyk account team. Open Preview features are not priced by default. If Snyk sets pricing for a feature, the feature can require a purchase.</td><td>Best effort, within the stated limits of the feature</td><td>Public documentation</td></tr><tr><td>Generally Available</td><td>The feature is fully available. Snyk continues to improve it based on customer feedback.</td><td>All customers, subject to standard feature availability for their plan</td><td>Your Snyk account team turns on the feature. Some features require a purchase.</td><td>Full support, covered by the Snyk support SLA</td><td>Public documentation</td></tr></tbody></table>

## Lifecycle statuses

After a feature is available, it can move to one of the following statuses. Before the change, Snyk gives notice based on its impact. For details, visit [Notice periods](snyk-release-process.md#notice-periods).

<table><thead><tr><th width="150">Status</th><th width="230">Description</th><th>Available to</th><th>Access</th><th>Support</th><th>Docs</th></tr></thead><tbody><tr><td>End of support</td><td>The feature still works, but Snyk stops development on it.</td><td>Existing users only</td><td>Available to existing users</td><td>Snyk does not answer new support tickets</td><td>Public documentation, with the release status at the top of the page</td></tr><tr><td>Discontinued</td><td>Snyk removes the feature from the product.</td><td>No users</td><td>Not available</td><td>No support</td><td>Snyk removes the documentation</td></tr></tbody></table>

## Notice periods

Snyk announces a lifecycle change in the documentation before the change takes effect. This also applies to other changes that break existing workflows, integrations, or API behavior. Snyk can also notify affected customers by email, through in-product messages, on [updates.snyk.io](https://updates.snyk.io), or on [trust.snyk.io/updates](https://trust.snyk.io/updates).

The notice period starts on the date of the public announcement. Snyk sets the length of the notice period by the impact of the change:

<table><thead><tr><th width="150">Notice period</th><th>When it applies</th></tr></thead><tbody><tr><td>Six months</td><td>Snyk discontinues a feature, or the change requires you to migrate. Six months' notice also applies when the change affects a large share of customers or cannot be reversed.</td></tr><tr><td>Three months</td><td>The change requires you to act, has a documented migration path, and affects a limited set of customers.</td></tr><tr><td>One month</td><td>The change has a low impact: it affects few customers, has a simple or automated fix, requires no migration, and can be reversed.</td></tr></tbody></table>

In exceptional cases, the notice period can be shorter than one month.

## Brownouts

Before Snyk removes a feature or an API endpoint, Snyk can schedule brownouts. A brownout is a short, announced period when Snyk turns off the feature or endpoint. Brownouts help you find any dependencies on the feature before the removal date. Snyk runs brownouts near the end of the notice period.

## Legacy release statuses

Some features show a status that Snyk announced before September 1, 2026. Snyk does not relabel these features. Use this table to find the current status that matches each legacy status.

<table><thead><tr><th width="200">Before</th><th>After</th></tr></thead><tbody><tr><td>Alpha</td><td>Prototype</td></tr><tr><td>Closed Beta</td><td>Experimental Preview</td></tr><tr><td>Early Access</td><td>Open Preview</td></tr><tr><td>General Availability</td><td>Generally Available</td></tr><tr><td>Deprecated</td><td>End of support</td></tr><tr><td>End of Life</td><td>Discontinued</td></tr></tbody></table>
