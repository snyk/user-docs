# Swift and Objective-C

## Applicability and integration

{% hint style="info" %}
Swift is supported for Snyk Code and Snyk Open Source.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

Available functions:

* Test your app's SBOM using `pkg:swift`, `pkg:cocoapods`
* Test your app's packages using `pkg:swift`, `pkg:cocoapods`

## Technical specifications

### Supported frameworks and libraries

For Swift and Objective-C, the following frameworks and libraries are supported:

* Swift standard library
* Foundation
* AppKit
* Swift UI
* UI Kit
* Asynchttpclient
* Commoncrypt
* Commoncrypto
* Cryptokit
* Cryptoswift
* Cryptor
* AlamoFire
* Filekit
* google-gemini/generative-ai-swift
* MacPaw/OpenAI
* dylanshine/openai-kit
* Pathos
* SQLite3
* Webkit
* SwiftCLI
* ShellOut
* SwiftShell
* Subprocess
* Shout
* Swiftline
* RNCryptor

### Supported package managers and registries

For Swift and Objective-C, Snyk supports the following package managers: CocoaPods, Swift Package Manager v3.0 or higher.

As package registries, Snyk for Swift and Objective-C uses multiple sources, including [cocoapods.org](https://cocoapods.org/) and [swiftpackageregistry.com](https://swiftpackageregistry.com/).

## Swift for Snyk Code

For Swift with Snyk Code, Snyk supports Swift versions up to 5.7.x. You can use higher versions, but Snyk does not provide support for them.

For Swift with Snyk Code, the following file formats are supported: .`swift`

Available features:

* Reports
* Interfile analysis

## Swift for Snyk Open Source

For Swift with Snyk Open Source, you can use Swift 3.0 or higher.

### Supported file formats

For Swift and Objective-C with Snyk Open Source, the following file formats are supported:

* For CocoaPods: `podfile` and `podfile.lock`
* For Swift: `package.swift`

For Swift and Objective-C with Snyk Open Source, Snyk provides support for package managers as follows:

* For CocoaPods: CLI support, Git support, License scanning
* For Swift Package Manager: CLI support

### Requirements for Swift Package Manager

{% hint style="info" %}
Snyk supports only Projects using Swift 3.0 or higher.
{% endhint %}

In order for Snyk to discover a Project, a `Package.swift` file must be present. Also, Snyk uses the `swift package show-dependencies` command to build the dependency graph.

When the .build folder of your Project is not present - in a pipeline, for example, it is possible that Snyk takes longer to scan your Swift Projects. When the CLI runs the command `swift package show-dependencies`, Swift must resolve the dependencies as part of this process, which can add to the overall time it takes Snyk to complete the scan. Therefore, ensuring that you have your Projects built first and that the .build folder is present can speed up the CLI processing time.

Swift Package Manager supports pre-processing and post-processing. For post-processing, custom commands can add extra dependencies. Detecting such dependencies is not supported.

It is not possible to scan Swift Package Manager Projects using Git import.

### Cocoapods and CLI

To build the dependency graph, Snyk examines the `Podfile` and `Podfile.lock` files.\
\
When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync by using the `--strict-out-of-sync=true`

For Cocoapods and Git, to test your Projects, Snyk analyzes the `Podfile` and `Podfile.lock` files.

The following summarizes support for Git import and testing.

### Features

Available features:

* License scanning (CocoaPods)
* Reports

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Swift and Objective-C. This means that you will not be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}
