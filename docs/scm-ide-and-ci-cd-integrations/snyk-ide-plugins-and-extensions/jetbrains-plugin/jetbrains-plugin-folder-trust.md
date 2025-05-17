# JetBrains plugin folder trust

Snyk Open Source may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager, for example, pip, Gradle, Maven, Yarn, npm, and so on, to get dependency information. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

To safeguard against using the extension on untrusted Projects, the Snyk plugin asks for Project trust before allowing you to run any scans against your code. When in doubt, do not proceed with a scan.

<figure><img src="../../../.gitbook/assets/modal-dialog copy.png" alt="Snyk extension prompt to trust a project"><figcaption><p>Snyk extension prompt to trust a project</p></figcaption></figure>

After a single Project trust is granted, Snyk will not ask for trust on the opened Project folder and its subfolders again. To revoke an existing folder trust, manually edit the `TRUSTED_PATHS` option in the `snyk.settings.xml` located in the [JetBrains IDE configuration options directory](https://www.jetbrains.com/help/idea/directories-used-by-the-ide-to-store-settings-caches-plugins-and-logs.html#config-directory).

Snyk recommends using the built-in [Trusted Projects](https://plugins.jetbrains.com/docs/intellij/trusted-projects.html) provided in the IntelliJ platform version 2021.2.4/2021.3.1 and higher for an extra layer of security.
