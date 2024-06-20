# Troubleshooting for the Eclipse plugin

## General troubleshooting

For general troubleshooting please see the troubleshooting pages [here](../troubleshooting-ides/).

## Logs

To determine where plugin logs are stored, navigate to **Preferences** > **Language Servers** > **Logs** and find the **Snyk Language Server** row in Eclipse. As it can be disabled, you may need to enable it to retrieve the logs. You will find the logs either in the console or in the file based on the preference set.

To see additional plugin error logs:

1. Navigate to **Window** > **Show View** > **Others...**.
2. In **type text filer** search for **Error Log**.
3. Click **Open** to see the error log tab. If you group the tab view by plugin (kebab menu in the top right corner > **Group By** > **Plug-in**), the `io.snyk.eclipse.plugin` row shows any plugin errors.

The log level can be set in the Eclipse settings via environment variable `SNYK_LOG_LEVEL` , e.g. `SNYK_LOG_LEVEL=debug`

Please determine the installation path of Language Server, as displayed in the plugin settings. The language server logs can be found in the same directory with the file name `snyk-ls.log`

Note: When working on the terminal, be careful to escape spaces in the folder path.

## Proxy Settings

Make sure to read our official [eclipse documentation](https://docs.snyk.io/ide-tools/eclipse-plugin)

#### Ensure the following details:

* Does the issue occurs on the CLI terminal, outside the IDE?
* Preferably use the latest Snyk CLI version
* Obtain the user snyk test/monitor output with the debug flag
* Set the proxy variable in the command line set http\_proxy=\<http….>

#### &#x20;Confirm if Proxy is being used

In eclipse, go to Windows → Preferences → General → Network Connections.

Ensure the configured proxy settings are identical to the ones that have been set in the CLI terminal.

The proxy settings can be overridden by adding the settings through the Snyk preferences Additional Environment:

`https_proxy=http://your-proxy.com:8080`

### Cannot install Eclipse plugin Installation

If you are facing issues installing Eclipse plugins, it might be due to the incompatibility of the JDK version used by Eclipse. This can be resolved by running Eclipse with JDK 17 or a higher version. Here are the steps to follow:

1. Download and install JDK 17 or a higher version
2. Download and install Eclipse IDE from the official Eclipse website.
3. Once installed, navigate to the Eclipse installation directory and locate the 'eclipse.ini' file.
4.  Open the 'eclipse.ini' file and add the following lines to it:

    * vm {path to JDK 17 or higher version}\bin

    Note: Make sure to replace {path to JDK 17 or higher version} with the actual path where JDK 17 or higher version is installed on your system.
5. Save and close the 'eclipse.ini' file.
6. Launch Eclipse IDE using the usual method.
7. Once Eclipse is launched, go to 'Window' > 'Preferences' > 'Java' > 'Installed JREs'.
8. Click on 'Add' and select the path to JDK 17 or higher version installed on your system.
9. Click 'OK' and close the Preferences dialogue box.
10. You are now ready to install any Eclipse plugin of your choice.

Here's an example of how the `eclipse.ini` file could look like after adding the path to JDK 17 or higher version. The last line is the relevant change, the rest of the file, can, but doesn’t need to be touched.

`-startup plugins/org.eclipse.equinox.launcher_1.6.200.v20210416-2027.jar --launcher.library plugins/org.eclipse.equinox.launcher.win32.win32.x86_64_1.2.200.v20210429-1609 -product org.eclipse.epp.package.jee.product -showsplash org.eclipse.epp.package.common --launcher.defaultAction openFile --launcher.appendVmargs -vmargs -Dosgi.requiredJavaVersion=11 -Xms256m -Xmx2048m --add-modules=ALL-SYSTEM --add-opens=java.base/java.util=ALL-UNNAMED --add-opens=java.base/java.lang=ALL-UNNAMED --add-opens=java.base/java.io=ALL-UNNAMED --add-opens=java.base/java.nio=ALL-UNNAMED --add-opens=java.base/sun.nio.ch=ALL-UNNAMED -vm C:\\Program Files\\Java\\jdk-17.0.1\\bin`

Make sure to replace `C:\\\\Program Files\\\\Java\\\\jdk-17.0.1\\\\bin` with the actual path where JDK 17 or higher version is installed on your system.

### Application Development <a href="#application-development" id="application-development"></a>

While running Eclipse with JDK 17 or a higher version can resolve issues related to plugin installation, it is important to note that it may not be compatible with all versions of Java used for developing applications. Depending on the specific requirements of your application, you may need to use a different version of Java, such as JDK 8 or 11, for development purposes. Be sure to consult the documentation and requirements of your application before making any changes to your Java environment.

It's also worth noting that while you can install multiple versions of JDK on your system, you can only use one version at a time. This means that if you need to switch between different versions of JDK for different projects, you'll need to update the Eclipse configuration accordingly.

To change the JDK version used by Eclipse for a specific project, follow these steps:

1. Open the project in Eclipse.
2. Right-click on the project and select 'Properties'.
3. In the Properties dialogue box, go to 'Java Build Path' > 'Libraries'.
4. Locate the JRE System Library and click on it to expand the options.
5. Click on 'Edit' and select the desired JDK version from the list of installed JREs.
6. Click 'Finish' to save the changes.

By following these steps, you can easily switch between different JDK versions for different projects in Eclipse. It's important to ensure that you're using the correct version of JDK for each project, as this can impact the compatibility and functionality of your application.

### Windows Defender warns of running a binary

Windows Defender can sometimes block or issue warnings when running Go binaries. This can happen due to various reasons, such as the binary being unrecognized or having behavior that triggers the antivirus software.

To resolve this issue, you can try the following solutions:

* You can use the code-signed CLI as language server binary, by specifying a CLI location as language server path in Snyk Settings (`Snyk Language Server`) and restarting Eclipse:

<figure><img src="../../../.gitbook/assets/image (479).png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (482).png" alt=""><figcaption></figcaption></figure>

* Add an exclusion: Exclude the Go binary or the directory containing the binary from Windows Defender's scanning. This can help prevent false positives and allow the binary to run without interference.
* Disable Windows Defender temporarily: If adding an exclusion doesn't solve the problem, you can temporarily disable Windows Defender while running the Go binary. However, exercise caution when disabling antivirus software and ensure that you have other security measures in place.
* Submit the binary to Microsoft: If you believe the detection is a false positive, you can report it to Microsoft. They have a process for submitting files to be reviewed, and if confirmed as a false positive, the detection may be updated in future antivirus definitions.

\
