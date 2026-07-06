# Code execution warning for Snyk CLI

{% hint style="warning" %}
As part of examining the codebase for vulnerabilities, Snyk CLI can automatically execute code on your computer to obtain additional data for analysis.
{% endhint %}

This includes invoking the package manager, for example, pip, Gradle, Maven, Yarn, npm, and so on, to get dependency information for Snyk Open Source.

Invoking these programs on untrusted code that has malicious configurations can expose your system to malicious code execution and exploits.

Always ensure you understand and trust the code in the directory you intend to scan with Snyk CLI.

When in doubt, do not proceed with a scan.
