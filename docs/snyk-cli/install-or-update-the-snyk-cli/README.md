# Install or update the Snyk CLI

You can install or update the [Snyk CLI](../) using the methods explained on this page.

After you install the Snyk CLI, you must [authenticate](../commands/auth.md). Then you can [get started](../getting-started-with-the-snyk-cli.md) testing and fixing your vulnerabilities, beginning with testing your installation.

{% hint style="info" %}
Snyk recommends always keeping your CLI installation updated to the latest version. You can check which version of the Snyk CLI you have installed by running `snyk --version`.
{% endhint %}

For information about installing the CLI for an IDE, see the IDE documentation.

## Install with standalone executables

Use [GitHub Releases](https://github.com/snyk/snyk/releases) to download a standalone executable (macOS, Linux, Windows) of Snyk CLI for your platform.

Snyk also provides these standalone executables on the Snyk Content Delivery Network (CDN). See the latest `release.json` [file](https://static.snyk.io/cli/latest/release.json) for the download links. Examples for a specific version or platform follow:

* [https://static.snyk.io/cli/v1.666.0/release.json](https://static.snyk.io/cli/v1.666.0/release.json)
* [https://static.snyk.io/cli/latest/snyk-macos](https://static.snyk.io/cli/latest/snyk-macos)

For example, to download and run the latest Snyk CLI on macOS, you could run:

```bash
curl --compressed https://static.snyk.io/cli/latest/snyk-macos -o snyk
chmod +x ./snyk
mv ./snyk /usr/local/bin/
```

You can also use these direct links to download the executables:

* **macOS**: [https://static.snyk.io/cli/latest/snyk-macos](https://static.snyk.io/cli/latest/snyk-macos)
* **Windows**: [https://static.snyk.io/cli/latest/snyk-win.exe](https://static.snyk.io/cli/latest/snyk-win.exe)\
  You can rename the file to snyk.exe so you can run snyk commands as documented, for example, `snyk test`.
* **Linux**: [https://static.snyk.io/cli/latest/snyk-linux](https://static.snyk.io/cli/latest/snyk-linux)
* **Linux/arm64**: [https://static.snyk.io/cli/latest/snyk-linux-arm64](https://static.snyk.io/cli/latest/snyk-linux-arm64)
* **Alpine**: [https://static.snyk.io/cli/latest/snyk-alpine](https://static.snyk.io/cli/latest/snyk-alpine)

{% hint style="info" %}
To use CLI releases before version 1.1230.0 on an Apple M1 or M2 machine, (darwin/arm64), see [Using CLI releases before version 1.1230.0 on an Apple M1 or M2 machine](using-cli-releases-before-version-1.1230.0-on-an-apple-m1-or-m2-machine.md).

To use the CLI with Alpine Linux, see [Prerequisites for CLI and Jenkins plugin on Alpine Linux operating system](prerequisites-for-cli-and-jenkins-plugin-on-alpine-linux-operating-system.md).

For more information see [verify CLI standalone binaries](verifying-cli-standalone-binaries.md).
{% endhint %}

{% hint style="warning" %}
The drawback of this method is that you must keep the Snyk CLI up to date manually.
{% endhint %}

## Install with Homebrew (macOS, Linux)

Install Snyk CLI from [Snyk's tap](https://github.com/snyk/homebrew-tap) with [Homebrew](https://brew.sh) by running the following. The tap is updated daily with the latest Snyk CLI release.

```bash
brew tap snyk/tap
brew install snyk
```

{% hint style="warning" %}
For Apple M1 or M2 (darwin/arm64), see: [Using CLI releases before version 1.1230.0 on an Apple M1 or M2 machine](using-cli-releases-before-version-1.1230.0-on-an-apple-m1-or-m2-machine.md).
{% endhint %}

## Install with Scoop (Windows)

Install Snyk CLI from [Snyk's bucket](https://github.com/snyk/scoop-snyk) with [Scoop](https://scoop.sh) by running the following. The bucket is updated daily with the latest Snyk CLI release.

```
scoop bucket add snyk https://github.com/snyk/scoop-snyk
scoop install snyk
```

## Install the Snyk CLI with npm or Yarn

Before installing the Snyk CLI using npm, be sure you have installed the **prerequisites**:

* Install the latest version of npm in your local environment, using Node version 12 or later. For information on how to update Node see [Install or upgrade to version of Node required for Snyk CLI](install-or-upgrade-to-version-of-node-required-for-snyk-cli.md).
* To run Snyk on Alpine Linux, first install libstdc++.\
  For more information see [Prerequisites for CLI and Jenkins plugin on Alpine Linux operating system](prerequisites-for-cli-and-jenkins-plugin-on-alpine-linux-operating-system.md).

Then follow these **steps to install with npm or Yarn**:

[Snyk CLI is available as an npm package](https://www.npmjs.com/package/snyk). If you have Node.js installed locally, you can **install** the npm package by running `npm install snyk -g`.

If you are using Yarn, **install** by running `yarn global add snyk`.

For additional information see [Installing Snyk CLI as a binary via npm](installing-snyk-cli-as-a-binary-via-npm.md).

## Snyk CLI in a Docker image

Snyk CLI can also be run from a Docker image. Snyk offers multiple Docker images under [snyk/snyk on Docker Hub](https://hub.docker.com/r/snyk/snyk). See [snyk/snyk-images on GitHub](https://github.com/snyk/snyk-images) for details.

These images wrap the Snyk CLI and depending on the Tag come with relevant tooling for different projects. An example follows for scanning a Gradle project with `snyk/snyk`:

```bash
docker run -it \
    -e "SNYK_TOKEN=<TOKEN>" \
    -v "<PROJECT_DIRECTORY>:/project" \
    -v "/home/user/.gradle:/home/node/.gradle" \
  snyk/snyk:gradle:6.4 test --org=my-org-name
```

This is an example for scanning a Maven project with `snyk/snyk`:

```
docker run --rm -r \
-e SNYK_TOKEN=<YOUR_SNYK_TOKEN> \
-v <PROJECT_DIRECTORY>:/app \
-v <PROJECT_DIRECTORY>/settings.xml:/root/.m2/settings.xml \
snyk/snyk:maven snyk monitor \
--all-projects=true \
--maven-aggregate-project
```

\
