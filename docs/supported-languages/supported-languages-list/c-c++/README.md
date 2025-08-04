# C/C++

## Applicability

C/C++ is supported for Snyk Open Source and Snyk Code.

Specific considerations apply for the [Snyk CLI for open-source C++ scans](snyk-cli-for-open-source-c++-scans.md). [Guidance for Snyk for C/C++](guidance-for-snyk-for-c-c++.md) is provided.

The following functions are available for C/C++:

* SCM import - available only for Snyk Code.
* Test or monitor your app through CLI and IDE
* Test your app's SBOM using `pkg:generic` or `pkg:conan`. For more information, see [Test an SBOM document for vulnerabilities](../../../snyk-api/using-specific-snyk-apis/sbom-apis/rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md).
* Test your app's packages using `pkg:generic` or `pkg:conan`. For more information, see [List issues for a package](../../../snyk-api/using-specific-snyk-apis/issues-list-issues-for-a-package.md).

## Package managers and supported file extensions

For Conan, Snyk supports [conan.io](https://conan.io/center) as a package registry.&#x20;

No additional options are required for the Snyk IDE. You can display results within the IDE using the Snyk plugin views.

For C/C++ for Snyk Code, Snyk supports the following file formats: `.c`, `.cc`, `.cpp`, `.cxx`, `.h`, `.hpp`, `.hxx`.

## Frameworks and libraries

For C/C++, Snyk supports the following frameworks and libraries:

* argparse parser - Comprehensive
* Asio Library - Comprehensive
* Boost Library - Partial
* Botan LIbrary - Comprehensive
* C Standard Library - Comprehensive
* C++ Standard Library - Comprehensive
* Curl library - Comprehensive
* fstream framework - Comprehensive
* grpc-cpp library - Comprehensive
* HTTPlib framework - Comprehensive
* JsonCpp library - Comprehensive
* liboai framework - Comprehensive
* libpq library - Comprehensive
* libpqxx framework - Comprehensive
* libsodium library - Comprehensive
* LibTomCrypt framework - Comprehensive
* libxml2 framework - Comprehensive
* MySQL framework - Comprehensive
* OpenSSL framework- Comprehensive
* POSIX LIbrary - Comprehensive
* pugixml library - Comprehensive
* SQLite library - Comprehensive
* WinHTTP framework - Comprehensive
* Xerces libraries - Comprehensive

## Features

&#x20;For C/C++, Snyk supports the following features:

| Snyk Open Source                                   | Snyk Code          |
| -------------------------------------------------- | ------------------ |
| <ul><li>License scanning</li><li>Reports</li></ul> | Interfile analysis |

{% hint style="info" %}
The **Snyk FixPR** feature is not available for C/C++. This means that you will not be notified if the PR checks fail when the following conditions are met

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}
