# Folder trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager (for example, pip, Gradle, Maven, Yarn, npm, and so on) to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

To safeguard against using the plugin on untrusted folders, the Snyk plugin asks for folder trust before allowing you to run scans against these folders. When in doubt, do not grant trust.

<figure><img src="../../../.gitbook/assets/image (4) (2) (1) (1).png" alt="Snyk extension prompt to trust a folder"><figcaption><p>Snyk extension prompt to trust a folder</p></figcaption></figure>

Once a single project trust is granted, Snyk will not ask for trust on the opened project folder and its subfolders again. If you didn’t grant trust the first time, the plugin will ask again the next time you restart your Eclipse instance.

To revoke an existing folder trust, you can navigate to the Snyk plugin preferences in Eclipse and edit the “Trusted Folder” setting.

<figure><img src="../../../.gitbook/assets/image (5) (3) (1) (1).png" alt="Snyk Eclipse plugin preferences Trusted Folders setting"><figcaption><p>Snyk Eclipse plugin preferences Trusted Folders setting</p></figcaption></figure>
