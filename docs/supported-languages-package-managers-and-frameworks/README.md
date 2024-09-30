# Supported languages, package managers, and frameworks

This section provides detailed information about language support for Snyk Code and Snyk Open Source. For information about language support for Snyk Container, see [Supported workloads, container registries, languages, and operating systems](../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/supported-workloads-container-registries-languages-and-operating-systems.md). See also [Operating system distributions supported by Snyk Container](../scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md). For IaC language support, see [Supported IaC languages, cloud providers, and cloud resources](../scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/).

## Code analysis (Snyk Code)

Snyk supports the following languages for code analysis: [Apex](apex.md), [C++](c-c++/c++-for-code-analysis.md), [Go](go/go-for-code-analysis.md), [Java and Kotlin](java-and-kotlin/java-and-kotlin-for-code-analysis.md), [JavaScript](javascript/javascript-for-code-analysis.md), [.NET](.net/.net-for-code-analysis.md), [PHP](php/php-for-code-analysis.md), [Python](python/python-for-code-analysis.md), [Ruby](ruby/ruby-for-code-analysis.md), [Scala](scala/scala-for-code-analysis.md). [Swift](swift-and-objective-c/swift-for-code-analysis.md), [TypeScript](typescript.md), and [VB NET](vb.net.md).

Interfile analysis is available for all languages supported except Ruby.

The individual language pages for code analysis provide details about fully supported frameworks and features for Snyk Code.

## Open source and licensing (Snyk Open Source)

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk supports the following for Snyk Open Source: Bazel, C.C++, Dart and Flutter, Elixir, Go, Java and Kotlin, Javascript. .NET, PHP, Python, Ruby, Swift, Typescript, and VB.NET.

The individual language pages for Open Source support provide details about fully supported package managers and features supported for Snyk Open Source.

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of projects that have a package manager, this means a release to the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the Github repo.
{% endhint %}

### Open source policy

For information on managing dependencies and vulnerabilities from your developer workflows through the use of policies, see the following:

* [Defining a secure open-source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

### Open source license compliance

To check compliance for open source licenses, see [Snyk License Compliance Management](../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md).

## Supported languages and the CLI, SCM integrations, and CI/CD

### Steps to start using the CLI

* [Create a Snyk account](../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
* [Install Snyk CLI and authenticate your machine](../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
* [Set the default Organization for the `snyk test` or `snyk code test` commands ](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md)

### CLI for code analysis

To start testing your code using Snyk Code through the CLI, open your repository in a terminal and run `snyk code test`.

For information about customizing test options, running other commands, excluding directories and files, and viewing and exploring the results in different formats, see the following:

* [Available commands](../snyk-cli/commands/#available-commands)
* [Exclude directories and files from Snyk Code CLI tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Output test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Export test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)
* [snyk-to-html](../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)

After you have run `snyk code test`, you can:

* [Open a Fix PR ](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/)
* [Configure PR Checks](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### Snyk CLI for open source

Ensure you have installed the relevant package manager and you have included the relevant manifest files supported by Snyk before testing.

To test your Open Source Project for vulnerabilities, run the `snyk test` command.

### Steps to start using SCM integrations

* [Set up an integration](../getting-started/quickstart/set-up-an-integration.md).
* For details, see [Snyk SCM integrations](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/).
* For language-specific information, see [Git repositories with Maven and Gradle](java-and-kotlin/git-repositories-with-maven-and-gradle.md), [Git repositories and JavaScript](javascript/git-repositories-and-javascript.md), and [Git repositories and Python](python/git-repositories-and-python.md).
