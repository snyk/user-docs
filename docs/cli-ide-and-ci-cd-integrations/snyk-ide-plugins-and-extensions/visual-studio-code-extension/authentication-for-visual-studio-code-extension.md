# Visual Studio Code extension authentication

To scan your Projects, you must authenticate with Snyk.&#x20;

Snyk supports the following protocols for authentication:

* OAuth 2.0 (default)
* Snyk API token (fallback option)

For both methods, Snyk uses the [Secret Storage API](https://code.visualstudio.com/api/references/vscode-api#SecretStorage) to store the token securely. This storage uses the system's keychain to manage the token.

{% hint style="warning" %}
Before authenticating, be sure you have set your region properly. For details, see [IDEs URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls).
{% endhint %}

## Steps to authenticate using the OAuth 2.0 protocol

To authenticate, follow these steps:

1.  After the extension is installed, click the **Snyk Icon** in the navigation bar, and then click **Connect & Trust Workspace**:



    <figure><img src="../../../.gitbook/assets/SCR-20240821-qmuv.png" alt="" width="359"><figcaption><p>Connect and trust workspace</p></figcaption></figure>
2. A new browser window opens; in response to the prompt, log in to your Snyk account.
3. In the next prompt, the Snyk IDE extension requests access to act on your behalf; click **Grant app access**.
4. When you have authenticated successfully, a confirmation message appears; close the browser window and return to the IDE.
5. The IDE reads and saves the authentication on your local machine. Close the browser window and return to the IDE.

The analysis starts automatically. The IDE reads and saves the authentication on your local machine.

{% hint style="info" %}
OAuth 2.0 tokens are not static and cannot be copied from the Snyk account page.
{% endhint %}

If you have problems, see [OAuth 2.0 authentication does not work](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

## Steps to authenticate using your Snyk API token

{% hint style="warning" %}
This method is inferior to the OAuth method.
{% endhint %}

To authenticate, follow these steps:

1. After the extension is installed, click the **Snyk Icon** in the navigation bar; then click the **Settings** icon, find **Authentication Method,** and change it to **Token authentication**:

<figure><img src="../../../.gitbook/assets/SCR-20240821-tarb.png" alt=""><figcaption><p>Change authentication method </p></figcaption></figure>

2. Copy your API token. For details, see [Obtain and use your Snyk API token](../../../getting-started/#obtain-and-use-your-snyk-api-token).
3. Run the`Snyk: Set Token command` and paste the token in the text field.

<figure><img src="../../../.gitbook/assets/image (224) (1) (1) (1) (1) (1) (1) (1) (1).png" alt=""><figcaption><p>Set token manually</p></figcaption></figure>

## How to switch accounts

To re-authenticate with a different account, follow these steps:

1. Run the provided `Snyk: Log Out` command.

<figure><img src="../../../.gitbook/assets/logging-out-command (1).png" alt=""><figcaption><p>Snyk: Log out</p></figcaption></figure>

2. When you have logged out, start authentication again from the beginning.

## Requirements for Linux and Unix

When authenticating with Snyk, users have the option to copy the authentication URL to their clipboard.

For Linux and Unix users, this requires the `xclip` or `xsel` utility to be installed.
