# Install or update the Snyk CLI

You can install the [Snyk CLI](./) using the methods explained on this page.

After you install the Snyk CLI, you must [authenticate](commands/auth.md). Then you can [get started](getting-started-with-the-cli.md) testing and fixing your vulnerabilities, beginning with testing your installation.

## Install the Snyk CLI with npm or Yarn

Before installing the Snyk CLI using npm, be sure you have installed the **prerequisites**:

* Install the latest version of npm in your local environment, using Node version 12 or later. See [What version of Node is required for Snyk?](https://support.snyk.io/hc/en-us/articles/360004183317-What-version-of-Node-is-required-for-Snyk-) for the steps to update Node.
* To run Snyk on Alpine Linux, first install libstdc++. See [How can I use CLI on an Alpine operating system?](https://support.snyk.io/hc/en-us/articles/360001929038) for more information.

Then follow these **steps to install with npm or Yarn**:

[Snyk CLI is available as an npm package](https://www.npmjs.com/package/snyk). If you have Node.js installed locally, you can **install** the npm package by running `npm install snyk -g`.

If you are using Yarn, **install** by running `yarn global add snyk`.

## Install with standalone executables

Use [GitHub Releases](https://github.com/snyk/snyk/releases) to download a standalone executable (macOS, Linux, Windows) of Snyk CLI for your platform.

Snyk also provides these standalone executables on the Snyk Content Delivery Network (CDN). See [https://static.snyk.io/cli/latest/release.json](https://static.snyk.io/cli/latest/release.json) for the download links. Examples for a specific version or platform follow:

* [https://static.snyk.io/cli/v1.666.0/release.json](https://static.snyk.io/cli/v1.666.0/release.json)
* [https://static.snyk.io/cli/latest/snyk-macos](https://static.snyk.io/cli/latest/snyk-macos)

For example, to download and run the latest Snyk CLI on macOS, you could run:

```bash
curl https://static.snyk.io/cli/latest/snyk-macos -o snyk
chmod +x ./snyk
mv ./snyk /usr/local/bin/
```

You can also use these direct links to download the executables:

* **macOS**: [https://static.snyk.io/cli/latest/snyk-macos](https://static.snyk.io/cli/latest/snyk-macos)
* **Windows**: [https://static.snyk.io/cli/latest/snyk-win.exe](https://static.snyk.io/cli/latest/snyk-win.exe)
* **Linux**: [https://static.snyk.io/cli/latest/snyk-linux](https://static.snyk.io/cli/latest/snyk-linux)
* **Linux/arm64**: [https://static.snyk.io/cli/latest/snyk-linux-arm64](https://static.snyk.io/cli/latest/snyk-linux-arm64)
* **Alpine**: [https://static.snyk.io/cli/latest/snyk-alpine](https://static.snyk.io/cli/latest/snyk-alpine)

{% hint style="warning" %}
For Apple M1 (darwin/arm64), see: [How do I run Snyk CLI on an Apple M1 machine?](https://support.snyk.io/hc/en-us/articles/5022278090397)
{% endhint %}

{% hint style="warning" %}
**Note:** The drawback of this method is that you must keep the Snyk CLI up to date manually.
{% endhint %}

## Install with Homebrew (macOS, Linux)

Install Snyk CLI from [Snyk's tap](https://github.com/snyk/homebrew-tap) with [Homebrew](https://brew.sh) by running the following. The tap is updated daily with the latest Snyk CLI release.

```bash
brew tap snyk/tap
brew install snyk
```

{% hint style="warning" %}
For Apple M1 (darwin/arm64), see: [How do I run Snyk CLI on an Apple M1 machine?](https://support.snyk.io/hc/en-us/articles/5022278090397)
{% endhint %}

## Install with Scoop (Windows)

Install Snyk CLI from [Snyk's bucket](https://github.com/snyk/scoop-snyk) with [Scoop](https://scoop.sh) by running the following. The bucket is updated daily with the latest Snyk CLI release.

```
scoop bucket add snyk https://github.com/snyk/scoop-snyk
scoop install snyk
```

## Snyk CLI in a Docker image

Snyk CLI can also be run from a Docker image. Snyk offers multiple Docker images under [snyk/snyk-cli](https://hub.docker.com/r/snyk/snyk-cli) and [snyk/snyk](https://hub.docker.com/r/snyk/snyk) (see [snyk/snyk-images on GitHub](https://github.com/snyk/snyk-images) for more details).

These images wrap the Snyk CLI and depending on the Tag come with relevant tooling for different projects, for example, for scanning a Gradle project with snyk/snyk-cli:

```bash
docker run -it \
    -e "SNYK_TOKEN=<TOKEN>" \
    -v "<PROJECT_DIRECTORY>:/project" \
    -v "/home/user/.gradle:/home/node/.gradle" \
  snyk/snyk-cli:gradle-5.4 test --org=my-org-name
```

## Install as a part of a Snyk integration

Snyk also offers many [integrations](../integrations/) into developer tooling. These integrations install and manage the Snyk CLI for you. Integrations include the following:

* [Snyk Jenkins plugin](https://github.com/jenkinsci/snyk-security-scanner-plugin)
* [CircleCI Orb](https://github.com/snyk/snyk-orb)
* [Azure Pipelines Task](https://github.com/snyk/snyk-azure-pipelines-task)
* [GitHub Actions](https://github.com/snyk/actions)
* [IntelliJ IDE Plugin](https://github.com/snyk/snyk-intellij-plugin)
* [VS Code Extension](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner)
* [Eclipse IDE Extension](https://github.com/snyk/snyk-eclipse-plugin)
* [Maven plugin](https://github.com/snyk/snyk-maven-plugin)

See the [integrations](../integrations/) docs for more details.
