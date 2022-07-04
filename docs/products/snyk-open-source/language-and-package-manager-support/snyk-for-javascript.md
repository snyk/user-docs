# Snyk for JavaScript

You can use Snyk to scan your JavaScript projects.

### Features

{% hint style="info" %}
**NOTE**\
Features might not be available, depending on your subscription plan.
{% endhint %}

| Package managers / Features | <p>CLI</p><p>support</p> | <p>Git</p><p>support</p> | <p>License</p><p>scanning</p> | Fix PRs |
| --------------------------- | ------------------------ | ------------------------ | ----------------------------- | ------- |
| npm                         | ✔︎                       | ✔︎                       | ✔︎                            | ✔       |
| Yarn                        | ✔︎                       | ✔︎                       | ✔︎                            | ✔︎      |
| Yarn v2                     | ✔︎                       | ✖️                       | ✔︎                            | ✖️      |
| Yarn Workspaces             | ✔︎                       | ✔︎                       | ✔                             | ✖️      |

### How it works

After Snyk builds a dependency tree, we use our [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

The way Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project.

See:

* [Snyk CLI tool for JavaScript projects](snyk-for-javascript.md#snyk-cli-tool-for-javascript-projects)
* [Git services for JavaScript projects](snyk-for-javascript.md#git-services-for-javascript-projects)

## Snyk CLI tool for JavaScript projects

The way Snyk analyzes and builds the tree varies based on the package manager of the project.

### npm

npm versions 6.x, 7.x are supported in the Snyk CLI.

{% hint style="info" %}
Workspaces npm 7.x is not supported.
{% endhint %}

We analyze your `package.json` and `package-lock.json` files, to build a full structured dependency tree. If the `package-lock.json` is missing, we analyze your `node_modules` folder.

Snyk can apply previously selected patches using the GNU patch utility. Patches are saved to the .snyk policy file.

**Peer dependencies** are scanned by default in npm v7.x _unless_ they are explicitly marked as optional in the `peerDependenciesMeta` object of the `package.json` (In npm v6.x add `--peer-dependencies` to your command, as these are not installed by default).

```
"peerDependenciesMeta": {
        "cache-manager": {
          "optional": true
        },
```

### Yarn

Yarn versions 1 & 2 are supported in the Snyk CLI.

We analyze your `package.json` and `yarn.lock` files, to build a full structured dependency tree. If the `yarn.lock` is missing, we analyze your `node_modules` folder.

Snyk supports resolutions only in Yarn v2. We do not support resolutions for Yarn v1.

### Yarn Workspaces

{% hint style="info" %}
Yarn v1 and 2 workspaces support is for `snyk test` and `snyk monitor` commands only at this time.
{% endhint %}

{% hint style="danger" %}
nohoist is **not** supported for Yarn workspaces.
{% endhint %}

For Yarn workspaces use the `--all-projects` flag to test and monitor your packages alongside other projects or `--yarn-workspaces` to specifically scan Yarn workspaces projects only. The root lock file is referenced when scanning all the packages. Use the `--detection-depth` parameter to find sub-folders that are not auto-discovered by default.

Example usage:\
`snyk test --all-projects --strict-out-of-sync=false --detection-depth=6` which will scan the packages that belong to any discovered workspaces this directly and 5 sub-directories deep as well as any other projects detected.

`snyk test --yarn-workspaces --strict-out-of-sync=false --detection-depth=6` which will scan only the Yarn workspace packages that belong to any discovered workspaces this directly and 5 sub-directories deep.

You may use a common **.snyk** policy file if you maintain ignores/patches in one place to be applied for all detected workspaces by providing the policy path:

`snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk`

### Lerna

{% hint style="info" %}
Snyk currently does not support Lerna. However, if your project is set up using Yarn Workspaces, you can scan the project in the same way you scan any Yarn Workspaces project. If your Lerna project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` with the following commands.
{% endhint %}

For each example-package you can use the following:

`snyk monitor --file=packages/example-package/package.json`

Alternatively, you can specify a script to automate scanning nested \``` package.json` ``files:

`ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json`

### CLI parameters for JavaScript

Prerequisites

* Install the relevant package manager.
* Include the relevant manifest files supported by Snyk.
* Install and authenticate [the Snyk CLI](https://docs.snyk.io/snyk-cli/install-the-snyk-cli) to start analyzing projects from your local environment.

Run `npm install` or `yarn install,` depending on the package manager you use for your JavaScript projects.

{% hint style="info" %}
As snyk test looks at the locally installed modules, it needs to be run after npm install or yarn install, and will seamlessly work with shrinkwrap, npm enterprise or any other custom installation logic you have.
{% endhint %}

**Parameters**

Use the following options to refine the scan:

| **Option**                         | <p><strong>Values</strong><br><strong>(default in bold)</strong></p> | **Description**                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------------------- | -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `--strict-out-of-sync`             | **true**,false                                                       | Prevent testing out-of-sync lockfiles (test fails when set to true if there are out-of-sync lockfiles in the project)                                                                                                                                                                                                                                                          |
| `--fail-on`                        | **all**, upgradable, patchable                                       | <p>Configure when a test should be failed if there are vulnerabilities as follows:</p><ul><li>All-fail for all projects containing vulnerabilities</li><li>Upgradable-fail only for projects with vulnerabilities that can be fixed with package upgrades</li><li>Patchable-fail for projects with vulnerabilities that can be fixed with either upgrades or patches</li></ul> |
| `--prune-repeated-subdependencies` | true, **false**                                                      | Use this flag if any big projects fail to be tested. Default is _false_                                                                                                                                                                                                                                                                                                        |
| `--dev`                            | true, **false**                                                      | Set to true if Snyk should scan dev dependencies                                                                                                                                                                                                                                                                                                                               |
| `--yarn-workspaces`                | **n/a**                                                              | <p>Provide this flag to only scan a Yarn workspace project where lockfile is in the root. By default <code>--all-projects</code> automatically detects and scans Yarn workspaces projects.</p><p>Note: <code>snyk test</code> and<code>snyk monitor</code> commands only.</p>                                                                                                  |
| `--all-projects`                   | **n/a**                                                              | <p>Provide this flag to detect and scan all Node and other projects.</p><p>Note: <code>snyk test</code> and<code>snyk monitor</code> commands only.</p>                                                                                                                                                                                                                        |

## Git services for JavaScript projects

JavaScript projects can be imported from any of the Git services we support. After import, Snyk analyzes your projects based on their supported manifest files.

### npm

npm versions 6.x, 7.x are supported in Git services.

{% hint style="info" %}
Workspaces npm 7.x is not supported.
{% endhint %}

We build the dependency tree based on these files:

* `package.json`
* `package-lock.json`

### Yarn

{% hint style="info" %}
Yarn version 1 is supported in Git services.
{% endhint %}

We build the dependency tree based on these files:

* `package.json`
* `yarn.lock`

### Yarn Workspaces

{% hint style="info" %}
**Note**\
Git support for Yarn Workspaces is enabled for all projects in organizations created after March 3rd 2021. To enable this feature for organizations created before this date [contact our Support team](https://support.snyk.io/hc/en-us/requests/new). Yarn version 1 is supported in Git services.
{% endhint %}

For Yarn Workspaces we scan each `package.json` that matches the `packages` pattern from the root level `package.json` and root level `yarn.lock`

Fix Pull/Merge Requests are not supported for Yarn Workspaces. The fix advice can be used to manually generate PRs.

Commit status checks always use the root level `yarn.lock` and workspace `package.json` for tests.

{% hint style="info" %}
**Warning**\
If your `package.json` and root `yarn.lock` are out-of-sync, we will have issues re-testing the projects. Snyk shows errors on project page and import logs when this happens.\
If you reference the locally installed packages which do not appear in a lockfile, you can disable the **Require package.json and yarn.lock files to be in sync** setting, on the **Languages Settings** page for JavaScript.
{% endhint %}

### Git settings for JavaScript

**Preferences**

From the Snyk UI, use these parameters to customize your language preferences for JavaScript-based projects:

| Option                                                                         | Description                                                                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Scan and fix devDependencies**                                               | If selected, Snyk reads the \`devDependencies\` property on the package.json and reports & fixes any vulnerabilities accordingly.                                                                                                                                                          |
| **Require package.json and package-lock.json to be in sync**                   | When selected if the package.json and package.lock files are out-of-sync, Snyk fails the import.                                                                                                                                                                                           |
| **Exclude package-lock.json from being generated when fixing vulnerabilities** | If you are using private mirrors or registries, a Snyk generated lockfile might not be appropriate for you because Snyk uses the npm registry to update the lockfile. This setting allows you to opt-out of getting lockfiles generated for you in our fix pull requests / merge requests. |

### Update language preferences

1. Log in to your account and navigate to the relevant group and organization that you want to manage
2. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **Languages**
3. Click **Edit settings** for JavaScript to configure preferences for your JavaScript (npm and Yarn) projects in this organization
