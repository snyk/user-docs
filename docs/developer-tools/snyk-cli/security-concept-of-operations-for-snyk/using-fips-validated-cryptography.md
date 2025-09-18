# Using FIPS-validated cryptography

## Availability of FIPS-validated cryptography in Snyk

Support for use of FIPS-validated cryptography is limited to the Windows and Linux operating systems.

| Operating System | FIPS support |
| ---------------- | ------------ |
| Windows          | ✅            |
| Linux            | ✅            |
| Alpine           | ⛔            |
| macOS            | ⛔            |

## FIPS-validated cryptography support and use in the Snyk CLI and Snyk Language Server

To optimize the developer experience, Snyk is combining the [Snyk Language Server](../../snyk-ide-plugins-and-extensions/snyk-language-server/) and the [Snyk CLI](../getting-started-with-the-snyk-cli.md). As a first step, Snyk is bringing FIPS binaries under one application. Later also non-FIPS CLI binaries will be used for Snyk Language Server.

The Snyk Language Server can now be executed as a CLI command.

```bash
> snyk language-server

# instead of 
> snyk-ls
```

As a consequence, instructions for using FIPS-validated cryptography are the same for the CLI and the Language Server.

### Prerequisites for FIPS-cryptography in the CLI and Snyk Language Server

**Linux operating systems**

On Linux, Snyk supports FIPS-validated cryptography through OpenSSL and its validated FIPS provider.

Ensure that your Linux system has OpenSSL installed and configured to meet FIPS validation requirements. For information about how to accomplish this, see the documentation from the [OpenSSL project](https://www.openssl.org/docs/fips.html).

**Windows operating systems**

On Windows, Snyk supports FIPS-validated cryptography through the Windows CNG API.

To enable FIPS on Windows, [use the Windows FIPS policy](https://docs.microsoft.com/en-us/windows/security/threat-protection/fips-140-validation#step-3-enable-the-fips-security-policy).

For testing, FIPS can be enabled using the following registry key   `HKLM\SYSTEM\CurrentControlSet\Control\Lsa\FipsAlgorithmPolicy` by setting the value of `Enabled` to 1.

#### Download FIPS-enabled binaries

Snyk binaries are available with and without FIPS support. They are all hosted on downloads.snyk.io, differentiated by their Base URL.

**FIPS Base URL:** https://downloads.snyk.io/fips/

**Regular Base URL:** https://downloads.snyk.io/

All instructions on [how to install and use Snyk](../install-or-update-the-snyk-cli/) remain the same. The only required change is using the appropriate Base URL.

### Example of downloading and running the FIPS-enabled Snyk CLI

The example that follows uses a [Microsoft Mariner](https://mcr.microsoft.com/en-us/product/cbl-mariner/base/core/about) image to download and run the FIPS-enabled Snyk CLI.

```bash
docker run -it mcr.microsoft.com/cbl-mariner/base/core:2.0 bash
> tdnf install -y ca-certificates
> curl --compressed https://downloads.snyk.io/fips/cli/latest/snyk-linux-arm64 -o snyk
> chmod +x snyk
> ./snyk -d
...
2023-09-13T11:02:49Z main - Features:
2023-09-13T11:02:49Z main -   oauth:               Disabled
2023-09-13T11:02:49Z main -   fips:                Enabled
...
```

### Troubleshooting FIPS-cryptography in the Snyk CLI

`not in FIPS mode` errors indicate that the underlying cryptography library is not in FIPS mode. To solve these issues ensure that the [prerequisites](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server) are met.

## FIPS-validated cryptography support and use in IDE Integrations

### Visual Studio Code

To make use of FIPS-validated cryptography in the [Snyk Visual Studio Code integration](../../snyk-ide-plugins-and-extensions/visual-studio-code-extension/), do the following:

* Ensure the [prerequisites](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server) are met.
* [Download the appropriate FIPS-enabled binaries](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* Disable automatic binary management in the Snyk settings.
* Configure the integration to use the binary by setting the Language Server Path and the CLI Path to the same binary.

### Eclipse

To make use of FIPS-validated cryptography in the [Snyk Eclipse integration](../../snyk-ide-plugins-and-extensions/eclipse-plugin/), do the following:

* Ensure the [prerequisites](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server) are met
* [Download the appropriate FIPS-enabled binaries.](using-fips-validated-cryptography.md#download-fips-enabled-binaries)
* Disable automatic binary management in the Snyk preferences.
* Configure the integration to use the binary by setting the Language Server Path and the CLI Path to the same executable.
* Configure the Java Runtime to use a FIPS-validated[ JCE (Java Cryptography Extension)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search?SearchMode=Basic\&ModuleName=java\&CertificateStatus=Active\&ValidationYear=0).

### JetBrains

To make use of FIPS-validated cryptography in the [Snyk JetBrains integration](../../snyk-ide-plugins-and-extensions/jetbrains-plugin/), do the following:

* Ensure the [prerequisites](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server) are met.
* [Download the appropriate FIPS-enabled binaries](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* Disable automatic binary management in the Snyk preferences.
* Configure the integration to use the binary by setting the CLI Path.
* Configure the Java Runtime to use a FIPS-validated[ JCE (Java Cryptography Extension)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search?SearchMode=Basic\&ModuleName=java\&CertificateStatus=Active\&ValidationYear=0).

### Visual Studio

To make use of FIPS-validated cryptography in the [Snyk Visual Studio integration](../../snyk-ide-plugins-and-extensions/visual-studio-extension/) do the following:

* Ensure the [prerequisites](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server) are met.
* [Download the appropriate FIPS-enabled binaries](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* Disable automatic binary management.
* Configure the integration to use the binary, by setting the CLI Path.

## FIPS-validated cryptography support and use in CI/CD Integrations

FIPS in [CI/CD Integrations](../../snyk-ci-cd-integrations/) is available only by using a FIPS-enabled CLI directly.

## FIPS-validated cryptography support and use in Package Repositories

The  [Artifactory Gatekeeper plugin](../../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/artifactory-gatekeeper-plugin.md) uses the Snyk API and runs on a Java VM. To make use of FIPS-validated cryptography, configure the Java Runtime to use a FIPS-validated[ JCE (Java Cryptography Extension)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search?SearchMode=Basic\&ModuleName=java\&CertificateStatus=Active\&ValidationYear=0).
