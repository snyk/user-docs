# Getting started with the CLI

## Introduction to the Snyk CLI

[Snyk](https://snyk.io/) scans and monitors your software development projects for security vulnerabilities.

The Snyk CLI brings the functionality of Snyk into your development workflow. You can run the CLI locally, or in your CI/CD pipeline. The following shows an example of Snyk CLI test command output.

<figure><img src="../.gitbook/assets/snyk-cli-screenshot.png" alt="Snyk CLI test command output example"><figcaption><p>Snyk CLI test command output</p></figcaption></figure>

## About [Snyk](https://snyk.io/)

[Snyk](https://snyk.io/) is a developer-first cloud-native security tool. It covers multiple areas of application security:

1. [**Snyk Open Source**](https://snyk.io/product/open-source-security-management/): Find and automatically fix open source vulnerabilities
2. [**Snyk Code**](https://snyk.io/product/snyk-code/): Find and fix vulnerabilities in your application code in real time
3. [**Snyk Container**](https://snyk.io/product/container-vulnerability-management/): Find and fix vulnerabilities in container images and Kubernetes applications
4. [**Snyk Infrastructure as Code**](https://snyk.io/product/infrastructure-as-code-security/): Find and fix insecure configurations in Terraform and Kubernetes code

[Learn more about what Snyk can do and sign up for a free account »](https://snyk.io/)

## Supported languages and tools

Snyk CLI scanning for Open Source supports many languages and tools, including Java, JavaScript, Python, PHP, Ruby, Scala, .NET, Golang, and C/C++. For more information see [Open source language and package manager support](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/).

Snyk CLI also supports:

* [Snyk Code scanning](../scan-application-code/snyk-code/cli-for-snyk-code/)
* [Snyk Container scanning](../scan-containers/snyk-cli-for-container-security/) including Docker scanning
* [Snyk Infrastructure as Code scanning](../scan-cloud-deployment/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/) including Terraform and Kubernetes (K8s) projects

For additional information see the [summary of supported languages](https://docs.snyk.io/getting-started/introducing-snyk#how-can-snyk-work-in-my-environment).

## Install the Snyk CLI and authenticate

To use the CLI you must install it and authenticate. See [Install the Snyk CLI](install-the-snyk-cli.md) and the [Auth](commands/auth.md) command help. You can refer to the [release notes](https://github.com/snyk/cli/releases) for a summary of changes in each release. Before scanning your code, review the [Code execution warning for Snyk CLI](code-execution-warning-for-snyk-cli.md).

**Note:** Before you can use the CLI you must install your package manager. The needed third-party tools such as Gradle or Maven must be in the PATH.

## Test your installation

After authenticating you can **test your installation**. Change directory into a folder containing a supported package manifest file such as package.json, pom.xml, or composer.lock) (`cd /my/project/`) and run `snyk test`.

Alternatively, you can perform a **quick test** on a public npm package, for example, `snyk test ionic`.

Look at the `test` command **report** in your terminal. The report shows the vulnerabilities Snyk found in the package. For each issue found, Snyk reports the severity of the issue, provides a link to a detailed description, reports the path through which the vulnerable module got into your system, and provides guidance on how to fix the problem.

{% hint style="info" %}
Before using the Snyk CLI to test your project for vulnerabilities you must **build your project**, with limited exceptions (for details see [Which projects must be built before testing with CLI?](https://support.snyk.io/hc/en-us/articles/360015552617-Which-projects-must-be-built-before-testing-with-CLI-)). Depending on the language of your project, you may need to **set up your language environment** before using the Snyk CLI. For details refer to [Open Source language and package manager support](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/).
{% endhint %}

## Scan your development project

Once you have installed and authenticated, change directory into a folder containing a supported package manifest file (for example, `package.json`, `pom.xml`, `composer.lock`) using `cd /my/project/`. Then run `snyk test`. All vulnerabilities identified are listed, including their path and fix guidance.

You can also scan a Docker image by its tag with [Snyk Container](https://snyk.io/product/container-vulnerability-management/) by running, for example: `snyk container test ubuntu:18.04`.

To test a Kubernetes (K8s) file run `snyk iac test /path/to/kubernetes_file.yaml`.

If you are using the Snyk CLI to scan for issues in Snyk Code, Container, or Infrastructure as Code projects, refer to the instructions for using the CLI that are specific to each product:

* [Using Snyk Code via the CLI](../scan-application-code/snyk-code/cli-for-snyk-code/)
* [Snyk CLI for container security](../scan-containers/snyk-cli-for-container-security/)
* [Snyk CLI for Infrastructure as Code](../scan-cloud-deployment/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/)

## Monitor your open source or container project

Snyk can monitor your project periodically and alert you to new vulnerabilities. To set up your project to be monitored run `snyk monitor` or `snyk container monitor`.

This creates a snapshot of your current dependencies so Snyk can regularly scan your code. Snyk can then alert you about newly disclosed vulnerabilities as they are introduced or when a previously unavailable patch or upgrade path is created. The following code shows an example of the output of the `snyk monitor` command.

```
> snyk monitor
Monitoring /project (project-name)...

Explore this snapshot at https://app.snyk.io/org/my-org/project/29361c2c-9005-4692-8df4-88f1c040fa7c/history/e1c994b3-de5d-482b-9281-eab4236c851e

Notifications about newly disclosed issues related to these dependencies will be emailed to you.
```

You can log in to the Snyk Web UI, and navigate to the [Projects page](https://app.snyk.io/projects) to see the latest snapshot and scan results:

<figure><img src="../.gitbook/assets/monitor (1).png" alt="Snyk monitor snapshot and scan results"><figcaption><p>Snyk monitor snapshot and scan results</p></figcaption></figure>

For more information see [Monitor your projects at regular intervals](test-for-vulnerabilities/monitor-your-projects-at-regular-intervals.md).

## Running out of tests

If you run out of tests on an open source project follow these steps:

* Run `snyk monitor`.
* Open the Snyk UI and go to the **settings** of the project.
* Enter the URL of your open source repo in **Git remote URI**.

## Additional information

Run `snyk help` or see the [CLI commands and options summary](cli-reference.md).

See [Introduction to the Snyk CLI](https://training.snyk.io/courses/intro-cli) for a quick video training session.

Snyk also provides a [cheat sheet](https://snyk.io/wp-content/uploads/cheat-sheet-snyk-cli-v3.pdf) ([blog post](https://snyk.io/blog/snyk-cli-cheat-sheet/)) and a [video tutorial](https://www.youtube.com/watch?v=xp\_LtchEkT8).

In particular see the information about the following options that you may find useful:

* `--severity-threshold=low|medium|high|critical`: Report only vulnerabilities of the specified level or higher
* `--json`: Print results in JSON format
* `--all-projects`: Auto-detect all projects in the working directory

## Getting support

[Submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support whenever you need help with the Snyk CLI or Snyk in general. Note that Snyk support does not actively monitor GitHub Issues on any [Snyk project](https://github.com/snyk).

## Contributing

The Snyk CLI project is open source, but Snyk does not encourage outside contributors.

You may look into [design decisions for the Snyk CLI](https://github.com/snyk/snyk/blob/master/help/\_about-this-project/README.md).

The Snyk CLI repository is a monorepo that also covers other projects and tools at this time:\
[@snyk/protect](https://github.com/snyk/snyk/tree/master/packages/snyk-protect); [npm package for snyk-protect command](https://www.npmjs.com/package/@snyk/protect)

## Security

For any security issues or concerns, see the [SECURITY.md](https://github.com/snyk/snyk/blob/master/SECURITY.md) file in the GitHub repository.
