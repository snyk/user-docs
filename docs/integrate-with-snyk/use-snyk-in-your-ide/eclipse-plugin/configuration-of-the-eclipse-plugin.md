# Configuration of the Eclipse plugin

You can set the following configuration options in the Snyk preferences.

<figure><img src="../../../.gitbook/assets/image (2).png" alt="Snyk preferences"><figcaption><p>Snyk preferences</p></figcaption></figure>

* `Snyk API Token`: Set the authentication token from Snyk.
* `Path`: Specify your additions to the path to find needed third-party tools such as Gradle or Maven.
* `Custom Endpoint`: Specify the custom endpoint for Single Tenant setups instead of `https://app.snyk.io/api`.
* `Allow unknown certificate authorities`: Disable certificate checks for SSL connections.
* `Snyk Open Source enabled`: Enable or Disable Snyk Open Source Dependency Scans through Language Server. Default: `Enabled` during beta
* `Snyk Code enabled`: Enable or Disable Snyk Code Scans via Language Server. Default: `Disabled` during beta.
* `Snyk Infrastructure-as-Code enabled` : Enable or Disable Snyk IaC Scans through Language Server. Default: `Enabled` during beta.
* `Organization`: Specify the Snyk Organization to use for scanning. Snyk recommends using the `ORG_ID`. If you specify an Organization slug name, the value must match the URL slug (`[orgslugname]`) as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`.
* `Additional Parameters`: Specify additional parameters to pass to the CLI , for example, `--all-projects` or `--debug.`
* `Additional Environment`: Add environment variables to Language Server; multiple variable can be separated by `;`. Example: `JAVA_HOME=/usr/local/bin;GOPATH=/usr/local/bin`
* `Update and install Snyk binaries automatically`: If `disabled`, no updates are downloaded, and updates must be performed manually. Snyk recommends always using the most recent version of the CLI. Ensure that the location for the CLI points to an existent, current binary.
* `Base URL for CLI download:` Specify an alternative download host for the CLI, for example,  `https://static.snyk.io/fips`. This must provide the CLI and necessary files as the default `https://static.snyk.io` does, that is, the following files. See also  GitHub[ releases](https://github.com/snyk/cli/releases).
  * %Base URL%/cli/v%VERSION%/%CLI-BINARY-NAME%
  * %Base URL%/cli/v%VERSION%/%CLI-BINARY-NAME%.sha256
  * %Base URL%/cli/v%VERSION%/sha256sums.txt.asc
  * %Base URL%/cli/v%VERSION%/release.json
  * %Base URL%/cli/latest/version
  * %Base URL%/cli/latest/%CLI-BINARY-NAME%
  * %Base URL%/cli/latest/%CLI-BINARY-NAME%.sha256
  * %Base URL%/cli/latest/ls-protocol-version-%PROTOCOL\_VERSION%
  * %Base URL%/cli/latest/release.json
  * %Base URL%/cli/latest/sha256sums.txt.asc
* `Snyk CLI`: Specify the location of the Snyk CLI.
* `Send error reports to Snyk`: Send errors from Language Server to Snyk to enable quick bug fixing. Default: `Enabled`.
* `Send usage statistics to Snyk`: Allow Snyk to get usage data to improve workflows. Default: `Enabled`.
* `Trusted Folders`:  specify, which directories should be considered safe, e.g. the parent directory of all your projects.

###
