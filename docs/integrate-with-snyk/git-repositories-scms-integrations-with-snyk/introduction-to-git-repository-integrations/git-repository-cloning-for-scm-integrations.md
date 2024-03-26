# Git repository cloning for SCM integrations

{% hint style="info" %}
**Feature availability**\
This feature is in early access for GitHub, GitHub Enterprise, GitLab, Bitbucket Server, Bitbucket Cloud App, Bitbucket Cloud (Legacy), and Azure Repos integrations.

To enable this feature, you must use Snyk Preview. See [Enable Git repository cloning](../../../snyk-admin/snyk-preview.md#enable-git-repository-cloning).
{% endhint %}

Full Git repository cloning allows Snyk to provide more reliable and accurate results when scanning your source code through the [SCM integrations](https://docs.snyk.io/integrations/git-repository-scm-integrations), helping you develop fast and stay secure.

## How Git cloning supports more reliable results

Traditionally, Snyk has accessed repository contents using SCM APIs, which impose primary and secondary rate limits, and content limits. For example, the GitHub.com APIs are rate-limited to allow only a certain number of requests per hour, and there is a limit on the number of tree entries that can be retrieved from the Git database.

When repository contents are retrieved over these APIs, these limitations inhibit Snyk’s providing a complete analysis in a number of ways, especially across a very large number of repositories, or for repositories containing more than 100,000 files, sometimes referred to as monorepos.

Through cloning, these limitations are removed.

## How Git cloning supports more accurate results

The accuracy of results is improved in a number of ways through cloning. Since Snyk can access a complete view of a source code repository at a specific commit SHA, including repositories containing more than 100,000 files, the analysis is also more complete through cloning.

## Snyk data ingestion

When Git repository cloning is enabled, Snyk will ingest, through configured SCM integrations, a temporary snapshot of repository contents at a given commit and all commit metadata (including the commit message, authors, and timestamp).

{% hint style="info" %}
For more information on Snyk data processing and safeguards concerning Git repo cloning, see [Cache retention period related to vulnerability source data](../../../working-with-snyk/how-snyk-handles-your-data.md#cache-retention-period-related-to-vulnerability-source-data) and [Safeguards Snyk puts in place to ensure data is secure](../../../working-with-snyk/how-snyk-handles-your-data.md#safeguards-snyk-puts-in-place-to-ensure-data-is-secure).
{% endhint %}

{% hint style="warning" %}
By enabling this feature, you agree that your Git repository is a [Protected Asset](../../../working-with-snyk/how-snyk-handles-your-data.md#git-cloning-applicable-contract-terms) as defined in the contract between your company and Snyk.
{% endhint %}

## Git repository cloning protocols

Repositories are cloned using HTTPS. SSH-based clones are currently unavailable.

## Flows used in Git repository cloning

[PR checks](../../../scan-with-snyk/pull-requests/pull-request-checks/), [import](../../../getting-started/quickstart/import-a-project.md), and [recurring tests](https://docs.snyk.io/scan-with-snyk/working-with-snyk-in-your-environment/running-scans) will all make use of Git repository cloning for all Projects imported using an SCM integration.

## Snyk Broker interactions

Brokered connections are supported when Git operations are allowed through Broker.

{% hint style="warning" %}
This will override restrictions from `accept.json`. For more information, see [Clone capability with Broker for Docker](../../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md).
{% endhint %}
