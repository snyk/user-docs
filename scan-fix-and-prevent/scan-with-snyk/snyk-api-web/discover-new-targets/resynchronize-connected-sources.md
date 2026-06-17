# Resynchronize with connected sources

You can scan a Cloudflare or AWS connection for asset discovery by setting up the respective integration on the **Integrations** page. Visit [Scan a Cloudflare connection for asset discovery](scan-cloudflare-connection-asset-discovery.md) and [Scan an AWS connection for asset discovery](scan-aws-connection-asset-discovery.md) to learn more.

When you have one of these integrations in place, changes made in those platforms are not immediately reflected in Snyk API & Web. Snyk runs periodic checks to keep information up to date, but you can speed up the synchronization with your connections at any time.

## Manually resynchronize with connected sources

To manually resynchronize with your connected sources:

1. Access the **Domains** page on your Snyk account by clicking the **Domains** tab that appears under the **Targets** list title.
2. On the **Domains** page, click the **Refresh** icon that appears next to the **Add domain** button.

Snyk then resynchronizes with your connected sources, fetches zones (domains) and DNS records from Cloudflare, and imports the Route 53 domains set up at AWS. The results load shortly after.
