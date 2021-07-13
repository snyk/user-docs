# Choose a provisioning option

* [ Introduction to Snyk Single Sign-On \(SSO\)](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360018025297-Introduction-to-Snyk-Single-Sign-On-SSO-/README.md)
* [ Choose a provisioning option](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360019607318-Choose-a-provisioning-option/README.md)
* [ Set up Snyk Single Sign-On \(SSO\)](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360017753618-Set-up-Snyk-Single-Sign-On-SSO-/README.md)
* [ Identity Provider \(IdP\) migration](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/4402292397969-Identity-Provider-IdP-migration/README.md)

## Choose a provisioning option

Determine how new users in your organization get access to Snyk:

* [Open to all](untitled-302.md)
* [Invitation required](untitled-302.md)
* [Custom](untitled-302.md)

With the Open option, all users have automatic access to Snyk by using SSO to log in. They will have access to all organizations in your selected group. Their accounts will all be provisioned with the same role, with two options:

* The **administrator** role allows all new users to manage group admins, view group reports, and work with organizations within the group.
* The **collaborator** role can access the organizations.

Let Snyk Support know whether new users will have the **administrator** role or the **collaborator** role. The selected role applies for all users.

With the invitation required option, admins can invite users or users can request access to an organization.

There are two ways to invite users to organizations. Invite members from the [Members Settings page](https://support.snyk.io/hc/en-us/articles/360004399238-Invite-and-collaborate-with-team-members) or automate the process using our [API endpoint.](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)

If users who have not been invited use SSO to log in, they will gain access to Snyk, but they will not see any projects until an admin invites them or manually adds them to the organization\(s\). You can show a list of organizations with the appropriate contact person so that new users can [request access](https://support.snyk.io/hc/en-us/articles/360016034417).

**Feature availability**  
This feature is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.

You can provision user accounts with customized rules.

You can configure SSO differently for each of your different Snyk groups. You can also map users to a specific organization and role assignment based on information from the identity provider.

Work with your Customer Success Manager and Snyk Technical Services to prepare for implementing this SSO option.

Read more about the next step, [set up single sign-on](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360017753618/README.md).

