---
description: The built-in Snyk API and Web scan profiles and their differences
nav_context: agnostic
---

# Built-in scan profiles and their differences

Snyk API & Web has four different built-in scan profiles:

* **Lightning**
* **Safe**
* **Normal**
* **Full**

## Lightning

The **Lightning** scan profile is designed for speed, taking less than a minute to complete, and focuses on identifying vulnerabilities related to SSL/TLS, HTTP headers, and cookies.

## Safe

The **Safe** scan profile minimizes the potential impact on the target application by testing for all supported vulnerabilities using a limited set of payloads. The scanner uses only GET requests and does not include POST, PUT, or DELETE requests. The crawler still makes requests using these methods when necessary, for example, to log in to the application.

## Normal

The **Normal** scan profile tests for all supported vulnerabilities and uses a more extensive set of payloads than the **Safe** profile without any restrictions on the methods used.

## Full

The **Full** scan profile includes all tests from the **Normal** profile and uses an even more extensive set of payloads.

## Customization

Snyk also lets you customize scan profiles when you need to adjust and fine-tune the scans for your targets. For details, visit [Customize a scan profile](customize-scan-profile.md).
