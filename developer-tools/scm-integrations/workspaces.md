# Workspaces

{% hint style="info" %}
Workspaces improve the accuracy and reliability of the SCM integration results in Snyk, especially for large-scale enterprise environments. This capability also supports additional functionality and improvements we have planned in the future. For these reasons, Snyk strongly recommends [enabling this option](workspaces.md#manage-workspaces) by default at your Group level.
{% endhint %}

Workspaces represents a significant step forward in providing you with the most reliable and accurate vulnerability detection possible.

Traditionally, Snyk accessed repository contents using SCM APIs, which impose primary and secondary rate limits, and content limits. For example, the GitHub.com APIs are rate-limited to allow only a certain number of requests per hour, and there is a limit on the number of tree entries that can be retrieved from the Git database.

When Snyk retrieves repository contents over these APIs, these limitations inhibit Snyk from providing a complete analysis in a number of ways, especially across a large number of repositories, or for repositories containing more than 100,000 files, sometimes referred to as monorepos.

Workspaces removes these limitations.

Cloning improves the accuracy of results in a number of ways. Because Snyk can access a complete view of a source code repository at a specific commit SHA, including repositories containing more than 100,000 files, the analysis is also more complete through cloning.

## Snyk data ingestion

When Workspaces is enabled, Snyk ingests, through configured SCM integrations, a temporary and shallow clone of repository contents at a given commit and all commit metadata (including the commit message, authors, and timestamp).

{% hint style="info" %}
For more information on Snyk data processing and safeguards concerning Git repo cloning, see [Cache retention period related to vulnerability source data](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-snyk-handles-your-data#cache-retention-period-related-to-vulnerability-source-data) and [Snyk integrations: workspaces](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-snyk-handles-your-data#snyk-integrations-workspaces).
{% endhint %}

## Git protocols used by Workspaces

Repositories are cloned using HTTPS. SSH-based clones are unavailable.

## Manage Workspaces

You manage the Workspaces feature through the **Integrations settings** page on the Group or Organization level. To do so, you must be a Snyk Group Admin or Snyk Organization Admin.

When Workspaces is enabled at the Group level, all Organizations in that Group have Workspaces enabled. The individual Organization level setting can override this Group level setting.

{% hint style="warning" %}
When an Organization level setting overrides a Group level setting, no future changes to the Group level setting can modify the Organization level setting.
{% endhint %}

While you can choose to disable the Workspaces feature, it is important to understand that doing so hinders the ability of Snyk to scan repositories in two specific ways:

1. Frequency of scanning: With Workspaces disabled, Snyk analyzes repository content through SCM APIs, which typically impose primary and secondary rate limits, and content limits. For example, the GitHub.com APIs are rate-limited to allow only a certain number of requests per hour, which hinders the ability of Snyk to analyze a large number of repositories in any one hour.
2. Completeness of scanning: With the Workspaces feature disabled, Snyk analyzes repo content through SCM APIs, which typically limit the number of tree entries that can be retrieved from the Git database, potentially leading to missed vulnerabilities. This impacts analysis of repositories containing a large number of files, sometimes referred to as monorepos.
