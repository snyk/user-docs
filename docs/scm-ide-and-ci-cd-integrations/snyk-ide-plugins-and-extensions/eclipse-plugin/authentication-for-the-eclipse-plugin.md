# Authentication for the Eclipse plugin

To scan your Projects, you must authenticate with Snyk.&#x20;

Snyk supports the following protocols for authentication:

* OAuth 2.0 (default)
* Snyk API token (fallback option)

## Steps to authenticate using the OAuth 2.0 protocol

After the extension is installed, to authenticate, follow these steps:

1. In the dialog that opens, set the Snyk API endpoint for a custom multi-tenant or single-tenant setup. For details, see [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls).\
   \
   Multi-tenant users who do not belong to the `SNYK-US-01` region ( `https://api.snyk.io`) will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.\
   \
   When you are finished with the settings on this page, click **Next**.

<figure><img src="../../../.gitbook/assets/SCR-20240822-mgxw (1) (1).png" alt="" width="563"><figcaption><p>Snyk endpoint configuration</p></figcaption></figure>

2. On the next page, follow the prompts and read the information; then click **Finish**.

<figure><img src="../../../.gitbook/assets/SCR-20240822-mibb (1) (1).png" alt="" width="563"><figcaption><p>Additional information and finish</p></figcaption></figure>

3. The extension opens a new browser page; in response to the prompt, log in to your Snyk account.
4. In the next prompt, the Snyk IDE plugin requests access to act on your behalf; click **Grant app access**.
5. When you have authenticated successfully, a confirmation message appears; close the browser window and return to the IDE.

The analysis starts automatically. The IDE reads and saves the authentication tokens on your local machine.&#x20;

{% hint style="info" %}
OAuth 2.0 tokens are not static and cannot be copied from the Snyk account page.
{% endhint %}

If you have problems, see [OAuth 2.0 authentication does not work](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

## Steps to authenticate using your Snyk API token

{% hint style="warning" %}
This method is inferior to the OAuth method.
{% endhint %}

To authenticate using the API token, follow these steps:

1. Navigate to **Preferences** > **Snyk**.
2. Set the flag to **Use token authentication**.
3. Copy your API token. For details, see [Obtain and use your Snyk API token](../../../getting-started/#obtain-and-use-your-snyk-api-token).
4. Be sure to use your personal token. Paste or enter the token in the **Token** field.
5. Click **Apply and Close.**

The analysis starts automatically.

## How to switch accounts

To re-authenticate with a different account, follow these steps:

1. Navigate to **Preferences** > **Snyk**.
2. Clear the value in the **Token** field
3. Click **Apply and Close**.
4. When you have logged out, start authentication again from the beginning.
