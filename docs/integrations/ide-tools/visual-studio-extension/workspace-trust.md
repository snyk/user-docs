# Workspace trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. For example, this includes invoking the package manager (such as pip, gradle, maven, yarn, npm, and so on) to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

To safeguard from using the extension on untrusted folders, our extension will ask for folder trust before allowing you to run any scans against your code. When in doubt, do not proceed with a scan.

<figure><img src="https://lh4.googleusercontent.com/zeg6dMXmw1BUVMTeCviRJ-b2ct0CcRXFdfm1AZJFfgzRfro6TzvFv2_4DwZbqcRRHUl20poHvThC-8TdJniBQ12zgFxgZCY7b5NLZcG4XOu2qd25avtIEZM9Hzq-LvqKmqDS4N3uAFydWTT99x4HMNi0g_2LwFH2LzKwD98KAt1gYxZHitFK3PwbpB7pFA" alt=""><figcaption></figcaption></figure>

After a single folder trust is granted, Snyk will not ask for trust on the opened folder and its subfolders again. To revoke an existing folder trust, manually edit the **trustedFolders** JSON element in the **settings.json** file located in the following path:

%LocalAppData%\Microsoft\VisualStudio\\\<Visual Studio version>\Extensions\Snyk\Snyk Security - Code and Open Source Dependencies\\\<Snyk extension version>
