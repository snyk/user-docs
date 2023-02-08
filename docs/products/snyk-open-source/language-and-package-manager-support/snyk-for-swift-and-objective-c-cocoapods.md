# Snyk for Swift and Objective-C

Snyk offers security scanning to test your Projects for vulnerabilities, both from the [Snyk CLI](../../../snyk-cli/) and the [Snyk Web UI](../../../getting-started/quickstart/create-a-snyk-account/snyk-web-ui.md) using different [Snyk Integrations](../../../integrations/).

## Features of Snyk for Swift and Objective-C

{% hint style="warning" %}
Swift Package Manager CLI support  is in beta
{% endhint %}

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features | CLI support | Git support | License scanning | Fix PRs |
| --------------------------- | ----------- | ----------- | ---------------- | ------- |
| Cocoapods                   | ✔︎          | ✔︎          | ✔︎               |         |
| Swift Package Manager       | ✔︎ Beta     |             |                  |         |

## Snyk CLI for Swift and Objective-C projects (CI/CD)

The way Snyk analyzes and builds the dependencies varies depending on the language and package manager of the Project.

* Snyk CLI with Swift Package Manager: to build the dependency graph, Snyk uses the `swift package show-dependencies` CLI command. A `Package.swift` file must be present for the Snyk CLI to discover the project.
* Snyk CLI with CocoaPods: to build the dependency graph, Snyk examines the `Podfile` and `Podfile.lock` files.

Once Snyk has built the tree, Snyk uses the vulnerability database to find vulnerabilities in any of the packages anywhere in the dependency tree.

### **Prerequisites for CLI for Swift and CocoaPods**

* Ensure you have installed the relevant package manager before you begin using the Snyk CLI.
* Ensure you have included the relevant manifest files supported by Snyk before testing.
* [Install](../../../snyk-cli/install-the-snyk-cli.md) and authenticate the Snyk CLI to start analyzing projects from your local environment. For more information about Snyk CLI see [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-cli.md).

### CLI help for Swift Package Manager projects

{% hint style="warning" %}
Swift Package Manager CLI support is in beta
{% endhint %}

There are a some limitations to using the Snyk CLI with Swift Package Manager projects.

* The `snyk monitor` CLI command is not currently supported.
* Projects must use Swift 3.0 or higher.&#x20;
* Swift Package Manager supports pre- or post-processing. In the case of post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.

### **CLI help for CocoaPods Projects**

When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync using the following option: `--strict-out-of-sync=true|false`

For details see [Option for CocoaPods projects](https://docs.snyk.io/snyk-cli/commands/test#option-for-cocoapods-projects) in the `snyk test` help.

## Git services for Swift and Objective-C projects

Testing Swift Package Manager projects via Git import is not currently supported.

Snyk supports testing CocoaPods projects imported from Git repositories.

Projects managed by CocoaPods can be imported from any of the Git repositories that Snyk  supports.&#x20;

In order to test your projects, Snyk analyzes the `Podfile` and `Podfile.lock` files.
