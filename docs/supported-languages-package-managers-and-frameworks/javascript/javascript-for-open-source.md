# JavaScript for open source

## Snyk for JavaScript support

**Package managers**: npm, Yarn

**Package manager versions**:&#x20;

npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`

Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

**Package registry**: [npmjs.org](https://www.npmjs.org)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:npm`

**Test your app's packages**: Available, `pkg:npm`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

## Open source and licensing

The following summarizes Snyk Open Source support for npm and Yarn and partial support for Lerna.

### npm

The following table shows the npm lockfile versions and Snyk features availability matrix.

| Lockfile version | CLI support | Git support | License scanning | Fix PRs |
| ---------------- | ----------- | ----------- | ---------------- | ------- |
| Lockfile v1      | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| Lockfile v2      | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| Lockfile v3      | ✔︎          | ✔︎          | ✔︎               | ✔︎      |

**Peer dependencies**

In npm v7 and above, the behavior of **peer dependencies** changes if they are being installed by default. To match this in npm v7+ Projects, Snyk assumes peer dependencies are installed and scans them by default.

The only case in which an npm v7+ Project ignores peer dependencies is if they are explicitly marked as optional in the `peerDependenciesMeta` object in the `package.json` as shown here for `cache-manager`:

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

**Lockfile versions**

Snyk uses the `package-lock.json` lockfile when present to generate a dependency tree for your Project. These lockfiles come in different versions.

Lockfile v1 was used in npm v5 and v6. Two new formats were introduced in npm v7 - lockfile v2 and lockfile v3 (see [lockfileVersion](https://docs.npmjs.com/cli/v9/configuring-npm/package-lock-json#lockfileversion)).

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

### **Yarn**

Snyk uses the Yarn lockfile (`yarn.lock`) to generate a representation of Project dependencies.

The files Snyk relies on to scan a Project may change on version upgrades of the package manager. Snyk lists only versions verified internally as supported.

If you are using a newer version of Yarn than is listed on this page, you may find Snyk performs as expected because Yarn is using a lockfile version that is already supported. That version of Yarn has likely not been evaluated and, thus, not added to this page.&#x20;

The following table shows the Yarn versions and Snyk features availability matrix.

| Yarn Version | CLI support | Git support | License scanning | Fix PRs |
| ------------ | ----------- | ----------- | ---------------- | ------- |
| Yarn 1       | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| Yarn 2       | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| Yarn 3       | ✔︎          | ✔︎          | ✔︎               | ✔︎      |

{% hint style="info" %}
Because different versions of Yarn have different feature sets, there are differences in Snyk support in order to match how the package manager works.

**Resolutions** are supported in Yarn v2 and above. Yarn v1 resolutions are not supported.
{% endhint %}

### Lerna

Snyk does not fully support **Lerna**. If your Project is set up using Yarn Workspaces, you can scan the Project in the same way you scan any Yarn Workspaces Project.

If your Lerna Project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` as follows:

For each example package, you can use the following:

<pre class="language-javascript"><code class="lang-javascript"><strong>snyk monitor --file=packages/example-package/package.json
</strong></code></pre>

Alternatively, you can specify a script to automate scanning of nested `package.json` files:

```javascript
ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json
```

### Steps to start scanning using npm and Yarn

The following table lists the steps to start scanning your dependencies. It covers basic commands, such as `snyk test` and `snyk monitor`. To check the full list, see [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md).

| Package manager | Getting started                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                    |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| npm             | <ol><li>Install NPM.</li><li>Make sure you are in a directory with NPM Project files, that is, <code>package.json</code> and <code>package-lock.json</code>.</li><li>Run <code>npm install</code>.</li><li>Run <a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> and <a href="../../snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>.<br></li></ol> | <p>Snyk analyzes your <code>package.json</code> and <code>package-lock.json</code> files to build a fully structured dependency tree. </p><p></p><p>If the <code>package-lock.json</code> is missing, Snyk analyzes your <code>node_modules</code> folder.</p> |
| Yarn            | <ol><li>Install Yarn.</li><li>Make sure you are in a directory with Yarn Project files, that is, <code>package.json</code> and <code>yarn.lock</code>.</li><li>Run <code>yarn</code></li><li>Run <a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../snyk-cli/commands/test.md#options-for-yarn-projects">snyk test</a> and <a href="../../snyk-cli/commands/monitor.md#options-for-yarn-projects">snyk monitor</a>.</li></ol>                 | <p>Snyk analyzes your <code>package.json</code> and <code>yarn.lock</code> files to build a fully structured dependency tree. </p><p></p><p>If the <code>yarn.lock</code> is missing, Snyk analyzes your <code>node_modules</code> folder.</p>                 |

For Yarn Workspaces, only the `package.json` file is updated for Snyk Fix PRs. The `yarn.lock` file is not updated.

## Monorepo Projects

Snyk scans one manifest file at a time. For information about scanning all manifest files, see [Does the Snyk CLI support monorepos or multiple manifest files?](https://support.snyk.io/hc/en-us/articles/360000910577-Does-the-Snyk-CLI-support-monorepos-or-multiple-manifest-files-)

### **Scan all npm workspaces**

NPM v7 introduced the concept of **workspaces**. See [lockfile version and Snyk feature availability matrix](javascript-for-open-source.md#npm).&#x20;

To detect and scan all workspaces in your npm Project, use the `--all-projects` Snyk CLI parameter, as follows:

```javascript
snyk test --all-projects
```

### **Scan all Yarn workspaces**

{% hint style="info" %}
`nohoist`is **not** supported by Yarn Workspaces.
{% endhint %}

For Yarn Workspaces, you can use the following options:

* `--all-projects` : Test and monitor your packages with other Projects or `--yarn-workspaces` , scanning only Yarn Workspaces Projects. The root lock file is referenced when scanning all the packages.&#x20;
* `--detection-depth` : Find sub-folders that are not auto-discovered by default.
* `--strict-out-of-sync=false` : Relax strict synchronization requirements for packages in a Yarn workspace. When set to `false` , you can run Snyk tests with unsynchronized `package.json` and the `package-lock.json` files without throwing errors. Using different dependency versions can introduce potential risks, such as compatibility issues or security vulnerabilities.
* `--policy-path` : Specify the path to a policy used by Snyk during testing.

Scan the packages that belong to any discovered workspaces in this directory, five deep sub-directories, and any other Projects detected.

```javascript
snyk test --all-projects --strict-out-of-sync=false --detection-depth=6 
```

### Examples of scanning Yarn workspaces

Scan only the Yarn Workspace packages that belong to any discovered workspaces in this directory and five deep sub-directories.

```javascript
snyk test --yarn-workspaces --strict-out-of-sync=false --detection-depth=6
```



Use a common `.snyk` policy file, if you maintain ignores and patches in one place to be applied for all detected workspaces by using the policy path (see [The .snyk file](../../manage-risk/policies/the-.snyk-file.md)).

```javascript
snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk
```
