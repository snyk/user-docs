# Tool: snyk-api-import

Snyk helps you find, fix, and monitor for known vulnerabilities in your dependencies, both on an ad hoc basis and as part of your Continuous Integration (CI) (build) system.

## snyk-api-import

The Snyk API project importer, `snyk-api-import`, is a script intended to help import projects into Snyk using available [Snyk APIs](https://snyk.docs.apiary.io/) at a controlled pace to avoid rate limiting from GitHub, GitLab, Bitbucket, and other systems and to provide a stable import. The script kicks off import in batches, waits for completion, and then keeps going. Any failed requests are retried before they are considered a failure and logged.

If you need to adjust concurrency you can stop the script, change the concurrency variable, and start again. The tool skips previous repos (targets) that have been requested for import.

To use `snyk-api-import` you must do the following in advance:

* Set up your Snyk organizations (Orgs) before running an import.
* Configure your Snyk organizations with some connection to SCM (GitHub, GitLab, Bitbucket, other) as you will need the `integrationId` to generate the import files.
* Disable [notifications](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings) for emails and so on, to avoid receiving import notifications (recommended).
* Disable the [fix PRs and PR checks](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update) until import is complete to avoid sending extra requests to SCMs (GitHub, GitLab, Bitbucket and so on).

## Installation

Snyk `snyk-api-import` CLI can be installed through multiple channels.

### Standalone executables (macOS, Linux, Windows)

Use the [GitHub Releases](https://github.com/snyk-tech-services/snyk-api-import/releases) to download a standalone executable of `snyk-api-import` CLI for your platform.

### More installation methods

<details>

<summary>Install with npm or Yarn</summary>

**Install with npm or Yarn**

[Snyk snyk-api-import CLI is available as an npm package](https://www.npmjs.com/package/snyk-api-import). If you have Node.js installed locally, you can install the package by running:

```bash
npm install snyk-api-import@latest -g
```

**or if you are using Yarn:**

```bash
yarn global add snyk-api-import
```

</details>

## Usage

By default the `import` command will run if no command is specified.

* `import` - kick off a an API powered import of repos (targets) into existing Snyk Orgs defined in the import configuration file. All support available for all project types is provided through the [Import API (Import Projects, Import)](https://snyk.docs.apiary.io/#reference/import-projects).
* `help` - show help and all available commands and their options.
* `orgs:data` util - use to generate data required to create Orgs using the API.
* `orgs:create` util - use to create the Orgs in Snyk based on the data file generated with `orgs:data` command.
* `import:data` util - use to generate data required to kick off an import. Note that archived repos are excluded by default.
* `list:imported` util - use to generate data to help skip previously imported targets during import.

The logs can be explored using [Bunyan CLI](http://trentm.com/node-bunyan/bunyan.1.html)

## Contents of the snyk-api-import instructions

* Utilities
  * [Creating orgs in Snyk](creating-orgs-in-snyk.md)
  * [Generating import data](creating-import-targets-data-for-import.md)
  * [Mirroring GitHub.com and GitHub Enterprise organizations and repos in Snyk](mirroring-github.com-github-enterprise-organizations-and-repos-in-snyk.md)
  * [Mirroring GitLab organizations and repos in Snyk](mirroring-gitlab-organizations-and-repos-in-snyk.md)
  * [Mirroring Bitbucket Server organizations and repos in Snyk](mirroring-bitbucket-server-organizations-and-repos-in-snyk.md)
  * [Mirroring Bitbucket Cloud organizations and repos in Snyk](mirroring-bitbucket-cloud-organizations-and-repos-in-snyk.md)
* [Kicking off an import](kicking-off-an-import.md)
* [Contributing](contributing.md)
* [Sync: detecting changes in monitored repos and updating Snyk projects](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/sync.md)
* Example workflow: [AWS automation](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/example-workflows/aws-automation-example.md)

## FAQ

<details>

<summary><code>Error: ENFILE: file table overflow, open</code> or <code>Error: EMFILE, too many open files</code></summary>

If you see these errors then you may need to bump **ulimit** to allow more open file operations. In order to keep the operations performing well, the tool logs as soon as it is convenient rather than waiting until very end of a loop and logging a huge data structure. This means that depending on the number of concurrent imports set, the tool may exceed the system default **ulimit**.

Some of these resources may help you bump the **ulimit**:

* [ss64.com](https://ss64.com/bash/ulimit.html)
* [StackOverflow](https://stackoverflow.com/questions/45004352/error-enfile-file-table-overflow-scandir-while-run-reaction-on-mac)
* [blog.mact.me](http://blog.mact.me/2014/10/22/yosemite-upgrade-changes-open-file-limit)

</details>

<details>

<summary><code>ERROR: HttpError: request to https://github.private.com failed, reason: self signed certificate in certificate chain</code></summary>

If your GitHub, GitLab, Bitbucket, or Azure instance is using a self-signed certificate, you can configure `snyk-api-import` to use this certificate when calling the HTTPS APIs.

`export NODE_EXTRA_CA_CERTS=./path-to-ca`

</details>

<details>

<summary>Does this work with brokered integrations?</summary>

Yes. because we reuse the existing integration with your SCM (git) repository to perform the imports, the brokered connection will be used when configured.

</details>

<details>

<summary>What is supported for the import command?</summary>

`snyk-api-import` supports all of the same integration types and project sources as identified in the [Import API documentation](https://snyk.docs.apiary.io/#reference/integrations/import-projects/import). If an example for your use case is not in these instructions please see the API documentation.

</details>
