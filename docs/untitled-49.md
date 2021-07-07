# Test public npm packages before use

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Test for vulnerabilities

* [ Use Snyk Open Source from the CLI](/hc/en-us/articles/360003812478-Use-Snyk-Open-Source-from-the-CLI)
* [ Test public repositories before use](/hc/en-us/articles/360003851277-Test-public-repositories-before-use)
* [ Test public npm packages before use](/hc/en-us/articles/360003812498-Test-public-npm-packages-before-use)
* [ Why is my setup.py file failing to scan or finding 0 dependencies?](/hc/en-us/articles/360005351598-Why-is-my-setup-py-file-failing-to-scan-or-finding-0-dependencies-)
* [ Scan all unmanaged jar files](/hc/en-us/articles/4402912058769-Scan-all-unmanaged-jar-files)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [Snyk CLI](/hc/en-us/categories/360000456217-Snyk-CLI)
3.  [Test for vulnerabilities](/hc/en-us/sections/360001100837-Test-for-vulnerabilities)

##  Test public npm packages before use

You can also use `snyk test` to **scrutinize a public package before installing it**, to see if it has known vulnerabilities or not. Using the package name will test the latest version of that package, and you can also provide a specific version or range using `snyk test module[@semver-range]`.

```text
snyk test ionic@1.6.5
```

See also

* [Getting started with the CLI](/hc/articles/360003812458#UUID-19fc37f2-b686-11ed-b85c-4789e90c8dfc)
* [Our full CLI reference](/hc/articles/360003812578#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0)
* [Set severity thresholds for CLI tests](/hc/articles/360003851337#UUID-844be978-191f-c813-ccd0-2221c04ca510)
* [Test your code from the CLI](/hc/articles/360003812478#UUID-2e8464f9-1d8a-5f79-466d-2ca5c5cf9f22)
* [Test dev dependencies](/hc/articles/360003812478#UUID-1070ae3e-5f30-cb4e-6350-a1c3f5c67511)
* [Test public repositories before use](/hc/articles/360003851277#UUID-ba99a73f-110d-1f1d-9e7a-1bad66bf0996)
* 
