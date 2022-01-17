# Auth

## Usage

`snyk auth [<OPTIONS>]`

## Description

The `snyk auth` command authenticates the Snyk CLI with a Snyk account. Running `$ snyk auth` opens a browser window with a prompt to log in to your Snyk account and authorize.

In some environments and configurations you must use the `<API_TOKEN>`; see [Authenticate the CLI with your account](https://docs.snyk.io/features/snyk-cli/authenticate-the-cli-with-your-account). Is is recommended that in a CI/CD environment you use the SNYK\_TOKEN environment variable; see [Configure the Snyk CLI](https://docs.snyk.io/features/snyk-cli/configure-the-snyk-cli).

## Debug

Use the `-d` option to output the debug logs.
