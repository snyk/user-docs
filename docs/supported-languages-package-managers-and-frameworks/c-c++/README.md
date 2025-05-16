# C/C++

## Applicability

C/C++ is supported for Snyk Open Source and Snyk Code.

Specific considerations apply for the [Snyk CLI for open-source C++ scans](snyk-cli-for-open-source-c++-scans.md). [Guidance for Snyk for C/C++](guidance-for-snyk-for-c-c++.md) is provided.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* SCM import is available only for Snyk Code.
* Test or monitor your app through CLI and IDE for both Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:generic.`
* Test your app's packages using `pkg:generic.`

For Snyk IDE, no additional options are required. The Snyk plugin has views within the IDE for displaying results.

## Package managers and supported file extensions

Snyk for C/C++ does not support any package managers but supports the following file formats:

* Snyk Open Source: N/A
* Snyk Code: `.c`, `.cc`, `.cpp`, `.cxx`, `.h`, `.hpp`, `.hxx`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for C/C++:

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

The following features are supported in Snyk for C/C++:

| Snyk Open Source                                   | Snyk Code          |
| -------------------------------------------------- | ------------------ |
| <ul><li>License scanning</li><li>Reports</li></ul> | Interfile analysis |

PR Checks configured to “Only fail when the issues found have a fix available” rely on Snyk FixPR support and will not alert for C/C++ or other languages that do not support FixPRs.
