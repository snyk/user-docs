# Troubleshooting for the JetBrains plugin

{% hint style="warning" %}
Snyk plugins are not supported on any operating system that has reached End Of Life (EOL) with the distributor.
{% endhint %}

## Get detailed logs: debug logs

{% hint style="warning" %}
When you enable `debug`, your code may be logged in the IDE log files, for example, the `idea.log` file.
{% endhint %}

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

* **Network or proxy settings:** If the network or proxy settings are not configured correctly, the checkboxes may be disabled. Check to see if there is an MITM proxy with certificate substitution. Also, verify whether connections to the endpoint API and deeproxy can be established using other tools, for example, the CLI or cURL.
* **Incorrect endpoint address:** If the endpoint address in the Snyk Code plugin configuration is incorrect, the checkboxes will be disabled. To fix this, check that the endpoint address is correct by following the instructions. Restart IntelliJ afterwards.
* **Snyk Code is disabled server-side:** If Snyk Code is disabled in the Snyk Organisation's settings, the checkboxes will be disabled. To fix this, follow the instructions shown in the IntelliJ settings. Restart your IDE.
* **Have a look at the JetBrains logs:** For additional information, see [Locating IDE log files](https://intellij-support.jetbrains.com/hc/en-us/articles/207241085-Locating-IDE-log-files).

## Undefined Python version

Snyk uses Python in order to scan and find your dependencies.

If you are using multiple Python versions, use the `--command` option to specify the correct Python command for execution. The plugin does not detect the Python version associated with the Project.

## JCEF problem and error when showing Issue Details in 2025.1

A `java.lang.NullPointerException` was encountered within IntelliJ IDEA 2025.1 when the Snyk Security plugin was installed. The user query indicated a suspicion that changes in the Java Chromium Embedded Framework (JCEF) API or its support within the IntelliJ Platform 2025.1 might be the underlying cause. This analysis dissects the provided stack trace, correlates it with known platform issues and recent changes in both IntelliJ IDEA and the Snyk plugin, and provides actionable recommendations to resolve the error.

### Initial analysis of the provided log

The exception encountered is a `java.lang.NullPointerException` with the message: `Cannot read field "objId" because "robj" is null`. An examination of the stack trace reveals the origin and propagation of this error:

```
java.lang.NullPointerException: Cannot read field "objId" because "robj" is null
at...[source](https://youtrack.jetbrains.com/issue/IJPL-186252/JCEF.-NullPointerException-Cannot-read-field-objId-because-robj-is-null)
at com.intellij.ui.jcef.JBCefApp.createMessageRouter(JBCefApp.java:399)
at com.intellij.ui.jcef.JBCefJSQuery$JSQueryFunc.<init>(JBCefJSQuery.java:246)
at com.intellij.ui.jcef.JBCefJSQuery$JSQueryFunc.<init>(JBCefJSQuery.java:236)
at com.intellij.ui.jcef.JBCefJSQuery.lambda$create$0(JBCefJSQuery.java:56)
at com.intellij.ui.jcef.JBCefJSQuery.create(JBCefJSQuery.java:59)
at io.snyk.plugin.ui.jcef.OpenFileLoadHandlerGenerator.generate(OpenFileLoadHandlerGenerator.kt:38)
at io.snyk.plugin.ui.toolwindow.panels.SuggestionDescriptionPanelFromLS$1.invoke(JCEFDescriptionPanel.kt:53)
at io.snyk.plugin.ui.toolwindow.panels.SuggestionDescriptionPanelFromLS$1.invoke(JCEFDescriptionPanel.kt:52)
at io.snyk.plugin.ui.jcef.JCEFUtils.getJBCefBrowserIfSupported(Utils.kt:27)
```

Key observations from the stack trace:

* The error originates deep within the JCEF remote messaging infrastructure, specifically in `com.jetbrains.cef.remote.router.RemoteMessageRouterImpl.create` at line 38. This class is part of the JetBrains implementation of JCEF, handling communication with an out-of-process browser.
* The Snyk plugin initiates the sequence of calls leading to the error. The trace shows `io.snyk.plugin.ui.jcef.OpenFileLoadHandlerGenerator.generate` attempting to create a `JBCefJSQuery` (via `com.intellij.ui.jcef.JBCefJSQuery.create`).
* `JBCefJSQuery.create` internally calls `com.intellij.ui.jcef.JBCefApp.createMessageRouter`, which in turn attempts to create a `CefMessageRouter`. This process ultimately fails within the remote router initialization.

The failure point, `RemoteMessageRouterImpl.create`, suggests an issue with setting up or accessing components related to a remote (out-of-process) JCEF instance. The field `robj` being null indicates that an expected remote object, crucial for establishing the message router, was not available or initialized at the point of access. This aligns with the user's initial hypothesis that changes to JCEF support in IntelliJ 2025.1 could be a factor.

The `NullPointerException: Cannot read field "objId" because "robj" is null` encountered in IntelliJ IDEA 2025.1 with the Snyk plugin installed is a manifestation of a known platform issue, tracked as IJPL-186252. This issue is strongly correlated with the new default out-of-process JCEF architecture in IntelliJ IDEA 2025.1. The Snyk plugin, in its operations involving JCEF for UI rendering, was triggering this underlying platform instability.

The most effective and recommended path to resolution is as follows:

1. Update the Snyk Security plugin to version 2.12.2 or later. This version contains a specific workaround implemented by Snyk to mitigate the JCEF initialization problem in IntelliJ 2025.1.
2. If updating the Snyk plugin does not resolve the issue, or as an alternative, apply the IntelliJ IDEA VM option `-Dide.browser.jcef.out-of-process.enabled=false` to revert to the in-process JCEF mode, which is known to bypass the bug described in IJPL-186252.

It is also advisable to keep IntelliJ IDEA updated to the latest patch release of version 2025.1.x or later, as future updates from JetBrains may include a permanent fix for IJPL-186252, rendering plugin-specific workarounds or manual VM option adjustments unnecessary. The dynamic between IDE platforms and their extensive plugin ecosystems often involves such diagnostic and adaptive challenges, where changes at one layer necessitate responses and adjustments at others to maintain overall system stability and user experience.

## How Snyk Container and Kubernetes JetBrains integration works

The JetBrains plugin scans your Kubernetes workload files and collects the images used. To troubleshoot whether the plugin is correctly scanning a container image, you can verify the following:

* Whether the image definition is in the Kubernetes YAML file in the Project. Ensure you have the image specified with an image name mapped in the format `imageValue:imageKey` for the image yaml attribute, for example, `image:nginx:1.17.1`.
* Whether the container image has been successfully built locally or pushed to a container registry or both. It is also a good practice to verify this before referring to the container image in the Kubernetes YAML file.

If you encounter an error [contact support](https://support.snyk.io).

For each image found, perform a test with the Snyk CLI.

* For more information about how Snyk Container performs a test on the image, Refer to [Snyk CLI for Snyk Container](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/).
* While testing the image, the CLI downloads the image if it is not already available locally in your Docker daemon.
* Snyk plans to expand the scope of container scanning, so if there are more files, such as Dockerfiles or workflows that you want to be supported, submit a feature request [to Snyk support](https://support.snyk.io).
