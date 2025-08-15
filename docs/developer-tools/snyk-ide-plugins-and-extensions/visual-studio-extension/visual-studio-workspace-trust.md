# Visual Studio workspace trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager, for example, pip, Gradle, Maven, Yarn, npm, and so on, to get dependency information. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

To safeguard against using the extension on untrusted folders, the Snyk extension asks for folder trust before allowing you to run any scans against your code. When in doubt, do not proceed with a scan.

After a single folder trust is granted, Snyk will not ask for trust on the opened folder and its subfolders again. To revoke an existing folder trust, manually edit the `trustedFolders` JSON element in the `settings.json` file located in the following path:

`%LocalAppData%\Microsoft\VisualStudio\<Visual Studio version>\Extensions\Snyk\Snyk Security\<Snyk extension version>`
