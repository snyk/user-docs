# Maven plugin integration

Snyk offers a [Maven plugin](https://github.com/snyk/snyk-maven-plugin), that's based on the [Snyk CLI](snyk-cli/guides-for-our-cli/cli-reference/). This plugins allows you to scan and monitor your Maven dependencies for vulnerabilities.

{% hint style="info" %}
[See the Snyk Maven Plugin repository for installation and usage instructions](https://github.com/snyk/snyk-maven-plugin/)
{% endhint %}

* The test goal presents a list of vulnerabilities in your project's dependencies, in either a developer's machine or in your CI process.
* The monitor goal records the state of dependencies and any vulnerabilities on `snyk.io` so you can be alerted when new vulnerabilities or updates/patches are disclosed that affect your repositories.
* Running `mvn snyk:test or mvn snyk:monitor` will run the desired goals \(either test or monitor\) outside the Maven build lifecycle.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

