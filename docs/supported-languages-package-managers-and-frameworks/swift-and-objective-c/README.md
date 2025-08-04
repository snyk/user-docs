# Swift and Objective-C

## Applicability

Snyk supports [Swift for code analysis](swift-for-code-analysis.md) and [Swift and Objective-C for open source](swift-and-objective-c-for-open-source.md). This includes support for CocoaPods.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

{% hint style="info" %}
**Supported Swift version**&#x20;

* For Snyk Open Source, you can use Swift 3.0 or higher.
* For Snyk Code, Snyk supports Swift versions up to 5.7.x. You can use higher versions but Snyk does not provide support for them.
{% endhint %}

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code. If used with Snyk Open Source, then the SCM import is available for CocoaPods.
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:swift`, `pkg:cocoapods`
* Test your app's packages using `pkg:swift`, `pkg:cocoapods`

## Package managers and supported files

Snyk for Swift and Objective-C supports CocoaPods, Swift Package Manager v3.0 or higher as package managers and it uses multiple sources including [cocoapods.org](https://cocoapods.org/) and [swiftpackageregistry.com](https://swiftpackageregistry.com/) for package registry.

Snyk for Swift and Objective-C supports the following file formats:

* Snyk Open Source:&#x20;
  * For CocoaPods: `podfile`, `podfile.lock`,&#x20;
  * For Swift: `package.swift`
* Snyk Code: .`swift`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Swift and Objective-C:&#x20;

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

## Features

The following features are supported in Snyk for Swift and Objective-C:

| Snyk Open Source                                                | Snyk Code                                            |
| --------------------------------------------------------------- | ---------------------------------------------------- |
| <ul><li>License scanning (CocoaPods) </li><li>Reports</li></ul> | <ul><li>Reports</li><li>Interfile analysis</li></ul> |

PR Checks configured to “Only fail when the issues found have a fix available” rely on Snyk FixPR support and will not alert for Swift and Objective-C or other languages that do not support FixPRs.
