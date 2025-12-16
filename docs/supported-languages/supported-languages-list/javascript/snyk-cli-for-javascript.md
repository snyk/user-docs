# Snyk CLI for JavaScript

To help generate reports locally or at build time, see the [snyk-to-html plugin](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).&#x20;

See `--json` and `--sarif` options for generating output that can be programmatically accessed.

For advanced filtering options, see[ snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).

## Open Source libraries

The `snyk test` command tests the first manifest it can find and performs a test on that singular entry point. To have Snyk analyze all manifests in the directory, use the following options:

* `--all-projects`: This option detects and scans all Yarn and other Projects in this directory.
* `--yarn-workspaces`: For Yarn Workspaces use the `--all-projects` flag to test and monitor your packages with other package managers or Yarn workspaces or use `--yarn-workspaces` to specifically scan Yarn Workspaces Projects only.

{% hint style="info" %}
If you are using a package manager that requires options, it is suggested to target them individually with `--file=`
{% endhint %}

### Codebase

* Framework support - see [Supported languages, frameworks, and feature availability overview](../../supported-languages-package-managers-and-frameworks.md).
* Use the `snyk code test` command from the root of the Project to perform source code analysis.

### Containers

* Snyk will automatically look for application (open source) vulnerabilities as part of a container scan. Consider having Snyk integrated through CLI earlier in the pipeline and utilize this for an additional signal of and insight into what is in production.
* If you ship your Node.JS application in a container, be aware that you might also be bundling insecure packages (Linux, open source), alongside your application in addition to what is brought in by the container base image. The Snyk Container CLI can help you identify a base image that minimizes the attack surface of your application.
* For more information on how you can filter to the layer you wish to work on, such as identifying a secure base image to build off of, the layers you are responsible for, or application (OS) vulnerabilities, see [Snyk CLI for container security](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)

### Infrastructure as Code

See [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/).

## Resources

See the [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md) and the [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/). Use the `--help` option in the CLI for details of Snyk CLI commands.
