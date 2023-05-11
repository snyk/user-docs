# Artifactory Registry setup

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
**Supported Projects**\
The Artifactory Package Repository integration currently supports [Node.js](../../../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-javascript/) (npm and Yarn) and [Maven](../../../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-java-gradle-maven.md) projects. Gradle projects are not currently supported.
{% endhint %}

Connecting a custom Artifactory Package Repository enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

You can configure these types of Artifactory Package Repository:

* Publicly accessible instances protected by basic authentication
* Instances on a private network accessed through Snyk Broker (with or without basic authentication).

These instructions apply to configuring publicly accessible instances. For instructions on configuring a brokered instance see [Set up Snyk Broker with Artifactory Repository](../../../snyk-admin/snyk-broker/snyk-broker-set-up-examples/set-up-snyk-broker-with-artifactory-repository.md).
