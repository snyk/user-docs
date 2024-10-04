# C/C++ for open source

## C/C++ for open source support

**Package manager**: NA

**Package registry**: No single registry, multiple sources&#x20;

**Import your app through SCM**: NA

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:generic`

**Test your app's packages**: Available, `pkg:generic`

**Features**:

* License scanning
* Reports

**Package manager versions**: NA

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of projects that have a package manager, this means a release to the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the GitHub repo.
{% endhint %}

## Open source dependency management

Snyk features that support the management of open-source dependencies include the following:

<table><thead><tr><th width="267">Package managers / Features</th><th>CLI support</th><th>Git support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>C/C++</td><td>✔︎</td><td></td><td>✔︎</td><td></td></tr></tbody></table>

For information about managing dependencies and licenses from your developer workflows through policy, see the following

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

## Open source license compliance

To check compliance for open source licenses, see [Snyk License Compliance Management.](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md)

## IDE for C++ for open-source dependencies

IDE Under **Additional Parameters** in the IDE settings, enter the **--unmanaged** option to scan for C/C++ open source dependencies.

<figure><img src="../../.gitbook/assets/Screenshot 2023-08-04 at 12.18.01.png" alt="Scan for dependencies"><figcaption><p>Scan for dependencies</p></figcaption></figure>

## Troubleshooting C++ for open source

See Troubleshooting C++ for open source. If you need additional help, [contact Snyk Support](https://support.snyk.io/hc/en-us).
