# Troubleshooting for the JetBrains plugin

{% hint style="warning" %}
Snyk plugins are not supported on any operating system that has reached End Of Life (EOL) with the distributor.&#x20;
{% endhint %}

## Get detailed logs: debug logs

Each JetBrains IDE provides easy access to the logs.

Follow these steps to set logging to the debug level.

To obtain plugin logs, navigate to **Help** > **Show Log in Finder** (Mac) or **Show Log in Explorer** (Windows).

<figure><img src="../../../.gitbook/assets/image (487).png" alt="Show log in Finder"><figcaption><p>Show log in Finder</p></figcaption></figure>

You can change the log level to debug using the IDE:

Press `Shift Shift` quickly and select the tab **Actions** . Then search for `debug` . Alternatively, select the debug log settings in the menu (not available in Jetbrains Rider).

<figure><img src="../../../.gitbook/assets/image (488).png" alt="Actions tab"><figcaption><p>Actions tab</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (489).png" alt="Search for action"><figcaption><p>Search for action</p></figcaption></figure>

* `Snyk Code` enables normal debug logging (only until plugin version 2.6.0).
* `Snyk CodeRequestLogging` enables a detailed protocol of the HTTP requests when communicating with the Snyk Code API (only until plugin version 2.6.0).
* `Snyk Language Server` enables debug logging of the language server in the background.

<figure><img src="../../../.gitbook/assets/image (490).png" alt="Snyk Language Server configuration"><figcaption><p>Snyk Language Server configuration</p></figcaption></figure>

## Trusted root certificates issues

Refer to the [JetBrains documentation](https://www.jetbrains.com/help/idea/ssl-certificates.html) on how the IDE resolves custom certificates and how to import them if the plugin experiences any network failures due to incorrect configuration.

## Snyk Code checkboxes disabled in IntelliJ

Sometimes the checkboxes for Snyk Code in the JetBrains IntelliJ plugin are disabled. Some possible reasons follow:

* **Network or proxy settings:** If the network or proxy settings are not configured correctly, the checkboxes may be disabled. Check to see if there is an MITM proxy with certificate substitution. Also, verify whether connections to the endpoint API and deeproxy can be established using other tools, for example, the CLI or cURL.&#x20;
* **Incorrect endpoint address:** If the endpoint address in the Snyk Code plugin configuration is incorrect, the checkboxes will be disabled. To fix this, check that the endpoint address is correct by following the instructions. Restart IntelliJ afterwards.
* **Snyk Code is disabled server-side:** If Snyk Code is disabled in the Snyk Organisation's settings, the checkboxes will be disabled. To fix this, follow the instructions shown in the IntelliJ settings. Restart your IDE.
* **Have a look at the JetBrains logs:** For additional information, see  [Locating IDE log files](https://intellij-support.jetbrains.com/hc/en-us/articles/207241085-Locating-IDE-log-files).

## Undefined Python version

Snyk uses Python in order to scan and find your dependencies.&#x20;

If you are using multiple Python versions, use the -`-command` option to specify the correct Python command for execution. The plugin does not detect the Python version associated with the Project.
