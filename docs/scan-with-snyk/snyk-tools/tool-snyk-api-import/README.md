# Tool: snyk-api-import

Snyk helps you find, fix, and monitor for known vulnerabilities in your dependencies, both on an ad hoc basis and as part of your Continuous Integration (CI) (build) system.

## snyk-api-import

The Snyk API Project importer, `snyk-api-import`, is a script intended to help import Projects into Snyk using available [Snyk APIs](../../../snyk-api/reference/) at a controlled pace to avoid rate limiting from GitHub, GitLab, Bitbucket, and other systems and to provide a stable import. The script kicks off import in batches, waits for completion, and then keeps going. Any failed requests are retried before they are considered a failure and logged.

If you need to adjust concurrency, you can stop the script, change the concurrency variable, and start again. The tool skips previous repositories (Targets) that have been requested for import.

To use `snyk-api-import` you must do the following in advance:

* Set up your Snyk Organizations (Orgs) before running an import.
* Configure your Snyk Organizations with some connection to an SCM (GitHub, GitLab, Bitbucket, other) as you will need the `integrationId` to generate the import files.
* Use the [Set notification settings](../../../snyk-api/reference/organizations-v1.md#org-orgid-notification-settings) API endpoint to disable notifications for emails and so on, to avoid receiving import notifications (recommended).
* Use the [Update](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-settings) (integration settings) endpoint to disable the fix PRs and PR checks until import is complete to avoid sending extra requests to SCMs (GitHub, GitLab, Bitbucket, and so on).

## Installation

`snyk-api-import` CLI can be installed through multiple channels.

### Standalone executables (macOS, Linux, Windows)

Use the [GitHub Releases](https://github.com/snyk/snyk-api-import/releases) to download a standalone executable of `snyk-api-import` CLI for your platform.

### Install with npm or Yarn

[snyk-api-import CLI is available as an npm package](https://www.npmjs.com/package/snyk-api-import). If you have Node.js installed locally, you can install it by running:

```
npm install snyk-api-import@latest -g
```

If you are using Yarn, run:

```
yarn global add snyk-api-import
```

## Usage

By default the `import` command will run if no command is specified.

* `import` - kick off an API-powered import of repos (Targets) into existing Snyk Organizations defined in the import configuration file. All support available for all Project types is provided through the [Import](../../../snyk-api/reference/import-projects-v1.md) API endpoints, [Import targets](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import) and [Get Import job details](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import-jobid). [Import API (Import Projects, Import)](https://snyk.docs.apiary.io/#reference/import-projects).
* `help` - show help and all available commands and their options.
* `orgs:data` utility - use to generate data required to create Organizations using the API.
* `orgs:create` utility - use to create the Organizations in Snyk based on the data file generated with `orgs:data` command.
* `import:data` utility - use to generate data required to kick off an import. Note that archived repositories are excluded by default.
* `list:imported` utility - use to generate data to help skip previously imported targets during import.

The logs can be explored using the [Bunyan CLI](http://trentm.com/node-bunyan/bunyan.1.html)

## Contents of the snyk-api-import instructions

* Utilities
  * [Creating Organizations in Snyk](creating-organizations-in-snyk.md)
  * [Creating import targets and data for import](creating-import-targets-data-for-import-command.md)
  * [Mirroring GitHub.com and GitHub Enterprise organizations and repos in Snyk](mirroring-github.com-and-github-enterprise-organizations-and-repos-in-snyk.md)
  * [Mirroring GitLab organizations and repos in Snyk](mirroring-gitlab-organizations-and-repos-in-snyk.md)
  * [Mirroring Bitbucket Server organizations and repos in Snyk](mirroring-bitbucket-server-organizations-and-repos-in-snyk.md)
  * [Mirroring Bitbucket Cloud organizations and repos in Snyk](mirroring-bitbucket-cloud-organizations-and-repos-in-snyk.md)
* [Kicking off an import](kicking-off-an-import.md)
* [Contributing to snyk-api-import](contributing-to-snyk-api-import.md)
* [Sync: detecting changes in monitored repos and updating Snyk Projects](https://github.com/snyk/snyk-api-import/blob/master/docs/sync.md)
* Example workflow: [AWS automation](https://github.com/snyk/snyk-api-import/blob/master/docs/example-workflows/aws-automation-example.md)

## FAQ

<details>

<summary><code>Error: ENFILE: file table overflow, open</code> or <code>Error: EMFILE, too many open files</code></summary>

If you see these errors, you may need to bump **ulimit** to allow more open file operations. In order to keep the operations performing well, the tool logs as soon as it is convenient rather than waiting until the very end of a loop and logging a huge data structure. This means that depending on the number of concurrent imports set, the tool may exceed the system default **ulimit**.

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

Yes. Because Snyk reuses the existing integration with your SCM (Git) repository to perform the imports, the brokered connection will be used when configured.

</details>

<details>

<summary>What is supported for the import command?</summary>

`snyk-api-import` supports all of the same integration types and Project sources as identified in the [Import API](../../../snyk-api/reference/import-projects-v1.md) documentation. If an example for your use case is not in these instructions, see the API documentation.

</details>
