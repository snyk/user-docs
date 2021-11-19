# Test public npm packages before use

You can also use `snyk test` to **scrutinize a public package before installing it**, to see if it has known vulnerabilities or not. Using the package name will test the latest version of that package, and you can also provide a specific version or range using `snyk test module[@semver-range]`.

```
snyk test ionic@1.6.5
```

See also

* [Getting started with the CLI](../)
* [Our full CLI reference](../guides-for-our-cli/cli-reference.md)
* [Set severity thresholds for CLI tests](set-severity-thresholds-for-cli-tests.md)
* [Test your code from the CLI](../../../tutorials/springone-workshop/developer-environment-and-snyk/snyk-test-using-cli.md)
* [Test dev dependencies](https://docs.snyk.io/features/snyk-cli/test-for-vulnerabilities/use-snyk-open-source-from-the-cli#test-dev-dependencies)
* [Test public repositories before use](test-public-repositories-before-use.md)
