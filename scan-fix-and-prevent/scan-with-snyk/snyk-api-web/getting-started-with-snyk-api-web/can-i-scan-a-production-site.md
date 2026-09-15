---
description: Whether you can scan a production site with Snyk API and Web
nav_context: agnostic
---

# Can I scan a production site?

Snyk cannot recommend or assume liability for any damage to your site resulting from a target scan. However, the risk varies significantly depending on your application type and configuration.

{% hint style="warning" %}
The following guidance helps you assess the risk, but you are responsible for determining whether scanning your production environment is appropriate for your use case.
{% endhint %}

## Administrative back offices (high risk)

Do not scan production environments where you can manage users or site content.

Scanning administrative interfaces in production can result in:
- Deleted users
- Unwanted content added to your site
- Modified permissions or configurations
- Corrupted transaction data

While Snyk API & Web blocklists URLs and buttons such as **Delete**, the risk of harmful actions remains significant.

## User-generated content applications (high risk)

Do not scan production environments where users can insert content that other users can view.

Scanning these applications in production results in:
- Attack payloads appearing as random content on your site
- Visible security test data in public-facing areas
- Potential confusion or concern for legitimate users

Examples include social platforms, forums, or content management systems with public-facing output.

## Isolated user or organization scopes (low risk)

Applications where users or organizations do not interact with each other present low to very low risk for production scanning.

In these applications:
- All actions are performed within the scope of a user or organization
- Data added or changed by users is visible only to users from the same entity
- Test accounts can be isolated from real user data

For example, in a customer relationship management (CRM) application where each organization has multiple users, create a testing organization and a testing user account specifically for Snyk. This isolates test data and minimizes risk.

Even in low-risk scenarios, use dedicated test accounts and configure authentication to ensure scans stay within isolated test environments. Visit [Best practices for deploying DAST](best-practices-for-deploying-dast.md) for additional guidance.
