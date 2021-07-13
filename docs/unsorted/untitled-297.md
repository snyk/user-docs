# Why can't I open a Pull Request / Merge Request for issues found by Snyk ?

#### [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

### [ ](untitled-297.md) <a id="category-name"></a>

#### Starting to fix vulnerabilities

* [ Prioritizing Snyk issues](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360009884837-Prioritizing-Snyk-issues/README.md)
* [ Why can't I open a Pull Request / Merge Request for issues found by Snyk ?](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360018829997-Why-can-t-I-open-a-Pull-Request-Merge-Request-for-issues-found-by-Snyk-/README.md)
* [ How Snyk finds out about new vulnerabilities](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360003923877-How-Snyk-finds-out-about-new-vulnerabilities/README.md)
* [ What languages do we support Fix Pull Requests or Merge Requests?](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360003044737-What-languages-do-we-support-Fix-Pull-Requests-or-Merge-Requests-/README.md)
* [ Fix your vulnerabilities](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360003891038-Fix-your-vulnerabilities/README.md)
* [ Upgrading package versions to fix vulnerabilities](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360003891058-Upgrading-package-versions-to-fix-vulnerabilities/README.md)
* [ Snyk patches to fix vulnerabilities](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360003891078-Snyk-patches-to-fix-vulnerabilities/README.md)
* [ Introduction to ignoring issues](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360003891098-Introduction-to-ignoring-issues/README.md)
* [ Merge advice](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360007389537-Merge-advice/README.md)
* [Docs Library \| Snyk](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/README.md)
* [Fixing and prioritizing issues](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/categories/360001328418-Fixing-and-prioritizing-issues/README.md)
* [Starting to fix vulnerabilities](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/sections/360001106758-Starting-to-fix-vulnerabilities/README.md)

## Why can't I open a Pull Request / Merge Request for issues found by Snyk ?

You would not necessarily get a Fix PR button for your project whenever you import it from your Source Control tool or run a scan. Snyk only opens PR for direct dependencies.

If your project has transitive dependencies that contains vulnerabilities, Snyk would not open PRs against them.

* 
