---
description: Snyk support for Dart and Flutter with Snyk Code and Snyk Open Source, including Early Access code analysis on Enterprise plans
---

# Dart and Flutter

{% hint style="info" %}
Dart and Flutter is supported for Snyk Code and Snyk Open Source.
{% endhint %}

## Dart and Flutter for Snyk Code

{% hint style="info" %}
Code analysis support for Dart is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

For an overview of the supported security rules, visit [Dart and Flutter rules](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-code/snyk-code-security-rules/dart-and-flutter-rules).

### Supported frameworks and libraries

For Dart and Flutter with Snyk Code, Snyk supports the following frameworks and libraries:

{% columns %}
{% column %}
* crypto
* encrypt
* uuid
* basic\_utils
* pointycastle
* cryptography
* sqflite
* sqlite3
* drift
* realm
* http
* dio
* cupertino\_http
* web\_socket\_channel
* web
{% endcolumn %}

{% column %}
* cronet\_http
* flutter\_inappwebview
* webview\_flutter
* twilio\_flutter
* dart\_eval
* google\_sign\_in
* flutter\_facebook\_auth
* sign\_in\_with\_apple
* flutter\_appauth
* openid\_client
* firebase\_auth
* amplify\_flutter
* flutter\_stripe
* nfc\_manager
* mobile\_scanner
* flutter
{% endcolumn %}
{% endcolumns %}

### Supported file formats

The following file formats are supported: `.dart`

### Available features

* Reports
* Interfile analysis

## Dart and Flutter for Snyk Open Source

### Available features

* Test your app's SBOM and packages using `pkg:pub` PURLs through the [SBOM test](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/sbom-test) CLI command
* Test & monitor your Flutter apps native platform dependencies using [`snyk test`](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/test) and [`snyk monitor`](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/monitor) commands

### Testing a Dart applications pub dependency tree

Activate the pub [`sbom`](https://pub.dev/packages/sbom) package & create a minimal `sbom.yaml` file in the root folder of the repository:

```
dart pub global activate sbom
cat << EOF > sbom.yaml
type: spdx
spdx:
  SPDXFormat: 'tagvalue'
EOF
```

Use the dart `sbom` command to create a SBOM file & test it using the [`sbom test`](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/sbom-test) command:

```
dart pub global run sbom
snyk sbom test --experimental --file sbom-pub.json
```

### Testing platform dependencies (iOS, macOS, Android) in Flutter apps

Flutter applications rely on native platform dependencies to handle lower-level tasks, such as analytics, hardware access, or integrating existing functionality. These dependencies can be added through pub packages to extend functionality or integrated directly into build systems like Gradle or Cocoapods.

Snyk’s regular open-source support can scan these packages; however, a complete app build is necessary to make them available in the repository and accessible to CLI tools.

You can start by building the application for all relevant platforms. This ensures that `pub` fetches all required packages, and the Flutter build system establishes the necessary links for the native build systems.

```
flutter build apk --debug
flutter build ios --debug --no-codesign
flutter build macos --debug
```

Next, run the `snyk monitor` command to scan for native dependencies:

```
snyk monitor --all-projects --exclude=example,.symlinks
```

The `--exclude` parameter removes duplicates and ignores example applications, which are part of the plugin source code but not included in regular application builds.

You are now able to view in the Snyk Web UI all native dependencies, including those introduced by third-party plugins.

<figure><img src="../../.gitbook/assets/image (197).png" alt=""><figcaption><p>Snyk Project page showing dependencies in Flutter apps</p></figcaption></figure>
