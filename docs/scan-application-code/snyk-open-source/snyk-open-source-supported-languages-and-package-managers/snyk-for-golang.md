# Snyk for Go

Snyk supports testing and monitoring of Go projects that have dependencies managed by [Go Modules](https://golang.org/ref/mod), [dep,](https://github.com/golang/dep) and [govendor](https://github.com/kardianos/govendor).

{% hint style="warning" %}
After January 1, 2023, Snyk will no longer support govendor Projects. As a general security best practice, Snyk recommends using tools that are consistently maintained and up-to-date.

Because govendor was [archived in GitHub](https://github.com/kardianos/govendor), it has not been receiving bug fixes and security updates, nor has it been improved and maintained.

Once Snyk no longer supports scanning of govendor Projects, a warning will be issued and no results will be received.

Snyk recommends using [Go Modules](https://go.dev/ref/mod), which is considered production beginning with go1.14 but which also works well in go1.13 and go1.12.
{% endhint %}

The following describes how to use Snyk to scan your Go Projects:

## Features of Snyk for Go <a href="#h_01esm3gfnmn0f7art59aek97tm" id="h_01esm3gfnmn0f7art59aek97tm"></a>

{% hint style="info" %}
Some of these features may not be available for your Snyk plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features                       | CLI support | Git support | License scanning | Fix PRs |
| ------------------------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [Go Modules](https://golang.org/ref/mod)          | ✔︎          | ✔︎          | ✔︎               |         |
| [dep](https://github.com/golang/dep)              | ✔︎          | ✔︎          | ✔︎               |         |
| [govendor](https://github.com/kardianos/govendor) | ✔︎          | ✔︎          | ✔︎               |         |

## **How Snyk for Go works**

Once Snyk builds the tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the modules or packages, anywhere in the dependency tree.

{% hint style="info" %}
To scan your dependencies in the CLI, ensure you have first installed the relevant package manager and that your project contains the supported manifest files.
{% endhint %}

How Snyk analyzes and builds the tree varies depending on the language and package manager for the Project, as well as on the location of your Project.

## Snyk CLI for Go projects

### **Go Modules and Snyk CLI**

{% hint style="info" %}
Snyk scans Go Modules projects in the CLI at the _package_ level rather than on the _module_ level, as Snyk has full access to your Project source code.\
This is beneficial since you might use a vulnerable module but not the vulnerable package.
{% endhint %}

To build the dependency tree, Snyk uses the `go list -json -deps` command.

Packages from the Go standard library are excluded from the dependency tree.

When you test Go Modules projects using the CLI, Snyk does not require that dependencies be installed, but you must have a `go.mod` file at the root of your Project; `go list` uses this and your project source code to build a complete dependency tree.

{% hint style="info" %}
Different versions of Go generate different results for the **go list -json -deps** command. This can affect the dependency tree and the vulnerabilities that the Snyk CLI finds.
{% endhint %}

### **Dep and Snyk CLI**

To build the dependency tree, Snyk analyzes your `Gopkg.lock` files.

When you test dep Projects using the CLI, Snyk requires dependencies to be installed. Run `dep ensure` to achieve this.

### **Govendor and Snyk CLI**

To build the dependency tree, Snyk analyzes your `vendor/vendor.json` files.

When you test Govendor projects using the CLI, Snyk requires that dependencies be installed. Run `govendor sync` to achieve this.

## Git services for Go projects

### **Go Modules and Git**

{% hint style="info" %}
For Go Modules projects imported via Git, dependencies are resolved at the _module_ level rather than at the _package_ level, because Snyk does not have full access to your project source code.

This means that you may see more issues, including potential false positives, reported than for projects tested in the CLI, because Snyk reports all the vulnerabilities for each module, and not just the package(s) referenced in your source code.
{% endhint %}

To build the dependency tree Snyk runs the `go mod graph` command using the `go.mod` files in the selected repository.

### **Private modules and Git**

Go Modules Projects that depend on modules from private Git repositories are supported when the private repositories are in the same Git organization as the main project repository.

{% hint style="info" %}
Snyk supports a single private Git repository for _all_ the Go Modules projects you have.
{% endhint %}

Imports for Projects with private modules from repos in other Git organizations will fail. Support for private module dependencies from other Git organizations is planned for the future.

Private modules are supported for GitHub, GitHub Enterprise, and Bitbucket Cloud. GitLab and Bitbucket Server are not currently supported.

### **Broker and Git with Go Projects**

Projects imported via the new [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker/broker-introduction) clients should work as expected.

To add support to existing clients created before December 30, 2020, add `go.mod` and `go.sum` to your `accept.json` file, as per the changes in this [pull request](https://github.com/snyk/broker/pull/299/files).

If you're using private Go Modules (repositories) integrated via the Broker, note that Snyk requires that each private module have a `go.mod` file defined.

### **Dep and Git**

To build the dependency tree, Snyk analyzes the `Gopkg.lock` files in the selected repository.

### **Govendor and Git**

To build the dependency tree, Snyk analyzes the `vendor/vendor.json` files in the selected repository.
