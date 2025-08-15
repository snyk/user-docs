# Requests for access to an Organization

Users who are not members of a Snyk Organization can request access.

When **Requesting access** is enabled for an Organization, after a user requests access, all administrators of the Organization are notified by email and can send an invitation to the user by email.

After the requester accepts the invitation, the requester has access to the Organization.

## How to request access to an Organization

Users who want to join an Organization in a specific Group can select the **Organizations** option on the Group menu and the **All Organizations** tab. Find the name of the Organization on the list displayed, click **Request**, and respond to the prompts to request access. This notifies the Organization Admins or users with the Organization-level **Invite Users** permission, who can respond to the request.

Additional common routes for users to request this access include the following:

* Using a link shared by someone
* Clicking through to Snyk for a pull request that Snyk has opened
* Clicking through from a status check from Snyk on any pull request in the repository

When **Requesting access** is enabled for an Organization and a non-member reaches a valid URL for a Project or status check result, the user can request access once per Organization in 48 hours.

{% hint style="info" %}
If you need to request access to an Organization you have access to, you may have logged in using an authentication method different from the one you usually use. For more information, see the [support article](https://support.snyk.io/s/article/Unable-to-display-this-organization-The-organization-does-not-exist-or-you-do-not-have-permission-to-access-it) that explains the error and how to proceed.
{% endhint %}

{% hint style="info" %}
Only a valid URL for a Project or status check allows users to request access. Using any other URL results in the error message. This prevents users from requesting access or getting confirmation that Snyk is in use by a company by guessing a URL.
{% endhint %}

## Enable the Request Access setting

In your Organization, navigate to **Settings** > **General** > **Requesting Access**. Select **Enabled** to allow access requests or **Disabled** to stop allowing access requests.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-09-27 at 09.44.52.png" alt=""><figcaption><p>Request Access setting</p></figcaption></figure>
