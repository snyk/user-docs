---
description: How Snyk Workspaces improve the accuracy and reliability of SCM integration results in large enterprises
nav_context: agnostic
last_reviewed: 2026-11-05
last_reviewed_source: focused-edit
---

# Workspaces

Workspaces represents a significant step forward in providing you with the most reliable and accurate vulnerability detection possible.

{% hint style="info" %}
Workspaces improve the accuracy and reliability of Snyk’s SCM integration results, especially for large-scale enterprise environments. This capability also supports additional functionality and improvements we have planned in the future. For these reasons, Snyk strongly recommends [enabling this option](workspaces.md#manage-workspaces) by default at your Group level.
{% endhint %}

Traditionally, Snyk has accessed repository contents using SCM APIs, which impose primary and secondary rate limits, and content limits. For example, the GitHub.com APIs are rate-limited to allow only a certain number of requests per hour, and there is a limit on the number of tree entries that can be retrieved from the Git database.

When repository contents are retrieved over these APIs, these limitations inhibit Snyk’s providing a complete analysis in a number of ways, especially across a very large number of repositories, or for repositories containing more than 100,000 files, sometimes referred to as monorepos.

Through Workspaces, these limitations are removed.

The accuracy of results is improved in a number of ways through cloning. Since Snyk can access a complete view of a source code repository at a specific commit SHA, including repositories containing more than 100,000 files, the analysis is also more complete through cloning.

## Snyk data ingestion

When Workspaces is enabled, Snyk will ingest, through configured SCM integrations, a temporary and shallow clone of repository contents at a given commit and all commit metadata (including the commit message, authors, and timestamp).

{% hint style="info" %}
For more information on Snyk data processing and safeguards concerning Git repo cloning, see [Cache retention period related to vulnerability source data](https://docs.snyk.io/snyk-data-and-governance/how-snyk-handles-your-data#cache-retention-period-related-to-vulnerability-source-data) and [Snyk integrations: workspaces](https://docs.snyk.io/snyk-data-and-governance/how-snyk-handles-your-data#snyk-integrations-workspaces).
{% endhint %}

## Git protocols used by Workspaces

Repositories are cloned using HTTPS. SSH-based clones are unavailable.

## Manage Workspaces

You manage Workspaces in the **Workspaces** section of the **Integration Settings** page, at the Group or Organization level. You must be a Snyk Group Admin or Snyk Organization Admin. To turn on Workspaces, click **Enable for group** or **Enable for organization**. After you turn on Workspaces, you cannot turn it off.

If your Group or Organization turned off Workspaces in the past, Workspaces stays off until you turn it on.

When you turn on Workspaces at the Group level, every Organization in the Group inherits the setting, unless the Organization has its own Workspaces setting.

{% hint style="warning" %}
An Organization-level setting always takes precedence over the Group-level setting. If an Organization turned off Workspaces in the past, Workspaces stays off for that Organization even after you turn it on for the Group. To turn on Workspaces for that Organization, turn it on at the Organization level.
{% endhint %}
