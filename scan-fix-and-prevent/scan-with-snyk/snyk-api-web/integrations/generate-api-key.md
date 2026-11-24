---
description: Learn how to generate API keys to integrate with Snyk API & Web
nav_context: agnostic
last_reviewed: 2026-08-03
last_reviewed_source: bulk-sync
---

# Generate API key

{% hint style="info" %}
For Snyk-managed accounts, you manage API keys in Snyk API & Web, not in Snyk.

The page lists the keys of one Group. Choose the Group in the bar at the top first.

When adding or editing a key, choose an Organization, then a role, then click **Add**. Add one pair for each Organization the key needs access to. At least one pair is required, and you cannot choose all Organizations in a Group at once.

The roles available come from your Snyk roles, and you cannot scope a key to a single target.

For details, visit [Understand Snyk-managed accounts](../managing-account/snyk-managed-accounts.md).
{% endhint %}

## Overview

Snyk API & Web provides API keys for authentication when you integrate with third-party systems, such as Slack, Jira, Azure Boards, Azure DevOps, or Jenkins. You also use API keys to integrate your own applications with Snyk using the Snyk API & Web API.

Generating an API key has two steps:

1. Access the API key configuration.
2. Configure, generate, and save the API key.

This article describes these steps in detail.

## Step 1: Access the API key configuration

Access the API key configuration as follows:

1. Open the Snyk API & Web application.
2. Navigate to the **Settings** dropdown menu in the bottom-left corner of the navigation bar and select **API Keys**.
3. Click **Add API key** to open the configuration form.

## Step 2: Configure, generate, and save the API key

In the configuration form, fill out the following fields:

1. **Name**: Type a meaningful name for the API key (for example, "Azure DevOps Integration" or "My App Integration with Snyk API & Web").
2. **Role / Scope**: Select the roles that can use the API key and at which level (scope), for example, "Admin" / "Global (account)". Click the plus (**+**) button to add the pairs of roles and scopes.
3. (Optional) **Labels**: Tag the API key with meaningful labels for filtering.

Click **Generate key** and copy and save the API key.

{% hint style="warning" %}
Snyk displays the API key only this one time. Ensure you copy and save it in a secure place.
{% endhint %}

After you save the API key, proceed to integrate with third-party tools or integrate your own applications with Snyk using the Snyk API & Web API.
