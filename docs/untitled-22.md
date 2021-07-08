# Why can't I open a Pull Request / Merge Request for issues found by Snyk ?

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Starting to fix vulnerabilities

* [ Prioritizing Snyk issues](/hc/en-us/articles/360009884837-Prioritizing-Snyk-issues)
* [ Why can't I open a Pull Request / Merge Request for issues found by Snyk ?](/hc/en-us/articles/360018829997-Why-can-t-I-open-a-Pull-Request-Merge-Request-for-issues-found-by-Snyk-)
* [ How Snyk finds out about new vulnerabilities](/hc/en-us/articles/360003923877-How-Snyk-finds-out-about-new-vulnerabilities)
* [ What languages do we support Fix Pull Requests or Merge Requests?](/hc/en-us/articles/360003044737-What-languages-do-we-support-Fix-Pull-Requests-or-Merge-Requests-)
* [ Fix your vulnerabilities](/hc/en-us/articles/360003891038-Fix-your-vulnerabilities)
* [ Upgrading package versions to fix vulnerabilities](/hc/en-us/articles/360003891058-Upgrading-package-versions-to-fix-vulnerabilities)
* [ Snyk patches to fix vulnerabilities](/hc/en-us/articles/360003891078-Snyk-patches-to-fix-vulnerabilities)
* [ Introduction to ignoring issues](/hc/en-us/articles/360003891098-Introduction-to-ignoring-issues)
* [ Merge advice](/hc/en-us/articles/360007389537-Merge-advice)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [Fixing and prioritizing issues](/hc/en-us/categories/360001328418-Fixing-and-prioritizing-issues)
3.  [Starting to fix vulnerabilities](/hc/en-us/sections/360001106758-Starting-to-fix-vulnerabilities)

##  Why can't I open a Pull Request / Merge Request for issues found by Snyk ?

You would not necessarily get a Fix PR button for your project whenever you import it from your Source Control tool or run a scan. Snyk only opens PR for direct dependencies. 

If your project has transitive dependencies that contains vulnerabilities, Snyk would not open PRs against them.

* 
