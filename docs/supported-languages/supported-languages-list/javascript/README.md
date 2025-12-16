# JavaScript

{% hint style="info" %}
JavaScript is supported for Snyk Code and Snyk Open Source.&#x20;
{% endhint %}

## JavaScript for Snyk Code

For an overview of the supported security rules, visit [JavaScript and TypeScript rules](../../../scan-with-snyk/snyk-code/snyk-code-security-rules/javascript-and-typescript-rules.md).

### Supported frameworks and libraries

The following frameworks and libraries are supported:

{% columns %}
{% column %}
* @Google Drive/generative-ai
* @anthropic-ai/sdk
* @huggingface/inference
* @mistralai/mistralai
* axios
* Angular
* apollo-server
* bcrypt-nodejs
* cross-spawn
* crypto-js
* date-fns
* dayjs
* dompurify
* electron
* ejs
* execa
* express
* express-mongo-sanitize
* express-graphql
* express-jwt
* fastMCP
* fs
* fs-extra
* fs-plus
* graceful-fs
* graphql-js
* grpc-js
* hapi.js
* jQuery
* js-yaml
* jzip
* koa
* koa-graphql
* libxml
* libxmljs
* lodash
* luxon
* minimongo
{% endcolumn %}

{% column %}
* minimist
* modelcontextprotocol/typescript-sdk
* mongodb
* Mongoose
* mercurius
* Nestjs
* Node Crypto
* node-buffer
* node-cmd
* Node Crypto
* node-dir
* node-forge
* node-pty
* node-serialize
* octokit
* openai
* pg
* pg-promise
* React - Partial
* request-promise
* restler
* rimraf
* sanitize-html
* shelljs
* Stanford JS Crypto
* superagent
* tar-stream
* TSOA
* unirest
* unzip
* underscore
* url
* vm
* webstomp-client
* WebCryptoAPI
* xpath
* yargs


{% endcolumn %}
{% endcolumns %}

### Supported file formats

The following file formats are supported:  `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

### Available features

* Reports
* Interfile analysis

## JavaScript for Snyk Open Source

### Supported package managers and package registries

Snyk supports the following package managers and versions:

*   npm: `npm 5`, `npm 6`, `npm 7`, `npm 8`, `npm 9`, `npm 10+`

    Supported Lockfile versions: `Lockfile v1`, `Lockfile v2`, `Lockfile v3`
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`, `pnpm 10`
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

Snyk's default package registry is [npmjs.org](https://www.npmjs.org/). Private package registries are supported. For more information, visit [Package repository integrations.](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/)

### Available integrations

* SCM import
* CLI and IDE: test or monitor your app

### Supported file formats

The following file formats are supported:

* For npm: `package.json` and `package-lock.json`
* For pnpm: `pnpm-lock.yaml`
* For yarn: `yarn.lock`

Lerna is partially supported.

### Available features

* Fix PRs
* License scanning
* Reports
* Test your app's SBOM and packages using `pkg:npm`  PURLs, using [SBOM test](../../../developer-tools/snyk-cli/commands/sbom-test.md) command.

### Language and package manager considerations

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of JavaScript packages this means a release to the npmjs.org package registry.&#x20;
{% endhint %}

#### devDependencies analysis

`devDependencies` analysis is disabled by default as these are not typically elevated to production, often seen as “noise” by both security and development. To enable testing on dev-dependencies:

* Use the `--dev` parameter for CLI and CI/CD integrations.
* For SCM integrations, set using **Settings** > **Languages** in the relevant configuration item.

#### optionalDependencies analysis

optionalDependencies are included by default for CLI and CI/CD, as well as SCM integrations.

#### Unmanaged JavaScript

If you are on the Enterprise plan and thus have access to the Snyk API, can use the API to get a full list of dependencies and their transitive dependencies.

To test for vulnerabilities, you can use the following API endpoints:

