---
nav_context: agnostic
description: >-
  Snyk support for JavaScript with Snyk Code and Snyk Open Source, including
  supported frameworks, libraries, and package managers
---

# JavaScript

{% hint style="info" %}
JavaScript is supported for Snyk Code and Snyk Open Source.
{% endhint %}

## JavaScript for Snyk Code

For an overview of the supported security rules, visit [JavaScript and TypeScript rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/javascript-and-typescript-rules).

### Supported frameworks and libraries

Snyk Code supports the following frameworks and libraries:

{% columns %}
{% column %}
* @google/generative-ai
* @anthropic-ai/sdk
* @huggingface/inference
* @mistralai/mistralai
* axios
* Angular
* apollo-server
* bcrypt-nodejs
* Bootbox
* cross-spawn
* crypto-js
* d3
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
* NestJS
* Node Crypto
* node-buffer
* node-cmd
* node-dir
* node-forge
* node-pty
* node-serialize
* octokit
* openai
* pg
* pg-promise
* React (partial)
* request-promise
* restler
* rimraf
* sanitize-html
* shelljs
* Stanford JavaScript Crypto
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

Snyk Code supports the following file formats: `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`, `.erb`.

To resolve imports that use path aliases when tracing data flow across modules, Snyk Code reads `tsconfig.json` and `jsconfig.json` in your Project.

### Available features

* Reports
* Interfile analysis

### Template file analysis

Snyk Code parses template files and follows data from your application code into the template, so it reports a cross-site scripting vulnerability that becomes exploitable only where the template renders its output.

For JavaScript, Snyk Code supports the following template engines:

* EJS
* Handlebars
* Mustache
* Pug

