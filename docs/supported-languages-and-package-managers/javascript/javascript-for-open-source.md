# JavaScript for open source

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## JavaScript support

**Package manager**: npm, Yarn

**Package registry**: [npmjs.org](https://www.npmjs.org)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:npm`

**Test your app's packages**: Available, `pkg:npm`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

**Package manager versions**:&#x20;

* npm
  * `Lockfile 1, Lockfile 2, Lockfile 3, 7.*`\
    For details, see the [Snyk Javascript ](./#npm)page.
* Yarn
  * `Yarn 1, Yarn 2, Yarn 3`. For more information, see the [Snyk Javascript ](./#yarn)page.

## Open source and licensing

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk Open Source provides full support for both npm and Yarn, as outlined below.

#### npm

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

#### **Yarn**

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

#### Partially supported package managers

Snyk currently does not fully support **Lerna**. If your Project is set up using Yarn Workspaces, you can scan the Project in the same way you scan any Yarn Workspaces Project.

If your Lerna Project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` as follows:

For each example package, you can use the following:

<pre class="language-javascript"><code class="lang-javascript"><strong>snyk monitor --file=packages/example-package/package.json
</strong></code></pre>

Alternatively, you can specify a script to automate scanning of nested `package.json` files:

```javascript
ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json
```

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).
