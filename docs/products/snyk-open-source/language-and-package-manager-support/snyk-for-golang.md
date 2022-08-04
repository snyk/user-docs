# Snyk for Go

Snyk supports testing and monitoring of Go projects that have their dependencies managed by [Go Modules](https://golang.org/ref/mod), [dep](https://github.com/golang/dep) and [govendor](https://github.com/kardianos/govendor).

The following describes how to use Snyk to scan your Go projects:

### Features <a href="#h_01esm3gfnmn0f7art59aek97tm" id="h_01esm3gfnmn0f7art59aek97tm"></a>

{% hint style="info" %}
Some of these features may not be available for your Snyk subscription plan.
{% endhint %}

| Package managers / Features                       | CLI support | Git support | License scanning | Fix PRs |
| ------------------------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [Go Modules](https://golang.org/ref/mod)          | ✔︎          | ✔︎          | ✔︎               |         |
| [dep](https://github.com/golang/dep)              | ✔︎          | ✔︎          | ✔︎               |         |
| [govendor](https://github.com/kardianos/govendor) | ✔︎          | ✔︎          | ✔︎               |         |

## **How it works**

Once Snyk builds the tree, the Snyk [vulnerability database](https://snyk.io/vuln) can be used to find vulnerabilities in any of the modules or packages, anywhere in the dependency tree.

{% hint style="info" %}
To scan your dependencies in the CLI, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

How Snyk analyzes and builds the tree varies depending on the language and package manager for the project, as well as on the location of your project.

## Snyk CLI tool for Go projects

**Go Modules**

To build the dependency tree, Snyk uses the `go list -json -deps` command.

{% hint style="info" %}
Snyk scans Go Modules projects in the CLI at the _package_ level rather than on the _module_ level, as we have full access to your project source code.\
This is beneficial since you might use a vulnerable module but not the vulnerable package.
{% endhint %}

When testing Go Modules projects via the CLI, Snyk does not require that dependencies be installed, but you must have a `go.mod` file at the root of your project: `go list` uses this and your project source code to build a complete dependency tree.

{% hint style="info" %}
Different versions of the Go generate different results for the **go list -json -deps** command. This can affect the dependency tree and the vulnerabilities that the Snyk CLI finds.
{% endhint %}

**Dep**

To build the dependency tree, Snyk analyzes your `Gopkg.lock` files.

When testing dep projects via the CLI Snyk requires dependencies to be installed, run `dep ensure` to achieve this.

**Govendor**

To build the dependency tree, Snyk analyzes your `vendor/vendor.json` files.

When testing Govendor projects via the CLI, Snyk requires that dependencies be installed:  run `govendor sync` to achieve this.

## Git services for Go projects

### **Go Modules**

{% hint style="info" %}
For Go Modules projects imported via Git, dependencies are resolved at the _module_ level rather than at the _package_ level, because Snyk does not have full access to your project source code.\
This means that you may see more issues (including potential false positives) reported than for projects tested in the CLI, because Snyk reports all the vulnerabilities for each module, and not just the package(s) referenced in your source code.
{% endhint %}

To build the dependency tree Snyk, runs the `go mod graph` command using the `go.mod` files in the selected repository.

**Private modules**

Go Modules projects that depend on modules from private Git repositories are supported when the private repositories are in the same Git organization as the main project repository.&#x20;

{% hint style="info" %}
Snyk supports a single private Git repository for Go Modules projects.
{% endhint %}

Imports for projects with private modules from repos in other Git organizations will fail. Support for private module dependencies from other Git organizations is planned for the future.

Private modules are supported for GitHub and Bitbucket Cloud. GitLab, GitHub Enterprise and Bitbucket Server are not currently supported.

**Broker**

Projects imported via the new [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker/broker-introduction) clients should work as expected.

To add support to existing clients created before Dec 30th 2020, you should add `go.mod` and `go.sum` to your `accept.json` file, as per the changes in this [pull request](https://github.com/snyk/broker/pull/299/files).

If you're using private Go Modules (repositories) integrated via the broker, note that Snyk requires that each private module have a `go.mod` file defined.

**Dep**

To build the dependency tree, Snyk analyzes the `Gopkg.lock` files in the selected repository.

**Govendor**

To build the dependency tree, Snyk analyzes the `vendor/vendor.json` files in the selected repository.
