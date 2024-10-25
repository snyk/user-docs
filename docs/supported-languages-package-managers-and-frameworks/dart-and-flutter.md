# Dart and Flutter

Snyk for Dart and Flutter is supported **only for Open Source**. The following summarizes support for Dart and Flutter:

**Package manager**: Pub

**Package registry**: [pub.dev](https://pub.dev/)

**Test your app's SBOM**: Available, `pkg:pub`

**Test your app's packages**: Available, `pkg:pub`

## Testing pub packages using the API

{% hint style="info" %}
The Snyk API is available only for Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Snyk supports the testing of open-source packages from the Pub package manager using the API endpoint [List issues for a package](../snyk-api/reference/issues.md#orgs-org\_id-packages-purl-issues):  &#x20;

```
GET /orgs/{org_id}/packages/{purl}/issues endpoint
```

The endpoint returns known vulnerabilities for the package. For more information, see the the page[List issues for a package](../snyk-api/how-to-use-snyk-sbom-and-list-issues-apis/list-issues-for-a-package.md).

## Testing pub packages using the SBOM CLI

You can also test an SBOM using the [SBOM CLI](../snyk-cli/commands/sbom.md). You must first create an SBOM file. For example, you can use`cdxgen`to extract the SBOM to be sent to the Snyk CLI as follows:

```
cdxgen -t pub -o pub-sbom.json \
  && snyk sbom test --experimental --file pub-sbom.json
```

## Testing platform dependencies (iOS, macOS, Android) in Flutter Apps

Flutter applications often rely on native platform dependencies to handle lower-level tasks, such as analytics, hardware access, or integrating existing functionality. These dependencies can be added through pub packages to extend functionality or integrated directly into build systems like Gradle or Cocoapods.

Snyk’s regular open-source support can scan these packages; however, a complete app build is necessary to make them available in the repository and accessible to CLI tools.

Start by building the application for all relevant platforms. This ensures that `pub` fetches all required packages, and the Flutter build system establishes the necessary links for the native build systems.

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

Now you will be able to view all native dependencies, including those introduced by third-party plugins, in the Snyk Web UI.

<figure><img src="../.gitbook/assets/image (571).png" alt="Snyk Project page showing dependencies in Flutter apps"><figcaption><p>Snyk Project page showing dependencies in Flutter apps</p></figcaption></figure>

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;





