# Snyk for npm

You can use Snyk to scan your JavaScript project managed by npm.

## Features of Snyk for npm

{% hint style="info" %}
**Feature availability**\
Features may not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

| Lockfile version / Feature | CLI Support | Git Support               | License Scanning | Fix Prs |
| -------------------------- | ----------- | ------------------------- | ---------------- | ------- |
| Lockfile v1                | ✔︎          | ✔︎                          | ✔︎               | ✔︎      |
| Lockfile v2                | ✔︎          | ✔︎                          | ✔︎               | ✔︎      |
| Lockfile v3                | ✔︎          | [Beta](#lockfile-versions) | ✔︎               |         |

## How Snyk for npm works

{% hint style="info" %}
Support may vary depending on the lockfile version used in your project, see the [Lockfile versions](snyk-for-npm.md#lockfile-versions) section for more information.
{% endhint %}

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

The way Snyk analyzes and builds the graph varies depending on the language and package manager of the Project, as well as the location of your project.&#x20;

For the ways you can scan Projects with Snyk see [Snyk CLI for npm projects](#snyk-cli-for-npm-projects) and [Git services for npm projects](snyk-for-npm.md#git-services-for-npm-projects).

## Snyk CLI for npm projects

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager and that your Project contains the supported manifest files.
{% endhint %}

Snyk analyzes your `package.json` and `package-lock.json` files to build a fully structured dependency tree. If the `package-lock.json` is missing, Snyk analyzes your `node_modules` folder.

To get started using the CLI for npm projects:

* Make sure npm is installed.
* Make sure you are in a directory with npm Project files, that is, `package.json` and `package-lock.json`.
* Run `npm install`.
* Install and authenticate Snyk CLI.

You can now test and monitor your project using `snyk test` or `snyk monitor`.

You can use options with Snyk CLI commands to refine your scan.&#x20;

For information about the `snyk test` options available for use with npm, see [Options for npm projects](https://docs.snyk.io/snyk-cli/commands/test#options-for-npm-projects) in the Test help.&#x20;

For the available `snyk monitor` options, see [Options for npm projects in the Monitor help](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-npm-projects).

## Git services for npm projects

JavaScript Projects can be imported from any of the Git services that Snyk supports. After import, Snyk analyzes your Projects based on their supported manifest files.

Snyk scans your Projects based on these files being present:

* `package.json`
* `package-lock.json`

## Git settings for npm

### Preferences for JavaScript-based Projects

From the Snyk UI, use these parameters to customize your language preferences for JavaScript-based projects:

| Preference                                                                 | Description                                                                                                                                                                                                                                                                                   |
| -------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Scan and fix devDependencies                                               | If this is selected, Snyk reads the "devDependencies" property on the `package.json` and reports and fixes any vulnerabilities accordingly.                                                                                                                                                   |
| Require package.json and package-lock.json to be in sync                   | When this is selected, if the `package.json` and `package-lock.json` files are out-of-sync, Snyk fails the import.                                                                                                                                                                            |
| Exclude package-lock.json from being generated when fixing vulnerabilities | If you are using private mirrors or registries, a Snyk-generated lockfile might not be appropriate for you because Snyk uses the npm registry to update the lockfile. This setting allows you to opt out of getting lockfiles generated for you in Snyk fix pull requests and merge requests. |

### Update language preferences for Snyk for npm

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages.**
3. Select **Edit settings** for JavaScript to configure preferences for your JavaScript (npm and Yarn) projects in this Organization.

## Additional support notes

### Lockfile versions

{% hint style="info" %}
**Lockfile v3** support for Git is currently in beta. To participate in the beta [contact us](mailto:support@snyk.io).
{% endhint %}

Snyk uses the `package-lock.json` lockfile when present to generate a dependency tree for your project. These lockfiles come in different versions.

Lockfile v1 was used in npm v5 and v6. Two new formats were introduced in npm v7 - lockfile v2 and lockfile v3 (see [lockfileVersion](https://docs.npmjs.com/cli/v9/configuring-npm/package-lock-json#lockfileversion)).

You can see which lockfile format you are using in the `package-lock.json`, as follows:

```json
{
    ...
    "lockfileVersion": 3,
    ...
}
```

If you want to force npm to create a specific lockfile version use the npm `--lockfile-version` parameter.

```bash
npm install --lockfile-version=2
```

### Workspaces

npm v7 introduced the concept of **workspaces**.&#x20;

Workspaces are supported in the Snyk CLI, but not currently in the Git integrations.

To detect and scan all workspaces in your npm project, use the `--all-projects` Snyk CLI parameter, as follows:

```bash
snyk test --all-projects
```

### Peer dependencies

In npm v7 and above, the behaviour of **peer dependencies** changes if they are being installed by default. To match this in npm v7+ projects, Snyk assumes peer dependencies are installed and scans them by default.

The only case in which an npm v7+ project ignores peer dependencies is if they are explicitly marked as optional in the `peerDependenciesMeta` object in the `package.json` as shown here for `cache-manager`:

```json
{
    ...
    "peerDependenciesMeta": {
        "cache-manager": {
            "optional": true
        }
    },
    ...
}
```

In npm v6 and below, peer dependencies are not scanned by default, as the package manager does not install them by default. To scan peer dependencies, make sure they are installed, and then run the CLI with the `--peer-dependencies` option.
