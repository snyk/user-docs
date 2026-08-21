---
description: How to migrate Snyk SSO custom role mapping from the legacy format to v2 colon-delimited syntax
nav_context: classic
---

# Migrating from legacy to v2 custom mapping

Single Sign-On (SSO) custom mapping dynamically provisions users into Snyk Groups and Organizations based on attributes passed from your Identity Provider (IdP).

If your company still uses the legacy format, the updated custom mapping gives you a standardized syntax, wildcard support, and streamlined assignments for custom roles.

For details on the legacy and v2 formats, visit [Legacy custom mapping](legacy-custom-mapping.md) and [Custom mapping](README.md).

## Benefits of v2 custom mapping

The v2 format uses an extensible, colon-delimited string syntax:

`snyk:{scope}:{target}:{role}`

This gives you:

* Support for multiple Groups
* Claims-based provisioning and deprovisioning of access
* Support for Tenant-level roles

## Migration IdP configuration process

### Step 1: Plan and audit

* Review your current IdP configuration (Okta, Entra ID, Google Workspace).
* Document existing legacy role strings.
* Ensure all users have appropriate role mappings configured in the IdP before activation.
* Maintain fallback access during the transition by leaving legacy mapping assignments as is within the IdP.
* Coordinate timing with internal stakeholders to minimize disruption—for example, work with security or compliance teams, as role changes during audits could temporarily affect access reporting or user permissions.

**Note:** These migration activities apply to the SSO connection directly and affect all Snyk Groups with which the connection is associated. All Groups using the same connection will transition to v2 custom mapping simultaneously. The process below is designed to make this as straightforward as possible.

### Step 2: Extract identifiers

The new format requires slugs, not IDs.

* **Org slugs**: Found in **Organization Settings** > **General**.
* **Group slugs**: Found in **Group Settings** > **General**.
* **Role names**: Found in **Group Settings** > **Member Roles** (for example, `developer_readonly`).

For more details, see [Slugs](README.md#slugs) and [Role normalized name](README.md#role-normalized-name).

### Step 3: Implement additional role mappings

Create new colon-delimited syntax strings based on logic from the old dash-delimited strings.

**Example transformation:**

* **Before:** `snyk-partner-plugins-admin`
* **After:** `snyk:org:partner-plugins:org_admin`

### Step 4: Configure IdP

Update your IdP to pass a multi-value attribute containing the new strings.

* **Mandatory prefix:** All strings must start with `snyk:`
* **Case sensitivity:** Role values are case-sensitive.

**Existing (legacy) role mappings:**

* **Coexistence:** New and legacy mappings can coexist in the assertion during the transition period.
* It is recommended to retain them temporarily to support rollback if needed.

## Implementation and rollout

### Step 5: Test pre-production activation

After you have set up a few new role mappings in your IdP, open a support case with the [Snyk Support team](https://support.snyk.io) for custom mapping activation. As part of this case, Support will validate the claims for compliance with the specification and subsequently activate custom mapping.

Once claims have been validated, you should complete role mapping setup in your IdP.

Your Snyk account team will perform final validation.

### Step 6: Production activation

Once validation is complete, Snyk Support will enable v2 custom mapping in your production environment.

**Best practice:** The existing SSO connection will be updated to use v2 custom role mapping.

Roles are automatically assigned upon the next user login.

**Warning:** Users without a valid mapping configured in the IdP may lose access upon login.

### Step 7: Production validation and clean up

Your team should confirm expected access levels and Organization assignment in Snyk.

It is recommended to clean up legacy mapping configuration in the IdP post-validation, since this is no longer needed.

## Syntax translation reference

### Group-level roles

The legacy format relied on strict strings or Group IDs. The v2 format targets the group scope and uses wildcards or Group slugs.

| Goal | Legacy format | v2 format |
| ---- | ------------- | --------- |
| Group Admin (all Groups in SSO) | `snyk-groupadmin` | `snyk:group:*:group_admin` |
| Group Viewer (all Groups in SSO) | `snyk-groupviewer` | `snyk:group:*:group_viewer` |
| Custom Group role | Not available | `snyk:group::custom:{custom_role}` |

{% hint style="info" %}
The v2 format replaces Group ID logic with explicit Organization-level wildcards.
{% endhint %}

### Organization-level roles

Legacy format used dashes, which made parsing custom roles difficult if the role or Organization name contained dashes. The new format uses a strict `snyk:org:{slug}:{role}` structure.

| Goal | Legacy format | New custom mapping format |
| ---- | ------------- | ------------------------- |
| Org Admin | `snyk-{orgslug}-admin` | `snyk:org:{orgslug}:org_admin` |
| Org Collaborator | `snyk-{orgslug}-collaborator` | `snyk:org:{orgslug}:org_collaborator` |
| Custom role | `snyk-{orgslug}-{custom_role}` | `snyk:org:{orgslug}:custom:{custom_role}` |

### Tenant-level roles

The new format introduces the tenant scope and uses an empty string `::` for the target, as an SSO connection is linked to a single Tenant.

| Goal | Legacy format | New custom mapping format |
| ---- | ------------- | ------------------------- |
| Tenant Admin | `snyk-tenantadmin` | `snyk:tenant::tenant_admin` |
| Tenant Viewer | `snyk-tenantviewer` | `snyk:tenant::tenant_viewer` |
| Tenant Member | `snyk-tenantmember` | `snyk:tenant::tenant_member` |
| Custom Tenant role | N/A (new in v2) | `snyk:tenant::custom:{custom_role}` |
