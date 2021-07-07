# How Snyk finds out about new vulnerabilities

* [ Prioritizing Snyk issues](/hc/en-us/articles/360009884837-Prioritizing-Snyk-issues)
* [ Why can't I open a Pull Request / Merge Request for issues found by Snyk ?](/hc/en-us/articles/360018829997-Why-can-t-I-open-a-Pull-Request-Merge-Request-for-issues-found-by-Snyk-)
* [ How Snyk finds out about new vulnerabilities](/hc/en-us/articles/360003923877-How-Snyk-finds-out-about-new-vulnerabilities)
* [ What languages do we support Fix Pull Requests or Merge Requests?](/hc/en-us/articles/360003044737-What-languages-do-we-support-Fix-Pull-Requests-or-Merge-Requests-)
* [ Fix your vulnerabilities](/hc/en-us/articles/360003891038-Fix-your-vulnerabilities)
* [ Upgrading package versions to fix vulnerabilities](/hc/en-us/articles/360003891058-Upgrading-package-versions-to-fix-vulnerabilities)
* [ Snyk patches to fix vulnerabilities](/hc/en-us/articles/360003891078-Snyk-patches-to-fix-vulnerabilities)
* [ Introduction to ignoring issues](/hc/en-us/articles/360003891098-Introduction-to-ignoring-issues)
* [ Merge advice](/hc/en-us/articles/360007389537-Merge-advice)

##  How Snyk finds out about new vulnerabilities

Snyk’s security team, based in Israel, maintains our vulnerability database.

This work includes curating vulnerabilities found or reported elsewhere on the web, as well as doing our own research to uncover previously unknown vulnerabilities, which we then responsibly disclose. Snyk Enterprise gets early notifications for issues our research uncovers alongside this responsible disclosure process.

Most of the vulnerabilities in our database originate from one of these sources:

1. _Monitoring other vulnerability databases_, such as CVEs from [NVD](https://nvd.nist.gov/) and many others.
2. _Monitoring user activity on GitHub_, including issues, PRs and commit messages that may indicate a vulnerability.
3. _Bulk research_, using tools that look for repeated security mistakes across open source package code
4. _Manual research_, investing our researchers time to manually audit more widely used packages for security flaws.

For every issue deemed to be a real vulnerability, we assign the right CVSS \(severity\) score and package version specification, create an advisory, and make it available in the product.

