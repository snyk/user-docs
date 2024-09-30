# Dart and Flutter

{% hint style="info" %}
The Snyk API is available only for Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Snyk for Dart and Flutter is supported **only for Open Source**. The following summarizes support for Dart and Flutter:

**Package manager**: Pub

**Package registry**: [pub.dev](https://pub.dev/)

**Test your app's packages**: Available, `pkg:pub`

Snyk supports the testing of open-source Dart and Flutter packages from the Pub package manager  using the endpoint [List issues for a package](../snyk-api/reference/issues.md#orgs-org\_id-packages-purl-issues):\
`GET /orgs/{org_id}/packages/{purl}/issues endpoint`

The endpoint returns known vulnerabilities for the package. For more information, see the the page[List issues for a package](../snyk-api/how-to-use-snyk-sbom-and-list-issues-apis/list-issues-for-a-package.md).