For file extensions and the full set of supported language and template engine combinations, visit [Template file analysis](../../technical-specifications-and-guidance.md#template-file-analysis).

## JavaScript for Snyk Open Source

### Supported package managers and package registries

Snyk supports the following package managers and versions:

*   npm: `npm 5`, `npm 6`, `npm 7`, `npm 8`, `npm 9`, `npm 10+`

    Supported lockfile versions: `Lockfile v1`, `Lockfile v2`, `Lockfile v3`
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`, `pnpm 10`
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`, `Yarn 4`

The default package registry for Snyk is [npmjs.org](https://www.npmjs.org/). Snyk supports private package registries. For more information, visit [Package repository integrations](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations).

### Available integrations

* SCM import
* CLI and IDE: test or monitor your app

### Supported file formats

Snyk Open Source supports the following file formats:

* For npm: `package.json` and `package-lock.json`
* For pnpm: `pnpm-lock.yaml`
* For Yarn: `yarn.lock`

Lerna is partially supported.

### Available features

* Automatic and manual Fix PRs (for npm, pnpm, and Yarn)
* License scanning
* Reports
* Test your app's software bill of materials (SBOM) and packages using `pkg:npm` PURLs through the [SBOM test](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/sbom-test) CLI command

### Language and package manager considerations

{% hint style="info" %}
Snyk tracks only official releases. Snyk does not identify commits, including commits to the default branch, unless they are included in an official release or tag.

For JavaScript packages, this means a release to the npmjs.org package registry.
{% endhint %}

#### devDependencies analysis

Snyk disables `devDependencies` analysis by default because these dependencies are not typically elevated to production, and security and development teams often consider them noise. To enable testing of dev dependencies:

* Use the `--dev` parameter for CLI and continuous integration/continuous delivery (CI/CD) integrations.
* For SCM integrations, configure the relevant setting in **Settings** > **Languages**.

#### optionalDependencies analysis

Snyk includes `optionalDependencies` by default for CLI, CI/CD, and SCM integrations.

#### Unmanaged JavaScript

If you have an Enterprise plan and therefore have access to the Snyk API, you can use the API to get a full list of dependencies and their transitive dependencies.

To test for vulnerabilities, you can use the following API endpoints:

* [Test for issues in a public package by name and version](https://docs.snyk.io/developer-tools/snyk-api/reference/test-v1#test-npm-packagename-version)
* [Test Dep Graph](https://docs.snyk.io/developer-tools/snyk-api/reference/test-v1#test-dep-graph)
* [List issues for a package](https://docs.snyk.io/developer-tools/snyk-api/reference/issues#orgs-org_id-packages-purl-issues)

#### Out of sync lockfiles

To control behavior when the lockfile and package file are out of sync, use the following:

* CLI additional values: `--strict-out-of-sync`, `--fail-on`
* Web UI for SCM scans: **Settings** > **Language** > **JavaScript**

### Support for npm

For all supported lockfile versions, the following features are available:

* CLI support
* SCM support
* License scanning
* Automatic and manual Fix PRs

Snyk can build a dependency tree with or without a lockfile. If a lockfile is present, Snyk uses it. If a lockfile is not present, Snyk builds the tree as follows:

* Locally and with CI/CD: if you scan with the CLI or an IDE, Snyk looks at `node_modules` to determine what is installed.
* With SCM integrations: Snyk approximates what the tree looks like at build time. This is valuable for getting insights into Projects in development, or into what the next build looks like.

Although `npm audit` is available whenever you work with your dependencies, Snyk provides the following capabilities for both individuals and companies:

* Security for your first-party code, not only open source. If you use infrastructure as code or containers, Snyk also provides visibility and remediation advice.
* For open source:
  * The curation, updates, and additional value that the Snyk Security Team adds, such as Known Exploit and Trending on X
  * Automated remediation
* Central reporting
* SCM integrations, plus integrations across your pipeline and visibility into production
* Broad support across programming languages and package managers
* Ignore capabilities

#### Peer dependencies

In npm v7 and later, npm installs peer dependencies by default. To match this in npm v7+ Projects, Snyk assumes peer dependencies are installed and scans them by default.

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

In npm v6 and earlier, Snyk does not scan peer dependencies by default, because the package manager does not install them by default. To scan peer dependencies, ensure they are installed, and then run the CLI with the `--peer-dependencies` option.

#### Lockfile versions

Snyk uses the `package-lock.json` lockfile when present to generate a dependency tree for your Project. These lockfiles come in different versions.

Lockfile v1 was used in npm v5 and v6. In npm v7, npm introduced two new formats: lockfile v2 and lockfile v3. For more information, visit [lockfileVersion](https://docs.npmjs.com/cli/v9/configuring-npm/package-lock-json#lockfileversion).

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

For all supported pnpm versions, the following features are available:

* CLI support
* SCM support
* License scanning
* Fix PRs

For Snyk to detect a Project as using pnpm, the Project must include the relevant lockfile and configuration files:

* Standard pnpm Projects must contain both `package.json` and `pnpm-lock.yaml` in the same directory.
* pnpm workspaces must include `package.json`, `pnpm-lock.yaml`, and `pnpm-workspace.yaml` in the root directory. Additionally, each package within the workspace must have its own `package.json`.

If the pnpm lockfile is not present, Snyk treats the Project as an `npm` Project and scans it as described earlier for `npm`.

Snyk uses the `pnpm-lock.yaml` lockfile to generate a dependency tree for your Project.

The supported lockfile versions are 5.4, 6.x, and 9.x, as used by pnpm 7, 8, 9, and 10.

Because pnpm lockfiles do not include [bundledDependencies](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#bundledependencies), Snyk does not include them in scans.

### Support for Yarn

Snyk uses the Yarn lockfile (`yarn.lock`) to generate a representation of Project dependencies.

The files Snyk relies on to scan a Project can change when the package manager is upgraded. Snyk lists only versions that Snyk has verified as supported.

If you use a newer version of Yarn that is not listed on this page, it is possible that Snyk performs as expected because that Yarn version uses a lockfile version that Snyk already supports. Snyk has likely not evaluated that version of Yarn, so it is not listed on this page.

For all supported Yarn versions, the following features are available:

* CLI support
* SCM support
* License scanning
* Automatic and manual Fix PRs

{% hint style="info" %}
Because different versions of Yarn have different feature sets, there are differences in Snyk support to match how the package manager works.

Snyk supports resolutions in Yarn v2 and later. Snyk does not support Yarn v1 resolutions.
{% endhint %}

Snyk does not support `nohoist` for Yarn Workspaces.

If the Yarn lockfile is not present, Snyk treats the Project as an `npm` Project and scans it as described earlier for `npm`.

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

### Scanning using npm, pnpm, and Yarn

The following table lists the steps to start scanning your dependencies. It covers basic commands, such as `snyk test` and `snyk monitor`. For a full list of CLI commands, visit the [CLI commands and options summary](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/cli-commands-and-options-summary).

<table><thead><tr><th width="153.43489583333331">Package manager</th><th>Getting started</th><th>Description</th></tr></thead><tbody><tr><td>npm</td><td><ol><li>Install npm.</li><li>Ensure you are in a directory with npm Project files, that is, <code>package.json</code> and <code>package-lock.json</code>.</li><li>(Optional) Run <code>npm install</code>.</li><li>Run <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/cli-commands-and-options-summary">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/test#options-for-npm-projects">snyk test</a> and <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/monitor#options-for-npm-projects">snyk monitor</a>.</li></ol></td><td><p>Snyk analyzes your <code>package-lock.json</code> files to build a dependency tree.</p><p>If the <code>package-lock.json</code> is missing, Snyk analyzes your <code>node_modules</code> folder.</p><p>Alternatively, run <code>npm install</code> to generate the lockfile first.</p></td></tr><tr><td>pnpm</td><td><ol><li>Install pnpm.</li><li>Ensure that you are in a directory with pnpm Project files, that is, <code>package.json</code> and <code>pnpm-lock.yaml</code>.</li><li>(Optional) Run <code>pnpm install</code>.</li><li>Run <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/cli-commands-and-options-summary">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/test#options-for-npm-projects">snyk test</a> and <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/monitor#options-for-npm-projects">snyk monitor</a>.</li></ol></td><td>Snyk analyzes your <code>pnpm-lock.yaml</code> files to build a dependency tree.<br><br>If the <code>pnpm-lock.yaml</code> is missing, Snyk analyzes your <code>node_modules</code> folder.<br><br>Alternatively, run <code>pnpm install</code> to generate the lockfile first.</td></tr><tr><td>Yarn</td><td><ol><li>Install Yarn.</li><li>Ensure you are in a directory with Yarn Project files, that is, <code>package.json</code> and <code>yarn.lock</code>.</li><li>(Optional) Run <code>yarn install</code>.</li><li>Run <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/cli-commands-and-options-summary">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/test#options-for-yarn-projects">snyk test</a> and <a href="https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/monitor#options-for-yarn-projects">snyk monitor</a>.</li></ol></td><td><p>Snyk analyzes your <code>yarn.lock</code> files to build a dependency tree.</p><p>If the <code>yarn.lock</code> is missing, Snyk analyzes your <code>node_modules</code> folder.</p><p>Alternatively, run <code>yarn install</code> to generate the lockfile first.</p></td></tr></tbody></table>

### Support for monorepos and workspaces

Yarn, npm, and pnpm support workspaces to help manage monorepos that contain multiple sub-Projects.

#### SCM scanning considerations

Snyk SCM integration scans do not explicitly support npm workspaces. Snyk scans root-level `package.json` manifest files with adjacent lockfiles as normal.

For nested manifest files with no lockfiles, Snyk approximates what the dependency tree looks like at build time without using the root lockfile.

Yarn workspaces Projects must have the `package.json` and `yarn.lock` files in the root directory.

Yarn 1.x workspaces fail to import using SCM and return an `Out of sync package.json and package-lock.json detected` error. This occurs with internal workspace packages when the `yarn.lock` file is at the repository root, but the `package.json` file is in a subdirectory.

As an SCM workaround, in the Snyk Web UI, navigate to **Settings** > **Snyk Open Source** > **Languages** > **JavaScript** and clear the **Require package.json and package-lock.json/yarn.lock files to be in sync** check box.

As a workaround using the CLI, use the `--strict-out-of-sync=false` flag to allow testing without causing errors.

{% hint style="info" %}
Upgrading to Yarn 2+ does not resolve this issue for SCM scans. SCM workspace Projects require both the `package.json` and `yarn.lock` files to be in the root directory.
{% endhint %}

For pnpm workspaces, the `package.json`, `pnpm-lock.yaml`, and `pnpm-workspace.yaml` files must be in the root directory.

SCM scans do not support the pnpm [workspace protocol](https://pnpm.io/workspaces#workspace-protocol-workspace).

Define dependencies explicitly with specific versions, or with versions that use standard semver, for example, `"foo": "^1.1.0"`.

SCM scans list dependencies that use the workspace protocol for the version, for example, `"foo" : "workspace:*"`, as having an undefined version.

For all workspaces, Fix PRs and Upgrade PRs do not support workspaces lockfile updates. PRs for these Projects update only the `package.json`.

#### CLI scanning considerations

Workspaces are supported in the Snyk CLI for the following CLI options:

* `--all-projects`: Discovers and scans all Yarn, npm, and pnpm workspaces Projects, along with Projects from other supported ecosystems. Snyk references the root lockfile when scanning the workspace Projects.
* `--detection-depth`: Specifies how many sub-directory levels to search.
* `--strict-out-of-sync=false`: Allows testing out-of-sync lockfiles for packages in a workspace. When this option is set to `false`, you can run Snyk tests with unsynchronized manifest and lock files without causing errors.
* `--policy-path`: Specifies the path to a policy used by Snyk during testing.

#### Examples of scanning workspaces with the CLI

To scan all workspaces Projects in the current directory and five sub-directories deep, plus any other Projects types detected, use the following command:

```bash
snyk test --all-projects --strict-out-of-sync=false --detection-depth=6 
```

If you maintain ignores and patches in one place to apply to all detected workspaces, use a common `.snyk` policy file with the policy path. For more information, visit [The .snyk file](https://docs.snyk.io/scan-fix-and-prevent/prevent/policies/the-.snyk-file).

```bash
snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk
```

#### Scanning npm workspaces

Support for workspaces was introduced in npm v7. To detect and scan all workspaces in your npm Project, use the CLI options described earlier.

#### Scanning pnpm workspaces

For pnpm workspaces, the `package.json`, `pnpm-lock.yaml`, and `pnpm-workspace.yaml` files must be in the root directory. To detect and scan all workspaces in your pnpm Project, use the CLI options described earlier.

#### Scanning Yarn workspaces

{% hint style="info" %}
Snyk does not support `nohoist` for Yarn Workspaces.
{% endhint %}

To detect and scan all workspaces in your Yarn Project, use the CLI options identified for monorepos and workspaces, and the following Yarn-specific option:

`--yarn-workspaces`: Use instead of `--all-projects` to detect and scan only Yarn workspaces Projects when a lockfile is present in the root. Snyk ignores other ecosystems.

## Validating, monitoring, alerting, and gating for JavaScript

For SCM integrations with a Snyk Enterprise plan only, Snyk can monitor container images and their open-source or Linux-based packages used in production through the Kubernetes integration, and notify you of known vulnerabilities in applications in production.

For all Snyk plans, where a production integration does not exist, use the [`snyk monitor`](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/monitor) CLI command to take a snapshot and monitor what is being pushed to production.
