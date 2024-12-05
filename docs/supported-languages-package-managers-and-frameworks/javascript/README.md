# JavaScript

Snyk supports [JavaScript for code analysis](javascript-for-code-analysis.md) and [JavaScript for open source](javascript-for-open-source.md). For information about importing Projects using SCM integrations, see [Git repositories and JavaScript](git-repositories-and-javascript.md).&#x20;

[Guidance for JavaScript and Node.js](best-practices-for-javascript-and-node.js.md) is available.

## Applicability

Snyk supports [JavaScript for code analysis](javascript-for-code-analysis.md) and [JavaScript for open source](javascript-for-open-source.md).&#x20;

For more information about importing Projects using SCM integrations, see [Git repositories and JavaScript](git-repositories-and-javascript.md).

[Guidance for JavaScript and Node.js](best-practices-for-javascript-and-node.js.md) is available.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.&#x20;
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM: using `pkg:npm`&#x20;
* Test your app's packages using `pkg:npm`

For information on using the Snyk CLI for code analysis, see [Snyk CLI for Snyk Code](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Package managers and supported file extensions

Snyk for JavaScript supports npm, pnpm, and Yarn as package managers with the following versions for them:&#x20;

* npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`&#x20;
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`&#x20;
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

As a package registry, it supports [npmjs.org](https://www.npmjs.org/).

Snyk for JavaScript supports the following file formats:

* Snyk Open Source:&#x20;
  * For npm package manager: `package.json`, `package-lock.json`
  * For pnpm package manager: `pnpm-lock.yaml`,&#x20;
  * For yarn package manager: `yarn.lock`
* Snyk Code: `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for JavaScript:&#x20;

* @Google Drive/generative-ai - Comprehensive&#x20;
* @anthropic-ai/sdk - Comprehensive&#x20;
* @huggingface/inference - Comprehensive&#x20;
* @mistralai/mistralai - Comprehensive&#x20;
* axios - Comprehensive&#x20;
* Angular - Partial&#x20;
* apollo-server - Partial&#x20;
* bcrypt-nodejs - Comprehensive&#x20;
* cross-spawn - Comprehensive&#x20;
* crypto-js - Comprehensive&#x20;
* date-fns - Comprehensive&#x20;
* dayjs - Comprehensive&#x20;
* dompurify - Comprehensive&#x20;
* electron - Partial&#x20;
* ejs - Partial&#x20;
* execa - Comprehensive&#x20;
* express - Comprehensive&#x20;
* express-graphql - Partial&#x20;
* express-jwt - Partial&#x20;
* fs - Comprehensive&#x20;
* fs-extra - Comprehensive&#x20;
* fs-plus - Comprehensive&#x20;
* graceful-fs - Comprehensive&#x20;
* graphql-js - Partial&#x20;
* jQuery - Comprehensive&#x20;
* js-yaml - Comprehensive&#x20;
* jzip - Comprehensive&#x20;
* koa - Comprehensive&#x20;
* koa-graphql - Comprehensive&#x20;
* libxml - Comprehensive&#x20;
* libxmljs - Comprehensive&#x20;
* lodash - Comprehensive&#x20;
* luxon - Comprehensive&#x20;
* minimongo - Comprehensive&#x20;
* minimist - Comprehensive&#x20;
* mongodb - Comprehensive&#x20;
* Mongoose - Comprehensive&#x20;
* mercurius - Partial&#x20;
* Nestjs - Partial&#x20;
* Node Crypto - Comprehensive&#x20;
* node-buffer - Partial&#x20;
* node-cmd - Comprehensive&#x20;
* Node Crypto - Comprehensive&#x20;
* node-dir - Comprehensive&#x20;
* node-forge - Comprehensive&#x20;
* node-pty - Comprehensive&#x20;
* node-serialize - Comprehensive&#x20;
* octokit - Comprehensive&#x20;
* openai - Comprehensive&#x20;
* pg - Comprehensive&#x20;
* pg-promise - Comprehensive&#x20;
* React - Partial&#x20;
* request-promise - Comprehensive&#x20;
* restler - Partial&#x20;
* rimraf - Comprehensive&#x20;
* sanitize-html - Comprehensive&#x20;
* shelljs - Comprehensive&#x20;
* Stanford JS Crypto - Comprehensive&#x20;
* superagent - Comprehensive&#x20;
* tar-stream - Comprehensive&#x20;
* unirest - Comprehensive&#x20;
* unzip - Comprehensive&#x20;
* underscore - Comprehensive&#x20;
* url - Comprehensive&#x20;
* vm - Comprehensive&#x20;
* webstomp-client - Partial&#x20;
* WebCryptoAPI - Comprehensive&#x20;
* xpath - Comprehensive&#x20;
* yargs - Comprehensive

## Features

The following features are supported in Snyk for JavaScript:

| Snyk Open Source                                                     | Snyk Code                                                                  |
| -------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| <ul><li>Fix PRs </li><li>License scanning </li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules </li><li>Interfile analysis</li></ul> |

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
