# Maven plugin integration

Snyk offers a [Maven plugin](https://github.com/snyk/snyk-maven-plugin), that's based on the [Snyk CLI](https://docs.snyk.io/snyk-cli/guides-for-our-cli/cli-reference). This plugins allows you to scan and monitor your Maven dependencies for vulnerabilities.

[See the Snyk Maven Plugin repository for installation and usage instructions](https://github.com/snyk/snyk-maven-plugin)

* The test goal presents a list of vulnerabilities in your project's dependencies, in either a developer's machine or in your CI process.
* The monitor goal records the state of dependencies and any vulnerabilities on `snyk.io` so you can be alerted when new vulnerabilities or updates/patches are disclosed that affect your repositories.
* Running `mvn snyk:test or mvn snyk:monitor` will run the desired goals \(either test or monitor\) outside the Maven build lifecycle.

