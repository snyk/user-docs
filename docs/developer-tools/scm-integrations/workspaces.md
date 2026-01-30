# Workspaces

{% hint style="info" %}
Workspaces improve the accuracy and reliability of Snyk’s SCM integration results, especially for large-scale enterprise environments. This capability also supports additional functionality and improvements we have planned in the future. For these reasons, Snyk strongly recommends [enabling this option](workspaces.md#manage-workspaces) by default at your Group level.
{% endhint %}

Workspaces represents a significant step forward in providing you with the most reliable and accurate vulnerability detection possible.

Traditionally, Snyk has accessed repository contents using SCM APIs, which impose primary and secondary rate limits, and content limits. For example, the GitHub.com APIs are rate-limited to allow only a certain number of requests per hour, and there is a limit on the number of tree entries that can be retrieved from the Git database.

When repository contents are retrieved over these APIs, these limitations inhibit Snyk’s providing a complete analysis in a number of ways, especially across a very large number of repositories, or for repositories containing more than 100,000 files, sometimes referred to as monorepos.

Through Workspaces, these limitations are removed.

The accuracy of results is improved in a number of ways through cloning. Since Snyk can access a complete view of a source code repository at a specific commit SHA, including repositories containing more than 100,000 files, the analysis is also more complete through cloning.

## Snyk data ingestion

When Workspaces is enabled, Snyk will ingest, through configured SCM integrations, a temporary and shallow clone of repository contents at a given commit and all commit metadata (including the commit message, authors, and timestamp).

{% hint style="info" %}
For more information on Snyk data processing and safeguards concerning Git repo cloning, see [Cache retention period related to vulnerability source data](../../snyk-data-and-governance/how-snyk-handles-your-data.md#cache-retention-period-related-to-vulnerability-source-data) and [Snyk integrations: workspaces](../../snyk-data-and-governance/how-snyk-handles-your-data.md#snyk-integrations-workspaces).
{% endhint %}

## Git protocols used by Workspaces

Repositories are cloned using HTTPS. SSH-based clones are unavailable.

## Manage Workspaces

The Workspaces feature is managed through the **Integrations settings** page on the Group or Organization level. To do so, you must be a Snyk Group Admin or Snyk Organization Admin.

When Workspaces is enabled at the Group level, all Organizations within that Group have Workspaces enabled. This Group level setting can be overridden by the individual Organization level setting.

{% hint style="warning" %}
When a Group level setting is overridden by an Organization level setting, the Organization level setting cannot be modified by any future changes made to the Group level setting.
{% endhint %}

While you can choose to disable the Workspaces feature, it is important to understand doing so will hinder Snyk’s ability to scan repositories in two specific ways:

1. Frequency of scanning: With Workspaces disabled, Snyk will analyze repository content through SCM APIs, which typically impose primary and secondary rate limits, and content limits. For example, the GitHub.com APIs are rate-limited to allow only a certain number of requests per hour, which hinders Snyk’s ability to analyze a large number of repositories in any one hour.
2. Completeness of scanning: With the Workspaces feature disabled, Snyk will analyze repo content through SCM APIs, which typically limit the number of tree entries that can be retrieved from the Git database, potentially leading to missed vulnerabilities. This impacts analysis of repositories containing a large number of files, sometimes referred to as monorepos.
