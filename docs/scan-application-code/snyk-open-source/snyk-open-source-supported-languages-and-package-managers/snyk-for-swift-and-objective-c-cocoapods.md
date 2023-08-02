# Snyk for Swift and Objective-C

{% hint style="success" %}
Snyk offers security scanning to test your Projects for vulnerabilities, both from the [Snyk CLI](../../../snyk-cli/) and the [Snyk Web UI](../../../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md), using different [Snyk Integrations](../../../integrations/).
{% endhint %}

## Features of Snyk for Swift and Objective-C

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features | CLI support | Git support | License scanning | Fix PRs |
| --------------------------- | ----------- | ----------- | ---------------- | ------- |
| Cocoapods                   | ✔︎          | ✔︎          | ✔︎               |         |
| Swift Package Manager       | ✔︎          |             |                  |         |

## Snyk CLI for Swift and Objective-C Projects (CI/CD)

The way Snyk analyzes and builds the dependency tree varies depending on the language and package manager of the Project.

After Snyk has built the tree, Snyk uses the vulnerability database to find vulnerabilities in any packages in the dependency tree.

### **Prerequisites for CLI for Swift and CocoaPods**

* Ensure you have installed the relevant package manager before using the Snyk CLI.
* Ensure you have included the relevant manifest files supported by Snyk before testing.
* [Install](../../../snyk-cli/install-the-snyk-cli/) and authenticate the Snyk CLI to analyze projects from your local environment. For more information about Snyk CLI see [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-cli.md).

### Swift Package Manager and Snyk CLI

A `Package.swift` file must be present for the Snyk CLI to discover the Project.

To build the dependency graph, Snyk uses the `swift package show-dependencies`  command.&#x20;

There are some limitations to using the Snyk CLI with Swift Package Manager Projects.

* Only Projects using Swift 3.0 or higher are supported.
* Swift Package Manager supports pre-processing and post-processing. In the case of post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.

### **CocoaPods** and Snyk CLI

To build the dependency graph, Snyk examines the `Podfile` and `Podfile.lock` files.

When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync using the following option: `--strict-out-of-sync=true|false`

For details, see [Option for CocoaPods projects](https://docs.snyk.io/snyk-cli/commands/test#option-for-cocoapods-projects) in the `snyk test` help.

## Git services for Swift and Objective-C

### Swift Package Manager and Git

Scanning Swift Package Manager Projects via Git import is not currently supported.

### CocoaPods and Git

To test your Projects, Snyk analyzes the `Podfile` and `Podfile.lock` files.

Projects managed by CocoaPods can be imported from any of the Git repositories that Snyk supports.
