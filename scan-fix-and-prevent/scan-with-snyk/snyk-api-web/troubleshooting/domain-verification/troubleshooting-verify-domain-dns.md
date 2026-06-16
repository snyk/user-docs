# Troubleshoot domain verification with DNS

In order for Snyk API & Web to run full-fledged scans on your target, you need to verify its domain. Visit [Verify domain ownership](../../configure-targets/verify-domain-ownership/) to learn more about why domain verification is required.

## The problem

Domain verification using CNAME or TXT records in your DNS fails with the following errors:

* DNS (CNAME) error: `The DNS query name does not exist: <your CNAME record>`
* DNS (TXT) error: `Token not found`

## Troubleshoot the problem

To troubleshoot this problem, go through the following steps to identify the possible causes and respective solutions to fix it.

### Check the TTL value

Check the Time to Live (TTL) as follows:

1. Go to [Google Admin Toolbox Dig](https://toolbox.googleapps.com/apps/dig/) and type the target's URL.
2. Depending on your domain verification method, click on **CNAME** or **TXT** to see the TTL value.

If the TTL value is high, the DNS configuration with the CNAME or TXT record has not been propagated yet, and the domain verification fails.

| Cause                                                                           | Solution                                                                                                                                          |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| The DNS configuration with the CNAME or TXT record has not been propagated yet. | Wait more time for the DNS configuration to propagate, or go to your authoritative DNS server (for example, Cloudflare) and reduce the TTL value. |

After following these steps, identifying the causes, and applying the respective solutions, you should be able to verify your domain using DNS.

## Related articles

* [Verify domain ownership](../../configure-targets/verify-domain-ownership/)
* [Verify domain with DNS CNAME record](../../configure-targets/verify-domain-ownership/verify-with-dns-cname.md)
* [Verify domain with DNS TXT record](../../configure-targets/verify-domain-ownership/verify-with-dns-txt.md)
