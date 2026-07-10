---
description: How to configure extra hosts for Snyk API and Web web targets
---

# Configure extra hosts

Include additional domains in the scan scope when your web application uses multiple hostnames, such as separate domains for front-end and back-end APIs.

## What are extra hosts?

When you add a target, Snyk API & Web scans only pages under that target URL. The scanner does not scan pages from different hostnames unless you add them as extra hosts.

This default behavior can present challenges for single-page applications (SPAs) where the front-end is built with JavaScript and the server-side functionality is invoked using an API. In SPAs, APIs are typically hosted under a different subdomain (such as `api.example.com`) while the front-end resides on `app.example.com`. In some cases, the API can be an entirely separate domain.

When you configure extra hosts, Snyk follows and scans any `XMLHttpRequest` performed to the specified hostnames. If a hostname is in the extra hosts list, the scanner regards it as within the scanning scope and analyzes the API responses for vulnerabilities.

## When to use extra hosts

Use extra hosts when your web application architecture includes:

* Separate API domains: front-end on `app.example.com` accessing an API on `api.example.com`.
* Multiple subdomains: different subdomains serving various application functions.
* Content delivery networks: assets served from different domains that need security testing.

If Snyk scans only the front-end domain, it can miss critical security tests in the API that drives much of your application's functionality. Adding extra hosts ensures complete coverage of your application's security surface.

## Add extra hosts

To add extra hosts to your target:

1. Navigate to the **Targets** list.
2. Find your target in the list and click the **gear icon** to access the target settings.
3. Navigate to the **Extra Hosts** tab.
4. In the **Add hostname** field, enter the hostname you want to include in the scan scope.
5. Click **Add**.

You can add multiple extra hosts to a single target. The scanner tests each configured hostname during scans.

## Verify domain ownership

After adding an extra host, ensure that its hostname is under a verified domain. Snyk API & Web requires domain ownership verification to prevent unauthorized scanning of domains you do not control.

Visit [Verify domain ownership](../verify-domain-ownership/) to learn how to verify ownership of your domains.

By including extra hosts and scanning both the front-end and associated APIs, you gain a complete picture of your web application's security posture.
