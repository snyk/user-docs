# Git repositories and JavaScript

You can import JavaScript repositories from any Git services [supported by Snyk](../../developer-tools/scms/organization-level-integrations/). After the import, Snyk analyzes your Projects based on their supported manifest files.

## Snyk for JavaScript language settings

Configure language settings for open source and licensing at the Organization level. The configuration settings apply to all Projects in that Organization.

1. In your Snyk account, navigate to **Settings** > **Languages**..
2. Under **Languages**, navigate to **JavaScript** and select **Edit settings.**
3. Configure the settings based on your package manager, npm or Yarn.
   * **Scan and fix dev dependencies**: If this is selected, Snyk reads the `devDependencies` property on the `package.json` and reports and fixes any vulnerabilities accordingly.
   * **Require package.json and package-lock.json/yarn.lock files to be in sync**: When this is selected, if the `package.json` and `package-lock.json`/`yarn.lock`files are out-of-sync, Snyk fails the import.
   * **Exclude package-lock.json from being generated when fixing vulnerabilities**: If you are using private mirrors or registries, a Snyk-generated lockfile might not be appropriate for you because Snyk uses the npm registry to update the lockfile. Enterprise customers can use [package repository integrations](../../scan-with-snyk/snyk-open-source/package-repository-integrations/) to ensure lockfiles are updated correctly. Alternatively, this setting allows you to opt out of getting lockfiles generated for you in Snyk fix pull requests and merge requests.
4. **Update Settings** to save changes.

## Workspaces in Snyk for JavaScript

{% hint style="info" %}
Yarn and npm workspaces are not explicitly supported in Snyk git repository integration scans.
{% endhint %}

Root-level `package.json`manifest files with adjacent lockfiles will be scanned as normal.&#x20;

For nested manifest files with no lockfiles, Snyk will approximate what the dependency tree will look like at build time without using the root lockfile.&#x20;

In addition, Fix PRs and Upgrade PRs are not supported for nested manifest files with no lockfiles.

## Fix PRs and npm save-prefix

When creating a fix for vulnerabilities using npm v7+ Projects, Snyk will use the default npm `save-prefix` rather than inferring it from your Project.

This means if you have dependencies using a range format other than the caret range (`^`), you may see additional changes to the `version` fields in the `package-lock.json` file.&#x20;

These changes should not affect day-to-day functionality, as the ranges will be read from the `package.json`.

## Fix PRs for Yarn zero-installs users

In Yarn v2, the [zero-installs](https://yarnpkg.com/features/zero-installs) feature was released, which allowed Yarn developers to work on a Project without having to run `yarn` to install dependencies on their machine.&#x20;

Zero-installs achieved this by installing all the dependencies of a Project inside of the `.yarn/cache` directory and asking users to commit this to their version control system , allowing the next developer to pull any new dependencies directly from the repo.

{% hint style="info" %}
If you use the zero-installs feature, Snyk Fix PRs do not update the .yarn/cache directory. You must run `yarn` to update this directory.
{% endhint %}
