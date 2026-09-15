---
description: How to verify domain ownership for Snyk API and Web targets
nav_context: classic
---

# Verify domain ownership

Verify domain ownership to authorize security testing on your domain.

Snyk API & Web performs extensive security tests that can appear as malicious attacks. Domain verification proves that you own the domain and are authorized to run security tests.

## Verification methods

Choose one of the following methods to verify domain ownership:

* **[TXT file](verify-with-txt-file.md)** - Add a `.txt` file to your website's root directory with verification content
* **[DNS TXT record](verify-with-dns-txt.md)** - Add a TXT record to your DNS configuration with the verification token
* **[DNS CNAME record](verify-with-dns-cname.md)** - Add a CNAME record to your DNS configuration pointing to the verification domain
* **[Meta tag](verify-with-meta-tag.md)** - Add a meta tag to your website's HTML with the verification token

Each method provides detailed step-by-step instructions and troubleshooting guidance.

## Scanning before verification

Before domain verification completes, you can only run Lightning scans on Web targets. Lightning scans are superficial scans designed for speed that focus on identifying vulnerabilities related to SSL/TLS, HTTP headers, and cookies.

To configure a Lightning scan profile:

1. Navigate to your target settings.
2. Access the **Profile** tab.
3. Select **Lightning**.

## Get help with verification

If the standard verification methods do not work for your setup, contact Snyk support for assistance with domain verification.
