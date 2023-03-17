# Artifactory Registry setup

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
The Artifactory Package Repository integration currently supports [Node.js](../../products/snyk-open-source/language-and-package-manager-support/snyk-for-javascript/) (npm and Yarn) and [Maven](../../products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven.md) projects. Gradle projects are not currently supported.
{% endhint %}

Connecting a custom Artifactory Package Repository enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

You can configure two types of Artifactory Package Repository:

1. Publicly accessible instances protected by basic authentication
2. Instances on a private network accessed through Snyk Broker (with or without basic authentication).

These instructions apply to configuring publicly accessible instances. For instructions on configuring a brokered instance see [Set up Snyk Broker with Artifactory](../../snyk-admin/snyk-broker/snyk-broker-set-up-examples/set-up-snyk-broker-with-artifactory.md).

1. Go to settings <img src="../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> > **Integrations > Package Repositories > Artifactory**.
2. Enter the URL of your Artifactory instance, this **must** end with **/artifactory**.
3. Enter your username and password.
4. Select **Save**.

<figure><img src="../../.gitbook/assets/screenshot_2020-04-17_at_14.38.12.png" alt="Artifactory integration setup"><figcaption><p>Artifactory integration setup</p></figcaption></figure>

{% hint style="info" %}
If you do not see the **Snyk Broker** on/off switch, you do not have the necessary permissions and can only add a publicly accessible instance.

Submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new) if you want to add a private registry.
{% endhint %}

When you have permissions to add a private registry, continue with the instructions on [Set up Snyk Broker with Artifactory](../../snyk-admin/snyk-broker/snyk-broker-set-up-examples/set-up-snyk-broker-with-artifactory.md).
