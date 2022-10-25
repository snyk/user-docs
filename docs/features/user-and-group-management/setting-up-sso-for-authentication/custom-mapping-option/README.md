# Custom Mapping Option

This option allows you to dynamically assign users to your Snyk Group(s) and Organizations based on data provided by your Identity Provider (IdP) to implement a scaled user provisioning and access model. To understand more about roles and permissions within Snyk, refer to [Managing permissions](../../managing-users-and-permissions/managing-permissions.md). Work with your Snyk account team to implement this option.

See also [Member Roles](../../managing-users-and-permissions/member-roles.md) and [Roles in Custom SSO](../../managing-users-and-permissions/member-roles.md#roles-in-custom-sso).

## Requirements

* Completion of the SSO information worksheet for appropriate IdP (identity provider) found in the Resources section of [Set up Snyk Single Sign-On (SSO)](../set-up-snyk-single-sign-on-sso.md)
* A proper configuration of the custom attributes within your IdP to populate the `roles` array mapping ([Example roles array mapping](./#example-roles-array-mapping))

## Understanding roles array mapping

Within the IdP, you must first pass a custom mapping called `roles` as a string array. Refer to [Example: Setting up custom mapping for Okta](example-setting-up-custom-mapping-for-okta.md). Refer to your IdP documentation on how to configure custom mappings for additional IdP providers.

### How Snyk handles roles array mapping

To configure this option, send the `roles` array within the SAML attributes or OIDC claims to adhere to one of the following patterns:

1\. {prefix}-groupadmin

* This role mapping assigns users with the Group Admin and Org Admin roles.
* **prefix** is present on every role mapping. It is an identifier that allows Snyk to identify which `roles` array values to process. By default this value is **snyk**. If another value is required, work with your Snyk account team. Note: **prefix** must be fully lowercase.
* **groupadmin** configures all users with this role as a Group Admin and Org Admin for all Group(s) that the user is assigned to and all Orgs that fall under the Group(s).

2\. {prefix}-{groupID}

* This role mapping assigns users with the Org Collaborator roles for all Organizations underneath the specified Group(s).
* **groupID** is the ID string for a group in Snyk. This can be found in the snyk URL at the Group level: https://app.snyk.io/group/\<Group ID> or Group dropdown -> ![Settings](../../../../.gitbook/assets/cog\_icon.png) -> General -> Group ID.

3\. {prefix}-{orgslug}-{role}

* This role mapping assigns users with the specified role of Collaborator or Admin for the Snyk Organization specified in **orgslug**.
* **orgslug** is the unique identifier of the Organization name in Snyk.
  * How to find the **orgslug**: https://app.snyk.io/org/{orgslug} OR by using the Snyk [API List all organizations in a group endpoint.](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)
  * **Note**: The **orgslug** is the name of the Organization in most cases; however, this may not always be the case.
  * Note: **orgslug** can be a value of up to 60 characters and must be fully lowercase
* **role** is either **collaborator or admin.**

### Roles array mapping format

To assign users with Group Admin role use the following format:

```
{
    "roles": [
        "{prefix}-groupadmin"
    ]
}
```

To assign users with Org Collaborator roles use the following format:

```
{
    "roles": [
        "{prefix}-{groupID}"
    ]
}
```

To assign users as Org Admin or Org Collaborator use the following format for the roles array. **Note**: You can assign different roles on a per-org basis. The following example assigns a user as Org Admin in the **orgslug** Org but a Collaborator in the **orgslug2** Org.

```
{
    "roles": [
        "{prefix}-{orgslug}-admin",
        "{prefix}-{orgslug2}-collaborator"
    ]
}
```

### Example roles array mapping

The following example shows how to assign roles to Snyk users under the mapping convention.

* The customer is named ABC and has one group called ABC.
* The customer has three Organizations within Snyk: Application-SecurityScanner1, Partner-Plugins, and Application-Payments.
* The customer has four teams: Business Development, Engineering, Security, and Product. Each has different needs:
  * Business Development team needs access to the ABC group and only the Partner-Plugins organization as Org Admin.
  * Engineering needs access to the ABC group, the Application-SecurityScanner1 organization as Org Admin, Partner-Plugins organization as Org Admin, and Application-Payments as Org Collaborator.
  * Security needs access to the ABC group as Group Admin and all three organizations as Org Admin.
  * Product team needs access to the ABC group and all three organizations as Org Collaborator,

For the Business Development Team, Snyk uses the {prefix}-{orgslug}-{role} mapping:

```
{
    "roles": [
        "snyk-partner-plugins-admin"
    ]
}
```

For the Engineering Team, Snyk uses the {prefix}-{orgslug}-{role} mapping:

```
{
    "roles": [
        "snyk-application-securityscanner1-admin",
        "snyk-partner-plugins-admin",
        "snyk-application-payments-collaborator"
    ]
}
```

For the Security Team, Snyk uses the {prefix}-groupadmin mapping:

```
{
    "roles": [
        "snyk-groupadmin"
    ]
}
```

For the Product Team, Snyk uses the {prefix}-{groupID} mapping, where the value of groupID must be inserted;

```
{
    "roles": [
        "snyk-{groupID}"
    ]
}
```

### Summary diagram of roles under custom mapping

<figure><img src="../../../../.gitbook/assets/custom-mapping-screenshot.png" alt="Summary diagram of role under custom mapping"><figcaption><p>Summary diagram of role under custom mapping</p></figcaption></figure>
