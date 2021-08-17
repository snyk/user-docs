# Artifactory Registry for Maven

**Feature availability**  
This feature is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Snyk can be configured to use custom package registries under specific conditions, allowing insight into dependencies that are not hosted in canonical registries.

The custom package registry feature currently supports [Artifactory](https://support.snyk.io/hc/en-us/articles/360013805638) with Maven.

Maven analysis can be configured to mirror all requests through a custom package repository, or you can specify additional repositories to use alongside Maven Central.

## **Setup custom Maven package registries**

If authentication is required to access your custom registry you will need to first configure the Artifactory package repository integration, see [Artifactory Registry Setup.](https://support.snyk.io/hc/en-us/articles/360013805638)

Once the integration is set up you can configure Maven settings by navigating to settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; Languages &gt; Java.

You can choose whether to use Artifactory as a mirror or as an additional repository where your artifacts will reside. These settings will be very similar to what you have in `~/.m2/settings.xml`.

**Mirrors**

Choose a value for the Type, either 'Direct' or if using authentication 'Integration'

If using Direct you will need to complete the URL, Repository Name and what it is a Mirror Of.

The Mirror Of value can either be a \* to mirror everything or you can type in a value for example 'central'.

![image2.png](https://support.snyk.io/hc/article_attachments/360007146318/uuid-fd027725-33b3-7f12-a921-d7fba9cedad8-en.png)

If using Type 'Integration', you will need to choose an integration type and provide the Repository Name and Mirror Of details.

The Repository Name should be set as whatever comes after 'artifactory/' in the internal repository URL.

For example, if the URL is '[http://artifactory.company.io/artifactory/libs-release'\[,\]\(http://artifactory.company.io/artifactory/jcenter](http://artifactory.company.io/artifactory/libs-release'[,]%28http://artifactory.company.io/artifactory/jcenter)',\) Repository Name should be set as 'libs-release'.

![image1.png](https://support.snyk.io/hc/article_attachments/360007064697/uuid-293cfd2b-2cd5-b8a3-0671-bf6d2798a3bc-en.png)

**Repositories**

Alternatively, you can configure repositories which will be used as additional locations to check for artifacts.

