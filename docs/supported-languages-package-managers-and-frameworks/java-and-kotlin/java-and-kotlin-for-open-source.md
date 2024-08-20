# Java and Kotlin for open source

## Snyk support for Java and Kotlin for open source

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

### **Java**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          | ✔︎          | ✔︎               | Fix advice only |

### **Kotlin**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          |             | ✔︎               | Fix advice only |

Gradle Projects imported using a Git integration are tested by parsing `build.gradle` files. You can resolve Gradle dependencies only by executing the tool itself, but even this method can sometimes provide incomplete results.

If possible, enable [lockfiles](java-and-kotlin-for-open-source.md#git-services-for-maven-and-gradle) in your Gradle Project to improve the accuracy of Git imports.

For Gradle projects without lockfiles, Snyk recommends using the Snyk CLI for the most accurate results.

