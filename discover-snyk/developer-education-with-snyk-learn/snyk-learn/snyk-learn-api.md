---
description: How to interact with Snyk Learn programmatically through the Beta Snyk Learn API endpoints in the Snyk REST API
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Snyk Learn API

{% hint style="info" %}
**Release status**

The Snyk Learn API endpoints are [Beta](https://docs.snyk.io/developer-tools/snyk-api/rest-api/about-the-rest-api#versioning) endpoints.
{% endhint %}

The Snyk Learn API endpoints are part of the [Snyk REST API](https://docs.snyk.io/developer-tools/snyk-api/rest-api/about-the-rest-api). Use them to interact with Snyk Learn programmatically.

The Snyk REST API requires authentication. For information about authentication, visit [Authentication for API](https://docs.snyk.io/developer-tools/snyk-api/authentication-for-api).

The following table shows the Snyk Learn API endpoints available for each plan. For more information about plans, visit [Plans and pricing](https://snyk.io/plans/).

<table><thead><tr><th width="294">Plan</th><th>Features</th></tr></thead><tbody><tr><td>Free and Team</td><td><ul><li>View the <a href="https://apidocs.snyk.io/?version=2024-10-15#get-/learn/catalog">content catalog</a>.</li></ul></td></tr><tr><td>Enterprise</td><td><ul><li>Track learning progress for <a href="https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/progress/users">individuals</a> in an Organization.</li><li>Track overall learning progress for your <a href="https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/progress/catalog">Organization</a>.</li></ul></td></tr><tr><td>Learning Management add-on</td><td><ul><li><a href="https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/assignments">Manage assignments</a>.</li></ul></td></tr></tbody></table>
