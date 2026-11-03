---
description: Snyk support for Swift and Objective-C with Snyk Code and Snyk Open Source, including Early Access Objective-C code analysis on Enterprise plans
nav_context: agnostic
last_reviewed: 2026-09-14
last_reviewed_source: focused-edit
---

# Swift and Objective-C

{% hint style="info" %}
Snyk supports Swift and Objective-C for Snyk Code and Snyk Open Source.
{% endhint %}

## Swift and Objective-C for Snyk Code

{% hint style="info" %}
Code analysis support for Objective-C is in Early Access and is available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://docs.snyk.io/platform-administration/snyk-hierarchy/snyk-preview).
{% endhint %}

For an overview of the supported security rules, visit [Swift rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/swift-rules) and [Objective-C rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/objective-c-rules).

For Swift with Snyk Code, Snyk supports Swift versions up to 5.7.x. You can use higher versions, but Snyk does not provide support for them.

### Supported frameworks and libraries

For Swift and Objective-C, Snyk supports the following frameworks and libraries:

{% columns %}
{% column %}
* AFNetworking
* Alamofire
* AppKit
* Asynchttpclient
* CocoaLumberjack
* CommonCrypto (Native API)
* Commoncrypt
* Cryptokit
* Cryptor
* Cryptoswift
* dylanshine/openai-kit
* Filekit
* FMDB
* Foundation
* google-gemini/generative-ai-swift
* IDZSwiftCommonCrypto
* JLRoutes
* MacPaw/OpenAI
* Mantle
* NSTask (Native API)
{% endcolumn %}

{% column %}
* OneSignal-iOS-SDK
* Pathos
* Realm
* RNCryptor
* SDWebImage
* ShellOut
* Shout
* SQLite3
* Subprocess
* Swift standard library
* SwiftUI
* SwiftCLI
* Swiftline
* SwiftShell
* UIKit
* WebKit
* XLForm
* ZXingObjC
{% endcolumn %}
{% endcolumns %}

### Supported file formats

For Swift, Snyk supports `.swift` files.

For Objective-C, Snyk supports `.m` files, and implicitly supports `.h` files.

### Available features

* Reports
* Interfile analysis

## Swift and Objective-C for Snyk Open Source

For Swift with Snyk Open Source, Snyk supports Swift versions from 3.0 up to 6.2.x.

### Supported package managers and registries

For Swift and Objective-C, Snyk supports the following package managers: CocoaPods and Swift Package Manager v3.0 or higher.

As package registries, Snyk for Swift and Objective-C uses multiple sources, including [cocoapods.org](https://cocoapods.org/) and [swiftpackageregistry.com](https://swiftpackageregistry.com/).

### Supported file formats

For Swift and Objective-C with Snyk Open Source, Snyk supports the following file formats:

* For CocoaPods: `Podfile` and `Podfile.lock`
* For Swift: `Package.swift`

For Swift and Objective-C with Snyk Open Source, Snyk provides support for package managers as follows:

* For CocoaPods: CLI support, SCM support, and license scanning
* For Swift Package Manager: CLI support

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Swift and Objective-C. This means that Snyk does not notify you if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available**.
* **"Fixed in" available** is set to **Yes**.
{% endhint %}

### Requirements for Swift Package Manager

{% hint style="info" %}
Snyk supports only Projects using Swift 3.0 or higher.
{% endhint %}

For Snyk to discover a Project, a `Package.swift` file must be present. Snyk uses the `swift package show-dependencies` command to build the dependency graph.

When the `.build` folder of your Project is not present, for example in a pipeline, Snyk can take longer to scan your Swift Projects. When the CLI runs the `swift package show-dependencies` command, Swift must resolve the dependencies as part of this process, which can add to the overall time it takes Snyk to complete the scan. To speed up CLI processing, build your Projects first so that the `.build` folder is present.

Swift Package Manager supports pre-processing and post-processing. For post-processing, custom commands can add extra dependencies. Snyk does not detect such dependencies.

You cannot scan Swift Package Manager Projects using SCM import.

### CocoaPods and CLI

To build the dependency graph, Snyk examines the `Podfile` and `Podfile.lock` files.

When working with Swift and Objective-C Projects from the Snyk CLI, you can prevent testing any lock files that are out of sync by using the `--strict-out-of-sync=true` option.

For CocoaPods with SCM integrations, Snyk analyzes the `Podfile` and `Podfile.lock` files to test your Projects.
