# PHP for open source

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## [PHP](../php.md) for open source

**Package manager**: Composer

**Package registry**: [packagist.org](https://packagist.org/)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:composer`

**Test your app's packages**: Available, `pkg:composer`

**Features**:&#x20;

* License scanning
* Reports

**Package manager versions**: NA

## Open source and licensing

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).

#### Open source supported features

{% hint style="info" %}
**Feature availability**\
Features might not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features         | <p>CLI</p><p>support</p> | <p>Git</p><p>support</p> | License scanning | Fix PRs |
| ----------------------------------- | ------------------------ | ------------------------ | ---------------- | ------- |
| [Composer](https://getcomposer.org) | ✔︎                       | ✔︎                       | ✔︎               |         |

After Snyk has built the dependencies tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

To scan your dependencies, you must ensure you have first installed the relevant package manager, and that your Project contains the supported manifest files.

The way Snyk analyzes and builds the dependencies tree varies depending on the language and package manager of the Project, as well as the location of your Project.&#x20;
