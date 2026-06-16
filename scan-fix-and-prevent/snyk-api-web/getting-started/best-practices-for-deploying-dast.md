# Best practices for deploying DAST

Learn how to optimize your dynamic application security testing (DAST) deployment and avoid common pitfalls.

## Overview

DAST simulates real-world attacks on your web applications and APIs to identify security vulnerabilities. While DAST provides valuable security insights, it performs invasive scans that can affect application performance and behavior.

During a scan, Snyk API & Web:
- Crawls your application URLs
- Interacts with application elements (filling forms, clicking buttons)
- Performs extensive security tests to identify vulnerabilities

Following these best practices helps you deploy DAST safely and effectively.

## Scan non-production environments

Run scans against development or staging environments, not production systems.

Avoid scanning production back offices where administrators manage content, users, permissions, or transactions. DAST scans can:
- Delete users or create unwanted records
- Degrade application performance due to high request volume
- Inject test data that becomes visible to users and potential attackers

{% hint style="warning" %}
When testing for Cross-site Scripting or SQL Injection vulnerabilities, the scanner attempts to inject data into your application. If a vulnerability exists, this test data can appear in your production environment.
{% endhint %}

Use production-like test environments that include web servers and databases that can be easily restored if needed.

## Use test data that replicates real application behavior

Avoid using real data for security testing.

Using production data in test environments can:
- Expose sensitive information in scan results
- Violate data privacy regulations
- Create data leakage risks

Create dedicated test data in a controlled, isolated environment using a separate test organization or user account. This approach enables thorough testing without putting sensitive information at risk.

## Configure authentication with test accounts

Many applications restrict access to authenticated users only. Configuring authentication allows Snyk API & Web to scan deeper into your application scope and identify more vulnerabilities.

Use dedicated test credentials to prevent mixing test activities with real user data.

To learn more about authentication configuration, see [Configure authentication](../configure-targets/configure-authentication.md).

## Exclude features that trigger external actions

During scans, Snyk API & Web interacts with every discovered element, including forms and buttons. These interactions can:
- Send email messages
- Create support tickets
- Make expensive API calls
- Generate garbage content

Configure your reject list to exclude URLs and features that trigger unwanted actions. In your target settings, add patterns for:
- Email notification endpoints
- Ticketing system integrations
- Payment processing endpoints
- External API calls

For more information, see [Use seeds and reject lists](../configure-targets/use-seeds-and-reject-lists.md).
