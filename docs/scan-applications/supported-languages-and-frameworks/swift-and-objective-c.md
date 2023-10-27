# Swift and Objective-C

## Supported frameworks and package managers

{% hint style="warning" %}
You might encounter false positives or false negatives for partially supported frameworks and package managers.
{% endhint %}

### Code analysis for Swift

Snyk Code for Swift supports the following frameworks:

* Alamofire
* Pathos&#x20;
* sqlite3
* CryptoKit

### Open source and licensing

Snyk Open Source supports the following frameworks:

<table><thead><tr><th>Package managers / Features</th><th width="151">CLI support</th><th>Git support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>Cocoapods</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td></td></tr><tr><td>Swift Package Manager</td><td>✔︎</td><td></td><td></td><td></td></tr></tbody></table>

## Getting started with Snyk for Swift and Objective-C language across environments

### Snyk CLI&#x20;

#### **Prerequisites for CLI for Swift and CocoaPods**

1. [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account.md)
2. [Install Snyk CLI and authenticate your machine](../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
3. [Set the default Organization for all Snyk tests](broken-reference) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../scan-application-code/snyk-code/cli-for-snyk-code/excluding-directories-and-files-from-the-snyk-code-cli-test.md)
* [Explore test results in a JSON or SARIF format in the terminal ](broken-reference)
* [Exporting the test results to a JSON or SARIF file](broken-reference)

#### Open source and licensing

The way Snyk analyzes and builds the dependency tree varies depending on the language and package manager of the Project.

After Snyk has built the tree, Snyk uses the vulnerability database to find vulnerabilities in any packages in the dependency tree.

| Swift Package Manager                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | CocoaPods and Snyk CLI                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <p>A <code>Package.swift</code> file must be present for the Snyk CLI to discover the Project.<br><br>Snyk uses the <code>swift package show-dependencies</code>  command to build the dependency graph.<br><br>Limitations:<br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ</span> Supports only Projects using Swift 3.0 or higher.<br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ</span> Swift Package Manager supports pre-processing and post-processing. For post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.<br></p> | <p>To build the dependency graph, Snyk examines the <code>Podfile</code> and <code>Podfile.lock</code> files.<br><br>When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync by using the <code>--strict-out-of-sync=true|false</code> option. </p><p>For details, see <a href="https://docs.snyk.io/snyk-cli/commands/test#option-for-cocoapods-projects">Option for CocoaPods projects</a> in the <code>snyk test</code> help.</p> |

### Snyk Web UI (Git repository integration)

| Swift Package Manager and Git                                                              | CocoaPods and Git                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| At the moment, it is not possible to scan Swift Package Manager Projects using Git import. | <p>To test your Projects, Snyk analyzes the <code>Podfile</code> and <code>Podfile.lock</code> files.<br><br>You can import projects managed by CocoaPods from any of the Snyk supported Git repositories.</p> |

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../../integrations/ide-tools/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
