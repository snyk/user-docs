# Troubleshooting for the JetBrains plugin

{% hint style="warning" %}
Snyk plugins are not supported on any operating system that has reached End Of Life (EOL) with the distributor.
{% endhint %}

## Get detailed debug logs

{% hint style="warning" %}
When you enable `debug`, your code may be logged in the IDE log files, for example, the `idea.log` file.
{% endhint %}

To enable Snyk Language Server debug logs in the JetBrains IDEs, change the log level to debug using the IDE by navigating to Debug Log Settings (Custom Debug Log Configuration).&#x20;

Press the Shift key twice quickly and select the **Actions** tab. Then search for **Debug**. Alternatively, select the Debug Log Settings in the menu (not available in JetBrains Rider).

<figure><img src="../../../.gitbook/assets/image (179).png" alt="Actions tab"><figcaption><p>Use Actions tab to open Debug Log Settings</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (180).png" alt="Search for action"><figcaption><p>Go to Custom Debug Log Configuration from Debug Log Settings</p></figcaption></figure>

Entering `Snyk Language Server` on its own line will enable debug logging of the Language Server. Restart the IDE to reload the Snyk Language Server with the new debug level logging enabled.

<figure><img src="../../../.gitbook/assets/image (181).png" alt="Snyk Language Server configuration"><figcaption><p>Snyk Language Server configuration set to Debug Mode in the Custom Debug Log Configuration</p></figcaption></figure>

To view the debug logs, navigate to **Help** > **Show Log** in the Finder (Mac) or **Show Log** in the Explorer (Windows). Then open the files idea.log, **idea.1.log**, and so on, in the folder.

<figure><img src="../../../.gitbook/assets/image (178).png" alt="Show log in Finder" width="323"><figcaption><p>Opens the Intellij logs folder</p></figcaption></figure>

## Trusted root certificates issues

See the [JetBrains documentation](https://www.jetbrains.com/help/idea/ssl-certificates.html) for information about how the IDE resolves custom certificates and how to import them if the plugin experiences any network failures due to incorrect configuration.

## Snyk Code checkboxes disabled in JetBrains plugin

Sometimes the checkboxes for Snyk Code in the JetBrains plugin are disabled. Some possible reasons follow:

* Network or proxy settings: If the network or proxy settings are not configured correctly, the checkboxes may be disabled. Check to see if there is an MITM proxy with certificate substitution. Also, verify whether connections to the endpoint API and deeproxy can be established using other tools, for example, the CLI or cURL.
* Incorrect endpoint addres&#x73;**:** If the endpoint address in the Snyk Code plugin configuration is incorrect, the checkboxes will be disabled. To fix this, refer to the instructions and check that the endpoint address is correct. Restart the plugin afterwards.
* **Snyk Code is disabled server-side:** If Snyk Code is disabled in the Snyk Organisation's settings, the checkboxes will be disabled. To fix this, follow the instructions shown in the IntelliJ settings. Restart your IDE.
* **Have a look at the JetBrains logs:** For additional information, see [Locating IDE log files](https://intellij-support.jetbrains.com/hc/en-us/articles/207241085-Locating-IDE-log-files).

## Undefined Python version

Snyk uses Python in order to scan and find your dependencies.

If you are using multiple Python versions, use the `--command` option to specify the correct Python command for execution. The plugin does not detect the Python version associated with the Project.

## JCEF problem and error when showing Issue Details

A `java.lang.NullPointerException` was encountered within **Jetbrains 2025.1** when the Snyk Security plugin was installed. Also, **Android Studio** does not enable JCEF out of the box.

### Initial analysis of the provided log (Jetbrains 2025.1 platform)

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

The `NullPointerException: Cannot read field "objId" because "robj" is null` encountered in IntelliJ IDEA 2025.1 with the Snyk plugin installed is a manifestation of a known platform issue, tracked as IJPL-186252. This issue is strongly correlated with the new default out-of-process JCEF architecture in IntelliJ IDEA 2025.1. The Snyk plugin, in its operations involving JCEF for UI rendering, was triggering this underlying platform instability.

### How to solve Android Studio issue

The most effective and recommended path to resolution is as follows:

As of **Android Studio Koala (2024.1.1), JCEF (Java Chromium Embedded Framework) is included** but **not enabled by default**. To utilize JCEF features, you need to manually enable it by following these steps:

#### Step 1: Disable the JCEF Sandbox

1. Open Android Studio.
2. Navigate to **Help** > **Find Action...**\
   (Or press `Ctrl+Shift+A` on Windows/Linux, or `Cmd+Shift+A` on macOS.)
3. In the search box, type **Registry...** and select it.
4. In the Registry dialog that appears, search for:`ide.browser.jcef.sandbox.enable`
5. Uncheck the box next to it to disable the sandbox.

#### Step 2: Choose a JCEF-Compatible Java Runtime

1. Again, navigate to **Help** > **Find Action...**\
   (Or press `Ctrl+Shift+A` on Windows/Linux, or `Cmd+Shift+A` on macOS.)
2. Search for and select **Choose Boot Java Runtime for the IDE...**
3. From the list, select a runtime labeled something like:**JetBrains Runtime with JCEF**
4. Click **OK** to apply the change.
5. Restart Android Studio when prompted.

### How to solve issues with Jetbrains 2025.1 platform

Apply the IntelliJ IDEA VM option `-Dide.browser.jcef.out-of-process.enabled=false` to revert to the in-process JCEF mode, which is known to bypass the bug described in IJPL-186252.

It is also advisable to keep IntelliJ IDEA updated to the latest patch release of version 2025.1.x or later, as future updates may include a permanent fix for IJPL-186252, rendering plugin-specific workarounds or manual VM option adjustments unnecessary. The dynamic between IDE platforms and their extensive plugin ecosystems often involves such diagnostic and adaptive challenges, where changes at one layer necessitate responses and adjustments at others to maintain overall system stability and user experience.

## How Snyk Container and Kubernetes JetBrains integration works

The JetBrains plugin scans your Kubernetes workload files and collects the images used. To troubleshoot whether the plugin is correctly scanning a container image, you can verify the following:

* Whether the image definition is in the Kubernetes YAML file in the Project. Ensure you have the image specified with an image name mapped in the format `imageValue:imageKey` for the image YAML attribute, for example, `image:nginx:1.17.1`.
* Whether the container image has been successfully built locally, pushed to a container registry, or both. It is also a good practice to verify this before referring to the container image in the Kubernetes YAML file.

If you encounter an error, [contact Snyk support](https://support.snyk.io).

For each image found, perform a test with the Snyk CLI.

* For more information about how Snyk Container performs a test on the image, refer to [Snyk CLI for Snyk Container](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/).
* While testing the image, the CLI downloads the image if it is not already available locally in your Docker daemon.
* Snyk plans to expand the scope of container scanning, so if there are more files, such as Dockerfiles or workflows that you want to be supported, submit a feature request to [Snyk support](https://support.snyk.io).
