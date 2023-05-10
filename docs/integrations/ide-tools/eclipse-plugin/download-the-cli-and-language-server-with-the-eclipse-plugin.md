# Download the CLI and language server with the Eclipse plugin

The Snyk Eclipse plugin works with an underlying language server for an optimal Eclipse experience. After restart, when you open a file that Snyk supports, the Eclipse plugin ensures the prerequisites for the plugin are satisfied.

The prerequisites include downloading the Snyk CLI, downloading the language server, and asking you to authenticate. All of these steps are shown on this page and the next, in the order they happen.

The process downloads the Snyk CLI unless you have opted out of download through the plugin, then downloads the language server and uses the CLI for authentication.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-10-19 at 09.10.10 (1).png" alt="Downloading Snyk CLI dependency"><figcaption><p>Downloading Snyk CLI dependency</p></figcaption></figure>

When the download is completed successfully, you will see a success notification in the bottom right corner.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-10-19 at 09.10.23.png" alt="Snyk CLI has been downloaded.notification"><figcaption><p>Snyk CLI has been downloaded.notification</p></figcaption></figure>

Continue with [Authentication for the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/authentication-for-the-eclipse-plugin).
