# Go for open source

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## Go for open source

**Package manager**: Go Modules, dep

**Package registry**: No single registry, multiple sources&#x20;

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:golang`

**Test your app's packages**: Available, `pkg:golang`

**Features**:&#x20;

* License scanning
* Reports

**Package manager versions**: NA

## Open source and licensing

{% hint style="warning" %}
Beginning on January 1 2023 Snyk no longer supports govendor Projects. As a general security best practice, Snyk recommends using tools that are consistently maintained and up-to-date.

Now that Snyk no longer supports scanning of govendor Projects, a warning is issued and no results are provided.
{% endhint %}

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk supports testing and monitoring of Go Projects with dependencies managed by [Go Modules](https://golang.org/ref/mod) and [dep](https://github.com/golang/dep).

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Snyk License Compliance Management](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md).

#### Open source supported features

{% hint style="info" %}
**Feature availability**\
Features might not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features              | CLI support | Git support | License scanning | Fix PRs |
| ---------------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [Go Modules](https://golang.org/ref/mod) | ✔︎          | ✔︎          | ✔︎               |         |
| [dep](https://github.com/golang/dep)     | ✔︎          | ✔︎          | ✔︎               |         |

After Snyk has built the dependencies tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

To scan your dependencies, ensure you have first installed the relevant package manager, and that your Project contains the supported manifest files.

The way Snyk analyzes and builds the dependencies tree varies depending on the language and package manager of the Project, as well as the location of your Project.&#x20;
