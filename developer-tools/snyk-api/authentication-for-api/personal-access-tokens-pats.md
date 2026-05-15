# Personal Access Tokens (PATs)

## What are PATs?

A Personal Access Token (PAT) is a credential used to authenticate API, CLI, and IDE workflows.

At any given time, you can have up to three active PATs. This is primarily to allow for downtime-free rotation journeys, where the new token can be created ahead of the old one expiring. PATs can be named, and you can choose their expiration up to a maximum of 90 days ahead. Names are solely there to help you identify tokens; they have no functional impact. Users cannot create PATs that live for more than 90 days.

PATs are revealed only upon creation, and never after this point. Users must ensure they safely store the token itself, as the only option after this will be to revoke and create another.

PATs can only be created through the UI, but can be revoked through the UI and API.

You can see the name, created date and expiry for each token of a PAT (but not the token), and these details can be seen in the UI and through the API.

## Generating a new token (UI Only)

1. In your account settings area, select the **Personal access tokens** tab.

<figure><img src="../../.gitbook/assets/unknown (20).png" alt=""><figcaption></figcaption></figure>

2. In the Name field, enter the name that you would like to provide your PAT. The names are solely there to help you distinguish between your PATs, and provide no other purpose.
3. To set an expiry date, select the Expiry drop-down and select a date for your PAT to expire. The maximum date you can pick is 90 days ahead. No PAT can be created with a longer expiry.

<figure><img src="../../.gitbook/assets/unknown (21).png" alt=""><figcaption></figcaption></figure>

4. Click **Generate new token**. You have generated your PAT and will see a window appear with your token. Click **Copy** to copy your token. This is the only time the Snyk Platform will share this token with you. You must store it in a secure and accessible location as you cannot retrieve it later.

<figure><img src="../../.gitbook/assets/unknown (22).png" alt=""><figcaption></figcaption></figure>

5. If you have three PATs already, you will not be able to create another one. You must revoke one of them before being able to create any more.

## Revoking a token&#x20;

1. In your account settings area, select the **Personal access tokens** tab.
2. Next to the token that you would like to revoke, click Revoke. This results in your token being revoked. Token revocation cannot be reverted.

<figure><img src="../../.gitbook/assets/unknown (23).png" alt=""><figcaption></figcaption></figure>
