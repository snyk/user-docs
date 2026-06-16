# Resynchronize with connected sources

You can scan a Cloudflare or AWS connection for asset discovery by setting up the respective integration on the **Integrations** page. Visit [Scan a Cloudflare connection for asset discovery](scan-cloudflare-connection-asset-discovery.md) and [Scan an AWS connection for asset discovery](scan-aws-connection-asset-discovery.md) to learn more.

When you have one of these integrations in place, you may notice that changes made in those platforms are not immediately reflected in Snyk API & Web. While Snyk API & Web runs periodic checks to ensure information is up to date, you may want to speed up the synchronization with your connections at any time.

## Manually resynchronize with connected sources

To manually resynchronize with your connected sources:

1. Access the **Domains** page on your Snyk API & Web account by clicking the **Domains** tab that appears under the **Targets** list title.

   <figure><img src="../../../.gitbook/assets/resynchronize-connected-sources-domains-tab.png" alt="Domains tab under Targets list"></figure>

1. On the **Domains** page, click the **Refresh** icon that appears next to the **Add domain** button.

   <figure><img src="../../../.gitbook/assets/resynchronize-connected-sources-refresh-icon.png" alt="Refresh icon next to Add domain button"></figure>

This enables Snyk API & Web to resynchronize with your connected sources, fetch zones (domains) and DNS records from Cloudflare, and import the Route 53 domains set up at AWS. The results load shortly after.
