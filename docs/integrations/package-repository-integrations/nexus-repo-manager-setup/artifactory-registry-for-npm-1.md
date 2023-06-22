# Nexus Repository Manager for npm

## O**verview**

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can use Nexus Repository Manager with npm and Yarn projects imported from Git.

This enables Snyk to regenerate lockfiles with the correct URLs when creating Pull/Merge Requests.

You can add configuration to tell Snyk where your private Nexus Repository Manager Node.js packages are hosted and what scope they are under.

This is the same information you would normally add in your `.yarnrc` or `.npmrc`

{% hint style="info" %}
This guide is relevant for Snyk Web UI integrations only, the Snyk CLI already supports Yarn and npm projects with private Nexus Repository Manager registries.
{% endhint %}

## JavaScript Language Settings

Go to settings <img src="../../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> **> Languages > JavaScript** and either the npm or yarn settings depending on your project types.

If you have not previously connected to Nexus Repository Manager you will be asked to configure an integration first, see [.](./ "mention")

![](<../../../.gitbook/assets/Screenshot 2022-07-15 at 14.18.43.png>)

Now follow the steps below, according to your version of Nexus.

{% tabs %}
{% tab title="Nexus 3" %}
1. Select “Add registry configuration”
2. Select "Nexus" as the Package source
3. If you want to configure this registry as **default registry url**, then leave scope blank.
4. If you want to configure **only scoped packages** to use this registry then add a scope.
5. If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations - one for the default and one per scope.
6. The **Repository** section should be set as whatever comes after `repository/` in the internal repository URL.\
   For example if the URL is `http://nexus.company.io/repository/npm-group`, Repository should be set as `npm-group`
7. When you have added all the registries and scopes you want, click **Update settings**.
{% endtab %}

{% tab title="Nexus 2" %}
1. Select “Add registry configuration”
2. Select "Nexus" as the Package source
3. If you want to configure this registry as **default registry url**, then leave scope blank.
4. If you want to configure **only scoped packages** to use this registry then add a scope.
5. If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations - one for the default and one per scope.
6. The **Repository** section should be set as whatever comes after `nexus/content/` in the internal repository URL.\
   For example if the URL is `http://nexus.company.io/nexus/content/groups/npm-group`, Repository should be set as `groups/npm-group` .\
   Or `http://nexus.company.io/nexus/content/repositories/npm-hosted`, Repository should be set as `repositories/npm-hosted`
7. When you have added all the registries and scopes you want, click **Update settings**.
{% endtab %}
{% endtabs %}

### Test it out

Open a Pull/Merge Request on a project that contains private dependencies that are hosted in Nexus to see **a lockfile updated and included in the Snyk Fix Pull Request with the correct URL to your repository.**

![](<../../../.gitbook/assets/Screenshot 2022-07-15 at 14.22.59.png>)
