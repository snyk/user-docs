# SCM, IDE, and CI/CD integrations

This section of the documentation provides information about [Snyk SCM](snyk-scm-integrations/), [IDE](snyk-ide-plugins-and-extensions/), and [CI/CD integrations](snyk-ci-cd-integrations/).

Snyk supports SCM, IDE, and CI/CD integration methods that allow you to implement security at each point in your workflow: importing a Project, writing your code, and building and deployment.

{% hint style="info" %}
**Feature availability**

The API and Snyk Essentials are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

There are two ways of implementing SCM integrations in a Snyk environment:

* Group level - At the Group level, you can set up the SCM integrations for Snyk Essentials.&#x20;
* Organization level - At the Organization level, you can set up the SCM integrations for all other Snyk products and all Snyk plans. See the [Manage your Integrations](../getting-started/snyk-web-ui.md) at the Organizational level page for more details.&#x20;

{% hint style="info" %}
If you want to use the same SCM integration at both Group and Organization levels, then you must set up that integration for both levels.
{% endhint %}

## Choose an Integration

If you are an Enterprise customer, see [Choose rollout integrations](../implement-snyk/team-implementation-guide/phase-1-discovery-and-planning/choose-rollout-integrations.md) in the Enterprise implementation guide for tips and considerations on import strategies, as well as context for which integrations suit your SDLC.

### GitHub vs GitHub Enterprise

As an Enterprise plan user, Snyk recommends using the GitHub Enterprise integration as it enables you to use a single GitHub service account personal access token (PAT) across your Snyk Organization rather than depending on a PAT for an individual user account. You can use this integration whether or not you have a GitHub Enterprise (GHE) license or subscription.

Another benefit to using the GitHub Enterprise integration is that you can choose to clone integration settings when you are creating new Snyk Organizations. This means you can use one GitHub Enterprise integration for all Organizations in your Snyk Group.

As a Free or Team plan user, Snyk recommends using the GitHub integration as it only requires a PAT for an individual user account, which should meet your needs at this level.

If you use the self-hosted GitHub Enterprise product, you must use the GitHub Enterprise integration.

{% hint style="info" %}
For detailed steps on migrating from GitHub to GitHub Enterprise, see [Migrate to GitHub Enterprise](snyk-scm-integrations/github.md#migrate-to-the-github-enterprise-integration).
{% endhint %}

### Bitbucket Cloud (PAT) vs Bitbucket Cloud App

In general, Snyk recommends using the new Bitbucket Cloud app integration. However, the new integration does not fit all cases. The information in this section is intended to help you decide which integration is right for you.

See [Migrate to the Snyk Bitbucket Cloud App](snyk-scm-integrations/bitbucket-cloud.md#migrate-to-the-snyk-bitbucket-cloud-app) for detailed migration instructions.

#### Main capabilities unlocked by the new app integration

* Allows using Snyk with Bitbucket's [allowlisting IP addresses](https://support.atlassian.com/bitbucket-cloud/docs/control-access-to-your-private-content/) premium tier feature.
* Helps handle rate-limiting issues for companies who spread their repos across multiple workspaces in Bitbucket Cloud.
* Supports the first-party interface in Bitbucket Cloud (Snyk's Security tab) out-of-the-box, meaning you need not install and maintain the first-party extension app to consume Snyk's security insights from Bitbucket Cloud.

#### Limitations of the new app integration

* In the new app integration, every Snyk Organization can connect to only one workspace in Bitbucket Cloud. If you want to import Projects from various workspaces in Bitbucket into the same single Organization in Snyk, use the PAT integration.
* The new app integration does not yet support Snyk Multi-Tenant EU, Snyk Multi-Tenant AUS, and Snyk Single-Tenant cloud deployments.
* For customers who are part of the custom branch closed beta, the new app integration's first-party interface in Bitbucket Cloud does not allow importing Projects from non-default branches. It is possible to import a non-default branch; you must do it from the Snyk.io import modal.

#### Are there any plans for end-of-life for the Personal Access Token (PAT) integration?

No, the Personal Access Token Bitbucket Cloud integration is fully supported, and there are no plans to stop supporting it.

However, there is a first-party interface _extension_ app that serves as an extension layer to the PAT integration, allowing developers to consume Snyk's findings from within the Bitbucket interface. This extension app was developed and supported by an external contractor company. As this functionality is now an integral part of the new app integration, the extension app has now moved to no-support mode, meaning that customers who use the PAT integration alongside the first-party extension app must migrate to the new app integration to get support for the first-party interface functionality.

## Pull Requests for Snyk integrations

Snyk can automatically create pull requests (PRs) on your behalf to upgrade your dependencies based on scan results. This is compatible with a variety of Snyk integrations. For more information, see [Upgrade open source dependencies with automatic PRs](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/upgrade-open-source-dependencies-with-automatic-prs.md).
