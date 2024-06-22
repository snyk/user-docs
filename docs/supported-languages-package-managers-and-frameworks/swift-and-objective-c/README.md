# Swift and Objective-C

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## Snyk CLI&#x20;

#### **Prerequisites for CLI for Swift and CocoaPods**

1. [Create a Snyk account](../../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
2. [Install Snyk CLI and authenticate your machine](../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
3. [Set the default Organization for all Snyk tests](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Explore test results in a JSON or SARIF format in the terminal ](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Exporting the test results to a JSON or SARIF file](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)

#### Open source and licensing

The way Snyk analyzes and builds the dependency tree varies depending on the language and package manager of the Project.

After Snyk has built the tree, Snyk uses the vulnerability database to find vulnerabilities in any packages in the dependency tree.

| Swift Package Manager                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | CocoaPods and Snyk CLI                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <p>A <code>Package.swift</code> file must be present for the Snyk CLI to discover the Project.<br><br>Snyk uses the <code>swift package show-dependencies</code>  command to build the dependency graph.<br><br>Limitations:<br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> Supports only Projects using Swift 3.0 or higher.<br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> Swift Package Manager supports pre-processing and post-processing. For post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.<br></p> | <p>To build the dependency graph, Snyk examines the <code>Podfile</code> and <code>Podfile.lock</code> files.<br><br>When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync by using the <code>--strict-out-of-sync=true|false</code> option. </p><p>For details, see <a href="https://docs.snyk.io/snyk-cli/commands/test#option-for-cocoapods-projects">Option for CocoaPods projects</a> in the <code>snyk test</code> help.</p> |

## Snyk Web UI (Git repository integration)

| Swift Package Manager and Git                                                              | CocoaPods and Git                                                            |
| ------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| At the moment, it is not possible to scan Swift Package Manager Projects using Git import. | To test your Projects, Snyk analyzes the `Podfile` and `Podfile.lock` files. |

## IDE and CI/CD

:link: For integrated development environments, see [Use Snyk in your IDE](../../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/).

:link: If you use continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software. See [Snyk CI/CD](../../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
