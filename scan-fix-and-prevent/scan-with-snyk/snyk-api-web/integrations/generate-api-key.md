---
description: How to generate an API key for Snyk API and Web
---

# Generate API key

Learn how to generate API keys to integrate with Snyk API & Web.

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
