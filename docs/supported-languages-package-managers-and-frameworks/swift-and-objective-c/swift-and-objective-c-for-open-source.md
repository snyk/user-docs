# Swift and Objective-C for open source

## Swift and Objective-C for open source support

**Package manager**: CocoaPods, Swift Package Manager

**Package manager versions**: CocoaPods, Swift Package Manager, Swift v3.0 or higher.

**Package registry**: No single registry, multiple sources including [cocoapods.org](https://cocoapods.org)

**Import your app through SCM**: Available for CocoaPods

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:swift`, `pkg:cocoapods`

**Test your app's packages**: Available, `pkg:swift`, `pkg:cocoapods`

**Features**:&#x20;

* License scanning (CocoaPods)
* Reports

## Open source and licensing

Snyk Open Source supports the following package managers.

<table><thead><tr><th>Package managers / Features</th><th width="151">CLI support</th><th>Git support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>Cocoapods</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td></td></tr><tr><td>Swift Package Manager</td><td>✔︎</td><td></td><td></td><td></td></tr></tbody></table>

The requirements follow.

| Swift Package Manager                                                                                                                                                                                                                                                                                                                                                                                                                                                              | CocoaPods and Snyk CLI                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <p>A <code>Package.swift</code> file must be present for the Snyk CLI to discover the Project.<br><br>Snyk uses the <code>swift package show-dependencies</code>  command to build the dependency graph.<br><br>Limitations:<br>Supports only Projects using Swift 3.0 or higher.<br><br>Swift Package Manager supports pre-processing and post-processing. For post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.<br></p> | <p>To build the dependency graph, Snyk examines the <code>Podfile</code> and <code>Podfile.lock</code> files.<br><br>When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync by using the <code>--strict-out-of-sync=true|false</code> option. </p><p>For details, see <a href="../../snyk-cli/commands/test.md#option-for-cocoapods-projects">Option for CocoaPods projects</a> in the <code>snyk test</code> help.</p> |

The following summarizes support for Git import and testing.

| Swift Package Manager and Git                                               | CocoaPods and Git                                                            |
| --------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| It is not possible to scan Swift Package Manager Projects using Git import. | To test your Projects, Snyk analyzes the `Podfile` and `Podfile.lock` files. |

