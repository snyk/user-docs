---
description: >-
  Record a login sequence once and map the session credential so Snyk API & Web
  can authenticate API targets at scan time without manual token refresh.
---

# Browser login for APIs

{% hint style="info" %}
Browser login for API targets is in **Early Access**. Contact your account team or [Snyk Support](https://support.snyk.io/s/contactsupport) if you want to try this feature.

This release does not support:

* OpenAPI or GraphQL API targets
* multi-factor authentication (MFA), including time-based one-time passwords (TOTP)
* AI-assisted login sequence configuration
* login video downloads for debugging
{% endhint %}

Snyk API & Web can replay a browser-driven sign-in flow at scan time, extract the resulting session credential, and attach it to your collection requests, so authenticated scans run on a schedule without manual token refresh. Use this when sign-in happens in a browser (for example, OAuth 2.0, OIDC, SAML, or SSO) rather than through a scripted collection folder. For other authentication options, visit [Authentication](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-authentication).

By the end of this guide, you will have recorded a login sequence and mapped one or more session credentials for a Postman or Bruno API target. See [Complete the setup](browser-login-for-apis.md#complete-the-setup) for the remaining steps.

## Prerequisites

{% hint style="info" %}
Your collection must already run end-to-end when the session variable is populated. Browser login obtains the credentials. It does not replace the collection structure or environment configuration.
{% endhint %}

Before you configure browser login for an API target, ensure you have:

* An existing Postman or Bruno API target with collection and environment variables configured. If you are starting from scratch, visit [Configure Postman Collection targets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-api-targets/configure-postman-collection-targets) for Postman or [Configure an API target with a Bruno Collection](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-api-targets/configure-an-api-target-with-a-bruno-collection) for Bruno.
* **Change Target Settings** permission for the target.
* A **login sequence** JSON ready to paste or upload. To record a new sequence, install the [Snyk API & Web Sequence Recorder](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-web-targets/use-sequence-recorder) browser plugin and follow the instructions in [Record a login sequence](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-authentication/configure-login-sequence#record-a-login-sequence). If you have already scanned the same application as a web target and have a working login sequence, you can reuse it.
* A working understanding of where your session credential appears after sign-in (response body, response header, redirect query parameter, cookie, or browser storage).
* Collection requests that reference the variable you map in the **Map session credential** (for example, `{{access_token}}` in Postman or `{{access_token}}` / `bru.getVar("access_token")` in Bruno).

## Add a login sequence

1. Navigate to your Postman or Bruno API target > **Settings** > **Authentication**.
2. Expand **Browser login** > **Record browser sign-in** > **Add login sequence**.
3. Paste or upload your login sequence JSON, then click **Submit sequence**.

For recording, reusing a sequence from a web target, and plugin setup, follow [Login sequence](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-authentication/configure-login-sequence).

## Map session credential

After sign-in, tell Snyk where to find the session credential in login traffic and how your collection requests should use it. Under the **Map session credential**, click **Add token**.

{% stepper %}
{% step %}
### URL to match

Enter the URL of the network request during login that carries the credentials. This is usually:

* The token endpoint response that contains the OAuth 2.0 `access_token`
* The first authenticated request after login for an auth header (for example, `Authorization` or `X-ACCESS-TOKEN`)
* A redirect to another app that returns the token as a query parameter (for example, `?access_key=...`)

Use `*` as a wildcard to match variable parts of the URL. For example:

* `https://*.example.com/oauth/callback*`: matches any callback path or query string on that host
* `https://idp.example.com/oauth/token*`: matches the token endpoint

Regular expressions are not supported. Use `*` only.
{% endstep %}

{% step %}
### Read the token from

Choose where on that matched request the credential lives:

* **Response body**: The token is in a network response body during login. Set **Path** to a dot-notation path to the value in the response body:
  * JSON: keys joined by dots, for example, `data.access_token`
  * XML: element names joined by dots, starting with the root element, for example, `root.user.access_token` (first matching child is used)
  * Plain text: leave Path empty to use the whole body as the value
* **Response header**: The token is returned in a response header on the matched request. Set **Header name** to that header (for example, `Authorization`). Header matching is case-insensitive.
* **Query parameter**: The token appears in the redirect URL after sign-in. Set **Query parameter name** to the parameter that holds it. Common for OAuth 2.0 authorization code flows.
* **Cookie**: A session cookie is set during login. Set **Cookie name** to that cookie (for example, `sessionid`). Set **URL to match** to the domain that sets the cookie.
* **Local storage** or **Session storage**: The app writes the token to browser storage during login. Set **Storage key name** to the key that holds the value, and **URL to match** to the domain associated with that storage entry.

{% hint style="info" %}
If you are unsure which request carries the credential, run **Test configuration** after saving (see [Complete the setup](browser-login-for-apis.md#complete-the-setup)) or inspect network traffic while recording the login sequence.
{% endhint %}
{% endstep %}

{% step %}
### Send on scan requests

In the **Add token** dialog, this section follows **Extract from login traffic**. Configure how scan requests carry the credential:

1. **Where to put the token**: **Header** or **Cookie**
2. **Header name** or **Cookie name**: The name on scan requests (for example, `Authorization` or `session`)
3. **Prefix (optional)**: Text before the token value on headers only (for example, `Bearer` for `Authorization: Bearer <token>`)
{% endstep %}

{% step %}
### Collection variable

Select the variable your Postman or Bruno collection expects to receive the extracted value.

After login, Snyk writes the extracted credential into the **Collection variable** at the environment scope, the same way a collection script would set it with `pm.environment.set()` or `bru.setEnvVar()`. When the scanner runs your collection, any request or script that references that variable gets the live value. This is the bridge between what Snyk extracted from login traffic and the variable names your collection already uses.

{% hint style="warning" %}
A **collection variable** is required for Postman and Bruno targets. The name must match the variable your collection uses. If it does not appear in the dropdown, re-upload or refresh the collection so Snyk can parse variables from `{{variable}}`, `pm.variables.get("variableName")`, or `bru.interpolate("{{variableName}}")` usage.
{% endhint %}

Example: An OAuth flow extracts `access_token` from a token endpoint response body, sets **Where to put the token** to header `Authorization` with prefix `Bearer`, and selects `access_token` as the **Collection variable** (because requests use `Authorization: Bearer {{access_token}}`).

Click **Save token**. You can add multiple tokens if your API needs more than one credential from the same login sequence.
{% endstep %}

{% step %}
### Save and enable

Click **Save** on the authentication page. A **Configured** badge appears when both a login sequence and at least one valid token exist. Enable **Browser login** flow.
{% endstep %}
{% endstepper %}

## Complete the setup

Complete the remaining steps in your target setup:

1. **Test configuration**: Use **Test configuration** in the page header to verify your setup before running a full scan. Visit the [Test target configuration](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/test-target-configuration).
2. **Logout detection** (optional): If your session can expire during a long scan, configure [Logout detection](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/configure-targets/configure-authentication/configure-logout-detection).
3. **Stop scan on login failure** (optional): On the **Authentication** tab, enable **Stop scan on login failure** if you want the scan to stop immediately when login fails instead of continuing without a valid session.
4. Start a scan: Run an authenticated scan from your target. Visit [Start scanning](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web/start-scanning).

## Disable Browser login

To stop using browser login on an API target, turn off **Browser login** flow on the **Authentication** tab. The change applies immediately. You do not need to click **Save**.

Snyk keeps your login sequence and tokens while the feature is off, so you can turn it on again without recording or adding them.

To remove a login sequence or token permanently, delete it on the **Authentication** tab.
