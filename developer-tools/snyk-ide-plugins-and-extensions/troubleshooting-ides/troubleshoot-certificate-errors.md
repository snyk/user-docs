# Troubleshoot certificate errors

## Problem <a href="#problem" id="problem"></a>

Error message **Unable to get local issuer certificate** or something with certificate path.

## Solution <a href="#solution" id="solution"></a>

* Add all certificates in the correct chain to the operating system, the environment (JDK), or both that the certificate is running on.
* Set **Allow unknown certificate authority** in the IDE or
* Pass the `--insecure` argument to the Snyk CLI executable using settings, **This does not solve Download, API access, and Snyk Code communication.**

## Custom Certificate Authorities <a href="#custom-certificate-authorities" id="custom-certificate-authorities"></a>

### Java <a href="#java.1" id="java.1"></a>

The CA and intermediate certificates must be added to the certificate store of each JDK used.

See [Import the Certificate as a Trusted Certificate (The Java™ Tutorials > Security Features in Java SE > Signing Code and Granting It Permissions)](https://docs.oracle.com/javase/tutorial/security/toolsign/rstep2.html).

### Eclipse <a href="#eclipse" id="eclipse"></a>

<figure><img src="../../../.gitbook/assets/image (173).png" alt="Allow unknown certificate authorities"><figcaption><p>Allow unknown certificate authorities</p></figcaption></figure>

* Update the JDK used by Snyk scans to add unknown certificates.
* Update to the latest CLI and pugin version.
* CLI download can be re-triggered by deleting the CLI. The path is shown in the Snyk Preferences.

### IntelliJ <a href="#intellij" id="intellij"></a>

<figure><img src="../../../.gitbook/assets/image (174).png" alt="IntelliJ setting"><figcaption><p>IntelliJ setting</p></figcaption></figure>

* See [Trusted root certificates | IntelliJ IDEA](https://www.jetbrains.com/help/idea/ssl-certificates.html).
* Updating the Jetbrains certificate handling only will most likely not be sufficient, as the CLI does not use Jetbrains settings, but JAVA\_HOME and PATH to determine a JDK. This JDKs certificate store must be updated.

### VSCode <a href="#vscode" id="vscode"></a>

If you face certificate validation errors in VS Code, ensure Node.js and the Snyk extension trust the certificate authority (CA) of your organization.

#### Configure the environment variable (Recommended)

Use the `NODE_EXTRA_CA_CERTS` environment variable to configure Node.js to trust your CA certificate. Snyk recommends this approach because it is reliable and works across the Snyk CLI, VS Code extension, JetBrains plugins, and other Node-based developer tools.

1. Set the variable: `NODE_EXTRA_CA_CERTS=path-to-your-ca-certificate`
2. Restart VS Code.

#### Windows-only workaround

Use a VS Code extension, for example, `win-ca`, to expose Windows trusted root certificates to Node.js. This extension is not actively maintained and might not work in all environments. Use this workaround only if you cannot configure the environment variable.

#### Verify the certificate chain

Ensure that:

* The full certificate chain (root and intermediates) is installed.
* The operating system trusts the certificates.
* You restart VS Code after you make changes.

#### Disable strict SSL

If certificate validation fails, disable strict SSL validation temporarily to troubleshoot. Snyk does not recommend this for long-term use due to security risks.
