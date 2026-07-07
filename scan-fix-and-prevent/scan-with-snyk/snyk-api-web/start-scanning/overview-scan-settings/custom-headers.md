# Custom headers

When performing a dynamic application security testing (DAST) scan, your scanner acts like an automated user, navigating your site and testing for vulnerabilities. However, modern security infrastructures (such as Web Application Firewalls) or complex authentication requirements can block these automated requests unless they are properly identified.

This page explains why you need custom headers and how to set them up in Snyk API & Web.

## What are custom headers used for?

In a DAST scan, custom headers serve three primary purposes:

1. **Identifying requests** - Because all requests include the defined custom headers, you can use them for many purposes, such as bypassing Web Application Firewalls. By adding a unique custom header (for example, `X-SnykApiWeb-Scan: true`), you can tell your Web Application Firewall to allow the scanner's traffic while still protecting the site from malicious bots.
2. **Authentication and authorization** - Some APIs require specific headers for every request to prove identity, such as `Authorization: Bearer <token>` or custom API keys.
3. **Environment signaling** - You can use headers to trigger specific behaviors in your application during a scan, such as disabling certain checks like CAPTCHAs or preventing the app from sending real emails during the test.

## Configure custom headers

In Snyk API & Web, you can configure custom headers at the **Target** level. This ensures that every request sent by the crawler and scanner to that specific application includes your required headers.

### Step-by-step configuration

1. Log in to Snyk API & Web.
2. From the Targets list, locate the target you want to configure and click the **gear icon** to navigate to the respective Settings.
3. Click the **Scanner** tab of the Target settings.
4. Locate the **CUSTOM HEADERS** module:
   * Enter the custom header name (for example, `X-Scan-Origin`).
   * Enter the corresponding value (for example, `Snyk-DAST-Scanner`).
   * Choose whether to test this header during the scan. When the check box is clear (default), the scanner sends the header as-is with every request but does not test it for vulnerabilities. When the check box is selected, the scanner also treats the header value as an attack surface and runs security checks against it.
5. Click **Add** to save the custom header.

You can also set up static custom headers for authentication purposes. Visit the [authentication documentation](../../configure-targets/configure-authentication/README.md) for more details.
