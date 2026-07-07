# Identify scanner requests

When running target scans, you can use the information in the user-agent header to identify HTTP requests from Snyk API & Web.

## User-agent header format

Here is an example of a user-agent header:

```
Mozilla/5.0 (compatible; +https://probely.com/sos) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36 ProbelySPDR/0.2.0
```

The header contains a specific string: `https://probely.com/sos`.

Use this string to identify Snyk API & Web requests. Do not rely on any other information in the header for identification.

## Related information

You can also identify Snyk requests through the scanner IP address. For more information, visit [Scanner IP address](scanner-ip-address.md).
