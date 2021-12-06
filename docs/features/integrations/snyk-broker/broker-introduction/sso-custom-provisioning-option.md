# SSO Custom Provisioning Option

The Custom provisioning option allows for customers to map users to Snyk organization and role assignments based on the customer's Identity Provider (IdP) settings. Work with your Customer Success Manager to implement this SSO option.

### Requirements:

* Completion of the SSO information worksheet for appropriate IdP located [here](../../../user-and-group-management/setting-up-sso-for-authentication/set-up-snyk-single-sign-on-sso.md)
* A proper configuration of the custom attributes within the customer's IdP for configuring roles array mapping (example below)

### Example use case of this provisioning option:&#x20;

* Customer has users configured to be part of a **security** group within their IdP. In this scenario, they may want to give access to all **orgs** and make them **group admin.**
* If there are users in the **finance group**, they may only want to give access to the finance org as **collaborators.**&#x20;

### Understanding the roles array mapping:

Organization mappings are to be passed as an array field called roles which must be a string array. In this array, we expect each mapping to adhere to the following pattern:

**PREFIX-ORGNAME-ROLE**

* **PREFIX** is present on every role mapping. It should be agreed upon between the customer and Snyk and will be used to identify the customer's role mapping.
  * Example: snyk-CUSTOMERNAME (With no spaces)
* **ORGNAME** is the organization name in Snyk and must be the slugified version of the org name with no spaces (The slugified org name can be found in the snyk URL within an organization)
  * Example: https://app.snyk.io/org/\<slugified org name>
* **ROLE** is the role a user has in that org or group. It is either **collaborator**, or **admin**, or **groupadmin.** (Avoid using any custom roles here)

### Example roles array mapping:

To assign roles to Snyk users under the mapping convention&#x20;

1. Apply the attribute name as: `snyk-{group_name}-{orgName}-{role}`
2. Use the following structure:

```
{
"roles": [
"snyk-{group_name}-{my-org}-admin",
"snyk-{group_name}-{my-other-org}-collaborator"
]
}
```

To assign users with group admin privileges&#x20;

1. Apply the attribute name as:  `snyk-{group_name}-groupadmin`
2. Use the following structure:

```
{
"roles": [
  "snyk-{group_name}-groupadmin"
]
}
```

### Understanding the Custom provisioning option:

* This option will allow customers to use their existing IdP groups and settings to enable a scaled deployment of Snyk users.&#x20;
* Snyk requires the use an agreed naming convention with the customer for the organization/role assignments.
* This option only modifies the SSO configuration within Snyk. It does not modify or impact the customer's IdP settings.&#x20;
* This option auto assigns users to orgs with given role. It does not prevent customers from using invitations or assigning users to organizations manually within Snyk.
* Currently, there are only 3 roles a Snyk user can be assigned: Group admin, Admin, or Collaborator. This will still be the case with this SSO provisioning option.

&#x20;
