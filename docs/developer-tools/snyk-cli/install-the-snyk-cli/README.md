# Install the Snyk CLI

The Snyk CLI can be installed across all major operating systems using several different methods, depending on your existing toolset:

* [Package managers](./#package-managers): Use tools like npm, Homebrew, or Scoop for a quick installation.
* [Docker](./#docker): Run Snyk in a container for a consistent scanning environment.
* [GitHub Actions and CI/CD](./#github-actions-and-ci-cd): Use pre-built integrations to automate scanning within your deployment pipelines.
* [Direct binary download](./#standalone-executables): Download a standalone executable for Windows, MacOS, or Linux.

After you install the CLI, authenticate your Snyk account to start scanning your Snyk Projects. For more information, visit [Authenticate to use the CLI](../authenticate-to-use-the-cli.md).

{% hint style="warning" %}
The Snyk CLI does not natively support Windows Subsystem for Linux (WSL). You can install and use the CLI in a WSL environment, but Snyk does not guarantee compatability. Consequently, Snyk integrations that rely on the CLI, such as IDE extensions and CI/CD plugins, do not natively support WSL.
{% endhint %}

## Common installation steps

Regardless of your choice of tool, there are common steps to ensuring your installation of the Snyk CLI is a success:

1. Install the Snyk CLI using your preferred method.
2. Verify the installed Snyk CLI version using `snyk --version`.
3. Authenticate the Snyk CLI to your account.

To learn how to upgrade the Snyk CLI for your chosen tool, visit Upgrade the Snyk CLI.

{% hint style="info" %}
To use CLI releases before version 1.1230.0 on an Apple M1 or M2 machine (darwin/arm64), visit [Using CLI releases before version 1.1230.0 on an Apple M1 or M2 machine](https://docs.snyk.io/developer-tools/snyk-cli/install-or-update-the-snyk-cli/using-cli-releases-before-version-1.1230.0-on-an-apple-m1-or-m2-machine).
{% endhint %}

## Installation methods

### Package managers

<details>

<summary>Install with Homebrew (MacOS and Linux)</summary>

{% hint style="info" %}
For Apple M1 or M2 (darwin/arm64), see: [Using CLI releases before version 1.1230.0 on an Apple M1 or M2 machine](https://docs.snyk.io/developer-tools/snyk-cli/install-or-update-the-snyk-cli/using-cli-releases-before-version-1.1230.0-on-an-apple-m1-or-m2-machine).
{% endhint %}

1. Install and verify Homebrew using any method specified in their documentation. Visit [Install Homebrew](https://brew.sh/) for more information.
2.  Install [Homebrew Tap for Snyk](https://github.com/snyk/homebrew-tap). The tap is updated daily with the latest Snyk CLI release.

    ```bash
    brew tap snyk/tap
    ```
3.  Install the Snyk CLI for Homebrew:

    ```bash
    brew install snyk
    ```
4.  Verify that the installation has succeeded:

    ```bash
    brew list snyk
    ```
5. Authenticate. For more details, visit [Authentication for npm, Node.js, Yarn, Homebrew, and Scoop](./#authentication-for-npm-node.js-yarn-homebrew-and-scoop).

</details>

<details>

<summary>Install with npm</summary>

{% hint style="info" %}
Ensure you have permission to install global npm packages and network access to the npm registry to have a successful installation.
{% endhint %}

1. For Snyk CLI version 1.853.0 or later, install the latest version of npm on your machine using Node.js version 12 or later and npm version 7 or later.
2.  Install the npm package:

    ```bash
    npm install snyk -g
    ```
3. Authenticate. For more details, visit [Authentication for npm, Node.js, Yarn, Homebrew, and Scoop](./#authentication-for-npm-node.js-yarn-homebrew-and-scoop).

</details>

<details>

<summary>Install with Yarn</summary>

1.  Install with Yarn by running the command:

    ```bash
    yarn global add snyk
    ```


2. Authenticate. For more details, visit [Authentication for npm, Node.js, Yarn, Homebrew, and Scoop](./#authentication-for-npm-node.js-yarn-homebrew-and-scoop).

</details>

<details>

<summary>Install with Scoop (Windows)</summary>

1.  Add the official Snyk bucket to your machine:

    ```powershell
    scoop bucket add snyk https://github.com/snyk/scoop-snyk
    ```
2.  Install the Snyk CLI:

    ```powershell
    scoop install snyk
    ```
3.  Verify that the installation has succeeded:

    ```powershell
    snyk --version
    ```
4. Authenticate. For more details, visit [Authentication for npm, Node.js, Yarn, Homebrew, and Scoop](./#authentication-for-npm-node.js-yarn-homebrew-and-scoop).

</details>

### Docker

{% hint style="info" %}
Snyk offers multiple Docker images under [snyk/snyk on Docker Hub](https://hub.docker.com/r/snyk/snyk). To learn more, visit [snyk/snyk-images on GitHub](https://github.com/snyk/snyk-images).
{% endhint %}

Snyk provides a universal image and specialized images for specific package managers:

* Universal: `snyk/snyk`
* Language specific: `snyk/snyk:node`, `snyk/snyk:maven`, `snyk/snyk:python`, and so on.

Download the pre-packaged tools using the following command, specifying the universal or language specific image of your choosing:

{% code title="Pull the Snyk universal image" %}
```bash
docker pull snyk/snyk
```
{% endcode %}

{% code title="Pull the Snyk language specific image" %}
```bash
docker pull snyk/snyk:node
```
{% endcode %}

Authenticate the connection. For more details, visit Authentication in Docker.

### GitHub actions and CI/CD

<details>

<summary>GitHub actions</summary>

1. In your Snyk account, navigate to **Account Settings** and copy your Auth token. Use a Service Account Token if you are an Enterprise plan user, to protect your pipeline from breaking if you as an individual leave the comapny.
2. Navigate to **Settings** > **Secrets and variables** > **Actions**, and create a **New repository secret** in the format:
   1. **Name**: `SNYK_TOKEN`
   2. **Secret**: Your Snyk Auth token or Service Account Token
3. In GitHub repository:
   1. Create a folder named `.github`.
   2. Within that folder, create another folder named `workflows`.
   3. Create a file named `snyk-scan.yml`. The entire filepath should be `.github/workflows/snyk.yml`.
   4.  Add your YAML code to this file and commit it:

       ```yaml
       name: Snyk Security Scan
       on: push
       jobs:
         snyk:
           runs-on: ubuntu-latest
           steps:
             - uses: actions/checkout@v4
             
             # Use the action matching your language (node, python, go, and so on)
             - name: Run Snyk to check for vulnerabilities
               uses: snyk/actions/node@master
               env:
                 SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
               with:
                 args: --severity-threshold=high
       ```

</details>

<details>

<summary>CI/CD</summary>

For detailed instructions on using the Snyk CLI with CI/CD integrations, visit [Snyk CI/CDs](../../snyk-ci-cd-integrations/).

</details>

### Direct binary download

{% hint style="warning" %}
When using standalone executables, you must keep the Snyk CLI up to date manually.
{% endhint %}

Snyk provides standalone executables on the Snyk Content Delivery Network (CDN). For the latest download links, visit the latest `release.json` [file](https://downloads.snyk.io/cli/stable/release.json).

For instructions on how to verify the shasum of downloaded binaries and their GPG signatures, visit [Verifying CLI standalone binaries](verifying-cli-standalone-binaries.md).

<details>

<summary>Install for MacOS and Linux</summary>

Use `curl` to download the binary, make it executable, and move it to your system path in one go.

For Intel/AMD (x86\_64) use the following commands:

```bash
# Download
curl --compressed https://downloads.snyk.io/cli/latest/snyk-linux -o snyk

# Make the file executable
chmod +x ./snyk

# Move to a folder in your PATH
sudo mv ./snyk /usr/local/bin/
```

For Apple Silicon (M1/M2/M3) use the following commands:

```bash
# Download
curl --compressed https://downloads.snyk.io/cli/latest/snyk-macos-arm64 -o snyk

# Make the file executable
chmod +x ./snyk

# Move to a folder in your PATH
sudo mv ./snyk /usr/local/bin/
```

</details>

<details>

<summary>Install for Alpine</summary>

Use `curl` to download the binary, make it executable, and move it to your system path in one go.

For Intel/AMD (x86\_64) use the following commands:

```bash
# Download

curl --compressed https://downloads.snyk.io/cli/stable/snyk-alpine -o snyk

# Make the file executable

chmod +x ./snyk

# Move to a folder in your PATH. May require sudo.

mv ./snyk /usr/local/bin/
```

For ARM64 use the following commands:

```bash
# Download

curl --compressed https://downloads.snyk.io/cli/stable/snyk-alpine-arm64 -o snyk

# Make the file executable

chmod +x ./snyk

# Move to a folder in your PATH. May require sudo.

mv ./snyk /usr/local/bin/
```

</details>

<details>

<summary>Install for Windows</summary>

For Windows, download the `.exe` and manually add its location to your system variables.

1. Download [snyk-win.exe](https://static.snyk.io/cli/latest/snyk-win.exe)
2. Rename the file from `snyk-win.exe` to just `snyk.exe`.
3. Move it to a permanent folder (for example, `C:\tools\snyk\`).
4. Add to `PATH`:
   1. Search for **Environment Variables** in your Start Menu.
   2. Under **System Variables**, find **Path** and click **Edit**.
   3. Add the path to the folder where you saved the `.exe` (e.g., `C:\tools\snyk\`).

</details>
