# Request access to an Organization

Users who are not members of a Snyk Organization can request access. Administrators are notified of a new request by email and can invite the user to join the Organization.

Common routes for users to request this access include the following:

* Using a link shared by someone
* Clicking through to Snyk for a pull request that Snyk has opened
* Clicking through from a status check from Snyk on any pull request in the repository

{% hint style="info" %}
If you need to request access to an Organization you have access to, you may have logged in using an authentication method different from the one you usually use. For more information, see the [support article](https://support.snyk.io/hc/en-us/articles/360001649558-Unable-to-display-this-organization) that explains the error and how to proceed.
{% endhint %}

## How requesting access to an Organization works

The following explains how a user requests access and the request is processed.

When **Requesting access** is enabled for an Organization and a non-member reaches a valid URL for a Project or status check result, the user can request access once per Organization in 48 hours.

The user browses to an address for a Snyk Organization and sees the message **Unable to display this organization. The organization does not exist, or you do not have permission to access it.**

{% hint style="info" %}
Only a valid URL for a Project or status check allows users to request access. Using any other URL results in the error message. This prevents users from requesting access or getting confirmation that Snyk is in use by a company by guessing a URL.
{% endhint %}

After a user requests access, all administrators of the Organization are notified by email and can send an invitation to the user by email.

After the requester accepts the invitation, the requester has access to the Organization.

## Enable the Request Access setting

In your Organization, navigate to the  **Settings** > **General** option and then to the **Requesting Access** section. Select **Enabled** or **Disabled** to stop allowing access requests.

<figure><img src="../../.gitbook/assets/Screenshot 2022-09-27 at 09.44.52.png" alt="Request Access setting"><figcaption><p>Request Access setting</p></figcaption></figure>
