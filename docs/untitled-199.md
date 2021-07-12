# npm Teams & npm Enterprise for npm

* [ Artifactory Registry Setup](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360013805638-Artifactory-Registry-Setup/README.md)
* [ Artifactory Registry for Maven](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360005507418-Artifactory-Registry-for-Maven/README.md)
* [ Artifactory Registry for npm](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007537418-Artifactory-Registry-for-npm/README.md)
* [ npm Teams & npm Enterprise for npm](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360009411777-npm-Teams-npm-Enterprise-for-npm/README.md)
* [ Private Gem Sources for Ruby](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360013742557-Private-Gem-Sources-for-Ruby/README.md)

## npm Teams & npm Enterprise for npm

### Overview

This guide is relevant for Snyk UI integrations only, the CLI already supports yarn and npm projects with private npm Teams and npm Enterprise registries

You can add configuration to tell Snyk where your private npm Teams and npm Enterprise Node.js packages are hosted and what scope they are under.

This is the same information you would normally add in your `.yarnrc` or `.npmrc`

Once configured, Snyk will use this information to access private dependencies when creating Pull/Merge Requests**,** by allowing yarn/npm to reach those deps in order to regenerate the lockfile.

**Feature availability**  
This feature is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.

### JavaScript Language Settings

1. Go to settings  &gt; **Languages &gt; JavaScript** and either the npm or yarn settings depending on your project types \(yarn shown in screenshots below\)
2. If you have not previously connected to npm Teams or npm Enterprise you will be asked to configure an integration first, see [npm Teams & npm Enterprise Registry Settings](untitled-199.md) below
3. Once you have set up the integration, select **Add registry configuration**  
   * Select "npm" as the Package source
   * If you want to configure this registry as **default registry url**, then leave scope blank
   * If you want to configure **only scoped packages** to use this registry then add a scope
   * If you want to add a mix of **default registry url** and **scoped packages**, add multiple configurations - one for the default and one per scope.
4. When you have added all the registries and scopes you want, click **Update settings**.
5. Now test it out - open a Pull/Merge Request on a project that contains private dependencies to see a lockfile updated and included in the Snyk Fix Pull Request where previously none was generated. _\*\*_

### npm Teams & npm Enterprise Registry Settings <a id="h_bab29371-80cb-4e6e-ad3b-e930b24700b1"></a>

You can configure token based authentication for npm Teams and npm Enterprise integrations.

#### Getting started

1. Go to settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Integrations &gt; Package Repositories &gt; npm**
2. You should see this screen at the beginning.
3. Enter **Public URL** and **Token** values.
4. Click **Save**.

