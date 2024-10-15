# Artifactory Package Repository connection setup

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available with Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans/).

**Supported projects**\
The Artifactory Package Repository integration supports [Node.js](../../../../supported-languages-package-managers-and-frameworks/javascript/#supported-frameworks-and-package-managers) (npm and Yarn) and [Maven](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-frameworks-and-package-managers) Projects.&#x20;

For [Improved Gradle SCM scanning](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning-early-access), use the Maven settings.
{% endhint %}

Connecting a custom Artifactory Package Repository enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

You can configure these types of Artifactory Package Repository:

* Publicly accessible instances protected by basic authentication
* Instances on a private network by using Snyk Broker (with or without basic authentication).

These instructions apply to configuring publicly accessible instances. For instructions on configuring a brokered instance, see the [setup instructions for Snyk Broker with Artifactory Repository](../../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/artifactory-repository-install-and-configure-broker/).

The steps to set up Artifactory Repository Manager follow.

1. Navigate to **Settings** > **Integrations > Package Repositories > Artifactory**.
2. Enter the URL of your Artifactory instance; this must end with `/artifactory`.
3. Enter your username and password.
4. Select **Save**.

<figure><img src="../../../../.gitbook/assets/screenshot_2020-04-17_at_14.38.12.png" alt="Artifactory repository setup"><figcaption><p>Artifactoryrepository setup</p></figcaption></figure>
