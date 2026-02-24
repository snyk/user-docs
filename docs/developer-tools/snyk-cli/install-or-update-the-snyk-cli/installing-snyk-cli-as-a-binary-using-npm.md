# Install the Snyk CLI with npm

Snyk is updating the CLI deployment using npm to support extensibility. The Extensible CLI now deploys as a binary using npm.

## Prerequisites

Ensure you have the following:

* Node.js v12 or higher
* npm v7 or higher
* Permission to install global npm packages
* Network access to the npm registry

To check your installed versions, run:

```bash
node -v
npm -v
```

## Install the CLI

Install the Snyk CLI globally:

```bash
npm install -g snyk
```

If you encounter permission issues on macOS or Linux, run:

```bash
npm install -g snyk --unsafe-perm
```

## Verify installation

Run the following commands to verify the version, authenticate, and test the installation:

```bash
snyk --version
snyk auth
snyk test
```

## Update the CLI

To update the CLI, run:

```bash
npm update -g snyk
```

Alternatively, install the latest version:

```bash
npm install -g snyk@latest
```

## Troubleshooting degraded CLI versions

This process includes a fallback mechanism. If the binary download fails, Snyk reverts to the previous method and displays the following warning:

```
------------------------------ Warning -------------------------------
You are currently running a degraded version of the Snyk CLI.
As a result, some features of the CLI will be unavailable.
For information on how to resolve this, please see this article: https://docs.snyk.io/snyk-cli/installing-snyk-cli-as-a-binary-via-npm
For any assistance, please check http://support.snyk.io/.
------------------------------- Warning -----
```

Follow the instructions in the error message. If the issue persists, try the following solutions:

### Verify your npm version

Ensure your npm version is supported and Node.js environment is up to date.

### Fix permissions

Permission issues can cause degraded installs. Ensure your global npm directory is writable or use the unsafe permissions flag:

```bash
npm install -g snyk --unsafe-perm
```

### Clear cache and reinstall

Try clearing the npm cache and reinstalling if the installation appears corrupted.

```bash
npm cache clean --force
npm uninstall -g snyk
npm install -g snyk
```

### Check network or proxy restrictions

Ensure npm can reach the public registry. Verify that a firewall or proxy does not block downloads.

### Use the legacy CLI

As a temporary workaround, use the legacy CLI:

```bash
snyk --legacy-cli
```

### Contact Snyk Support

If the CLI continues to run in degraded mode, try another installation method, such as standalone binaries.

If you need to contact Snyk Support, include the following information to help diagnose the issue:

* Full error message
* Output of `snyk --version`
* Node.js and npm versions
* OS and architecture
* Proxy or firewall details
