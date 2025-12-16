# SCM integrations for JavaScript

You can import JavaScript repositories from any SCM integration supported by Snyk. See [Organization level integrations](../../../developer-tools/scm-integrations/organization-level-integrations/). After the import, Snyk analyzes your Projects based on their supported manifest files.

## Configure language settings for Snyk with JavaScript

You can configure language settings for open source and licensing at the Organization level. The configuration settings apply to all Projects in that Organization. To configure:

1. In the Snyk Web UI, navigate to **Settings** > **Snyk Open Source** > **Languages** > **JavaScript**.
2. Click **Edit settings.**
3. Configure the settings based on your package manager - npm, pnpm, or Yarn.
   * **Scan and fix dev dependencies**: if you check this option, Snyk reads the `devDependencies` property on the `package.json` and reports and fixes any vulnerabilities accordingly.
   * **Require package.json and package-lock.json/yarn.lock files to be in sync**: when this is checked if the `package.json` and `package-lock.json`/`yarn.lock/pnpm-lock.yaml` files are out-of-sync, Snyk fails the import.
   * **Exclude package-lock.json from being generated when fixing vulnerabilities**: if you are using private mirrors or registries, it is possible that a Snyk-generated lockfile is not appropriate for you because Snyk uses the npm registry to update the lockfile. Enterprise customers can use [package repository integrations](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/) to ensure lockfiles are updated correctly. Alternatively, this setting allows you to opt out of getting lockfiles generated for you in Snyk fix pull requests and merge requests.
4. Click **Update Settings** to save changes.

## Package registry integrations (Artifactory/Nexus)

{% hint style="info" %}
Artifactory, Nexus, npm Teams, and npm Enterprise Package Registry integrations are available only for Snyk Enterprise plans.
{% endhint %}

Snyk Open Source Gatekeeper plugins integrates with Artifactory to block builds from downloading packages with vulnerability and license issues.

Snyk Open Source can also integrate with Artifactory, Nexus, npm Teams, and npm Enterprise to assist in the security testing of your applications. Snyk uses this integration for dependency resolution, fix calculation, and re-locking lock files.

If your Projects reference private dependencies in these repositories but you are not a Snyk Enterprise user, you can use the Snyk CLI in a properly configured local environment (such as your build pipeline) so these dependencies can be resolved and included in the test.

For more information, see the following external resources:

* Package registry integrations: [npm Teams and npm Enterprise](https://docs.snyk.io/scan-with-snyk/snyk-open-source/package-repository-integrations/npm-teams-and-npm-enterprise-integration), [Artifactory Registry setup](https://docs.snyk.io/scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup) and [Nexus Repository Manager setup](https://docs.snyk.io/scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup).
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](https://docs.snyk.io/scan-with-snyk/snyk-open-source/manage-vulnerabilities/artifactory-gatekeeper-plugin)

## Fix PRs and npm save-prefix

When creating a fix for vulnerabilities using npm v7+ Projects, Snyk uses the default npm `save-prefix` rather than inferring it from your Project.

This means that if you have dependencies using a range format other than the caret range (`^`), it is possible that you see additional changes to the `version` fields in the `package-lock.json` file.&#x20;

These changes do not affect day-to-day functionality, as the ranges will be read from the `package.json`.

## Fix PRs for Yarn zero-installs users

In Yarn v2, the [zero-installs](https://yarnpkg.com/features/zero-installs) feature was released, which allowed Yarn developers to work on a Project without having to run `yarn` to install dependencies on their machine.

Zero-installs achieved this by installing all the dependencies of a Project inside of the `.yarn/cache` directory and asking users to commit this to their version control system, allowing the next developer to pull any new dependencies directly from the repository.

{% hint style="info" %}
If you use the zero-installs feature, Snyk Fix PRs do not update the .yarn/cache directory. You must run `yarn` to update this directory.
{% endhint %}
