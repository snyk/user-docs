# Artifactory Registry for Maven

* [ Artifactory Registry Setup](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360013805638-Artifactory-Registry-Setup/README.md)
* [ Artifactory Registry for Maven](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360005507418-Artifactory-Registry-for-Maven/README.md)
* [ Artifactory Registry for npm](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007537418-Artifactory-Registry-for-npm/README.md)
* [ npm Teams & npm Enterprise for npm](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360009411777-npm-Teams-npm-Enterprise-for-npm/README.md)
* [ Private Gem Sources for Ruby](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360013742557-Private-Gem-Sources-for-Ruby/README.md)

## Artifactory Registry for Maven

**Feature availability**  
This feature is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Snyk can be configured to use custom package registries under specific conditions, allowing insight into dependencies that are not hosted in canonical registries.

The custom package registry feature currently supports [Artifactory](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360013805638/README.md) with Maven.

Maven analysis can be configured to mirror all requests through a custom package repository, or you can specify additional repositories to use alongside Maven Central.

### **Setup custom Maven package registries**

If authentication is required to access your custom registry you will need to first configure the Artifactory package repository integration, see [Artifactory Registry Setup.](https://support.snyk.io/hc/en-us/articles/360013805638)

Once the integration is set up you can configure Maven settings by navigating to settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; Languages &gt; Java.

You can choose whether to use Artifactory as a mirror or as an additional repository where your artifacts will reside. These settings will be very similar to what you have in `~/.m2/settings.xml`.

**Mirrors**

Choose a value for the Type, either 'Direct' or if using authentication 'Integration'

If using Direct you will need to complete the URL, Repository Name and what it is a Mirror Of.

The Mirror Of value can either be a \* to mirror everything or you can type in a value for example 'central'.

If using Type 'Integration', you will need to choose an integration type and provide the Repository Name and Mirror Of details.

The Repository Name should be set as whatever comes after 'artifactory/' in the internal repository URL.

For example, if the URL is '\[[http://artifactory.company.io/artifactory/libs-release'\[,\]\(http://artifactory.company.io/artifactory/jcenter\]\(http://artifactory.company.io/artifactory/libs-release'\[,\]%28http://artifactory.company.io/artifactory/jcenter\)',\](http://artifactory.company.io/artifactory/libs-release'[,]%28http://artifactory.company.io/artifactory/jcenter]%28http://artifactory.company.io/artifactory/libs-release'[,]%28http://artifactory.company.io/artifactory/jcenter%29',\)\) Repository Name should be set as 'libs-release'.

**Repositories**

Alternatively, you can configure repositories which will be used as additional locations to check for artifacts.

