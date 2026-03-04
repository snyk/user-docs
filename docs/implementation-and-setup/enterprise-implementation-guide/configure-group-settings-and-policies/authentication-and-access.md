# Authentication and access

## Set up SSO

{% stepper %}
{% step %}
### Configure SSO settings at the Group level

{% hint style="success" %}
**Key decision:** Choose between **Open to all** or **Require an invite** based on your security policy and license management needs.
{% endhint %}

Snyk recommends using Self-Serve Single Sign-On to establish a SAML connection with your identity provider (IdP). To do this:

1. In the Snyk web UI, navigate to Group **Settings** > **SSO**. If the SSO option is missing, verify your license or contact Snyk Support.
2. Configure your SAML connection, valid email domains, and default permissions.
3. Contact Snyk Support if you want to use other protocols such as OIDC and ADFS.

{% hint style="info" %}
Any identity provider is supported, including tools such as Entra ID, OKTA, and Google Workspace.
{% endhint %}
{% endstep %}

{% step %}
### Manage user accounts

{% hint style="success" %}
**Key decision**: Identify which administrators require Group Admin status and ensure they transition from personal accounts to SSO accounts before the general rollout.
{% endhint %}

To ensure a clean transition to Enterprise authentication, manage your administrative accounts before the wider rollout:

1. **Promote SSO accounts:** Create SSO-linked accounts for your Snyk administrators and grant them the Administrator role.
2. **Remove personal accounts**: Remove any non-SSO or personal accounts used during the pilot phase to ensure all access is governed by your IdP.
3. **Map custom roles**: If you use custom roles, configure your IdP to pass these roles to Snyk.

{% hint style="info" %}
Custom mapping requires Snyk professional services. Contact your account team for assistance.
{% endhint %}
{% endstep %}

{% step %}
### (Optional) Provision users with the API

{% hint style="success" %}
**Key decision**: Determine if you need to pre-allocate users to specific Organizations and roles before their first login to prevent broad default access.
{% endhint %}

If you need to define specific permissions before users first log in, use the Snyk API. This allows you to:

* Assign a specific role to each user.
* Grant access to specific Organizations automatically.
* Control the user footprint from day one.
{% endstep %}
{% endstepper %}

## Set pre-defined user roles

Determine if Snyk pre-defined roles meet your requirements or if you must create custom roles. Snyk uses role-based access control (RBAC) to manage permissions across the Tenant, Group, and Organization levels.

{% stepper %}
{% step %}
### Review Tenant-level roles

{% hint style="success" %}
**Key decision:** Determine if you need centralized oversight across multiple Groups to manage analytics and global user membership.
{% endhint %}

The Tenant is the highest level of the Snyk hierarchy, and it encompasses all Groups and Organizations. Tenant-level roles are essential for large enterprises that require cross-group reporting and user management.

* **Tenant Admin:** Can manage users across the entire Tenant, assign roles, and remove members.
* **Tenant Viewer:** Provides read-only access to Tenant-level features like Snyk Analytics.
* **Tenant Member**: Allows access to the Tenant level but requires specific Group or Organization permissions to take action.

{% hint style="info" %}
Features like Snyk Analytics are available only on Enterprise plans. You can switch between Tenants by selecting the Tenant name in the navigation menu.
{% endhint %}
{% endstep %}

{% step %}
### Review Group-level roles

{% hint style="success" %}
**Key decision:** Determine if your team leads can operate with the fixed permissions of an Organization Admin or if they require a restricted custom role.
{% endhint %}

Pre-defined roles at these levels have fixed permissions that cannot be modified.

#### Group-level roles

* **Group Admin**: Provides full permissions at the Group and Organization levels. Assign this role to Snyk administrators.
* **Group Member**: Allows access to the Group but requires specific Organization-level permissions to interact with Projects.
* **Group Viewer**: Provides read-only access to the Group level. Use this to audit Group settings without making changes.
{% endstep %}
{% endstepper %}

## Set custom roles

{% hint style="success" %}
**Key decision**: Identify if you need a "middle-ground" role, such as a Senior Developer who can, for example, override checks but cannot delete Projects.
{% endhint %}

Large enterprises often require more granular control. Because you cannot change permissions for pre-defined roles, use custom roles to limit or expand specific actions and gain:

* **Customization**: Assign a specific set of permissions to a role name you define.
* **Flexibility**: Update permissions for a custom role at any time. Changes apply immediately to all assigned users.
* **Efficiency**: Use custom role templates provided by Snyk to jumpstart your configuration.
