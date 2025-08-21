# Swift and Objective-C

## Applicability and integration

{% hint style="info" %}
Swift is supported for Snyk Code and Snyk Open Source.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

Available functions:&#x20;

* Test your app's SBOM using `pkg:swift`, `pkg:cocoapods`
* Test your app's packages using `pkg:swift`, `pkg:cocoapods`

## Technical specifications

### Supported package managers and registries

For Swift and Objective-C, Snyk supports the following package managers: CocoaPods, Swift Package Manager v3.0 or higher.

As package registries, Snyk for Swift and Objective-C uses multiple sources, including [cocoapods.org](https://cocoapods.org/) and [swiftpackageregistry.com](https://swiftpackageregistry.com/).

### Supported frameworks and libraries

For Swift and Objective-C, the following frameworks and libraries are supported:&#x20;

* Swift standard library - Comprehensive&#x20;
* Foundation- Comprehensive&#x20;
* Appkit - Comprehensive&#x20;
* Swift UI - Comprehensive&#x20;
* UI Kit - Comprehensive&#x20;
* Asynchttpclient - Comprehensive&#x20;
* Commoncrypt - Comprehensive&#x20;
* Commoncrypto - Comprehensive&#x20;
* Cryptokit - Comprehensive&#x20;
* Cryptoswift - Comprehensive&#x20;
* Cryptor - Comprehensive&#x20;
* AlamoFire - Comprehensive&#x20;
* Filekit - Comprehensive&#x20;
* google-gemini/generative-ai-swift - Comprehensive&#x20;
* MacPaw/OpenAI - Comprehensive&#x20;
* dylanshine/openai-kit - Comprehensive&#x20;
* Pathos - Comprehensive&#x20;
* SQLite3 - Comprehensive&#x20;
* Webkit - Comprehensive&#x20;
* SwiftCLI - Comprehensive&#x20;
* ShellOut - Comprehensive&#x20;
* SwiftShell - Comprehensive&#x20;
* Subprocess - Comprehensive&#x20;
* Shout - Comprehensive
* Swiftline - Comprehensive&#x20;
* RNCryptor - Comprehensive

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

* For CocoaPods: `podfile`, `podfile.lock`,&#x20;
* For Swift: `package.swift`

For Swift and Objective-C with Snyk Open Source, Snyk provides support for package managers as follows:&#x20;

* For CocoaPods: CLI support, Git support, License scanning
* For Swift Package Manager: CLI support

### Requirements for Swift Package Manager

{% hint style="info" %}
Snyk supports only Projects using Swift 3.0 or higher.
{% endhint %}

A `Package.swift` file must be present for the Snyk CLI to discover the Project.\
\
Snyk uses the `swift package show-dependencies` command to build the dependency graph.\
\
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

* License scanning (CocoaPods)&#x20;
* Reports

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Swift and Objective-C. This means that you will not be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}



