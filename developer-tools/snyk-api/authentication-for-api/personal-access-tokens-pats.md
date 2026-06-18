# Personal Access Tokens (PATs)

## What are PATs?

A Personal Access Token (PAT) is a credential used to authenticate API, CLI, and IDE workflows.

At any given time, you can have up to three active PATs. This allows for downtime-free rotation, where you can create the new token before the old one expires. You can name PATs and choose their expiration up to a maximum of 90 days ahead. Names help you identify tokens and have no functional impact. You cannot create PATs that live for more than 90 days.

Snyk reveals PATs only upon creation, never after this point. You must safely store the token itself, because the only option afterward is to revoke and create another.

You can create PATs only through the UI, but you can revoke them through the UI and API.

You can see the name, created date, and expiry for each PAT, but not the token, in the UI and through the API.

## Generating a new token (UI only)

1. In your account settings area, select the **Personal access tokens** tab.

<figure><img src="../../.gitbook/assets/unknown (20).png" alt=""><figcaption></figcaption></figure>

2. In the **Name** field, enter a name for your PAT. Names help you distinguish between your PATs and serve no other purpose.
3. To set an expiry date, select the **Expiry** dropdown and select a date for your PAT to expire. The maximum date you can pick is 90 days ahead. You cannot create a PAT with a longer expiry.

<figure><img src="../../.gitbook/assets/unknown (21).png" alt=""><figcaption></figcaption></figure>

4. Click **Generate new token**. Snyk generates your PAT and a window appears with your token. Click **Copy** to copy your token. This is the only time the Snyk Platform shares this token with you. You must store it in a secure and accessible location, because you cannot retrieve it later.

<figure><img src="../../.gitbook/assets/unknown (22).png" alt=""><figcaption></figcaption></figure>

5. If you already have three PATs, you cannot create another one. You must revoke one of them before you can create any more.

## Revoking a token&#x20;

1. In your account settings area, select the **Personal access tokens** tab.
2. Next to the token you want to revoke, click **Revoke**. Snyk revokes your token. You cannot revert token revocation.

<figure><img src="../../.gitbook/assets/unknown (23).png" alt=""><figcaption></figcaption></figure>
