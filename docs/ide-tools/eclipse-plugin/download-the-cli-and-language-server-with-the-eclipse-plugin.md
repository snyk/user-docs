# Download the CLI and language server with the Eclipse plugin

The Snyk Eclipse plugin works with an underlying language server for an optimal Eclipse experience. After restart, on opening a file that Snyk supports, the Eclipse plugin attempts to start a workspace scan:

![Eclipse plugin starting a scan](<../../.gitbook/assets/Screenshot 2022-05-13 at 09.28.30 (1).png>)

Once the plugin is installed, all of the plugin's prerequisites are triggered on opening a file that Snyk supports (Snyk hooks into this action). The prerequisites include downloading the Snyk CLI, the language server, and asking you to authenticate. All of these steps are shown in the following sections, in the order they happen.

### Snyk CLI and language server download

The process downloads the Snyk CLI and the language server and uses the CLI for authentication.

![Download the Snyk CLI](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.27.00.png>)

###
