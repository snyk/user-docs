# Test public repositories before use

To test a public Github, BitBucket or GitLab repository, run `snyk test` and include the GitHub URL to the repo:

`$ snyk test https://github.com/snyk/snyk`

The following git URL formats are supported:

* `git://github.com/user/project.git#commit-ish`
* `https://github.com/user/project#commit-ish`
* `user/project#commit-ish`

This also works for Bitbucket and GitLab.

You can also test a public npm package or Github project via the Test page on [snyk.io](https://snyk.io/test/).

See also

* [Getting started with the CLI](../getting-started-with-the-cli.md)
* [CLI reference](../cli-reference.md)
* [Use Snyk Open Source from the CLI](../../products/snyk-open-source/use-snyk-open-source-from-the-cli/)
* [Set severity thresholds for CLI tests](set-severity-thresholds-for-cli-tests.md)
* [Test public npm packages before use](test-public-npm-packages-before-use.md)
