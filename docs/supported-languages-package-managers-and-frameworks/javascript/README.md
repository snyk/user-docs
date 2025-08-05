# JavaScript

Snyk supports [JavaScript for code analysis](javascript-for-code-analysis.md) and [JavaScript for open source](javascript-for-open-source.md). For information about importing Projects using SCM integrations, see [Git repositories and JavaScript](git-repositories-and-javascript.md).

[Guidance for JavaScript and Node.js](best-practices-for-javascript-and-node.js.md) is available.

## Applicability

Snyk supports [JavaScript for code analysis](javascript-for-code-analysis.md) and [JavaScript for open source](javascript-for-open-source.md).

For more information about importing Projects using SCM integrations, see [Git repositories and JavaScript](git-repositories-and-javascript.md).

[Guidance for JavaScript and Node.js](best-practices-for-javascript-and-node.js.md) is available.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM: using `pkg:npm`
* Test your app's packages using `pkg:npm`

For information on using the Snyk CLI for code analysis, see [Snyk CLI for Snyk Code](../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Package managers and supported file extensions

Snyk for JavaScript supports npm, pnpm, and Yarn as package managers with the following versions for them:

* npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`, `pnpm 10`
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

As a package registry, it supports [npmjs.org](https://www.npmjs.org/).

Snyk for JavaScript supports the following file formats:

* Snyk Open Source:
  * For npm package manager: `package.json`, `package-lock.json`
  * For pnpm package manager: `pnpm-lock.yaml`
  * For yarn package manager: `yarn.lock`
* Snyk Code: `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for JavaScript:

* @Google Drive/generative-ai - Comprehensive
* @anthropic-ai/sdk - Comprehensive
* @huggingface/inference - Comprehensive
* @mistralai/mistralai - Comprehensive
* axios - Comprehensive
* Angular - Partial
* apollo-server - Partial
* bcrypt-nodejs - Comprehensive
* cross-spawn - Comprehensive
* crypto-js - Comprehensive
* date-fns - Comprehensive
* dayjs - Comprehensive
* dompurify - Comprehensive
* electron - Partial
* ejs - Partial
* execa - Comprehensive
* express - Comprehensive
* express-mongo-sanitize - Comprehensive
* express-graphql - Partial
* express-jwt - Partial
* fastMCP - Comprehensive
* fs - Comprehensive
* fs-extra - Comprehensive
* fs-plus - Comprehensive
* graceful-fs - Comprehensive
* graphql-js - Partial
* grpc-js - Comprehensive
* hapi.js - Comprehensive
* jQuery - Comprehensive
* js-yaml - Comprehensive
* jzip - Comprehensive
* koa - Comprehensive
* koa-graphql - Comprehensive
* libxml - Comprehensive
* libxmljs - Comprehensive
* lodash - Comprehensive
* luxon - Comprehensive
* minimongo - Comprehensive
* minimist - Comprehensive
* modelcontextprotocol/typescript-sdk - Comprehensive
* mongodb - Comprehensive
* Mongoose - Comprehensive
* mercurius - Partial
* Nestjs - Partial
* Node Crypto - Comprehensive
* node-buffer - Partial
* node-cmd - Comprehensive
* Node Crypto - Comprehensive
* node-dir - Comprehensive
* node-forge - Comprehensive
* node-pty - Comprehensive
* node-serialize - Comprehensive
* octokit - Comprehensive
* openai - Comprehensive
* pg - Comprehensive
* pg-promise - Comprehensive
* React - Partial
* request-promise - Comprehensive
* restler - Partial
* rimraf - Comprehensive
* sanitize-html - Comprehensive
* shelljs - Comprehensive
* Stanford JS Crypto - Comprehensive
* superagent - Comprehensive
* tar-stream - Comprehensive
* TSOA - Comprehensive
* unirest - Comprehensive
* unzip - Comprehensive
* underscore - Comprehensive
* url - Comprehensive
* vm - Comprehensive
* webstomp-client - Partial
* WebCryptoAPI - Comprehensive
* xpath - Comprehensive
* yargs - Comprehensive

## Features

The following features are supported in Snyk for JavaScript:

| Snyk Open Source                                                   | Snyk Code                                                                 |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------- |
| <ul><li>Fix PRs</li><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis</li></ul> |
