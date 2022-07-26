# Configuration of the Eclipse plugin

### Configuration

In the Snyk preferences, the following configuration options are available.

* `Snyk API Token`: The authentication token from Snyk.
* `Path`: Your additions to the path to find needed third party tools such as Gradle or Maven.
* `Custom Endpoint`: The custom endpoint for Single Tenant setups instead of https://app.snyk.io.
* `Allow unknown certificate authorities`: Disable certificate checks for SSL connections.
* `Custom Snyk LS Path`: Specify the location of your language server binaries. If set, no updates are downloaded and updates must be performed manually to synchronize Eclipse features and preferences with the Language Server.
* `Snyk Open Source enabled`: Enable/Disable Snyk Open Source Dependency Scans via Language Server. Default: `Enabled` during beta
* `Snyk Code enabled`: Enable/Disable Snyk Code Scans via Language Server. Default: `Disabled` during beta.
* `Snyk Infrastructure-as-Code enabled` : Enable/Disable Snyk IaC Scans via Language Server. Default: `Enabled` during beta.
* `Organization`: Specify the Snyk Organization to use for scanning.
* `Additional Parameters`: Specify additional parameters to pass to the CLI (for example, `--all-projects` or `-d.`
* `Additional Environment`: Add environment variables to Language Server.
* `Send error reports to Snyk`: Send errors from Language Server to Snyk to enable quicker bug fixing. Default: `Enabled` during beta.
* `Send usage statistics to Snyk`: Allow Snyk to get usage data to improve workflows. Default: `Disabled` during beta.

###
