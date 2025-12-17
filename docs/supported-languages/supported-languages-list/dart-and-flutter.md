# Dart and Flutter

{% hint style="info" %}
Dart and Flutter are supported for Snyk Code and Snyk Open Source. 
{% endhint %}

{% hint style="info" %}
Code analysis support for Dart is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](../../snyk-platform-administration/snyk-preview.md).
{% endhint %}

## Dart and Flutter for Snyk Code

For an overview of the supported security rules, visit [Dart/Flutter rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/dart-flutter-rules.md).

### Supported frameworks and libraries

The following frameworks and libraries are supported:

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

## Applicability and integration

The following functions are available for Dart and Flutter:

Available functions: 

* Test your app's SBOM using `pkg:pub`
* Test your app's packages using `pkg:pub`

## Technical specifications

* Supported package managers:  Pub
* Supported package registry: [pub.dev](https://pub.dev/)

Snyk features are not supported for Dart and Flutter. You can test an open-source package from Pub package manager:

* Using the API endpoint [List issues for a package](../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues)
* Using the [SBOM CLI](../../developer-tools/snyk-cli/commands/sbom.md).

## Testing platform dependencies (iOS, macOS, Android) in Flutter apps

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