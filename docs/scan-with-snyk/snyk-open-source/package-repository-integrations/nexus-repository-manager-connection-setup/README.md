# Nexus repository manager connection setup

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).

**Supported projects**\
The Nexus Repository Manager integration supports [Node.js](../../../../supported-languages/supported-languages-list/javascript/#supported-package-managers-and-package-registries) (npm and Yarn) and [Maven](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-package-managers-and-package-registries) Projects. For [Improved Gradle SCM scanning](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning), use the Maven settings.
{% endhint %}

Connecting Nexus repository manager enables Snyk to resolve all direct and transitive dependencies of packages hosted on the Nexus registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

You can configure these types of Nexus repository manager:

* Publicly accessible instances protected by basic authentication
* Instances on a private network by using Snyk Broker (with or without basic authentication).

{% hint style="info" %}
**Versions supported**

* Nexus Repository Manager version 3.0+
* Nexus Repository Manager version 2.15+
{% endhint %}

These instructions apply to configuring publicly accessible instances. For instructions on configuring a brokered instance, see the [setup instructions for Snyk Broker with Nexus Repository Manager](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/).

## Set up publicly accessible instances

1. Navigate to **Settings** > **Integrations** > **Package Repositories** > **Nexus**
2. Verify that you see the screen to configure Nexus.

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 15.15.11.png" alt="Configure Nexus"><figcaption><p>Configure Nexus</p></figcaption></figure>

On the page to configure Nexus, enter the information for the version you are using.

{% tabs %}
{% tab title="Nexus 3" %}
* Enter the URL of your Nexus instance; this must end with `/repository`
* Enter the Username.
* Enter the Password.
* Click **Save**
{% endtab %}

{% tab title="Nexus 2" %}
* Enter the URL of your Nexus instance; this must end with `/nexus/content.`
* Enter the Username.
* Enter the Password.
* Click **Save.**
{% endtab %}
{% endtabs %}

## Nexus behind a reverse proxy

If your Nexus server is running behind a reverse proxy, for example, Nginx, the URL might not end with the default `/repository` for Nexus 3 or `/nexus/content` for Nexus 2, depending on what routes have been configured in the reverse proxy. If this is the case, ensure you use the URL configured in the reverse proxy.

Example: for Nexus 3, if `http://nexus.company.io/repository` is mapped to `http://nexus.company.io/my-company/my-repository`, use `http://nexus.company.io/my-company/my-repository`.

Example: for Nexus 2, if `http://nexus.company.io/nexus/content` is mapped to `http://nexus.company.io/my-nexus-content`, use `http://nexus.company.io/my-nexus-content`.

{% hint style="info" %}
A green success message appears if Snyk can contact your repository.

If you see a yellow warning message, check your URL and credentials and try again.
{% endhint %}

## Nexus user permissions

The Nexus user needs the following privileges, either as part of the Role or added individually:

{% tabs %}
{% tab title="Nexus 3" %}
* `nx-metrics-all` (for the [system status check endpoint](https://support.sonatype.com/hc/en-us/articles/226254487-System-Status-and-Metrics-REST-API))
* `nx-repository-view-[*-* | <ecosystem-repo-name>]-read`
* `nx-repository-view-[*-* | <ecosystem-repo-name>]-browse`
{% endtab %}

{% tab title="Nexus 2" %}
* `Status - Read`
* `All [<ecosystem>] Repositories - (read)`
* `[All Repositories | <repoName>] - (view)`
{% endtab %}
{% endtabs %}
