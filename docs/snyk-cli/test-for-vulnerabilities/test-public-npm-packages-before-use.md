# Test public npm packages before use

You can also use `snyk test` to **scrutinize a public package before installing it**, to see if it has known vulnerabilities or not. Using the package name will test the latest version of that package, and you can also provide a specific version or range using `snyk test module[@semver-range]`.

```text
snyk test ionic@1.6.5
```

See also

* [Getting started with the CLI](https://support.snyk.io/hc/articles/360003812458#UUID-19fc37f2-b686-11ed-b85c-4789e90c8dfc)
* [Our full CLI reference](https://support.snyk.io/hc/articles/360003812578#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0)
* [Set severity thresholds for CLI tests](https://support.snyk.io/hc/articles/360003851337#UUID-844be978-191f-c813-ccd0-2221c04ca510)
* [Test your code from the CLI](https://support.snyk.io/hc/articles/360003812478#UUID-2e8464f9-1d8a-5f79-466d-2ca5c5cf9f22)
* [Test dev dependencies](https://support.snyk.io/hc/articles/360003812478#UUID-1070ae3e-5f30-cb4e-6350-a1c3f5c67511)
* [Test public repositories before use](https://support.snyk.io/hc/articles/360003851277#UUID-ba99a73f-110d-1f1d-9e7a-1bad66bf0996)

