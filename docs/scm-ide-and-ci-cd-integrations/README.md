# SCM, IDE, and CI/CD integrations

This section of the documentation provides information about [Snyk SCM](snyk-scm-integrations/), [IDE](snyk-ide-plugins-and-extensions/), and [CI/CD integrations](snyk-ci-cd-integrations/).

Snyk supports SCM, IDE, and CI/CD integration methods that allow you to implement security at each point in your workflow: importing a Project, writing your code, and building and deployment.

{% hint style="warning" %}
Enterprise plan users have access to all of the functionality. The API and Snyk AppRisk are not available to Free and Team plan users. See [Plans and pricing](https://snyk.io/plans/) for more information.
{% endhint %}

There are two ways of implementing SCM integrations in a Snyk environment:

* **Group level** - At the Group level, you can set up the SCM integrations for Snyk AppRisk.&#x20;
* **Organization level** - At the Organization level, you can set up the SCM integrations for all Snyk plans. See the [Manage your Integrations](../getting-started/snyk-web-ui.md) at the Organizational level page for more details.&#x20;

## Choose an Integration

If you are an Enterprise customer, see [Choose rollout integrations](../implement-snyk/team-implementation-guide/phase-1-discovery-and-planning/choose-rollout-integrations.md) in the Enterprise implementation guide for tips and considerations on import strategies, as well as context for which integrations suit your SDLC.

### GitHub vs GitHub Enterprise

As an **Enterprise plan user**, Snyk recommends using the GitHub Enterprise integration as it enables you to use a single GitHub service account personal access token (PAT) across your Snyk Organization rather than depending on a PAT for an individual user account. You can use this integration whether or not you have a GitHub Enterprise (GHE) license or subscription.

Another benefit to using the GitHub Enterprise integration is that you can choose to clone integration settings when you are creating new Snyk Organizations. This means you can use one GitHub Enterprise integration for all Organizations in your Snyk Group.

As a **Free or Team plan user**, Snyk recommends using the GitHub integration as it only requires a PAT for an individual user account, which should meet your needs at this level.

If you use the **self-hosted** GitHub Enterprise product, you must use the GitHub Enterprise integration.

{% hint style="info" %}
For detailed steps on migrating from GitHub to GitHub Enterprise, see [Migrate to GitHub Enterprise](snyk-scm-integrations/github.md#migrate-to-the-github-enterprise-integration).
{% endhint %}

## Pull Requests for Snyk integrations

Snyk can automatically create pull requests (PRs) on your behalf to upgrade your dependencies based on scan results. This is compatible with a variety of Snyk integrations. For more information, see [View and understand Snyk upgrade pull requests for integrations](broken-reference).
