---
nav_context: classic
---

# Access and authentication

{% hint style="info" %}
**Feature availability**

You must have Evo enabled for your Snyk Tenant. To enable Evo, contact your Snyk account team.

Snyk supports private, single-Tenant deployments. Contact your Snyk representative for more details.
{% endhint %}

After you enable Evo, follow these steps:

1. Navigate to `https://evo.snyk.io`.
2. Log in using your existing Snyk authentication workflow.

The following table maps regions to environment URLs:

| Region              | URL                                              |
| ------------------- | ------------------------------------------------ |
| Default, SNYK-US-01 | [https://evo.snyk.io](https://evo.snyk.io)       |
| SNYK-US-02          | [https://evo.us.snyk.io](https://evo.us.snyk.io) |
| SNYK-AU-01          | [https://evo.au.snyk.io](https://evo.au.snyk.io) |
| SNYK-EU-01          | [https://evo.eu.snyk.io](https://evo.eu.snyk.io) |

### Add members

You must be a Tenant Admin to add members or modify roles.

A user joins a Tenant as a Tenant Member. This role does not include Evo access. Navigate to **Tenant** > **Members** in app.snyk.io and assign the role that matches the user's required access:

| Tenant role                   | Evo access                                                                    |
| ----------------------------- | ----------------------------------------------------------------------------- |
| Tenant Admin                  | Every asset and issue in the Tenant, and every write action                   |
| Tenant Viewer                 | Read access to the assets and issues in the Organizations the user belongs to |
| Tenant Member                 | No access to Evo                                                              |
| Tenant Member with Evo access | Every asset and issue in the Tenant, and every write action                   |
| Tenant Viewer with Evo access | Every asset and issue in the Tenant, and every write action                   |

Snyk creates two roles with Evo access in your tenant when a Tenant Admin first logs in to Evo. Each role includes the permissions of its base role and full Evo access.

To limit user access to specific Organizations, see Restrict access to specific Organizations.

For more information, visit [Manage users in a Tenant](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/tenant/manage-users-in-a-tenant).

### Restrict access to specific Organizations

Users with Organization-scoped access see only the assets and issues within the Organizations they can read.

The Tenant Viewer role is the only role limited by Organization. To scope user access:

* Assign the user the Tenant Viewer role.
* Remove Tenant Member with Evo access or Tenant Viewer with Evo access if the user holds one. These roles grant access to the entire tenant. The user retains unrestricted access until you remove the role. To change a role, visit [Manage users in a Tenant](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/tenant/manage-users-in-a-tenant).
* Add the user to each Organization that requires access to Evo data. Ensure the role includes the View Evo Organization permission. Organization Admin and Organization Collaborator include this permission.

#### Restricted actions

Organization-scoped users have read-only access. They cannot perform the following actions:

* Create, edit, or delete policies
* Save or delete reports
* Approve or reject custom discovery results
* Delete a repository added with the Snyk CLI
* Open the ADS **Settings** page (activation, deployment, and product selection)
* Start or cancel a Continuous Offensive Security scan, or create, edit, or delete a target
* Change Agent Behavior Governance policies, retention settings, or guardrail settings, or redact or delete events

Evo disables these actions in the interface and explains the reason. Evo chat applies the same limitation. It labels policy changes as restricted, keeps policy reads available, and offers to draft policy clauses for an administrator rather than attempting changes the platform rejects.

#### Coverage and limitations

Organization-scoped users do not see assets that do not belong to an Organization. Agent Supply Chain Security discovers assets on end-user machines. Continuous Offensive Security creates targets, scans, and findings. None of these carry an Organization, so these surfaces appear empty. Viewing these requires a Tenant role with full Evo access.

Policies apply to the entire tenant. Therefore, an Organization-scoped user reads every tenant policy, including policies that cover Organizations they cannot access. Evo raises issues against assets. A policy can appear with no issues against it.

A Tenant Viewer sees no data in Evo if they do not belong to an Organization, or if their role in every Organization they belong to omits the **View Evo Organization** permission.

Counts and totals reflect the access of the user viewing them. Two users on the same page can see different totals. The lower number is correct for the scoped user.

Evo caches a user's Organization access for one minute, so a membership change takes effect within a minute rather than immediately.

The Organization is the smallest unit of access. A user who can read an Organization reads every Evo asset in it. Evo does not support scoping a user to a subset of the repositories inside one Organization.
