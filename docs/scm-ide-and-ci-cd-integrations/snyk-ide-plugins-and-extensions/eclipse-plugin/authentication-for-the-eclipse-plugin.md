# Authentication for the Eclipse plugin

To scan your Projects, you must authenticate with Snyk.&#x20;

Snyk supports the following protocols for authentication:

* OAuth 2.0 (Recommended)
* Personal Access Token
* Snyk API token (Legacy)

{% include "../../../.gitbook/includes/before-authenticating.md" %}

<figure><img src="../../../.gitbook/assets/image (59).png" alt=""><figcaption><p>Authentication methods available in the Snyk plugin in Eclipse</p></figcaption></figure>

## Steps to authenticate using the OAuth 2.0 protocol

After the plugin is installed, follow these steps to authenticate:

1. In the dialog that opens, set the Snyk API endpoint for a custom multi-tenant or single-tenant setup. For details, see [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls).\
   \
   Multi-tenant users who do not belong to the `SNYK-US-01` region ( `https://api.snyk.io`) will be automatically redirected to the correct domain for the email with which the user authenticated. This redirect will not happen if users are expected to use a custom URL, such as in single-tenant company configurations.\
   \
   When you are finished with the settings on this page, click **Next**.

<figure><img src="../../../.gitbook/assets/SCR-20240822-mgxw (1) (1).png" alt="" width="563"><figcaption><p>Snyk endpoint configuration</p></figcaption></figure>

2. On the next page, follow the prompts, then click **Finish**.

<figure><img src="../../../.gitbook/assets/SCR-20240822-mibb (1) (1).png" alt="" width="563"><figcaption><p>Additional information and finish</p></figcaption></figure>

3. A new browser page opens, requiring you to log in to your Snyk account.
4. In the next prompt, the Snyk IDE plugin requests access to act on your behalf. Click **Grant app access**.
5. After you have successfully authenticated, a confirmation message appears. Close the browser window and return to the IDE.

The analysis starts automatically. The IDE reads and saves the authentication tokens on your local machine.&#x20;

{% hint style="info" %}
OAuth 2.0 tokens are not static and cannot be copied from the Snyk account page.
{% endhint %}

If you have problems, see [OAuth 2.0 authentication does not work](../../../cli-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

## Steps to authenticate using your Personal Access Token

{% include "../../../.gitbook/includes/this-method-is-inferior-to-....md" %}

{% hint style="warning" %}
The Personal Access Token (PAT) authentication is progressively rolled out to all Enterprise customers. To check if this feature is available for your Organization at this time, please reach out to your Snyk account team.
{% endhint %}

To authenticate using the Personal Access Token, follow these steps:

1. Navigate to **Eclipse** > **Settings** > **Snyk**. \
   (On Windows/Linux navigate to **Window** > **Preferences** > **Snyk**)
2.  Set the **Authentication Method** to **Use Personal Access Token**.

    <figure><img src="../../../.gitbook/assets/image (56).png" alt=""><figcaption></figcaption></figure>
3. Click the **Connect IDE to Snyk** button.
4. Create your **Personal Access** **Token**. For details, see the [Authentication for API](../../../snyk-api/authentication-for-api/) page.&#x20;
5. Add the token in the **Token** field.
6. Click **Apply and Close.**

The analysis starts automatically.

## Steps to authenticate using your Snyk API token

{% include "../../../.gitbook/includes/this-method-is-inferior-to-....md" %}

To authenticate using the API token, follow these steps:

1. Navigate to **Eclipse** > **Settings** > **Snyk**.\
   (On Windows/Linux navigate to **Window** > **Preferences** > **Snyk**)
2.  Set the **Authentication Method** to **API token**.

    <figure><img src="../../../.gitbook/assets/image (57).png" alt=""><figcaption></figcaption></figure>
3. Click the **Connect IDE to Snyk** button.
4.  Click **Authenticate** in the web browser window that opens.

    <figure><img src="../../../.gitbook/assets/image (47).png" alt=""><figcaption></figcaption></figure>
5. The API token is automatically updated in the **API Token field**.
6. Click **Apply and Close.**

The analysis starts automatically.

{% hint style="info" %}
Alternatively, copy the personal API token from your Snyk Web UI instance (default is [https://app.snyk.io](https://app.snyk.io/)). Paste the token in the **API Token** field.  For details, see [Obtain and use your Snyk API token](../../../getting-started/#obtain-and-use-your-snyk-api-token).
{% endhint %}

## How to switch accounts

To re-authenticate with a different account, follow these steps:

1. Navigate to **Preferences** > **Snyk**.
2. Clear the value in the **Token** field.
3. Click **Apply and Close**.
4. When you have logged out, start authentication again from the beginning.
