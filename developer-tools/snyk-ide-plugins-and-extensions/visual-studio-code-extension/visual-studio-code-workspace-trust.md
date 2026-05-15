# Visual Studio Code workspace trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager, for example, pip, Gradle, Maven, Yarn, npm, and so on, to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

In addition to the built-in [Workspace Trust](https://code.visualstudio.com/docs/editor/workspace-trust) feature of Visual Studio Code, the Snyk extension will ask for folder trust before allowing you to run any scans against your code. When in doubt, do not proceed with a scan.

<figure><img src="../../../.gitbook/assets/vscode-trust.png" alt=""><figcaption><p>Request to trust folders</p></figcaption></figure>

After a single Project trust is granted, Snyk will not ask for trust on the opened folder and its subfolders again. If you did not grant trust the first time, the extension will ask again the next time when you restart your Visual Studio Code instance.

To revoke an existing folder trust or grant trust to more folders, you can navigate to the Snyk extension settings and edit the **Trusted Folders** (`snyk.trustedFolders`) setting.
