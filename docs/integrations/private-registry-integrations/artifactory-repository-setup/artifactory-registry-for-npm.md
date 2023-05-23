# Artifactory Registry for npm

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can use Artifactory Package Repositories with npm and Yarn Projects.

This enables Snyk to regenerate lockfiles with the correct URLs when creating Pull/Merge Requests.

You can add configuration to tell Snyk where your private Artifactory Node.js packages are hosted and what scope they are under.

This is the same information you would normally add in your `.yarnrc` or `.npmrc`

{% hint style="info" %}
Note that this guide is relevant for Snyk UI integrations only. The CLI already supports Yarn and npm projects with private Artifactory registries.
{% endhint %}

## JavaScript language settings

1. Go to settings <img src="../../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> **> Languages > JavaScript** and either the npm or Yarn settings depending on your project types
2. If you have not previously connected to Artifactory you will be asked to configure an integration first; see [Artifactory Registry setup](./).
3. Select **Add registry configuration**.
   1. Select **Artifactory** as the Package source.
   2. If you want to configure this registry as **default registry url**, leave **scope** blank.
   3. If you want to configure **only scoped packages** to use this registry, add a **scope**.
   4. If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations, one for the default and one per scope.
4. When you have added all the registries and scopes you want, click **Update settings**.

## Test the integration

Open a Pull/Merge Request on a Project that contains private dependencies that are hosted in Artifactory to see **a lockfile updated and included in the Snyk Fix Pull Request with the correct URL to your repository**.

<figure><img src="../../../.gitbook/assets/image4-3-.png" alt="Pull request to test Artifactory integration"><figcaption><p>Pull request to test Artifactory integration</p></figcaption></figure>
