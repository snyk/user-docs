# Dart and Flutter

{% hint style="info" %}
Dart and Flutter are supported only for Open Source.
{% endhint %}

## Applicability

The following functions are available for Dart and Flutter:

* Test your app's SBOM using `pkg:pub`
* Test your app's packages using `pkg:pub`

## Package managers and supported file extensions

For Dart and Flutter, Snyk supports Pub as a package manager and [pub.dev](https://pub.dev/) as a package registry. It does not support any file formats.

## Frameworks and libraries

Snyk does not support frameworks and libraries for Dart and Flutter.&#x20;

## Features

Snyk features are not supported for Dart and Flutter. You can test an open-source package from Pub package manager:

* Using the API endpoint [List issues for a package](../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues)
* Using the [SBOM CLI](../snyk-cli/commands/sbom.md).

## Testing platform dependencies (iOS, macOS, Android) in Flutter Apps

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

<figure><img src="../.gitbook/assets/image (571).png" alt=""><figcaption><p>Snyk Project page showing dependencies in Flutter apps</p></figcaption></figure>
