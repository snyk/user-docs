# Evaluating and prioritizing vulnerabilities

* [ Evaluating and prioritizing vulnerabilities](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360006113978-Evaluating-and-prioritizing-vulnerabilities/README.md)
* [ Remediate your vulnerabilities](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360006113798-Remediate-your-vulnerabilities/README.md)
* [ Upgrading package versions to fix](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360005993658-Upgrading-package-versions-to-fix/README.md)
* [ Snyk patches to fix](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360004032437-Snyk-patches-to-fix/README.md)
* [ Ignoring issues not prioritized for your project](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360004002718-Ignoring-issues-not-prioritized-for-your-project/README.md)

## Evaluating and prioritizing vulnerabilities

### Prioritize and remediate by exploit maturity

You can filter detected vulnerabilities in your projects according to exploit maturity to see whether a specific vulnerability has an exploit in the wild and if so, how mature that exploit is.

In this way, you can prioritize and attend to the most important and risky vulnerabilities first.

The filter appears as follows:

The **Exploit maturity** filter is available from any detailed **Projects** page, from our **Reports,** and from our **Vulnerabilities DB**. Furthermore, an API is now available.

**Proof of Concept vulnerability patches cannot be disabled** and will appear in fix PRs where they are found.

**Prerequisites:** Projects imported prior to the implementation of his feature cannot be evaluated for exploit maturity. Reimport the project in order to scan for this data.

**Steps:**

1. Log in to Snyk.
2. Go to the detailed **Projects** page for any of your projects
3. Work with and remediate vulnerabilities from the **Issues** tab of the **Reports** area as well:
   1. Filter reports by exploit maturity:
   2. View exploit maturity data from the **Issues** list in **Grouped** mode:
   3. View exploit maturity data from the **Issues** list in **Ungrouped** mode:

