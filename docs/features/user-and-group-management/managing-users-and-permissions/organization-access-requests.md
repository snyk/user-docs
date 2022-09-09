# Organization access requests

Organization access requests allow users to request access to a Snyk Organization of which they are not a member. Administrators are notified of the new request by email, and can invite the user to join the Organization.

The most common routes for users to request this access include:

* Someone sharing a link.
* Clicking through to Snyk for a pull request Snyk has raised.
* Clicking through from a status check from Snyk on any pull request in the repository.

{% hint style="info" %}
If you find yourself having to request access to an org you have access to, it may be because you logged in with a different authentication method than you usually use. [Find out more](https://support.snyk.io/hc/en-us/articles/360001649558-Unable-to-display-this-organization).
{% endhint %}

## How it works

A user browses to an address for a Snyk organization and sees a message that they "either do not have access or the organization does not exist".

When request access is enabled for an organization and a non-member reaches a valid URL for a project or status check result, the user can request access.

{% hint style="info" %}
Only a valid URL for a project or status check allows a user to request access. Any other URL continues to inform the user that they "either do not have access or the organization does not exist". This prevents users from requesting access or getting confirmation that Snyk is in use by a company by guessing a URL.
{% endhint %}

A user can only request access once per Organization in 48 hours.

After a user requests access, all administrators of the org are notified by email, and can opt to send an invitation to the user by email.

After the requester accepts the invitation, they have access to the Organization.

## Settings

Request access can be enabled or disabled at either the Snyk Organization level or for an entire Group, on the **Members** page in either the org or group settings.