* [Test for issues in a public package by name and version](../../../snyk-api/reference/test-v1.md#test-npm-packagename-version)
* [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph)
* [List issues for a package](../../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues)

#### Out of sync lockfiles

Control behavior when the lockfile and package file are in sync can be done using:

* CLI additional values: `--strict-out-of-sync`, `--fail-on`
* Web UI for SCM scans: **Settings** > **Language** > **JavaScript**

### Support for npm

For all supported lockfile versions, the following features are available:

* CLI support
* SCM support
* License scanning
* Fix PRs

Snyk can build a dependency tree with or without a lockfile. If a lockfile is present, Snyk uses it as follows:

* Locally and with CI/CD: if a lockfile is not present and the scan is performed with the CLI or an IDE, Snyk looks at `node_modules` to determine what is installed.
* With SCM integrations: if a lockfile is not present, Snyk approximates what the tree will look like at build time. This is highly valuable for getting insights into Projects in development or what the next build will look like when there is no lockfile present

As a user of npm, even though npm-audit is at hand anytime you are working with your dependencies, Snyk provides the following capabilities. These are designed for both individuals and companies:

* It helps secure not only open source, but also your first-party code. If you are using infrastructure as code and/or containers, Snyk also provides visibility and remediation advice.
* In the context of Open Source:
  * You receive all the benefits of the curation, updates, and additional value that the Snyk Security Team adds, such as Known Exploit and Trending on X.
  * You have Automated Remediation.
* Central reporting
* Git Code repository integration, but not just there, Snyk has integrations across your pipeline and visibility into production.
* Broad support across programming languages and package managers.
* Ignore capabilities.

#### **Peer dependencies**

In npm v7 and above, the behavior of peer dependencies changes if they are being installed by default. To match this in npm v7+ Projects, Snyk assumes peer dependencies are installed and scans them by default.

An npm v7+ Project ignores peer dependencies only if they are explicitly marked as optional in the `peerDependenciesMeta` object in the `package.json` as shown here for `cache-manager`:

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

Snyk uses the `package-lock.json` lockfile when present to generate a dependency tree for your Project. These lockfiles come in different versions.

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

### Support for pnpm

For all supported pnpm versions, the following features are available:&#x20;

* CLI support
* SCM support
* License scanning
* Fix PRs

For Snyk to detect a Project as using pnpm, the Project must include the relevant lockfile and configuration files:

* Standard pnpm Projects must contain both `package.json` and `pnpm-lock.yaml` in the same directory.
* pnpm workspaces must include `package.json`, `pnpm-lock.yaml`, and `pnpm-workspace.yaml` in the root directory. Additionally, each package within the workspace must have its own `package.json`.

If the mentioned pnpm lockfile is not present, Snyk treats the Project as an `npm` Project and scans it according to the details mentioned above, for `npm`.

**Lockfile versions**

Snyk uses the `pnpm-lock.yaml` lockfile to generate a dependency tree for your Project.&#x20;

The supported lockfile versions are 5.4, 6.x and 9.x, as used by pnpm 7, 8, 9 and 10.

pnpm lockfiles do not include [bundledDependencies](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#bundledependencies), so Snyk does not include them in scans.

### **Support for Yarn**

Snyk uses the Yarn lockfile (`yarn.lock`) to generate a representation of Project dependencies.

The files Snyk relies on to scan a Project may change on version upgrades of the package manager. Snyk lists only versions verified internally as supported.

If you are using a newer version of Yarn than is not listed on this page, it is possible that Snyk performs as expected because Yarn is using a lockfile version that is already supported. That version of Yarn has likely not been evaluated and, thus not added to this page.&#x20;

For all supported Yarn versions, the following features are available:&#x20;

* CLI support
* SCM support
* License scanning
* Fix PRs

{% hint style="info" %}
Because different versions of Yarn have different feature sets, there are differences in Snyk support in order to match how the package manager works.

Resolutions are supported in Yarn v2 and above. Yarn v1 resolutions are not supported.
{% endhint %}

If a lock file is not present in CLI, the `node_modules` folder will be used to construct the dependency tree.

`nohoist` is not supported for Yarn Workspaces.

### Support for Lerna

Snyk does not fully support Lerna. If your Project is set up using Yarn Workspaces, you can scan the Project in the same way you scan any Yarn Workspaces Project.

If your Lerna Project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` as follows.

For each example package, you can use the following command:

<pre class="language-shellscript"><code class="lang-shellscript"><strong>snyk monitor --file=packages/example-package/package.json
</strong></code></pre>

Alternatively, you can specify a script to automate scanning of nested `package.json` files:

```shellscript
ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json
```

### Scanning using npm, pnpm and Yarn

The following table lists the steps to start scanning your dependencies. It covers basic commands, such as `snyk test` and `snyk monitor`. For a full list of CLI commands, see the [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md).

<table><thead><tr><th width="153.43489583333331">Package manager</th><th>Getting started</th><th>Description</th></tr></thead><tbody><tr><td>npm</td><td><ol><li>Install npm.</li><li>Ensure you are in a directory with npm Project files, that is, <code>package.json</code> and <code>package-lock.json</code>.</li><li>(Optional) Run <code>npm install</code>.</li><li>Run <a href="../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../../developer-tools/snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> and <a href="../../../developer-tools/snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>.</li></ol></td><td><p>Snyk analyzes your <code>package-lock.json</code> files to build a dependency tree. </p><p></p><p>If the <code>package-lock.json</code> is missing, Snyk analyzes your <code>node_modules</code> folder. </p><p></p><p>Alternatively, run <code>npm install</code> to generate the lockfile first.</p></td></tr><tr><td>pnpm</td><td><ol><li>Install pnpm.</li><li>Ensure that you are in a directory with pnpm Project files, that is, <code>package.json</code> or <code>pnpm</code>and <code>pnpm-lock.yaml</code>.</li><li>(Optional)  Run <code>pnpm install</code>.</li><li>Run <a href="../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../../developer-tools/snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> and <a href="../../../developer-tools/snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>.</li></ol></td><td>Snyk analyzes your<code>pnpm-lock.yaml</code> files to build a dependency tree. <br><br>If the <code>pnpm-lock.yaml</code> is missing, Snyk analyzes your <code>node_modules</code> folder. <br><br>Alternatively, run <code>pnpm install</code> to generate the lockfile first.</td></tr><tr><td>Yarn</td><td><ol><li>Install Yarn.</li><li>Ensure you are in a directory with Yarn Project files, that is, <code>package.json</code> and <code>yarn.lock</code>.</li><li>(Optional) Run <code>yarn install</code></li><li>Run <a href="../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../../developer-tools/snyk-cli/commands/test.md#options-for-yarn-projects">snyk test</a> and <a href="../../../developer-tools/snyk-cli/commands/monitor.md#options-for-yarn-projects">snyk monitor</a>.</li></ol></td><td><p>Snyk analyzes your <code>yarn.lock</code> files to build a dependency tree. </p><p></p><p>If the <code>yarn.lock</code> is missing, Snyk analyzes your <code>node_modules</code> folder. </p><p></p><p>Alternatively, run <code>yarn install</code> to generate the lockfile first.</p></td></tr></tbody></table>

### Support for monorepos and workspaces

Yarn, npm, and pnpm support workspaces, to help manage monorepos containing multiple sub-Projects.

Workspaces are supported in the Snyk CLI for the following CLI options:

* `--all-projects` : Discovers and scan all Yarn,  npm and pnpm workspaces Projects, along with Projects from other supported ecosystems. The root lock file is referenced when scanning the workspace Projects.
* `--detection-depth` : Specifies how many sub-directory levels to search.
* `--strict-out-of-sync=false` :  Allows testing out-of-sync lockfiles for packages in a  workspace. When this option is set to `false` , you can run Snyk tests with unsynchronized manifest and lock files without causing errors.
* `--policy-path` : Specifies the path to a policy used by Snyk during testing.

{% hint style="info" %}
Yarn and npm workspaces are not explicitly supported in Snyk SCM integrations scans.

pnpm workspaces must have the `package.json`, `pnpm-lock.yaml` and `pnpm-workspace.yaml` files in the root directory. Fix and Upgrade PRs do not support lockfile update for pnpm workspaces.
{% endhint %}

Root-level `package.json` manifest files with adjacent lockfiles are scanned as normal.&#x20;

For nested manifest files with no lockfiles, Snyk approximates what the dependency tree looks like at build time without using the root lockfile.&#x20;

In addition, Fix PRs and Upgrade PRs are not supported for nested manifest files with no lockfiles.

#### Examples of scanning workspaces

To scan all workspaces Projects in the current directory and five sub-directories deep, plus any other Projects types detected, use the following command:

```bash
snyk test --all-projects --strict-out-of-sync=false --detection-depth=6 
```

Use a common `.snyk` policy file, if you maintain ignores and patches in one place to be applied for all detected workspaces by using the policy path. See [The .snyk file](../../../manage-risk/policies/the-.snyk-file.md).

```bash
snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk
```

#### **Scanning npm workspaces**

npm v7 introduced support for workspaces. To detect and scan all workspaces in your npm Project, use the CLI options described above.

#### **Scanning pnpm workspaces**

pnpm workspaces must have the `package.json`, `pnpm-lock.yaml` and `pnpm-workspace.yaml` files in the root directory. To detect and scan all workspaces in your pnpm Project, use the CLI options described above.

#### **Scanning Yarn workspaces**

{% hint style="info" %}
`nohoist` is not supported for Yarn Workspaces.
{% endhint %}

To detect and scan all workspaces in your Yarn Project, use the CLI options identified for monorepos and workspaces, as well as this Yarn-specific option.

`--yarn-workspaces` : Uses instead of `--all-projects` to detect and scan only Yarn workspaces Projects when a lockfile is present in the root. Other ecosystems will be ignored.

## Validating, monitoring, alerting, and gating for JavaScript

For SCM integrations with a Snyk Enterprise plan only, Snyk can monitor container images and their open source or Linux-based packages being used in production using Kubernetes integration to notify customers of known vulnerabilities for applications in production.

For all Snyk plans, where a production integration does not exist, use the [`snyk monitor`](../../../developer-tools/snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.
