# Understanding Linux vulnerability severity

* [ Snyk Container security overview](/hc/en-us/articles/360003946897-Snyk-Container-security-overview)
* [ How Snyk Container works](/hc/en-us/articles/360003915918-How-Snyk-Container-works)
* [ Supported operating system distributions](/hc/en-us/articles/360017545417-Supported-operating-system-distributions)
* [ Understanding Linux vulnerability severity](/hc/en-us/articles/360013304357-Understanding-Linux-vulnerability-severity)

##  Understanding Linux vulnerability severity

A single vulnerability can impact different Linux distributions in different ways, giving both a general severity value, and specific severity data based on specific Linux security advisories.

If these separate sets of data are represented by a single **severity** score, this detailed information is lost.

### Relative importance

Relative Importance asserts a common severity for a vulnerability and shows the underlying detailed information for that severity, based on multiple sources. This helps developers and analysts view a common level of importance, and exposes the underlying information that helped form the given severity.

### View relative importance

New information appears in the **Security information** section of the project page, for each issue:

We currently support Ubuntu and Debian. More Linux distributions to follow.

### External information sources

We use the following external sources to provide this information for Debian and Ubuntu:

* [NVD severity](https://nvd.nist.gov/vuln)
*  [Debian urgency score](https://security-team.debian.org/security_tracker.html#severity-levels)
*  [Ubuntu priority score](https://people.canonical.com/~ubuntu-security/cve/priority.html)

