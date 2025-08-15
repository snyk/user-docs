# JavaScript for open source

## Snyk for JavaScript support

Refer to the [JavaScript details](./) for supported package managers and features.

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of JavaScript packages this means a release to the npmjs package registry.&#x20;
{% endhint %}

## Open source and licensing

The following summarizes Snyk Open Source support for npm, pnpm and Yarn, and partial support for Lerna.

### npm

The following table provides a matrix for npm lockfile versions and Snyk features availability.

<table><thead><tr><th>Lockfile version</th><th width="144">CLI support</th><th>SCM support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>Lockfile v1</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>Lockfile v2</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>Lockfile v3</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr></tbody></table>

#### **Peer dependencies**

In npm v7 and above, the behavior of peer dependencies changes if they are being installed by default. To match this in npm v7+ projects, Snyk assumes peer dependencies are installed and scans them by default.

An npm v7+ project ignores peer dependencies only if they are explicitly marked as optional in the `peerDependenciesMeta` object in the `package.json` as shown here for `cache-manager`:

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

In npm v6 and below, peer dependencies are not scanned by default, as the package manager does not install them by default. To scan peer dependencies, ensure they are installed, and then run the CLI with the `--peer-dependencies` option.

#### **Lockfile versions**

Snyk uses the `package-lock.json` lockfile when present to generate a dependency tree for your project. These lockfiles come in different versions.

