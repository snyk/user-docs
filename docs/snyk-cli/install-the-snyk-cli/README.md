# Install the Snyk CLI

Use our Snyk CLI tool to find and fix known vulnerabilities in your dependencies, both ad hoc and as part of your CI \(Build\) system.

The Snyk CLI requires you to authenticate with your account before using it. See [Authenticate the CLI with your account.](authenticate-the-cli-with-your-account.md)

Install the Snyk CLI using one of these options:

* [Install the Snyk CLI with npm](./)
* [Install the Snyk CLI using  the prebuilt binaries](./)
* [Install the Snyk CLI tool with a container](./)
* [Install the Snyk CLI with Homebrew](./)
* [Install the Snyk CLI with the Windows Scoop package manager](./)

## Install the Snyk CLI with npm

Install our Snyk CLI tool using npm.

**Prerequisites:**

* Ensure you’ve installed the latest version of npm on your local environment, using Node version 10 or later \(see [What version of Node is required for Snyk?](https://support.snyk.io/hc/en-us/articles/360004183317-What-version-of-Node-is-required-for-Snyk-)\).
* To run Snyk on Alpine Linux, first install libstdc++. See [this doc](https://support.snyk.io/hc/en-us/articles/360001929038) for more help.

**Steps:**

Run these commands to install it for local use:

```text
npm install -g snyk
```

Once installed, you need to authenticate with your Snyk account:

```text
snyk auth
```

![](../../.gitbook/assets/uuid-7f427e54-45f8-910e-98c5-2016a27d29b0-en.gif)

To test your installation change directory into a folder containing a supported package manifest file \(package.json, pom.xml, composer.lock, etc.\) and run:

```text
cd /my/project/
snyk test
```

Alternatively, you can perform a quick test on a public npm package, for instance:

```text
snyk test ionic
```

As you can see, Snyk found and reported several vulnerabilities in the package. For each issue found, Snyk provides the severity of the issue, a link to a detailed description, the path through which the vulnerable module got into your system, and guidance on how to fix the problem.

## Install the Snyk CLI using the prebuilt binaries

You can download and use Snyk's prebuilt binaries already containing npm, the Snyk CLI and other necessary components. To download the prebuilt binary, visit the [**Releases tab**](https://github.com/snyk/snyk/releases) in the CLI repository page in GitHub:

![](../../.gitbook/assets/image%20%2812%29.png)

Once you've completed installation, get started testing and remediating your vulnerabilities with our [Getting started](https://support.snyk.io/hc/articles/360003812458#UUID-19fc37f2-b686-11ed-b85c-4789e90c8dfc) guide and [our full list](https://support.snyk.io/hc/articles/360003812578#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0) of our CLI commands, options and arguments.

## Install the Snyk CLI tool with a container

You can use a Snyk created Docker container already containing npm, the Snyk CLI and other necessary components. Follow the detailed instructions here [https://hub.docker.com/r/snyk/snyk-cli](https://hub.docker.com/r/snyk/snyk-cli)

Once you've completed installation, get started testing and remediating your vulnerabilities with our [Getting started](https://support.snyk.io/hc/articles/360003812458#UUID-19fc37f2-b686-11ed-b85c-4789e90c8dfc) guide and [our full list](https://support.snyk.io/hc/articles/360003812578#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0)[ ](https://snyk.gitbook.io/user-docs/snyk-cli/guides-for-our-cli/cli-reference)of our CLI commands, options and arguments.

## Install the Snyk CLI with Homebrew

From MAC OSx and Linux environments, you can use Homebrew to install our Snyk CLI tool. The repository for installation is stored in [our GitHub](https://github.com/snyk/homebrew-tap).

**Prerequisites:**

* Supported for MAC OSx and Linux environments only.
* Ensure [Homebrew](https://brew.sh/index_he) has already been installed:

  ```text
  brew tap snyk/tap
  ```

**Steps:**

1. Install Snyk as follows:

   ```text
   brew install snyk
   ```

## Install the Snyk CLI with the Windows Scoop package manager

From Windows environments, you can use Scoop to install our Snyk CLI tool. The repository for installation is stored in [our GitHub](https://github.com/snyk/scoop-snyk).

**Prerequisites:**

* Supported for Windows environments only.
* Ensure [Scoop](https://scoop.sh/) has already been installed:

  ```text
  scoop bucket add snyk https://github.com/snyk/scoop-snyk
  ```

**Steps:**

1. Install Snyk:

   ```text
   scoop install snyk
   ```

## See also

* [Getting started with the CLI](../guides-for-our-cli/getting-started-with-the-cli.md)
* [CLI reference](../guides-for-our-cli/cli-reference.md)

