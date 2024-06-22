# Java and Kotlin for open source

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## [Java and Kotlin](./) for open source

**Package manager**: Maven, Gradle

**Package registry**: [maven.org](https://maven.org/)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:maven`

**Test your app's packages**: Available, `pkg:maven`

**Features**:&#x20;

* Fix PRs (Maven)
* License scanning
* Reports

**Package manager versions**:&#x20;

* Maven
  * `3.*` For details, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
* Gradle
  * `4.*`, `5.*`, `6.*`, `7.*`\
    For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).

## Open source and licensing

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk Open Source provides full support for Java and Kotlin, as outlined below.

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

{% hint style="info" %}
Gradle Projects imported via Git are tested by parsing `build.gradle` files. You can resolve Gradle dependencies only by executing the tool itself, but even this method can sometimes provide incomplete results.

If possible, enable [lockfiles](java-and-kotlin-for-open-source.md#git-services-for-maven-and-gradle) in your Gradle Project to improve the accuracy of Git imports.

For Gradle projects without lockfiles, Snyk recommends using the Snyk CLI for the most accurate results.
{% endhint %}

{% tabs %}
{% tab title="Java" %}
**Java**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          | ✔︎          | ✔︎               | Fix advice only |
{% endtab %}

{% tab title="Kotlin" %}
**Kotlin**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          |             | ✔︎               | Fix advice only |
{% endtab %}
{% endtabs %}

#### Supported versions of Maven and Gradle

| Maven                                                                                                               | Gradle                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CLI - Maven `3.*` For details, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support). | CLI - Gradle `4.*`, `5.*`, `6.*`, `7.*`, `8.*` For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support). |
| Git - Maven `3.*`                                                                                                   | Git - Gradle `4.*`, `5.*`, `6.*`, `7.*`, `8.*`                                                                                                                |

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).