Lockfile v1 was used in npm v5 and v6. Two new formats were introduced in npm v7 - lockfile v2 and lockfile v3. For more information, see [lockfileVersion](https://docs.npmjs.com/cli/v9/configuring-npm/package-lock-json#lockfileversion).

You can see which lockfile format you are using in the `package-lock.json`, as follows:

```json
{
    ...
    "lockfileVersion": 3,
    ...
}
```

If you want to force npm to create a specific lockfile version, use the npm `--lockfile-version` parameter.

```bash
npm install --lockfile-version=2
```

### pnpm

{% hint style="info" %}
**Release status**

Snyk CLI pnpm support is in Early Access.

Enable it using [Snyk Preview](../../../snyk-platform-administration/snyk-preview.md) and install CLI v1.1293.0 or later.
{% endhint %}

The following table shows a matrix of pnpm versions and Snyk features availability.

<table><thead><tr><th>pnpm version</th><th>CLI support</th><th width="151">SCM support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>pnpm 7</td><td>✔︎ (Early Access)</td><td></td><td>✔︎ (Early Access)</td><td></td></tr><tr><td>pnpm 8</td><td>✔︎ (Early Access)</td><td></td><td>✔︎ (Early Access)︎</td><td></td></tr><tr><td>pnpm 9</td><td>✔︎ (Early Access)︎</td><td></td><td>✔︎ (Early Access)︎</td><td></td></tr><tr><td>pnpm 10</td><td>✔︎ (Early Access)︎</td><td></td><td>✔︎ (Early Access)︎</td><td></td></tr></tbody></table>

**Lockfile versions**

Snyk uses the `pnpm-lock.yaml` lockfile to generate a dependency tree for your project.&#x20;

The supported lockfile versions are 5.4, 6.x and 9.x, as used by pnpm 7, 8, 9 and 10.

pnpm lockfiles do not include [bundledDependencies](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#bundledependencies), so Snyk does not include them in scans.

### **Yarn**

Snyk uses the Yarn lockfile (`yarn.lock`) to generate a representation of project dependencies.

The files Snyk relies on to scan a project may change on version upgrades of the package manager. Snyk lists only versions verified internally as supported.

If you are using a newer version of Yarn than is listed on this page, you may find Snyk performs as expected because Yarn is using a lockfile version that is already supported. That version of Yarn has likely not been evaluated and, thus, not added to this page.&#x20;

The following table shows the Yarn versions and Snyk features availability matrix.

<table><thead><tr><th>Yarn Version</th><th width="148">CLI support</th><th>SCM support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>Yarn 1</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>Yarn 2</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>Yarn 3</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr></tbody></table>

{% hint style="info" %}
Because different versions of Yarn have different feature sets, there are differences in Snyk support in order to match how the package manager works.

**Resolutions** are supported in Yarn v2 and above. Yarn v1 resolutions are not supported.
{% endhint %}

### Lerna

Snyk does not fully support Lerna. If your project is set up using Yarn Workspaces, you can scan the project in the same way you scan any Yarn Workspaces project.

If your Lerna project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` as follows.

For each example package, you can use the following:

<pre class="language-javascript"><code class="lang-javascript"><strong>snyk monitor --file=packages/example-package/package.json
</strong></code></pre>

Alternatively, you can specify a script to automate scanning of nested `package.json` files:

```javascript
ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json
```

## Steps to start scanning using npm, pnpm and Yarn

The following table lists the steps to start scanning your dependencies. It covers basic commands, such as `snyk test` and `snyk monitor`. For a full list of CLI commands, see the [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md).

| Package manager | Getting started                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                               |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| npm             | <ol><li>Install npm.</li><li>Ensure you are in a directory with npm project files, that is, <code>package.json</code> and <code>package-lock.json</code>.</li><li>(Optional) Run <code>npm install</code>.</li><li>Run <a href="../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../../developer-tools/snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> and <a href="../../../developer-tools/snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>.<br></li></ol>                         | <p>Snyk analyzes your <code>package-lock.json</code> files to build a dependency tree. </p><p></p><p>If the <code>package-lock.json</code> is missing, Snyk analyzes your <code>node_modules</code> folder. </p><p></p><p>Alternatively, run <code>npm install</code> to generate the lockfile first.</p> |
| pnpm            | <ol><li>Install pnpm</li><li>Make sure that you are in a directory with pnpm project files, that is, <code>package.json</code> or <code>pnpm</code>and <code>pnpm-lock.yaml</code>.</li><li>(Optional)  Run <code>pnpm install</code>.</li><li>Run <a href="../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../../developer-tools/snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> and <a href="../../../developer-tools/snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>.</li></ol> | <p>Snyk analyzes your<code>pnpm-lock.yaml</code> files to build a dependency tree. <br><br>If the <code>pnpm-lock.yaml</code> is missing, Snyk analyzes your <code>node_modules</code> folder. <br><br>Alternatively, run <code>pnpm install</code> to generate the lockfile first.</p>                   |
| Yarn            | <ol><li>Install Yarn.</li><li>Make sure you are in a directory with Yarn project files, that is, <code>package.json</code> and <code>yarn.lock</code>.</li><li>(Optional) Run <code>yarn install</code></li><li>Run <a href="../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../../developer-tools/snyk-cli/commands/test.md#options-for-yarn-projects">snyk test</a> and <a href="../../../developer-tools/snyk-cli/commands/monitor.md#options-for-yarn-projects">snyk monitor</a>.</li></ol>                              | <p>Snyk analyzes your <code>yarn.lock</code> files to build a dependency tree. </p><p></p><p>If the <code>yarn.lock</code> is missing, Snyk analyzes your <code>node_modules</code> folder. </p><p></p><p>Alternatively, run <code>yarn install</code> to generate the lockfile first.</p>                |

## Support for monorepos and workspaces

Yarn, npm, and pnpm support workspaces, to help manage monorepos containing multiple sub-projects.

Workspaces are supported in the Snyk CLI for the following CLI options:

* `--all-projects` : Discover and scan all Yarn,  npm and pnpm workspaces projects, along with projects from other supported ecosystems. The root lock file is referenced when scanning the workspace projects.
* `--detection-depth` : Specify how many sub-directory levels to search.
* `--strict-out-of-sync=false` :  Allow testing out-of-sync lockfiles for packages in a  workspace. When this option is set to `false` , you can run Snyk tests with unsynchronized manifest and lock files without causing errors.
* `--policy-path` : Specify the path to a policy used by Snyk during testing.

### Examples of scanning workspaces

Scan all workspaces projects in the current directory and five sub-directories deep, plus any other Projects types detected.

```bash
snyk test --all-projects --strict-out-of-sync=false --detection-depth=6 
```

Use a common `.snyk` policy file, if you maintain ignores and patches in one place to be applied for all detected workspaces by using the policy path. For details about this file, see [The .snyk file](../../../manage-risk/policies/the-.snyk-file.md).

```bash
snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk
```

### **npm workspaces example**

npm v7 introduced support for workspaces. See the [lockfile version and Snyk feature availability matrix](javascript-for-open-source.md#npm).&#x20;

To detect and scan all workspaces in your npm Project, use the CLI options described above.

### **pnpm workspaces example**

pnpm workspaces must have the `package.json`, `pnpm-lock.yaml` and `pnpm-workspace.yaml` files in the root directory.

To detect and scan all workspaces in your pnpm Project, use the CLI options described above.

### **Yarn workspaces example**

{% hint style="info" %}
`nohoist`is not supported for Yarn Workspaces.
{% endhint %}

To detect and scan all workspaces in your Yarn Project, use the CLI options identified for monorepos and workspaces, as well as this Yarn-specific opton:

`--yarn-workspaces` : Use instead of `--all-projects` to detect and scan only Yarn workspaces projects when a lockfile is present in the root. Other ecosystems will be ignored.
