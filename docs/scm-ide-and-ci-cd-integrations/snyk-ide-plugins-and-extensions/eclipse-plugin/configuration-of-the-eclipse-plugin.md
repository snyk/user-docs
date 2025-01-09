# Configuration of the Eclipse plugin

You can configure both [Global settings](configuration-of-the-eclipse-plugin.md#global-settings) and [Project-specific propertie](configuration-of-the-eclipse-plugin.md#project-specific-properties).

## Global settings

You can set the following global configuration settings in the Snyk preferences.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-01-09 at 8.35.31 AM.png" alt=""><figcaption><p>Snyk preferences</p></figcaption></figure>

* `Token`: Set the authentication token for Snyk.
* `Use token authentication`: Select to override the default OAuth2 authentication in order to use a Snyk API token. Snyk recommends keeping this setting off because the default OAuth2 authentication is more secure.
* `Path`: Specify your additions to the path to find needed third-party tools such as Gradle or Maven.
* `Custom Endpoint`: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls). Multi-tenant users who do not belong to the default region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* `Allow unknown certificate authorities`: Disable certificate checks for SSL connections.
* `Snyk Open Source enabled`: Enable or disable Snyk Open Source Dependency Scans through the Language Server. Default: `Enabled` during beta.
* `Snyk Code Security enabled`: Enable or disable Snyk Code Security Issues through the Language Server. Default: `Disabled` by default.
* `Snyk Code Quality enabled`: Enable or Disable Snyk Code Quality Issues through the Language Server. Default: `Disabled` by default.
* `Snyk Infrastructure-as-Code enabled`: Enable or disable Snyk IaC scans through the Language Server. Default: `Enabled` during beta.
* `Scan automatically on start-up and save`: Scan automatically or not when you start the extension.
* `Organization`: Specify the Snyk Organization to use for scanning. Snyk recommends using the `ORG_ID`. If you specify an Organization slug name, the value must match the URL slug (`[orgslugname]`) as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`.
* `Additional Parameters`: Specify additional parameters to pass to the CLI, for example, `--file=pom.xml` or `--debug.` \
  **Note:** When you enable `debug`, your code may be logged in the IDE log files, for example, the `idea.log` file.
* `Additional Environment`: Add environment variables to the Language Server; multiple variables can be separated by `;`. Example: `JAVA_HOME=/Library/JDK/bin;GOPATH=/usr/local/bin`
* `Update and install Snyk binaries automatically`: If `disabled`, no updates are downloaded, and updates must be performed manually. Snyk recommends always using the most recent version of the CLI. Ensure that the location for the CLI points to an existing current binary.
* `Base URL for CLI download:` Specify an alternative download host for the CLI, for example, `https://downloads.snyk.io/fips`. This must provide the CLI and necessary files as the default `https://downloads.snyk.io` does, that is, the following files. See also GitHub[ releases](https://github.com/snyk/cli/releases).
  * %Base URL%/cli/v%VERSION%/%CLI-BINARY-NAME%
  * %Base URL%/cli/v%VERSION%/%CLI-BINARY-NAME%.sha256
  * %Base URL%/cli/v%VERSION%/sha256sums.txt.asc
  * %Base URL%/cli/v%VERSION%/release.json
  * %Base URL%/cli/stable/version
  * %Base URL%/cli/stable/%CLI-BINARY-NAME%
  * %Base URL%/cli/stable/%CLI-BINARY-NAME%.sha256
  * %Base URL%/cli/stable/ls-protocol-version-%PROTOCOL\_VERSION%
  * %Base URL%/cli/stable/release.json
  * %Base URL%/cli/stable/sha256sums.txt.asc
* `Snyk CLI`: Specify the location of the Snyk CLI, where it is searched for, and where it is downloaded to, if automatic management of Snyk binaries is enabled.
* `Send error reports to Snyk`: Send errors from the Language Server to Snyk to enable quick bug fixing. Default: `Enabled`.
* `Send usage statistics to Snyk`: Allow Snyk to get usage data to improve workflows. Default: `Enabled`.
* `Trusted Folders`: Specify, which directories should be considered safe, for example, the parent directory of all your Projects.

## Project-specific properties

When you select Project **Properties**, you can specify project-specific scan settings on the Snyk properties page.

<figure><img src="../../../.gitbook/assets/image (652).png" alt=""><figcaption><p>Project properties can be opened e.g. via the context menu of the project</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (651).png" alt=""><figcaption><p>The Snyk Project-specific properties page allows entering scan parameters for the Project.</p></figcaption></figure>

