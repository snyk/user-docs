# npm Teams and npm Enterprise integration

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can use custom npm Teams and npm Enterprise repositories with npm and Yarn Projects.

This enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

After configuration, Snyk will also use this information to access private dependencies when creating Pull/Merge Requests, by allowing npm and yarn to reach those dependencies in order to regenerate the lockfile.

You can add configuration to tell Snyk where your private npm Teams and npm Enterprise Node.js packages are hosted and what scope they are under.

This is the same information you would normally add in your `.yarnrc` or `.npmrc`

{% hint style="info" %}
This guide is relevant for Snyk Web UI integrations only. The Snyk CLI already supports Yarn and npm projects with private npm Teams and npm Enterprise registries.
{% endhint %}

## JavaScript language settings

1. Go to settings <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> **> Languages > JavaScript** and either the npm or Yarn settings, depending on your Project types. Yarn is shown in the screenshots that follow.
2. If you have not previously connected to npm Teams or npm Enterprise, you will be asked to configure an integration first. See [npm Teams and npm Enterprise registry settings](npm-teams-and-npm-enterprise-integration.md#npm-teams-and-npm-enterprise-registry-settings).
3. After you have set up the integration, select **Add registry configuration.**
   1. Select **npm** as the Package source.
   2. If you want to configure this registry as **default registry url**, leave **scope** blank.
   3. If you want to configure **only scoped packages** to use this registry, add a **scope.**
   4. If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations, one for the default and one per scope.
4. When you have added all the registries and scopes you want, click **Update settings**.
5. to test the integration, open a Pull/Merge Request on a Project that contains private dependencies to see a lockfile updated and included in the Snyk Fix Pull Request where previously none was generated.

<figure><img src="../../.gitbook/assets/image (34) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (2).png" alt="Test of npm integration"><figcaption><p>Test of npm integration</p></figcaption></figure>

## npm Teams and npm Enterprise registry settings

You can configure token-based authentication for npm Teams and npm Enterprise integrations.

1. Go to settings <img src="../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> > **Integrations > Package Repositories > npm.**
2. You should see the **Credentials** screen at the beginning.
3. Enter **Public URL** and **Token** values.
4. Click **Save**.

<figure><img src="../../.gitbook/assets/image (35) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="npm credentials screen"><figcaption><p>npm credentials screen</p></figcaption></figure>
