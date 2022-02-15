# How Snyk finds out about new vulnerabilities

Snyk’s security team, based in Israel, maintains our [vulnerability database](https://security.snyk.io).

This work includes curating vulnerabilities found or reported elsewhere on the web, as well as doing our own research to uncover previously unknown vulnerabilities, which we then responsibly disclose. Snyk Enterprise gets early notifications for issues our research uncovers alongside this responsible disclosure process.

Most of the vulnerabilities in our database originate from one of these sources:

1. _Monitoring other vulnerability databases_, such as CVEs from [NVD](https://nvd.nist.gov) and many others.
2. _Monitoring user activity on GitHub_, including issues, PRs and commit messages that may indicate a vulnerability.
3. _Bulk research_, using tools that look for repeated security mistakes across open source package code
4. _Manual research_, investing our researchers time to manually audit more widely used packages for security flaws.

For every issue deemed to be a real vulnerability, we assign the right CVSS (severity) score and package version specification, create an advisory, and make it available in the product.
