# Snyk Open Source language support technical specifications and guidance

## How Snyk for Open Source and licensing works

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk builds a dependency graph and (dependencies tree) and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in that tree.

The way Snyk analyzes and builds the dependencies tree varies depending on the language and package manager for the Project, as well as the location of the Project.

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of projects that have a package manager, this means a release to the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the GitHub repo.
{% endhint %}

## Snyk policies in Open Source

For information on managing dependencies and vulnerabilities from your developer workflows through the use of policies, see the following:

* [Defining a secure open-source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

## Open Source license compliance

To check compliance for open source licenses, see [Snyk License Compliance Management](../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md).
