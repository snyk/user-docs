---
description: How to set Snyk roles and access at the Organization level using role-based access control
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Authentication and access

## Set pre-defined user roles

Determine if Snyk pre-defined roles meet your requirements or if you must create custom roles. Snyk uses role-based access control (RBAC) to manage permissions at the Organization level.

## Review Organization-level roles

{% hint style="success" %}
**Key decision:** Determine if your team leads can operate with the fixed permissions of an Organization Admin or if they require a restricted custom role.
{% endhint %}

Pre-defined Organization-level roles have fixed permissions that you cannot modify.

* **Organization Admin**: Allows users to add or delete Projects, override Snyk checks, and provision users. Assign this role to team leads.
* **Organization Collaborator**: Grants standard developer access. Use this for small teams or a developer-first rollout.

## Align roles with your Organization structure

{% hint style="success" %}
**Key decision:** Decide who is responsible for provisioning new users as your Snyk footprint grows.
{% endhint %}

Your choice of roles depends on how you structured your Snyk Organizations in the previous steps.

| Structure                 | Typical role assignment                                                          |
| ------------------------- | -------------------------------------------------------------------------------- |
| **Team-based**            | Assign **Organization Admin** to the specific team lead for that Organization.   |
| **Product-based**         | Assign developers as **Collaborators** across multiple product Orgs.             |
| **SCM organization-based** | Use custom roles and the Snyk API to automate role assignment during SCM import. |

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1775661961/2._Setting_Up_Roles_and_Permissions_h69yve.mp4" %}
Setting up roles and permissions video guide
{% endembed %}
