# Unable to test Go Modules repository using CLI

##  Unable to test Go Modules repository using CLI

When attempting to test your Go Modules repository you may encounter the following error:

```text
Testing /Users/SnykUser/Documents/Snyk/reponame...Could not detect supported target files in /Users/SnykUser/Documents/Snyk/reponame.
Please see our documentation for supported languages and target files: https://support.snyk.io/hc/en-us/articles/360000911957-Language-support and make sure you are in the right directory.
```

This is generally because you do not have a `Go.mod` file present in this repo.

You can read more about our [Go package manager support](/hc/en-us/articles/360001378818) page.

### Creating a Go.mod <a id="bb30"></a>

If your project is already in version control, you can simply run

```text
go mod init
```

Or you can supply module path manually. It’s kinda like a name, URL and import path for your package:

```text
go mod init github.com/your/repo
```

This command will create _go.mod_ file which both defines projects requirements and locks dependencies to their correct versions

