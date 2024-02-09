# Workspace trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager (for example, pip, gradle, maven, yarn, npm, and so on) to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

In addition to the built-in [Workspace Trust](https://code.visualstudio.com/docs/editor/workspace-trust) feature of Visual Studio Code, our extension will ask for folder trust before allowing to run any scans against your code. When in doubt, do not proceed with a scan.

<figure><img src="../../../.gitbook/assets/vscode-trust (1) (1).png" alt=""><figcaption><p>Request to trust folders</p></figcaption></figure>

Once a single project trust is granted, Snyk will not ask for trust on the opened folder and its subfolders any more. If you didn’t grant trust the first time, the plugin will ask next time again when you restart your Visual Studio Code instance.

To revoke an existing folder trust or grant trust to more folders, you can navigate to the Snyk extension settings and edit the **Trusted Folders** (`snyk.trustedFolders`) setting.
