# Verify domain with meta tag

Verify domain ownership by adding a meta tag to your website's HTML.

This verification method is quick and straightforward for websites where you can modify the HTML source code.

## Get the meta tag

1. In Snyk API & Web, navigate to the **Targets** page.
1. Click the **Domains** tab.
1. Locate your domain and click **Verify**.
1. Select **Meta Tag** as the verification method.
1. Copy the **META TAG** content provided.
1. Keep the verification dialog open.

## Add the meta tag to your website

1. Access your website's source code.
1. Open the index page (typically `index.html` or your homepage template).
1. Locate the `<head>` section.
1. Paste the meta tag before the closing `</head>` tag:

```html
<head>
  <!-- other head content -->
  <meta name="snyk-site-verification" content="your-verification-token" />
</head>
```

5. Save and deploy the changes to your live website.

## Complete verification

1. Confirm the meta tag is visible in your website's source code by viewing the page source in a browser.
1. Return to the verification dialog in Snyk API & Web.
1. Click **Verify**.

After successful verification, you can remove the meta tag if desired, though it is harmless to leave it in place.

## Troubleshooting

### Meta tag not found

If verification fails, check that:
- The meta tag is on the index/homepage of your domain
- The meta tag is inside the `<head>` section
- There are no typos in the meta tag content
- The changes have been deployed to your live website (not just saved locally)
- Caching is not serving an old version of the page

