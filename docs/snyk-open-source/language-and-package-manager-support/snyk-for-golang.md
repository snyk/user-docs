# Snyk for Golang

Snyk supports testing and monitoring of Go projects that have their dependencies managed by [Go Modules](https://golang.org/ref/mod), [dep](https://github.com/golang/dep) and [govendor](https://github.com/kardianos/govendor).

The following describes how to use Snyk to scan your Go projects:

#### Features <a id="h_01ESM3GFNMN0F7ART59AEK97TM"></a>

{% hint style="info" %}

---
**NOTE**  
Features might not be available, depending on your subscription plan.

---
{% endhint %}

| Package managers / Features  | CLI support  | Git support  | License scanning  | Remediation  | Runtime monitoring  |
| :--- | :--- | :--- | :--- | :--- | :--- |
| [Go Modules](https://golang.org/ref/mod) | ✔︎ | ✔︎  | ✔︎ |  |  |
| [dep](https://github.com/golang/dep) | ✔︎ | ✔︎ | ✔︎ |  |  |
| [govendor](https://github.com/kardianos/govendor) | ✔︎ | ✔︎ | ✔︎ |  |  |

### **How it works**

Once we've built the tree, we can use our [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the modules or packages anywhere in the dependency tree.

{% hint style="info" %}

---
**NOTE**  
In order to scan your dependencies in the CLI, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files.

---

{% endhint %}

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project.

## Snyk CLI tool for Go projects

**Go Modules**

In order to build the dependency tree Snyk uses the `go list -json -deps` command.

### Note

{% hint style="info" %}
Note:  
Snyk scans Go Modules projects in the CLI at the _package_ level rather than on the _module_ level, as we have full access to your project source code.  
This is beneficial since you might use a vulnerable module but not the vulnerable package.
{% endhint %}

When testing Go Modules projects via the CLI Snyk does not require dependencies to be installed, but you must have a `go.mod` file at the root of your project, `go list` uses this and your project source code to build a complete dependency tree.

{% hint style="info" %}
Different versions of the Go \(i.e. 1.15 vs 1.16\) will generate different results for the \`go list -json -deps\` command. This will likely affect the dependency tree and in turn, the vulnerabilities that the Snyk CLI will find.
{% endhint %}

**Dep**

In order to build the dependency tree Snyk analyzes your `Gopkg.lock` files.

When testing dep projects via the CLI Snyk requires dependencies to be installed, run `dep ensure` to achieve this.

**Govendor**

In order to build the dependency tree Snyk analyzes your `vendor/vendor.json` files.

When testing Govendor projects via the CLI Snyk requires dependencies to be installed, run `govendor sync` to achieve this.

#### Git services for Go projects <a id="h_01EFH2KBK8MDKV4M7YW8CTE25Z"></a>

**Go Modules**

### Note

{% hint style="info" %}
For Go Modules projects imported via Git, dependencies are resolved at the _module_ level rather than the _package_ level, as we do not have full access to your project source code.   
This means you may see more issues reported than for projects tested in the CLI, as we report all vulnerabilities for each module not just the package\(s\) referenced in your source code.
{% endhint %}

In order to build the dependency tree Snyk runs the `go mod graph` command using the `go.mod` files in the selected repository.

**Private modules**

Go Modules projects that depend on modules from private Git repositories are supported where the private repositories are in the same Git organisation as the main project repository. 

Imports for projects with private modules from repos in other Git organisations will fail. Support for private module dependencies from other Git organisations is planned for the future.

Private modules are supported for GitHub and Bitbucket Cloud. GitLab and Bitbucket Server are not currently supported.

**Broker**

Projects imported via the new [Snyk Broker](https://support.snyk.io/hc/en-us/articles/360015367178) clients should work as expected. 

To add support to existing clients created before Dec 30th 2020, you should add `go.mod` and `go.sum` to your `accept.json` file, as per the changes in this [pull request](https://github.com/snyk/broker/pull/299/files).

If you are using private Go Modules \(repositories\) integrated via the broker, please note that we require each private module to have a `go.mod` file defined.

**Dep**

In order to build the dependency tree Snyk analyzes the `Gopkg.lock` files in the selected repository.

**Govendor**

In order to build the dependency tree Snyk analyzes the `vendor/vendor.json` files in the selected repository.

