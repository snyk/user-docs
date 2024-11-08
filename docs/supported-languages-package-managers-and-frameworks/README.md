# Supported languages, package managers, and frameworks

{% hint style="info" %}
The language support pages are being extensively revised and updated.

The language support pages provide detailed information about language support for **Snyk Code** and **Snyk Open Source**.

For information about language support for **Snyk Container**, see [Supported workloads, container registries, languages, and operating systems](../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/supported-workloads-container-registries-languages-and-operating-systems.md) and [Operating system distributions supported by Snyk Container](../scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md). For **IaC** language support, see [Supported IaC languages, cloud providers, and cloud resources](../scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/).
{% endhint %}

The list of languages supported for Snyk Code and Snyk Open Source follows. For details of what is supported, see [Snyk language support details](snyk-language-support-details.md). For guidance on using Snyk with each language, see the individual language pages.

| Language              | Snyk Code | Open Source |
| --------------------- | --------- | ----------- |
| Apex                  | ✔️        |             |
| Bazel                 |           | ✔️          |
| C++                   | ✔️        | ✔️          |
| Dart and Flutter      |           | ✔️          |
| Elixir                |           | ✔️          |
| Go                    | ✔️        | ✔️          |
| Java and Kotlin       | ✔️        | ✔️          |
| JavaScript            | ✔️        | ✔️          |
| .NET                  | ✔️        | ✔️          |
| PHP                   | ✔️        | ✔️          |
| Python                | ✔️        | ✔️          |
| Ruby                  | ✔️        | ✔️          |
| Rust                  |           | ✔️          |
| Scala                 | ✔️        | ✔️          |
| Swift and Objective-C | ✔️        | ✔️          |
| TypeScript            | ✔️        | ✔️          |
| VB NET                | ✔️        |             |

{% hint style="info" %}
Interfile analysis in Snyk Code is available for all languages supported except Ruby.

For Snyk Open Source, only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of Projects that have a package manager, this means a release to the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++) this requires an official release or tag on the GitHub repository.
{% endhint %}
