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

For information on using the Snyk CLI for code analysis, see [Snyk CLI for Snyk Code](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Package managers and supported file extensions

Snyk for JavaScript supports npm, pnpm, and Yarn as package managers with the following versions for them:

* npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`, `pnpm 10`
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

As a package registry, it supports [npmjs.org](https://www.npmjs.org/).

Snyk for JavaScript supports the following file formats:

* Snyk Open Source:
  * For npm package manager: `package.json` and `package-lock.json`
  * For pnpm package manager: `pnpm-lock.yaml`
  * For yarn package manager: `yarn.lock`
* Snyk Code: `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for JavaScript:

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
* minimist
* modelcontextprotocol/typescript-sdk
* mongodb
* Mongoose
* mercurius
* mysql2/mysql2-promise
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
* React
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

## Features

The following features are supported in Snyk for JavaScript:

| Snyk Open Source                                                   | Snyk Code                                                                 |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------- |
| <ul><li>Fix PRs</li><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis</li></ul> |
