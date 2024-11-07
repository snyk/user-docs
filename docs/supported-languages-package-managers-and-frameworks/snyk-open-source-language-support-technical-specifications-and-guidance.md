# Snyk Open Source language support technical specifications and guideance

Snyk for Open Source and licensing works as follows.

Snyk builds a dependency graph and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any packages anywhere in that tree.

After Snyk has built the dependencies tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

The way Snyk analyzes and builds the dependencies tree varies depending on the language and package manager of the Project, as well as the location of your Project.

For more information about language support for Snyk Open Source, see [Supported languages, package managers, and frameworks (Snyk Open Source)](./#open-source-and-licensing-snyk-open-source).

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of projects that have a package manager, this means a release to the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the GitHub repo.
{% endhint %}
