# Scan a Cloudflare connection for asset discovery

Organizations often lack visibility into all their assets (web applications and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

Scanning a Cloudflare connection for asset discovery involves two steps:

1. Obtain the Cloudflare API Token.
2. Add the Cloudflare connection.

## Obtain the Cloudflare API Token

To add a Cloudflare connection, you need the Cloudflare API Token. To obtain it, follow these steps:

1.  Navigate to your Cloudflare account, click **My Profile**, and then click **API Tokens** to access the **User API Tokens** tab.

    <div data-gb-custom-block data-tag="hint" data-style="info" class="hint hint-info"><p>Account API Tokens will not work in this scenario.</p></div>
2. Click **Create Token** and then click **Get Started** in the **Create Custom Token** configuration.
3. Under the **Permissions** section, choose the **Zone** permission group, the item **Zone**, and the **Read** access permission. Then, add another **Zone** permission with the item **DNS** and the **Read** access permission as well.
4. Under the **Zone Resources** section, select the zones you want to include in the scan. Snyk recommends choosing **All zones** to include current and future zones from your Cloudflare account.
5. Click **Continue to summary** to review the details and then click **Create Token**.
6. After the token is created, you are presented with the token value. Copy the token and store it securely, as you will not be able to view it again. You will need this token value in the next step.

## Add the Cloudflare connection

In Snyk API & Web, add a Cloudflare connection for asset discovery:

1. Select the **Discovery** page and click **Add source** to open the configuration.
2. Select **Connect with Cloudflare** and click **Next**.
3. Paste the Cloudflare API Token (obtained in the previous step) into the **Cloudflare API Token** field and click **Connect**.

After successfully connecting with Cloudflare, Snyk API & Web starts running regular discovery scans automatically on your Cloudflare account. In Snyk API & Web, check the **Discovery** page. Once the asset discovery is finished, all newly found assets are added to the list. At the top of the page, information about the number of newly found assets is displayed. Click on it to filter the list.

If you wish to update or remove Snyk API & Web's connection to Cloudflare, navigate to the **Integrations** page in Snyk API & Web.
