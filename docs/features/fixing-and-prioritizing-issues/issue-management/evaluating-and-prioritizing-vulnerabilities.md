# Evaluating and prioritizing vulnerabilities

**Prioritize and fix by exploit maturity**

You can filter detected vulnerabilities in your projects according to exploit maturity to see whether a specific vulnerability has an exploit in the wild and if so, how mature that exploit is.

This way, you can prioritize and attend to the most important and risky vulnerabilities first.

The filter appears as follows:

![](<../../../.gitbook/assets/image (53) (1).png>)

These values are available:

* **Mature:** Snyk has a published code exploit for this vulnerability.
* **Proof of concept:** Snyk has a proof-of-concept or detailed explanation of how to exploit this vulnerability.
* **No known exploit:** Snyk did not find a proof-of-concept or a published exploit for this vulnerability.
* **No data**: The issue is not a vulnerability (but rather, a license issue or a vulnerability advisory)

![Exploit Maturity in the Vulnerability Card](<../../../.gitbook/assets/image (1) (3) (1).png>)

The **Exploit maturity** filter is available from any detailed **Projects** page, from our **Reports,** and from our **Vulnerabilities DB**. Furthermore, an API is now available.

**Proof of Concept vulnerability patches cannot be disabled** and will appear in fix PRs where they are found.

**Prerequisites:** Projects imported prior to the implementation of this feature cannot be evaluated for exploit maturity. Reimport the project to scan for this data.

**Exploit sources:**

Information about the exploit's existence and status are collected from various sources.

The security analysts at Snyk hand-curated information on new exploits and an automated process that explores structured and unstructured data from multiple exploit sources.

Examples of structured data are the [Known Exploited Vulnerabilities Catalog](https://www.cisa.gov/known-exploited-vulnerabilities-catalog) by CISA (Cybersecurity and Infrastructure Security Agency) [Exploit DB](https://www.exploit-db.com/), and others. Examples of unstructured data are blogs, forums, and social media sites like Twitter.

![Exmpale of Exploit Maturity for CVE-2022-22965](<../../../.gitbook/assets/image (105) (1) (1) (1) (1) (1) (1) (2).png>)

**Steps:**

1. Log in to Snyk.
2. Go to the detailed **Projects** page for any of your projects
3. Work with and fix vulnerabilities from the **Issues** tab of the **Reports** area as well:
   1. Filter reports by exploit maturity
   2. View exploit maturity data from the **Issues** list in **Grouped** mode
   3. View exploit maturity data from the **Issues** list in **Ungrouped** mode
