# Visual Studio Code extension configuration

## Environment variables

To analyze projects, the plugin uses the Snyk CLI which requires environment variables:

* `PATH`: the path to needed binaries, for example, to maven
* `JAVA_HOME`: the path to the JDK you want to use to analyze Java dependencies

Setting these variables only in a shell environment (for example, using `~/.bashrc`) is not sufficient, if you do not start the IDE from the command line or create a script file that starts the IDE using a shell environment.

* On `Windows`, you can set the variables using the GUI or on the command line using the `setx` tool.
* On `macOS`, the process `launchd` must know the environment variables to launch the IDE from Finder directly. You can set environment variables for applications launched using Finder by using the `launchctl setenv` command, for example, on start-up or through a script you launch at user login.\
  **Note:** The provision of environment variables to the macOS UI can change between operating system releases, so it may be easier to create a small shell script that launches the IDE to leverage the shell environment that can be defined via `~/.bashrc`.
* On `Linux`, updating the file /etc/environment can propagate the environment variables to the windows manager and UI.

## Proxy

If you are behind a proxy, configure the proxy settings using VS Code proxy settings or set the proxy settings using `http_proxy` and `https_proxy` environment variables.

## Visual Studio Code extension configuration

After the extension is installed, you can set the following configurations for the extension:

* **Features**
  * **Code Security**: configure if code security analysis should run over your code.
  * **Code Quality**: configure if code quality analysis should run over your code.
  * **Open Source Security**: configure if security analysis should run over your open source dependencies.
* **Severity**: set severity level to display in the analysis result tree.
* **Advanced**
  * **Advanced mode**: toggle a panel to allow the user to manually control when the analysis should be run.
  * **Auto Scan Open Source Security**: set severity level to display in the analysis result tree.
  * **Additional Parameters**: set parameters to be passed to Snyk CLI for Open Source Security tests. For the full list see the [CLI commands and options summary](../../snyk-cli/cli-reference.md).
  * **Organization**: specify an organization slug name to run tests for that organization. The value of organization setting `snyk.advanced.organization` must match the URL slug as displayed in the URL of your org in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`. If not specified, the preferred organization as defined in your [web account settings](https://app.snyk.io/account) is used to run tests.
  * **Custom endpoint**: Specify the custom Snyk API endpoint for your organization. Use this field for the custom endpoint for Single Tenant setups as well instead of https://app.snyk.io.
