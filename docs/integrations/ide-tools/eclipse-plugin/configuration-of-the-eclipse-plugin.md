# Configuration of the Eclipse plugin

You can set the following configuration options in the Snyk preferences.

<figure><img src="../../../.gitbook/assets/image (314).png" alt=""><figcaption><p>The Snyk preference page</p></figcaption></figure>

* `Snyk API Token`: Set the authentication token from Snyk.
* `Path`: Specify your additions to the path to find needed third party tools such as Gradle or Maven.
* `Custom Endpoint`: Specify the custom endpoint for Single Tenant setups instead of https://app.snyk.io.
* `Allow unknown certificate authorities`: Disable certificate checks for SSL connections.
* `Snyk Open Source enabled`: Enable/Disable Snyk Open Source Dependency Scans via Language Server. Default: `Enabled` during beta
* `Snyk Code enabled`: Enable/Disable Snyk Code Scans via Language Server. Default: `Disabled` during beta.
* `Snyk Infrastructure-as-Code enabled` : Enable/Disable Snyk IaC Scans via Language Server. Default: `Enabled` during beta.
* `Organization`: Specify the Snyk Organization to use for scanning. Snyk recommends using the `ORG_ID`. If you specify an organization slug name, the value must match the URL slug as displayed in the URL of your org in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`.
* `Additional Parameters`: Specify additional parameters to pass to the CLI (for example, `--all-projects` or `-d.`
* `Additional Environment`: Add environment variables to Language Server, multiple can be separated by `;`. Example: `JAVA_HOME=/usr/local/bin;GOPATH=/usr/local/bin`
* `Update and install Snyk binaries automatically`: If `disabled`, no updates are downloaded and updates must be performed manually. Snyk recommends always using the most recent version of the CLI. Please make sure that the locations for Language Server and CLI point to an existent, current binary.
* `Snyk Language Server`: Specify the location of your language server binary.
* `Snyk CLI`: Specify the location of the Snyk CLI .
* `Send error reports to Snyk`: Send errors from Language Server to Snyk to enable quicker bug fixing. Default: `Enabled`.
* `Send usage statistics to Snyk`: Allow Snyk to get usage data to improve workflows. Default: `Enabled`.

###
