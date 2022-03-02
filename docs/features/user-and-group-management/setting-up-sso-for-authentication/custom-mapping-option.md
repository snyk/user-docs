# Custom Mapping Option

This option allows you to dynamically assign users to your Snyk group(s) and organizations based on data provided by your Identity Provider (IdP) to implement a scaled user provisioning and access model. To understand more about roles and permissions within Snyk, refer to [this](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions) article. Work with your Customer Success Manager to implement this option.

### Requirements

* Completion of the SSO information worksheet for appropriate IdP located [here](set-up-snyk-single-sign-on-sso.md)
* A proper configuration of the custom attributes within your IdP to populate the `roles` array mapping ([example](custom-mapping-option.md#example-roles-array-mapping) below)

### Understanding roles array mapping

Within the IdP, you must first pass a custom mapping called `roles` as a string array. [Here is an example document on configuring custom mapping on Okta](https://docs.snyk.io/features/user-and-group-management/setting-up-sso-for-authentication/example-setting-up-custom-mapping-for-okta). Refer to your IdP documentation on how to configure custom mappings for additional IdP providers.  &#x20;

### How Snyk handles roles array mapping

To configure this option, we expect you to send the `roles` array within the SAML attributes or OIDC claims to adhere to one of the following patterns:

**1. {prefix}-groupadmin**

* This role mapping will assign users with the Group Admin and Org Admin roles
* **prefix** is present on every role mapping. It is an identifier that allows Snyk to identify which `roles` array values to process. By default, we expect this value to be **snyk**. If another value is required, work with your Customer Success Manager.
  * Note: **prefix** must be fully lowercase
* **groupadmin** will configure all users with this role as a Group Admin and Org Admin for all group(s) that the user is assigned to and all orgs that fall under the group(s)

**2.  {prefix}-{groupID}**

* This role mapping will assign users with the Org Collaborator roles for all organizations underneath the specified group(s)
* **groupID** is the ID string for a group in Snyk (This can be found in the snyk URL at the group level)
  * How to find Group ID: https://app.snyk.io/group/\<Group ID> or Group dropdown -> ![](../../../.gitbook/assets/cog\_icon.png) -> General -> Group ID

**3.** **{prefix}-{orgslug}-{role}**

* This role mapping will assign users with the specified role of Collaborator or Admin for the Snyk organization specified in **orgslug**.&#x20;
* **orgslug** is the unique identifier of the organization name in Snyk&#x20;
  * How to find the **orgslug**: https://app.snyk.io/org/{orgslug} OR via our [API ](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)(Note: While the **orgslug** will be the name of the organization for most cases, this may not always be the case)
  * Note: **orgslug** can be a value of up to 60 characters and must be fully lowercase&#x20;
* **role** is either **collaborator or admin**

### Roles array mapping format

To assign users with Group Admin role use the following format

```
{
    "roles": [
        "{prefix}-groupadmin"
    ]
}
```

To assign users with Org Collaborator roles use the following format

```
{
    "roles": [
        "{prefix}-{groupID}"
    ]
}
```

To assign users as Org Admin or Org Collaborator use the following format for the roles array\
\
Note: You can assign different roles at a per-org basis. Example below will assign a user as Org Admin in the **orgslug** org but a Collaborator in the **orgslug2** org

```
{
    "roles": [
        "{prefix}-{orgslug}-admin",
        "{prefix}-{orgslug2}-collaborator"
    ]
}
```

### Example roles array mapping

The following example will show how to assign roles to Snyk users under the mapping convention.

* The customer is named ABC and has one group called ABC
* The customer has 3 organizations within Snyk: Application-SecurityScanner1, Partner-Plugins, and Application-Payments
* The customer has 4 teams: Business Development, Engineering, Security, and Product that have different needs:
  * Business Development team will need access to the ABC group and only the Partner-Plugins organization as Org Admin
  * Engineering will need access to the ABC group, the Application-SecurityScanner1 organization as Org Admin**,** Partner-Plugins organization as Org Admin, and Application-Payments as Org Collaborator
  * Security will need access to the ABC group as Group Admin and all 3 organizations as Org Admin
  * Product team will need access to the ABC group and all 3 organizations as Org Collaborator

For the Business Development Team, we will be using the {prefix}-{orgslug}-{role} mapping

```
{
    "roles": [
        "snyk-partner-plugins-admin"
    ]
}
```

For Engineering Team, we will be using the {prefix}-{orgslug}-{role} mapping

```
{
    "roles": [
        "snyk-application-securityscanner1-admin",
        "snyk-partner-plugins-admin",
        "snyk-application-payments-collaborator"
    ]
}
```

For Security Team, we will be using the {prefix}-groupadmin mapping

```
{
    "roles": [
        "snyk-groupadmin"
    ]
}
```

For Product Team, we will be using the {prefix}-{groupID} mapping, where the value of groupID would need to be inserted

```
{
    "roles": [
        "snyk-{groupID}"
    ]
}
```

### Summary Diagram of Roles under Custom Mapping

![](../../../.gitbook/assets/custom-mapping-screenshot.png)

