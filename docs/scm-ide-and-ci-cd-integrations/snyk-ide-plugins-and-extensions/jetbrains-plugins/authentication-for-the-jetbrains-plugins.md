# JetBrains plugin authentication

To scan your Projects you must authenticate with Snyk.&#x20;

{% hint style="info" %}
Before authenticating ensure you configure your Snyk region properly if you use Snyk on the EU and AU tenants. For more information, see [Regional hosting and data residency](../../../working-with-snyk/regional-hosting-and-data-residency.md#cli-and-ci-pipelines-urls).
{% endhint %}

Snyk supports the following protocols for authentication:

* OAuth 2.0 (default)
* Snyk API token (fallback option)

## Steps to authenticate using the OAuth 2.0 protocol

To authenticate follow these steps:

1. After the extension is installed, click the **Snyk icon** in the  navigation bar, then click **Trust project and scan**.

<figure><img src="../../../.gitbook/assets/SCR-20240821-twbu.png" alt="Snyk idon and connect and trust"><figcaption><p>Snyk icon and connect and trust</p></figcaption></figure>

2. The extension opens a new page in a default browser and asks you to log in to your Snyk account:

<figure><img src="../../../.gitbook/assets/SCR-20240821-qogt.png" alt="Snyk login" width="375"><figcaption><p>Snyk login</p></figcaption></figure>

3. The next page asks for your authorization for the IDE to act on your behalf. Click **Grant app access**.

<figure><img src="../../../.gitbook/assets/SCR-20240821-qnpy.png" alt="Grant app access" width="375"><figcaption><p>Grant app access</p></figcaption></figure>

4. After you authenticate successfully, view the confirmation message.

<figure><img src="../../../.gitbook/assets/SCR-20240821-qrgp.png" alt="Successful authentication" width="375"><figcaption><p>Successful authentication</p></figcaption></figure>

5. The IDE reads and saves the authentication on your local machine. Close the browser window and return to the IDE.&#x20;

The analysis starts automatically. If you have problems, see [OAuth 2.0 authentication does not work](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

{% hint style="info" %}
OAuth 2.0 tokens are not static and cannot be copied from Snyk account page.
{% endhint %}

## Steps to authenticate using your Snyk API token

{% hint style="warning" %}
This method is inferior to the OAuth method.
{% endhint %}

To authenticate follow these steps:

1. In the JetBrains plugin navigate to **Preferences** > **Tools** > **Snyk**.
2. Find the **Authentication Method** and change it to **Token authentication**
3. Copy your API token. For details, see [How to obtain your Snyk API token](../../../getting-started/how-to-obtain-and-authenticate-with-your-snyk-api-token.md).
4. Paste or enter the token in the **Token** field.
5. Click **Apply** or **OK**. The analysis starts automatically.

## How to switch accounts

To re-authenticate with a different account, follow these steps:

1. In the JetBrains plugin, navigate to **Preferences** > **Tools** > **Snyk**.
2. Clear the value of the **Token** field.
3. Click **Apply** or **OK**.
4. When you have logged out, start authentication from scratch.

## Requirements for Linux and Unix

When authenticating with Snyk, users have the option to copy the authentication URL to their clipboard.

For Linux and Unix users, this requires that the `xclip` or `xsel` utility be installed.
