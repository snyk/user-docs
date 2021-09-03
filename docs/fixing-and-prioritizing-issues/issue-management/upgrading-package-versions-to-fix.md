# Upgrading package versions to fix

Snyk will always recommend the smallest upgrade of a dependency to resolve the vulnerability. To resolve a vulnerability in a transitive dependency Snyk will calculate the dependency tree for your project and determine the minimum upgrade to the direct dependency which will result in a vulnerability free version of the indirect dependency.

Some remediations may require a major upgrade of a dependency. In this situation, we indicate on the Fix PR screen if we suspect a major change causing breakage.

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}