# Team Lead role template

A Team Lead is equivalent to the default [**Organization Admin**](../pre-defined-roles.md) rol&#x65;**,** providing additional capabilities to customize or change user permissions for other users.\
\
Sample permissions include the ability to **Mark as Successful** in Git repositories for pull request and merge request checks, with the ignore capability disabled.

This Group-level role has all Organization-level permissions plus the following additional Group level permissions:

## Group-level permissions

To create this role, enable the following permissions in the relevant categories:

### Group Management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Groups</td><td>true</td></tr><tr><td>Edit Group details</td><td>true</td></tr><tr><td>View Group settings</td><td>true</td></tr><tr><td>Edit settings</td><td>false</td></tr><tr><td>View Group notification settings</td><td>false</td></tr><tr><td>Edit Group notification settings</td><td>false</td></tr></tbody></table>

### Role management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled</th></tr></thead><tbody><tr><td>Read roles</td><td>true</td></tr><tr><td>Create roles</td><td>true</td></tr><tr><td>Edit roles</td><td>true</td></tr><tr><td>Remove roles</td><td>true</td></tr></tbody></table>

### User management

<table><thead><tr><th>Permission</th><th data-type="checkbox"></th></tr></thead><tbody><tr><td>View users</td><td>true</td></tr><tr><td>Add users to the Group</td><td>true</td></tr><tr><td>Edit users in the Group</td><td>true</td></tr><tr><td>Remove users</td><td>true</td></tr><tr><td>Delete users</td><td>true</td></tr><tr><td>Provision users</td><td>true</td></tr><tr><td>Assign and unassign roles</td><td>true</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Organization management
* Snyk Essentials management
* Audit Log management
* IaC settings management
* Insights management
* Issue management
* Reports management
* Request access management
* Security and licence policies
* Service account management
* Snyk Apps management
* Snyk Preview management
* SSO settings management
* Tags management

## Organization-level permissions

This role should have all Organization-level permissions enabled.

You can set this quickly using the **Enable all** button in the Organization-level permissions section. Be sure to update the permissions using the button at the bottom of the section.

<figure><img src="../../../.gitbook/assets/image (150).png" alt=""><figcaption><p>Organization-level permissions section</p></figcaption></figure>
