# Snyk for Swift and Objective-C \(CocoaPods\)

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your CocoaPods projects:

## Features

{% hint style="info" %}
Features might not be available, depending on your subscription plan.
{% endhint %}

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| ![i\_icon\_cocoapods.png](../../.gitbook/assets/uuid-6de05da9-de7e-11cc-4316-8459517aaf57-en.png) | Cocoapods | ✔︎ | ✔︎ | ✔︎ |  |  |

## **How it works**

Once we’ve built the tree, we can use our [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

{% hint style="info" %}
**Note**  
In order to scan your dependencies, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project:

* [analysis with our CLI tool](https://support.snyk.io/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-#UUID-2f6bc912-9512-0d39-56f9-01b0926ef65d)
* [import from our app](https://support.snyk.io/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-#UUID-3276e078-ac51-90e8-3204-3cbc39b81034)

## Snyk CLI tool for CocoaPods projects

We scan CocoaPods projects and examine your Podfile and Podfile.lock files. We then compare the specific versions of every direct and deep dependency in your project against our vulnerability database in order to build the project dependency tree accordingly.

### **CLI parameters for Swift and Objective-C**

#### **Prerequisites**

* Ensure you've installed the relevant package manager before you begin using the Snyk CLI tool.
* Ensure you've included the relevant manifest files supported by Snyk before testing.
* Install and authenticate the Snyk CLI to start analyzing projects from your local environment. Read more about our CLI in [Getting started with the CLI](https://support.snyk.io/hc/articles/360003812458#UUID-6d3e2b39-daa0-f2f1-19d2-b9107b678c81) as well.

#### **Parameters**

When working with Swift and Objective-C projects from our CLI, you can prevent testing any lockfiles that are out-of-sync, as follows:

<table>
  <thead>
    <tr>
      <th style="text-align:left">Option</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><code>--strict-out-of-sync=</code>
      </td>
      <td style="text-align:left">
        <p>Prevent testing out-of-sync lockfiles.</p>
        <p>Defaults to <b>true</b>.</p>
      </td>
    </tr>
  </tbody>
</table>

## Git services for CocoaPods projects

We scan CocoaPods projects and examine your Podfile and Podfile.lock files. We then compare the specific versions of every direct and deep dependency in your project against our vulnerability database in order to build the project dependency tree accordingly.

**Git services**

Swift and Objective-C projects managed by CocoaPods can be imported from any of the Git repositories we support. In order to test your projects, we analyze your Podfile and Podfile.lock files.

