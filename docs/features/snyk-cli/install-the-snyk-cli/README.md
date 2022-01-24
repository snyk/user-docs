# Install the Snyk CLI

Use our Snyk CLI tool to find and fix known vulnerabilities in your dependencies, both ad hoc and as part of your CI (Build) system.

You can install the Snyk CLI using multiple [installation methods](./#snyk-cli-installation-methods).

After you install the Snyk CLI, you can [get started](../getting-started-with-the-cli/) testing and fixing your vulnerabilities.

## Snyk CLI Installation methods

### Install the Snyk CLI with npm or Yarn

Install our Snyk CLI tool using npm.

**Prerequisites:**

* Ensure you’ve installed the latest version of npm on your local environment, using Node version 10 or later (see [What version of Node is required for Snyk?](https://support.snyk.io/hc/en-us/articles/360004183317-What-version-of-Node-is-required-for-Snyk-)).
* To run Snyk on Alpine Linux, first install libstdc++. See [How can I use CLI on an Alpine operating system?](https://support.snyk.io/hc/en-us/articles/360001929038) for more help.

**Steps:**

[Snyk CLI is available as an npm package](https://www.npmjs.com/package/snyk). If you have Node.js installed locally, you can install it by running:

```bash
npm install snyk@latest -g
```

or if you are using Yarn:

```bash
yarn global add snyk
```

Once installed, you need to authenticate with your Snyk account:

```
snyk auth
```

To test your installation change directory into a folder containing a supported package manifest file (package.json, pom.xml, composer.lock, etc.) and run:

```
cd /my/project/
snyk test
```

Alternatively, you can perform a quick test on a public npm package, for instance:

```
snyk test ionic
```

As you can see, Snyk found and reported several vulnerabilities in the package. For each issue found, Snyk provides the severity of the issue, a link to a detailed description, the path through which the vulnerable module got into your system, and guidance on how to fix the problem.

### Install with standalone executables

Use [GitHub Releases](https://github.com/snyk/snyk/releases) to download a standalone executable (macOS, Linux, Windows) of Snyk CLI for your platform.

We also provide these standalone executables on our official CDN. See [the `release.json` file](https://static.snyk.io/cli/latest/release.json) for the download links:

```http
https://static.snyk.io/cli/latest/release.json

# Or for specific version or platform
https://static.snyk.io/cli/v1.666.0/release.json
https://static.snyk.io/cli/latest/snyk-macos
```

For example, to download and run the latest Snyk CLI on macOS, you could run:

```bash
curl https://static.snyk.io/cli/latest/snyk-macos -o snyk
chmod +x ./snyk
mv ./snyk /usr/local/bin/
```

#### Direct download methods

You can also use these direct links to download the executables:

* **macOS**: [https://static.snyk.io/cli/latest/snyk-macos](https://static.snyk.io/cli/latest/snyk-macos)
* **Windows**: [https://static.snyk.io/cli/latest/snyk-win.exe](https://static.snyk.io/cli/latest/snyk-win.exe)
* **Linux**: [https://static.snyk.io/cli/latest/snyk-linux](https://static.snyk.io/cli/latest/snyk-linux)
* **Alpine**: [https://static.snyk.io/cli/latest/snyk-alpine](https://static.snyk.io/cli/latest/snyk-alpine)

{% hint style="warning" %}
The drawback of this method is that you will have to manually keep the Snyk CLI up to date.
{% endhint %}

### Install with Homebrew (macOS, Linux)

Install Snyk CLI from [Snyk tap](https://github.com/snyk/homebrew-tap) with [Homebrew](https://brew.sh) by running:

```bash
brew tap snyk/tap
brew install snyk
```

### Scoop (Windows)

Install Snyk CLI from our [Snyk bucket](https://github.com/snyk/scoop-snyk) with [Scoop](https://scoop.sh) on Windows:

```
scoop bucket add snyk https://github.com/snyk/scoop-snyk
scoop install snyk
```

### Snyk CLI in a Docker image

Snyk CLI can also be run from a Docker image. Snyk offers multiple Docker images under [snyk/snyk-cli](https://hub.docker.com/r/snyk/snyk-cli) and [snyk/snyk](https://hub.docker.com/r/snyk/snyk) ([snyk/images on GitHub](https://github.com/snyk/snyk-images) for more details).

These images wrap the Snyk CLI and depending on the Tag come with a relevant tooling for different projects. For example scanning a Gradle project with snyk/snyk-cli:

```bash
docker run -it
    -e "SNYK_TOKEN=<TOKEN>"
    -e "USER_ID=1234"
    -v "<PROJECT_DIRECTORY>:/project"
    -v "/home/user/.gradle:/home/node/.gradle"
  snyk/snyk-cli:gradle-5.4 test --org=my-org-name
```

### Install as a part of a Snyk integration

Snyk also offers many [integrations](../../integrations/) into developer tooling. These integrations install and manage the Snyk CLI for you. Integrations include the following:

* [Snyk Jenkins plugin](https://github.com/jenkinsci/snyk-security-scanner-plugin)
* [CircleCI Orb](https://github.com/snyk/snyk-orb)
* [Azure Pipelines Task](https://github.com/snyk/snyk-azure-pipelines-task)
* [GitHub Actions](https://github.com/snyk/actions)
* [IntelliJ IDE Plugin](https://github.com/snyk/snyk-intellij-plugin)
* [VS Code Extension](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner)
* [Eclipse IDE Extension](https://github.com/snyk/snyk-eclipse-plugin)
* [Maven plugin](https://github.com/snyk/snyk-maven-plugin)

See the [integrations](../../integrations/) docs for more details.
