# Dart and Flutter

{% hint style="info" %}
Dart and Flutter is supported for Snyk Code and Snyk Open Source.
{% endhint %}

## Dart and Flutter for Snyk Code

{% hint style="info" %}
Code analysis support for Dart is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](../../snyk-platform-administration/snyk-preview.md).
{% endhint %}

For an overview of the supported security rules, visit [Dart and Flutter rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/dart-and-flutter-rules.md).

### Supported frameworks and libraries

For Dart and Flutter, Snyk supports the following frameworks and libraries:

{% columns %}
{% column %}
* crypto
* encrypt
* uuid
* basic_utils
* pointycastle
* cryptography
* sqflite
* sqlite3
* drift
* realm
* http
* dio
* cupertino_http
* web_socket_channel
* web
{% endcolumn %}

{% column %}
* cronet_http
* flutter_inappwebview
* webview_flutter
* twilio_flutter
* dart_eval
* google_sign_in
* flutter_facebook_auth
* sign_in_with_apple
* flutter_appauth
* openid_client
* firebase_auth
* amplify_flutter
* flutter_stripe
* nfc_manager
* mobile_scanner
* flutter
{% endcolumn %}
{% endcolumns %}

### Supported file formats

The following file formats are supported: `.dart`

### Available features

- Reports
- Interfile analysis

## Dart and Flutter for Snyk Open Source

### Available features

* Test your app's SBOM and packages using `pkg:pub` PURLs, using the [`sbom test`](../../developer-tools/snyk-cli/commands/sbom-test.md) command
* Test & monitor your Flutter apps native platform dependencies using [`snyk test`](../../developer-tools/snyk-cli/commands/test.md) and [`snyk monitor`](../../developer-tools/snyk-cli/commands/monitor.md) commands

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

Use the dart `sbom` command to create a SBOM file & test it using the [`sbom test`](../../developer-tools/snyk-cli/commands/sbom-test.md) command:

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
