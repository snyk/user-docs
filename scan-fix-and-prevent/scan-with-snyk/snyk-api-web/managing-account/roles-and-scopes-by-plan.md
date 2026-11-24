---
description: Learn what type of roles and scopes you can set up for your users and API keys
nav_context: agnostic
last_reviewed: 2026-06-16
last_reviewed_source: bulk-sync
---

# Roles and scopes by plan

The roles and scopes you can assign your users and API keys differ depending on your account plan:

* With the **Free** plan, users and API keys you add are automatically assigned the **SecOps** role in the context of the **Global (account)** scope.
* With the **Enterprise** plan, you can select a **built-in role** for your users and API keys or create your own **custom roles**, based on selected permissions. As for the scope, it can be **Global (account)**, one or more **Teams**, or one or more **Targets**.

{% hint style="info" %}
For Snyk-managed accounts, roles come from Snyk and apply to your Tenant, a Group, or an Organization. You cannot scope a role to a single target, for a user or for an API key.
{% endhint %}
