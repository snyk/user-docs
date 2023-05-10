# Using GitHub or GitHub Enterprise

{% hint style="info" %}
GitHub Enterprise integration is available to Enterprise customers. If you have a Legacy Business plan, contact [Snyk support](https://support.snyk.io/hc/en-us) for access. See [Pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

Snyk has integrations for both GitHub and GitHub Enterprise.

Snyk recommends GitHub Enterprise integration for most customers with access to the feature, because this integration allows use of a single Personal Access Token (PAT) across an Organization, rather than a PAT tied to an individual user account.

By using GitHub Enterprise integration, you can choose to clone integration settings when you are creating new Organizations. This means that you can use one GitHub Enterprise integration for all Organizations in your Group.

{% hint style="info" %}
You do not need a GitHub Enterprise license or subscription to use the GitHub Enterprise integration within Snyk.
{% endhint %}

## Migrating from GitHub integration to GitHub Enterprise integration

The [GitHub integration](github-integration.md) with Snyk is a simple way of connecting a Snyk user account with a GitHub account, using a GitHub personal access token, so you can import projects into Snyk.

This is a common starting point for new users to Snyk. However, there are some downsides to having individual users import Projects using their own tokens.

The most common issue relates to permissions of the user who imported a Project. If that user has permissions changed or removed in GitHub, the Project may no longer be accessible to Snyk for monitoring.

Thus Snyk recommends that you create a single service account user in GitHub that has the needed permissions on all repositories that you wish to monitor using Snyk.

Once this service account user is created, you can use the single personal access token from this service account to configure the GitHub Enterprise integration in Snyk, so that all Project imports are linked to the single account in GitHub.

## Ways to migrate from GitHub to GitHub Enterprise integration

Before you configure and Import Projects using the GitHub Enterprise integration, Snyk recommends that you remove all Projects imported using the GitHub integration. This avoids having duplicate projects in Snyk.

{% hint style="warning" %}
Before deleting Projects imported using the GitHub integration, consider the impact on reporting. For details contact your Snyk account team or [Snyk support](https://support.snyk.io/hc/en-us).
{% endhint %}

You can remove Projects imported using the GitHub integration manually by removing the projects from Snyk. However, depending on how many projects you have already imported using the GitHub integration, it may be easier for you to create a new Organization.

If you have already created multiple Organizations and GitHub Projects have been imported to each, then it may be easier to re-create the Organizations to use the GitHub Enterprise integration, rather than updating each one manually. If you choose to do this, you can copy integration settings from an existing Organization to avoid having to re-configure other integrations.

## Steps to migrate from GitHub to GitHub Enterprise integration

If you already have multiple Organizations with Projects imported using the GitHub integration, follow these steps to migrate from GitHub to GitHub Enterprise integration.

1. Create a new Organization that will be used as the template for all others. You can copy integration settings from an existing Organization if required.
2. In this new template Organization, set up the GitHub Enterprise integration using the steps at the top of the page [GitHub Enterprise integration](github-enterprise-integration.md). The dedicated service account in those steps is a separate user account in GitHub that you will use as the connection between Snyk and GitHub.
3. When the GitHub Enterprise integration is configured, you can import a Project to your template Organization to test that the integration is working as expected.
4. You can now create new Organizations that will replace the existing Organizations that were configured using the GitHub integration. As you create each new Organisation, ensure that you copy the integration settings from this template Organisation so that the GitHub Enterprise integration is available.
5. Now that your new Organizations are created, you can import your projects, choosing the GitHub Enterprise integration when you select the source.
6. You can now remove the previous Organizations that were configured using the GitHub integration.

You may want to [disconnect your GitHub integration](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration#disconnecting-the-github-integration) to avoid unintentionally importing Projects using the GitHub integration in the future. Because the GitHub integration is configured per user account, rather than per Organization, each user who has set up the GitHub integration must complete this disconnection process individually.
