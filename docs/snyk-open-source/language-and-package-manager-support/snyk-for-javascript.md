# Snyk for JavaScript

You can use Snyk to scan your JavaScript projects.

### Note

Features might not be available, depending on your subscription plan.

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| ![i\_icon\_npm.png](https://support.snyk.io/hc/article_attachments/360007183097/uuid-aa98c079-4d3e-5e58-f450-08985b5d4b8b-en.png) | npm | ✔︎ | ✔︎ | ✔︎ | ✔︎ | ✔︎ |
| ![i\_icon\_yarn.png](https://support.snyk.io/hc/article_attachments/360007259138/uuid-374be14a-4c05-fdb4-a25c-e8fbdedc2990-en.png) | Yarn | ✔︎ | ✔︎ | ✔︎ | ✔︎ | ✔︎ |
| ![i\_icon\_yarn.png](https://support.snyk.io/hc/article_attachments/360007259138/uuid-374be14a-4c05-fdb4-a25c-e8fbdedc2990-en.png) | Yarn Workspaces | ✔︎ | ✔︎ | ✔︎ | ✔︎ \(advice\) | ✔︎ |

After Snyk builds a dependency tree, we use our [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

To scan your dependencies, ensure you install the relevant package manager, and that your project contains the supported manifest files.

The way Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project.

See:

* [Snyk CLI tool for JavaScript projects](snyk-for-javascript.md)
* [Git services for JavaScript projects](snyk-for-javascript.md)

## Snyk CLI tool for JavaScript projects

The way Snyk analyzes and builds the tree varies based on the package manager of the project.

### npm

npm versions 6.x, 7.x are supported in the Snyk CLI.

Workspaces npm 7.x is not supported.

We analyze your `package.json` and `package-lock.json` files, to build a full structured dependency tree. If the `package-lock.json` is missing, we analyze your `node_modules` folder.

Snyk can apply previously selected patches using the GNU patch utility. Patches are saved to the .snyk policy file.

### Yarn

Yarn versions 1 & 2 are supported in the Snyk CLI.

We analyze your `package.json` and `yarn.lock` files, to build a full structured dependency tree. If the `yarn.lock` is missing, we analyze your `node_modules` folder.

Snyk supports resolutions only in Yarn v2 in repositories _without_ workspaces. We do not support resolutions for Yarn v1.

### Yarn workspaces

Yarn v1 workspaces support is for `snyk test` and `snyk monitor` commands only at this time. We do not support workspaces for Yarn v2.

For Yarn workspaces use the `--yarn-workspaces` flag to test and monitor your packages. The root lockfile is referenced when scanning all the packages. Use the `--detection-depth` parameter to find sub-folders which are not auto discovered by default.

Example usage:  
`snyk test --yarn-workspaces --strict-out-of-sync=false --detection-depth=6` which will scan only the packages that belong to any discovered workspaces this directly and 5 sub-directories deep.

You may use a common **.snyk** policy file if you maintain ignores/patches in one place to be applied for all detected workspaces by providing the policy path:

`snyk test --yarn-workspaces --strict-out-of-sync=false --policy-path=src/.snyk`

### CLI parameters for JavaScript

Prerequisites

* Install the relevant package manager.
* Include the relevant manifest files supported by Snyk.
* Install and authenticate [the Snyk CLI](https://support.snyk.io/hc/en-us/sections/360001295038-Install-the-Snyk-CLI) to start analyzing projects from your local environment.

Run `npm install` or `yarn install,` depending on the package manager you use for your JavaScript projects.

As snyk test looks at the locally installed modules, it needs to be run after npm install or yarn install, and will seamlessly work with shrinkwrap, npm enterprise or any other custom installation logic you have.

**Parameters**

Use the following options to refine the scan:

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Option</b>
      </th>
      <th style="text-align:left"><b>Values<br />(default in bold)</b>
      </th>
      <th style="text-align:left"><b>Description</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><code>--strict-out-of-sync</code>
      </td>
      <td style="text-align:left"><b>true</b>,false</td>
      <td style="text-align:left">Prevent testing out-of-sync lockfiles (test fails when set to true if
        there are out-of-sync lockfiles in the project)</td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--fail-on</code>
      </td>
      <td style="text-align:left"><b>all</b>, upgradable, patchable</td>
      <td style="text-align:left">
        <p>Configure when a test should be failed if there are vulnerabilities as
          follows:</p>
        <ul>
          <li>All-fail for all projects containing vulnerabilities</li>
          <li>Upgradable-fail only for projects with vulnerabilities that can be fixed
            with package upgrades</li>
          <li>Patchable-fail for projects with vulnerabilities that can be fixed with
            either upgrades or patches</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--prune-repeated-subdependencies</code>
      </td>
      <td style="text-align:left">true, <b>false</b>
      </td>
      <td style="text-align:left">Use this flag if any big projects fail to be tested. Default is <em>false</em>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--dev</code>
      </td>
      <td style="text-align:left">true, <b>false</b>
      </td>
      <td style="text-align:left">Set to true if Snyk should scan dev dependencies</td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--yarn-workspaces</code>
      </td>
      <td style="text-align:left"><b>n/a</b>
      </td>
      <td style="text-align:left">
        <p>Provide this flag to signal that the scanned project is a Yarn workspace
          so the root level lockfile can be used for all packages.</p>
        <p>Note: <code>snyk test</code> and<code>snyk monitor</code> commands only.</p>
      </td>
    </tr>
  </tbody>
</table>

## Git services for JavaScript projects

JavaScript projects can be imported from any of the Git services we support. After import, Snyk analyzes your projects based on their supported manifest files.

### npm

npm versions 6.x, 7.x are supported in Git services.

Workspaces npm 7.x is not supported.

We build the dependency tree based on these files:

* `package.json`
* `package-lock.json`

### Yarn

Yarn version 1 is supported in Git services.

We build the dependency tree based on these files:

* `package.json`
* `yarn.lock`

### Yarn Workspaces

### Note

Git support for Yarn Workspaces is enabled for all projects in organisations created after March 3rd 2021. To enable this feature for organisations created before this date please contact support@snyk.io.

For Yarn Workspaces we scan each `package.json` that matches the `packages` pattern from the root level `package.json` and root level `yarn.lock`.

Fix Pull/Merge Requests are not supported for Yarn Workspaces. The remediation advice can be used to manually generate PRs.

Commit status checks fail if the `package.json` and the root `yarn.lock` are out of sync.

Commit status checks always use the root level `yarn.lock` and workspace `package.json` for tests.

### Warning

If your `package.json` and root `yarn.lock` are out of sync, we will have issues re-testing the projects. Snyk shows errors on project page and import logs when this happens.

If you reference the locally installed packages which do not appear in a lockfile, you can disable the **Require package.json and yarn.lock files to be in sync** setting, on the **Languages Settings** page for JavaScript.

### Git settings for JavaScript

**Preferences**

From the Snyk UI, use these parameters to customize your language preferences for JavaScript-based projects:

| Option | Description |
| :--- | :--- |
| **Scan and fix devDependencies** | If selected, Snyk reads the \`devDependencies\` property on the package.json and reports & fixes any vulnerabilities accordingly. |
| **Require package.json and package-lock.json to be in sync** | When selected if the package.json and package.lock files are out-of-sync, Snyk fails the import. |
| **Exclude package-lock.json from being generated when fixing vulnerabilities** | If you are using private mirrors or registries, a Snyk generated lockfile might not be appropriate for you because Snyk uses the npm registry to update the lockfile. This setting allows you to opt-out of getting lockfiles generated for you in our fix pull requests / merge requests. |

### Update language preferences

1. Log in to your account and navigate to the relevant group and organization that you want to manage.

   ![AddProjectMenu.gif](https://support.snyk.io/hc/article_attachments/360007259158/uuid-da316a4a-c823-cf03-f37f-5305446dc970-en.gif)

2. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Languages.**
3. Click **Edit settings** for JavaScript to configure preferences for your JavaScript \(npm and Yarn\) projects in this organization.

