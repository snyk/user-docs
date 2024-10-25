# Run an analysis with Visual Studio Code extension

{% hint style="info" %}
Ensure the Snyk extension is configured, authenticated, and trusted for your current project, as described in the authentication and configuration pages.
{% endhint %}

You can trigger `snyk test` using one of these methods:

* automatic (default)
* manual

Snyk Code and Iac (configuration) scans are triggered automatically when your Project is opened and when any supported files are saved. This behavior can be turned off using an [existing configuration](visual-studio-code-extension-configuration.md#user-experience).

Snyk Open Source does not automatically run on save by default, but you can enable auto-scan in settings.

To manually trigger `snyk test`, click  the **Rescan** (play) button at the top of the plugin's panels.

{% hint style="info" %}
Ensure your files are saved before manually running an analysis.
{% endhint %}

<figure><img src="../../../.gitbook/assets/SCR-20241024-qqsi.png" alt="How to manually trigger a Snyk analysis" width="355"><figcaption><p>How to manually trigger a Snyk analysis</p></figcaption></figure>

## User experiences while scanning

In the IDE, observe that the extension is already picking up the files and uploading them for analysis.

Snyk Open Source requires the Snyk CLI, so it is downloaded in the background.

Snyk Code analysis runs quickly without the CLI, so results may already be available. Otherwise, you see the following screen while Snyk scans your workspace for vulnerabilities and issues:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-16 at 14.53.38.png" alt="Snyk scan in progress"><figcaption><p>Snyk scan in progress</p></figcaption></figure>

