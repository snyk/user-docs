# Java and Kotlin for open source

{% hint style="warning" %}
**Release status: Improved Gradle SCM scanning - Early Access**

You can now obtain more accurate results for your Gradle Projects imported through Git integrations by using Improved Gradle SCM scanning. this Early Access feature. For more information, see [Git repositories with Maven and Gradle](git-repositories-with-maven-and-gradle.md).
{% endhint %}

**Build tools**: Maven, Gradle

**Build tool versions**:&#x20;

* Maven: `3.*.` For more information, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
* Gradle: `4.*`, `5.*`, `6.*`, `7.*, 8.*.` For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).

**Package registry**: [maven.org](https://maven.org/)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:maven`

**Test your app's packages**: Available, `pkg:maven`

**Features**:&#x20;

* Fix PRs (Maven)
* License scanning
* Reports

The following summarizes Snyk support for Java and Kotlin.

## **Java**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          | ✔︎          | ✔︎               | Fix advice only |

## **Kotlin**

| Package managers / Features       | CLI support | Git support  | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ------------ | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎           | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          | Early Access | ✔︎               | Fix advice only |



