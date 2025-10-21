# Artifactory registry for Maven

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).

**Supported projects**\
The Artifactory Package Repository integration supports [Node.js](../../../../supported-languages/supported-languages-list/javascript/#package-managers-and-supported-file-extensions) (npm and Yarn) and [Maven](../../../../supported-languages/supported-languages-list/java-and-kotlin/#package-managers-and-supported-file-extensions) Projects. For [Improved Gradle SCM scanning](../../../../supported-languages/supported-languages-list/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning), use the Maven settings on this page.
{% endhint %}

Snyk can use custom Artifactory Package Repositories with Maven Projects.

This enables Snyk to resolve all direct and transitive dependencies of packages hosted on the custom registry and calculate a more complete, accurate dependency graph and related vulnerabilities.

Maven Projects can be configured to mirror all requests through a custom package repository, or you can specify additional repositories to use alongside Maven Central.

## **Set up custom Maven package registries**

If authentication is required for access to your custom registry, you must configure the Artifactory package repository integration; see [Artifactory Package Repository connection setup](./).

After the integration is set up, you can configure Maven settings by navigating to **Settings** > **Languages** > **Java**.

You can choose whether to use Artifactory as a mirror or as an additional repository where your artifacts will reside. These settings will be very similar to what you have in `~/.m2/settings.xml`.

## Mirrors

<figure><img src="../../../../.gitbook/assets/image (244).png" alt=""><figcaption><p>Maven settings, choose Type</p></figcaption></figure>

Choose a value for the Type, either **Direct** or, if you are using authentication, **Integration.**

If you are using **Direct,** you must complete the **URL**, **Repository Name,** and what it is a **Mirror Of**.

The **Mirror Of** value can either be a `*` to mirror everything, or you can type in a value, for example, `central`.

If you are using the Type **Integration**, you must choose an integration type and provide the **Repository Name** and **Mirror Of** details.

Set the **Repository Name** as whatever comes after `artifactory/` in the internal repository URL.

For example, if the URL is `http://artifactory.company.io/artifactory/libs-release` **Repository Name** should be set as `libs-release`.

## **Additional repositories**

Alternatively, you can configure repositories that will be used as additional locations to check for artifacts.

Repositories are configured in the same way as [Mirrors](artifactory-registry-for-maven.md#mirrors) but do not require **Mirror Of**.
