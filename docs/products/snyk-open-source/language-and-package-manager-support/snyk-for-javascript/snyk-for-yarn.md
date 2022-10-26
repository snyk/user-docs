# Snyk for Yarn

You can use Snyk to scan your JavaScript projects managed by Yarn.

### Feature support

{% hint style="info" %}
**NOTE**\
Features might not be available, depending on your subscription plan.
{% endhint %}

| Yarn Version / Feature | CLI Support | Git Support | License Scanning | Fix Prs |
| ---------------------- | ----------- | ----------- | ---------------- | ------- |
| Yarn 1                 |  ✔︎         |  ✔︎         |  ✔︎              |  ✔︎     |
| Yarn 2                 |  ✔︎         | ✖️          |  ✔︎              | ✖️      |

### Yarn version and how it affects Snyk support

Snyk uses the Yarn lockfile (`yarn.lock`) to generate representation of project dependencies. With the release of Yarn 2 a new format for this lockfile was released; hence the disparity in feature support shown in the preceding table.

With changes to the files Snyk relies on to scan changing on version upgrades, Snyk lists only versions verified internally as supported. If you are using a newer version of Yarn than is listed on this page, you may find Snyk performs as expected because Yarn is using an already supported lockfile version. That version of Yarn has likely not been evaluated and thus not been added to this page.&#x20;

To see if you are using the new lockfile format, look for the string `__metadata` in the `yarn.lock` file itself.

### How Snyk for Yarn works

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

The way Snyk analyzes and builds the graph varies depending on the language and package manager of the project, as well as the location of your project.

See the following:

* [Snyk CLI for Yarn projects](snyk-for-yarn.md#snyk-cli-tool-for-yarn-projects)
* [Git services for Yarn projects](snyk-for-yarn.md#git-services-for-npm-projects)

## Snyk CLI for Yarn projects

Snyk analyzes your `package.json` and `yarn.lock` files, to build a fully structured dependency tree. If the `yarn.lock` is missing, Snyk analyzes your `node_modules` folder.

To get started using the CLI for Yarn projects:

* Make sure Yarn is installed.
* Make sure you are in a directory with a yarn project files, that is,`package.json` and `yarn.lock`.&#x20;
* Run `yarn`.
* Install and authenticate [Snyk CLI](https://docs.snyk.io/snyk-cli/install-the-snyk-cli).

You can now test and monitor your project using `snyk test` or `snyk monitor` .

### CLI Options

There are options you can use with the CLI commands to refine your scan:

| Option                             | Values (default in bold)         | Description                                                                                                                                                                                                                                                                                                                                                               |
| ---------------------------------- | -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--strict-out-of-sync`             | **true** / false                 | Prevent testing out-of-sync lockfiles (test fails when set to true if there are out-of-sync lockfiles in the project).                                                                                                                                                                                                                                                    |
| `--fail-on`                        | **all** / upgradable / patchable | <p>Configure when a test should fail if there are vulnerabilities as follows:</p><ul><li>All-fail for all projects containing vulnerabilities</li><li>Upgradable-fail only for projects with vulnerabilities that can be fixed with package upgrades</li><li>Patchable-fail for projects with vulnerabilities that can be fixed with either upgrades or patches</li></ul> |
| `--prune-repeated-subdependencies` | true / **false**                 | Use this option if any big projects fail to be tested.                                                                                                                                                                                                                                                                                                                    |
| `--dev`                            | true / **false**                 | Set to true if Snyk should scan dev dependencies.                                                                                                                                                                                                                                                                                                                         |
| `--all-projects`                   | N/A                              | Use this option to detect and scan all Yarn and other projects in this directory.                                                                                                                                                                                                                                                                                         |
| `--yarn-workspaces`                | N/A                              | Use this option only to scan a Yarn Workspace project where lockfile is in the root. By default `--all-projects` automatically detects and scans Yarn Workspace projects.                                                                                                                                                                                                 |

### Differences due to Yarn versions

Because different versions of Yarn have different feature sets, there are differences in Snyk support in order to best match how the package manager works.

**Resolutions** are supported in Yarn v2 only. Yarn v1 resolutions are not supported.

### Yarn Workspaces In CLI

{% hint style="danger" %}
`nohoist` is **not** supported for Yarn Workspaces.
{% endhint %}

For Yarn Workspaces use the `--all-projects` flag to test and monitor your packages alongside other projects or `--yarn-workspaces` to specifically scan Yarn Workspaces projects only. The root lock file is referenced when scanning all the packages. Use the `--detection-depth` option to find sub-folders that are not auto-discovered by default.

Example usage:\
`snyk test --all-projects --strict-out-of-sync=false --detection-depth=6` scans the packages that belong to any discovered workspaces in this directory and 5 sub-directories deep, as well as any other projects detected.

`snyk test --yarn-workspaces --strict-out-of-sync=false --detection-depth=6` scans only the Yarn Workspace packages that belong to any discovered workspaces in this directory and 5 sub-directories deep.

You may use a common `.snyk` policy file if you maintain ignores and patches in one place to be applied for all detected workspaces by using the policy path:

`snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk`

## Git services for Yarn projects

Yarn projects can be imported from any of the Git services Snyk supports. After import, Snyk analyzes your projects based on their supported manifest files.

Snyk scans based on these files being present:

* `package.json`
* `yarn.lock`

### Git settings for Yarn

From the Snyk UI, use these parameters to customize your language preferences for JavaScript-based projects:

#### Preferences

| Preference                                                         | Description                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Scan and fix devDependencies                                       | If this is selected, Snyk reads the "devDependencies" property on the `package.json` and reports and fixes any vulnerabilities accordingly.                                                                                                                                                   |
| Require package.json and yarn.lock to be in sync                   | When this is selected, if the `package.json` and `yarn.lock` files are out of sync, Snyk fails the import.                                                                                                                                                                                    |
| Exclude yarn.lock from being generated when fixing vulnerabilities | If you are using private mirrors or registries, a Snyk-generated lockfile might not be appropriate for you because Snyk uses the npm registry to update the lockfile. This setting allows you to opt out of getting lockfiles generated for you in Snyk fix pull requests and merge requests. |

#### Update language preferences

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Click on settings <img src="../../../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line">  > **Languages**
3. Click **Edit settings** for JavaScript to configure preferences for your JavaScript (npm and Yarn) projects in this Organization.
