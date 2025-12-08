# TypeScript

## Applicability and integration

{% hint style="info" %}
TypeScript is supported for Snyk Open Source and Snyk Code.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

Available functions:

* Test your app's SBOM: using `pkg:npm`
* Test your app's packages using `pkg:npm`

## Technical specifications

You can use TypeScript versions up to 4.2

### Supported frameworks and libraries

Snyk supports all JavaScript frameworks and libraries. For more information, see [JavaScript frameworks and libraries](javascript/#frameworks-and-libraries).

### Supported package managers and registries

For TypeScript, Snyk supports npm, pnpm, Yarn as package managers, with the following versions:

* npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3`
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

As a package registry, Snyk supports [npmjs.org](https://www.npmjs.org/).

## TypeScript for Snyk Code

For TypeScript with Snyk Code, the following file formats are supported: `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

Available features:

* Reports
* Custom rules
* Interfile analysis

## TypeScript for Snyk Open Source

For TypeScript with Snyk Open Source, the following file formats are supported:

* npm: `package.json` and `package-lock.json`
* pnpm: `pnpm-lock.yaml`
* Yarn: `yarn.lock`

Available features:

* License scanning
* Reports

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for TypeScript. This means that you will not be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}
