# Why do yarn and npm report more dependencies than snyk?

##  Why do yarn and npm report more dependencies than snyk?

You'll often find that the number of dependencies that Snyk reports will vary from other tools such as NPM or Yarn.

Tools like NPM and Yarn tend to count the number of dependencies in their tree, but not necessarily doing a distinct count - there tend to be many duplicates as packages tend to be called on several or even many times within the same project.

Snyk also by default does not scan dependencies in the devDependencies section of the package.json.

### Enabling devDependencies in your scans

In the CLI you can scan devDependencies by appending --dev at the end of your command, for example:

```text
$ snyk test --dev
```

In the Snyk.io portal this can be configured in settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Languages**

