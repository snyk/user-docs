# Swift and Objective-C

{% hint style="info" %}
Swift and Objective-C are supported for Snyk Code and Snyk Open Source.
{% endhint %}

## Swift and Objective-C for Snyk Code

{% hint style="info" %}
Code analysis support for Objective-C is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](../../snyk-platform-administration/snyk-preview.md).
{% endhint %}

For an overview of the supported security rules, see the following pages:
* [Swift rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/swift-rules.md)
* [Objective-C rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/objective-c-rules.md)

For Swift with Snyk Code, Snyk supports Swift versions up to 5.7.x. You can use higher versions, but Snyk does not provide support for them.

### Supported frameworks and libraries

For Swift and Objective-C, the following frameworks and libraries are supported:

{% columns %}
{% column %}
* AFNetworking
* AlamoFire
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
{% endcolumn %}

{% column %}
* Mantle
* NSTask (Native API)
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
* Swift UI
* SwiftCLI
* Swiftline
* SwiftShell
* UI Kit
* Webkit
* XLForm
* ZXingObjC
{% endcolumn %}
{% endcolumns %}

### Supported file formats

* For Swift: .`swift`
* For Objective-C: Snyk supports `.m` files, and implicitly supports `.h` files.

### Available features

* Reports
* Interfile analysis

## Swift and Objective-C for Snyk Open Source

### Supported package managers and registries

For Swift and Objective-C, Snyk supports the following package managers: CocoaPods, Swift Package Manager v3.0 or higher.

As package registries, Snyk for Swift and Objective-C uses multiple sources, including [cocoapods.org](https://cocoapods.org/) and [swiftpackageregistry.com](https://swiftpackageregistry.com/).

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

To build the dependency graph, Snyk examines the `Podfile` and `Podfile.lock` files.

When working with Swift and Objective-C projects from the Snyk CLI, you can prevent testing any lock files that are out-of-sync by using the `--strict-out-of-sync=true`

For Cocoapods and Git, to test your Projects, Snyk analyzes the `Podfile` and `Podfile.lock` files.

### Features

Available features:

* License scanning (CocoaPods)
* Reports

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Swift and Objective-C. This means that you will not be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available**.
* **"Fixed in" available** is set to **Yes**.
{% endhint %}
