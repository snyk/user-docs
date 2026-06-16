# Generate an API key

Learn how to generate API keys to integrate with Snyk API & Web.

## Overview

Snyk API & Web provides API keys for authentication purposes in the integration with third-party systems (for example, Slack, Jira, Azure Boards, Azure DevOps, or Jenkins), as well as the integration of your own applications with Snyk API & Web using the Snyk API & Web API.

Generating an API key has two steps:

1. Access the API key configuration.
1. Configure, generate, and save the API key.

This article describes these steps in detail.

## Step 1: Access the API key configuration

Start by accessing the API key configuration as follows:

1. Go to the Snyk API & Web application.
1. Go to the **Settings** dropdown menu at the bottom-left corner of the navigation bar and click **API Keys**.

<figure><img src="../../../.gitbook/assets/how-to-generate-an-api-key_image1.png" alt="Settings dropdown menu with API Keys option highlighted"></figure>

1. Click **Add API key** to open the configuration form.

## Step 2: Configure, generate, and save the API key

In the configuration form, fill out the following fields:

<figure><img src="../../../.gitbook/assets/how-to-generate-an-api-key_image2.png" alt="API key configuration form"></figure>

1. **Name**: Type a meaningful name for the API key (for example, "Azure DevOps Integration" or "My App Integration with Snyk API & Web").
1. **Role / Scope**: Choose the roles that can use the API key and at which level (scope) (for example, "Admin" / "Global (account)"). Click the plus (**+**) button to add the pairs of roles and scopes.
1. (Optional) **Labels**: Tag the API key with meaningful labels for filtering purposes.

Click **Generate key** and copy and save the API key.

<figure><img src="../../../.gitbook/assets/how-to-generate-an-api-key_image3.png" alt="Generated API key displayed"></figure>

{% hint style="warning" %}
This is the only time the API key is displayed. Make sure you copy and save it in a secure place.
{% endhint %}

With the API key saved, proceed to the integration with third-party tools or the integration of your own applications with Snyk API & Web using the Snyk API & Web API.
