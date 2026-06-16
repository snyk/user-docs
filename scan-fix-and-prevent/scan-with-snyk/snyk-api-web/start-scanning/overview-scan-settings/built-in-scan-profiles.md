# Built-in scan profiles and their differences

Snyk API & Web has four different built-in scan profiles:

* **Lightning**
* **Safe**
* **Normal**
* **Full**

## Lightning

The **Lightning** scan profile is designed for speed, taking less than a minute to complete, and focuses on identifying vulnerabilities related to SSL/TLS, HTTP headers, and cookies.

## Safe

The **Safe** scan profile is designed to minimize the potential impact on the target application by testing for all supported vulnerabilities but using a limited set of payloads. Additionally, the scanner only employs GET requests and does not include POST, PUT, or DELETE requests. Nevertheless, the crawler will still make requests using these methods if necessary, for example, to log in to the application.

## Normal

The **Normal** scan profile tests for all supported vulnerabilities and uses a more extensive set of payloads than the **Safe** profile without any restrictions on the methods used.

## Full

The **Full** scan profile includes all tests from the **Normal** profile and utilizes an even more extensive set of payloads.

## Customization

Snyk API & Web also allows the customization of scan profiles in case you need to adjust and fine-tune the scans for your targets. Visit [Customize a scan profile](customize-scan-profile.md) for more details.
