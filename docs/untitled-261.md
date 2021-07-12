# Organization access requests

* [ Managing permissions](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360006548637-Managing-permissions/README.md)
* [ Organization access requests](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360016034417-Organization-access-requests/README.md)
* [ Session length](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004008358-Session-length/README.md)
* [ Switch between Snyk organizations](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003915618-Switch-between-Snyk-organizations/README.md)
* [ Audit logs](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004133117-Audit-logs/README.md)

## Organization access requests

Organization access requests allows a user to request access to a Snyk org that they're not a member of \(and so don't have access to\). Administrators will be notified of the new request by email, and can choose to invite the user to join the organization.

The most common routes for a user to reach the address of an org they're not a member of include:

* Someone sharing a link
* Clicking through to Snyk for a pull request Snyk has raised
* Clicking through from a status check from Snyk on any pull request in the repository

**💡 If you find yourself having to request access to an org you have access too, it may be because you've logged in with a different authentication method than you usually use.** [**Find out more**](https://support.snyk.io/hc/en-us/articles/360001649558-Unable-to-display-this-organization)

### How it works:

Previously, when a user browses to any address for a Snyk organization, they are presented with a page informing them that they "either do not have access or the organization does not exist".

When request access is enabled for an organization and a non-member reaches a valid URL for a project or status check result, the user will be given the chance to request access.

**💡 Only a valid URL for a project or status check result will allow them to request access. Any other URL will continue to inform them that they "either do not have access or the organization does not exist". This protects users requesting access \(or even just getting confirmation that Snyk is in use by a company\) by guessing a URL.**

A user can only request access once per organization in 48 hours.

Once a user requests access, all administrators of the org will be notified by email, and can opt to send an invitation to the user.

Once the requester accepts the invitation \(sent to them by email\), they will have access to the org.

### Settings

This can be enabled/disabled at either the Snyk organization level, or for an entire group.

Go to the **Members** page in either the org or group settings.

