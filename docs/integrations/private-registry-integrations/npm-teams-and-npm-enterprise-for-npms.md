# npm Teams & npm Enterprise for npms

## **Overview**

{% hint style="info" %}
This guide is relevant for Snyk UI integrations only, the CLI already supports yarn and npm projects with private npm Teams and npm Enterprise registries.
{% endhint %}

You can add configuration to tell Snyk where your private npm Teams and npm Enterprise Node.js packages are hosted and what scope they are under.

This is the same information you would normally add in your `.yarnrc` or `.npmrc`

Once configured, Snyk will use this information to access private dependencies when creating Pull/Merge Requests**,** by allowing yarn/npm to reach those deps in order to regenerate the lockfile.

{% hint style="info" %}
**Feature availability**  
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## JavaScript Language Settings

1. Go to settings ![](../../.gitbook/assets/cog_icon.png) **&gt; Languages &gt; JavaScript** and either the npm or yarn settings depending on your project types \(yarn shown in screenshots below\) 
2. If you have not previously connected to npm Teams or npm Enterprise you will be asked to configure an integration first, see [npm Teams & npm Enterprise Registry Settings](npm-teams-and-npm-enterprise-for-npms/) below 
3. Once you have set up the integration, select **Add registry configuration**
   1. Select "npm" as the Package source
   2. If you want to configure this registry as **default registry url**, then leave scope blank
   3. If you want to configure **only scoped packages** to use this registry then add a scope
   4. If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations - one for the default and one per scope.
4. When you have added all the registries and scopes you want, click **Update settings**.
5. Now test it out - open a Pull/Merge Request on a project that contains private dependencies to see a lockfile updated and included in the Snyk Fix Pull Request where previously none was generated.

![](../../.gitbook/assets/image%20%2834%29.png)

## npm Teams & npm Enterprise Registry Settings

You can configure token based authentication for npm Teams and npm Enterprise integrations.

## Getting started

1. Go to settings ![](../../.gitbook/assets/cog_icon.png) &gt; **Integrations &gt; Package Repositories &gt; npm** 
2. You should see the **Credentials** screen at the beginning \(see below\/)
3. Enter **Public URL** and **Token** values. 
4. Click **Save**.

![](../../.gitbook/assets/image%20%2835%29.png)

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

