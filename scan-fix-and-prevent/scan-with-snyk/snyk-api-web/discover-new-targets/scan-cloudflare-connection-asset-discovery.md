---
description: How to scan a Cloudflare connection for asset discovery in Snyk API and Web
nav_context: classic
---

# Scan a Cloudflare connection for asset discovery

Organizations often lack visibility into all their assets (web applications and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

Scanning a Cloudflare connection for asset discovery involves two steps:

1. Obtain the Cloudflare API Token.
2. Add the Cloudflare connection.

## Obtain the Cloudflare API Token

To add a Cloudflare connection, you need the Cloudflare API Token. To obtain it, follow these steps:

1.  Navigate to your Cloudflare account, click **My Profile**, and then click **API Tokens** to access the **User API Tokens** tab.

    <div data-gb-custom-block data-tag="hint" data-style="info" class="hint hint-info"><p>Account API Tokens do not work in this scenario.</p></div>
2. Click **Create Token** and then click **Get Started** in the **Create Custom Token** configuration.
3. Under the **Permissions** section, select the **Zone** permission group, the item **Zone**, and the **Read** access permission. Then, add another **Zone** permission with the item **DNS** and the **Read** access permission as well.
4. Under the **Zone Resources** section, select the zones you want to include in the scan. Snyk recommends selecting **All zones** to include current and future zones from your Cloudflare account.
5. Click **Continue to summary** to review the details and then click **Create Token**.
6. After Cloudflare creates the token, it displays the token value. Copy the token and store it securely, because you cannot view it again. You need this token value in the next step.

## Add the Cloudflare connection

In Snyk API & Web, add a Cloudflare connection for asset discovery:

1. Select the **Discovery** page and click **Add source** to open the configuration.
2. Select **Connect with Cloudflare** and click **Next**.
3. Paste the Cloudflare API Token (obtained in the previous step) into the **Cloudflare API Token** field and click **Connect**.

After successfully connecting with Cloudflare, Snyk starts running regular discovery scans automatically on your Cloudflare account. In Snyk, check the **Discovery** page. After the asset discovery finishes, Snyk adds all newly found assets to the list. At the top of the page, Snyk displays information about the number of newly found assets. Click it to filter the list.

To update or remove the Snyk connection to Cloudflare, navigate to the **Integrations** page in Snyk.
