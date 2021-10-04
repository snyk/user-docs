# Test public repositories before use

To test a public Github, BitBucket or GitLab repository, run `snyk test` and include the Github URL to the repo.

```text
$ snyk test https://github.com/snyk/snyk
```

The following git URL formats are supported:

* `git://github.com/user/project.git#commit-ish`
* `https://github.com/user/project#commit-ish`
* `user/project#commit-ish`

This also works for Bitbucket and GitLab.

You can also test a public npm package or Github project via the Test page on [snyk.io](https://snyk.io/test/)

See also

* [Getting started with the CLI](https://support.snyk.io/hc/articles/360003812458#UUID-19fc37f2-b686-11ed-b85c-4789e90c8dfc)
* [Our full CLI reference](https://support.snyk.io/hc/articles/360003812578#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0)
* [Set severity thresholds for CLI tests](https://support.snyk.io/hc/articles/360003851337#UUID-844be978-191f-c813-ccd0-2221c04ca510)
* [Test your code from the CLI](https://support.snyk.io/hc/articles/360003812478#UUID-2e8464f9-1d8a-5f79-466d-2ca5c5cf9f22)
* [Test dev dependencies](https://support.snyk.io/hc/articles/360003812478#UUID-1070ae3e-5f30-cb4e-6350-a1c3f5c67511)
* [Test public npm packages before use](https://support.snyk.io/hc/articles/360003812498#UUID-0ab434a8-c1b5-873d-cbf6-7a61a07c4ad8)

