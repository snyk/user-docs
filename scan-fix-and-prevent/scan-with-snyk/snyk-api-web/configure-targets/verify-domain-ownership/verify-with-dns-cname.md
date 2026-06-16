# Verify domain with DNS CNAME record

Verify domain ownership by adding a CNAME record to your DNS configuration.

This verification method proves you have control over the DNS records for your domain without making changes to your website. Use this method when your domain FQDN is a CNAME and TXT record verification is not available.

## Get verification details

1. In Snyk API & Web, navigate to the **Targets** page.
1. Click the **Domains** tab.
1. Locate your domain and click **Verify**.
1. Select **DNS (CNAME)** as the verification method.
1. Note both the **NAME** and **VALUE** provided.
1. Keep the verification dialog open.

## Create the DNS CNAME record

1. Access your DNS management console (provided by your domain registrar or DNS hosting service).
1. Add a new CNAME record with these details:
   - **Name/Host**: The name value from Snyk API & Web
   - **Type**: CNAME
   - **Value/Points to**: The value from Snyk API & Web (verification domain)

## Complete verification

1. Return to the verification dialog in Snyk API & Web.
1. Click **Verify**.

## Troubleshooting

### DNS propagation delay

DNS records can take time to propagate across all DNS servers. If verification fails immediately:
- Wait up to 48 hours for full propagation (typically much faster)
- Use a DNS lookup tool to check if the CNAME record is visible
- Try verification again after waiting

If you encounter CNAME restrictions, use an alternative verification method:
- [TXT file verification](verify-with-txt-file.md)
- [Meta tag verification](verify-with-meta-tag.md)
