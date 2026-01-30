# Authentication for the JetBrains plugin

To scan your Projects, you must authenticate with Snyk.

Snyk supports the following protocols for authentication:

* OAuth 2.0 (Recommended)
* Personal Access Token
* Snyk API token (Legacy)

{% hint style="warning" %}
Before authenticating, ensure your region is properly set. For more details, see [IDEs URLs](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls).
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (326).png" alt=""><figcaption><p>Authentication methods available in the Snyk plugin in Jetbrains IDEs</p></figcaption></figure>

## Steps to authenticate using the OAuth 2.0 protocol

Follow the next steps to authenticate:

1. After the extension is installed, click the Snyk icon in the navigation bar, then click **Trust project and scan**.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-06-03 at 11.17.13 AM.png" alt=""><figcaption><p>Snyk icon and connect and trust</p></figcaption></figure>

2. A new browser window opens, requiring you to log in to your Snyk account.
3. In the next prompt, the Snyk IDE plugin requests access to act on your behalf. Click **Grant app access**.
4. When you have authenticated successfully, a confirmation message appears. Close the browser window and return to the IDE.

The analysis starts automatically. The IDE reads and saves the authentication tokens on your local machine.

{% hint style="info" %}
OAuth 2.0 tokens are not static and cannot be copied from the Snyk account page.
{% endhint %}

If you have problems, see [OAuth 2.0 authentication does not work](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

## Steps to authenticate using your Personal Access Token

{% include "../../../.gitbook/includes/this-method-is-inferior-to-....md" %}

{% hint style="warning" %}
The Personal Access Token (PAT) authentication is progressively rolled out to all Enterprise customers. To check if this feature is available for your Organization at this time, please reach out to your Snyk account team.
{% endhint %}

To authenticate using the Personal Access token, follow these steps:

1. Navigate to **Settings** > **Tools** > **Snyk**.
2.  Set the **Authentication Method** to **Use Personal Access Token**.

    <figure><img src="../../../.gitbook/assets/image (326) (1).png" alt=""><figcaption></figcaption></figure>
3. Click the **Connect IDE to Snyk** button.
4. Create your **Personal Access** **Token**. For details, see the [Authentication for API](../../../snyk-api/authentication-for-api/) page.
5. Add the token in the **Token** field.
6. Click **Apply and Close.**

## Steps to authenticate using your Snyk API token

{% include "../../../.gitbook/includes/this-method-is-inferior-to-....md" %}

To authenticate, follow these steps:

1. In the JetBrains plugin, navigate to **Settings** > **Tools** > **Snyk**.
2.  Set the **Authentication Method** to **API token**.

    <figure><img src="../../../.gitbook/assets/image (328).png" alt=""><figcaption></figcaption></figure>
3. Click the **Connect IDE to Snyk** button.
4.  Click **Authenticate** in the web browser window that opens.

    <figure><img src="../../../.gitbook/assets/image (317).png" alt=""><figcaption></figcaption></figure>
5. The API token is automatically updated in the **API Token field**.
6. Click **Apply** or **OK.**

The analysis starts automatically.

{% hint style="info" %}
Alternatively, copy the personal API token from your Snyk Web UI instance (default is [https://app.snyk.io](https://app.snyk.io/)). Paste the token in the **API Token** field. For details, see [Obtain and use your Snyk API token](../../../discover-snyk/getting-started/#obtain-and-use-your-snyk-api-token).
{% endhint %}

## How to switch accounts

To re-authenticate with a different account, follow these steps:

1. In the JetBrains plugin, navigate to **Settings** > **Tools** > **Snyk**.
2. Clear the value of the **Token** field.
3. Click **Apply** or **OK**.
4. When you have logged out, start authentication again from the beginning.

## Requirements for Linux and Unix

When authenticating with Snyk, users have the option to copy the authentication URL to their clipboard.

For Linux and Unix users, this requires the `xclip` or `xsel` utility to be installed.
