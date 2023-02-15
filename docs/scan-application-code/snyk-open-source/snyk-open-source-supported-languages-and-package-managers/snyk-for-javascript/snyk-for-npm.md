# Snyk for npm

You can use Snyk to scan your JavaScript project managed by npm.

## Features of Snyk for npm

{% hint style="info" %}
**Feature availability**\
Features might not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| npm Version / Feature | CLI Support | Git Support | License Scanning | Fix Prs |
| --------------------- | ----------- | ----------- | ---------------- | ------- |
| V6 and Below          | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| V7                    | ✔︎          | ✔︎          | ✔︎               | ✔︎      |

## npm version and how it affects Snyk support

Snyk uses the npm lockfile (`package-lock.json`) when it is present to generate representation of a Project's dependencies. Until npm v7, this lockfile was always the same format. In npm v7 a new lockfile format, lockfile v2, was introduced.

With changes to the files Snyk relies on to scan changing on version upgrades, Snyk lists only versions verified internally as supported. If you are using a newer version of npm than is listed on this page, you may find Snyk performs as expected and this newer version of npm has not been evaluated and added to this page.

To see if you are using the new lockfile format, look in in the `package-lock.json` itself:

```json
{
    ...
    "lockfileVersion": 2,
    ...
}
```

## How Snyk for npm works

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

The way Snyk analyzes and builds the graph varies depending on the language and package manager of the Project, as well as the location of your project. For the ways Snyk can be run see [Snyk CLI for npm projects](snyk-for-npm.md#snyk-cli-for-npm-projects) and [Git services for npm projects](snyk-for-npm.md#git-services-for-npm-projects).

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

There are **CLI options** you can use with the CLI commands to refine your scan:

`--strict-out-of-sync` **true** / false

Prevent testing out-of-sync lockfiles (test fails when set to true if there are out-of-sync lockfiles in the project).

`--fail-on`**all** / upgradable / patchable

Configure when a test should be failed if there are vulnerabilities as follows:

* All-fail for all projects containing vulnerabilities
* Upgradable-fail only for projects with vulnerabilities that can be fixed with package upgrades
* Patchable-fail for projects with vulnerabilities that can be fixed with either upgrades or patches

`--prune-repeated-subdependencies` true / **false**

Use this option if any big projects fail to be tested.

`--dev` true / **false**

Set to true if Snyk should scan dev dependencies.

`--all-projects`

Use this option to detect and scan all npm and other projects in this directory.

## Git services for npm projects

JavaScript Projects can be imported from any of the Git services, GitHub, GitLab, BitBucket, and so on, that Snyk supports. After import, Snyk analyzes your Projects based on their supported manifest files.

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

### Update language preferences

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages.**
3. Select **Edit settings** for JavaScript to configure preferences for your JavaScript (npm and Yarn) projects in this Organization.

## Differences due to npm versions

Due to different versions of npm having different feature sets there are differences in Snyk support to best match how the package manager works.

In npm v7 and above the behavior of **peer dependencies** changed if they are are installed by default. To match this in npm v7+ projects, Snyk assumes peer dependencies are installed and scans them by default. The only case in which an npm v7+ Project ignores peer dependencies is if they are explicitly marked as optional in the `peerDependenciesMeta` object in the `package.json` as shown here for `cache-manager`:

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

In npm v6 and below peer dependencies are not scanned by default, as the package manager does not install them by default. If you wish to scan peer dependencies, make sure they have been installed and then run the CLI with the `--peer-dependencies` option.

## Limitations of Snyk for npm

npm v7 introduced the concept of **workspaces**. These are not yet supported by Snyk
