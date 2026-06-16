# Troubleshooting domain verification with meta tag

In order for Snyk API & Web to run full-fledged scans, you need to verify your domains. Visit [Verify domain ownership](../../configure-targets/verify-domain-ownership/verify-domain-ownership.md) to learn more about why domain verification is required.

## The problem

Domain verification using a meta tag fails with the error: `Meta tag not found.`

<figure><img src="/.gitbook/assets/troubleshooting-verify-domain-meta-tag-error.png" alt="Error message showing Meta tag not found during domain verification"></figure>

## Troubleshoot the problem

To troubleshoot this problem, go through the following steps to identify the possible causes and respective solutions to fix it.

### Check the index page source

Test if the meta tag is present in the index page source in one of the following ways:

* Use the browser:
  1. Type the URL to the index page: `https://<domain_to_verify>`, for example: `https://example.com`.
  1. Right-click on the page and select **View Page Source**.
* Use a curl command like this: `curl -i https://<domain_to_verify>`, for example: `curl -i https://example.com`.

Analyze the HTML of the index page source, check the following possible causes, and apply the respective solution:

| Cause | Solution |
|-------|----------|
| The meta tag is not on the index page. | Review the implementation to add the meta tag before the closing `</head>` tag on the index page. |
| The meta tag is on the index page, but not inside the `<head>` tag. | Review the implementation to have the meta tag before the closing `</head>` tag on the index page. |
| The meta tag is on the index page inside the `<head>` tag, but it is being injected using JavaScript (for example, a Single-Page Application). | Review the implementation to add the meta tag before the closing `</head>` tag on the index page without using JavaScript. |

After following these steps, identifying the causes, and applying the respective solutions, you should be able to verify your domain using a meta tag.

## Related articles

* [Verify domain ownership](../../configure-targets/verify-domain-ownership/verify-domain-ownership.md)
* [Verify domain with meta tag](../../configure-targets/verify-domain-ownership/verify-with-meta-tag.md)
