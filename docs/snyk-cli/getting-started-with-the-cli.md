# Getting started with the CLI

To use the CLI you must install it and authenticate. See [Install the Snyk CLI](install-the-snyk-cli.md) and the [Auth](commands/auth.md) command help. You can refer to the [release notes](https://github.com/snyk/cli/releases) for a summary of changes in each release.

After authenticating you can **test your installation**. Change directory into a folder containing a supported package manifest file such as package.json, pom.xml, or composer.lock) (`cd /my/project/`) and run `snyk test`.

Alternatively, you can perform a **quick test** on a public npm package, for example, `snyk test ionic`.

Look at the `test` command **report** in your terminal. The report shows the vulnerabilities Snyk found in the package. For each issue found, Snyk reports the severity of the issue, provides a link to a detailed description, reports the path through which the vulnerable module got into your system, and provides guidance on how to fix the problem.

Before using the Snyk CLI to test your project for vulnerabilities you must **build your project**, with limited exceptions (for details see [Which projects must be built before testing with CLI?](https://support.snyk.io/hc/en-us/articles/360015552617-Which-projects-must-be-built-before-testing-with-CLI-)). Depending on the language of your project, you may need to **set up your language environment** before using the Snyk CLI. For details refer to [Open Source language and package manager support](../products/snyk-open-source/language-and-package-manager-support/).

If you are using the Snyk CLI to scan for issues in Snyk Code, Container, or Infrastructure as Code projects, refer to the instructions for using the CLI that are specific to each product:

* [Using Snyk Code via the CLI](../products/snyk-code/cli-for-snyk-code/)
* [Snyk CLI for container security](../products/snyk-container/snyk-cli-for-container-security/)
* [Snyk CLI for Infrastructure as Code](../products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/)

## Scan your project

Once you have installed and authenticated, change directory into a folder containing a supported package manifest file (for example, `package.json`, `pom.xml`, `composer.lock`) using `cd /my/project/`. Then run `snyk test`. All vulnerabilities identified are listed, including their path and fix guidance.

You can also scan a Docker image by its tag with [Snyk Container](https://snyk.io/product/container-vulnerability-management/) by running, for example: `snyk container test ubuntu:18.04`.

To test a Kubernetes (k8s) file run `snyk iac test /path/to/kubernetes_file.yaml`.

## Monitor your project

Snyk can monitor your project periodically and alert you to new vulnerabilities. To set up your project to be monitored run `snyk monitor`.

This creates a snapshot of your current dependencies so Snyk can regularly scan your code. Snyk can then alert you about newly disclosed vulnerabilities as they are introduced or when a previously unavailable patch or upgrade path is created. The following code shows an example of the output of the `snyk monitor` command.

```
> snyk monitor
Monitoring /project (project-name)...

Explore this snapshot at https://app.snyk.io/org/my-org/project/29361c2c-9005-4692-8df4-88f1c040fa7c/history/e1c994b3-de5d-482b-9281-eab4236c851e

Notifications about newly disclosed issues related to these dependencies will be emailed to you.
```

You can log in to the Snyk Web UI, and navigate to the [Projects page](https://app.snyk.io/projects) to see the latest snapshot and scan results:

![Snyk monitor snapshot and scan results](<../.gitbook/assets/monitor (1).png>)

For more information see [Monitor your projects at regular intervals](test-for-vulnerabilities/monitor-your-projects-at-regular-intervals.md).

## Running out of tests

If you run out of tests on an open source project follow these steps:

* Run `snyk monitor`.
* Open the Snyk UI and go to the **settings** of the project.
* Enter the URL of your open source repo in **Git remote URI**.

## Additional information

Run `snyk help` or see the [CLI reference](cli-reference.md)..

See [Introduction to the Snyk CLI](https://training.snyk.io/courses/intro-cli) for a quick video training session.

Snyk also provides a [cheat sheet](https://snyk.io/wp-content/uploads/cheat-sheet-snyk-cli-v3.pdf) ([blog post](https://snyk.io/blog/snyk-cli-cheat-sheet/)) and a [video tutorial](https://www.youtube.com/watch?v=xp\_LtchEkT8).

In particular see the information about the following options that you may find useful:

* `--severity-threshold=low|medium|high|critical`: Report only report vulnerabilities of the specified level or higher
* `--json`: Print results in JSON format
* `--all-projects`: Auto-detect all projects in the working directory

## Getting support

[Submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support whenever you need help with the Snyk CLI or Snyk in general. Note that Snyk support does not actively monitor GitHub Issues on any [Snyk project](https://github.com/snyk).

## Contributing

The Snyk CLI project is open source, but Snyk does not encourage outside contributors.

You may look into [design decisions for the Snyk CLI](https://github.com/snyk/snyk/blob/master/help/\_about-this-project/README.md).

The Snyk CLI repository is a monorepo that also covers other projects and tools at this time:

* [`@snyk/protect`](https://github.com/snyk/snyk/tree/master/packages/snyk-protect); [npm package for `snyk-protect` command](https://www.npmjs.com/package/@snyk/protect)

## Security

For any security issues or concerns, see the [SECURITY.md](https://github.com/snyk/snyk/blob/master/SECURITY.md) file in the GitHub repository.
