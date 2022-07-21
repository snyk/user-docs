# Nexus Repository Manager for Maven



{% hint style="info" %}
**Feature availability**\
****This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can use Nexus Repository Manager with Maven projects.

This enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

Maven projects can be configured to mirror all requests through a custom package repository, or you can specify additional repositories to use alongside Maven Central.

## **Setup custom Maven package registries**

If authentication is required to access your Nexus registry you will need to first configure the Nexus Repository Manager integration, see [nexus-repo-manager-setup.md](nexus-repo-manager-setup.md "mention")

Once the integration is set up you can configure Maven settings by navigating to settings ![](../../.gitbook/assets/cog\_icon.png) **> Languages > Java**.

You can choose whether to use Nexus as a mirror or as an additional repository where your artifacts will reside.&#x20;

These settings will be very similar to what you have in `~/.m2/settings.xml`.

### **Mirrors**

![](<../../.gitbook/assets/Screenshot 2022-07-15 at 15.10.52.png>)

Choose a value for the Type, either **Direct** or if using authentication **Integration**

If using **Direct** you will need to complete the **URL**, **Repository Name** and what it is a **Mirror Of**.

The **Mirror Of** value can either be a `*` to mirror everything or you can type in a value for example `central`.

If using Type **Integration**, you will need to choose the Nexus integration type and provide the **Repository Name** and **Mirror Of** details.

The **Repository Name** should be set according to your version of Nexus.

{% tabs %}
{% tab title="Nexus 3" %}
**Repository Name** should be set as whatever come after `repository/` in the internal repository URL.

For example, if the URL is `http://nexus.company.io/repository/libs-release` , Repository Name should be set as `libs-release`.
{% endtab %}

{% tab title="Nexus 2" %}
**Repository Name** should be set as whatever come after: `nexus/content/` in the internal repository URL.

For example, if the URL is `http://nexus.company.io/nexus/content/groups/public`, Repository Name should be set as `groups/public`.&#x20;

Or if the URL is `http://nexus.company.io/nexus/content/repositories/releases`, Repository should be set as `repositories/releases`.
{% endtab %}
{% endtabs %}

### **Repositories**

Alternatively, you can configure repositories which will be used as additional locations to check for artifacts.

Repositories are configured in the same way as [Mirrors](artifactory-registry-for-maven-1.md#mirrors), but do not require **Mirror Of**.
