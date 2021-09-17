# Artifactory Registry for Maven

{% hint style="info" %}
**Feature availability**  
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can be configured to use custom package registries under specific conditions, allowing insight into dependencies that are not hosted in canonical registries.

The custom package registry feature currently supports [Artifactory](integrations/private-registry-integrations/artifactory-registry-setup/) with Maven.

Maven analysis can be configured to mirror all requests through a custom package repository, or you can specify additional repositories to use alongside Maven Central.

## **Setup custom Maven package registries**

If authentication is required to access your custom registry you will need to first configure the Artifactory package repository integration, see [Artifactory Registry Setup.](integrations/private-registry-integrations/artifactory-registry-setup)

Once the integration is set up you can configure Maven settings by navigating to settings ![](../../.gitbook/assets/cog_icon.png/) &gt; Languages &gt; Java.

You can choose whether to use Artifactory as a mirror or as an additional repository where your artifacts will reside. These settings will be very similar to what you have in `~/.m2/settings.xml`.

**Mirrors**

Choose a value for the Type, either 'Direct' or if using authentication 'Integration'

If using Direct you will need to complete the URL, Repository Name and what it is a Mirror Of.

The Mirror Of value can either be a \* to mirror everything or you can type in a value for example 'central'.

![](../../.gitbook/assets/uuid-fd027725-33b3-7f12-a921-d7fba9cedad8-en.png)

If using Type **Integration**, you will need to choose an integration type and provide the Repository Name and Mirror Of details.

The Repository Name should be set as whatever comes after **artifactory/** in the internal repository URL.

For example, if the URL is '[http://artifactory.company.io/artifactory/libs-release](http://artifactory.company.io/artifactory/libs-release)' Repository Name should be set as **libs-release**.

![](../../.gitbook/assets/uuid-293cfd2b-2cd5-b8a3-0671-bf6d2798a3bc-en.png)

**Repositories**

Alternatively, you can configure repositories which will be used as additional locations to check for artifacts.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

