# Using GitHub or GitHub Enterprise integration

{% hint style="info" %}
**Feature availability**\
GitHub Enterprise integration is available to Snyk Enterprise plan customers. If you have a Legacy Business plan, contact [Snyk support](https://support.snyk.io/hc/en-us) for access. See the  [Plans and pricing](https://snyk.io/plans/) page for details.
{% endhint %}

Snyk has an integration for GitHub and a GitHub Enterprise integration for Snyk Enterprise plan customers. When you navigate to the Integrations page in your Snyk account, you see these integrations represented by two tiles, **GitHub Configured**, and **GitHub Enterprise**.

The GitHub Enterprise integration is the Snyk integration for Snyk Enterprise plan customers. You can use this integration whether or not you have a GitHub Enterprise (GHE) license or subscription.  If you use the self-hosted GitHub Enterprise product, you must use the Snyk GitHub Enterprise integration.

If you are a Snyk Enterprise plan customer, Snyk recommends that you click the **GitHub Enterprise** tile and set up that integration. The GitHub Enterprise integration allows you to use a single GitHub service account personal access token (PAT) across your Snyk Organization rather than depending on a PAT for an individual user account.

When you use the GitHub Enterprise integration, you can choose to clone integration settings when you are creating new Snyk Organizations. This means you can use one GitHub Enterprise integration for all Organizations in your Snyk Group.

## Migration from GitHub integration to GitHub Enterprise integration

The [GitHub integration](snyk-github-integration.md) with Snyk is a simple way of connecting a Snyk user account with a GitHub account using a GitHub personal access token so that you can import Projects into Snyk.

This is a common starting point for new users to Snyk. However, there are some downsides to having individual users import Projects using their own tokens.

The most common issue relates to the permissions of the user who imported a Project. If that user has permissions changed or removed in GitHub, the Project may no longer be accessible to Snyk for monitoring.

Thus. Snyk recommends creating a single GitHub service account user with the needed permissions on all repositories you monitor using Snyk.

After this GitHub service account user is created, you can use the single personal access token from this service account to configure the GitHub Enterprise integration in Snyk so that all Snyk Projects imported from GitHub are linked to a single GitHub service account.

## Ways to migrate from GitHub integration to GitHub Enterprise integration

Before you configure and Import Projects using the Snyk GitHub Enterprise integration, Snyk recommends removing all Projects imported using the GitHub integration. This avoids having duplicate Snyk Projects.

{% hint style="warning" %}
Before deleting Snyk Projects imported using the GitHub integration, consider the impact on reporting. For details, contact your Snyk account team or [Snyk support](https://support.snyk.io/hc/en-us).
{% endhint %}

You can remove Projects imported using the GitHub integration manually by removing the Projects from Snyk. However, depending on how many Projects you have already imported using the GitHub integration, it may be easier for you to create a new Snyk Organization.

If you have already created multiple Snyk Organizations and Projects have been imported from GitHub to each Organization, it may be easier to re-create the Snyk Organizations to use the Snyk GitHub Enterprise integration than to update each one manually. If you choose to do this, you can copy integration settings from an existing Organization to avoid having to re-configure other integrations.

## Steps to migrate from GitHub integration to the Snyk GitHub Enterprise integration

If you already have multiple Snyk Organizations with Projects imported using the GitHub integration, follow these steps to migrate from GitHub integration to GitHub Enterprise integration.

1. Create a new Snyk Organization that will be used as the template for all others. You can copy integration settings from an existing Organization if required.
2. In this new template Organization, set up the Snyk GitHub Enterprise integration using the steps on the page [GitHub Enterprise integration](github-enterprise-integration.md#how-to-set-up-a-github-enterprise-integration). The dedicated GitHub service account in those steps is a separate user account that you will use as the connection between Snyk and GitHub.
3. When the Snyk GitHub Enterprise integration is configured, you can import a Project to your template Organization to test that the integration is working as expected.
4. You can now create new Organizations that will replace the existing Organizations that were configured using the GitHub integration. As you create each new Organization, ensure that you copy the integration settings from this template Organization so that the GitHub Enterprise integration is available.
5. Now that your new Organizations are created, you can import your Projects, choosing the GitHub Enterprise integration when you select the source.
6. You can now remove the previous Organizations that were configured using the GitHub integration.

You may want to [disconnect your GitHub integration](snyk-github-integration.md#disconnecting-the-github-integration) to avoid unintentionally importing Projects using the GitHub integration in the future. Because the GitHub integration is configured per user account, rather than per Organization, each user who has set up the GitHub integration must complete this disconnection process individually.
