# Troubleshooting for Visual Studio Code extension

The logs for VSCode can be found in the output channels for the Snyk Extension and Snyk Language Server. Both are needed for troubleshooting.

To enable debug log level, start VS Code from the terminal like this:

`cd your-repo SNYK_LOG_LEVEL=debug code .`

Alternatively, you can add `-d` to additional parameters in the Snyk Settings.

You can obtain additional logs by navigating to **Help** > **Toggle Developer Tools**, and opening the **Console** tab for additional output.

<figure><img src="../../../.gitbook/assets/image (484).png" alt="Additional outout"><figcaption><p>Additional outout</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (485).png" alt="Selected confiuration, Snyk Security - Code, Open Source Dependencies, IaC Configurations"><figcaption><p><code>Selected confiuration, Snyk Security - Code, Open Source Dependencies, IaC Configurations</code></p></figcaption></figure>

## Command not found (getActiveUser, LoginCommand, and so on)

This indicates that the required binaries are not available. To start the extension and use it, the following are required:

* Snyk Language Server (snyk-ls) available and executable
* Snyk CLI available and executable

These are usually downloaded automatically by the extension. If this is not possible, the language server cannot be started, and thus not provide the indicated commands to VS Code.

## To fix <a href="#to-fix" id="to-fix"></a>

* Check to see if automatic download is activated and if it works.
  * Proxy/SSL problems could be a reason why the download does not finish.
  * Missing write permissions in the configured path could be another reason why automatic download does not work
* Before reaching out for help, check:
  * secure debug logs (`-d` in `additional parameters` ) and
  * configuration
  * manual CLI run logs and information on the network:
    * do they have a proxy?
    * do they use a mitm proxy that intercepts ssl connections?
    * if yes, have they made the custom certificate authority available to the OS?
* If automatic management of dependencies is not enabled,
  * download the Language Server and the CLI manually
  * make them executable, and
  * provide the path to the binary in the snyk settings in the corresponding fields.
