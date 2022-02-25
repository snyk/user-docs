# Artifactory Registry for npm

## **Overview**

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can use custom Artifactory Package Repositories with npm and Yarn projects.

This enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

Once configured, Snyk will also use this information to access private dependencies when creating Pull/Merge Requests\*\*,\*\* by allowing npm and yarn to reach those deps in order to regenerate the lockfile.

You can add configuration to tell Snyk where your private Artifactory Node.js packages are hosted and what scope they are under.

This is the same information you would normally add in your `.yarnrc` or `.npmrc`

{% hint style="info" %}
**Note**\
This guide is relevant for Snyk UI integrations only, the CLI already supports yarn and npm projects with private Artifactory registries.
{% endhint %}

## JavaScript Language Settings

1. Go to settings ![](../../../.gitbook/assets/cog\_icon.png) **> Languages > JavaScript** and either the npm or yarn settings depending on your project types (yarn shown in screenshots below)
2. If you have not previously connected to Artifactory you will be asked to configure an integration first, see [Artifactory Registry Setup.](https://docs.snyk.io/integrations/private-registry-integrations/artifactory-registry-setup)
3. Select “Add registry configuration”
   1. Select "Artifactory" as the Package source
   2. If you want to configure this registry as **default registry url**, then leave scope blank
   3. If you want to configure **only scoped packages** to use this registry then add a scope
   4. If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations - one for the default and one per scope.
4. Once connected, you can go back to the JavaScript Language Setting and select "Artifactory" from the Integration type drop-down and configure scope
5. When you have added all the registries and scopes you want, click **Update settings**.
6. Now test it out - open a Pull/Merge Request on a project that contains private dependencies that are hosted in Artifactory to see **a lockfile updated and included in the Snyk Fix Pull Request where previously none was generated**

![](../../../.gitbook/assets/image4-3-.png)
