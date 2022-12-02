# Download the CLI and language server with the Eclipse plugin

The Snyk Eclipse plugin works with an underlying language server for an optimal Eclipse experience. After restart, when you open a file that Snyk supports, the Eclipse plugin attempts to start a workspace scan:

![Eclipse plugin starting a scan](<../../../.gitbook/assets/Screenshot 2022-05-13 at 09.28.30 (1).png>)

Once the plugin is installed, all of the plugin's prerequisites are triggered. The prerequisites include downloading the Snyk CLI, downloading the language server, and asking you to authenticate. All of these steps are shown on this page and the next, in the order they happen.

The process downloads the Snyk CLI unless you have opted out of download through the plugin, downloads the language server, and uses the CLI for authentication.

![Download the Snyk CLI](<../../../.gitbook/assets/Screenshot 2022-05-13 at 11.27.00.png>)

Continue with [Authentication for the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/authentication-for-the-eclipse-plugin).
