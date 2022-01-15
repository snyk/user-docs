# Getting started with the CLI

The Snyk Command Line Interface (CLI) helps you find and fix known vulnerabilities in your dependencies, both manually and as part of your Continuous Integration (CI) build system.

See [Language and package manager support](../../../products/snyk-open-source/language-and-package-manager-support/) for details about package managers and languages that the CLI supports.

## Install and authenticate the Snyk CLI

1. Install the Snyk CLI using npm, Homebrew, Scoop or a manual installer from the Snyk GitHub repository. See [Install the Snyk CLI](../install-the-snyk-cli/).
2. To associate your Snyk account with the CLI, authenticate your account, either through your browser or using your API token. See [Authenticate the CLI with your account](../authenticate-the-cli-with-your-account/).

## Build your project

Before testing for vulnerabilities you must build your project, with limited exceptions. For details refer to [Language and package manager support](../../../products/snyk-open-source/language-and-package-manager-support/).

## Test for vulnerabilities

Once you have installed and authenticated, change directory into a folder containing a supported package manifest file (for example, `package.json`, `pom.xml`, `composer.lock`) and run:

`cd /my/project/`

`snyk test`

All vulnerabilities identified are listed, including their path, and fix guidance.

## Monitor your project

Snyk uses monitoring to regularly test your code and notify you when new vulnerabilities are introduced. To set up your project to be monitored:

1.  Open a terminal from your project directory and run `snyk monitor`.

    This command takes a snapshot of your current dependencies so Snyk can regularly scan your code, and notify you about newly disclosed vulnerabilities as they are introduced, or when a previously unavailable patch or upgrade path is created.
2. Log in to the Snyk app and navigate to the [Projects page](https://app.snyk.io/projects) to see the latest snapshot and scan results.

For more information see [Monitor your projects at regular intervals](../secure-your-projects-in-the-long-term/monitor-your-projects-at-regular-intervals.md).

## **Additional information**

Run `snyk help` or see the [CLI reference](../cli-reference/).

Snyk also provides a [handy cheat sheet](https://res.cloudinary.com/snyk/image/upload/v1551195097/Snyk\_CLI\_Cheat\_Sheet.pdf) and a [video tutorial](https://www.youtube.com/watch?v=xp\_LtchEkT8).
