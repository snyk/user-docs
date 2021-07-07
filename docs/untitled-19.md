# Mismatch of Vulnerabilities. Why does the CLI show a different number of vulnerabilities than throug

##  Mismatch of Vulnerabilities. Why does the CLI show a different number of vulnerabilities than through the Snyk App?

There are a few reasons why the Command Line Interface \(CLI\) tool may indicate a different number of Vulnerabilities found than you will through the Snyk App or via an Integration with your Code repo.

* Your project may have devDependencies that are being ignored
* The CLI has direct access to your Private dependencies and build environment
* When comparing to something like **npm audit**, Snyk reports different results
* If the project was scanned with a Lockfile, the results may be different than when scanned using a manifest file such as a package.json

**Your Project may have devDependencies that are being ignored**

By default, Snyk will not scan for dependencies listed in the manifest file's devDependencies section. You can enable scanning these files a couple of ways:

1. Through the Snyk CLI, simply add the command line option: **--dev**
2. Via the Snyk App you can enable by click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Languages**: 

**The CLI has direct access to your Private dependencies and build environment**

The Snyk CLI has direct access to all of your Projects files, build environment and private dependencies.

Integrations such as GitHub, GitLab and BitBucket we only process the project's dependency files \(e.g. `package.json`, `Gemfile.lock` etc.\). We then infer the dependency graph of your project, mimicking the operation of the build tool. We do not have access to private dependencies, specifics of your build environment and so on. Hence the results may be partial compared to our CLI based analysis.

**When comparing to something like npm audit, Snyk reports different results**

When comparing Snyk to npm audit, we have about one third more vulnerabilities in our database. For a third party perspective, checkout Nearform's blog where they compared Snyk and Npm audit here: [https://www.nearform.com/blog/comparing-npm-audit-with-snyk/](https://www.nearform.com/blog/comparing-npm-audit-with-snyk/)

