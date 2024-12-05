# Custom mapping

Custom mapping allows you to dynamically assign users to your Snyk Groups and Organizations based on data provided by your Identity Provider (IdP), in order to implement a scaled user provisioning and access model.

{% hint style="info" %}
Work with your Snyk account team to implement this option.
{% endhint %}

To understand more about roles and permissions within Snyk, see [Pre-defined roles](../../../snyk-admin/user-roles/pre-defined-roles.md). See also [user role management](../../../snyk-admin/user-roles/user-role-management.md).

## Requirements for custom mapping

* Complete the SSO information worksheet for the appropriate IdP (identity provider) found in [Resources for SSO setup](../set-up-snyk-single-sign-on-sso.md#resources-for-sso-setup).
* Properly configure the custom attributes in your IdP to populate the `roles` array mapping. See [Example: roles array mapping](./#example-roles-array-mapping).

## Custom Mapping options

Snyk offers an updated custom mapping option explained on this page, with increased flexibility, including the ability to grant users Group-level custom roles as well as pre-defined roles.&#x20;

The Snyk [Legacy custom mapping](legacy-custom-mapping.md) option is still supported.

## Roles array mapping with Snyk

In the IdP, you must first pass a custom mapping called `roles` as a string array. [Examples](examples-setting-up-custom-mapping-for-idps/) of how to set this up for different IdPs are provided.

Refer to your identity provider documentation for further information on how to configure custom mappings.

## Custom mapping assertions

This section documents the role assertions that Snyk expects in order to map correctly to Snyk roles within Snyk Groups and Organizations.

### Role assertion format

Role assertions should be provided to Snyk in the following format:

`snyk:{scope}:{target}:{role}`

Where:

`snyk` is a fixed prefix for role mapping. **Required**.

`scope` can be one of `org` or `group.`**Required**; if a role mapping does not contain a valid scope, it will be ignored.

`target` must be a **slug** of either `org` or `group` where the role will be granted, or a **wildcard**.

* See the [Slugs](./#slugs) section for details on how to find this information.
* **Optional**; may be an asterisk `*` or empty string `::`to apply as a [wildcard](./#wildcards) for all resources within the defined `scope` that are associated with the SSO connection.

`role` is the normalized name of the required role. See [Role normalized name](./#role-normalized-name) to find this information.

* **Required**; if no role is present, the role mapping is ignored.
* If the role is a custom role, that is, a role created in the Group Member Roles admin panel that can be of either `Org` or `Group` type,  then it must have a `custom:` prefix. See the [example role assertions](./#example-role-assertions).
* Built-in roles do not have the `custom:` prefix, so values like `org_admin`, `org_collaborator`, `group_viewer` will refer to the Snyk pre-defined roles, which are shown with a padlock symbol in the Member Roles page.

{% hint style="warning" %}
Users must only have one role mapped per Organization or Group. Mapping multiple roles except when using wildcards is not supported and can lead to unexpected behavior.
{% endhint %}

{% hint style="info" %}
Any user that is granted a role in an Organization within the SSO without an explicit Group-level role in the role assertion, will also be implicitly assigned the **Group Member** Group-level role for that Group. This is the pre-defined Group-level role with the fewest permissions and ensures that the user becomes a member of the Group.
{% endhint %}

### Example role assertions

* `snyk:group:*:group_admin` Assign the user the pre-defined **Group Admin** role for all groups associated with the SSO connection.
* `snyk:group::custom:sys_admin` Assign the user the custom Group-level role `Sys Admin` for all groups associated with the SSO connection.
  * Note that `::` here indicates an empty string for the target, and so is treated as a wildcard in the preceding example.
  * Note that this Group-level custom role must be created manually before it can be assigned.
* `snyk:org:my-default-org:org_admin` Assign the user the pre-defined **Organization Admin** Organization-level role for the Organization `my-default-org`.
* `snyk:org:my-default-org:custom:code_editor`   Assign the user the custom Org-level role `Code Editor` for the Organization `my-default-org` .

### Example role assertions array

An example of a set of role assertions for a user follows:

<pre><code><strong>{
</strong><strong>    "roles": [
</strong>        "snyk:group:*:group_viewer",
        "snyk:org:development:org_admin",
        "snyk:org:test-org-N58YhztauHcaMiNfvi5fbL:custom:developer_readonly",
    ]
}
</code></pre>

{% hint style="info" %}
The system also supports comma-separated lists of roles instead of an array.

```
{
  "roles": "snyk:group:*:group_viewer, snyk:org:development:org_admin, 
snyk:org:test-org-N58YhztauHcaMiNfvi5fbL:custom:developer_readonly"
}
```
{% endhint %}

These assertions will assign the user:

* The pre-defined Group-level role **Group Viewer** for all groups in the SSO. See [pre-defined roles](../../../snyk-admin/user-roles/pre-defined-roles.md) for the permission this grants
* The pre-defined Organization-level role **Organization Admin** for the Organization with the name **Development**.
* The custom Organization-level role **Developer ReadOnly** for the Organization with the name **Test Org**, which has the slug `test-org-N58YhztauHcaMiNfvi5fbL`.

## Wildcards

Custom mapping introduces wildcards, which allow one assertion to assign a user the same role in all Organizations or Groups.

Assertions using wildcards take a lower priority than assertions with a specific target. This means that it is possible, for example, to assign a user a default role for all Organizations, and specific roles in others:

```
roles: [
        "snyk:org:*:custom:developer_readonly",
        "snyk:org:development:org_admin",
    ]
```

These role assertions will:

* Grant the user the pre-defined Organization-level role **Organization Admin** in the **Development** Organization.
* Grant the user the custom Organization-level role **Developer ReadOnly** on all other organizations within the SSO connection.
* Grant the user the pre-defined Group-level role **Group Member** on all groups in the SSO connection. For more details, see the note that follows.

{% hint style="info" %}
Any user that is granted a role in an Organization within the SSO without an explicit Group-level role in the role assertion, will also be implicitly assigned the **Group Member** Group-level role for that Group. This is the pre-defined Group-level role with the fewest permissions and ensures that the user becomes a member of the Group.
{% endhint %}

## Slugs

For a valid role assertion, the Organization or Group slug may be required, where a wildcard is not used. The slug is the canonical name for the Organization or Group within Snyk.

To find an Organization slug, navigate to the **Settings** page for the Organization, and under **General** settings, the Organization slug value is visible. This can then be copied and used in role assertions in custom mapping.&#x20;

<figure><img src="../../../.gitbook/assets/settings_org_slug.png" alt="Organization general settings page, showing the Organization slug"><figcaption><p>Organization general settings page, showing the Organization slug</p></figcaption></figure>

To find the slug of a Group, navigate to the Group Settings, and find the Group slug under General Settings, which you can copy.

<figure><img src="../../../.gitbook/assets/image (374).png" alt="Group general settings page, showing the Group slug"><figcaption><p>Group general settings page, showing the Group slug</p></figcaption></figure>

## Role normalized name

To find the normalized name of a role for use in custom mapping, first confirm that the role exists for the Snyk Group by navigating to it in the Group settings: **Group Settings** > **Member Roles** > {**Role**}.&#x20;

This will open the role details page that shows which permissions are enabled for the role and also shows the normalized name. Copy this normalized name and use it in custom mapping.

For more details on roles, and specifically, custom roles, see [user role management](../../../snyk-admin/user-roles/user-role-management.md).

<figure><img src="../../../.gitbook/assets/image (376).png" alt="Role details page for Organization Admin role"><figcaption><p>Role details page for Organization Admin role</p></figcaption></figure>

## Pre-defined role slugs

Snyk has a set of [pre-defined roles](../../../snyk-admin/user-roles/pre-defined-roles.md). Their corresponding normalized names are listed below.

| Role Type    | Role Name        | Role Slug          |
| ------------ | ---------------- | ------------------ |
| Organization | Org Admin        | `org_admin`        |
| Organization | Org Collaborator | `org_collaborator` |
| Group        | Group Admin      | `group_admin`      |
| Group        | Group Viewer     | `group_viewer`     |
| Group        | Group Member     | `group_member`     |
