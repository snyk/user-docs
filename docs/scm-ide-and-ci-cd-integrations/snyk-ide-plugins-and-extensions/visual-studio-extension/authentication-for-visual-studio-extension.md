# Authentication for Visual Studio extension

To scan your Projects, you must authenticate with Snyk.&#x20;

Snyk supports the following protocols for authentication:

* OAuth 2.0 (default)
* Snyk API token (fallback option)

{% hint style="warning" %}
Before authenticating, be sure you have set your region properly. For details, see [IDEs URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#ides-urls).
{% endhint %}

## Steps to authenticate using the OAuth 2.0 protocol

To authenticate, follow these steps:

1. After the extension is installed, navigate to **Extensions** > **Snyk** > **Windows**, and then **Snyk** to open the Snyk panel.&#x20;

<figure><img src="../../../.gitbook/assets/SCR-20240822-llxy.png" alt="" width="563"><figcaption><p>Snyk extension navigation</p></figcaption></figure>

2. On the welcome screen, click **Trust project and scan.**

<figure><img src="../../../.gitbook/assets/SCR-20240822-lmdw.png" alt="" width="563"><figcaption><p>Trust project and scan</p></figcaption></figure>

3. A new browser window opens; in response to the prompt, log in to your Snyk account.
4. In the next prompt, the Snyk IDE extension requests access to act on your behalf; click **Grant app access**.
5. When you have authenticated successfully, a confirmation message appears; close the browser window and return to the IDE.

The analysis starts automatically. The IDE reads and saves the authentication on your local machine.&#x20;

{% hint style="info" %}
OAuth 2.0 tokens are not static and cannot be copied from the Snyk account page.
{% endhint %}

If you have problems, see [OAuth 2.0 authentication does not work](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

## Steps to authenticate using your Snyk API token

{% hint style="warning" %}
This method is inferior to the OAuth method.
{% endhint %}

To authenticate, follow these steps:

1. After the extension is installed, navigate to **Extensions > Snyk** > **Settings**:

<figure><img src="../../../.gitbook/assets/SCR-20240822-lyzs.png" alt="" width="375"><figcaption><p>Snyk Settings navigation</p></figcaption></figure>

2. Find the **Authentication Method** and change it to **Token authentication**
3. Copy your API token. For details, see [Obtain and use your Snyk API token](../../../getting-started/#obtain-and-use-your-snyk-api-token).
4. Be sure to use your personal token. Paste or enter the token in the **Token** field.
5. Click **OK.**&#x20;

## How to switch accounts

To re-authenticate with a different account, follow these steps:

1. Navigate to **Extensions > Snyk > Settings.**
2. Clear the value of the **Token** field.
3. Click **OK**.
4. When you have logged out, start authentication again from the beginning.
