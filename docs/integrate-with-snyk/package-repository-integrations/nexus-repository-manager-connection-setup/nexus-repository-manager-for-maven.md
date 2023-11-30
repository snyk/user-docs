# Nexus Repository Manager for Maven

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can use Nexus Repository Manager with Maven Projects.

This enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

Maven Projects can be configured to mirror all requests through a custom package repository, or you can specify additional repositories to use alongside Maven Central.

## **Setup custom Maven package registries**

If authentication is required to access your Nexus registry, you must first configure the Nexus Repository Manager integration. See [Nexus Repository Manager setup](./).

Once the integration is set up, you can configure Maven settings by navigating to settings ![Settings icon](../../../.gitbook/assets/cog\_icon.png) **> Languages > Java**.

You can choose whether to use Nexus as a mirror or as an additional repository where your artifacts will reside.

These settings are very similar to what you have in `~/.m2/settings.xml`.

### **Mirrors**

<figure><img src="../../../.gitbook/assets/Screenshot 2022-07-15 at 15.10.52.png" alt="Set up for Mirrors"><figcaption><p>Set up for Mirrors</p></figcaption></figure>

Choose a value for the Type, either **Direct** or, if you are using using authentication, **Integration**.

If you are using **Direct,** you must complete the **URL**, **Repository Name** and what it is a **Mirror Of**.

The **Mirror Of** value can either be a `*` to mirror everything, or you can type in a value, for example,`central`.

If you are using Type **Integration**, you must choose the Nexus integration type and provide the **Repository Name** and **Mirror Of** details.

Set **Repository Name** according to your version of Nexus.

{% tabs %}
{% tab title="Nexus 3" %}
Set **Repository Name** as whatever comes after `repository/` in the internal repository URL.

For example, if the URL is `http://nexus.company.io/repository/libs-release` , Repository Name should be set as `libs-release`.
{% endtab %}

{% tab title="Nexus 2" %}
Set **Repository Name** as whatever comes after `nexus/content/` in the internal repository URL.

For example, if the URL is `http://nexus.company.io/nexus/content/groups/public`, Repository Name should be set as `groups/public`.

If the URL is `http://nexus.company.io/nexus/content/repositories/releases`, Repository Name should be set as `repositories/releases`.
{% endtab %}
{% endtabs %}

### **Repositories**

Alternatively, you can configure repositories which will be used as additional locations to check for artifacts.

Repositories are configured in the same way as [Mirrors](nexus-repository-manager-for-maven.md#mirrors), but do not require **Mirror Of**.
