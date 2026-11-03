---
description: How to import and scan JavaScript repositories through Snyk SCM integrations, including Organization-level language settings
nav_context: agnostic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# SCM integrations for JavaScript

You can import JavaScript repositories from any SCM integration supported by Snyk. For more information, visit [Organization level integrations](https://docs.snyk.io/developer-tools/integrations/scm-integrations/organization-level-integrations). After the import, Snyk analyzes your Projects based on their supported manifest files.

## Configure language settings for JavaScript

You can configure language settings for open source and licensing at the Organization level. The configuration settings apply to all Projects in that Organization. To configure:

1. In the Snyk Web UI, navigate to **Settings** > **Snyk Open Source** > **Languages** > **JavaScript**.
2. Click **Edit settings**.
3. Configure the settings for your package manager: npm, pnpm, or Yarn.
   * **Scan and fix dev dependencies**: if you select this option, Snyk reads the `devDependencies` property on the `package.json` and reports and fixes any vulnerabilities accordingly.
   * **Require package.json and package-lock.json/yarn.lock files to be in sync**: if you select this option and the `package.json` and `package-lock.json`/`yarn.lock/pnpm-lock.yaml` files are out of sync, Snyk fails the import.
   * **Exclude package-lock.json from being generated when fixing vulnerabilities**: if you are using private mirrors or registries, it is possible that a Snyk-generated lockfile is not appropriate for you because Snyk uses the npm registry to update the lockfile. Enterprise customers can use [package repository integrations](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations) to ensure lockfiles are updated correctly. Alternatively, use this setting to opt out of lockfile generation in Snyk fix pull requests and merge requests.
4. Click **Update Settings** to save changes.

## Package registry integrations (Artifactory/Nexus)

{% hint style="info" %}
Artifactory, Nexus, npm Teams, and npm Enterprise Package Registry integrations are available only for Snyk Enterprise plans.
{% endhint %}

The Snyk Open Source Gatekeeper plugin integrates with Artifactory to block builds from downloading packages with vulnerabilities and license issues.

Snyk Open Source can also integrate with Artifactory, Nexus, npm Teams, and npm Enterprise to assist in the security testing of your applications. Snyk uses this integration for dependency resolution, fix calculation, and re-locking lockfiles.

If your Projects reference private dependencies in these repositories but you are not a Snyk Enterprise user, you can use the Snyk CLI in a properly configured local environment (such as your build pipeline) so these dependencies can be resolved and included in the test.

For more information, visit the following resources:

* Package registry integrations: [npm Teams and npm Enterprise](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations/npm-teams-and-npm-enterprise-integration), [Artifactory Registry setup](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup), and [Nexus Repository Manager setup](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup)
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/manage-vulnerabilities/artifactory-gatekeeper-plugin)

## Fix PRs and npm save-prefix

When Snyk creates a fix for vulnerabilities in npm v7+ Projects, Snyk uses the default npm `save-prefix` rather than inferring it from your Project.

This means that if you have dependencies using a range format other than the caret range (`^`), it is possible that you see additional changes to the `version` fields in the `package-lock.json` file.

These changes do not affect day-to-day functionality, because the ranges are read from the `package.json`.

## Fix PRs for Yarn zero-installs users

Yarn v2 introduced the [zero-installs](https://yarnpkg.com/features/zero-installs) feature, which lets Yarn developers work on a Project without running `yarn` to install dependencies on their machine.

Zero-installs stores all the dependencies of a Project in the `.yarn/cache` directory, and users commit this directory to their version control system. The next developer can then pull any new dependencies directly from the repository.

{% hint style="info" %}
If you use the zero-installs feature, Snyk Fix PRs do not update the `.yarn/cache` directory. You must run `yarn` to update this directory.
{% endhint %}
