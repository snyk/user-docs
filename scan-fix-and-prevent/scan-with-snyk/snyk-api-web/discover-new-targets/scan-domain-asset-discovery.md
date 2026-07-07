# Scan a domain for asset discovery

Organizations often lack visibility into all their assets (web applications and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

To scan a domain for asset discovery, follow these steps:

1. Add a domain.
2. Verify the domain.

## Add a domain

In Snyk API & Web, add a domain for asset discovery:

1. Select the **Discovery** page.
2. Click **Add source** to open the configuration modal.
3. Select **Add a domain** and click **Next**.
4. On the next screen, enter the hostname and click **Add**.

## Verify the domain

After adding a domain, you must verify it by following the instructions on the screen to complete the scan. For more information, visit [Verify domain ownership](../configure-targets/verify-domain-ownership/README.md).

After the domain is added and verified, Snyk API & Web starts running regular discovery scans automatically on your account.

In Snyk, check the **Discovery** page. After the asset discovery finishes, Snyk adds all newly found assets to the list. At the top of the page, Snyk displays information about the number of newly found assets. Click it to filter the list.
