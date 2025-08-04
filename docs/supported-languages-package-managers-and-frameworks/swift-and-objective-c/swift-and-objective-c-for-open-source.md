# Swift and Objective-C for open source

## Swift and Objective-C for open source support

Refer to the[ Swift and Objective-C detail](./) for supported package managers and features.

If you need help, [contact Snyk Support](https://support.snyk.io).

## Open source and licensing

Snyk Open Source features support the following package managers.

<table><thead><tr><th>Package managers / Features</th><th width="151">CLI support</th><th>Git support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>Cocoapods</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td></td></tr><tr><td>Swift Package Manager</td><td>✔︎</td><td></td><td></td><td></td></tr></tbody></table>

The requirements follow.

| Swift Package Manager                                                                                                                                                                                                                                                                                                                                                                                                                                                             | CocoaPods and Snyk CLI                                                                                                                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <p>A <code>Package.swift</code> file must be present for the Snyk CLI to discover the Project.<br><br>Snyk uses the <code>swift package show-dependencies</code> command to build the dependency graph.<br><br>Limitations:<br>Supports only Projects using Swift 3.0 or higher.<br><br>Swift Package Manager supports pre-processing and post-processing. For post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.<br></p> | <p>To build the dependency graph, Snyk examines the <code>Podfile</code> and <code>Podfile.lock</code> files.<br><br>When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync by using the <code>--strict-out-of-sync=true</code></p> |

The following summarizes support for Git import and testing.

| Swift Package Manager and Git                                               | CocoaPods and Git                                                            |
| --------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| It is not possible to scan Swift Package Manager Projects using Git import. | To test your Projects, Snyk analyzes the `Podfile` and `Podfile.lock` files. |
