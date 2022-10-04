# Snyk for Swift and Objective-C (CocoaPods)

Snyk offers security scanning to test your projects for vulnerabilities, both through the [Snyk CLI ](../../../run-snyk/snyk-cli/)and from the Snyk Web UI through different [Snyk Integrations](../../../integrate-with-snyk/).

The following describes how to use Snyk to scan your CocoaPods projects:

### Features

{% hint style="info" %}
Features might not be available, depending on your subscription plan.
{% endhint %}

| Package managers / Features | CLI support | Git support | License scanning | Fix PRs |
| --------------------------- | ----------- | ----------- | ---------------- | ------- |
| Cocoapods                   | ✔︎          | ✔︎          | ✔︎               |         |

### **How it works**

Once we’ve built the tree, we can use our vulnerability database to find vulnerabilities in any of the packages anywhere in the dependency tree.

{% hint style="info" %}
**Note**\
In order to scan your dependencies, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project

### Snyk CLI tool for CocoaPods projects

We scan CocoaPods projects and examine your Podfile and Podfile.lock files. We then compare the specific versions of every direct and deep dependency in your project against our vulnerability database in order to build the project dependency tree accordingly.

### **CLI parameters for Swift and Objective-C**

#### **Prerequisites**

* Ensure you've installed the relevant package manager before you begin using the Snyk CLI tool.
* Ensure you've included the relevant manifest files supported by Snyk before testing.
* Install and authenticate the Snyk CLI to start analyzing projects from your local environment. Read more about Snyk CLI in [Getting started with the CLI](../../../run-snyk/snyk-cli/getting-started-with-the-cli.md) as well.

#### **Parameters**

When working with Swift and Objective-C projects from our CLI, you can prevent testing any lockfiles that are out-of-sync, as follows:

| Option                  | Description                                                                            |
| ----------------------- | -------------------------------------------------------------------------------------- |
| `--strict-out-of-sync=` | <p>Prevent testing out-of-sync lockfiles.</p><p>Defaults to <strong>true</strong>.</p> |

### Git services for CocoaPods projects

We scan CocoaPods projects and examine your Podfile and Podfile.lock files. We then compare the specific versions of every direct and deep dependency in your project against our vulnerability database in order to build the project dependency tree accordingly.

**Git services**

Swift and Objective-C projects managed by CocoaPods can be imported from any of the Git repositories we support. In order to test your projects, we analyze your Podfile and Podfile.lock files.
