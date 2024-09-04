# Download the CLI and language server with the Eclipse plugin

The Snyk Eclipse plugin works with an underlying [Language Server](../snyk-language-server/) for an optimal Eclipse experience. After restart, when you open a file that Snyk supports, the Eclipse plugin ensures the prerequisites for the plugin are satisfied.

The prerequisites include downloading the [Snyk CLI](../../../snyk-cli/) and authenticating in response to the prompt. These steps are shown on this page and the next, in the order they happen.

The process downloads the Snyk CLI unless you have opted out of downloading through the plugin, and then downloads the Language Server and uses the CLI for authentication.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-10-19 at 09.10.10 (1).png" alt="Downloading Snyk CLI"><figcaption><p>Downloading the Snyk CLI</p></figcaption></figure>

Continue with [Authentication for the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/authentication-for-the-eclipse-plugin).
