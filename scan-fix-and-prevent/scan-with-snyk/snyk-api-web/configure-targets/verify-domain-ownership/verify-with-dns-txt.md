---
description: How to verify domain ownership with a DNS TXT record for Snyk API and Web
---

# Verify domain with DNS TXT record

Verify domain ownership by adding a TXT record to your DNS configuration.

This verification method proves you have control over the DNS records for your domain without making changes to your website.

## Get verification details

1. In Snyk API & Web, navigate to the **Targets** page.
1. Click the **Domains** tab.
1. Locate your domain and click **Verify**.
1. Select **DNS (TXT)** as the verification method.
1. Note the **CONTENT** value provided.
1. Keep the verification dialog open.

## Create the DNS TXT record

1. Access your DNS management console (provided by your domain registrar or DNS hosting service).
1. Add a new TXT record with these details:
   - **Name/Host**: Your domain's FQDN (for example, `example.com`)
   - **Type**: TXT
   - **Value/Content**: The content value from Snyk API & Web

## Complete verification

1. Return to the verification dialog in Snyk API & Web.
1. Click **Verify**.

## Troubleshooting

### DNS propagation delay

DNS records can take time to propagate across all DNS servers. If verification fails immediately:
- Wait up to 48 hours for full propagation (typically much faster)
- Use a DNS lookup tool to check if the TXT record is visible
- Try verification again after waiting

### CNAME limitations

If your domain FQDN is a CNAME record, this verification method does not work because of DNS limitations. CNAMEs cannot have TXT records.

Use one of these alternative methods instead:
- [TXT file verification](verify-with-txt-file.md)
- [DNS CNAME verification](verify-with-dns-cname.md)
- [Meta tag verification](verify-with-meta-tag.md)
