# How to set environment variables by operating system for IDEs and CLI

## Windows <a href="#windows" id="windows"></a>

* Can be centrally defined in System Settings and should be set there.
* Windows WSL2 env variables should be set like in Linux. For details, see [Share Environment Vars between WSL and Windows](https://devblogs.microsoft.com/commandline/share-environment-vars-between-wsl-and-windows/)

## Mac <a href="#mac" id="mac"></a>

There are several places where you can set environment variables.

For details, see [HowTo: Set an Environment Variable in Mac OS X - /etc/\<wbr>launchd.conf - Dowd and Associates](http://www.dowdandassociates.com/blog/content/howto-set-an-environment-variable-in-mac-os-x-slash-etc-slash-launchd-dot-conf/).

* `~/.profile`: use this for variables you want to set in all programs launched from the terminal. Note that, unlike on Linux, all shells opened in the Terminal.app are login shells.
* `~/.zshrc`: this is invoked for shells that are not login shells. Use this for aliases and other things that need to be redefined in subshells, not for environment variables that are inherited.
* `/etc/profile`: this is loaded before `~/.profile`, but is otherwise equivalent. Use it when you want the variable to apply to terminal programs launched by all users on the machine (assuming they use bash).
* Your user’s `launchd` instance: this applies to all programs launched by the user, GUI, and CLI. You can apply changes at any time by using the `setenv` command in `launchctl`. In theory, you should be able to put `setenv` commands in `~/.launchd.conf`, and `launchd` would read them automatically when the user logs in, but in practice support for this file was never implemented. Instead, you can use another mechanism to execute a script at login and have that script call `launchctl` to set up the `launchd` environment.
* `/etc/launchd.conf:` this is read by `launchd` when the system starts up and when a user logs in. These variables affect every single process on the system, because `launchd` is the root process. To apply changes to the running root `launchd` you can pipe the commands into `sudo launchctl`.

The fundamental things to understand are:

* Environment variables are inherited by a process’s children at the time they are forked.
* The root process is a `launchd` instance, and there is also a separate `launchd` instance per user session.
* `launchd` allows you to change its current environment variables using `launchctl`; the updated variables are then inherited by all new processes it forks from then on.

You can verify if IntelliJ runs using this environment variable using `ps eww -o command <PID> | tr ' ' '\\n'` command.

## Linux <a href="#linux" id="linux"></a>

* For details, see [EnvironmentVariables - Community Help Wiki](https://help.ubuntu.com/community/EnvironmentVariables#System-wide_environment_variables).
* Setting `.profile / .bashrc` is not enough for UI apps, as the system does not launch them from the terminal. The relevant variables need to be available to the process that is launching the apps, for example, the window manager.

## Important environment variables for IDEs and CLI <a href="#important-environment-variables-for-ides-cli" id="important-environment-variables-for-ides-cli"></a>

### CLI <a href="#cli" id="cli"></a>

* `HTTP_PROXY`
* `HTTPS_PROXY`
* `NO_PROXY`
* `PATH` (needs to include directories to Maven and Gradle)
* `JAVA_HOME`
* `...`

### Java <a href="#java" id="java"></a>

* `http.proxyHos`t
* `https.proxyHost`
* `http.proxyPort`
* `https.proxyPort`

### Golang <a href="#golang" id="golang"></a>

* `GOPATH` (path to go binaries)
* `GOROOT` (current go installation)

### Python <a href="#python" id="python"></a>

`PYTHONPATH`

## Proxy <a href="#proxy" id="proxy"></a>

### Set proxy in Java <a href="#proxy-in-java" id="proxy-in-java"></a>

See [Configure HTTP/HTTPS Proxy Settings Java](https://memorynotfound.com/configure-http-proxy-settings-java/).

### Set proxy in Visual Studio Code <a href="#set-proxy-in-visual-studio-code" id="set-proxy-in-visual-studio-code"></a>

See [JohannesHoppe/settings.json](https://gist.github.com/JohannesHoppe/23105342f6580847578701f0ced9d5b0).

### Set proxy in Jetbrains apps <a href="#set-proxy-in-jetbrains-apps" id="set-proxy-in-jetbrains-apps"></a>

See [HTTP Proxy | IntelliJ IDEA](https://www.jetbrains.com/help/idea/settings-http-proxy.html).

The Snyk Jetbrains plugin **does not read the proxy settings from this configuration.** **You must set** **both** JAVA Proxy environment variables and the CLI environment variables.
