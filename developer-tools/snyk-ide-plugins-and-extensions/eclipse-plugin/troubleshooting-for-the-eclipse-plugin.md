# Troubleshooting for the Eclipse plugin

{% hint style="warning" %}
Snyk plugins are not supported on any operating system that has reached End Of Life (EOL) with the distributor.
{% endhint %}

## General troubleshooting

For general troubleshooting, see the [IDE troubleshooting pages](../troubleshooting-ides/).

## Logs

{% hint style="warning" %}
When you enable `debug`, your code may be logged in the IDE log files, for example, the `io.snyk.languageserver.log` file.
{% endhint %}

To enable Snyk Language Server debug logs in Eclipse, navigate to the Snyk Preferences, set the **Additional Environment** preference to `SNYK_LOG_LEVEL=debug` and restart Eclipse.

<figure><img src="../../../.gitbook/assets/image (165).png" alt=""><figcaption><p>Snyk Preferences, Eclipse</p></figcaption></figure>

To view Snyk plugin logs, open the **Error Log** view:

1. Navigate to **Window** > **Show View** > **Others...**.
2. In **type filter text,** search for **Error Log**.
3. Click **Open** to see the error log tab. If you group the tab view by plugin (on the three dots menu in the top right corner, select **Group By** > **Plug-in**), the `io.snyk.eclipse.plugin` rows show any plugin errors.

To get the **Snyk Language Server log file**, click on the **Open Log** button in the **Error Log** view.

Notice that **\[Info]** lines contain **INF**, **WRN**, and **DBG** log rows if the SNYK\_LOG\_LEVEL is set to debug.

<figure><img src="../../../.gitbook/assets/image (166).png" alt=""><figcaption><p>Error Log view</p></figcaption></figure>

## Proxy settings

Ensure that you read the Snyk [Eclipse documentation](./).

### Details to address

* Determine whether the issue occurs on the CLI terminal, outside of the IDE.
* Preferably, use the latest Snyk CLI version.
* Use the debug option to obtain the user's `snyk test` and `snyk monitor` output.
* Set the proxy variable in the command line: s`et http_proxy=<http….>`

### Proxy setting confirmation

In Eclipse, navigate to **Windows** > **Preferences** > **General** > **Network Connections**.

Ensure the configured proxy settings are identical to the ones that have been set in the CLI terminal.

You can override the proxy settings by adding the settings through the Snyk preferences **Additional Environment**:

`https_proxy=http://your-proxy.com:8080`

## Issues with installing the Eclipse plugin

If you are facing issues installing Eclipse plugins, it might be due to the incompatibility of the JDK version used by Eclipse. You can resolve this by running Eclipse with JDK 17 or a higher version. Follow these steps:

1. Download and install JDK 17 or a higher version.
2. Download and install the Eclipse IDE from the official Eclipse website.
3. After the Eclipse IDE is installed, navigate to the Eclipse installation directory and locate the `eclipse.ini` file.
4.  Open the `eclipse.ini` file and add the following to it:

    `vm {path to JDK 17 or higher version}\bin`

    Ensure you replace `{path to JDK 17 or higher version}` with the actual path where the JDK 17 or higher version is installed on your system.
5. Save and close the `eclipse.ini` file.
6. Launch the Eclipse IDE by the usual method.
7. After Eclipse is launched, navigate to **Window** > **Preferences** > **Java** > **Installed JREs**.
8. Click on **Add** and select the **path to JDK 17 or higher version** installed on your system.
9. Click **OK** and close the Preferences dialogue box.

You are now ready to install any Eclipse plugin of your choice.

An example follows showing how the `eclipse.ini` file could look after adding the path to the JDK 17 or higher version. The last line is the relevant change; the rest of the file can be touched, but does not need to be.

`-startup plugins/org.eclipse.equinox.launcher_1.6.200.v20210416-2027.jar --launcher.library plugins/org.eclipse.equinox.launcher.win32.win32.x86_64_1.2.200.v20210429-1609 -product org.eclipse.epp.package.jee.product -showsplash org.eclipse.epp.package.common --launcher.defaultAction openFile --launcher.appendVmargs -vmargs -Dosgi.requiredJavaVersion=11 -Xms256m -Xmx2048m --add-modules=ALL-SYSTEM --add-opens=java.base/java.util=ALL-UNNAMED --add-opens=java.base/java.lang=ALL-UNNAMED --add-opens=java.base/java.io=ALL-UNNAMED --add-opens=java.base/java.nio=ALL-UNNAMED --add-opens=java.base/sun.nio.ch=ALL-UNNAMED -vm C:\\Program Files\\Java\\jdk-17.0.1\\bin`

Ensure you replace `C:\\\\Program Files\\\\Java\\\\jdk-17.0.1\\\\bin` with the actual path where the JDK 17 or higher version is installed on your system.

## Application development and JDK version <a href="#application-development" id="application-development"></a>

While running Eclipse with JDK 17 or a higher version can resolve issues related to plugin installation, it is important to note that it may not be compatible with all versions of Java used for developing applications. Depending on the specific requirements of your application, you may need to use a different version of Java, such as JDK 8 or 11, for development purposes. Be sure to consult the documentation and requirements for your application before making any changes to your Java environment.

Note also that while you can install multiple versions of JDK on your system, you can use only one version at a time. If you need to switch between different versions of JDK for different Projects, you must update the Eclipse configuration accordingly.

To change the JDK version used by Eclipse for a specificProject, follow these steps:

1. Open the Project in Eclipse.
2. Right-click on the Project and select **Properties**.
3. In the Properties dialogue box, navigate to J**ava Build Path** > **Libraries**.
4. Locate the **JRE System Library** and click on it to expand the options.
5. Click on **Edit** and select the desired JDK version from the list of installed JREs.
6. Click **Finish** to save the changes.

By following these steps, you can easily switch between different JDK versions for different Projects in Eclipse. Ensure that you are using the correct version of JDK for each Project, as this can impact the compatibility and functionality of your application.

## Windows Defender warns of running a binary

Windows Defender can sometimes block or issue warnings when running Go binaries. This can happen for various reasons, for example, the binary may be unrecognized or have behavior that triggers the antivirus software.

To resolve this issue, you can try the following solutions:

* Add an exclusion: Exclude the Go binary or the directory containing the binary from Windows Defender scanning. This can help prevent false positives and allow the binary to run without interference.
* Disable Windows Defender temporarily: If adding an exclusion does not solve the problem, you can temporarily disable Windows Defender while running the Go binary. However, exercise caution when disabling antivirus software and ensure that you have other security measures in place.
* Submit the binary to Microsoft: If you believe the detection is a false positive, you can report it to Microsoft. They have a process for submitting files to be reviewed, and if confirmed as a false positive, the detection may be updated in future antivirus definitions.

## **Signing information for Jars**

If you want to verify the correct provenance of your plugin, verify the signing details from the Eclipse dialog using this data.

<figure><img src="../../../.gitbook/assets/image (23).png" alt="The signing key details to verify the integrity and origin of the download plugin"><figcaption><p>The signing key details to verify the integrity and origin of the downloaded plugin</p></figcaption></figure>
