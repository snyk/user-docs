# Environment variables for the Eclipse plugin

To analyze open-source dependencies and IAC template files, the plugin uses the Snyk CLI. The CLI needs the following environment variables:

* `PATH`: The path to needed binaries, for example, to Maven. You can also adjust the `PATH` variable manually using the **`Path`** field in the settings dialog.
* `JAVA_HOME`: The path to the JDK you want to use to analyze Java dependencies.
* `http_proxy` and `https_proxy`: Proxy set using the value in the format `http://username:password@proxyhost:proxyport`.\
  Note that the leading `http://` in the value does not change to `https://` for an `https_proxy`. If you populate the proxy settings in Eclipse, the Snyk plugin will set the environment variables automatically for Language Server and CLI.

Setting these variables only in a shell environment (for example, using `~/.bashrc`) is not enough, if you do not start Eclipse from the command line or create a script file that starts Eclipse using a shell environment.

* On Windows, set the variables using the GUI or on the command line using the [setx](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/setx) tool.
* On macOS, the process `launchd` needs to know the environment variables to launch Eclipse directly from the Finder. Set these environment variables using the `launchctl setenv` command, for example, on startup or using a script you launch at user login.\
  Note that providing environment variables to the macOS UI may change between operating system releases, so it can be easier to create a small shell script that launches the Eclipse app to leverage the shell environment. This can be defined using `~/.bashrc`.
* On Linux, you can update the file `/etc/environment` to propagate the environment variables to the window manager and UI.

{% hint style="info" %}
The Eclipse plugin uses the proxy settings from Eclipse, but also picks up the proxy settings from the environment variables.
{% endhint %}
