# Workspace trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager (for example, pip, gradle, maven, yarn, npm, and so on) to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

In addition to providing the built-in [Workspace Trust feature](https://code.visualstudio.com/docs/editor/workspace-trust) of Visual Studio Code, the Snyk extension asks for folder trust before allowing any scans to run against your code. When in doubt, do not grant trust.

<figure><img src="../../.gitbook/assets/vscode-trust.png" alt="Request to trust folders"><figcaption><p>Request to trust folders</p></figcaption></figure>

Once folder trust is granted, Snyk will not ask for trust on the opened folder and its subfolders again. If you didn’t grant trust the first time, the plugin will ask again when you restart your Visual Studio Code instance.

To revoke an existing folder trust or grant trust to more folders, navigate to the Snyk extension settings and edit the **Trusted Folders** (`snyk.trustedFolders`) setting.
