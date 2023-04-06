# Snyk for Yarn

You can use Snyk to scan your JavaScript projects managed by Yarn.

## Features of Snyk for Yarn

{% hint style="info" %}
**Feature availability**\
Features may not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

| Yarn Version / Feature | CLI Support | Git Support | License Scanning | Fix Prs |
| ---------------------- | ----------- | ----------- | ---------------- | ------- |
| Yarn 1                 | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| Yarn 2                 | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| Yarn 3                 | ✔︎          | ✔︎          | ✔︎               | ✔︎      |

## Yarn version and how it affects Snyk support

Snyk uses the Yarn lockfile (`yarn.lock`) to generate representation of Project dependencies.

The files Snyk relies on to scan a Project may change on version upgrades of the package manager.Therefore Snyk lists only versions verified internally as supported.

If you are using a newer version of Yarn than is listed on this page, you may find Snyk performs as expected because Yarn is using a lockfile version that is already supported. That version of Yarn has likely not been evaluated and thus not been added to this page.

## How Snyk for Yarn works

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager, and that your Project contains the supported manifest files.
{% endhint %}

The way Snyk analyzes and builds the graph varies depending on the language and package manager of the Project, as well as the location of your Project.

For the ways you can scan Projects with Snyk, see [Snyk CLI for Yarn projects](snyk-for-yarn.md#snyk-cli-for-yarn-projects) and [Git services for Yarn projects](snyk-for-yarn.md#git-services-for-yarn-projects).

## Snyk CLI for Yarn projects

Snyk analyzes your `package.json` and `yarn.lock` files to build a fully structured dependency tree. If the `yarn.lock` is missing, Snyk analyzes your `node_modules` folder.

To get started using the CLI for Yarn projects:

* Make sure Yarn is installed.
* Make sure you are in a directory with Yarn Project files, that is, `package.json` and `yarn.lock`.
* Run `yarn`.
* [Install](../../../../snyk-cli/install-the-snyk-cli/) and authenticate the Snyk CLI.

You can now test and monitor your Project using `snyk test` or `snyk monitor` .

### CLI options for Snyk for Yarn

For information about the `snyk test` options available for use with Yarn, see [Options for Yarn projects in the Test help](https://docs.snyk.io/snyk-cli/commands/test#options-for-yarn-projects). For the available `snyk monitor` options, see [Options for Yarn projects in the Monitor help](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-yarn-projects).

### Differences due to Yarn versions

Because different versions of Yarn have different feature sets, there are differences in Snyk support in order to best match how the package manager works.

**Resolutions** are supported in Yarn v2 only. Yarn v1 resolutions are not supported.

### Yarn Workspaces In CLI

{% hint style="danger" %}
`nohoist` is **not** supported for Yarn Workspaces.
{% endhint %}

For Yarn Workspaces use the `--all-projects` flag to test and monitor your packages with other Projects or `--yarn-workspaces` to specifically scan Yarn Workspaces Projects only. The root lock file is referenced when scanning all the packages. Use the `--detection-depth` option to find sub-folders that are not auto-discovered by default.

Example usage:\
`snyk test --all-projects --strict-out-of-sync=false --detection-depth=6` scans the packages that belong to any discovered workspaces in this directory and five sub-directories deep, as well as any other Projects detected.

`snyk test --yarn-workspaces --strict-out-of-sync=false --detection-depth=6` scans only the Yarn Workspace packages that belong to any discovered workspaces in this directory and five sub-directories deep.

You may use a common `.snyk` policy file if you maintain ignores and patches in one place to be applied for all detected workspaces by using the policy path:

`snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk`

## Git services for Yarn projects

Yarn Projects can be imported from any of the Git services Snyk supports. After import, Snyk analyzes your Projects based on their supported manifest files.

Snyk scans based on these files being present:

* `package.json`
* `yarn.lock`

{% hint style="info" %}
For Yarn Workspaces, only the `package.json` file is updated for Snyk fix PRs. The `yarn.lock` file is not updated.
{% endhint %}

### Fix PRs for [zero-installs](https://yarnpkg.com/features/zero-installs) users

In Yarn V2 the [zero-installs](https://yarnpkg.com/features/zero-installs) feature was released, which allowed Yarn developers to work on a project without having to run `yarn` to install dependencies on their machine. It achieved this by installing all the dependencies of a project inside of the `.yarn/cache` directory and asking users to commit this to their version control system - allowing the next developer to pull any new dependencies directly from the repo.

{% hint style="warning" %}
If you are using the **zero-installs** feature, any Snyk fix PRs do not update the **.yarn/cache** directory. You must run `yarn` to update this directory.
{% endhint %}

### Git settings for Yarn

From the Snyk UI, use these parameters to customize your language preferences for JavaScript-based projects:

#### Preferences for Snyk for Yarn

| Preference                                                         | Description                                                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Scan and fix devDependencies                                       | If this is selected, Snyk reads the "devDependencies" property on the `package.json` and reports and fixes any vulnerabilities accordingly.                                                                                                                                                    |
| Require package.json and yarn.lock to be in sync                   | When this is selected, if the `package.json` and `yarn.lock` files are out of sync, Snyk fails the import.                                                                                                                                                                                     |
| Exclude yarn.lock from being generated when fixing vulnerabilities | If you are using private mirrors or registries, a Snyk-generated lock file might not be appropriate for you because Snyk uses the npm registry to update the lockfile. This setting allows you to opt out of getting lockfiles generated for you in Snyk fix pull requests and merge requests. |

## Update language preferences for Snyk for Yarn

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages**
3. Select **Edit settings** for JavaScript to configure preferences for your JavaScript (npm and Yarn) projects in this Organization.
